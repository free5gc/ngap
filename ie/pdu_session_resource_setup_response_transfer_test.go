package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceSetupResponseTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceSetupResponseTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionEstablishment",
			input: &PDUSessionResourceSetupResponseTransfer{
				DLQosFlowPerTNLInformation: &QosFlowPerTNLInformation{
					UPTransportLayerInformation: &UPTransportLayerInformation{
						Choice: &GTPTunnel{
							TransportLayerAddress: &TransportLayerAddress{
								Value: aper.BitString{
									Bytes:     []byte{0xac, 0x10, 0x1f, 0x32},
									BitLength: 32,
								},
							},
							GTPTEID: &GTPTEID{
								Value: aper.OctetString{0x00, 0x00, 0x00, 0x01},
							},
						},
					},
					AssociatedQosFlowList: &AssociatedQosFlowList{
						List: []AssociatedQosFlowItem{
							{
								QosFlowIdentifier: &QosFlowIdentifier{
									Value: 1,
								},
							},
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x03, 0xe0, 0xac, 0x10, 0x1f, 0x32, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
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

func TestPDUSessionResourceSetupResponseTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceSetupResponseTransfer
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionEstablishment",
			input: []byte{
				0x00, 0x03, 0xe0, 0xac, 0x10, 0x1f, 0x32, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01,
			},
			expected: &PDUSessionResourceSetupResponseTransfer{
				DLQosFlowPerTNLInformation: &QosFlowPerTNLInformation{
					UPTransportLayerInformation: &UPTransportLayerInformation{
						Choice: &GTPTunnel{
							TransportLayerAddress: &TransportLayerAddress{
								Value: aper.BitString{
									Bytes:     []byte{0xac, 0x10, 0x1f, 0x32},
									BitLength: 32,
								},
							},
							GTPTEID: &GTPTEID{
								Value: aper.OctetString{0x00, 0x00, 0x00, 0x01},
							},
						},
					},
					AssociatedQosFlowList: &AssociatedQosFlowList{
						List: []AssociatedQosFlowItem{
							{
								QosFlowIdentifier: &QosFlowIdentifier{
									Value: 1,
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
			ie := new(PDUSessionResourceSetupResponseTransfer)
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
