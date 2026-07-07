package aper

import (
	"reflect"

	"github.com/pkg/errors"
)

/*** General bit/byte-operation functions ***/
func (pd *PerBitData) getSingleBit() (byte, error) {
	if int(pd.byteOffset) > len(pd.bytes)-1 {
		return 0, errors.Errorf(
			"getSingleBit(): pd.byteOffset>len(pd.bytes)-1 (%d > %d)",
			pd.byteOffset,
			len(pd.bytes)-1,
		)
	}
	dstByte := (pd.bytes[pd.byteOffset] >> (8 - pd.bitsOffset - 1)) & 0x01
	pd.bitsOffset += 1
	pd.bitCarry()
	return dstByte, nil
}

func (pd *PerBitData) getBitsBase(
	bitLength uint64,
	// isPeek: if true, bit/byte offset of pd won't be updated
	isPeek bool,
	// isLeftAligned: if false, bits are right-aligned in []bytes, which is better for bits value calculation;
	//      otherwise, bits are left-aligned, which is used for e.g. BitString format with length specified
	isLeftAligned bool,
) ([]byte, error) {
	byteFirst := pd.byteOffset
	bitOffsetForFirstByte := pd.bitsOffset
	byteLast := byteFirst + (bitLength+uint64(bitOffsetForFirstByte)+7)>>3 - 1
	numBitsForLastByte := (uint64(pd.bitsOffset) + bitLength) & 0x7
	if numBitsForLastByte == 0 {
		numBitsForLastByte = 8
	}
	numBitsLeftForLastByte := 8 - numBitsForLastByte

	// some checks
	if byteLast < byteFirst {
		return nil, errors.Errorf("getBits error: byteLast < byteFirst")
	}
	if int(byteLast) > len(pd.bytes)-1 {
		return nil, errors.Errorf("getBits error: outside the valid range")
	}

	// get target bytes
	dstBytes := append([]byte{}, pd.bytes[byteFirst:byteLast+1]...) // upper bound excluded

	// Align bits to left/right of []byte
	if !isLeftAligned {
		// Align bits to right of []byte

		// remove first <bitOffsetForFirstByte> bits
		var firstByteBitMask byte = 0xff
		firstByteBitMask >>= bitOffsetForFirstByte
		dstBytes[0] &= firstByteBitMask

		// shifting <numBitsLeftForLastByte> bits right
		if numBitsLeftForLastByte > 0 {
			shiftLeft(&dstBytes, int(numBitsLeftForLastByte)*(-1))
		}

		// remove redundant bytes after shifting
		byteLen := (bitLength + 7) >> 3
		numRedundantBytes := len(dstBytes) - int(byteLen)
		dstBytes = dstBytes[numRedundantBytes:]
	} else {
		// Align bits to left of []byte

		// remove last <numBitsLeftForLastByte> bits
		var lastByteBitMask byte = 0xff
		lastByteBitMask <<= numBitsLeftForLastByte
		dstBytes[len(dstBytes)-1] &= lastByteBitMask

		// shifting <bitOffsetForFirstByte> bits left
		if bitOffsetForFirstByte > 0 {
			shiftLeft(&dstBytes, int(bitOffsetForFirstByte))
		}

		// remove redundant bytes after shifting
		byteLen := (bitLength + 7) >> 3
		dstBytes = dstBytes[:byteLen]
	}

	// if isPeek, no need to update the bit/byte offset of pd
	if !isPeek {
		pd.bitsOffset += uint(bitLength)
		pd.bitCarry()
	}

	return dstBytes, nil
}

// bits are right-aligned in []byte
func (pd *PerBitData) getBits(bitLength uint64) ([]byte, error) {
	return pd.getBitsBase(bitLength, false, false)
}

// bits are left-aligned in []byte
func (pd *PerBitData) getBitsForOctetString(bitLength uint64) ([]byte, error) {
	return pd.getBitsBase(bitLength, false, true)
}

// // bits are right-aligned in []byte, and bit/byte offset of pd won't be updated.
func (pd *PerBitData) peekBits(bitLength uint64) ([]byte, error) {
	return pd.getBitsBase(bitLength, true, false)
}

// getBitsValueBase: int64
func (pd *PerBitData) getBitsValueBase(bitLength uint64, isPeek bool, is2sComplement bool) (int64, error) {
	// get bits of bitLength (in []byte)
	var bytes []byte
	var err error
	if isPeek {
		bytes, err = pd.peekBits(bitLength)
	} else {
		bytes, err = pd.getBits(bitLength)
	}
	if err != nil {
		return -1, errors.Wrap(err, "git bit value failed")
	}
	return bytesToInt64(bytes, is2sComplement)
}

func (pd *PerBitData) getBitsValue(bitLength uint64, is2sComplement bool) (int64, error) {
	return pd.getBitsValueBase(bitLength, false, is2sComplement)
}

func (pd *PerBitData) peekBitsValue(bitLength uint64, is2sComplement bool) (int64, error) {
	return pd.getBitsValueBase(bitLength, true, is2sComplement)
}

// getBitsValueBase: uint64
func (pd *PerBitData) getBitsValueBaseForUint64(bitLength uint64, isPeek bool) (uint64, error) {
	// get bits of bitLength (in []byte)
	var bytes []byte
	var err error
	if isPeek {
		bytes, err = pd.peekBits(bitLength)
	} else {
		bytes, err = pd.getBits(bitLength)
	}
	if err != nil {
		return 0, errors.Wrap(err, "get bit value (uint64) failed")
	}
	return bytesToUint64(bytes)
}

func (pd *PerBitData) getBitsValueForUint64(bitLength uint64) (uint64, error) {
	return pd.getBitsValueBaseForUint64(bitLength, false)
}

// func (pd *PerBitData) peekBitsValueForUint64(bitLength uint64) (uint64, error) {
//  return pd.getBitsValueBaseForUint64(bitLength, true)
// }

func (pd *PerBitData) getAlignBits() error {
	if alignBits := uint64(8-pd.bitsOffset&0x7) & 0x7; alignBits != 0 {
		// pd.getBitsValue includes bit/byte offset calculation
		if val, err := pd.getBitsValue(alignBits, false); err != nil {
			return err
		} else if val != 0 {
			if !skipPaddingCheck {
				return errors.Errorf("getAlignBits error: Align Bit is not zero")
			}
		}
	}
	return nil
}

/*** End of general bit/byte-operation functions ***/

/*** Decoding different types for ASN.1 ***/

func (pd *PerBitData) ReadExtensible() (bool, error) {
	b, err := pd.getSingleBit()
	if err != nil {
		return false, errors.Wrap(err, "ReadExtensible()")
	}
	return b == 1, nil
}

// 有約束、半約束、無約束、常見小自然數
func (pd *PerBitData) ReadWholeNum(isNormallySmallNonNegativeNum bool,
	lbPtr *int64, ubPtr *int64,
) (int64, error) {
	if isNormallySmallNonNegativeNum {
		// IV. 常见小自然数编码
		//  This procedure is used when encoding a non-negative whole number that is expected to be small,
		//  but whose size is potentially unlimited due to the presence of an extension marker
		// 这种情况经常出现在对表征SEQUENCE、SET类型可选成员的Bitmap长度进行编码时；
		// 或者CHOICE类型序号编码时。这种长度相当小，但是却没有一种限定。
		// | 当0≤n≤63时，n以6个比特编码，并且在前面增加一个 0-bit（八位组不对齐）
		// | 当64≤n时，n以半约束数方式编码，下边界为0，并且在前面增加一个 1-bit
		firstBit, err := pd.getSingleBit()
		if err != nil {
			return -1, errors.Wrap(err, "ReadWholeNum()")
		}
		switch firstBit {
		case 0:
			// represent value in 6 bits
			val, err := pd.getBitsValue(6, false)
			return val, err
		case 1:
			var tmpLb int64 = 0
			val, err := pd.ReadWholeNum(false, &tmpLb, nil)
			return val, err
		}
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
		if d < 0 {
			err := errors.Errorf("ReadWholeNum error: value range is negative")
			return 0, err
		} else if d == 1 { // ub/lb itself is the value (no bits encoded)
			return *ubPtr, nil
		} else if d <= 255 {
			numBits := getRequiredNumBits(d-1, false) // num of bits required to represent d numbers = log2(d-1)
			val, err := pd.getBitsValue(numBits, false)
			val += *lbPtr
			return val, err
		} else if d == 256 {
			err := pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number failed")
			}
			val, err := pd.getBitsValue(8*1, false)
			val += *lbPtr
			return val, err
		} else if d <= 65536 {
			err := pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number failed")
			}
			val, err := pd.getBitsValue(8*2, false)
			val += *lbPtr
			return val, err
		} else {
			// encode length
			// (12.2.6: 用 constrained length determinate 編碼，lb=1，ub=log256d)。
			dNumBytes := getRequiredNumBytes(d-1, false)
			lmin := uint64(1)
			encodeValueNumBytes, err := pd.ReadLength(false, &lmin, &dNumBytes)
			if err != nil {
				return 0, errors.Wrap(err, "read whold number failed")
			}
			// decode value
			err = pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number failed")
			}
			val, err := pd.getBitsValue(8*encodeValueNumBytes, false)
			val += *lbPtr
			return val, err
		}
	} else if lbPtr == nil {
		// III. 无约束数编码
		// 无约束指值域没有下边界（即使存在上边界）。
		// n encoded into octects as a 2’s-complement-binary-integer
		// n 的编码占用log256d个八位组，并且在前面增加L的编码。
		numBytes, err := pd.ReadLength(false, nil, nil)
		if err != nil {
			return 0, errors.Wrap(err, "read whole number failed")
		}

		// This check is out of spec regulation. We set this check for IOT flexibility to allow
		// invalid input []byte{0x00} for unconstrained int decode. Aper decoding []byte{0x00}
		// will get 0 instead of an error after adding this check.
		if numBytes == 0 {
			return 0, nil
		}

		return pd.getBitsValue(8*numBytes, true)
	} else if ubPtr == nil {
		// II. 半约束数编码
		// 半约束指值域没有上边界（上边界为+∞）。
		// n-bmin encoded into octects as a non-negative-binary-integer
		// n-bmin的编码占用log256d个八位组，并且在前面增加L的编码。
		lb := uint64(*lbPtr)
		numBytes, err := pd.ReadLength(false, &lb, nil)
		if err != nil {
			return 0, errors.Wrap(err, "read whole number failed")
		}
		val, err := pd.getBitsValue(8*numBytes, false)
		val += *lbPtr
		return val, err
	} else {
		err := errors.Errorf("ReadWholeNum error: write whole number error.")
		return 0, err
	}
}

func (pd *PerBitData) ReadUint64WholeNum(isNormallySmallNonNegativeNum bool,
	lbPtr *uint64, ubPtr *uint64,
) (uint64, error) {
	if isNormallySmallNonNegativeNum {
		// IV. 常见小自然数编码
		//  This procedure is used when encoding a non-negative whole number that is expected to be small,
		//  but whose size is potentially unlimited due to the presence of an extension marker
		// 这种情况经常出现在对表征SEQUENCE、SET类型可选成员的Bitmap长度进行编码时；
		// 或者CHOICE类型序号编码时。这种长度相当小，但是却没有一种限定。
		// | 当0≤n≤63时，n以6个比特编码，并且在前面增加一个 0-bit（八位组不对齐）
		// | 当64≤n时，n以半约束数方式编码，下边界为0，并且在前面增加一个 1-bit
		firstBit, err := pd.getSingleBit()
		if err != nil {
			return 0, errors.Wrap(err, "ReadUint64WholeNum()")
		}
		switch firstBit {
		case 0:
			// represent value in 6 bits
			val, err := pd.getBitsValueForUint64(6)
			return val, err
		case 1:
			var tmpLb uint64 = 0
			val, err := pd.ReadUint64WholeNum(false, &tmpLb, nil)
			return val, err
		}
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
			// decode length
			// (12.2.6: 用 constrained length determinate 編碼，lb=1，ub=log256d)。
			dNumBytes := uint64(8)
			lmin := uint64(1)
			encodeValueNumBytes, err := pd.ReadLength(false, &lmin, &dNumBytes)
			if err != nil {
				return 0, errors.Wrap(err, "read whole number (uint64) failed")
			}
			// decode value
			err = pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number (uint64) failed")
			}
			val, err := pd.getBitsValueForUint64(8 * encodeValueNumBytes)
			val += *lbPtr
			return val, err
		}

		d := *ubPtr - *lbPtr + 1
		if d == 0 {
			err := errors.Errorf("ReadUint64WholeNum error: value range is zero")
			return 0, err
		} else if d == 1 { // ub/lb itself is the value (no bits encoded)
			return *ubPtr, nil
		} else if d <= 255 {
			numBits := getRequiredNumBitsForUint64(d - 1) // num of bits required to represent d numbers = log2(d-1)
			val, err := pd.getBitsValueForUint64(numBits)
			val += *lbPtr
			return val, err
		} else if d == 256 {
			err := pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number (uint64) failed")
			}
			val, err := pd.getBitsValueForUint64(8 * 1)
			val += *lbPtr
			return val, err
		} else if d <= 65536 {
			err := pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number (uint64) failed")
			}
			val, err := pd.getBitsValueForUint64(8 * 2)
			val += *lbPtr
			return val, err
		} else {
			// decode length
			// (12.2.6: 用 constrained length determinate 編碼，lb=1，ub=log256d)。
			dNumBytes := getRequiredNumBytesForUint64(d - 1)
			lmin := uint64(1)
			encodeValueNumBytes, err := pd.ReadLength(false, &lmin, &dNumBytes)
			if err != nil {
				return 0, errors.Wrap(err, "read whole number (uint64) failed")
			}
			// decode value
			err = pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read whole number (uint64) failed")
			}
			val, err := pd.getBitsValueForUint64(8 * encodeValueNumBytes)
			val += *lbPtr
			return val, err
		}
	} else if lbPtr == nil {
		// III. 无约束数编码
		// 无约束指值域没有下边界（即使存在上边界）。
		// n encoded into octects as a 2’s-complement-binary-integer
		// n 的编码占用log256d个八位组，并且在前面增加L的编码。
		numBytes, err := pd.ReadLength(false, nil, nil)
		if err != nil {
			return 0, errors.Wrap(err, "read whole number (uint64) failed")
		}
		return pd.getBitsValueForUint64(8 * numBytes)
	} else if ubPtr == nil {
		// II. 半约束数编码
		// 半约束指值域没有上边界（上边界为+∞）。
		// n-bmin encoded into octects as a non-negative-binary-integer
		// n-bmin的编码占用log256d个八位组，并且在前面增加L的编码。
		lb := *lbPtr
		numBytes, err := pd.ReadLength(false, &lb, nil)
		if err != nil {
			return 0, errors.Wrap(err, "read whole number (uint64) failed")
		}
		val, err := pd.getBitsValueForUint64(8 * numBytes)
		val += *lbPtr
		return val, err
	} else {
		err := errors.Errorf("ReadUint64WholeNum unexpected error.")
		return 0, err
	}
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
func (pd *PerBitData) ReadInteger(valueExtensible bool,
	lbPtr *int64, ubPtr *int64,
) (int64, error) {
	var lb, ub int64
	if ubPtr != nil {
		ub = *ubPtr
	}
	if lbPtr != nil {
		lb = *lbPtr
	}

	// Case 1.
	// Decide the value constraint (-1 means no valid size constraint)
	if valueExtensible {
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return -1, errors.Wrap(err, "ReadInteger()")
		}
		if isExt { // Case 1-2.
			return pd.ReadWholeNum(false, nil, nil)
		} // else: Case 1-1.,Continue to use Case 2 to decode value
	}

	// Case 2-1.
	if lbPtr != nil && ubPtr != nil && (lb == ub) {
		return lb, nil
	}
	// Case 2-2.
	return pd.ReadWholeNum(false, lbPtr, ubPtr)
}

func (pd *PerBitData) ReadUint64Integer(valueExtensible bool,
	lbPtr *uint64, ubPtr *uint64,
) (uint64, error) {
	var lb, ub uint64
	if ubPtr != nil {
		ub = *ubPtr
	}
	if lbPtr != nil {
		lb = *lbPtr
	}

	// Case 1.
	// Decide the value constraint (-1 means no valid size constraint)
	if valueExtensible {
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return 0, errors.Wrap(err, "ReadUint64Integer()")
		}
		if isExt { // Case 1-2.
			return pd.ReadUint64WholeNum(false, nil, nil)
		} // else: Case 1-1.,Continue to use Case 2 to decode value
	}

	// Case 2-1.
	if lbPtr != nil && ubPtr != nil && (lb == ub) {
		return lb, nil
	}
	// Case 2-2.
	return pd.ReadUint64WholeNum(false, lbPtr, ubPtr)
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
func (pd *PerBitData) ReadLength(isForBitMapLen bool, lbPtr *uint64, ubPtr *uint64) (uint64, error) {
	if isForBitMapLen { // case 1: for set or sequence type with extension marker
		l, err := pd.ReadWholeNum(true, nil, nil)
		return uint64(l + 1), err
	} else {
		if ubPtr != nil && *ubPtr <= 65535 { // case 2.
			lb := int64(*lbPtr)
			ub := int64(*ubPtr)
			l, err := pd.ReadWholeNum(false, &lb, &ub)
			return uint64(l), err
		} else { // case 3. (ubPtr == nil || *ubPtr > 65535)
			err := pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read length failed")
			}
			peekByte, err := pd.peekBitsValue(8, false)
			if err != nil {
				return 0, errors.Wrap(err, "read length failed")
			} else {
				if peekByte>>7 == 0 { // case 3-1.
					val, err := pd.getBitsValue(8*1, false)
					return uint64(val), err
				} else if peekByte>>6 == 2 { // case 3-2.
					// remove the first two bits (10)
					bits, err := pd.getBits(8 * 2)
					if err != nil {
						return 0, errors.Wrap(err, "read length failed")
					}
					bits[0] &= 0x3F
					val, err := bytesToInt64(bits, false)
					if err != nil {
						return 0, errors.Wrap(err, "read length failed")
					}
					return uint64(val), err
				} else if peekByte>>6 == 3 { // case 3-3 (l >= 16384)
					// remove the first two bits (c0)
					bits, err := pd.getBits(8)
					if err != nil {
						return 0, errors.Wrap(err, "read length failed")
					}
					bits[0] &= 0x3F
					val, err := bytesToInt64(bits, false)
					if err != nil {
						return 0, errors.Wrap(err, "read length failed")
					}
					return uint64(val) * 16384, err
				}
			}
		}
	}

	return 0, errors.Errorf("ReadLength error: fail to parse Length")
}

// Case 3-3.
// 当16,384≤l，整个编码以f*16K为单位分割（f取值为1，2，3或者4）。
// 除最后的片段外，其余每段，长度為 f 值以一个八位组编码 (用無約束方式)，最高两个比特位为11。
// 如果编码恰好时16K的整倍数，则在最后补充一个全空的八位组；否则最后一个片段按照前两条进行编码。
func (pd *PerBitData) ReadLargeBitStringWithFragment(dstBsPtr *BitString, fragmentLen uint64,
	lbPtr *uint64, ubPtr *uint64,
) error {
	for {
		// Decide fragmentation size and write fragmentation length
		(*dstBsPtr).BitLength += fragmentLen

		err := pd.getAlignBits()
		if err != nil {
			return errors.Wrap(err, "read large bitstring failed")
		}
		// Read fragmentation data
		bits, err := pd.getBits(fragmentLen)
		if err != nil {
			return errors.Wrap(err, "read large bitstring failed")
		}
		(*dstBsPtr).Bytes = append((*dstBsPtr).Bytes, bits...)

		// Check whether this is the last fragment
		if fragmentLen >= 16384 {
			// TO-BE-TEST
			peekBitsVal, tmpErr := pd.peekBitsValue(8, false)
			if tmpErr == nil && peekBitsVal == 0 {
				// if fragment size > 16383 and the next byte is 0 -> last byte (case: the bitstring length is times of 16K)
				pd.bitsOffset += 8
				pd.bitCarry()
				tmpErr = pd.getAlignBits()
				return tmpErr
			} else {
				// otherwise, this is not the last fragment
				fragmentLen, err = pd.ReadLength(false, lbPtr, ubPtr)
				if err != nil {
					return errors.Wrap(err, "read large bitstring failed")
				}
				continue
			}
		} else {
			// last fragment
			err := pd.getAlignBits()
			return err
		}
	}
}

func (pd *PerBitData) ReadLargeOctetStringWithFragment(dstOsPtr *OctetString, fragmentLen uint64,
	lbPtr *uint64, ubPtr *uint64,
) error {
	for {
		// Decide fragmentation size and write fragmentation length
		err := pd.getAlignBits()
		if err != nil {
			return errors.Wrap(err, "read large octetstring failed")
		}
		// Read fragmentation data
		bits, err := pd.getBits(fragmentLen * 8)
		if err != nil {
			return errors.Wrap(err, "read large octetstring failed")
		}
		(*dstOsPtr) = append((*dstOsPtr), bits...)

		// Check whether this is the last fragment
		if fragmentLen >= 16384 {
			// TO-BE-TEST
			peekBitsVal, tmpErr := pd.peekBitsValue(8, false)
			if tmpErr == nil && peekBitsVal == 0 {
				// if fragment size > 16383 and the next byte is 0 -> last byte (case: the bitstring length is times of 16K)
				pd.bitsOffset += 8
				pd.bitCarry()
				tmpErr = pd.getAlignBits()
				return tmpErr
			} else {
				// otherwise, this is not the last fragment
				fragmentLen, err = pd.ReadLength(false, lbPtr, ubPtr)
				if err != nil {
					return errors.Wrap(err, "read large octetstring failed")
				}
				continue
			}
		} else {
			// last fragment
			err := pd.getAlignBits()
			return err
		}
	}
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
func (pd *PerBitData) ReadBitString(sizeExtensible bool, lbPtr *uint64, ubPtr *uint64) (BitString, error) {
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

	// Case 1.
	if sizeExtensible {
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return BitString{}, errors.Wrap(err, "ReadBitString()")
		}
		if isExt { // Case 1-2.
			l, err := pd.ReadLength(false, lbPtr, nil)
			if err != nil {
				return BitString{}, errors.Wrap(err, "read bitstring failed")
			}
			err = pd.getAlignBits()
			if err != nil {
				return BitString{}, errors.Wrap(err, "read bitstring failed")
			}
			bits, err := pd.getBitsForOctetString(l)
			return BitString{bits, l}, err
		} // else: Case 1-1.,Continue to use Case 2 to decode value
	}

	// Decode Bitstring
	// case 2.
	if ubPtr != nil && *ubPtr == 0 { // case 2-1.   // TO-BE-TEST
		return BitString{}, nil
	}

	if lbPtr != nil && ubPtr != nil && lb == ub { // ub == lb
		if ub <= 16 { // case 2-2
			bits, err := pd.getBitsForOctetString(ub)
			if err != nil {
				return BitString{}, errors.Wrap(err, "read bitstring failed")
			}
			return BitString{bits, ub}, nil
		} else if ub <= 65536 { // case 2-3
			err := pd.getAlignBits()
			if err != nil {
				return BitString{}, errors.Wrap(err, "read bitstring failed")
			}
			bits, err := pd.getBitsForOctetString(ub)
			if err != nil {
				return BitString{}, errors.Wrap(err, "read bitstring failed")
			}
			return BitString{bits, ub}, nil
		}
	}

	// case 2-4
	l, err := pd.ReadLength(false, lbPtr, ubPtr)
	if err != nil {
		return BitString{}, errors.Wrap(err, "read bitstring failed")
	}
	if l <= 16383 {
		if l == 0 { // empty bit string
			return BitString{[]byte{}, 0}, err
		}
		tmpErr := pd.getAlignBits()
		if tmpErr != nil {
			return BitString{}, tmpErr
		}
		bits, tmpErr := pd.getBitsForOctetString(l)
		return BitString{bits, l}, tmpErr
	} else { // apply fragmentation
		bs := BitString{[]byte{}, 0}
		tmpErr := pd.ReadLargeBitStringWithFragment(&bs, l, lbPtr, ubPtr)
		return bs, tmpErr
	}
}

// OCTET STRING
// 与BIT STRING规则相同，不过长度域L表征的是八位组的个数而不是比特数。
func (pd *PerBitData) ReadOctetString(sizeExtensible bool, lbPtr *uint64, ubPtr *uint64) (OctetString, error) {
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

	// Case 1.
	if sizeExtensible {
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return OctetString{}, errors.Wrap(err, "ReadOctetString()")
		}
		if isExt { // Case 1-2.
			l, err := pd.ReadLength(false, lbPtr, nil)
			if err != nil {
				return OctetString{}, errors.Wrap(err, "read octetstring failed")
			}
			err = pd.getAlignBits()
			if err != nil {
				return OctetString{}, errors.Wrap(err, "read octetstring failed")
			}
			bits, err := pd.getBitsForOctetString(l * 8)
			return bits, err
		} // else: Case 1-1.,Continue to use Case 2 to decode value
	}

	// Decode OctetString
	// case 2.
	if ubPtr != nil && *ubPtr == 0 { // case 2-1.   // TO-BE-TEST
		return OctetString{}, nil
	}

	if lbPtr != nil && ubPtr != nil && lb == ub { // ub == lb
		if ub <= 2 { // case 2-2
			bits, err := pd.getBitsForOctetString(ub * 8)
			if err != nil {
				return OctetString{}, errors.Wrap(err, "read octetstring failed")
			}
			return bits, nil
		} else if ub <= 65536 { // case 2-3
			err := pd.getAlignBits()
			if err != nil {
				return OctetString{}, errors.Wrap(err, "read octetstring failed")
			}
			bits, err := pd.getBitsForOctetString(ub * 8)
			if err != nil {
				return OctetString{}, errors.Wrap(err, "read octetstring failed")
			}
			return bits, nil
		}
	}

	// case 2-4
	l, err := pd.ReadLength(false, lbPtr, ubPtr)
	if err != nil {
		return OctetString{}, errors.Wrap(err, "read octetstring failed")
	}
	if l <= 16383 {
		if l == 0 { // empty octet string
			return []byte{}, err
		}
		tmpErr := pd.getAlignBits()
		if tmpErr != nil {
			return OctetString{}, tmpErr
		}
		bits, tmpErr := pd.getBitsForOctetString(l * 8)
		return bits, tmpErr
	} else { // apply fragmentation
		os := OctetString{}
		tmpErr := pd.ReadLargeOctetStringWithFragment(&os, l, lbPtr, ubPtr)
		return os, tmpErr
	}
}

// Characters: PrintableString, VisibleString, UTF8String
// PER-visible constraints only apply to known-multiplier character string types.
// For other restricted character string types
// "aub" will be unset and "alb" will be zero. known-multiplier character string types: NumericString,
// PrintableString, VisibleString (ISO646String), IA5String, BMPString, and UniversalString
func (pd *PerBitData) ReadPrintableString(sizeExtensible bool, lbPtr *uint64, ubPtr *uint64) (PrintableString, error) {
	os, err := pd.ReadOctetString(sizeExtensible, lbPtr, ubPtr)
	return PrintableString(os), err
}

func (pd *PerBitData) ReadVisibleString(sizeExtensible bool, lbPtr *uint64, ubPtr *uint64) (VisibleString, error) {
	os, err := pd.ReadOctetString(sizeExtensible, lbPtr, ubPtr)
	return VisibleString(os), err
}

func (pd *PerBitData) ReadUTF8String(sizeExtensible bool, lbPtr *uint64, ubPtr *uint64) (UTF8String, error) {
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

	os, err := pd.ReadOctetString(sizeExtensible, lbPtr, ubPtr)
	return UTF8String(os), err
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
//	     v2 ENUMERATED {orange(56), green(-2), red(2476), …, yellow, purple} ::= yellow
//	的PER编码则为“10000000”。
func (pd *PerBitData) ReadEnumerated(valueExtensible bool,
	lbPtr *int64, ubPtr *int64,
) (Enumerated, error) {
	// Check lb and ub
	if lbPtr == nil || ubPtr == nil {
		return 0, errors.Errorf("ReadEnumerated error: ENUMERATED value constraint is error")
	}

	lb, ub := *lbPtr, *ubPtr

	if lb < 0 || lb > ub {
		return 0, errors.Errorf("ReadEnumerated error: ENUMERATED value constraint is error")
	}

	if !valueExtensible { // Case 1
		val, err := pd.ReadWholeNum(false, &lb, &ub)
		if err != nil {
			return Enumerated(val), errors.Wrap(err, "read enumerated failed")
		}
		if lb > val || val > ub {
			return Enumerated(val), errors.Errorf("ReadEnumerated error: " +
				"ENUMERATED value is out of valid range")
		}
		return Enumerated(val), nil
	} else { // Case 2
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return Enumerated(0), errors.Wrap(err, "ReadEnumerated()")
		}
		if isExt { // Case 2-2.
			val, err := pd.ReadWholeNum(true, nil, nil)
			if err != nil {
				return Enumerated(0), errors.Wrap(err, "read enumerated failed")
			}
			return Enumerated(val + ub + 1), nil
		} else { // Case 2-1.
			val, err := pd.ReadWholeNum(false, &lb, &ub)
			if err != nil {
				return Enumerated(val), errors.Wrap(err, "read enumerated failed")
			}
			if lb > val || val > ub {
				return Enumerated(val), errors.Errorf("ReadEnumerated error: " +
					"ENUMERATED value is out of valid range")
			}
			return Enumerated(val), nil
		}
	}
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
func (pd *PerBitData) ReadSequenceOfPreambleBitMap(sizeExtensible bool, lbPtr *uint64, ubPtr *uint64) (uint64, error) {
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

	// Write Preamble Bit if is size extensible
	if sizeExtensible {
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return 0, errors.Wrap(err, "ReadSequenceOfPreambleBitMap()")
		}
		if isExt { // Case 1-2.
			// get number of components for sequence of
			numComponents, err := pd.ReadLength(false, lbPtr, nil)
			if err != nil {
				return 0, errors.Wrap(err, "read seqof failed")
			}
			err = pd.getAlignBits()
			if err != nil {
				return 0, errors.Wrap(err, "read seqof failed")
			}
			// Read preamble bitmap for sequenco of
			// Please continue with using other pd.Read functions according to
			// the sequence of components type to complete the decoding process.
			return numComponents, nil
		} // else: Case 1-1., Continue to encode using Case 2.
	}

	// Decode Length determinant if required
	if ubPtr != nil && lbPtr != nil && (ub == lb) && ub < 65536 { // Case 2-1.
		// Read preamble bitmap for sequenco of
		// Please continue with using other pd.Read functions according to
		// the sequence of components type to complete the decoding process.
		return ub, nil
	} else {
		numComponents, err := pd.ReadLength(false, lbPtr, ubPtr)
		if err != nil {
			return 0, errors.Wrap(err, "read seqof failed")
		}
		if numComponents <= 16383 { // Case 2-2.
			// Read preamble bitmap for sequenco of
			// Please continue with using other pd.Read functions according to
			// the sequence of components type to complete the decoding process.
			return numComponents, nil
		} else { // Case 2-3.
			// TODO: Fragmentation is not supported here since length interleaving with data.
			// Encoding requires knowledge about sequence of type (e.g. using reflect)
			// to encode fragmented components preceded with length determinant of each fragment.
			return numComponents, errors.Errorf("ReadSequenceOfPreambleBitMap error: " +
				"unsupport sequence of size larger than 16383." +
				"Require fragmentation implementation.")
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
func (pd *PerBitData) ReadChoicePreambleBitMap(valueExtensible bool, ubPtr *int64) (int64, error) {
	// Check ub
	if ubPtr == nil {
		return 0, errors.Errorf("ReadChoicePreambleBitMap error: the upper bound of CHIOCE is missing")
	}

	lb, ub := int64(0), *ubPtr

	if ub < 0 {
		return 0, errors.Errorf("ReadChoicePreambleBitMap error: the upper bound of CHIOCE is negative")
	}

	if valueExtensible {
		isExt, err := pd.ReadExtensible()
		if err != nil {
			return -1, errors.Wrap(err, "ReadChoicePreambleBitMap()")
		}
		if isExt { // Case 2-2.
			index, err := pd.ReadWholeNum(true, &lb, nil)
			if err != nil {
				return index, errors.Wrap(err, "read choice failed")
			}
			// Read preamble bitmap for choice
			// Please continue with using other pd.Read functions according to
			// the sequence of components type to complete the decoding process.
			return index, err
		} // else: Case 2-1
	} // else: Case 2-1

	// Case 2-1
	// Case 1 must happens along with Case 2-1 (selected choice is in root)
	if ub == 0 {
		// Only one choice in root and the choice is selected, no index encoded
		return 0, nil // index is 0
	}

	index, err := pd.ReadInteger(false, &lb, &ub)
	return index, err
}

// SEQUENCE
//  0. 如果SEQUENCE类型是可扩展的，则在编码的头部加上一个比特的bit-field，
//     (Unsupported) 如果SEQUENCE的取值中有属于扩展附加部分的成员，则该比特等于1，否则等于0。
//  1. 如果SEQUENCE的定义中在扩展根部（extension root）有"n"个成员被置为OPTIONAL或DEFAULT，
//     则在编码头部再添加"n"个比特的bit-field，该bit-field从第一个bit开始，依次指示被标记为OPTIONAL或DEFAULT的成员是否出现。
//     - 如果“n”小于64K，则这个 bit-field 应该直接添到码流中。
//     - (Unsupported) 如果"n"大于等于64K，按照前面提到的处理方法把“n”个bit的bit-field分段并添加到域序列中，
//     前面的长度字段L就作为一个有约束的整数编码，而约束的上限和下限都等于n。
func (pd *PerBitData) ReadSequencePreambleBitMap(dstOptionalPresentFlag *[]bool,
	valueExtensible bool,
) error { // return optional present flag
	// TO-BE-TEST

	// Case 0.: write extensible bits
	if valueExtensible {
		_, err := pd.ReadExtensible()
		if err != nil {
			return errors.Wrap(err, "ReadSequencePreambleBitMap()")
		}
		// TODO: Support extensible sequence
		// Extensible Sequence is not supported. Extensible bit is always 0.
	}

	// Case 1: write optional bit-field
	numOptionalFields := len(*dstOptionalPresentFlag)
	if numOptionalFields == 0 {
		// No optional field for sequence
		return nil
	} else if numOptionalFields > 0 && numOptionalFields < 65536 {
		optionalBits, err := pd.getBitsValue(uint64(numOptionalFields), false)
		if err != nil {
			return errors.Wrap(err, "read sequence failed")
		}
		for i := numOptionalFields - 1; i >= 0; i-- {
			if optionalBits&0x1 == 1 {
				(*dstOptionalPresentFlag)[i] = true
			} else {
				(*dstOptionalPresentFlag)[i] = false
			}
			optionalBits >>= 1
		}
		return err
	} else { // numOptionalFields > 64K (Unsupported)
		return errors.Errorf("ReadSequencePreambleBitMap error: " +
			"number of optional fields more than 64K is not supported.")
	}
}

// OPEN TYPE
// 0. 真正類型的內容要先被 encode 成一個長度為 "n" 的 octet string
// 1. 在 0. 的 octet string 前面用無約束編碼的方式編碼長度 (n，in units of octets)
func (pd *PerBitData) ReadOpenType() ([]byte, error) {
	byteLen, err := pd.ReadLength(false, nil, nil)
	if err != nil {
		return []byte{}, errors.Wrap(err, "read opentype failed")
	}
	if byteLen <= 16383 {
		if byteLen == 0 { // empty octect string: don't encode
			return []byte{}, nil
		}
		err = pd.getAlignBits()
		if err != nil {
			return []byte{}, errors.Wrap(err, "read opentype failed")
		}
		return pd.getBits(byteLen * 8)
	} else {
		openTypeBytes := []byte{}
		err = pd.ReadLargeOctetStringWithFragment((*OctetString)(&openTypeBytes),
			byteLen, nil, nil)
		return openTypeBytes, err
	}
}

// BOOLEAN
func (pd *PerBitData) ReadBool() (bool, error) {
	b, err := pd.getSingleBit()
	if err != nil {
		return false, errors.Wrap(err, "ReadBool()")
	}
	if b == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

/*** End of encoding different types for ASN.1 ***/

/*** Legacy fuction for encoding an empty interface (using reflect) ***/
// parseField is the main parsing function. Given a byte slice and an offset
// into the array, it will try to parse a suitable ASN.1 value out and store it
// in the given Value. TODO : ObjectIdenfier, handle extension Field
func legacyParseField(v reflect.Value, pd *PerBitData, params fieldParameters) error {
	fieldType := v.Type()

	// If we have run out of data return error.
	if pd.byteOffset == uint64(len(pd.bytes)) {
		return errors.Errorf("legacyParseField error: sequence truncated")
	}
	if v.Kind() == reflect.Ptr {
		ptr := reflect.New(fieldType.Elem())
		v.Set(ptr)
		return legacyParseField(v.Elem(), pd, params)
	}

	// We deal with the structures defined in this package first.
	switch fieldType {
	case BitStringType:
		bitString, err1 := pd.ReadBitString(params.sizeExtensible, params.sizeLowerBound, params.sizeUpperBound)

		if err1 != nil {
			return err1
		}
		v.Set(reflect.ValueOf(bitString))
		return nil
	case ObjectIdentifierType:
		return errors.Errorf("legacyParseField error: unsupport ObjectIdenfier type")
	case OctetStringType:
		if octetString, err := pd.ReadOctetString(params.sizeExtensible,
			params.sizeLowerBound, params.sizeUpperBound); err != nil {
			return err
		} else {
			v.Set(reflect.ValueOf(octetString))
			return nil
		}
	case EnumeratedType:
		if parsedEnum, err := pd.ReadEnumerated(params.valueExtensible, params.valueLowerBound,
			params.valueUpperBound); err != nil {
			return err
		} else {
			v.SetUint(uint64(parsedEnum))
			return nil
		}
	}
	switch val := v; val.Kind() {
	case reflect.Bool:
		if parsedBool, err := pd.ReadBool(); err != nil {
			return err
		} else {
			val.SetBool(parsedBool)
			return nil
		}
	case reflect.Int, reflect.Int32, reflect.Int64:
		if parsedInt, err := pd.ReadInteger(params.valueExtensible,
			params.valueLowerBound, params.valueUpperBound); err != nil {
			return err
		} else {
			val.SetInt(parsedInt)
			return nil
		}
	case reflect.Struct:

		structType := fieldType
		var structParams []fieldParameters
		var handledOptionalCount uint64
		var optionalPresentsFlag []bool
		// var optionalPresents uint64

		// pass tag for optional
		for i := 0; i < structType.NumField(); i++ {
			if structType.Field(i).PkgPath != "" {
				return errors.Errorf("legacyParseField error: struct contains unexported fields : %s",
					structType.Field(i).PkgPath)
			}
			tempParams := parseFieldParameters(structType.Field(i).Tag.Get("aper"))
			// for optional flag
			if tempParams.optional {
				optionalPresentsFlag = append(optionalPresentsFlag, false)
			}
			structParams = append(structParams, tempParams)
		}

		// CHOICE or OpenType
		if structType.NumField() > 0 && structType.Field(0).Name == "Present" {
			present := 0
			if params.openType {
				if params.referenceFieldValue == nil {
					return errors.Errorf("legacyParseField error: OpenType reference value is empty")
				}
				refValue := *params.referenceFieldValue

				for j, param := range structParams {
					if j == 0 {
						continue
					}
					if param.referenceFieldValue != nil && *param.referenceFieldValue == refValue {
						present = j
						break
					}
				}
				if present == 0 {
					return errors.Errorf("legacyParseField error: " +
						"OpenType reference value does not match any field")
				} else if present >= structType.NumField() {
					return errors.Errorf("legacyParseField error: " +
						"OpenType Present is bigger than number of struct field")
				} else {
					val.Field(0).SetInt(int64(present))
					if opentypeBytes, err := pd.ReadOpenType(); err != nil {
						return err
					} else {
						pdOpenType := NewPerBitData(opentypeBytes)
						err := legacyParseField(val.Field(present), pdOpenType, structParams[present])
						return err
					}
				}
			} else {
				if index, err := pd.ReadChoicePreambleBitMap(params.valueExtensible, params.valueUpperBound); err != nil {
					return errors.Errorf("legacyParseField error: pd.getChoiceIndex Error")
				} else {
					present = int(index) + 1
				}
				val.Field(0).SetInt(int64(present))
				if present == 0 {
					return errors.Errorf("legacyParseField error: CHOICE present is 0(present's field number)")
				} else if present >= structType.NumField() {
					return errors.Errorf("legacyParseField error: CHOICE Present is bigger than number of struct field")
				} else {
					return legacyParseField(val.Field(present), pd, structParams[present])
				}
			}
		}

		// struct that is neither CHOICE nor OPEN TYPE (SEQUENCE)
		if err := pd.ReadSequencePreambleBitMap(&optionalPresentsFlag, params.valueExtensible); err != nil {
			return err
		} else {
			for i := 0; i < structType.NumField(); i++ {
				// optional fields
				if structParams[i].optional && handledOptionalCount < uint64(len(optionalPresentsFlag)) {
					if !optionalPresentsFlag[handledOptionalCount] {
						handledOptionalCount++
						continue
					} else {
						handledOptionalCount++
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
						return errors.Errorf("legacyParseField error:" +
							" Open type is not reference to the other field in the struct")
					}
					structParams[i].referenceFieldValue = new(int64)
					if referenceFieldValue, err := legacyGetReferenceFieldValue(val.Field(index)); err != nil {
						return err
					} else {
						*structParams[i].referenceFieldValue = referenceFieldValue
					}
				}
				if err := legacyParseField(val.Field(i), pd, structParams[i]); err != nil {
					return err
				}
			}
			return nil
		}
	case reflect.Slice:
		if numComponents, err := pd.ReadSequenceOfPreambleBitMap(params.sizeExtensible,
			params.sizeLowerBound, params.sizeUpperBound); err != nil {
			return err
		} else {
			sliceType := fieldType
			params.sizeExtensible = false
			params.sizeUpperBound = nil
			params.sizeLowerBound = nil
			sliceContent := reflect.MakeSlice(sliceType, int(numComponents), int(numComponents))
			for i := 0; i < int(numComponents); i++ {
				err := legacyParseField(sliceContent.Index(i), pd, params)
				if err != nil {
					return errors.Wrap(err, "legacy parse field failed")
				}
			}
			val.Set(sliceContent)
			return nil
		}
	case reflect.String:
		// Decoding PrintableString using Octet String decoding method

		if octetString, err := pd.ReadOctetString(params.sizeExtensible,
			params.sizeLowerBound, params.sizeUpperBound); err != nil {
			return err
		} else {
			printableString := string(octetString)
			val.SetString(printableString)
			return nil
		}
	}
	return errors.Errorf("legacyParseField error: unsupported: %s", v.Type().String())
}

// Unmarshal parses the APER-encoded ASN.1 data structure b
// and uses the reflect package to fill in an arbitrary value pointed at by value.
// Because Unmarshal uses the reflect package, the structs
// being written to must use upper case field names.
//
// An ASN.1 INTEGER can be written to an int, int32, int64,
// If the encoded value does not fit in the Go type,
// Unmarshal returns a parse error.
//
// An ASN.1 BIT STRING can be written to a BitString.
//
// An ASN.1 OCTET STRING can be written to a []byte.
//
// An ASN.1 OBJECT IDENTIFIER can be written to an
// ObjectIdentifier.
//
// An ASN.1 ENUMERATED can be written to an Enumerated.
//
// Any of the above ASN.1 values can be written to an interface{}.
// The value stored in the interface has the corresponding Go type.
// For integers, that type is int64.
//
// An ASN.1 SEQUENCE OF x can be written
// to a slice if an x can be written to the slice's element type.
//
// An ASN.1 SEQUENCE can be written to a struct
// if each of the elements in the sequence can be
// written to the corresponding element in the struct.
//
// The following tags on struct fields have special meaning to Unmarshal:
//
//	optional            OPTIONAL tag in SEQUENCE
//	sizeExt             specifies that size  is extensible
//	valueExt            specifies that value is extensible
//	sizeLB              set the minimum value of size constraint
//	sizeUB              set the maximum value of value constraint
//	valueLB             set the minimum value of size constraint
//	valueUB             set the maximum value of value constraint
//	default             sets the default value
//	openType            specifies the open Type
//	referenceFieldName  the string of the reference field for this type (only if openType used)
//	referenceFieldValue the corresponding value of the reference field for this type (only if openType used)
//
// Other ASN.1 types are not supported; if it encounters them,
// Unmarshal returns a parse error.
func LegacyUnmarshal(b []byte, value interface{}) error {
	return LegacyUnmarshalWithParams(b, value, "")
}

// UnmarshalWithParams allows field parameters to be specified for the
// top-level element. The form of the params is the same as the field tags.
func LegacyUnmarshalWithParams(b []byte, value interface{}, params string) error {
	v := reflect.ValueOf(value).Elem()
	pd := NewPerBitData(b)
	return legacyParseField(v, pd, parseFieldParameters(params))
}

/*** End of legacy function ***/
