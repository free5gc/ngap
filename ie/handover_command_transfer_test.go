package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestHandoverCommandTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *HandoverCommandTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover",
			input: &HandoverCommandTransfer{
				DLForwardingUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes: []uint8{
									0xac, 0x10, 0x1f, 0x33,
								},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{0x00, 0x00, 0x00, 0x02},
						},
						IEExtensions: nil,
					},
				},
				QosFlowToBeForwardedList: &QosFlowToBeForwardedList{
					List: []QosFlowToBeForwardedItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1, // Assuming QFI is 1
							},
							IEExtensions: nil,
						},
					},
				},
				DataForwardingResponseDRBList: nil,
				IEExtensions:                  nil,
			},
			expected: []byte{
				0x60, 0x0f, 0x80, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02,
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

func TestHandoverCommandTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverCommandTransfer
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover",
			input: []byte{
				0x60, 0x0f, 0x80, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02,
			},
			expected: &HandoverCommandTransfer{
				DLForwardingUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes: []uint8{
									0xac, 0x10, 0x1f, 0x33,
								},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{0x00, 0x00, 0x00, 0x02},
						},
						IEExtensions: nil,
					},
				},
				QosFlowToBeForwardedList: &QosFlowToBeForwardedList{
					List: []QosFlowToBeForwardedItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1, // Assuming QFI is 1
							},
							IEExtensions: nil,
						},
					},
				},
				DataForwardingResponseDRBList: nil,
				IEExtensions:                  nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(HandoverCommandTransfer)
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
