package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/ie"
)

func TestHandoverFailureMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *HandoverFailure
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestHandoverFailure",
			input: &HandoverFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1009452830939,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem,
					},
				},
			},
			expected: []byte{
				0x40, 0x0d, 0x00, 0x13, 0x00, 0x00, 0x02, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xeb, 0x08, 0x13, 0xd0,
				0xdb, 0x00, 0x0f, 0x40, 0x02, 0x01, 0xc0,
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

func TestHandoverFailureUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverFailure
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestHandoverFailure",
			input: []byte{
				0x40, 0x0d, 0x00, 0x13, 0x00, 0x00, 0x02, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xeb, 0x08, 0x13, 0xd0,
				0xdb, 0x00, 0x0f, 0x40, 0x02, 0x01, 0xc0,
			},
			expected: &HandoverFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1009452830939,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem,
					},
				},
			},
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
