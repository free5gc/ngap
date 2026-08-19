package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestInitialUEMessageMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *InitialUEMessage
		expected []byte
	}{
		{
			name: "Case 1",
			input: &InitialUEMessage{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3456106496,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d, 0x01, 0x13, 0x00, 0x13, 0x0f, 0xff,
						0x00, 0x00, 0x41, 0x00, 0x00, 0x21, 0xf0, 0x2e, 0x04, 0x80, 0x20, 0xe0, 0xe0, 0x17, 0x07, 0xe0, 0xe0,
						0xc0, 0x40, 0x00, 0x80, 0x20,
					}),
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x13, 0x30, 0x01},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x13, 0x30, 0x01},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
						},
					},
				},
				RRCEstablishmentCause: &ie.RRCEstablishmentCause{
					Value: ie.RRCEstablishmentCausePresentMoSignalling,
				},
				UEContextRequest: &ie.UEContextRequest{
					Value: ie.UEContextRequestPresentRequested,
				},
			},
			expected: []byte{
				0x00, 0x0f, 0x40, 0x50, 0x00, 0x00, 0x05, 0x00, 0x55,
				0x00, 0x05, 0xc0, 0xce, 0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x23, 0x22,
				0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d, 0x01, 0x13, 0x00, 0x13, 0x0f, 0xff,
				0x00, 0x00, 0x41, 0x00, 0x00, 0x21, 0xf0, 0x2e, 0x04, 0x80, 0x20, 0xe0,
				0xe0, 0x17, 0x07, 0xe0, 0xe0, 0xc0, 0x40, 0x00, 0x80, 0x20, 0x00, 0x79,
				0x00, 0x0f, 0x40, 0x13, 0x30, 0x01, 0x00, 0x00, 0x00, 0x00, 0x10, 0x13,
				0x30, 0x01, 0x00, 0x00, 0x01, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x70,
				0x40, 0x01, 0x00,
			},
		},
		{
			name: "Case 2 (R16)",
			input: &InitialUEMessage{
				// Contains R16 Fields
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 234881024,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d, 0x01,
						0x13, 0x00, 0x13, 0xf0, 0xff, 0x00, 0x00, 0x41, 0x00, 0x00, 0x21, 0xf0,
						0x2e, 0x02, 0xf0, 0xf0,
					}),
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x00, 0x00, 0x00},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x00, 0x00, 0x00},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x00}),
							},
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
									NID: &ie.NID{
										Value: aper.BitString{
											Bytes:     []byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x90},
											BitLength: 44,
										},
									},
								},
							},
						},
					},
				},
				RRCEstablishmentCause: &ie.RRCEstablishmentCause{
					Value: ie.RRCEstablishmentCausePresentMoSignalling,
				},
				UEContextRequest: &ie.UEContextRequest{
					Value: ie.UEContextRequestPresentRequested,
				},
			},
			expected: []byte{
				// Contains R16 Fields
				0x00, 0x0f, 0x40, 0x51, 0x00, 0x00, 0x05, 0x00, 0x55, 0x00, 0x05, 0xc0,
				0x0e, 0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x18, 0x17, 0x7e, 0x00, 0x41,
				0x79, 0x00, 0x0d, 0x01, 0x13, 0x00, 0x13, 0xf0, 0xff, 0x00, 0x00, 0x41,
				0x00, 0x00, 0x21, 0xf0, 0x2e, 0x02, 0xf0, 0xf0, 0x00, 0x79, 0x00, 0x1b,
				0x48, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x07, 0x00, 0x06, 0x99, 0x99, 0x99,
				0x99, 0x99, 0x90, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x70, 0x40, 0x01,
				0x00,
			},
		},
		{
			name: "Case 3: SNPN mode (from ueranemu SNPN pipeline TestRegistration)",
			input: &InitialUEMessage{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x00, 0x41, 0x79, 0x00, 0x0c, 0x01, 0x02,
						0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00,
						0x47, 0x78, 0x2e, 0x02, 0xe0, 0xe0,
					}),
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x04, 0x00, 0x10},
									BitLength: 36,
								},
							},
							IEExtensions: nil,
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
							},
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
							IEExtensions: nil,
						},
						TimeStamp: nil,
						IEExtensions: &ie.ProtocolExtensionContainerUserLocationInformationNRExtIEs{
							List: []ie.UserLocationInformationNRExtIEs{
								{
									Id: &ie.ProtocolExtensionID{
										Value: 263,
									},
									Criticality: &ie.Criticality{
										Value: aper.Enumerated(0),
									},
									PSCellInformation: nil,
									NID: &ie.NID{
										Value: aper.BitString{
											Bytes:     []byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x10},
											BitLength: 44,
										},
									},
									NRNTNTAIInformation: nil,
								},
							},
						},
					},
				},
				RRCEstablishmentCause: &ie.RRCEstablishmentCause{
					Value: aper.Enumerated(3),
				},
				FiveGSTMSI:                          nil,
				AMFSetID:                            nil,
				UEContextRequest:                    &ie.UEContextRequest{Value: aper.Enumerated(0)},
				AllowedNSSAI:                        nil,
				SourceToTargetAMFInformationReroute: nil,
				SelectedPLMNIdentity:                nil,
				IABNodeIndication:                   nil,
				CEmodeBSupportIndicator:             nil,
				LTEMIndication:                      nil,
				EDTSession:                          nil,
				AuthenticatedIndication:             nil,
				NPNAccessInformation:                nil,
				RedCapIndication:                    nil,
			},
			expected: []byte{
				0x00, 0x0f, 0x40, 0x4d, 0x00, 0x00, 0x05, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x26, 0x00, 0x17, 0x16, 0x7e, 0x00,
				0x41, 0x79, 0x00, 0x0c, 0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x47, 0x78, 0x2e, 0x02, 0xe0, 0xe0,
				0x00, 0x79, 0x00, 0x1b, 0x48, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x04, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01, 0x00,
				0x00, 0x01, 0x07, 0x00, 0x06, 0x10, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x70, 0x40, 0x01,
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

func TestInitialUEMessageUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *InitialUEMessage
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x0f, 0x40, 0x50, 0x00, 0x00, 0x05, 0x00, 0x55,
				0x00, 0x05, 0xc0, 0xce, 0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x23, 0x22,
				0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d, 0x01, 0x13, 0x00, 0x13, 0x0f, 0xff,
				0x00, 0x00, 0x41, 0x00, 0x00, 0x21, 0xf0, 0x2e, 0x04, 0x80, 0x20, 0xe0,
				0xe0, 0x17, 0x07, 0xe0, 0xe0, 0xc0, 0x40, 0x00, 0x80, 0x20, 0x00, 0x79,
				0x00, 0x0f, 0x40, 0x13, 0x30, 0x01, 0x00, 0x00, 0x00, 0x00, 0x10, 0x13,
				0x30, 0x01, 0x00, 0x00, 0x01, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x70,
				0x40, 0x01, 0x00,
			},
			expected: &InitialUEMessage{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3456106496,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d, 0x01, 0x13, 0x00, 0x13, 0x0f, 0xff,
						0x00, 0x00, 0x41, 0x00, 0x00, 0x21, 0xf0, 0x2e, 0x04, 0x80, 0x20, 0xe0, 0xe0, 0x17, 0x07, 0xe0, 0xe0,
						0xc0, 0x40, 0x00, 0x80, 0x20,
					}),
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x13, 0x30, 0x01},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x13, 0x30, 0x01},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
						},
					},
				},
				RRCEstablishmentCause: &ie.RRCEstablishmentCause{
					Value: ie.RRCEstablishmentCausePresentMoSignalling,
				},
				UEContextRequest: &ie.UEContextRequest{
					Value: ie.UEContextRequestPresentRequested,
				},
			},
		},
		{
			name: "Case 2 (R16)",
			input: []byte{
				// Contains R16 Fields
				0x00, 0x0f, 0x40, 0x51, 0x00, 0x00, 0x05, 0x00, 0x55, 0x00, 0x05, 0xc0,
				0x0e, 0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x18, 0x17, 0x7e, 0x00, 0x41,
				0x79, 0x00, 0x0d, 0x01, 0x13, 0x00, 0x13, 0xf0, 0xff, 0x00, 0x00, 0x41,
				0x00, 0x00, 0x21, 0xf0, 0x2e, 0x02, 0xf0, 0xf0, 0x00, 0x79, 0x00, 0x1b,
				0x48, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x07, 0x00, 0x06, 0x99, 0x99, 0x99,
				0x99, 0x99, 0x90, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x70, 0x40, 0x01,
				0x00,
			},
			expected: &InitialUEMessage{
				// Contains R16 Fields
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 234881024,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x00, 0x41, 0x79, 0x00, 0x0d, 0x01,
						0x13, 0x00, 0x13, 0xf0, 0xff, 0x00, 0x00, 0x41, 0x00, 0x00, 0x21, 0xf0,
						0x2e, 0x02, 0xf0, 0xf0,
					}),
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x00, 0x00, 0x00},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
									BitLength: 36,
								},
							},
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x00, 0x00, 0x00},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x00}),
							},
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
									NID: &ie.NID{
										Value: aper.BitString{
											Bytes:     []byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x90},
											BitLength: 44,
										},
									},
								},
							},
						},
					},
				},
				RRCEstablishmentCause: &ie.RRCEstablishmentCause{
					Value: ie.RRCEstablishmentCausePresentMoSignalling,
				},
				UEContextRequest: &ie.UEContextRequest{
					Value: ie.UEContextRequestPresentRequested,
				},
			},
		},
		{
			name: "Case 3: SNPN mode (from ueranemu SNPN pipeline TestRegistration)",
			input: []byte{
				0x00, 0x0f, 0x40, 0x4d, 0x00, 0x00, 0x05, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x26, 0x00, 0x17, 0x16, 0x7e, 0x00,
				0x41, 0x79, 0x00, 0x0c, 0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x47, 0x78, 0x2e, 0x02, 0xe0, 0xe0,
				0x00, 0x79, 0x00, 0x1b, 0x48, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x04, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01, 0x00,
				0x00, 0x01, 0x07, 0x00, 0x06, 0x10, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x70, 0x40, 0x01,
				0x00,
			},
			expected: &InitialUEMessage{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				NASPDU: &ie.NASPDU{
					Value: aper.OctetString([]byte{
						0x7e, 0x00, 0x41, 0x79, 0x00, 0x0c, 0x01, 0x02,
						0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00,
						0x47, 0x78, 0x2e, 0x02, 0xe0, 0xe0,
					}),
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x04, 0x00, 0x10},
									BitLength: 36,
								},
							},
							IEExtensions: nil,
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
							},
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
							IEExtensions: nil,
						},
						TimeStamp: nil,
						IEExtensions: &ie.ProtocolExtensionContainerUserLocationInformationNRExtIEs{
							List: []ie.UserLocationInformationNRExtIEs{
								{
									Id: &ie.ProtocolExtensionID{
										Value: 263,
									},
									Criticality: &ie.Criticality{
										Value: aper.Enumerated(0),
									},
									PSCellInformation: nil,
									NID: &ie.NID{
										Value: aper.BitString{
											Bytes:     []byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x10},
											BitLength: 44,
										},
									},
									NRNTNTAIInformation: nil,
								},
							},
						},
					},
				},
				RRCEstablishmentCause: &ie.RRCEstablishmentCause{
					Value: aper.Enumerated(3),
				},
				FiveGSTMSI:                          nil,
				AMFSetID:                            nil,
				UEContextRequest:                    &ie.UEContextRequest{Value: aper.Enumerated(0)},
				AllowedNSSAI:                        nil,
				SourceToTargetAMFInformationReroute: nil,
				SelectedPLMNIdentity:                nil,
				IABNodeIndication:                   nil,
				CEmodeBSupportIndicator:             nil,
				LTEMIndication:                      nil,
				EDTSession:                          nil,
				AuthenticatedIndication:             nil,
				NPNAccessInformation:                nil,
				RedCapIndication:                    nil,
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
