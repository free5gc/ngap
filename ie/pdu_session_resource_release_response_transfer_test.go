package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceReleaseResponseTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceReleaseResponseTransfer
		expected []byte
	}{
		{
			name:  "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionReleaseRequest",
			input: &PDUSessionResourceReleaseResponseTransfer{},
			expected: []byte{
				0x00,
			},
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

func TestPDUSessionResourceReleaseResponseTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceReleaseResponseTransfer
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionReleaseRequest",
			input: []byte{
				0x00,
			},
			expected: &PDUSessionResourceReleaseResponseTransfer{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(PDUSessionResourceReleaseResponseTransfer)
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
