package message

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceModifyResponseMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceModifyResponse
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionModification",
			input: &PDUSessionResourceModifyResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 665250635458,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceModifyListModRes: &ie.PDUSessionResourceModifyListModRes{
					List: []ie.PDUSessionResourceModifyItemModRes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceModifyResponseTransferMarshalBinary
							PDUSessionResourceModifyResponseTransfer: &aper.OctetString{0x10, 0x00, 0x08},
						},
					},
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x20},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x11},
							},
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x1a, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x40, 0x06, 0x80, 0x9a, 0xe4, 0x07, 0x1e,
				0xc2, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x41, 0x40, 0x07, 0x00, 0x00, 0x0a, 0x03, 0x10,
				0x00, 0x08, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00, 0x20, 0x02,
				0xf8, 0x39, 0x00, 0x00, 0x11,
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

func TestPDUSessionResourceModifyResponseUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceModifyResponse
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionModification",
			input: []byte{
				0x20, 0x1a, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x40, 0x06, 0x80, 0x9a, 0xe4, 0x07, 0x1e,
				0xc2, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x41, 0x40, 0x07, 0x00, 0x00, 0x0a, 0x03, 0x10,
				0x00, 0x08, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00, 0x20, 0x02,
				0xf8, 0x39, 0x00, 0x00, 0x11,
			},
			expected: &PDUSessionResourceModifyResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 665250635458,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceModifyListModRes: &ie.PDUSessionResourceModifyListModRes{
					List: []ie.PDUSessionResourceModifyItemModRes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceModifyResponseTransferUnmarshalBinary
							PDUSessionResourceModifyResponseTransfer: &aper.OctetString{0x10, 0x00, 0x08},
						},
					},
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x20},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x11},
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
