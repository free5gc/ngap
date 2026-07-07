package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCauseRadioNetworkMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *CauseRadioNetwork
		expected []byte
	}{
		{
			name: "Case 1",
			input: &CauseRadioNetwork{
				Value: CauseRadioNetworkPresentUnspecified,
			},
			expected: []byte{0x00},
		},
		{
			name: "Case 2",
			input: &CauseRadioNetwork{
				Value: CauseRadioNetworkPresentReleaseDueToCnDetectedMobility,
			},
			expected: []byte{0x58},
		},
		{
			name: "Case 3",
			input: &CauseRadioNetwork{
				Value: CauseRadioNetworkPresentN26InterfaceNotAvailable,
			},
			expected: []byte{0x80},
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

func TestCauseRadioNetworkUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *CauseRadioNetwork
	}{
		{
			name:  "Case 1",
			input: []byte{0x00},
			expected: &CauseRadioNetwork{
				Value: CauseRadioNetworkPresentUnspecified,
			},
		},
		{
			name:  "Case 2",
			input: []byte{0x58},
			expected: &CauseRadioNetwork{
				Value: CauseRadioNetworkPresentReleaseDueToCnDetectedMobility,
			},
		},
		{
			name:  "Case 3",
			input: []byte{0x80},
			expected: &CauseRadioNetwork{
				Value: CauseRadioNetworkPresentN26InterfaceNotAvailable,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(CauseRadioNetwork)
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
