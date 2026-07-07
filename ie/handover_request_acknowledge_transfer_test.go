package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestHandoverRequestAcknowledgeTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *HandoverRequestAcknowledgeTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu basic-k8s TestN2Handover",
			input: &HandoverRequestAcknowledgeTransfer{
				DLNGUUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []uint8{0xac, 0x10, 0x1f, 0x33},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{
								0x00, 0x00, 0x00, 0x01,
							},
						},
						IEExtensions: nil,
					},
				},
				DLForwardingUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []uint8{0xac, 0x10, 0x1f, 0x33},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{
								0x00, 0x00, 0x00, 0x02,
							},
						},
						IEExtensions: nil,
					},
				},
				SecurityResult: nil,
				QosFlowSetupResponseList: &QosFlowListWithDataForwarding{
					List: []QosFlowItemWithDataForwarding{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
							DataForwardingAccepted: nil,
							IEExtensions:           nil,
						},
					},
				},
				QosFlowFailedToSetupList: nil,
			},
			expected: []byte{
				0x40, 0x07, 0xc0, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x01, 0x01, 0xf0, 0xac, 0x10, 0x1f,
				0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := MarshalBinary(tc.input)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, b)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestHandoverRequestAcknowledgeTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverRequestAcknowledgeTransfer
	}{
		{
			name: "Case 1: from ueranemu basic-k8s TestN2Handover",
			input: []byte{
				0x40, 0x07, 0xc0, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x01, 0x01, 0xf0, 0xac, 0x10, 0x1f,
				0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x01,
			},
			expected: &HandoverRequestAcknowledgeTransfer{
				DLNGUUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []uint8{0xac, 0x10, 0x1f, 0x33},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{
								0x00, 0x00, 0x00, 0x01,
							},
						},
						IEExtensions: nil,
					},
				},
				DLForwardingUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []uint8{0xac, 0x10, 0x1f, 0x33},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{
								0x00, 0x00, 0x00, 0x02,
							},
						},
						IEExtensions: nil,
					},
				},
				SecurityResult: nil,
				QosFlowSetupResponseList: &QosFlowListWithDataForwarding{
					List: []QosFlowItemWithDataForwarding{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
							DataForwardingAccepted: nil,
							IEExtensions:           nil,
						},
					},
				},
				QosFlowFailedToSetupList: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := &HandoverRequestAcknowledgeTransfer{}
			err := UnmarshalBinary(tc.input, ie)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, ie)
			} else {
				require.Error(t, err)
			}
		})
	}
}
