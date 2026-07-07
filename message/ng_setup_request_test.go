package message

import (
	"encoding/hex"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestNGSetupRequestMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *NGSetupRequest
		expected []byte
	}{
		{
			name: "Case 1",
			input: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: []byte{0x13, 0x30, 0x01},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00},
									BitLength: 32,
								},
							},
						},
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("gnb.spirent.com"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: []byte{0x13, 0x30, 0x01},
										},
										TAISliceSupportList: &ie.SliceSupportList{
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
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: ie.PagingDRXPresentV32,
				},
			},
			expected: []byte{
				0x00, 0x15, 0x00, 0x43, 0x00, 0x00, 0x04, 0x00, 0x1b, 0x00, 0x09,
				0x00, 0x13, 0x30, 0x01, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52,
				0x40, 0x11, 0x07, 0x00, 0x67, 0x6e, 0x62, 0x2e, 0x73, 0x70, 0x69,
				0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x66, 0x00,
				0x15, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x13, 0x30, 0x01, 0x00,
				0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
				0x00, 0x15, 0x40, 0x01, 0x00,
			},
		},
		{
			name: "Case 2 (R16)",
			input: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: []byte{0x00, 0x00, 0x00},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00},
									BitLength: 32,
								},
							},
						},
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("src0.gNB.Spirent.com"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: []byte{0x00, 0x00, 0x00},
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: []byte{0x00, 0x00, 0x00},
										},
										TAISliceSupportList: &ie.SliceSupportList{
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
										IEExtensions: &ie.ProtocolExtensionContainerBroadcastPLMNItemExtIEs{
											List: []ie.BroadcastPLMNItemExtIEs{
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
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: ie.PagingDRXPresentV32,
				},
			},
			expected: []byte{ // Contains R16 Fields
				0x00, 0x15, 0x00, 0x50, 0x00, 0x00, 0x04, 0x00, 0x1b, 0x00, 0x09,
				0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52,
				0x40, 0x16, 0x09, 0x80, 0x73, 0x72, 0x63, 0x30, 0x2e, 0x67, 0x4e,
				0x42, 0x2e, 0x53, 0x70, 0x69, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63,
				0x6f, 0x6d, 0x00, 0x66, 0x00, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x04, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x08, 0x00, 0x80, 0x06,
				0x00, 0x00, 0x01, 0x02, 0x00, 0x07, 0x00, 0x99, 0x99, 0x99, 0x99,
				0x99, 0x90, 0x00, 0x15, 0x40, 0x01, 0x00,
			},
		},
		{
			name: "Case 3: SNPN mode (from ueranemu SNPN pipeline TestRegistration)",

			expected: []byte{
				0x00, 0x15, 0x00, 0x44, 0x00, 0x00, 0x04, 0x00, 0x1b, 0x00, 0x08, 0x00, 0x02, 0xf8, 0x39, 0x00,
				0x00, 0x00, 0x04, 0x00, 0x52, 0x40, 0x0b, 0x04, 0x00, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63,
				0x2d, 0x31, 0x00, 0x66, 0x00, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x01, 0x04, 0x02, 0xf8, 0x39, 0x00,
				0x00, 0x10, 0x08, 0x01, 0x02, 0x03, 0x00, 0x00, 0x01, 0x02, 0x00, 0x07, 0x00, 0x10, 0x00, 0x00,
				0x00, 0x00, 0x10, 0x00, 0x15, 0x40, 0x01, 0x40,
			},
			input: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: aper.OctetString{
								0x02, 0xf8, 0x39,
							},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x04},
									BitLength: 22,
								},
							},
						},
						IEExtensions: nil,
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("free5gc-1"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: aper.OctetString{
									0x00, 0x00, 0x01,
								},
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: aper.OctetString{
												0x02, 0xf8, 0x39,
											},
										},
										TAISliceSupportList: &ie.SliceSupportList{
											List: []ie.SliceSupportItem{
												{
													SNSSAI: &ie.SNSSAI{
														SST: &ie.SST{
															Value: aper.OctetString{
																0x01,
															},
														},
														SD: &ie.SD{
															Value: aper.OctetString{
																0x01, 0x02, 0x03,
															},
														},
														IEExtensions: nil,
													},
													IEExtensions: nil,
												},
											},
										},
										IEExtensions: &ie.ProtocolExtensionContainerBroadcastPLMNItemExtIEs{
											List: []ie.BroadcastPLMNItemExtIEs{
												{
													Id: &ie.ProtocolExtensionID{
														Value: int64(258),
													},
													Criticality: &ie.Criticality{
														Value: aper.Enumerated(0),
													},
													NPNSupport: &ie.NPNSupport{
														Choice: &ie.NID{
															Value: aper.BitString{
																Bytes: []byte{
																	0x10, 0x00, 0x00, 0x00, 0x00, 0x10,
																},
																BitLength: 44,
															},
														},
													},
													ExtendedTAISliceSupportList: nil,
													TAINSAGSupportList:          nil,
												},
											},
										},
									},
								},
							},
							IEExtensions: nil,
						},
					},
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: aper.Enumerated(2),
				},
				UERetentionInformation: nil,
				NBIoTDefaultPagingDRX:  nil,
				ExtendedRANNodeName:    nil,
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

func TestNGSetupRequestUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *NGSetupRequest
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x15, 0x00, 0x43, 0x00, 0x00, 0x04, 0x00, 0x1b, 0x00, 0x09,
				0x00, 0x13, 0x30, 0x01, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52,
				0x40, 0x11, 0x07, 0x00, 0x67, 0x6e, 0x62, 0x2e, 0x73, 0x70, 0x69,
				0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x66, 0x00,
				0x15, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x13, 0x30, 0x01, 0x00,
				0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
				0x00, 0x15, 0x40, 0x01, 0x00,
			},
			expected: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: []byte{0x13, 0x30, 0x01},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00},
									BitLength: 32,
								},
							},
						},
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("gnb.spirent.com"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: []byte{0x13, 0x30, 0x01},
										},
										TAISliceSupportList: &ie.SliceSupportList{
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
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: ie.PagingDRXPresentV32,
				},
			},
		},
		{
			name: "Case 2 (R16)",
			input: []byte{ // Contains R16 Fields
				0x00, 0x15, 0x00, 0x50, 0x00, 0x00, 0x04, 0x00, 0x1b, 0x00, 0x09,
				0x00, 0x00, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52,
				0x40, 0x16, 0x09, 0x80, 0x73, 0x72, 0x63, 0x30, 0x2e, 0x67, 0x4e,
				0x42, 0x2e, 0x53, 0x70, 0x69, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63,
				0x6f, 0x6d, 0x00, 0x66, 0x00, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x04, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x08, 0x00, 0x80, 0x06,
				0x00, 0x00, 0x01, 0x02, 0x00, 0x07, 0x00, 0x99, 0x99, 0x99, 0x99,
				0x99, 0x90, 0x00, 0x15, 0x40, 0x01, 0x00,
			},
			expected: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: []byte{0x00, 0x00, 0x00},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00},
									BitLength: 32,
								},
							},
						},
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("src0.gNB.Spirent.com"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: []byte{0x00, 0x00, 0x00},
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: []byte{0x00, 0x00, 0x00},
										},
										TAISliceSupportList: &ie.SliceSupportList{
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
										IEExtensions: &ie.ProtocolExtensionContainerBroadcastPLMNItemExtIEs{
											List: []ie.BroadcastPLMNItemExtIEs{
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
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: ie.PagingDRXPresentV32,
				},
			},
		},
		{
			name: "Case 3 (With Unknown IEs of Criticality ignore)",
			input: []byte{
				0x00, 0x15, 0x00, 0x58, // len=0x58 bytes
				0x00, 0x00,
				0x05, // 5 ies
				0x00, 0x1b, 0x00, 0x09,
				0x00, 0x13, 0x30, 0x01, 0x50, 0x00, 0x00, 0x00, 0x00,
				// Ran Node Name starts from here
				0x00, 0x52,
				0x40, 0x11, 0x07, 0x00, 0x67, 0x6e, 0x62, 0x2e, 0x73, 0x70, 0x69,
				0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
				// Add an unknown IE (ID=0x99, Criticality=ignore) here
				0x00, 0x99,
				0x40, 0x11, 0x07, 0x00, 0x67, 0x6e, 0x62, 0x2e, 0x73, 0x70, 0x69,
				0x72, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d,
				// Supported TAI List starts from here
				0x00, 0x66, 0x00,
				0x15, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x13, 0x30, 0x01, 0x00,
				0x01, 0x10, 0x08, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
				0x00, 0x15, 0x40, 0x01, 0x00,
			},
			expected: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: []byte{0x13, 0x30, 0x01},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x00, 0x00},
									BitLength: 32,
								},
							},
						},
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("gnb.spirent.com"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: aper.OctetString([]byte{0x00, 0x00, 0x01}),
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: []byte{0x13, 0x30, 0x01},
										},
										TAISliceSupportList: &ie.SliceSupportList{
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
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: ie.PagingDRXPresentV32,
				},
			},
		},
		{
			name: "Case 4: SNPN mode (from ueranemu SNPN pipeline TestRegistration)",
			input: []byte{
				0x00, 0x15, 0x00, 0x44, 0x00, 0x00, 0x04, 0x00, 0x1b, 0x00, 0x08, 0x00, 0x02, 0xf8, 0x39, 0x00,
				0x00, 0x00, 0x04, 0x00, 0x52, 0x40, 0x0b, 0x04, 0x00, 0x66, 0x72, 0x65, 0x65, 0x35, 0x67, 0x63,
				0x2d, 0x31, 0x00, 0x66, 0x00, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x01, 0x04, 0x02, 0xf8, 0x39, 0x00,
				0x00, 0x10, 0x08, 0x01, 0x02, 0x03, 0x00, 0x00, 0x01, 0x02, 0x00, 0x07, 0x00, 0x10, 0x00, 0x00,
				0x00, 0x00, 0x10, 0x00, 0x15, 0x40, 0x01, 0x40,
			},
			expected: &NGSetupRequest{
				GlobalRANNodeID: &ie.GlobalRANNodeID{
					Choice: &ie.GlobalGNBID{
						PLMNIdentity: &ie.PLMNIdentity{
							Value: aper.OctetString{
								0x02, 0xf8, 0x39,
							},
						},
						GNBID: &ie.GNBID{
							Choice: &ie.GNBIDForGNBID{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x00, 0x04},
									BitLength: 22,
								},
							},
						},
						IEExtensions: nil,
					},
				},
				RANNodeName: &ie.RANNodeName{
					Value: aper.PrintableString("free5gc-1"),
				},
				SupportedTAList: &ie.SupportedTAList{
					List: []ie.SupportedTAItem{
						{
							TAC: &ie.TAC{
								Value: aper.OctetString{
									0x00, 0x00, 0x01,
								},
							},
							BroadcastPLMNList: &ie.BroadcastPLMNList{
								List: []ie.BroadcastPLMNItem{
									{
										PLMNIdentity: &ie.PLMNIdentity{
											Value: aper.OctetString{
												0x02, 0xf8, 0x39,
											},
										},
										TAISliceSupportList: &ie.SliceSupportList{
											List: []ie.SliceSupportItem{
												{
													SNSSAI: &ie.SNSSAI{
														SST: &ie.SST{
															Value: aper.OctetString{
																0x01,
															},
														},
														SD: &ie.SD{
															Value: aper.OctetString{
																0x01, 0x02, 0x03,
															},
														},
														IEExtensions: nil,
													},
													IEExtensions: nil,
												},
											},
										},
										IEExtensions: &ie.ProtocolExtensionContainerBroadcastPLMNItemExtIEs{
											List: []ie.BroadcastPLMNItemExtIEs{
												{
													Id: &ie.ProtocolExtensionID{
														Value: int64(258),
													},
													Criticality: &ie.Criticality{
														Value: aper.Enumerated(0),
													},
													NPNSupport: &ie.NPNSupport{
														Choice: &ie.NID{
															Value: aper.BitString{
																Bytes: []byte{
																	0x10, 0x00, 0x00, 0x00, 0x00, 0x10,
																},
																BitLength: 44,
															},
														},
													},
													ExtendedTAISliceSupportList: nil,
													TAINSAGSupportList:          nil,
												},
											},
										},
									},
								},
							},
							IEExtensions: nil,
						},
					},
				},
				DefaultPagingDRX: &ie.PagingDRX{
					Value: aper.Enumerated(2),
				},
				UERetentionInformation: nil,
				NBIoTDefaultPagingDRX:  nil,
				ExtendedRANNodeName:    nil,
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

func TestUnknownIEID(t *testing.T) {
	hs := "001500" + "46" + "000005" +
		"001b00080002f8390000000400524006018053494d300066001000000000010002f83900001008010203" +
		"01f440" + "1000000000010000f11000001008000001" +
		"0015400140"
	bs, _ := hex.DecodeString(hs)
	msg, err := Parse(bs)
	require.NoError(t, err)
	spew.Dump(msg)
}
