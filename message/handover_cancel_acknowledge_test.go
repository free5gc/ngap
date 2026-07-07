package message

import (
	"testing"

	"github.com/free5gc/ngap/ie"
	"github.com/stretchr/testify/require"
)

func TestHandoverCancelAcknowledgeMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *HandoverCancelAcknowledge
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestHandoverCancel",
			input: &HandoverCancelAcknowledge{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 776276781756,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
			},
			expected: []byte{
				0x20, 0x0a, 0x00, 0x13, 0x00, 0x00, 0x02, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xb4, 0xbd, 0xb3, 0xaa,
				0xbc, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01,
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

func TestHandoverCancelAcknowledgeUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverCancelAcknowledge
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestHandoverCancel",
			input: []byte{
				0x20, 0x0a, 0x00, 0x13, 0x00, 0x00, 0x02, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xb4, 0xbd, 0xb3, 0xaa,
				0xbc, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01,
			},
			expected: &HandoverCancelAcknowledge{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 776276781756,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
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
