package message

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
	"github.com/stretchr/testify/require"
)

func TestPathSwitchRequestFailureMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PathSwitchRequestFailure
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestXnIPv6PathSwitchFail",
			input: &PathSwitchRequestFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 693350441842,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceReleasedListPSFail: &ie.PDUSessionResourceReleasedListPSFail{
					List: []ie.PDUSessionResourceReleasedItemPSFail{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested in TestPathSwitchRequestUnsuccessfulTransferMarshalBinary
							PathSwitchRequestUnsuccessfulTransfer: &aper.OctetString{0x08},
							IEExtensions:                          nil,
						},
					},
				},
				CriticalityDiagnostics: nil,
			},
			expected: []byte{
				0x40, 0x19, 0x00, 0x1c, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xa1, 0x6e, 0xe8, 0x23,
				0x72, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x45, 0x40, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x08,
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

func TestPathSwitchRequestFailureUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PathSwitchRequestFailure
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestXnIPv6PathSwitchFail",
			input: []byte{
				0x40, 0x19, 0x00, 0x1c, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xa1, 0x6e, 0xe8, 0x23,
				0x72, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x45, 0x40, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x08,
			},
			expected: &PathSwitchRequestFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 693350441842,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceReleasedListPSFail: &ie.PDUSessionResourceReleasedListPSFail{
					List: []ie.PDUSessionResourceReleasedItemPSFail{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested in TestPathSwitchRequestUnsuccessfulTransferUnmarshalBinary
							PathSwitchRequestUnsuccessfulTransfer: &aper.OctetString{0x08},
							IEExtensions:                          nil,
						},
					},
				},
				CriticalityDiagnostics: nil,
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
