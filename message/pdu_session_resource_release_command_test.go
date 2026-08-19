package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestPDUSessionResourceReleaseCommandMarshalBinary(t *testing.T) {
	t.Parallel()

	// pDUSessionResourceReleaseCommandTransfer
	pDUSessionResourceReleaseCommandTransfer := ie.PDUSessionResourceReleaseCommandTransfer{
		Cause: &ie.Cause{
			Choice: &ie.CauseNas{
				Value: ie.CauseNasPresentNormalRelease,
			},
		},
	}
	pDUSessionResourceReleaseCommandTransferBytes, err := ie.MarshalBinary(&pDUSessionResourceReleaseCommandTransfer)
	require.NoError(t, err)
	pDUSessionResourceReleaseCommandTransferOS := aper.OctetString(
		pDUSessionResourceReleaseCommandTransferBytes)

	testCases := []struct {
		name     string
		input    *PDUSessionResourceReleaseCommand
		expected []byte
	}{
		{
			name: "Case 1",
			input: &PDUSessionResourceReleaseCommand{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 140,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x02, 0xef, 0x95, 0x54, 0xaf, 0x03, 0x7e, 0x00, 0x68, 0x01, 0x00,
						0x05, 0x2e, 0x01, 0x10, 0xd3, 0x00, 0x12, 0x01,
					}),
				},
				PDUSessionResourceToReleaseListRelCmd: &ie.PDUSessionResourceToReleaseListRelCmd{
					List: []ie.PDUSessionResourceToReleaseItemRelCmd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 1,
							},
							PDUSessionResourceReleaseCommandTransfer: &pDUSessionResourceReleaseCommandTransferOS,
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x1c, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00,
				0x02, 0x00, 0x02, 0x00, 0x55, 0x00, 0x02, 0x00, 0x8c, 0x00,
				0x26, 0x40, 0x15, 0x14, 0x7e, 0x02, 0xef, 0x95, 0x54, 0xaf,
				0x03, 0x7e, 0x00, 0x68, 0x01, 0x00, 0x05, 0x2e, 0x01, 0x10,
				0xd3, 0x00, 0x12, 0x01, 0x00, 0x4f, 0x00, 0x05, 0x00, 0x00,
				0x01, 0x01, 0x10,
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestPDUSessionReleaseRequest",
			input: &PDUSessionResourceReleaseCommand{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 966541003059,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString{0x7e, 0x02, 0x6c, 0xed, 0xba, 0x34, 0x04, 0x7e, 0x00, 0x68, 0x01, 0x00, 0x05, 0x2e, 0x0a, 0x00, 0xd3, 0x24, 0x12, 0x0a},
				},
				PDUSessionResourceToReleaseListRelCmd: &ie.PDUSessionResourceToReleaseListRelCmd{
					List: []ie.PDUSessionResourceToReleaseItemRelCmd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceReleaseCommandTransferMarshalBinary
							PDUSessionResourceReleaseCommandTransfer: &aper.OctetString{0x10},
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x1c, 0x00, 0x35, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xe1, 0x0a, 0x55, 0x49,
				0x33, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x26, 0x40, 0x15, 0x14, 0x7e, 0x02, 0x6c, 0xed,
				0xba, 0x34, 0x04, 0x7e, 0x00, 0x68, 0x01, 0x00, 0x05, 0x2e, 0x0a, 0x00, 0xd3, 0x24, 0x12, 0x0a,
				0x00, 0x4f, 0x00, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x10,
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

func TestPDUSessionResourceReleaseCommandUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// pDUSessionResourceReleaseCommandTransfer
	pDUSessionResourceReleaseCommandTransfer := ie.PDUSessionResourceReleaseCommandTransfer{
		Cause: &ie.Cause{
			Choice: &ie.CauseNas{
				Value: ie.CauseNasPresentNormalRelease,
			},
		},
	}
	pDUSessionResourceReleaseCommandTransferBytes, err := ie.MarshalBinary(&pDUSessionResourceReleaseCommandTransfer)
	require.NoError(t, err)
	pDUSessionResourceReleaseCommandTransferOS := aper.OctetString(
		pDUSessionResourceReleaseCommandTransferBytes)

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceReleaseCommand
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x1c, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00,
				0x02, 0x00, 0x02, 0x00, 0x55, 0x00, 0x02, 0x00, 0x8c, 0x00,
				0x26, 0x40, 0x15, 0x14, 0x7e, 0x02, 0xef, 0x95, 0x54, 0xaf,
				0x03, 0x7e, 0x00, 0x68, 0x01, 0x00, 0x05, 0x2e, 0x01, 0x10,
				0xd3, 0x00, 0x12, 0x01, 0x00, 0x4f, 0x00, 0x05, 0x00, 0x00,
				0x01, 0x01, 0x10,
			},
			expected: &PDUSessionResourceReleaseCommand{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 140,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x02, 0xef, 0x95, 0x54, 0xaf, 0x03, 0x7e, 0x00, 0x68, 0x01, 0x00,
						0x05, 0x2e, 0x01, 0x10, 0xd3, 0x00, 0x12, 0x01,
					}),
				},
				PDUSessionResourceToReleaseListRelCmd: &ie.PDUSessionResourceToReleaseListRelCmd{
					List: []ie.PDUSessionResourceToReleaseItemRelCmd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 1,
							},
							PDUSessionResourceReleaseCommandTransfer: &pDUSessionResourceReleaseCommandTransferOS,
						},
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestPDUSessionReleaseRequest",
			input: []byte{
				0x00, 0x1c, 0x00, 0x35, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xe1, 0x0a, 0x55, 0x49,
				0x33, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x26, 0x40, 0x15, 0x14, 0x7e, 0x02, 0x6c, 0xed,
				0xba, 0x34, 0x04, 0x7e, 0x00, 0x68, 0x01, 0x00, 0x05, 0x2e, 0x0a, 0x00, 0xd3, 0x24, 0x12, 0x0a,
				0x00, 0x4f, 0x00, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x10,
			},
			expected: &PDUSessionResourceReleaseCommand{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 966541003059,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString{0x7e, 0x02, 0x6c, 0xed, 0xba, 0x34, 0x04, 0x7e, 0x00, 0x68, 0x01, 0x00, 0x05, 0x2e, 0x0a, 0x00, 0xd3, 0x24, 0x12, 0x0a},
				},
				PDUSessionResourceToReleaseListRelCmd: &ie.PDUSessionResourceToReleaseListRelCmd{
					List: []ie.PDUSessionResourceToReleaseItemRelCmd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceReleaseCommandTransferUnmarshalBinary
							PDUSessionResourceReleaseCommandTransfer: &aper.OctetString{0x10},
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
