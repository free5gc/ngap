package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/ie"
)

func TestErrorIndicationMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *ErrorIndication
		expected []byte
	}{
		{
			name: "Case 1",
			input: &ErrorIndication{
				Cause: &ie.Cause{
					Choice: &ie.CauseProtocol{
						Value: ie.CauseProtocolPresentMessageNotCompatibleWithReceiverState,
					},
				},
			},
			expected: []byte{
				0x00, 0x09, 0x40, 0x08, 0x00, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x01, 0x66,
			},
		},
		{
			name: "Case 2",
			input: &ErrorIndication{
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentUnknownLocalUENGAPID,
					},
				},
			},
			expected: []byte{
				0x00, 0x09, 0x40, 0x09, 0x00, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x03, 0x80,
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

func TestErrorIndicationUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *ErrorIndication
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x09, 0x40, 0x08, 0x00, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x01, 0x66,
			},
			expected: &ErrorIndication{
				Cause: &ie.Cause{
					Choice: &ie.CauseProtocol{
						Value: ie.CauseProtocolPresentMessageNotCompatibleWithReceiverState,
					},
				},
			},
		},
		{
			name: "Case 2",
			input: []byte{
				0x00, 0x09, 0x40, 0x09, 0x00, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x03, 0x80,
			},
			expected: &ErrorIndication{
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentUnknownLocalUENGAPID,
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
