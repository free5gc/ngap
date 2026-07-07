package message

import (
	"testing"

	"github.com/free5gc/ngap/ie"
	"github.com/stretchr/testify/require"
)

func TestUplinkRANStatusTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	intPtr := func(i int64) *int64 {
		return &i
	}

	testCases := []struct {
		name     string
		input    *UplinkRANStatusTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover",
			input: &UplinkRANStatusTransfer{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 910057611806,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				RANStatusTransferTransparentContainer: &ie.RANStatusTransferTransparentContainer{
					DRBsSubjectToStatusTransferList: &ie.DRBsSubjectToStatusTransferList{
						List: []ie.DRBsSubjectToStatusTransferItem{
							{
								DRBID: &ie.DRBID{
									Value: 123,
								},
								DRBStatusUL: &ie.DRBStatusUL{
									Choice: &ie.DRBStatusUL12{
										ULCOUNTValue: &ie.COUNTValueForPDCPSN12{
											PDCPSN12:     intPtr(898),
											HFNPDCPSN12:  intPtr(345),
											IEExtensions: nil,
										},
										ReceiveStatusOfULPDCPSDUs: nil,
										IEExtension:               nil,
									},
								},
								DRBStatusDL: &ie.DRBStatusDL{
									Choice: &ie.DRBStatusDL12{
										DLCOUNTValue: &ie.COUNTValueForPDCPSN12{
											PDCPSN12:     intPtr(907),
											HFNPDCPSN12:  intPtr(987),
											IEExtensions: nil,
										},
										IEExtension: nil,
									},
								},
								IEExtension: nil,
							},
						},
					},
					IEExtensions: nil,
				},
			},
			expected: []byte{
				0x00, 0x31, 0x40, 0x27, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xd3, 0xe3, 0xa9, 0x22,
				0x1e, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x54, 0x00, 0x10, 0x00, 0x40, 0x01, 0x7b, 0x00,
				0x03, 0x82, 0x40, 0x01, 0x59, 0x00, 0x03, 0x8b, 0x40, 0x03, 0xdb,
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

func TestUplinkRANStatusTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	intPtr := func(i int64) *int64 {
		return &i
	}

	testCases := []struct {
		name     string
		input    []byte
		expected *UplinkRANStatusTransfer
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover",
			input: []byte{
				0x00, 0x31, 0x40, 0x27, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xd3, 0xe3, 0xa9, 0x22,
				0x1e, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x54, 0x00, 0x10, 0x00, 0x40, 0x01, 0x7b, 0x00,
				0x03, 0x82, 0x40, 0x01, 0x59, 0x00, 0x03, 0x8b, 0x40, 0x03, 0xdb,
			},
			expected: &UplinkRANStatusTransfer{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 910057611806,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				RANStatusTransferTransparentContainer: &ie.RANStatusTransferTransparentContainer{
					DRBsSubjectToStatusTransferList: &ie.DRBsSubjectToStatusTransferList{
						List: []ie.DRBsSubjectToStatusTransferItem{
							{
								DRBID: &ie.DRBID{
									Value: 123,
								},
								DRBStatusUL: &ie.DRBStatusUL{
									Choice: &ie.DRBStatusUL12{
										ULCOUNTValue: &ie.COUNTValueForPDCPSN12{
											PDCPSN12:     intPtr(898),
											HFNPDCPSN12:  intPtr(345),
											IEExtensions: nil,
										},
										ReceiveStatusOfULPDCPSDUs: nil,
										IEExtension:               nil,
									},
								},
								DRBStatusDL: &ie.DRBStatusDL{
									Choice: &ie.DRBStatusDL12{
										DLCOUNTValue: &ie.COUNTValueForPDCPSN12{
											PDCPSN12:     intPtr(907),
											HFNPDCPSN12:  intPtr(987),
											IEExtensions: nil,
										},
										IEExtension: nil,
									},
								},
								IEExtension: nil,
							},
						},
					},
					IEExtensions: nil,
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
