package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/ie"
)

func TestHandoverPreparationFailureMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *HandoverPreparationFailure
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestHandoverFailure",
			input: &HandoverPreparationFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 742710891979,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem,
					},
				},
			},
			expected: []byte{
				0x40, 0x0c, 0x00, 0x19, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xac, 0xed, 0x04, 0xd5,
				0xcb, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x01, 0xc0,
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

func TestHandoverPreparationFailureUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverPreparationFailure
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestHandoverFailure",
			input: []byte{
				0x40, 0x0c, 0x00, 0x19, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xac, 0xed, 0x04, 0xd5,
				0xcb, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x01, 0xc0,
			},
			expected: &HandoverPreparationFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 742710891979,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
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
