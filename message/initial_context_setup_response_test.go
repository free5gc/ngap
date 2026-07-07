package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestInitialContextSetupResponseMarshalBinary(t *testing.T) {
	t.Parallel()

	// PDUSessionResourceSetupResponseTransfer
	// 172.16.3.1
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x01},
			BitLength: 32,
		},
	}
	gtpteid := &ie.GTPTEID{
		Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01}),
	}
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
	marshalledPduSessionResourceSetupResponseTransfer, err := ie.MarshalBinary(
		&pDUSessionResourceSetupResponseTransfer)
	require.NoError(t, err)
	marshalledPduSessionResourceSetupResponseTransferOS := aper.OctetString(
		marshalledPduSessionResourceSetupResponseTransfer)

	testCases := []struct {
		name     string
		input    *InitialContextSetupResponse
		expected []byte
	}{
		{
			name: "Case 1",
			input: &InitialContextSetupResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3456106496,
				},
				PDUSessionResourceSetupListCxtRes: &ie.PDUSessionResourceSetupListCxtRes{
					List: []ie.PDUSessionResourceSetupItemCxtRes{
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
				0x20, 0x0e, 0x00, 0x27, 0x00, 0x00, 0x03, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x01, 0x00, 0x55, 0x40, 0x05, 0xc0, 0xce, 0x00, 0x00,
				0x00, 0x00, 0x48, 0x40, 0x11, 0x00, 0x00, 0x05, 0x0d, 0x00, 0x03, 0xe0,
				0xac, 0x10, 0x03, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x09,
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

func TestInitialContextSetupResponseUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// PDUSessionResourceSetupResponseTransfer
	// 172.16.3.1
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x01},
			BitLength: 32,
		},
	}
	gtpteid := &ie.GTPTEID{
		Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01}),
	}
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
	marshalledPduSessionResourceSetupResponseTransfer, err := ie.MarshalBinary(&pDUSessionResourceSetupResponseTransfer)
	require.NoError(t, err)
	marshalledPduSessionResourceSetupResponseTransferOS := aper.OctetString(
		marshalledPduSessionResourceSetupResponseTransfer)

	testCases := []struct {
		name     string
		input    []byte
		expected *InitialContextSetupResponse
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x0e, 0x00, 0x27, 0x00, 0x00, 0x03, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x01, 0x00, 0x55, 0x40, 0x05, 0xc0, 0xce, 0x00, 0x00,
				0x00, 0x00, 0x48, 0x40, 0x11, 0x00, 0x00, 0x05, 0x0d, 0x00, 0x03, 0xe0,
				0xac, 0x10, 0x03, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x09,
			},
			expected: &InitialContextSetupResponse{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3456106496,
				},
				PDUSessionResourceSetupListCxtRes: &ie.PDUSessionResourceSetupListCxtRes{
					List: []ie.PDUSessionResourceSetupItemCxtRes{
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
