package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestPathSwitchRequestTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PathSwitchRequestTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic TestXnHandover",
			input: &PathSwitchRequestTransfer{
				DLNGUUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []byte{0xac, 0x10, 0x1f, 0x33},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{0x00, 0x00, 0x00, 0x02},
						},
						IEExtensions: nil,
					},
				},
				DLNGUTNLInformationReused:    nil,
				UserPlaneSecurityInformation: nil,
				QosFlowAcceptedList: &QosFlowAcceptedList{
					List: []QosFlowAcceptedItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
							IEExtensions: nil,
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x1f, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02,
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

func TestPathSwitchRequestTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PathSwitchRequestTransfer
	}{
		{
			name: "Case 1: from ueranemu k8s-basic TestXnHandover",
			input: []byte{
				0x00, 0x1f, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02,
			},
			expected: &PathSwitchRequestTransfer{
				DLNGUUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []byte{0xac, 0x10, 0x1f, 0x33},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{0x00, 0x00, 0x00, 0x02},
						},
						IEExtensions: nil,
					},
				},
				DLNGUTNLInformationReused:    nil,
				UserPlaneSecurityInformation: nil,
				QosFlowAcceptedList: &QosFlowAcceptedList{
					List: []QosFlowAcceptedItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
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
			ie := new(PathSwitchRequestTransfer)
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
