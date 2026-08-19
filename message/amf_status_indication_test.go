package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestAMFStatusIndicationMarshalBinary(t *testing.T) {
	t.Parallel()

	plmnIdentity := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x00, 0xF1, 0x53})}
	// AMFID: "cafe00"
	amfRegionId := &ie.AMFRegionID{Value: aper.BitString{Bytes: []byte{0xCA}, BitLength: 8}}
	amfSetId := &ie.AMFSetID{Value: aper.BitString{Bytes: []byte{0xFE, 0x00}, BitLength: 10}}
	amfPointer := &ie.AMFPointer{Value: aper.BitString{Bytes: []byte{0x00}, BitLength: 6}}

	testCases := []struct {
		name     string
		input    *AMFStatusIndication
		expected []byte
	}{
		{
			name: "Case 1",
			input: &AMFStatusIndication{
				UnavailableGUAMIList: &ie.UnavailableGUAMIList{
					List: []ie.UnavailableGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: plmnIdentity,
								AMFRegionID:  amfRegionId,
								AMFSetID:     amfSetId,
								AMFPointer:   amfPointer,
							},
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x01, 0x40, 0x0f, 0x00, 0x00, 0x01, 0x00, 0x78,
				0x00, 0x08, 0x00, 0x00, 0x00, 0xf1, 0x53, 0xca, 0xfe,
				0x00,
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

func TestAMFStatusIndicationUnmarshalBinary(t *testing.T) {
	t.Parallel()

	plmnIdentity := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x00, 0xF1, 0x53})}
	// AMFID: "cafe00"
	amfRegionId := &ie.AMFRegionID{Value: aper.BitString{Bytes: []byte{0xCA}, BitLength: 8}}
	amfSetId := &ie.AMFSetID{Value: aper.BitString{Bytes: []byte{0xFE, 0x00}, BitLength: 10}}
	amfPointer := &ie.AMFPointer{Value: aper.BitString{Bytes: []byte{0x00}, BitLength: 6}}

	testCases := []struct {
		name     string
		input    []byte
		expected *AMFStatusIndication
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x01, 0x40, 0x0f, 0x00, 0x00, 0x01, 0x00, 0x78,
				0x00, 0x08, 0x00, 0x00, 0x00, 0xf1, 0x53, 0xca, 0xfe,
				0x00,
			},
			expected: &AMFStatusIndication{
				UnavailableGUAMIList: &ie.UnavailableGUAMIList{
					List: []ie.UnavailableGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: plmnIdentity,
								AMFRegionID:  amfRegionId,
								AMFSetID:     amfSetId,
								AMFPointer:   amfPointer,
							},
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
