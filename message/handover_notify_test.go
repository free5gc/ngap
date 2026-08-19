package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestHandoverNotifyMarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	nrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x0d, 0x90, 0x30},
			BitLength: 36,
		},
	}

	testCases := []struct {
		name     string
		input    *HandoverNotify
		expected []byte
	}{
		{
			name: "Case 1",
			input: &HandoverNotify{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3439329281,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case1,
							NRCellIdentity: nrCellIdentity,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x0b, 0x40, 0x25, 0x00, 0x00, 0x03, 0x00, 0x0a,
				0x00, 0x02, 0x00, 0x02, 0x00, 0x55, 0x00, 0x05, 0xc0,
				0xcd, 0x00, 0x00, 0x01, 0x00, 0x79, 0x40, 0x0f, 0x40,
				0x13, 0x30, 0x01, 0x00, 0x00, 0x0d, 0x90, 0x30, 0x13,
				0x30, 0x01, 0x00, 0x00, 0x01,
			},
		},
		{
			name: "Case 2: from ueranemu basic-k8s pipeline TestN2Handover",
			input: &HandoverNotify{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 887252622092,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x01, 0x02, 0x00, 0x10},
									BitLength: 36,
								},
							},
							IEExtensions: nil,
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x01},
							},
							IEExtensions: nil,
						},
						TimeStamp:    nil,
						IEExtensions: nil,
					},
				},
				NotifySourceNGRANNode: nil,
			},
			expected: []byte{
				0x00, 0x0b, 0x40, 0x26, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xce, 0x94, 0x60, 0x9b,
				0x0c, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00,
				0x01, 0x02, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01,
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

func TestHandoverNotifyUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	nrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x0d, 0x90, 0x30},
			BitLength: 36,
		},
	}

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverNotify
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x0b, 0x40, 0x25, 0x00, 0x00, 0x03, 0x00, 0x0a,
				0x00, 0x02, 0x00, 0x02, 0x00, 0x55, 0x00, 0x05, 0xc0,
				0xcd, 0x00, 0x00, 0x01, 0x00, 0x79, 0x40, 0x0f, 0x40,
				0x13, 0x30, 0x01, 0x00, 0x00, 0x0d, 0x90, 0x30, 0x13,
				0x30, 0x01, 0x00, 0x00, 0x01,
			},
			expected: &HandoverNotify{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3439329281,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case1,
							NRCellIdentity: nrCellIdentity,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu basic-k8s pipeline TestN2Handover",
			input: []byte{
				0x00, 0x0b, 0x40, 0x26, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xce, 0x94, 0x60, 0x9b,
				0x0c, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00,
				0x01, 0x02, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01,
			},
			expected: &HandoverNotify{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 887252622092,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x01, 0x02, 0x00, 0x10},
									BitLength: 36,
								},
							},
							IEExtensions: nil,
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x01},
							},
							IEExtensions: nil,
						},
						TimeStamp:    nil,
						IEExtensions: nil,
					},
				},
				NotifySourceNGRANNode: nil,
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
