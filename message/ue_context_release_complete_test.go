package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestUEContextReleaseCompleteMarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	nrCellIdentity_case1 := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x02, 0xb6, 0x70},
			BitLength: 36,
		},
	}

	// Case 2
	plmnIdentity_case2 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x00, 0x00, 0x00})}
	tac_case2 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x00})}
	nrCellIdentity_case2 := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
			BitLength: 36,
		},
	}
	nid_case2 := &ie.NID{
		Value: aper.BitString{
			Bytes:     []byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x90},
			BitLength: 44,
		},
	}

	testCases := []struct {
		name     string
		input    *UEContextReleaseComplete
		expected []byte
	}{
		{
			name: "Case 1",
			input: &UEContextReleaseComplete{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3405774848,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case1,
							NRCellIdentity: nrCellIdentity_case1,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
				PDUSessionResourceListCxtRelCpl: &ie.PDUSessionResourceListCxtRelCpl{
					List: []ie.PDUSessionResourceItemCxtRelCpl{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x29, 0x00, 0x2c, 0x00, 0x00, 0x04, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x01, 0x00, 0x55, 0x40, 0x05, 0xc0,
				0xcb, 0x00, 0x00, 0x00, 0x00, 0x79, 0x40, 0x0f, 0x40,
				0x13, 0x30, 0x01, 0x00, 0x00, 0x02, 0xb6, 0x70, 0x13,
				0x30, 0x01, 0x00, 0x00, 0x01, 0x00, 0x3c, 0x00, 0x03,
				0x00, 0x00, 0x05,
			},
		},
		{
			name: "Case 2 (R16)",
			input: &UEContextReleaseComplete{
				// Contains R16 Field
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1157627904,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 234881024,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case2,
							NRCellIdentity: nrCellIdentity_case2,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case2,
							TAC:          tac_case2,
						},
						IEExtensions: &ie.ProtocolExtensionContainerUserLocationInformationNRExtIEs{
							List: []ie.UserLocationInformationNRExtIEs{
								{
									Id: &ie.ProtocolExtensionID{
										Value: ie.ProtocolIEIDNID,
									},
									Criticality: &ie.Criticality{
										Value: ie.CriticalityReject,
									},
									NID: nid_case2,
								},
							},
						},
					},
				},
				PDUSessionResourceListCxtRelCpl: &ie.PDUSessionResourceListCxtRelCpl{
					List: []ie.PDUSessionResourceItemCxtRelCpl{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
						},
					},
				},
			},
			expected: []byte{
				// Contains R16 Field
				0x20, 0x29, 0x00, 0x3b, 0x00, 0x00, 0x04, 0x00, 0x0a,
				0x40, 0x05, 0x60, 0x45, 0x00, 0x00, 0x00, 0x00, 0x55,
				0x40, 0x05, 0xc0, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x79,
				0x40, 0x1b, 0x48, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x01, 0x07, 0x00, 0x06, 0x99, 0x99, 0x99, 0x99,
				0x99, 0x90, 0x00, 0x3c, 0x00, 0x03, 0x00, 0x00, 0x05,
			},
		},
		{
			name: "Case 3: from ueranemu k8s-basic pipeline TestServiceRequest",
			input: &UEContextReleaseComplete{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 885695317650,
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
									Bytes:     []uint8{0x00, 0x00, 0x04, 0x00, 0x10},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x01},
							},
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x29, 0x00, 0x26, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xce, 0x37, 0x8e, 0x06,
				0x92, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00,
				0x00, 0x04, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01,
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

func TestUEContextReleaseCompleteUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	nrCellIdentity_case1 := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x02, 0xb6, 0x70},
			BitLength: 36,
		},
	}

	// Case 2
	plmnIdentity_case2 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x00, 0x00, 0x00})}
	tac_case2 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x00})}
	nrCellIdentity_case2 := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
			BitLength: 36,
		},
	}
	nid_case2 := &ie.NID{
		Value: aper.BitString{
			Bytes:     []byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x90},
			BitLength: 44,
		},
	}

	testCases := []struct {
		name     string
		input    []byte
		expected *UEContextReleaseComplete
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x29, 0x00, 0x2c, 0x00, 0x00, 0x04, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x01, 0x00, 0x55, 0x40, 0x05, 0xc0,
				0xcb, 0x00, 0x00, 0x00, 0x00, 0x79, 0x40, 0x0f, 0x40,
				0x13, 0x30, 0x01, 0x00, 0x00, 0x02, 0xb6, 0x70, 0x13,
				0x30, 0x01, 0x00, 0x00, 0x01, 0x00, 0x3c, 0x00, 0x03,
				0x00, 0x00, 0x05,
			},
			expected: &UEContextReleaseComplete{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3405774848,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case1,
							NRCellIdentity: nrCellIdentity_case1,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
				PDUSessionResourceListCxtRelCpl: &ie.PDUSessionResourceListCxtRelCpl{
					List: []ie.PDUSessionResourceItemCxtRelCpl{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "Case 2 (R16)",
			input: []byte{
				// Contains R16 Field
				0x20, 0x29, 0x00, 0x3b, 0x00, 0x00, 0x04, 0x00, 0x0a,
				0x40, 0x05, 0x60, 0x45, 0x00, 0x00, 0x00, 0x00, 0x55,
				0x40, 0x05, 0xc0, 0x0e, 0x00, 0x00, 0x00, 0x00, 0x79,
				0x40, 0x1b, 0x48, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x01, 0x07, 0x00, 0x06, 0x99, 0x99, 0x99, 0x99,
				0x99, 0x90, 0x00, 0x3c, 0x00, 0x03, 0x00, 0x00, 0x05,
			},
			expected: &UEContextReleaseComplete{
				// Contains R16 Field
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1157627904,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 234881024,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case2,
							NRCellIdentity: nrCellIdentity_case2,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case2,
							TAC:          tac_case2,
						},
						IEExtensions: &ie.ProtocolExtensionContainerUserLocationInformationNRExtIEs{
							List: []ie.UserLocationInformationNRExtIEs{
								{
									Id: &ie.ProtocolExtensionID{
										Value: ie.ProtocolIEIDNID,
									},
									Criticality: &ie.Criticality{
										Value: ie.CriticalityReject,
									},
									NID: nid_case2,
								},
							},
						},
					},
				},
				PDUSessionResourceListCxtRelCpl: &ie.PDUSessionResourceListCxtRelCpl{
					List: []ie.PDUSessionResourceItemCxtRelCpl{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "Case 3: from ueranemu k8s-basic pipeline TestServiceRequest",
			input: []byte{
				0x20, 0x29, 0x00, 0x26, 0x00, 0x00, 0x03, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xce, 0x37, 0x8e, 0x06,
				0x92, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00,
				0x00, 0x04, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01,
			},
			expected: &UEContextReleaseComplete{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 885695317650,
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
									Bytes:     []uint8{0x00, 0x00, 0x04, 0x00, 0x10},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x01},
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
