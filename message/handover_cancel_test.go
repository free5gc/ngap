package message

import (
	"testing"

	"github.com/free5gc/ngap/ie"
	"github.com/stretchr/testify/require"
)

func TestHandoverCancelMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *HandoverCancel
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestHandoverCancel",
			input: &HandoverCancel{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 776276781756,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHandoverCancelled,
					},
				},
			},
			expected: []byte{
				0x00, 0x0a, 0x00, 0x19, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xb4, 0xbd, 0xb3, 0xaa,
				0xbc, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x01, 0x40,
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

func TestHandoverCancelUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverCancel
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestHandoverCancel",
			input: []byte{
				0x00, 0x0a, 0x00, 0x19, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xb4, 0xbd, 0xb3, 0xaa,
				0xbc, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x01, 0x40,
			},
			expected: &HandoverCancel{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 776276781756,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHandoverCancelled,
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
