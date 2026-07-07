package ie

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
)

func TestUserLocationInformationMarshalBinary(t *testing.T) {
	t.Parallel()

	getByteArray := func(hexStr string) []byte {
		bs, err := hex.DecodeString(hexStr)
		require.NoError(t, err)
		return bs
	}

	tcs := []struct {
		name     string
		input    *UserLocationInformation
		expected []byte
	}{
		{
			name: "UserLocationInformationTwif",
			input: &UserLocationInformation{
				Choice: &ProtocolIESingleContainerUserLocationInformationExtIEs{
					UserLocationInformationExtIEs{
						UserLocationInformationTWIF: &UserLocationInformationTWIF{
							TWAPID: &TWAPID{
								Value: []byte{0x90, 0x3c, 0xb3, 0xb1, 0x6d, 0xa0},
							},
							IPAddress: &TransportLayerAddress{
								Value: aper.BitString{
									Bytes:     []byte{0xc0, 0xa8, 0x02, 0x01},
									BitLength: 32,
								},
							},
						},
					},
				},
			},
			expected: getByteArray("c000f8400e0006903cb3b16da00f80c0a80201"),
		},
	}

	for _, tc := range tcs {
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

func TestUserLocationInformationUnmarshalBinary(t *testing.T) {
	t.Parallel()

	getByteArray := func(hexStr string) []byte {
		bs, err := hex.DecodeString(hexStr)
		require.NoError(t, err)
		return bs
	}

	tcs := []struct {
		name     string
		input    []byte
		expected *UserLocationInformation
	}{
		{
			name:  "UserLocationInformationTwif",
			input: getByteArray("c000f8400e0006903cb3b16da00f80c0a80201"),
			expected: &UserLocationInformation{
				Choice: &ProtocolIESingleContainerUserLocationInformationExtIEs{
					UserLocationInformationExtIEs{
						UserLocationInformationTWIF: &UserLocationInformationTWIF{
							TWAPID: &TWAPID{
								Value: []byte{0x90, 0x3c, 0xb3, 0xb1, 0x6d, 0xa0},
							},
							IPAddress: &TransportLayerAddress{
								Value: aper.BitString{
									Bytes:     []byte{0xc0, 0xa8, 0x02, 0x01},
									BitLength: 32,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(UserLocationInformation)
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
