package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/ie"
)

func TestNGResetMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *NGReset
		expected []byte
	}{
		{
			name: "Case 1",
			input: &NGReset{
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason,
					},
				},
				ResetType: &ie.ResetType{
					Choice: &ie.UEAssociatedLogicalNGConnectionList{
						List: []ie.UEAssociatedLogicalNGConnectionItem{
							{
								RANUENGAPID: &ie.RANUENGAPID{
									Value: 94,
								},
							},
							{
								RANUENGAPID: &ie.RANUENGAPID{
									Value: 95,
								},
							},
							{
								RANUENGAPID: &ie.RANUENGAPID{
									Value: 296,
								},
							},
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x14, 0x00, 0x16, 0x00, 0x00, 0x02, 0x00, 0x0f, 0x40,
				0x02, 0x00, 0xc0, 0x00, 0x58, 0x00, 0x09, 0x40, 0x03, 0x20,
				0x5e, 0x20, 0x5f, 0x24, 0x01, 0x28,
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

func TestNGResetUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *NGReset
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x14, 0x00, 0x16, 0x00, 0x00, 0x02, 0x00, 0x0f, 0x40,
				0x02, 0x00, 0xc0, 0x00, 0x58, 0x00, 0x09, 0x40, 0x03, 0x20,
				0x5e, 0x20, 0x5f, 0x24, 0x01, 0x28,
			},
			expected: &NGReset{
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason,
					},
				},
				ResetType: &ie.ResetType{
					Choice: &ie.UEAssociatedLogicalNGConnectionList{
						List: []ie.UEAssociatedLogicalNGConnectionItem{
							{
								RANUENGAPID: &ie.RANUENGAPID{
									Value: 94,
								},
							},
							{
								RANUENGAPID: &ie.RANUENGAPID{
									Value: 95,
								},
							},
							{
								RANUENGAPID: &ie.RANUENGAPID{
									Value: 296,
								},
							},
						},
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
