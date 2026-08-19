package message

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestNGSetupResponseMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *NGSetupResponse
		expected []byte
	}{
		{
			name: "Case 1",
			input: &NGSetupResponse{
				AMFName: &ie.AMFName{
					Value: aper.PrintableString("AMF"),
				},
				ServedGUAMIList: &ie.ServedGUAMIList{
					List: []ie.ServedGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: []byte{0x13, 0x30, 0x01},
								},
								// AMFID: "cafe00",
								AMFRegionID: &ie.AMFRegionID{
									Value: aper.BitString{
										Bytes:     []byte{0xca},
										BitLength: 8,
									},
								},
								AMFSetID: &ie.AMFSetID{
									Value: aper.BitString{
										Bytes:     []byte{0xfe, 0x00},
										BitLength: 10,
									},
								},
								AMFPointer: &ie.AMFPointer{
									Value: aper.BitString{
										Bytes:     []byte{0x00},
										BitLength: 6,
									},
								},
							},
						},
					},
				},
				RelativeAMFCapacity: &ie.RelativeAMFCapacity{
					Value: 255,
				},
				PLMNSupportList: &ie.PLMNSupportList{
					List: []ie.PLMNSupportItem{
						{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x13, 0x30, 0x01},
							},
							SliceSupportList: &ie.SliceSupportList{
								List: []ie.SliceSupportItem{
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
											SD: &ie.SD{
												Value: aper.OctetString([]byte{0x01, 0x02, 0x03}),
											},
										},
									},
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
											SD: &ie.SD{
												Value: aper.OctetString([]byte{0x11, 0x22, 0x33}),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x15, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
				0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
				0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
				0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
			},
		},
		{
			name: "Case 2 (R16)",
			input: &NGSetupResponse{
				AMFName: &ie.AMFName{
					Value: aper.PrintableString("src0.AMF.Spirent.com"),
				},
				ServedGUAMIList: &ie.ServedGUAMIList{
					List: []ie.ServedGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: []byte{0x00, 0x00, 0x00},
								},
								// AMFID: "000000",
								AMFRegionID: &ie.AMFRegionID{
									Value: aper.BitString{
										Bytes:     []byte{0x00},
										BitLength: 8,
									},
								},
								AMFSetID: &ie.AMFSetID{
									Value: aper.BitString{
										Bytes:     []byte{0x00, 0x00},
										BitLength: 10,
									},
								},
								AMFPointer: &ie.AMFPointer{
									Value: aper.BitString{
										Bytes:     []byte{0x00},
										BitLength: 6,
									},
								},
							},
						},
					},
				},
				RelativeAMFCapacity: &ie.RelativeAMFCapacity{
					Value: 10,
				},
				PLMNSupportList: &ie.PLMNSupportList{
					List: []ie.PLMNSupportItem{
						{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x00, 0x00, 0x00},
							},
							SliceSupportList: &ie.SliceSupportList{
								List: []ie.SliceSupportItem{
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
										},
									},
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x02}),
											},
										},
									},
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x03}),
											},
										},
									},
								},
							},
							IEExtensions: &ie.ProtocolExtensionContainerPLMNSupportItemExtIEs{
								List: []ie.PLMNSupportItemExtIEs{
									{
										Id: &ie.ProtocolExtensionID{
											Value: ie.ProtocolIEIDNPNSupport,
										},
										Criticality: &ie.Criticality{
											Value: ie.CriticalityReject,
										},
										NPNSupport: &ie.NPNSupport{
											Choice: &ie.NID{
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
					},
				},
			},
			expected: []byte{ // Contains R16 Fields
				0x20, 0x15, 0x00, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x16, 0x09, 0x80, 0x73,
				0x72, 0x63, 0x30, 0x2e, 0x41, 0x4d, 0x46, 0x2e, 0x53, 0x70, 0x69, 0x72, 0x65, 0x6e,
				0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x56, 0x40, 0x01, 0x0a, 0x00, 0x50, 0x00, 0x18, 0x04, 0x00,
				0x00, 0x00, 0x00, 0x02, 0x00, 0x08, 0x00, 0x80, 0x06, 0x00, 0x00, 0x01, 0x02, 0x00,
				0x07, 0x00, 0x99, 0x99, 0x99, 0x99, 0x99, 0x90,
			},
		},
		{
			name: "Case 3: SNPN mode (from ueranemu SNPN pipeline TestRegistration)",
			input: &NGSetupResponse{
				AMFName: &ie.AMFName{
					Value: aper.PrintableString("amf1"),
				},
				ServedGUAMIList: &ie.ServedGUAMIList{
					List: []ie.ServedGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
								},
								AMFRegionID: &ie.AMFRegionID{
									Value: aper.BitString{
										Bytes:     []byte{0x01},
										BitLength: 8,
									},
								},
								AMFSetID: &ie.AMFSetID{
									Value: aper.BitString{
										Bytes:     []byte{0x00, 0x40},
										BitLength: 10,
									},
								},
								AMFPointer: &ie.AMFPointer{
									Value: aper.BitString{
										Bytes:     []byte{0x04},
										BitLength: 6,
									},
								},
								IEExtensions: nil,
							},
							BackupAMFName: nil,
							IEExtensions:  nil,
						},
					},
				},
				RelativeAMFCapacity: &ie.RelativeAMFCapacity{
					Value: 255,
				},
				PLMNSupportList: &ie.PLMNSupportList{
					List: []ie.PLMNSupportItem{
						{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
							},
							SliceSupportList: &ie.SliceSupportList{
								List: []ie.SliceSupportItem{
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
											SD: &ie.SD{
												Value: aper.OctetString([]byte{0x01, 0x02, 0x03}),
											},
											IEExtensions: nil,
										},
										IEExtensions: nil,
									},
								},
							},
							IEExtensions: &ie.ProtocolExtensionContainerPLMNSupportItemExtIEs{
								List: []ie.PLMNSupportItemExtIEs{
									{
										Id: &ie.ProtocolExtensionID{
											Value: 258,
										},
										Criticality: &ie.Criticality{
											Value: aper.Enumerated(0),
										},
										NPNSupport: &ie.NPNSupport{
											Choice: &ie.NID{
												Value: aper.BitString{
													Bytes:     []byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x10},
													BitLength: 44,
												},
											},
										},
										ExtendedSliceSupportList: nil,
										OnboardingSupport:        nil,
									},
								},
							},
						},
					},
				},
				CriticalityDiagnostics: nil,
				UERetentionInformation: nil,
				IABSupported:           nil,
				ExtendedAMFName:        nil,
			},
			expected: []byte{
				0x20, 0x15, 0x00, 0x3a, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x06, 0x01, 0x80, 0x61, 0x6d, 0x66,
				0x31, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x01, 0x00, 0x41, 0x00, 0x56, 0x40,
				0x01, 0xff, 0x00, 0x50, 0x00, 0x18, 0x04, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x10, 0x08, 0x01, 0x02,
				0x03, 0x00, 0x00, 0x01, 0x02, 0x00, 0x07, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x10,
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

func TestNGSetupResponseUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *NGSetupResponse
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x15, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
				0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
				0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
				0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
			},
			expected: &NGSetupResponse{
				AMFName: &ie.AMFName{
					Value: aper.PrintableString("AMF"),
				},
				ServedGUAMIList: &ie.ServedGUAMIList{
					List: []ie.ServedGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: []byte{0x13, 0x30, 0x01},
								},
								// AMFID: "cafe00",
								AMFRegionID: &ie.AMFRegionID{
									Value: aper.BitString{
										Bytes:     []byte{0xca},
										BitLength: 8,
									},
								},
								AMFSetID: &ie.AMFSetID{
									Value: aper.BitString{
										Bytes:     []byte{0xfe, 0x00},
										BitLength: 10,
									},
								},
								AMFPointer: &ie.AMFPointer{
									Value: aper.BitString{
										Bytes:     []byte{0x00},
										BitLength: 6,
									},
								},
							},
						},
					},
				},
				RelativeAMFCapacity: &ie.RelativeAMFCapacity{
					Value: 255,
				},
				PLMNSupportList: &ie.PLMNSupportList{
					List: []ie.PLMNSupportItem{
						{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x13, 0x30, 0x01},
							},
							SliceSupportList: &ie.SliceSupportList{
								List: []ie.SliceSupportItem{
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
											SD: &ie.SD{
												Value: aper.OctetString([]byte{0x01, 0x02, 0x03}),
											},
										},
									},
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
											SD: &ie.SD{
												Value: aper.OctetString([]byte{0x11, 0x22, 0x33}),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Case 2 (R16)",
			input: []byte{ // Contains R16 Fields
				0x20, 0x15, 0x00, 0x4a, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x16, 0x09, 0x80, 0x73,
				0x72, 0x63, 0x30, 0x2e, 0x41, 0x4d, 0x46, 0x2e, 0x53, 0x70, 0x69, 0x72, 0x65, 0x6e,
				0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x56, 0x40, 0x01, 0x0a, 0x00, 0x50, 0x00, 0x18, 0x04, 0x00,
				0x00, 0x00, 0x00, 0x02, 0x00, 0x08, 0x00, 0x80, 0x06, 0x00, 0x00, 0x01, 0x02, 0x00,
				0x07, 0x00, 0x99, 0x99, 0x99, 0x99, 0x99, 0x90,
			},
			expected: &NGSetupResponse{
				AMFName: &ie.AMFName{
					Value: aper.PrintableString("src0.AMF.Spirent.com"),
				},
				ServedGUAMIList: &ie.ServedGUAMIList{
					List: []ie.ServedGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: []byte{0x00, 0x00, 0x00},
								},
								// AMFID: "000000",
								AMFRegionID: &ie.AMFRegionID{
									Value: aper.BitString{
										Bytes:     []byte{0x00},
										BitLength: 8,
									},
								},
								AMFSetID: &ie.AMFSetID{
									Value: aper.BitString{
										Bytes:     []byte{0x00, 0x00},
										BitLength: 10,
									},
								},
								AMFPointer: &ie.AMFPointer{
									Value: aper.BitString{
										Bytes:     []byte{0x00},
										BitLength: 6,
									},
								},
							},
						},
					},
				},
				RelativeAMFCapacity: &ie.RelativeAMFCapacity{
					Value: 10,
				},
				PLMNSupportList: &ie.PLMNSupportList{
					List: []ie.PLMNSupportItem{
						{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x00, 0x00, 0x00},
							},
							SliceSupportList: &ie.SliceSupportList{
								List: []ie.SliceSupportItem{
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
										},
									},
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x02}),
											},
										},
									},
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x03}),
											},
										},
									},
								},
							},
							IEExtensions: &ie.ProtocolExtensionContainerPLMNSupportItemExtIEs{
								List: []ie.PLMNSupportItemExtIEs{
									{
										Id: &ie.ProtocolExtensionID{
											Value: ie.ProtocolIEIDNPNSupport,
										},
										Criticality: &ie.Criticality{
											Value: ie.CriticalityReject,
										},
										NPNSupport: &ie.NPNSupport{
											Choice: &ie.NID{
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
					},
				},
			},
		},
		{
			name: "Case 3: SNPN mode (from ueranemu SNPN pipeline TestRegistration)",
			input: []byte{
				0x20, 0x15, 0x00, 0x3a, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x06, 0x01, 0x80, 0x61, 0x6d, 0x66,
				0x31, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x01, 0x00, 0x41, 0x00, 0x56, 0x40,
				0x01, 0xff, 0x00, 0x50, 0x00, 0x18, 0x04, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x10, 0x08, 0x01, 0x02,
				0x03, 0x00, 0x00, 0x01, 0x02, 0x00, 0x07, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x10,
			},
			expected: &NGSetupResponse{
				AMFName: &ie.AMFName{
					Value: aper.PrintableString("amf1"),
				},
				ServedGUAMIList: &ie.ServedGUAMIList{
					List: []ie.ServedGUAMIItem{
						{
							GUAMI: &ie.GUAMI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
								},
								AMFRegionID: &ie.AMFRegionID{
									Value: aper.BitString{
										Bytes:     []byte{0x01},
										BitLength: 8,
									},
								},
								AMFSetID: &ie.AMFSetID{
									Value: aper.BitString{
										Bytes:     []byte{0x00, 0x40},
										BitLength: 10,
									},
								},
								AMFPointer: &ie.AMFPointer{
									Value: aper.BitString{
										Bytes:     []byte{0x04},
										BitLength: 6,
									},
								},
								IEExtensions: nil,
							},
							BackupAMFName: nil,
							IEExtensions:  nil,
						},
					},
				},
				RelativeAMFCapacity: &ie.RelativeAMFCapacity{
					Value: 255,
				},
				PLMNSupportList: &ie.PLMNSupportList{
					List: []ie.PLMNSupportItem{
						{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString([]byte{0x02, 0xf8, 0x39}),
							},
							SliceSupportList: &ie.SliceSupportList{
								List: []ie.SliceSupportItem{
									{
										SNSSAI: &ie.SNSSAI{
											SST: &ie.SST{
												Value: aper.OctetString([]byte{0x01}),
											},
											SD: &ie.SD{
												Value: aper.OctetString([]byte{0x01, 0x02, 0x03}),
											},
											IEExtensions: nil,
										},
										IEExtensions: nil,
									},
								},
							},
							IEExtensions: &ie.ProtocolExtensionContainerPLMNSupportItemExtIEs{
								List: []ie.PLMNSupportItemExtIEs{
									{
										Id: &ie.ProtocolExtensionID{
											Value: 258,
										},
										Criticality: &ie.Criticality{
											Value: aper.Enumerated(0),
										},
										NPNSupport: &ie.NPNSupport{
											Choice: &ie.NID{
												Value: aper.BitString{
													Bytes:     []byte{0x10, 0x00, 0x00, 0x00, 0x00, 0x10},
													BitLength: 44,
												},
											},
										},
										ExtendedSliceSupportList: nil,
										OnboardingSupport:        nil,
									},
								},
							},
						},
					},
				},
				CriticalityDiagnostics: nil,
				UERetentionInformation: nil,
				IABSupported:           nil,
				ExtendedAMFName:        nil,
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

func TestNGSetupResponseUnmarshalBinaryFailure(t *testing.T) {
	// Correct version:
	// 0x20, 0x15, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
	// 0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
	// 0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
	// 0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,

	// Transfer Syntax Error:
	// third byte will cause aper.ReadEnumerated failed
	fmt.Println("=========================")
	fmt.Println("Transfer Syntax Error")
	bytes := []byte{
		0x20, 0x15, 0xFF, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
		0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
		0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
		0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
	}
	_, err := Parse(bytes)
	xferSyntaxErr, ok := errors.Cause(err).(*ie.TransferSyntaxErr)
	require.True(t, ok)
	fmt.Println("--- xferSyntaxErr.Error() rsult:")
	fmt.Println(xferSyntaxErr.Error())
	fmt.Println()
	fmt.Println("--- xferSyntaxErr.ErrorTrace() result:")
	fmt.Println(xferSyntaxErr.ErrorTrace())
	fmt.Println("--- xferSyntaxErr.GetCause() result:")
	fmt.Printf("\tcause: %+v\n", xferSyntaxErr.GetCause())

	// Abstract Syntax Error:
	// case a. Not Comprehended IE/IE group
	// 	a-1. procedure code not comprehended (modify second byte)
	fmt.Println("=========================")
	fmt.Println("Abstract Syntax Error (a-1.)")
	bytes = []byte{
		0x20, 0x40, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
		0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
		0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
		0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
	}
	_, err = Parse(bytes)
	abstractSyntaxErr, ok := errors.Cause(err).(*ie.AbstractSyntaxErr)
	require.True(t, ok)
	fmt.Println("--- abstractSyntaxErr.Error() rsult:")
	fmt.Println(abstractSyntaxErr.Error())
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.ErrorTrace() result:")
	fmt.Println(abstractSyntaxErr.ErrorTrace())
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.GetCause() result:")
	// fmt.Println("++++++++++++++++++++")
	// fmt.Printf("%v", err)
	// fmt.Println()
	// fmt.Println("-------------------")
	// fmt.Printf("%+v", err)
	cause, err := abstractSyntaxErr.GetCause()
	require.Error(t, err)
	require.Nil(t, cause)
	fmt.Printf("cause: %+v, err: %v\n", cause, err)
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.GetCritDiag() result:")
	critDiag, err := abstractSyntaxErr.GetCritDiag(true)
	require.NoError(t, err)
	fmt.Printf("critDiag: %+v, err: %v\n", critDiag, err)

	// 	a-2. message type not comprehended
	fmt.Println("=========================")
	fmt.Println("Abstract Syntax Error (a-2.)")
	bytes = []byte{
		0x99, 0x15, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
		0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
		0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
		0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
	}
	_, err = Parse(bytes)
	abstractSyntaxErr, ok = errors.Cause(err).(*ie.AbstractSyntaxErr)
	require.True(t, ok)
	fmt.Println("--- abstractSyntaxErr.Error() rsult:")
	fmt.Println(abstractSyntaxErr.Error())
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.ErrorTrace() result:")
	fmt.Println(abstractSyntaxErr.ErrorTrace())
	fmt.Println("--- abstractSyntaxErr.GetCause() result:")
	cause, err = abstractSyntaxErr.GetCause()
	require.NoError(t, err)
	fmt.Printf("cause: %+v, err: %v\n", cause, err)
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.GetCritDiag() result:")
	critDiag, err = abstractSyntaxErr.GetCritDiag(true)
	require.Error(t, err)
	fmt.Printf("critDiag: %+v, err: %v\n", critDiag, err)

	// 	a-3. ies other than procedure code and message type not comprehended
	// 	change the IEID of ServedGUAMIList (0x60) to 0x61
	fmt.Println("=========================")
	fmt.Println("Abstract Syntax Error (a-3.)")
	bytes = []byte{
		0x20, 0x15, 0x00, 0x31, 0x00, 0x00, 0x04, 0x00, 0x01, 0x00, 0x05,
		0x01, 0x00, 0x41, 0x4d, 0x46, 0x00, 0x61, 0x00, 0x08, 0x00, 0x00, 0x13, 0x30, 0x01,
		0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50, 0x00, 0x10, 0x00, 0x13,
		0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
	}
	_, err = Parse(bytes)
	abstractSyntaxErr, ok = errors.Cause(err).(*ie.AbstractSyntaxErr)
	require.True(t, ok)
	fmt.Println("--- abstractSyntaxErr.Error() rsult:")
	fmt.Println(abstractSyntaxErr.Error())
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.ErrorTrace() result:")
	fmt.Println(abstractSyntaxErr.ErrorTrace())
	fmt.Println("--- abstractSyntaxErr.GetCause() result:")
	cause, err = abstractSyntaxErr.GetCause()
	require.NoError(t, err)
	fmt.Printf("cause: %+v, err: %v\n", cause, err)
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.GetCritDiag() result:")
	critDiag, err = abstractSyntaxErr.GetCritDiag(true)
	require.NoError(t, err)
	fmt.Printf("critDiag: %+v, err: %v\n", critDiag, err)

	// Abstract Syntax Error:
	// case b. Missing IE/IE group
	fmt.Println("=========================")
	fmt.Println("Abstract Syntax Error (b.)")
	bytes = []byte{
		0x20, 0x15, 0x00, 0x28, 0x00, 0x00, 0x03, 0x00, 0x60, 0x00, 0x08, 0x00, 0x00,
		0x13, 0x30, 0x01, 0xca, 0xfe, 0x00, 0x00, 0x56, 0x40, 0x01, 0xff, 0x00, 0x50,
		0x00, 0x10, 0x00, 0x13, 0x30, 0x01, 0x00, 0x01, 0x10, 0x08, 0x01, 0x02, 0x03,
		0x10, 0x08, 0x11, 0x22, 0x33,
	}
	_, err = Parse(bytes)
	abstractSyntaxErr, ok = errors.Cause(err).(*ie.AbstractSyntaxErr)
	require.True(t, ok)
	fmt.Println("--- abstractSyntaxErr.Error() rsult:")
	fmt.Println(abstractSyntaxErr.Error())
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.ErrorTrace() result:")
	fmt.Println(abstractSyntaxErr.ErrorTrace())
	fmt.Println("--- abstractSyntaxErr.GetCause() result:")
	cause, err = abstractSyntaxErr.GetCause()
	require.NoError(t, err)
	fmt.Printf("cause: %+v, err: %v\n", cause, err)
	fmt.Println()
	fmt.Println("--- abstractSyntaxErr.GetCritDiag() result:")
	critDiag, err = abstractSyntaxErr.GetCritDiag(true)
	require.NoError(t, err)
	fmt.Printf("critDiag: %+v, err: %v\n", critDiag, err)
}
