package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestPathSwitchRequestAcknowledgeTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PathSwitchRequestAcknowledgeTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestXnHandover",
			input: &PathSwitchRequestAcknowledgeTransfer{
				ULNGUUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []byte{0xac, 0x10, 0x1f, 0x3d},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{0x92, 0xf7, 0x52, 0x1f},
						},
						IEExtensions: nil,
					},
				},
			},
			expected: []byte{0x40, 0x1f, 0xac, 0x10, 0x1f, 0x3d, 0x92, 0xf7, 0x52, 0x1f},
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

func TestPathSwitchRequestAcknowledgeTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PathSwitchRequestAcknowledgeTransfer
	}{
		{
			name:  "Case 1: from ueranemu k8s-basic pipeline TestXnHandover",
			input: []byte{0x40, 0x1f, 0xac, 0x10, 0x1f, 0x3d, 0x92, 0xf7, 0x52, 0x1f},
			expected: &PathSwitchRequestAcknowledgeTransfer{
				ULNGUUPTNLInformation: &UPTransportLayerInformation{
					Choice: &GTPTunnel{
						TransportLayerAddress: &TransportLayerAddress{
							Value: aper.BitString{
								Bytes:     []byte{0xac, 0x10, 0x1f, 0x3d},
								BitLength: 32,
							},
						},
						GTPTEID: &GTPTEID{
							Value: aper.OctetString{0x92, 0xf7, 0x52, 0x1f},
						},
						IEExtensions: nil,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(PathSwitchRequestAcknowledgeTransfer)
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
