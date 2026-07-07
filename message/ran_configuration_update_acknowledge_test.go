package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRANConfigurationUpdateAcknowledgeMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *RANConfigurationUpdateAcknowledge
		expected []byte
	}{
		{
			name:  "Case 1",
			input: &RANConfigurationUpdateAcknowledge{},
			expected: []byte{
				0x20, 0x23, 0x00, 0x03, 0x00, 0x00, 0x00,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.input.MarshalBinary()
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, b)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestRANConfigurationUpdateAcknowledgeUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *RANConfigurationUpdateAcknowledge
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x23, 0x00, 0x03, 0x00, 0x00, 0x00,
			},
			expected: &RANConfigurationUpdateAcknowledge{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg, err := Parse(tc.input)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, msg)
			} else {
				require.Error(t, err)
			}
		})
	}
}
