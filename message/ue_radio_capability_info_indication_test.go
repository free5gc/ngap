package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestUERadioCapabilityInfoIndMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *UERadioCapabilityInfoIndication
		expected []byte
	}{
		{
			name: "Case 1",
			input: &UERadioCapabilityInfoIndication{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3456106496,
				},
				UERadioCapability: &ie.UERadioCapability{
					Value: aper.OctetString([]byte{0x10, 0x14, 0x00, 0x00, 0x1f, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xff, 0xc0}),
				},
			},
			expected: []byte{
				0x00, 0x2c, 0x40, 0x23, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x02, 0x00,
				0x01, 0x00, 0x55, 0x00, 0x05, 0xc0, 0xce, 0x00, 0x00, 0x00, 0x00, 0x75,
				0x40, 0x0d, 0x0c, 0x10, 0x14, 0x00, 0x00, 0x1f, 0xf8, 0x00, 0x00, 0x00,
				0x00, 0xff, 0xc0,
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

func TestUERadioCapabilityInfoIndUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *UERadioCapabilityInfoIndication
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x2c, 0x40, 0x23, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x02, 0x00,
				0x01, 0x00, 0x55, 0x00, 0x05, 0xc0, 0xce, 0x00, 0x00, 0x00, 0x00, 0x75,
				0x40, 0x0d, 0x0c, 0x10, 0x14, 0x00, 0x00, 0x1f, 0xf8, 0x00, 0x00, 0x00,
				0x00, 0xff, 0xc0,
			},
			expected: &UERadioCapabilityInfoIndication{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3456106496,
				},
				UERadioCapability: &ie.UERadioCapability{
					Value: aper.OctetString([]byte{0x10, 0x14, 0x00, 0x00, 0x1f, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xff, 0xc0}),
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
