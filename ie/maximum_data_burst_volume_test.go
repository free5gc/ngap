package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumDataBurstVolumeMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *MaximumDataBurstVolume
		expected []byte
	}{
		{
			name: "Case 1",
			input: &MaximumDataBurstVolume{
				Value: 2000,
			},
			expected: []byte{0x00, 0x07, 0xD0},
		},
		{
			name: "Case 2",
			input: &MaximumDataBurstVolume{
				Value: 4095,
			},
			expected: []byte{0x00, 0x0F, 0xFF},
		},
		{
			name: "Case 3",
			input: &MaximumDataBurstVolume{
				Value: 4096,
			},
			expected: []byte{0x80, 0x02, 0x10, 0x00},
		},
		{
			name: "Case 4",
			input: &MaximumDataBurstVolume{
				Value: 2000000,
			},
			expected: []byte{0x80, 0x03, 0x1E, 0x84, 0x80},
		},
		{
			name: "Case 5",
			input: &MaximumDataBurstVolume{
				Value: 2000001,
			},
			expected: []byte{0x80, 0x03, 0x1E, 0x84, 0x81},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := MarshalBinary(tc.input)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, b)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMaximumDataBurstVolumeUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *MaximumDataBurstVolume
	}{
		{
			name:  "Case 1",
			input: []byte{0x00, 0x07, 0xD0},
			expected: &MaximumDataBurstVolume{
				Value: 2000,
			},
		},
		{
			name:  "Case 2",
			input: []byte{0x00, 0x0F, 0xFF},
			expected: &MaximumDataBurstVolume{
				Value: 4095,
			},
		},
		{
			name:  "Case 3",
			input: []byte{0x80, 0x02, 0x10, 0x00},
			expected: &MaximumDataBurstVolume{
				Value: 4096,
			},
		},
		{
			name:  "Case 4",
			input: []byte{0x80, 0x03, 0x1E, 0x84, 0x80},
			expected: &MaximumDataBurstVolume{
				Value: 2000000,
			},
		},
		{
			name:  "Case 5",
			input: []byte{0x80, 0x03, 0x1E, 0x84, 0x81},
			expected: &MaximumDataBurstVolume{
				Value: 2000001,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(MaximumDataBurstVolume)
			err := UnmarshalBinary(tc.input, ie)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, ie)
			} else {
				require.Error(t, err)
			}
		})
	}
}
