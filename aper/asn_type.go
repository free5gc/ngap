package aper

import (
	"encoding/binary"
	"encoding/hex"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// BIT STRING

// BitString is for an ASN.1 BIT STRING type, BitLength means the effective bits.
type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength uint64 // length in bits.
}

// Format BitString from HexString
// hexString of odd digits is possible here since bitLength may not be a multiple of 8
// hexString of bitLength larger than itself will be prepended with zeros
func (bitString *BitString) FromHex(hexString string, bitLength int) error {
	if len(hexString)%2 != 0 {
		// if hexstring is odd length, prepend a zero
		hexString = "0" + hexString
	}

	// convert to bitString
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return errors.Wrap(err, "decode hexstring failed")
	}

	// check if hexstring value requires more bits than bitLength
	// find number of leading zero-bits
	numLeadingZeroBits := 0
	mask := byte(0x80)
	for i := 0; i < len(bytes); {
		if bytes[i]&mask == 0 {
			numLeadingZeroBits++
			mask >>= 1
			if mask == 0 {
				mask = 0x80
				i++
			}
		} else {
			break
		}
	}
	if len(bytes)*8-numLeadingZeroBits > bitLength {
		return errors.Errorf("Build BitString from hexString failed: "+
			"hexString value requires more than %d bits", bitLength)
	}

	// prepend 0x00 if required bitLength is longer than the converted byte array size
	prependedZerosNum := (bitLength+7)/8 - len(bytes)
	for i := 0; i < prependedZerosNum; i++ {
		bytes = append([]byte{0x00}, bytes...)
	}

	err = bitString.FromBytes(bytes, len(bytes)*8-bitLength, bitLength)
	if err != nil {
		return errors.Wrap(err, "build BitString from hexString failed")
	}
	return nil
}

func (bitString *BitString) ToHex() (string, error) {
	bytes, err := bitString.ToBytes()
	if err != nil {
		return "", errors.Wrap(err, "convert bitString to hexString failed")
	}
	hexString := hex.EncodeToString(bytes)
	// remove extra hexstring digit according to bitLength
	numDigit := (bitString.BitLength + 3) / 4
	hexString = hexString[len(hexString)-int(numDigit):]
	return hexString, nil
}

// Format BitString from Bytes
// Notice: resulting byteArray will have no redundant bit value and byte array outside the valid bitLength
// e.g. FromBytes([]byte{0xFF, 0xFF}, 0, 1) -> bitString{Bytes: []byte{0x80}, 1}
func (bitString *BitString) FromBytes(byteArray []byte, bitsOffset int, bitLength int) error {
	if bitsOffset < 0 {
		return errors.Errorf("Negative bitsOffset %d to build bitString from bytes",
			bitsOffset)
	}
	// bitsOffset >= 1 byte => shift srcBytes and calculate new bitsOffset
	if bitsOffset > 7 {
		byteArray = byteArray[bitsOffset/8:]
		bitsOffset %= 8
	}

	bitsLeft := len(byteArray)*8 - bitsOffset
	if bitLength > bitsLeft {
		return errors.Errorf("Require bitLength is not available, requireBits: %d, leftBits: %d",
			bitLength, bitsLeft)
	}
	byteLen := (bitsOffset + bitLength + 7) >> 3
	numBitsByteLen := (bitLength + 7) >> 3
	dstBytes := make([]byte, numBitsByteLen)
	numBitsMask := byte(0xff)
	if modEight := bitLength & 0x7; modEight != 0 {
		numBitsMask <<= uint8(8 - (modEight))
	}
	for i := 1; i < byteLen; i++ {
		dstBytes[i-1] = byteArray[i-1]<<bitsOffset | byteArray[i]>>(8-bitsOffset)
	}
	if byteLen == numBitsByteLen {
		dstBytes[byteLen-1] = byteArray[byteLen-1] << bitsOffset
	}
	dstBytes[numBitsByteLen-1] &= numBitsMask

	bitString.Bytes = dstBytes
	bitString.BitLength = uint64(bitLength)
	return nil
}

func (bitString *BitString) ToBytes() ([]byte, error) {
	dstByteLen := (bitString.BitLength + 7) / 8
	if int(dstByteLen) > len(bitString.Bytes) {
		err := errors.Errorf("Invalid bitString: bitLength is longer than the bytes length")
		return []byte{}, err
	}
	dstBytes := make([]byte, len(bitString.Bytes))
	copy(dstBytes, bitString.Bytes)
	// // check if there are values left after bitLength
	// for i := dstByteLen - 1; i < uint64(len(dstBytes)); i++ {
	// 	if i == dstByteLen-1 {
	// 		if bitString.BitLength%8 == 0 {
	// 			// No bit left after bitLength at the last byte of dstBytes
	// 			continue
	// 		}
	// 		tmpMask := byte(0xff) >> (bitString.BitLength % 8)
	// 		if dstBytes[i]&tmpMask != 0 {
	// 			err := errors.Errorf("Invalid bitString: there are value left after BitLength in Bytes")
	// 			return []byte{}, err
	// 		}
	// 	} else {
	// 		if dstBytes[i] != 0 {
	// 			err := errors.Errorf("Invalid bitString: there are value left after BitLength in Bytes")
	// 			return []byte{}, err
	// 		}
	// 	}
	// }
	// remove eliminated bytes based on bitString.BitLength
	dstBytes = dstBytes[:dstByteLen]

	// get number of bits to be shifted right
	if bitString.BitLength%8 != 0 {
		bitsLeft := 8 - (bitString.BitLength % 8)
		for i := dstByteLen - 1; i > 0; i-- {
			dstBytes[i] = dstBytes[i-1]<<byte(8-bitsLeft) | dstBytes[i]>>byte(bitsLeft)
		}
		dstBytes[0] = dstBytes[0] >> byte(bitsLeft)
	}

	return dstBytes, nil
}

// e.g. "1011" -> BitString(0xb0, 4)
func (bitString *BitString) FromBits(bits string) error {
	bitString.Bytes = []byte{}
	bitString.BitLength = uint64(len(bits))

	i := 0
	for i < len(bits) {
		// last byte: append trailing 0s
		if (i + 8) > len(bits) {
			for (i + 8) > len(bits) {
				bits = bits + "0"
			}
		}

		val, err := strconv.ParseUint(bits[i:i+8], 2, 8)
		if err != nil {
			return errors.Wrap(err, "convert string to bitString failed")
		}
		bitString.Bytes = append(bitString.Bytes, byte(val))

		i += 8
	}

	return nil
}

func (bitString *BitString) ToBits() (string, error) {
	// check whether the bitstring is valid
	dstByteLen := (bitString.BitLength + 7) / 8
	if int(dstByteLen) > len(bitString.Bytes) {
		err := errors.Errorf("Invalid bitString: bitLength is longer than the bytes length")
		return "", err
	}

	s := ""
	for i := 0; i < len(bitString.Bytes); i++ {
		for j := 0; j < 8; j++ {
			mask := byte(1 << (7 - j))
			if bitString.Bytes[i]&mask > 0 {
				s += "1"
			} else {
				s += "0"
			}
		}
	}
	return s[:bitString.BitLength], nil
}

func (bitString *BitString) FromUint(val uint64, numBits int) error {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, val)

	bitLength := int(getRequiredNumBitsForUint64(val))
	if numBits < 0 {
		numBits = bitLength
	}
	if numBits < bitLength {
		return errors.Errorf("Build BitString from uint64 failed: "+
			"provided numBits %d is too short to accommodate uint64 value %d", numBits, val)
	}
	err := bitString.FromBytes(bytes, len(bytes)*8-numBits, numBits)
	if err != nil {
		return errors.Wrap(err, "build bitString from uint64 failed")
	}

	return nil
}

func (bitString *BitString) ToUint() (uint64, error) {
	s, err := bitString.ToBits()
	if err != nil {
		return 0, errors.Wrap(err, "convert BitString to uint64 failed")
	}
	u, err := strconv.ParseUint(s, 2, len(s))
	if err != nil {
		return 0, errors.Wrap(err, "convert BitString to uint64 failed")
	}

	return u, nil
}

// OCTET STRING

// OctetString is for an ASN.1 OCTET STRING type
type OctetString []byte

func (octetString *OctetString) FromHex(hexString string) error {
	if len(hexString)%2 != 0 {
		// if hexstring is odd length, prepend a zero
		hexString = "0" + hexString
	}
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return errors.Wrap(err, "build octetString from hexString failed")
	}
	*octetString = OctetString(decoded)
	return nil
}

func (octetString *OctetString) ToHex() string {
	return hex.EncodeToString(*octetString)
}

// Get byte array for a string ('a' -> 0x61, 'A' -> 0x41 ...)
func (octetString *OctetString) FromString(s string) {
	*octetString = []byte(s)
}

func (octetString *OctetString) ToString() string {
	return string(*octetString)
}

// OBJECT IDENTIFIER

// ObjectIdentifier is for an ASN.1 OBJECT IDENTIFIER type
type ObjectIdentifier []byte

// ENUMERATED

// An Enumerated is represented as a plain uint64.
type Enumerated uint64

var (
	// BitStringType is the type of BitString
	BitStringType = reflect.TypeOf(BitString{})
	// OctetStringType is the type of OctetString
	OctetStringType = reflect.TypeOf(OctetString{})
	// ObjectIdentifierType is the type of ObjectIdentify
	ObjectIdentifierType = reflect.TypeOf(ObjectIdentifier{})
	// EnumeratedType is the type of Enumerated
	EnumeratedType = reflect.TypeOf(Enumerated(0))
)

// Characters
type (
	PrintableString string
	VisibleString   string
	UTF8String      string
)

// NULL
type NULL struct{}
