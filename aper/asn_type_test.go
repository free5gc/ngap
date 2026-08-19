package aper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBitStringFromHex(t *testing.T) {
	// Case 1
	bs := BitString{}
	err := bs.FromHex("12", 6)
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0x48}, 6}, bs)

	// Case 2
	bs = BitString{}
	err = bs.FromHex("12", 8)
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0x12}, 8}, bs)
}

func TestBitStringToHex(t *testing.T) {
	// Case 1
	bs := BitString{
		Bytes:     []byte{0xAB, 0xAB},
		BitLength: 16,
	}
	hexString, err := bs.ToHex()
	require.NoError(t, err)
	require.Equal(t, "abab", hexString)

	// Case 2
	bs = BitString{
		Bytes:     []byte{0xAB, 0xAB},
		BitLength: 5,
	}
	hexString, err = bs.ToHex()
	require.NoError(t, err)
	require.Equal(t, "15", hexString)

	// Case 3
	bs = BitString{
		Bytes:     []byte{0xAB, 0xAB},
		BitLength: 24,
	}
	hexString, err = bs.ToHex()
	require.Error(t, err)
	require.Equal(t, "", hexString)
}

func TestBitStringToBytes(t *testing.T) {
	// Case 1
	bs := BitString{
		Bytes:     []byte{0x18, 0x23},
		BitLength: 11,
	}
	bytes, err := bs.ToBytes()
	require.NoError(t, err)
	require.Equal(t, []byte{0x00, 0xC1}, bytes)

	// Case 2
	bs = BitString{
		Bytes:     []byte{0x18, 0x23},
		BitLength: 8,
	}
	bytes, err = bs.ToBytes()
	require.NoError(t, err)
	require.Equal(t, []byte{0x18}, bytes)

	bs = BitString{
		Bytes:     []byte{0x18, 0x23},
		BitLength: 16,
	}
	bytes, err = bs.ToBytes()
	require.NoError(t, err)
	require.Equal(t, []byte{0x18, 0x23}, bytes)

	bs = BitString{
		Bytes:     []byte{0x18, 0x23},
		BitLength: 17,
	}
	bytes, err = bs.ToBytes()
	require.Error(t, err)
	require.Equal(t, []byte{}, bytes)
}

func TestBitStringFromString(t *testing.T) {
	// Case 1
	bs := BitString{}
	err := bs.FromBits("1010")
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0xA0}, 4}, bs)

	// Case 2
	bs = BitString{}
	err = bs.FromBits("10101010")
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0xAA}, 8}, bs)

	// Case 3
	bs = BitString{}
	err = bs.FromBits("1010101011")
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0xAA, 0xC0}, 10}, bs)
}

func TestBitStringToString(t *testing.T) {
	// Case 1
	bs := BitString{
		Bytes:     []byte{0xA0},
		BitLength: 4,
	}
	s, err := bs.ToBits()
	require.NoError(t, err)
	require.Equal(t, "1010", s)

	// Case 2
	bs = BitString{
		Bytes:     []byte{0xAA},
		BitLength: 8,
	}
	s, err = bs.ToBits()
	require.NoError(t, err)
	require.Equal(t, "10101010", s)

	// Case 3
	bs = BitString{
		Bytes:     []byte{0xAA, 0xC0},
		BitLength: 10,
	}
	s, err = bs.ToBits()
	require.NoError(t, err)
	require.Equal(t, "1010101011", s)
}

func TestBitStringFromUint(t *testing.T) {
	// Case 1
	bs := BitString{}
	err := bs.FromUint(1, 1)
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0x80}, 1}, bs)

	// Case 2
	bs = BitString{}
	err = bs.FromUint(1, 3)
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0x20}, 3}, bs)

	// Case 3
	bs = BitString{}
	err = bs.FromUint(123456, -1)
	require.NoError(t, err)
	require.Equal(t, BitString{[]byte{0xF1, 0x20, 0x00}, 17}, bs)

	// Case 4: invalid input
	bs = BitString{}
	err = bs.FromUint(8, 1)
	require.Error(t, err)
	require.Equal(t, BitString{}, bs)
}

func TestBitStringToUint(t *testing.T) {
	// Case 1
	bs := BitString{[]byte{0x80}, 1}
	val, err := bs.ToUint()
	require.NoError(t, err)
	require.Equal(t, uint64(1), val)

	// Case 2
	bs = BitString{[]byte{0xF1, 0x20, 0x00}, 17}
	val, err = bs.ToUint()
	require.NoError(t, err)
	require.Equal(t, uint64(123456), val)
}
