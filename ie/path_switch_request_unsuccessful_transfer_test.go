package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathSwitchRequestUnsuccessfulTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PathSwitchRequestUnsuccessfulTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic TestXnHandover",
			input: &PathSwitchRequestUnsuccessfulTransfer{
				Cause: &Cause{
					Choice: &CauseTransport{
						Value: CauseTransportPresentTransportResourceUnavailable,
					},
				},
			},
			expected: []byte{0x08},
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

func TestPathSwitchRequestUnsuccessfulTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PathSwitchRequestUnsuccessfulTransfer
	}{
		{
			name:  "Case 1: from ueranemu k8s-basic TestXnHandover",
			input: []byte{0x08},
			expected: &PathSwitchRequestUnsuccessfulTransfer{
				Cause: &Cause{
					Choice: &CauseTransport{
						Value: CauseTransportPresentTransportResourceUnavailable,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(PathSwitchRequestUnsuccessfulTransfer)
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
