package aper

import (
	"math"
	"reflect"

	"github.com/pkg/errors"
)

/*** General bit/byte-operation functions ***/
func (pd *PerBitData) appendSingleBit(bit byte) {
	if pd.bitsOffset == 0 {
		pd.bytes = append(pd.bytes, []byte{0x00}...)
	}
	pd.bytes[pd.byteOffset] |= bit << (8 - pd.bitsOffset - 1)
	pd.bitsOffset += 1
	pd.bitCarry()
}

func (pd *PerBitData) appendBits(bits []byte, bitLength uint64) {
	pd.bytes = append(pd.bytes, bits...)
	if pd.bitsOffset != 0 {
		bitsLeft := 8 - pd.bitsOffset
		// Shifting <bitsLeft> bits left
		for i := int(pd.byteOffset); i < len(pd.bytes)-1; i++ {
			shiftedBits := pd.bytes[i+1] >> pd.bitsOffset
			pd.bytes[i] |= shiftedBits
			pd.bytes[i+1] = pd.bytes[i+1] << bitsLeft
		}
	}
	pd.bitsOffset += uint(bitLength)
	pd.bitCarry()

	// Remove empty byte if it exists after shifting
	lastByteIdx := pd.byteOffset
	if pd.bitsOffset > 0 {
		lastByteIdx += 1
	}
	pd.bytes = pd.bytes[:lastByteIdx]
}

func (pd *PerBitData) appendBitsWithValue(value int64, bitLength uint64) error {
	var err error
	if bitLength == 0 {
		err = errors.Errorf("appendBitsWithValue error: can't append value with 0-bit")
		return err
	}
	// Check if value is over capacity
	if value >= 0 {
		if value>>bitLength != 0 {
			err = errors.Errorf("appendBitsWithValue error: bits Value is over capacity")
			return err
		}
	} else { // handled in 2's complement format
		if (value*-1)>>(bitLength-1) != 0 {
			err = errors.Errorf("appendBitsWithValue error: bits Value is over capacity")
			return err
		}
	}

	int64ToBytes := int64ToBytes(value)

	leadingBitsNum := 8*8 - bitLength
	if leadingBitsNum > 0 {
		leadingBytesNum := leadingBitsNum / 8
		leadingBitsOffset := leadingBitsNum % 8
		// remove leading bytes
		int64ToBytes = int64ToBytes[leadingBytesNum:]
		// remove remaining leading bits and shift all the following bits
		if leadingBitsOffset > 0 {
			int64ToBytes[0] = int64ToBytes[0] << byte(leadingBitsOffset)
			// Shifting <bitsLeft> bits left
			bitsLeft := 8 - leadingBitsOffset
			for i := 0; i < len(int64ToBytes)-1; i++ {
				shiftedBits := int64ToBytes[i+1] >> bitsLeft
				int64ToBytes[i] |= shiftedBits
				int64ToBytes[i+1] = int64ToBytes[i+1] << bitsLeft
			}
		}
	}
	pd.appendBits(int64ToBytes, bitLength)

	return err
}

func (pd *PerBitData) appendBitsWithUint64Value(value uint64, bitLength uint64) error {
	var err error
	if bitLength == 0 {
		err = errors.Errorf("appendBitsWithUint64Value error: can't append value with 0-bit")
		return err
	}
	// Check if value is over capacity
	if value>>bitLength != 0 {
		err = errors.Errorf("appendBitsWithUint64Value error: bits Value is over capacity")
		return err
	}

	uint64ToBytes := uint64ToBytes(value)

	leadingBitsNum := 8*8 - bitLength
	if leadingBitsNum > 0 {
		leadingBytesNum := leadingBitsNum / 8
		leadingBitsOffset := leadingBitsNum % 8
		// remove leading bytes
		uint64ToBytes = uint64ToBytes[leadingBytesNum:]
		// remove remaining leading bits and shift all the following bits
		if leadingBitsOffset > 0 {
			uint64ToBytes[0] = uint64ToBytes[0] << byte(leadingBitsOffset)
			// Shifting <bitsLeft> bits left
			bitsLeft := 8 - leadingBitsOffset
			for i := 0; i < len(uint64ToBytes)-1; i++ {
				shiftedBits := uint64ToBytes[i+1] >> bitsLeft
				uint64ToBytes[i] |= shiftedBits
				uint64ToBytes[i+1] = uint64ToBytes[i+1] << bitsLeft
			}
		}
	}
	pd.appendBits(uint64ToBytes, bitLength)

	return err
}

func (pd *PerBitData) appendAlignBits() {
	if alignBits := uint64(8-pd.bitsOffset&0x7) & 0x7; alignBits != 0 {
		pd.byteOffset += 1
	}
	pd.bitsOffset = 0
}

/*** End of general bit/byte-operation functions ***/

/*** Encoding different types for ASN.1 ***/

func (pd *PerBitData) WriteExtensible(valInExtensed bool) {
	if valInExtensed { // value is within the range of exension root
		pd.appendSingleBit(1)
	} else { // value is in extensed additions
		pd.appendSingleBit(0)
	}
}

// 有約束、半約束、無約束、常見小自然數
func (pd *PerBitData) WriteWholeNum(isNormallySmallNonNegativeNum bool,
	value int64, lbPtr *int64, ubPtr *int64,
) error {
	var err error

	if isNormallySmallNonNegativeNum {
		// IV. 常见小自然数编码
		//  This procedure is used when encoding a non-negative whole number that is expected to be small,
		//  but whose size is potentially unlimited due to the presence of an extension marker
		// 这种情况经常出现在对表征SEQUENCE、SET类型可选成员的Bitmap长度进行编码时；
		// 或者CHOICE类型序号编码时。这种长度相当小，但是却没有一种限定。
		// | 当0≤n≤63时，n以6个比特编码，并且在前面增加一个 0-bit（八位组不对齐）
		// | 当64≤n时，n以半约束数方式编码，下边界为0，并且在前面增加一个 1-bit
		if value < 0 {
			err = errors.Errorf("An negative value is invalid to write normally small non-negative whole number")
			return err
		} else if value <= 63 {
			pd.appendSingleBit(0)
			// represent value in 6 bits
			err = pd.appendBitsWithValue(value, 6)
		} else {
			pd.appendSingleBit(1)
			var tmpLb int64 = 0
			err = pd.WriteWholeNum(false, value, &tmpLb, nil)
		}
		return err
	}
	if lbPtr != nil && ubPtr != nil {
		// I. 有约束数编码
		// 有约束指值域的上、下边界都有限。设d为取值范围的大小（上下边界的差值 + 1）；
		// 在对齐方式下：
		// n- bmin encoded into octects as a non-negative-binary-integer
		// 当2≤d≤255，n- bmin的编码占用log2d个比特。这些比特不进行八位组对齐，不编码L；
		// (If "range" has the value 1, then the result of the encoding shall be an empty bit-field (no bits))
		// | 当d=256，n- bmin的编码占用一个八位组，不编码L；
		// | 当257≤d≤65,536，n- bmin的编码占用两个八位组，不编码L；
		// | 当65,537≤d，n- bmin的编码占用log256d个八位组，并且在前面增加L的编码
		d := *ubPtr - *lbPtr + 1
		encodeValue := value - *lbPtr
		if d < 0 {
			err = errors.Errorf("WriteWholeNum error: Value range is negative")
			return err
		} else if d == 1 {
			return err
		} else if d <= 255 {
			numBits := getRequiredNumBits(d-1, false) // num of bits required to represent d numbers = log2(d-1)
			err = pd.appendBitsWithValue(encodeValue, numBits)
			return err
		} else if d == 256 {
			pd.appendAlignBits()
			pd.appendBits([]byte{byte(encodeValue)}, 8*1)
		} else if d <= 65536 {
			pd.appendAlignBits()
			firstByte := byte(encodeValue >> 8)
			secondByte := byte(encodeValue & 255)
			pd.appendBits([]byte{firstByte, secondByte}, 8*2)
			return err
		} else {
			// encode length
			// (12.2.6: 用 constrained length determinate 編碼，lb=1，ub=log256d)。
			dNumBytes := getRequiredNumBytes(d-1, false)
			encodeValueNumBytes := getRequiredNumBytes(encodeValue, false)
			lmin := uint64(1)
			if err = pd.WriteLength(false, encodeValueNumBytes, &lmin, &dNumBytes); err != nil {
				return err
			}
			// encode value
			pd.appendAlignBits()
			err = pd.appendBitsWithValue(encodeValue, 8*encodeValueNumBytes)
			return err
		}
	} else if lbPtr == nil {
		// III. 无约束数编码
		// 无约束指值域没有下边界（即使存在上边界）。
		// n encoded into octects as a 2’s-complement-binary-integer
		// n 的编码占用log256d个八位组，并且在前面增加L的编码。
		numBytes := getRequiredNumBytes(value, true)
		if err = pd.WriteLength(false, numBytes, nil, nil); err != nil {
			return err
		}
		err = pd.appendBitsWithValue(value, 8*numBytes)
		return err
	} else if ubPtr == nil {
		// II. 半约束数编码
		// 半约束指值域没有上边界（上边界为+∞）。
		// n-bmin encoded into octects as a non-negative-binary-integer
		// n-bmin的编码占用log256d个八位组，并且在前面增加L的编码。
		numBytes := getRequiredNumBytes(value-*lbPtr, false)
		lb := uint64(*lbPtr)
		if err = pd.WriteLength(false, numBytes, &lb, nil); err != nil {
			return err
		}
		err = pd.appendBitsWithValue(value, 8*numBytes)
	} else {
		err = errors.Errorf("Write whole number unexpected error.")
	}

	return err
}

func (pd *PerBitData) WriteUint64WholeNum(isNormallySmallNonNegativeNum bool,
	value uint64, lbPtr *uint64, ubPtr *uint64,
) error {
	var err error

	if isNormallySmallNonNegativeNum {
		// IV. 常见小自然数编码
		//  This procedure is used when encoding a non-negative whole number that is expected to be small,
		//  but whose size is potentially unlimited due to the presence of an extension marker
		// 这种情况经常出现在对表征SEQUENCE、SET类型可选成员的Bitmap长度进行编码时；
		// 或者CHOICE类型序号编码时。这种长度相当小，但是却没有一种限定。
		// | 当0≤n≤63时，n以6个比特编码，并且在前面增加一个 0-bit（八位组不对齐）
		// | 当64≤n时，n以半约束数方式编码，下边界为0，并且在前面增加一个 1-bit
		if value <= 63 {
			pd.appendSingleBit(0)
			// represent value in 6 bits
			err = pd.appendBitsWithUint64Value(value, 6)
		} else {
			pd.appendSingleBit(1)
			var tmpLb uint64 = 0
			err = pd.WriteUint64WholeNum(false, value, &tmpLb, nil)
		}
		return err
	}
	if lbPtr != nil && ubPtr != nil {
		// I. 有约束数编码
		// 有约束指值域的上、下边界都有限。设d为取值范围的大小（上下边界的差值 + 1）；
		// 在对齐方式下：
		// n- bmin encoded into octects as a non-negative-binary-integer
		// 当2≤d≤255，n- bmin的编码占用log2d个比特。这些比特不进行八位组对齐，不编码L；
		// (If "range" has the value 1, then the result of the encoding shall be an empty bit-field (no bits))
		// | 当d=256，n- bmin的编码占用一个八位组，不编码L；
		// | 当257≤d≤65,536，n- bmin的编码占用两个八位组，不编码L；
		// | 当65,537≤d，n- bmin的编码占用log256d个八位组，并且在前面增加L的编码

		// Special Case: ub = uint64 上界， lb = 0
		// since d := *ubPtr - *lbPtr + 1 will overflow, separate this from other cases
		if *ubPtr == 18446744073709551615 && *lbPtr == 0 {
			// case: d > 65536
			// encode length
			// (12.2.6: 用 constrained length determinate 編碼，lb=1，ub=log256d)。
			dNumBytes := uint64(8)
			encodeValue := value
			encodeValueNumBytes := getRequiredNumBytesForUint64(encodeValue)
			lmin := uint64(1)
			if err = pd.WriteLength(false, encodeValueNumBytes, &lmin, &dNumBytes); err != nil {
				return err
			}
			// encode value
			pd.appendAlignBits()
			err = pd.appendBitsWithUint64Value(encodeValue, 8*encodeValueNumBytes)
			return err
		}

		d := *ubPtr - *lbPtr + 1
		encodeValue := value - *lbPtr
		if d == 0 {
			err = errors.Errorf("WriteUint64WholeNum error: value range is zero")
			return err
		} else if d == 1 {
			return err
		} else if d <= 255 {
			numBits := getRequiredNumBitsForUint64(d - 1) // num of bits required to represent d numbers = log2(d-1)
			err = pd.appendBitsWithUint64Value(encodeValue, numBits)
			return err
		} else if d == 256 {
			pd.appendAlignBits()
			pd.appendBits([]byte{byte(encodeValue)}, 8*1)
		} else if d <= 65536 {
			pd.appendAlignBits()
			firstByte := byte(encodeValue >> 8)
			secondByte := byte(encodeValue & 255)
			pd.appendBits([]byte{firstByte, secondByte}, 8*2)
			return err
		} else {
			// encode length
			// (12.2.6: 用 constrained length determinate 編碼，lb=1，ub=log256d)。
			dNumBytes := getRequiredNumBitsForUint64(d - 1)
			encodeValueNumBytes := getRequiredNumBytesForUint64(encodeValue)
			lmin := uint64(1)
			if err = pd.WriteLength(false, encodeValueNumBytes, &lmin, &dNumBytes); err != nil {
				return err
			}
			// encode value
			pd.appendAlignBits()
			err = pd.appendBitsWithUint64Value(encodeValue, 8*encodeValueNumBytes)
			return err
		}
	} else if lbPtr == nil {
		// III. 无约束数编码
		// 无约束指值域没有下边界（即使存在上边界）。
		// n encoded into octects as a 2’s-complement-binary-integer
		// n 的编码占用log256d个八位组，并且在前面增加L的编码。
		numBytes := getRequiredNumBytesForUint64(value)
		if err = pd.WriteLength(false, numBytes, nil, nil); err != nil {
			return err
		}
		err = pd.appendBitsWithUint64Value(value, 8*numBytes)
		return err
	} else if ubPtr == nil {
		// II. 半约束数编码
		// 半约束指值域没有上边界（上边界为+∞）。
		// n-bmin encoded into octects as a non-negative-binary-integer
		// n-bmin的编码占用log256d个八位组，并且在前面增加L的编码。
		numBytes := getRequiredNumBytesForUint64(value - *lbPtr)
		lb := *lbPtr
		if err = pd.WriteLength(false, numBytes, &lb, nil); err != nil {
			return err
		}
		err = pd.appendBitsWithUint64Value(value, 8*numBytes)
	} else {
		err = errors.Errorf("Write whole number unexpected error.")
	}

	return err
}

// INTEGER
// 1. 如果有 extension marker 在約束裡，最前面要加上一個 preamble bit
// 1-1. 如果要編碼的值在 extension root 裡，preamble bit = 0，
//
//	並用 2. 的規則編碼。
//
// 1-2. 反之，如果要編碼的值不在 extension root 裡，preamble bit = 1，
//
//	並用無約束的方式對值編碼，編碼結束。
//
// 2. 如果沒有 extension marker 的編碼規則：
// 2-1. 如果PER可見約束限制 Integer 是單一的值 (ub==lb)，不編碼
// 2-2. 其他根據PER可見約束限制用 whole number 的編碼方式編碼
func (pd *PerBitData) WriteInteger(value int64, valueExtensible bool,
	lbPtr *int64, ubPtr *int64,
) error {
	var err error
	var lb, ub int64
	if ubPtr != nil {
		ub = *ubPtr
	}
	if lbPtr != nil {
		lb = *lbPtr
	}

	// Case 1.
	// Decide the value constraint (-1 means no valid size constraint)
	var validValueConstraint int64 = -1
	if lbPtr != nil {
		if ubPtr != nil {
			if lb <= value && value <= ub {
				validValueConstraint = ub - lb + 1
			}
		}
	}

	// Write Preamble Bit if is value extensible
	if valueExtensible {
		if validValueConstraint != -1 { // Case 1-1.
			pd.WriteExtensible(false) // Continue to use Case 2 to encode value
		} else {
			pd.WriteExtensible(true)                       // Case 1-2.
			err = pd.WriteWholeNum(false, value, nil, nil) // Encode with unconstrained whold number
			return err
		}
	} else {
		if ubPtr != nil && value > ub {
			return errors.Errorf("WriteInteger error: integer value is over upperbound")
		}
		if lbPtr != nil && value < lb {
			return errors.Errorf("WriteInteger error: integer value is under lowerbound")
		}
	}

	// Case 2-1.
	if validValueConstraint == 1 {
		return err
	}
	// Case 2-2.
	err = pd.WriteWholeNum(false, value, lbPtr, ubPtr)

	return err
}

func (pd *PerBitData) WriteUint64Integer(value uint64, valueExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	var err error
	var lb, ub uint64
	if ubPtr != nil {
		ub = *ubPtr
	}
	if lbPtr != nil {
		lb = *lbPtr
	}

	// Case 1.
	// Decide the value constraint (-1 means no valid size constraint)
	isValueConstraintValid := false
	if lbPtr != nil {
		if ubPtr != nil {
			if lb <= value && value <= ub {
				isValueConstraintValid = true
			}
		}
	}

	// Write Preamble Bit if is value extensible
	if valueExtensible {
		if isValueConstraintValid { // Case 1-1.
			pd.WriteExtensible(false) // Continue to use Case 2 to encode value
		} else {
			pd.WriteExtensible(true)                             // Case 1-2.
			err = pd.WriteUint64WholeNum(false, value, nil, nil) // Encode with unconstrained whold number
		}
	} else {
		if ubPtr != nil && value > ub {
			return errors.Errorf("WriteInteger: integer value is over upperbound")
		}
		if lbPtr != nil && value < lb {
			return errors.Errorf("WriteInteger: integer value is under lowerbound")
		}
	}

	// Case 2-1.
	if isValueConstraintValid && *ubPtr == *lbPtr {
		return err
	}
	// Case 2-2.
	err = pd.WriteUint64WholeNum(false, value, lbPtr, ubPtr)

	return err
}

// 長度編碼
//  1. 当l是一个bitmap的长度，l-1作为常见小自然数编码；
//     (Normally small lengths are only used to indicate the length of the bitmap that prefixes the extension
//     addition values of a "set" or "sequence type". )
//  2. 当lmax≤65,535，l作为有约束的数编码（约束为(lmin…lmax)）；
//  3. 当65,535≤lmax，或者lmax是无穷大：
//
// 3-1. 当l≤127，l以一个八位组编码（八位组对齐），最高比特位为0；
// 3-2. 当128≤l≤16,383，l以两个八位组编码（八位组对齐），最高两个比特位为10；
// 3-3. 当16,384≤l，見 writeLargeXXXWithFragmentation()
func (pd *PerBitData) WriteLength(isForBitMapLen bool, l uint64, lbPtr *uint64, ubPtr *uint64) error {
	var err error
	if isForBitMapLen { // case 1: for set or sequence type with extension marker
		err = pd.WriteWholeNum(true, int64(l-1), nil, nil)
		return err
	} else {
		if ubPtr != nil && *ubPtr <= 65535 { // case 2.
			lb := int64(*lbPtr)
			ub := int64(*ubPtr)
			err = pd.WriteWholeNum(false, int64(l), &lb, &ub)
		} else { // case 3. (ubPtr == nil || *ubPtr > 65535)
			if l <= 127 { // case 3-1.
				pd.appendAlignBits()
				pd.appendBits([]byte{byte(l)}, 8*1)
			} else if l <= 16383 { // case 3-2.
				pd.appendAlignBits()
				firstByte := byte(l >> 8)
				firstByte |= 0x80 // first two bits are 10
				secondByte := byte(l & 255)
				pd.appendBits([]byte{firstByte, secondByte}, 8*2)
			} else { // case 3-3 (l >= 16384)
				pd.appendAlignBits()
				f := l / 16384
				pd.appendBits([]byte{byte(f) | 0xc0}, 8)
			}
		}
	}

	return err
}

// Case 3-3.
// 当16,384≤l，整个编码以f*16K为单位分割（f取值为1，2，3或者4）。
// 除最后的片段外，其余每段，长度為 f 值以一个八位组编码 (用無約束方式)，最高两个比特位为11。
// 如果编码恰好时16K的整倍数，则在最后补充一个全空的八位组；否则最后一个片段按照前两条进行编码。
func (pd *PerBitData) WriteLargeBitStringWithFragment(l uint64, bytes []byte,
	lbPtr *uint64, ubPtr *uint64,
) error { // l in bit, lUnit = 1; l in byte, lUnit = 8
	var err error
	for {
		// Decide fragmentation size and write fragmentation length
		var fragmentLen uint64
		pd.appendAlignBits()
		if l > 65536 {
			fragmentLen = 65536
			err = pd.WriteLength(false, fragmentLen, nil, nil)
		} else if l >= 16384 {
			fragmentLen = l
			err = pd.WriteLength(false, fragmentLen, nil, nil)
			fragmentLen = (l / 16384) * 16384
		} else {
			fragmentLen = l
			err = pd.WriteLength(false, fragmentLen, lbPtr, ubPtr)
		}
		if err != nil {
			return errors.Wrap(err, "write large bitString failed")
		}

		// Write fragmentation data
		pd.appendAlignBits()
		numBytes := uint64(math.Ceil(float64(fragmentLen) / 8))
		pd.appendBits(bytes[:numBytes], fragmentLen)
		bytes = bytes[numBytes:]

		l -= fragmentLen
		if l <= 0 {
			if fragmentLen%16384 == 0 { // TO-BE-TEST
				err = pd.appendBitsWithValue(0, 8)
			}
			pd.appendAlignBits()
			break
		}
	}
	return err
}

func (pd *PerBitData) WriteLargeOctetStringWithFragment(l uint64, bytes []byte,
	lbPtr *uint64, ubPtr *uint64,
) error { // l in bit, lUnit = 1; l in byte, lUnit = 8
	var err error
	for {
		// Decide fragmentation size and write fragmentation length
		var fragmentLen uint64
		pd.appendAlignBits()
		if l > 65536 {
			fragmentLen = 65536
			err = pd.WriteLength(false, fragmentLen, nil, nil)
		} else if l >= 16384 {
			fragmentLen = l
			err = pd.WriteLength(false, fragmentLen, nil, nil)
			fragmentLen = (l / 16384) * 16384
		} else {
			fragmentLen = l
			err = pd.WriteLength(false, fragmentLen, lbPtr, ubPtr)
		}
		if err != nil {
			return errors.Wrap(err, "write large octet string failed")
		}

		// Write fragmentation data
		pd.appendAlignBits()
		numBytes := fragmentLen
		pd.appendBits(bytes[:numBytes], fragmentLen*8)
		bytes = bytes[numBytes:]

		l -= fragmentLen
		if l <= 0 {
			if fragmentLen%16384 == 0 { // TO-BE-TEST
				err = pd.appendBitsWithValue(0, 8)
			}
			pd.appendAlignBits()
			break
		}
	}
	return err
}

// BIT STRING
// 0. 如果 lb 沒有設定，則 lb = 0
// 1. 如果有 extension marker 在約束裡，最前面要加上一個 preamble bit
// 1-1. 如果要編碼的 bitString 長度在 extension root 裡，preamble bit = 0，
//
//	並用 2. 的規則對長度和 bitstring 編碼。
//
// 1-2. 反之，如果要編碼的 bitString 長度不在 extension root 裡，preamble bit = 1，
//
//	用半約束的方式編碼長度，並加上 bitstring 的值，編碼結束。
//
// 2. 如果沒有 extension marker 的編碼規則：
// 2-1. 如果 bitstring 的長度限制為 zero length (ub=0)，則不編碼
// 2-2. 当lmin＝lmax≤16比特，不发送长度，直接编码（不是八位组对齐的）
// 2-3. 当17≤lmin＝lmax≤65,536比特，不发送长度，直接编码（在八位组对齐方式下是八位组对齐的）；
// 2-4. 如果 2-1.~2-3. 都不符合，用加上長度編碼的方式編碼
func (pd *PerBitData) WriteBitString(bs BitString, sizeExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	var err error
	var lb, ub uint64

	// ITU X.691 15.4
	// If there is no finite maximum we say that "ub" is unset.
	// If there is no constraint on the minimum, then "lb" has the value zero.
	if ubPtr != nil {
		ub = *ubPtr
	}
	if lbPtr != nil {
		lb = *lbPtr
	} else { // Case 0.
		lb = 0
	}

	// Decide the size constraint (-1 means no valid size constraint)
	var validSizeConstraint int64 = -1
	if lbPtr != nil {
		if ubPtr != nil {
			if lb <= bs.BitLength && bs.BitLength <= ub {
				validSizeConstraint = int64(ub - lb + 1)
			}
		}
	}

	// Guarantee that the bits out of bitLength are 0s
	bsBytes := (bs.BitLength + 7) >> 3         // (bs.BitLength + 7)/8
	bsLastByteOffset := (8 - bs.BitLength&0x7) // 8 - (bs.BitLength%8)
	if bsLastByteOffset != 8 {
		bs.Bytes[bsBytes-1] &= (0xff << bsLastByteOffset)
	}

	// Write Preamble Bit if is size extensible
	if sizeExtensible {
		if validSizeConstraint == -1 { // Case 1-2.
			pd.WriteExtensible(true)
			err = pd.WriteLength(false, bs.BitLength, lbPtr, nil)
			pd.appendAlignBits()
			pd.appendBits(bs.Bytes, bs.BitLength)
			return err
		} else { // Case 1-1.
			pd.WriteExtensible(false) // Continue to encode using Case 2.
		}
	} else {
		if ubPtr != nil && bs.BitLength > ub {
			return errors.Errorf("WriteBitString error: bitString Length is over upperbound")
		}
	}

	// Encode Bitstring
	// case 2.
	if ubPtr != nil && *ubPtr == 0 { // case 2-1.	// TO-BE-TEST
		return err
	}

	if validSizeConstraint == 1 { // ub == lb
		if bs.BitLength != ub {
			err = errors.Errorf("WriteBitString error: "+
				"bitString Length(%d) is not match fix-sized : %d", bs.BitLength, ub)
		}
		if ub <= 16 { // case 2-2
			pd.appendBits(bs.Bytes, bs.BitLength)
			return err
		} else if ub <= 65536 { // case 2-3
			pd.appendAlignBits()
			pd.appendBits(bs.Bytes, bs.BitLength)
			return err
		}
	}

	// case 2-4
	if bs.BitLength <= 16383 {
		err = pd.WriteLength(false, bs.BitLength, lbPtr, ubPtr)
		if bs.BitLength == 0 { // empty bit string: don't encode
			return nil
		}
		pd.appendAlignBits()
		pd.appendBits(bs.Bytes, bs.BitLength)
	} else { // apply fragmentation
		err = pd.WriteLargeBitStringWithFragment(bs.BitLength, bs.Bytes, lbPtr, ubPtr)
	}

	return err
}

// OCTET STRING
// 与BIT STRING规则相同，不过长度域L表征的是八位组的个数而不是比特数。
func (pd *PerBitData) WriteOctetString(os []byte, sizeExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	var err error
	var lb, ub uint64
	byteLen := uint64(len(os))

	// ITU X.691 15.4
	// If there is no finite maximum we say that "ub" is unset.
	// If there is no constraint on the minimum, then "lb" has the value zero.
	if ubPtr != nil {
		ub = *ubPtr
	}
	if lbPtr != nil {
		lb = *lbPtr
	} else { // Case 0.
		lb = 0
	}

	// Decide the size constraint (-1 means no valid size constraint)
	var validSizeConstraint int64 = -1
	if lbPtr != nil {
		if ubPtr != nil {
			if lb <= byteLen && byteLen <= ub {
				validSizeConstraint = int64(ub - lb + 1)
			}
		}
	}

	// Write Preamble Bit if is size extensible
	if sizeExtensible {
		if validSizeConstraint == -1 { // Case 1-2.
			pd.WriteExtensible(true)
			err = pd.WriteLength(false, byteLen, lbPtr, nil)
			pd.appendAlignBits()
			pd.appendBits(os, byteLen*8)
			return err
		} else { // Case 1-1.
			pd.WriteExtensible(false) // Continue to encode using Case 2.
		}
	} else {
		if ubPtr != nil && byteLen > ub {
			return errors.Errorf("WriteOctetString error: Octet String Length is over upperbound")
		}
	}

	// Encode Octet String
	// case 2.
	if ubPtr != nil && *ubPtr == 0 { // case 2-1.	// TO-BE-TEST
		return err
	}

	if validSizeConstraint == 1 { // ub == lb
		if byteLen != ub {
			err = errors.Errorf("WriteOctetString error: "+
				"Octet String Length(%d) is not match fix-sized : %d", byteLen, ub)
		}
		if ub <= 2 { // case 2-2
			pd.appendBits(os, byteLen*8)
			return err
		} else if ub <= 65536 { // case 2-3
			pd.appendAlignBits()
			pd.appendBits(os, byteLen*8)
			return err
		}
	}

	// case 2-4
	if byteLen <= 16383 {
		err = pd.WriteLength(false, byteLen, lbPtr, ubPtr)
		if byteLen == 0 { // empty octect string: don't encode
			return nil
		}
		pd.appendAlignBits()
		pd.appendBits(os, byteLen*8)
	} else { // apply fragmentation
		err = pd.WriteLargeOctetStringWithFragment(byteLen, os, lbPtr, ubPtr)
	}
	return err
}

// Characters: PrintableString, VisibleString, UTF8String
// PER-visible constraints only apply to known-multiplier character string types.
// For other restricted character string types
// "aub" will be unset and "alb" will be zero. known-multiplier character string types: NumericString,
// PrintableString, VisibleString (ISO646String), IA5String, BMPString, and UniversalString
func (pd *PerBitData) WritePrintableString(s PrintableString, sizeExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	err := pd.WriteOctetString([]byte(string(s)), sizeExtensible, lbPtr, ubPtr)
	return err
}

func (pd *PerBitData) WriteVisibleString(s VisibleString, sizeExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	err := pd.WriteOctetString([]byte(string(s)), sizeExtensible, lbPtr, ubPtr)
	return err
}

func (pd *PerBitData) WriteUTF8String(s UTF8String, sizeExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	// constraint not applied for UTF8String
	if ubPtr != nil {
		ubPtr = nil
	}
	if lbPtr == nil || *lbPtr != 0 {
		lb := uint64(0)
		lbPtr = &lb
	}
	if sizeExtensible {
		sizeExtensible = false
	}

	err := pd.WriteOctetString([]byte(string(s)), sizeExtensible, lbPtr, ubPtr)
	return err
}

// ENUMERATED
//  1. 如果该ENUMERATED类型不是可扩展的，则先按照数值大小做升序排列，然后以0为起点，步长为1给每个成员编上序号。
//     对该类型的值编码时，只将序号以值域约束(0…Indexmax)编码。如：
//     v ENUMERATED {orange(56), green(-2), red(2476)} ::= orange
//     的PER编码为“01”（因为此时为Indexmax 2，需要两个比特）。
//  2. 如果该ENUMERATED类型是可扩展的，那么要在编码前增加一个前导（Preamble）比特，
//     同时对扩展部分的成员重新进行编号，擴展後的第一個 item 編號仍为0，步长为1。
//
// 2-1. 当值在扩展的根部分时，前導比特为0，编码时对值在根部分的情况，就和该类型是不可扩展时一样；如：
//
//	 v1 ENUMERATED {orange(56), green(-2), red(2476), …, yellow} ::= orange
//	的PER编码仍然为“01”；
//
// 2-2. 值在扩展部分的情况，对序号按照自然数方式编码。
//
//			v2 ENUMERATED {orange(56), green(-2), red(2476), …, yellow, purple} ::= yellow
//	   的PER编码则为“10000000”。
func (pd *PerBitData) WriteEnumerated(enum Enumerated, valueExtensible bool,
	lbPtr *int64, ubPtr *int64,
) error {
	var err error

	val := int64(enum)

	// Check lb and ub
	if lbPtr == nil || ubPtr == nil {
		return errors.Errorf("WriteEnumerated error: ENUMERATED value constraint is error")
	}

	lb, ub := *lbPtr, *ubPtr

	if lb < 0 || lb > ub {
		return errors.Errorf("WriteEnumerated error: ENUMERATED value constraint is error")
	}

	if !valueExtensible { // Case 1
		if lb <= val && val <= ub {
			err = pd.WriteWholeNum(false, val, &lb, &ub)
		} else {
			return errors.Errorf("WriteEnumerated error: ENUMERATED value is out of valid range")
		}
	} else {
		if lb <= val && val <= ub { // Case 2-1.
			pd.WriteExtensible(false)
			err = pd.WriteWholeNum(false, val, &lb, &ub)
		} else if ub < val { // Case 2-2.
			pd.WriteExtensible(true)
			err = pd.WriteWholeNum(true, val-ub-1, nil, nil)
		}
	}

	return err
}

// SEQUENCE OF
// 显式PER限制仅仅对SEQUENCE OF类型的成员个数有效。
// 1. 如果有 extension marker 在約束裡，最前面要加上一個 preamble bit
// 1-1. 如果要 SEQUENCE OF 的 components 數量 在 extension root 裡，preamble bit = 0，
//
//	並用 2. 的規則對 components 數量和 compoenents 編碼。
//
// 1-2. 反之，如果 components 數量不在 extension root 裡，preamble bit = 1，
//
//	用半約束的方式編碼數量，並加上 components 的值，編碼結束。
//
// 2. 如果沒有 extension marker 的編碼規則：
// 2-1. 如果 components 的數量為 fixed (ub==lb)，且 ub < 64K，則不編碼 L 直接編碼 components 的值
// 2-2. 否則用加上 components 數量的長度編碼的方式編碼
// 2-3. The fragmentation procedures may apply after 16K, 32K, 48K, or 64K components.
func (pd *PerBitData) WriteSequenceOfPreambleBitMap(numComponents uint64, sizeExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) error {
	var err error
	var lb, ub uint64

	// ITU X.691 20.2
	// If there is no finite maximum or "ub" is greater than or equal to 64K we say that "ub" is unset.
	// If there is no constraint on the minimum, then "lb" has the value zero.
	if ubPtr != nil && *ubPtr < 65536 {
		ub = *ubPtr
	} else {
		ubPtr = nil
	}
	if lbPtr != nil {
		lb = *lbPtr
	} else {
		lb = 0
	}

	// Decide the size constraint (-1 means no valid size constraint)
	var validSizeConstraint int64 = -1
	if lbPtr != nil {
		if ubPtr != nil {
			if lb <= numComponents && numComponents <= ub {
				validSizeConstraint = int64(ub - lb + 1)
			}
		}
	}

	// Write Preamble Bit if is size extensible
	if sizeExtensible {
		if validSizeConstraint == -1 { // Case 1-2.
			pd.WriteExtensible(true)
			err = pd.WriteLength(false, numComponents, lbPtr, nil)
			pd.appendAlignBits()
			// "Write preamble bitmap for sequenco of"
			// "Please continue with using other pd.Write functions according to "
			// "the sequence of components type to complete the encoding process."
			return err
		} else { // Case 1-1.
			pd.WriteExtensible(false) // Continue to encode using Case 2.
		}
	} else {
		if ubPtr != nil && numComponents > ub {
			return errors.Errorf("WriteSequenceOfPreambleBitMap error: Sequence of size is over upperbound")
		}
	}

	// Encode Length determinant if required
	if validSizeConstraint == 1 && ub < 65536 { // Case 2-1.
		// "Write preamble bitmap for sequenco of"
		// "Please continue with using other pd.Write functions according to "
		// "the sequence of components type to complete the encoding process."
		return err
	} else {
		if numComponents <= 16383 { // Case 2-2.
			err = pd.WriteLength(false, numComponents, lbPtr, ubPtr)
			// "Write preamble bitmap for sequenco of"
			// "Please continue with using other pd.Write functions according to "
			// "the sequence of components type to complete the encoding process."
			return err
		} else { // Case 2-3.
			// TODO: Fragmentation is not supported here since length interleaving with data.
			// Encoding requires knowledge about sequence of type (e.g. using reflect)
			// to encode fragmented components preceded with length determinant of each fragment.
			return errors.Errorf("WriteSequenceOfPreambleBitMap error: " +
				"Unsupport sequence of size larger than 16383. Require fragmentation implementation.")
		}
	}
}

// CHOICE
//  0. 如果该Choice类型不是可扩展的，则先按照数值大小做升序排列，然后以0为起点，步长为1给每个成员编上序号。
//     如果 Choice 类型是可扩展的，且有擴展附加選項，也從 0 開始幫這些擴展附加選項編號
//  1. 如果 extension root 只有一個選項，且該選項被選擇了 (一定同時是 2-1.)，不用編碼索引值
//  2. 如果CHOICE類型是可扩展的，那么要在编码前增加一个前导（Preamble）比特，選擇屬於擴展附加選項時設為 1，否則為 0。
//
// 2-1. 如果不可擴展或当值在扩展的根部分时，選項的索引值作為 Integer 編碼 (約束 (0..根部索引最大值))，然後編碼選項的值
// 2-2. 如果選項在扩展部分的情况，索引作為常見小自然數編碼，lb 設為 0，選項作為 open type 編碼
func (pd *PerBitData) WriteChoicePreambleBitMap(index int64, valueExtensible bool, ubPtr *int64) error {
	var err error

	// Check ub
	if ubPtr == nil {
		return errors.Errorf("WriteChoicePreambleBitMap error: The upper bound of CHIOCE is missing")
	}

	lb, ub := int64(0), *ubPtr

	if ub < 0 {
		return errors.Errorf("WriteChoicePreambleBitMap error: The upper bound of CHIOCE is negative")
	}

	// Case 2.
	if valueExtensible {
		if ub < index { // Case 2-2.
			pd.WriteExtensible(true)
			err = pd.WriteWholeNum(true, index, &lb, nil)
			if err != nil {
				return errors.Wrap(err, "write choice failed")
			}
			// "Write preamble bitmap for choice"
			// "Please continue with using other pd.Write functions according to "
			// "the sequence of components type to complete the encoding process."
			return nil
		} else {
			pd.WriteExtensible(false)
			// Case 2-1.
		}
	} // else: Case 2-1

	// Case 2-1
	// Case 1 must happens along with Case 2-1 (selected choice is in root)
	if ub == 0 && index == 0 {
		// "Only one choice and the choice is selected, no index encoding required"
		return nil
	}

	if lb <= index && index <= ub { // Case 2-1
		err = pd.WriteInteger(index, false, &lb, &ub)
		if err != nil {
			return errors.Wrap(err, "write choice failed")
		}
		// "Write preamble bitmap for choice"
		// "Please continue with using other pd.Write functions according to "
		// "the sequence of components type to complete the encoding process."
		return nil
	} else {
		return errors.Errorf("WriteChoicePreambleBitMap error: Choice index is out of valid range")
	}
}

// SEQUENCE
//  0. 如果SEQUENCE类型是可扩展的，则在编码的头部加上一个比特的bit-field，
//     (Unsupported) 如果SEQUENCE的取值中有属于扩展附加部分的成员，则该比特等于1，否则等于0。
//  1. 如果SEQUENCE的定义中在扩展根部（extension root）有"n"个成员被置为OPTIONAL或DEFAULT，
//     则在编码头部再添加"n"个比特的bit-field，该bit-field从第一个bit开始，依次指示被标记为OPTIONAL或DEFAULT的成员是否出现。
//     - 如果“n”小于64K，则这个 bit-field 应该直接添到码流中。
//     - (Unsupported) 如果"n"大于等于64K，按照前面提到的处理方法把“n”个bit的bit-field分段并添加到域序列中，
//     前面的长度字段L就作为一个有约束的整数编码，而约束的上限和下限都等于n。
func (pd *PerBitData) WriteSequencePreambleBitMap(optionalPresentFlag []bool, valueExtensible bool) error {
	// TO-BE-TEST

	// Case 0.: write extensible bits
	if valueExtensible {
		pd.WriteExtensible(false)
		// TODO: Support extensible sequence
		// "Extensible Sequence is not supported. Extensible bit is always 0.
	}

	// Case 1: write optional bit-field
	numOptionalFields := len(optionalPresentFlag)
	if numOptionalFields == 0 {
		// No optional field for sequence
		return nil
	} else if numOptionalFields > 0 && numOptionalFields < 65536 {
		var optionalBits uint64
		for _, flag := range optionalPresentFlag {
			optionalBits <<= 1
			if flag {
				optionalBits++
			}
		}
		err := pd.appendBitsWithValue(int64(optionalBits), uint64(numOptionalFields))
		return err
	} else { // numOptionalFields > 64K (Unsupported)
		return errors.Errorf("WriteSequencePreambleBitMap error: " +
			"Number of optional fields more than 64K is not supported.")
	}
}

// OPEN TYPE
// 0. 真正類型的內容要先被 encode 成一個長度為 "n" 的 octet string
// 1. 在 0. 的 octet string 前面用無約束編碼的方式編碼長度 (n，in units of octets)
func (pd *PerBitData) WriteOpenType(openTypeBytes []byte) error {
	var err error
	byteLen := uint64(len(openTypeBytes))
	if byteLen <= 16383 {
		err = pd.WriteLength(false, byteLen, nil, nil)
		if byteLen == 0 { // empty octect string: don't encode
			return nil
		}
		pd.appendAlignBits()
		pd.appendBits(openTypeBytes, byteLen*8)
	} else {
		err = pd.WriteLargeOctetStringWithFragment(byteLen, openTypeBytes, nil, nil)
	}
	return err
}

// BOOLEAN
func (pd *PerBitData) WriteBool(val bool) error {
	var err error
	if val {
		// Encoded BOOLEAN Value : 0x1
		err = pd.appendBitsWithValue(1, 1)
	} else {
		// Encoded BOOLEAN Value : 0x0
		err = pd.appendBitsWithValue(0, 1)
	}
	return err
}

/*** End of encoding different types for ASN.1 ***/

/*** Legacy fuction for encoding an empty interface (using reflect) ***/
func (pd *PerBitData) legacyMakeField(v reflect.Value, params fieldParameters) error {
	if !v.IsValid() {
		return errors.Errorf("legacyMakeField error: cannot marshal nil value")
	}
	// If the field is an interface{} then recurse into it.
	if v.Kind() == reflect.Interface && v.Type().NumMethod() == 0 {
		return pd.legacyMakeField(v.Elem(), params)
	}
	if v.Kind() == reflect.Ptr {
		return pd.legacyMakeField(v.Elem(), params)
	}
	fieldType := v.Type()

	// We deal with the structures defined in this package first.
	switch fieldType {
	case BitStringType:
		bs := BitString{v.Field(0).Bytes(), v.Field(1).Uint()}
		err := pd.WriteBitString(bs, params.sizeExtensible, params.sizeLowerBound,
			params.sizeUpperBound)
		return err
	case ObjectIdentifierType:
		err := errors.Errorf("legacyMakeField error: Unsupport ObjectIdenfier type")
		return err
	case OctetStringType:
		err := pd.WriteOctetString(v.Bytes(), params.sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		return err
	case EnumeratedType:
		err := pd.WriteEnumerated(Enumerated(v.Uint()), params.valueExtensible,
			params.valueLowerBound, params.valueUpperBound)
		return err
	}

	switch val := v; val.Kind() {
	case reflect.Bool:
		err := pd.WriteBool(v.Bool())
		return err
	case reflect.Int, reflect.Int32, reflect.Int64:
		err := pd.WriteInteger(v.Int(), params.valueExtensible, params.valueLowerBound, params.valueUpperBound)
		return err

	case reflect.Struct:

		structType := fieldType
		var structParams []fieldParameters
		var handledOptionalCount uint
		var optionalPresentsFlag []bool
		var sequenceType bool
		// struct extensive TODO: support extensed type
		// if params.valueExtensible {
		// 	perTrace(2, fmt.Sprintf("Encoding Value Extensive Bit : %t (Extensed type is not yet supported)", false))
		// 	// pd.WriteExtensible(false)
		// }
		sequenceType = (structType.NumField() <= 0 || structType.Field(0).Name != "Present")
		// pass tag for optional
		for i := 0; i < structType.NumField(); i++ {
			if structType.Field(i).PkgPath != "" {
				return errors.Errorf("legacyMakeField error: struct contains unexported fields : %s", structType.Field(i).PkgPath)
			}
			tempParams := parseFieldParameters(structType.Field(i).Tag.Get("aper"))
			if sequenceType {
				// for optional flag
				if tempParams.optional {
					if !v.Field(i).IsNil() {
						optionalPresentsFlag = append(optionalPresentsFlag, true)
					} else {
						optionalPresentsFlag = append(optionalPresentsFlag, false)
					}
				} else if v.Field(i).Type().Kind() == reflect.Ptr && v.Field(i).IsNil() {
					return errors.Errorf("legacyMakeField error: nil element in SEQUENCE type")
				}
			}

			structParams = append(structParams, tempParams)
		}
		if sequenceType {
			if err := pd.WriteSequencePreambleBitMap(optionalPresentsFlag, params.valueExtensible); err != nil {
				return err
			}
		}

		// CHOICE or OpenType
		if !sequenceType {
			present := int(v.Field(0).Int())
			if present == 0 {
				return errors.Errorf("legacyMakeField error: " +
					"CHOICE or OpenType present is 0(present's field number)")
			} else if present >= structType.NumField() {
				return errors.Errorf("legacyMakeField error: " +
					"Present is bigger than number of struct field")
			} else if params.openType {
				if params.referenceFieldValue == nil {
					return errors.Errorf("legacyMakeField error: " +
						"OpenType reference value is empty")
				}
				refValue := *params.referenceFieldValue

				if structParams[present].referenceFieldValue == nil ||
					*structParams[present].referenceFieldValue != refValue {
					return errors.Errorf("legacyMakeField error: " +
						"reference value and present reference value is not match")
				}

				pdOpenType := NewPerBitData(nil)

				if err := pdOpenType.legacyMakeField(val.Field(present), structParams[present]); err != nil {
					return err
				}
				if err := pd.WriteOpenType(pdOpenType.Bytes()); err != nil {
					return err
				}
			} else {
				if err := pd.WriteChoicePreambleBitMap(int64(present-1), params.valueExtensible,
					params.valueUpperBound); err != nil {
					return err
				}
				if err := pd.legacyMakeField(val.Field(present), structParams[present]); err != nil {
					return err
				}
			}
			return nil
		}

		// struct that is neither CHOICE nor OPEN TYPE (SEQUENCE)
		for i := 0; i < structType.NumField(); i++ {
			// optional
			if structParams[i].optional {
				if !optionalPresentsFlag[handledOptionalCount] {
					handledOptionalCount++
					continue
				} else {
					handledOptionalCount++
					// perTrace(3, fmt.Sprintf("Field \"%s\" in %s is OPTIONAL and present", structType.Field(i).Name, structType))
				}
			}
			// for open type reference
			if structParams[i].openType {
				fieldName := structParams[i].referenceFieldName
				var index int
				for index = 0; index < i; index++ {
					if structType.Field(index).Name == fieldName {
						break
					}
				}
				if index == i {
					return errors.Errorf("legacyMakeField error: " +
						"Open type is not reference to the other field in the struct")
				}
				structParams[i].referenceFieldValue = new(int64)
				if value, err := legacyGetReferenceFieldValue(val.Field(index)); err != nil {
					return err
				} else {
					*structParams[i].referenceFieldValue = value
				}
			}
			if err := pd.legacyMakeField(val.Field(i), structParams[i]); err != nil {
				return err
			}
		}
		return nil
	case reflect.Slice:
		err := pd.WriteSequenceOfPreambleBitMap(uint64(v.Len()),
			params.sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)
		// encode components of SEQUENCE OF
		params.sizeExtensible = false
		params.sizeUpperBound = nil
		params.sizeLowerBound = nil
		for i := 0; i < v.Len(); i++ {
			if err = pd.legacyMakeField(v.Index(i), params); err != nil {
				return err
			}
		}
		return err
	case reflect.String:
		printableString := v.String()
		err := pd.WriteOctetString([]byte(printableString), params.sizeExtensible, params.sizeLowerBound,
			params.sizeUpperBound)
		return err
	}
	return errors.Errorf("legacyMakeField error: unsupported: %s", v.Type().String())
}

// Marshal returns the ASN.1 encoding of val.
func LegacyMarshal(val interface{}) ([]byte, error) {
	return LegacyMarshalWithParams(val, "")
}

// MarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func LegacyMarshalWithParams(val interface{}, params string) ([]byte, error) {
	pd := NewPerBitData(nil)
	err := pd.legacyMakeField(reflect.ValueOf(val), parseFieldParameters(params))
	if err != nil {
		return nil, errors.Wrap(err, "legacy marshal failed")
	}
	return pd.Bytes(), nil
}

/*** End of legacy function ***/
