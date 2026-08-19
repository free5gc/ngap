package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceReleaseCommandTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceReleaseCommandTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionReleaseRequest",
			input: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{
						Value: CauseNasPresentNormalRelease,
					},
				},
			},
			expected: []byte{
				0x10,
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

func TestPDUSessionResourceReleaseCommandTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceReleaseCommandTransfer
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionReleaseRequest",
			input: []byte{
				0x10,
			},
			expected: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{
						Value: CauseNasPresentNormalRelease,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(PDUSessionResourceReleaseCommandTransfer)
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
