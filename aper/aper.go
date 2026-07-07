package aper

import (
	"encoding/binary"
	"math"
	"reflect"

	"github.com/pkg/errors"
)

type PerBitData struct {
	bytes      []byte
	byteOffset uint64
	bitsOffset uint
}

func (pd *PerBitData) Bytes() []byte {
	// It is possible that the value is not encoded (e.g. definite-value integer),
	// and pd.bytes is still nil.
	// In this case, return 0x00 other than nil.
	if len(pd.bytes) == 0 {
		return make([]byte, 1)
	}
	return pd.bytes
}

func (pd *PerBitData) ByteOffset() uint64 {
	return pd.byteOffset
}

func (pd *PerBitData) BitOffset() uint {
	return pd.bitsOffset
}

func NewPerBitData(b []byte) *PerBitData {
	if b != nil {
		return &PerBitData{b, 0, 0}
	}
	return &PerBitData{[]byte(""), 0, 0}
}

func (pd *PerBitData) bitCarry() {
	pd.byteOffset += uint64(pd.bitsOffset >> 3)
	pd.bitsOffset = pd.bitsOffset & 0x07
}

/*** Common Utility functions ***/
func getRequiredNumBits(val int64, is2Complement bool) uint64 {
	// 2's complement: 1 byte accommadates -128~+127 (one bit spared for sign bit)
	// non-negative binary: 1 byte accommadates 0~255

	// turn into non-negative binary case first for simplicity
	if val < 0 {
		val += 1
		val *= -1 // negative value is not allowed for log
	}

	if val == 0 { // log2(0) = inf
		return 1 // val = 0 or -1 => 1 bit
	}
	numBits := uint64(math.Floor(math.Log2(float64(val)) + 1))

	if is2Complement {
		numBits += 1 // 2's complement requires extra sign bit
	}
	return numBits
}

func getRequiredNumBitsForUint64(val uint64) uint64 {
	// must not be 2's complement
	numBits := uint64(math.Log2(float64(val))) + 1
	return numBits
}

func getRequiredNumBytes(val int64, is2sComplement bool) uint64 {
	numBits := getRequiredNumBits(val, is2sComplement)
	numBytes := uint64(math.Ceil(float64(numBits) / 8))
	return numBytes
}

func getRequiredNumBytesForUint64(val uint64) uint64 {
	// must not be 2's complement
	numBits := getRequiredNumBitsForUint64(val)
	numBytes := uint64(math.Ceil(float64(numBits) / 8))
	return numBytes
}

// ShiftLeft performs a left bit shift operation on the provided bytes.
// If the bits count is negative, a right bit shift is performed.
func shiftLeft(dataPtr *[]byte, bits int) {
	n := len(*dataPtr)
	if bits < 0 {
		// shift right
		bits = -bits
		for i := n - 1; i > 0; i-- {
			(*dataPtr)[i] = (*dataPtr)[i]>>bits | (*dataPtr)[i-1]<<(8-bits)
		}
		(*dataPtr)[0] >>= bits
	} else {
		// shift left
		for i := 0; i < n-1; i++ {
			(*dataPtr)[i] = (*dataPtr)[i]<<bits | (*dataPtr)[i+1]>>(8-bits)
		}
		(*dataPtr)[n-1] <<= bits
	}
}

func bytesToInt64(bytes []byte, is2sComplement bool) (int64, error) {
	if len(bytes) > 8 {
		return 0, errors.Errorf("Bytes value overflows int64")
	} else if len(bytes) == 0 {
		return 0, errors.Errorf("Bytes length is zero")
	}
	var val int64 = 0
	if is2sComplement && bytes[0]>>7 == 1 { // first bit = 1 -> negative value
		val = -1 // ffff ffff ffff ffff
	}
	for i := 0; i < len(bytes); i++ {
		val <<= 8
		val |= int64(bytes[i])
	}
	return val, nil
}

func bytesToUint64(bytes []byte) (uint64, error) {
	if len(bytes) > 8 {
		return 0, errors.Errorf("bytesToUint64 error: bytes value overflows uint64")
	} else if len(bytes) == 0 {
		return 0, errors.Errorf("bytesToUint64 error: bytes length is zero")
	}

	// pad bytes to 8 bytes
	nBytesPad := 8 - len(bytes)
	for nBytesPad > 0 {
		bytes = append([]byte{byte(0)}, bytes...)
		nBytesPad -= 1
	}

	return binary.BigEndian.Uint64(bytes), nil
}

func int64ToBytes(val int64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(val))
	return bytes
}

func uint64ToBytes(val uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, val)
	return bytes
}

/*** End of common utility functions ***/

/*** Common Legacy fuctions ***/
func legacyGetReferenceFieldValue(v reflect.Value) (value int64, err error) {
	fieldType := v.Type()
	switch v.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		value = v.Int()
	case reflect.Struct:
		if fieldType.Field(0).Name == "Present" {
			present := int(v.Field(0).Int())
			if present == 0 {
				err = errors.Errorf("legacyGetReferenceFieldValue error: " +
					"ReferenceField Value present is 0 (present's field number)")
			} else if present >= fieldType.NumField() {
				err = errors.Errorf("legacyGetReferenceFieldValue error: " +
					"Present is bigger than number of struct field")
			} else {
				value, err = legacyGetReferenceFieldValue(v.Field(present))
			}
		} else {
			value, err = legacyGetReferenceFieldValue(v.Field(0))
		}
	default:
		err = errors.Errorf("legacyGetReferenceFieldValue error: " +
			"OpenType reference only support INTEGER")
	}
	return
}

/*** End of common legacy function ***/
