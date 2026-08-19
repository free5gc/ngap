package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestPDUSessionResourceReleaseResponseMarshalBinary(t *testing.T) {
	t.Parallel()

	PDUSessionResourceReleaseResponseTransferOS := aper.OctetString{0x00}

	testCases := []struct {
		name     string
		input    *PDUSessionResourceReleaseResponse
		expected []byte
	}{
		{
			name: "Case 1",
			input: &PDUSessionResourceReleaseResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 140,
				},
				PDUSessionResourceReleasedListRelRes: &ie.PDUSessionResourceReleasedListRelRes{
					List: []ie.PDUSessionResourceReleasedItemRelRes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 1,
							},
							PDUSessionResourceReleaseResponseTransfer: &PDUSessionResourceReleaseResponseTransferOS,
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x1c, 0x00, 0x18, 0x00, 0x00, 0x03, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x02, 0x00, 0x55, 0x40, 0x02, 0x00,
				0x8c, 0x00, 0x46, 0x40, 0x05, 0x00, 0x00, 0x01, 0x01,
				0x00,
			},
		},
		{
			name: "Case 2",
			input: &PDUSessionResourceReleaseResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 966541003059,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceReleasedListRelRes: &ie.PDUSessionResourceReleasedListRelRes{
					List: []ie.PDUSessionResourceReleasedItemRelRes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceReleaseResponseTransferMarshalBinary
							PDUSessionResourceReleaseResponseTransfer: &aper.OctetString{0x00},
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x1c, 0x00, 0x1c, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xe1, 0x0a, 0x55, 0x49,
				0x33, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x46, 0x40, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x00,
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

func TestPDUSessionResourceReleaseResponseUnmarshalBinary(t *testing.T) {
	t.Parallel()

	PDUSessionResourceReleaseResponseTransferOS := aper.OctetString{0x00}

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceReleaseResponse
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x1c, 0x00, 0x18, 0x00, 0x00, 0x03, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x02, 0x00, 0x55, 0x40, 0x02, 0x00,
				0x8c, 0x00, 0x46, 0x40, 0x05, 0x00, 0x00, 0x01, 0x01,
				0x00,
			},
			expected: &PDUSessionResourceReleaseResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 140,
				},
				PDUSessionResourceReleasedListRelRes: &ie.PDUSessionResourceReleasedListRelRes{
					List: []ie.PDUSessionResourceReleasedItemRelRes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 1,
							},
							PDUSessionResourceReleaseResponseTransfer: &PDUSessionResourceReleaseResponseTransferOS,
						},
					},
				},
			},
		},
		{
			name: "Case 2",
			input: []byte{
				0x20, 0x1c, 0x00, 0x1c, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xe1, 0x0a, 0x55, 0x49,
				0x33, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x46, 0x40, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x00,
			},
			expected: &PDUSessionResourceReleaseResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 966541003059,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceReleasedListRelRes: &ie.PDUSessionResourceReleasedListRelRes{
					List: []ie.PDUSessionResourceReleasedItemRelRes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceReleaseResponseTransferUnmarshalBinary
							PDUSessionResourceReleaseResponseTransfer: &aper.OctetString{0x00},
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
