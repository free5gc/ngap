package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestPDUSessionResourceSetupResponseMarshalBinary(t *testing.T) {
	t.Parallel()

	// PDUSessionResourceSetupResponseTransfer
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x01},
			BitLength: 32,
		},
	}
	gtpteid := &ie.GTPTEID{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	pDUSessionResourceSetupResponseTransfer := ie.PDUSessionResourceSetupResponseTransfer{
		DLQosFlowPerTNLInformation: &ie.QosFlowPerTNLInformation{
			UPTransportLayerInformation: &ie.UPTransportLayerInformation{
				Choice: &ie.GTPTunnel{
					TransportLayerAddress: ipv4Addr,
					GTPTEID:               gtpteid,
				},
			},
			AssociatedQosFlowList: &ie.AssociatedQosFlowList{
				List: []ie.AssociatedQosFlowItem{
					{
						QosFlowIdentifier: &ie.QosFlowIdentifier{
							Value: 9,
						},
					},
				},
			},
		},
	}
	marshalledPduSessionResourceSetupResponseTransferBytes, err := ie.MarshalBinary(&pDUSessionResourceSetupResponseTransfer)
	require.NoError(t, err)
	marshalledPduSessionResourceSetupResponseTransferOS := aper.OctetString(
		marshalledPduSessionResourceSetupResponseTransferBytes)

	testCases := []struct {
		name     string
		input    *PDUSessionResourceSetupResponse
		expected []byte
	}{
		{
			name: "Case 1",
			input: &PDUSessionResourceSetupResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 721420288,
				},
				PDUSessionResourceSetupListSURes: &ie.PDUSessionResourceSetupListSURes{
					List: []ie.PDUSessionResourceSetupItemSURes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							PDUSessionResourceSetupResponseTransfer: &marshalledPduSessionResourceSetupResponseTransferOS,
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x1d, 0x00, 0x27, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x02, 0x00,
				0x01, 0x00, 0x55, 0x40, 0x05, 0xc0, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x4b,
				0x40, 0x11, 0x00, 0x00, 0x05, 0x0d, 0x00, 0x03, 0xe0, 0xac, 0x10, 0x03,
				0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x09,
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestPDUSessionEstablishment",
			input: &PDUSessionResourceSetupResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1015608786343,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceSetupListSURes: &ie.PDUSessionResourceSetupListSURes{
					List: []ie.PDUSessionResourceSetupItemSURes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceSetupResponseTransferMarshalBinary
							PDUSessionResourceSetupResponseTransfer: &aper.OctetString{
								0x00, 0x03, 0xe0, 0xac, 0x10, 0x1f, 0x32, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
							},
							IEExtensions: nil,
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x1d, 0x00, 0x28, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xec, 0x77, 0x00, 0x3d,
				0xa7, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x4b, 0x40, 0x11, 0x00, 0x00, 0x0a, 0x0d, 0x00,
				0x03, 0xe0, 0xac, 0x10, 0x1f, 0x32, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
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

func TestPDUSessionResourceSetupResponseUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// PDUSessionResourceSetupResponseTransfer
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x01},
			BitLength: 32,
		},
	}
	gtpteid := &ie.GTPTEID{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	pDUSessionResourceSetupResponseTransfer := ie.PDUSessionResourceSetupResponseTransfer{
		DLQosFlowPerTNLInformation: &ie.QosFlowPerTNLInformation{
			UPTransportLayerInformation: &ie.UPTransportLayerInformation{
				Choice: &ie.GTPTunnel{
					TransportLayerAddress: ipv4Addr,
					GTPTEID:               gtpteid,
				},
			},
			AssociatedQosFlowList: &ie.AssociatedQosFlowList{
				List: []ie.AssociatedQosFlowItem{
					{
						QosFlowIdentifier: &ie.QosFlowIdentifier{
							Value: 9,
						},
					},
				},
			},
		},
	}
	marshalledPduSessionResourceSetupResponseTransferBytes, err := ie.MarshalBinary(
		&pDUSessionResourceSetupResponseTransfer,
	)
	require.NoError(t, err)
	marshalledPduSessionResourceSetupResponseTransferOS := aper.OctetString(
		marshalledPduSessionResourceSetupResponseTransferBytes)

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceSetupResponse
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x1d, 0x00, 0x27, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x02, 0x00,
				0x01, 0x00, 0x55, 0x40, 0x05, 0xc0, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x4b,
				0x40, 0x11, 0x00, 0x00, 0x05, 0x0d, 0x00, 0x03, 0xe0, 0xac, 0x10, 0x03,
				0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x09,
			},
			expected: &PDUSessionResourceSetupResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 721420288,
				},
				PDUSessionResourceSetupListSURes: &ie.PDUSessionResourceSetupListSURes{
					List: []ie.PDUSessionResourceSetupItemSURes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							PDUSessionResourceSetupResponseTransfer: &marshalledPduSessionResourceSetupResponseTransferOS,
						},
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestPDUSessionEstablishment",
			input: []byte{
				0x20, 0x1d, 0x00, 0x28, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xec, 0x77, 0x00, 0x3d,
				0xa7, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x4b, 0x40, 0x11, 0x00, 0x00, 0x0a, 0x0d, 0x00,
				0x03, 0xe0, 0xac, 0x10, 0x1f, 0x32, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
			},
			expected: &PDUSessionResourceSetupResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1015608786343,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceSetupListSURes: &ie.PDUSessionResourceSetupListSURes{
					List: []ie.PDUSessionResourceSetupItemSURes{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPDUSessionResourceSetupResponseTransferUnmarshalBinary
							PDUSessionResourceSetupResponseTransfer: &aper.OctetString{
								0x00, 0x03, 0xe0, 0xac, 0x10, 0x1f, 0x32, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
							},
							IEExtensions: nil,
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
