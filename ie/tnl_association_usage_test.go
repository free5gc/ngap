package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTNLAssociationUsageMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *TNLAssociationUsage
		expected []byte
	}{
		{
			name: "Case 1",
			input: &TNLAssociationUsage{
				Value: TNLAssociationUsagePresentUe,
			},
			expected: []byte{0x00},
		},
		{
			name: "Case 2",
			input: &TNLAssociationUsage{
				Value: TNLAssociationUsagePresentNonUe,
			},
			expected: []byte{0x20},
		},
		{
			name: "Case 3",
			input: &TNLAssociationUsage{
				Value: TNLAssociationUsagePresentBoth,
			},
			expected: []byte{0x40},
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

func TestTNLAssociationUsageUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *TNLAssociationUsage
	}{
		{
			name:  "Case 1",
			input: []byte{0x00},
			expected: &TNLAssociationUsage{
				Value: TNLAssociationUsagePresentUe,
			},
		},
		{
			name:  "Case 2",
			input: []byte{0x20},
			expected: &TNLAssociationUsage{
				Value: TNLAssociationUsagePresentNonUe,
			},
		},
		{
			name:  "Case 3",
			input: []byte{0x40},
			expected: &TNLAssociationUsage{
				Value: TNLAssociationUsagePresentBoth,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(TNLAssociationUsage)
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
