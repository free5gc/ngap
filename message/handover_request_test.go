package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestHandoverRequestMarshalBinary(t *testing.T) {
	t.Parallel()

	// 172.16.3.100
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x64},
			BitLength: 32,
		},
	}
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	sst_case1 := &ie.SST{Value: aper.OctetString([]byte{0x01})}
	sd1_case1 := &ie.SD{Value: aper.OctetString([]byte{0x01, 0x02, 0x03})}
	sd2_case1 := &ie.SD{Value: aper.OctetString([]byte{0x11, 0x22, 0x33})}
	nrEncryptAlgo := &ie.NRencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	nrIntegrityAlgo := &ie.NRintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x40, 0x00},
			BitLength: 16,
		},
	}
	eutraEncryptAlgo := &ie.EUTRAencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	eutraIntegrityAlgo := &ie.EUTRAintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	securityKey := &ie.SecurityKey{
		Value: aper.BitString{
			Bytes: []byte{
				0xe8, 0x23, 0xf9, 0x91, 0xbe, 0xc1, 0x19, 0xb6, 0xe0, 0xd1,
				0x08, 0x1b, 0xe6, 0x34, 0x09, 0xeb, 0x93, 0x16, 0x84, 0x84, 0x8f, 0x9b,
				0x76, 0x4f, 0x80, 0x44, 0x5e, 0x2d, 0x64, 0xb3, 0x71, 0x53,
			},
			BitLength: 256,
		},
	}

	// AMFID: "cafe00"
	amfRegionId := &ie.AMFRegionID{Value: aper.BitString{Bytes: []byte{0xCA}, BitLength: 8}}
	amfSetId := &ie.AMFSetID{Value: aper.BitString{Bytes: []byte{0xFE, 0x00}, BitLength: 10}}
	amfPointer := &ie.AMFPointer{Value: aper.BitString{Bytes: []byte{0x00}, BitLength: 6}}

	// pduSessionResourceSetupRequestTransfer
	gtpteid := &ie.GTPTEID{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	pduSessionResourceSetupRequestTransfer := ie.PDUSessionResourceSetupRequestTransfer{
		ProtocolIEs: &ie.ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
			List: []ie.PDUSessionResourceSetupRequestTransferIEs{
				{
					PDUSessionAggregateMaximumBitRate: &ie.PDUSessionAggregateMaximumBitRate{
						PDUSessionAggregateMaximumBitRateDL: &ie.BitRate{
							Value: 100000000,
						},
						PDUSessionAggregateMaximumBitRateUL: &ie.BitRate{
							Value: 200000000,
						},
					},
				},
				{
					ULNGUUPTNLInformation: &ie.UPTransportLayerInformation{
						Choice: &ie.GTPTunnel{
							TransportLayerAddress: ipv4Addr,
							GTPTEID:               gtpteid,
						},
					},
				},
				{
					PDUSessionType: &ie.PDUSessionType{
						Value: ie.PDUSessionTypePresentIpv4,
					},
				},
				{
					QosFlowSetupRequestList: &ie.QosFlowSetupRequestList{
						List: []ie.QosFlowSetupRequestItem{
							{
								QosFlowIdentifier: &ie.QosFlowIdentifier{
									Value: 9,
								},
								QosFlowLevelQosParameters: &ie.QosFlowLevelQosParameters{
									QosCharacteristics: &ie.QosCharacteristics{
										Choice: &ie.NonDynamic5QIDescriptor{
											FiveQI: &ie.FiveQI{
												Value: 9,
											},
										},
									},
									AllocationAndRetentionPriority: &ie.AllocationAndRetentionPriority{
										PriorityLevelARP: &ie.PriorityLevelARP{
											Value: 15,
										},
										PreEmptionCapability: &ie.PreEmptionCapability{
											Value: ie.PreEmptionCapabilityPresentShallNotTriggerPreEmption,
										},
										PreEmptionVulnerability: &ie.PreEmptionVulnerability{
											Value: ie.PreEmptionVulnerabilityPresentNotPreEmptable,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	pduSessionResourceSetupRequestTransferBytes, err := ie.MarshalBinary(&pduSessionResourceSetupRequestTransfer)
	require.NoError(t, err)
	pduSessionResourceSetupRequestTransferOS := aper.OctetString(
		pduSessionResourceSetupRequestTransferBytes)

	// sourceNGRANNodeToTargetNGRANNodeTransparentContainer
	targetCellNrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x0d, 0x90, 0x30},
			BitLength: 36,
		},
	}
	lastVisitedCellNrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x02, 0xb6, 0x70},
			BitLength: 36,
		},
	}
	sourceNGRANNodeToTargetNGRANNodeTransparentContainer := ie.SourceNGRANNodeToTargetNGRANNodeTransparentContainer{
		RRCContainer: &ie.RRCContainer{
			Value: aper.OctetString([]byte{
				0x10, 0x00, 0x09,
				0x33, 0x31, 0x30, 0x33, 0x31, 0x30, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30,
				0x31, 0x32, 0x30,
			}),
		},
		PDUSessionResourceInformationList: &ie.PDUSessionResourceInformationList{
			List: []ie.PDUSessionResourceInformationItem{
				{
					PDUSessionID: &ie.PDUSessionID{
						Value: 5,
					},
					QosFlowInformationList: &ie.QosFlowInformationList{
						List: []ie.QosFlowInformationItem{
							{
								QosFlowIdentifier: &ie.QosFlowIdentifier{
									Value: 9,
								},
							},
						},
					},
				},
			},
		},
		TargetCellID: &ie.NGRANCGI{
			Choice: &ie.NRCGI{
				PLMNIdentity:   plmnIdentity_case1,
				NRCellIdentity: targetCellNrCellIdentity,
			},
		},
		UEHistoryInformation: &ie.UEHistoryInformation{
			List: []ie.LastVisitedCellItem{
				{
					LastVisitedCellInformation: &ie.LastVisitedCellInformation{
						Choice: &ie.LastVisitedNGRANCellInformation{
							GlobalCellID: &ie.NGRANCGI{
								Choice: &ie.NRCGI{
									PLMNIdentity:   plmnIdentity_case1,
									NRCellIdentity: lastVisitedCellNrCellIdentity,
								},
							},
							CellType: &ie.CellType{
								CellSize: &ie.CellSize{
									Value: ie.CellSizePresentMedium,
								},
							},
							TimeUEStayedInCell: &ie.TimeUEStayedInCell{
								Value: 100,
							},
						},
					},
				},
			},
		},
	}
	sourceNGRANNodeToTargetNGRANNodeTransparentContainerBytes, err := ie.MarshalBinary(
		&sourceNGRANNodeToTargetNGRANNodeTransparentContainer,
	)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    *HandoverRequest
		expected []byte
	}{
		{
			name: "Case 1",
			input: &HandoverRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				HandoverType: &ie.HandoverType{
					Value: ie.HandoverTypePresentIntra5gs,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentNgIntraSystemHandoverTriggered,
					},
				},
				UEAggregateMaximumBitRate: &ie.UEAggregateMaximumBitRate{
					UEAggregateMaximumBitRateDL: &ie.BitRate{
						Value: 2000000000,
					},
					UEAggregateMaximumBitRateUL: &ie.BitRate{
						Value: 1000000000,
					},
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms:             nrEncryptAlgo,
					NRintegrityProtectionAlgorithms:    nrIntegrityAlgo,
					EUTRAencryptionAlgorithms:          eutraEncryptAlgo,
					EUTRAintegrityProtectionAlgorithms: eutraIntegrityAlgo,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: securityKey,
				},
				PDUSessionResourceSetupListHOReq: &ie.PDUSessionResourceSetupListHOReq{
					List: []ie.PDUSessionResourceSetupItemHOReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd1_case1,
							},
							HandoverRequestTransfer: &pduSessionResourceSetupRequestTransferOS,
						},
					},
				},
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd1_case1,
							},
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd2_case1,
							},
						},
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: sourceNGRANNodeToTargetNGRANNodeTransparentContainerBytes,
				},
				GUAMI: &ie.GUAMI{
					PLMNIdentity: plmnIdentity_case1,
					AMFRegionID:  amfRegionId,
					AMFSetID:     amfSetId,
					AMFPointer:   amfPointer,
				},
			},
			expected: []byte{
				0x00, 0x0d, 0x00, 0x80, 0xdd, 0x00, 0x00, 0x0a, 0x00, 0x0a,
				0x00, 0x02, 0x00, 0x02, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00,
				0x0f, 0x40, 0x02, 0x07, 0xc0, 0x00, 0x6e, 0x00, 0x0a, 0x0c,
				0x77, 0x35, 0x94, 0x00, 0x30, 0x3b, 0x9a, 0xca, 0x00, 0x00,
				0x77, 0x00, 0x09, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x5d, 0x00, 0x21, 0x10, 0xe8, 0x23, 0xf9,
				0x91, 0xbe, 0xc1, 0x19, 0xb6, 0xe0, 0xd1, 0x08, 0x1b, 0xe6,
				0x34, 0x09, 0xeb, 0x93, 0x16, 0x84, 0x84, 0x8f, 0x9b, 0x76,
				0x4f, 0x80, 0x44, 0x5e, 0x2d, 0x64, 0xb3, 0x71, 0x53, 0x00,
				0x49, 0x00, 0x38, 0x00, 0x00, 0x05, 0x40, 0x20, 0x01, 0x02,
				0x03, 0x2f, 0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x0a, 0x0c,
				0x05, 0xf5, 0xe1, 0x00, 0x30, 0x0b, 0xeb, 0xc2, 0x00, 0x00,
				0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x03, 0x64, 0x00,
				0x00, 0x00, 0x01, 0x00, 0x86, 0x00, 0x01, 0x00, 0x00, 0x88,
				0x00, 0x07, 0x00, 0x09, 0x00, 0x00, 0x09, 0x38, 0x00, 0x00,
				0x00, 0x00, 0x0a, 0x22, 0x01, 0x01, 0x02, 0x03, 0x10, 0x08,
				0x11, 0x22, 0x33, 0x00, 0x65, 0x00, 0x30, 0x2f, 0x40, 0x12,
				0x10, 0x00, 0x09, 0x33, 0x31, 0x30, 0x33, 0x31, 0x30, 0x31,
				0x34, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x30, 0x00, 0x00,
				0x05, 0x00, 0x09, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x0d,
				0x90, 0x30, 0x00, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x02,
				0xb6, 0x71, 0x00, 0x00, 0x64, 0x00, 0x1c, 0x00, 0x07, 0x00,
				0x13, 0x30, 0x01, 0xca, 0xfe, 0x00,
			},
		},
		{
			name: "Case 2: from ueranemu pipeline basic-k8s TestN2Handover",
			input: &HandoverRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 887252622092,
				},
				HandoverType: &ie.HandoverType{
					Value: aper.Enumerated(0),
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHandoverDesirableForRadioReason,
					},
				},
				UEAggregateMaximumBitRate: &ie.UEAggregateMaximumBitRate{
					UEAggregateMaximumBitRateDL: &ie.BitRate{
						Value: 1000000,
					},
					UEAggregateMaximumBitRateUL: &ie.BitRate{
						Value: 1000000,
					},
					IEExtensions: nil,
				},
				CoreNetworkAssistanceInformationForInactive: nil,
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms: &ie.NRencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0xc0, 0x00},
							BitLength: 16,
						},
					},
					NRintegrityProtectionAlgorithms: &ie.NRintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0xc0, 0x00},
							BitLength: 16,
						},
					},
					EUTRAencryptionAlgorithms: &ie.EUTRAencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0x00, 0x00},
							BitLength: 16,
						},
					},
					EUTRAintegrityProtectionAlgorithms: &ie.EUTRAintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0x00, 0x00},
							BitLength: 16,
						},
					},
					IEExtensions: nil,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: &ie.SecurityKey{
						Value: aper.BitString{
							Bytes: []uint8{
								0x48, 0xb9, 0xb0, 0xeb, 0xe8, 0xfc, 0x7f, 0x0c, 0x1f, 0x78, 0x42, 0xb6, 0x6e, 0x11,
								0xed, 0x29, 0x65, 0x78, 0x37, 0x39, 0x4a, 0xc9, 0x2d, 0x0c, 0x76, 0xee, 0x75, 0x0b,
								0x8c, 0x65, 0x4d, 0x48,
							},
							BitLength: 256,
						},
					},
					IEExtensions: nil,
				},
				NewSecurityContextInd: nil,
				NASC:                  nil,
				PDUSessionResourceSetupListHOReq: &ie.PDUSessionResourceSetupListHOReq{
					List: []ie.PDUSessionResourceSetupItemHOReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x01, 0x02, 0x03},
								},
							},
							HandoverRequestTransfer: &aper.OctetString{
								0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08, 0x0f, 0x42, 0x40, 0x20,
								0x0f, 0x42, 0x40, 0x00, 0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f,
								0x3d, 0x26, 0x3f, 0x25, 0x91, 0x00, 0x86, 0x00, 0x01, 0x00, 0x00, 0x88,
								0x00, 0x07, 0x00, 0x01, 0x00, 0x00, 0x09, 0x1d, 0x00,
							},
							IEExtensions: nil,
						},
					},
				},
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x01, 0x02, 0x03},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x11, 0x22, 0x33},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
					},
				},
				TraceActivation: nil,
				MaskedIMEISV: &ie.MaskedIMEISV{
					Value: aper.BitString{
						Bytes:     []uint8{0x11, 0x10, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00},
						BitLength: 64,
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: aper.OctetString{
						0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02,
						0xf8, 0x39, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00,
						0x00, 0x10, 0x00, 0x00, 0x0a,
					},
				},
				MobilityRestrictionList: &ie.MobilityRestrictionList{
					ServingPLMN: &ie.PLMNIdentity{
						Value: aper.OctetString{0x02, 0xf8, 0x39},
					},
					EquivalentPLMNs:          nil,
					RATRestrictions:          nil,
					ForbiddenAreaInformation: nil,
					ServiceAreaInformation:   nil,
					IEExtensions:             nil,
				},
				LocationReportingRequestType:       nil,
				RRCInactiveTransitionReportRequest: nil,
				GUAMI: &ie.GUAMI{
					PLMNIdentity: &ie.PLMNIdentity{
						Value: aper.OctetString{0x02, 0xf8, 0x39},
					},
					AMFRegionID: &ie.AMFRegionID{
						Value: aper.BitString{
							Bytes:     []uint8{0x01},
							BitLength: 8,
						},
					},
					AMFSetID: &ie.AMFSetID{
						Value: aper.BitString{
							Bytes:     []uint8{0x00, 0x40},
							BitLength: 10,
						},
					},
					AMFPointer: &ie.AMFPointer{
						Value: aper.BitString{
							Bytes:     []uint8{0x04},
							BitLength: 6,
						},
					},
					IEExtensions: nil,
				},
				RedirectionVoiceFallback:               nil,
				CNAssistedRANTuning:                    nil,
				SRVCCOperationPossible:                 nil,
				IABAuthorized:                          nil,
				EnhancedCoverageRestriction:            nil,
				UEDifferentiationInfo:                  nil,
				NRV2XServicesAuthorized:                nil,
				LTEV2XServicesAuthorized:               nil,
				NRUESidelinkAggregateMaximumBitrate:    nil,
				LTEUESidelinkAggregateMaximumBitrate:   nil,
				PC5QoSParameters:                       nil,
				CEmodeBrestricted:                      nil,
				UEUPCIoTSupport:                        nil,
				ManagementBasedMDTPLMNList:             nil,
				UERadioCapabilityID:                    nil,
				ExtendedConnectedTime:                  nil,
				TimeSyncAssistanceInfo:                 nil,
				UESliceMaximumBitRateList:              nil,
				FiveGProSeAuthorized:                   nil,
				FiveGProSeUEPC5AggregateMaximumBitRate: nil,
				FiveGProSePC5QoSParameters:             nil,
			},
			expected: []byte{
				0x00, 0x0d, 0x00, 0x80, 0xe2, 0x00, 0x00, 0x0c, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xce, 0x94, 0x60,
				0x9b, 0x0c, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00, 0x0f, 0x40, 0x02, 0x04, 0x00, 0x00, 0x6e, 0x00,
				0x08, 0x08, 0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00, 0x77, 0x00, 0x09, 0x18, 0x00, 0x0c,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x21, 0x10, 0x48, 0xb9, 0xb0, 0xeb, 0xe8,
				0xfc, 0x7f, 0x0c, 0x1f, 0x78, 0x42, 0xb6, 0x6e, 0x11, 0xed, 0x29, 0x65, 0x78, 0x37, 0x39, 0x4a,
				0xc9, 0x2d, 0x0c, 0x76, 0xee, 0x75, 0x0b, 0x8c, 0x65, 0x4d, 0x48, 0x00, 0x49, 0x00, 0x36, 0x00,
				0x00, 0x0a, 0x40, 0x20, 0x01, 0x02, 0x03, 0x2d, 0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08,
				0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00, 0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f,
				0x3d, 0x26, 0x3f, 0x25, 0x91, 0x00, 0x86, 0x00, 0x01, 0x00, 0x00, 0x88, 0x00, 0x07, 0x00, 0x01,
				0x00, 0x00, 0x09, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x22, 0x01, 0x01, 0x02, 0x03, 0x10, 0x08,
				0x11, 0x22, 0x33, 0x00, 0x22, 0x40, 0x08, 0x11, 0x10, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00,
				0x65, 0x00, 0x21, 0x20, 0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02,
				0xf8, 0x39, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00,
				0x10, 0x00, 0x00, 0x0a, 0x00, 0x24, 0x40, 0x04, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x1c, 0x00, 0x07,
				0x00, 0x02, 0xf8, 0x39, 0x01, 0x00, 0x41,
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

func TestHandoverRequestUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// 172.16.3.100
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x64},
			BitLength: 32,
		},
	}
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	sst_case1 := &ie.SST{Value: aper.OctetString([]byte{0x01})}
	sd1_case1 := &ie.SD{Value: aper.OctetString([]byte{0x01, 0x02, 0x03})}
	sd2_case1 := &ie.SD{Value: aper.OctetString([]byte{0x11, 0x22, 0x33})}
	nrEncryptAlgo := &ie.NRencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	nrIntegrityAlgo := &ie.NRintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x40, 0x00},
			BitLength: 16,
		},
	}
	eutraEncryptAlgo := &ie.EUTRAencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	eutraIntegrityAlgo := &ie.EUTRAintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	securityKey := &ie.SecurityKey{
		Value: aper.BitString{
			Bytes: []byte{
				0xe8, 0x23, 0xf9, 0x91, 0xbe, 0xc1, 0x19, 0xb6, 0xe0, 0xd1,
				0x08, 0x1b, 0xe6, 0x34, 0x09, 0xeb, 0x93, 0x16, 0x84, 0x84, 0x8f, 0x9b,
				0x76, 0x4f, 0x80, 0x44, 0x5e, 0x2d, 0x64, 0xb3, 0x71, 0x53,
			},
			BitLength: 256,
		},
	}

	// AMFID: "cafe00"
	amfRegionId := &ie.AMFRegionID{Value: aper.BitString{Bytes: []byte{0xCA}, BitLength: 8}}
	amfSetId := &ie.AMFSetID{Value: aper.BitString{Bytes: []byte{0xFE, 0x00}, BitLength: 10}}
	amfPointer := &ie.AMFPointer{Value: aper.BitString{Bytes: []byte{0x00}, BitLength: 6}}

	// pduSessionResourceSetupRequestTransfer
	gtpteid := &ie.GTPTEID{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	pduSessionResourceSetupRequestTransfer := ie.PDUSessionResourceSetupRequestTransfer{
		ProtocolIEs: &ie.ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
			List: []ie.PDUSessionResourceSetupRequestTransferIEs{
				{
					PDUSessionAggregateMaximumBitRate: &ie.PDUSessionAggregateMaximumBitRate{
						PDUSessionAggregateMaximumBitRateDL: &ie.BitRate{
							Value: 100000000,
						},
						PDUSessionAggregateMaximumBitRateUL: &ie.BitRate{
							Value: 200000000,
						},
					},
				},
				{
					ULNGUUPTNLInformation: &ie.UPTransportLayerInformation{
						Choice: &ie.GTPTunnel{
							TransportLayerAddress: ipv4Addr,
							GTPTEID:               gtpteid,
						},
					},
				},
				{
					PDUSessionType: &ie.PDUSessionType{
						Value: ie.PDUSessionTypePresentIpv4,
					},
				},
				{
					QosFlowSetupRequestList: &ie.QosFlowSetupRequestList{
						List: []ie.QosFlowSetupRequestItem{
							{
								QosFlowIdentifier: &ie.QosFlowIdentifier{
									Value: 9,
								},
								QosFlowLevelQosParameters: &ie.QosFlowLevelQosParameters{
									QosCharacteristics: &ie.QosCharacteristics{
										Choice: &ie.NonDynamic5QIDescriptor{
											FiveQI: &ie.FiveQI{
												Value: 9,
											},
										},
									},
									AllocationAndRetentionPriority: &ie.AllocationAndRetentionPriority{
										PriorityLevelARP: &ie.PriorityLevelARP{
											Value: 15,
										},
										PreEmptionCapability: &ie.PreEmptionCapability{
											Value: ie.PreEmptionCapabilityPresentShallNotTriggerPreEmption,
										},
										PreEmptionVulnerability: &ie.PreEmptionVulnerability{
											Value: ie.PreEmptionVulnerabilityPresentNotPreEmptable,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	pduSessionResourceSetupRequestTransferBytes, err := ie.MarshalBinary(&pduSessionResourceSetupRequestTransfer)
	require.NoError(t, err)
	pduSessionResourceSetupRequestTransferOS := aper.OctetString(
		pduSessionResourceSetupRequestTransferBytes)

	// sourceNGRANNodeToTargetNGRANNodeTransparentContainer
	targetCellNrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x0d, 0x90, 0x30},
			BitLength: 36,
		},
	}
	lastVisitedCellNrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x02, 0xb6, 0x70},
			BitLength: 36,
		},
	}
	sourceNGRANNodeToTargetNGRANNodeTransparentContainer := ie.SourceNGRANNodeToTargetNGRANNodeTransparentContainer{
		RRCContainer: &ie.RRCContainer{
			Value: aper.OctetString([]byte{
				0x10, 0x00, 0x09,
				0x33, 0x31, 0x30, 0x33, 0x31, 0x30, 0x31, 0x34, 0x30, 0x30, 0x30, 0x30,
				0x31, 0x32, 0x30,
			}),
		},
		PDUSessionResourceInformationList: &ie.PDUSessionResourceInformationList{
			List: []ie.PDUSessionResourceInformationItem{
				{
					PDUSessionID: &ie.PDUSessionID{
						Value: 5,
					},
					QosFlowInformationList: &ie.QosFlowInformationList{
						List: []ie.QosFlowInformationItem{
							{
								QosFlowIdentifier: &ie.QosFlowIdentifier{
									Value: 9,
								},
							},
						},
					},
				},
			},
		},
		TargetCellID: &ie.NGRANCGI{
			Choice: &ie.NRCGI{
				PLMNIdentity:   plmnIdentity_case1,
				NRCellIdentity: targetCellNrCellIdentity,
			},
		},
		UEHistoryInformation: &ie.UEHistoryInformation{
			List: []ie.LastVisitedCellItem{
				{
					LastVisitedCellInformation: &ie.LastVisitedCellInformation{
						Choice: &ie.LastVisitedNGRANCellInformation{
							GlobalCellID: &ie.NGRANCGI{
								Choice: &ie.NRCGI{
									PLMNIdentity:   plmnIdentity_case1,
									NRCellIdentity: lastVisitedCellNrCellIdentity,
								},
							},
							CellType: &ie.CellType{
								CellSize: &ie.CellSize{
									Value: ie.CellSizePresentMedium,
								},
							},
							TimeUEStayedInCell: &ie.TimeUEStayedInCell{
								Value: 100,
							},
						},
					},
				},
			},
		},
	}
	sourceNGRANNodeToTargetNGRANNodeTransparentContainerBytes, err := ie.MarshalBinary(
		&sourceNGRANNodeToTargetNGRANNodeTransparentContainer,
	)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverRequest
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x0d, 0x00, 0x80, 0xdd, 0x00, 0x00, 0x0a, 0x00, 0x0a,
				0x00, 0x02, 0x00, 0x02, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00,
				0x0f, 0x40, 0x02, 0x07, 0xc0, 0x00, 0x6e, 0x00, 0x0a, 0x0c,
				0x77, 0x35, 0x94, 0x00, 0x30, 0x3b, 0x9a, 0xca, 0x00, 0x00,
				0x77, 0x00, 0x09, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x5d, 0x00, 0x21, 0x10, 0xe8, 0x23, 0xf9,
				0x91, 0xbe, 0xc1, 0x19, 0xb6, 0xe0, 0xd1, 0x08, 0x1b, 0xe6,
				0x34, 0x09, 0xeb, 0x93, 0x16, 0x84, 0x84, 0x8f, 0x9b, 0x76,
				0x4f, 0x80, 0x44, 0x5e, 0x2d, 0x64, 0xb3, 0x71, 0x53, 0x00,
				0x49, 0x00, 0x38, 0x00, 0x00, 0x05, 0x40, 0x20, 0x01, 0x02,
				0x03, 0x2f, 0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x0a, 0x0c,
				0x05, 0xf5, 0xe1, 0x00, 0x30, 0x0b, 0xeb, 0xc2, 0x00, 0x00,
				0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x03, 0x64, 0x00,
				0x00, 0x00, 0x01, 0x00, 0x86, 0x00, 0x01, 0x00, 0x00, 0x88,
				0x00, 0x07, 0x00, 0x09, 0x00, 0x00, 0x09, 0x38, 0x00, 0x00,
				0x00, 0x00, 0x0a, 0x22, 0x01, 0x01, 0x02, 0x03, 0x10, 0x08,
				0x11, 0x22, 0x33, 0x00, 0x65, 0x00, 0x30, 0x2f, 0x40, 0x12,
				0x10, 0x00, 0x09, 0x33, 0x31, 0x30, 0x33, 0x31, 0x30, 0x31,
				0x34, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x30, 0x00, 0x00,
				0x05, 0x00, 0x09, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x0d,
				0x90, 0x30, 0x00, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x02,
				0xb6, 0x71, 0x00, 0x00, 0x64, 0x00, 0x1c, 0x00, 0x07, 0x00,
				0x13, 0x30, 0x01, 0xca, 0xfe, 0x00,
			},
			expected: &HandoverRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				HandoverType: &ie.HandoverType{
					Value: ie.HandoverTypePresentIntra5gs,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentNgIntraSystemHandoverTriggered,
					},
				},
				UEAggregateMaximumBitRate: &ie.UEAggregateMaximumBitRate{
					UEAggregateMaximumBitRateDL: &ie.BitRate{
						Value: 2000000000,
					},
					UEAggregateMaximumBitRateUL: &ie.BitRate{
						Value: 1000000000,
					},
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms:             nrEncryptAlgo,
					NRintegrityProtectionAlgorithms:    nrIntegrityAlgo,
					EUTRAencryptionAlgorithms:          eutraEncryptAlgo,
					EUTRAintegrityProtectionAlgorithms: eutraIntegrityAlgo,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: securityKey,
				},
				PDUSessionResourceSetupListHOReq: &ie.PDUSessionResourceSetupListHOReq{
					List: []ie.PDUSessionResourceSetupItemHOReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd1_case1,
							},
							HandoverRequestTransfer: &pduSessionResourceSetupRequestTransferOS,
						},
					},
				},
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd1_case1,
							},
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd2_case1,
							},
						},
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: sourceNGRANNodeToTargetNGRANNodeTransparentContainerBytes,
				},
				GUAMI: &ie.GUAMI{
					PLMNIdentity: plmnIdentity_case1,
					AMFRegionID:  amfRegionId,
					AMFSetID:     amfSetId,
					AMFPointer:   amfPointer,
				},
			},
		},
		{
			name: "Case 2: from ueranemu pipeline basic-k8s TestN2Handover",
			input: []byte{
				0x00, 0x0d, 0x00, 0x80, 0xe2, 0x00, 0x00, 0x0c, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xce, 0x94, 0x60,
				0x9b, 0x0c, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00, 0x0f, 0x40, 0x02, 0x04, 0x00, 0x00, 0x6e, 0x00,
				0x08, 0x08, 0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00, 0x77, 0x00, 0x09, 0x18, 0x00, 0x0c,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x21, 0x10, 0x48, 0xb9, 0xb0, 0xeb, 0xe8,
				0xfc, 0x7f, 0x0c, 0x1f, 0x78, 0x42, 0xb6, 0x6e, 0x11, 0xed, 0x29, 0x65, 0x78, 0x37, 0x39, 0x4a,
				0xc9, 0x2d, 0x0c, 0x76, 0xee, 0x75, 0x0b, 0x8c, 0x65, 0x4d, 0x48, 0x00, 0x49, 0x00, 0x36, 0x00,
				0x00, 0x0a, 0x40, 0x20, 0x01, 0x02, 0x03, 0x2d, 0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08,
				0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00, 0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f,
				0x3d, 0x26, 0x3f, 0x25, 0x91, 0x00, 0x86, 0x00, 0x01, 0x00, 0x00, 0x88, 0x00, 0x07, 0x00, 0x01,
				0x00, 0x00, 0x09, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x22, 0x01, 0x01, 0x02, 0x03, 0x10, 0x08,
				0x11, 0x22, 0x33, 0x00, 0x22, 0x40, 0x08, 0x11, 0x10, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00,
				0x65, 0x00, 0x21, 0x20, 0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02,
				0xf8, 0x39, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00,
				0x10, 0x00, 0x00, 0x0a, 0x00, 0x24, 0x40, 0x04, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x1c, 0x00, 0x07,
				0x00, 0x02, 0xf8, 0x39, 0x01, 0x00, 0x41,
			},
			expected: &HandoverRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 887252622092,
				},
				HandoverType: &ie.HandoverType{
					Value: aper.Enumerated(0),
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHandoverDesirableForRadioReason,
					},
				},
				UEAggregateMaximumBitRate: &ie.UEAggregateMaximumBitRate{
					UEAggregateMaximumBitRateDL: &ie.BitRate{
						Value: 1000000,
					},
					UEAggregateMaximumBitRateUL: &ie.BitRate{
						Value: 1000000,
					},
					IEExtensions: nil,
				},
				CoreNetworkAssistanceInformationForInactive: nil,
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms: &ie.NRencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0xc0, 0x00},
							BitLength: 16,
						},
					},
					NRintegrityProtectionAlgorithms: &ie.NRintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0xc0, 0x00},
							BitLength: 16,
						},
					},
					EUTRAencryptionAlgorithms: &ie.EUTRAencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0x00, 0x00},
							BitLength: 16,
						},
					},
					EUTRAintegrityProtectionAlgorithms: &ie.EUTRAintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []uint8{0x00, 0x00},
							BitLength: 16,
						},
					},
					IEExtensions: nil,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: &ie.SecurityKey{
						Value: aper.BitString{
							Bytes: []uint8{
								0x48, 0xb9, 0xb0, 0xeb, 0xe8, 0xfc, 0x7f, 0x0c, 0x1f, 0x78, 0x42, 0xb6, 0x6e, 0x11,
								0xed, 0x29, 0x65, 0x78, 0x37, 0x39, 0x4a, 0xc9, 0x2d, 0x0c, 0x76, 0xee, 0x75, 0x0b,
								0x8c, 0x65, 0x4d, 0x48,
							},
							BitLength: 256,
						},
					},
					IEExtensions: nil,
				},
				NewSecurityContextInd: nil,
				NASC:                  nil,
				PDUSessionResourceSetupListHOReq: &ie.PDUSessionResourceSetupListHOReq{
					List: []ie.PDUSessionResourceSetupItemHOReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x01, 0x02, 0x03},
								},
							},
							HandoverRequestTransfer: &aper.OctetString{
								0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08, 0x0f, 0x42, 0x40, 0x20,
								0x0f, 0x42, 0x40, 0x00, 0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f,
								0x3d, 0x26, 0x3f, 0x25, 0x91, 0x00, 0x86, 0x00, 0x01, 0x00, 0x00, 0x88,
								0x00, 0x07, 0x00, 0x01, 0x00, 0x00, 0x09, 0x1d, 0x00,
							},
							IEExtensions: nil,
						},
					},
				},
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x01, 0x02, 0x03},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x11, 0x22, 0x33},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
					},
				},
				TraceActivation: nil,
				MaskedIMEISV: &ie.MaskedIMEISV{
					Value: aper.BitString{
						Bytes:     []uint8{0x11, 0x10, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00},
						BitLength: 64,
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: aper.OctetString{
						0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02,
						0xf8, 0x39, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00,
						0x00, 0x10, 0x00, 0x00, 0x0a,
					},
				},
				MobilityRestrictionList: &ie.MobilityRestrictionList{
					ServingPLMN: &ie.PLMNIdentity{
						Value: aper.OctetString{0x02, 0xf8, 0x39},
					},
					EquivalentPLMNs:          nil,
					RATRestrictions:          nil,
					ForbiddenAreaInformation: nil,
					ServiceAreaInformation:   nil,
					IEExtensions:             nil,
				},
				LocationReportingRequestType:       nil,
				RRCInactiveTransitionReportRequest: nil,
				GUAMI: &ie.GUAMI{
					PLMNIdentity: &ie.PLMNIdentity{
						Value: aper.OctetString{0x02, 0xf8, 0x39},
					},
					AMFRegionID: &ie.AMFRegionID{
						Value: aper.BitString{
							Bytes:     []uint8{0x01},
							BitLength: 8,
						},
					},
					AMFSetID: &ie.AMFSetID{
						Value: aper.BitString{
							Bytes:     []uint8{0x00, 0x40},
							BitLength: 10,
						},
					},
					AMFPointer: &ie.AMFPointer{
						Value: aper.BitString{
							Bytes:     []uint8{0x04},
							BitLength: 6,
						},
					},
					IEExtensions: nil,
				},
				RedirectionVoiceFallback:               nil,
				CNAssistedRANTuning:                    nil,
				SRVCCOperationPossible:                 nil,
				IABAuthorized:                          nil,
				EnhancedCoverageRestriction:            nil,
				UEDifferentiationInfo:                  nil,
				NRV2XServicesAuthorized:                nil,
				LTEV2XServicesAuthorized:               nil,
				NRUESidelinkAggregateMaximumBitrate:    nil,
				LTEUESidelinkAggregateMaximumBitrate:   nil,
				PC5QoSParameters:                       nil,
				CEmodeBrestricted:                      nil,
				UEUPCIoTSupport:                        nil,
				ManagementBasedMDTPLMNList:             nil,
				UERadioCapabilityID:                    nil,
				ExtendedConnectedTime:                  nil,
				TimeSyncAssistanceInfo:                 nil,
				UESliceMaximumBitRateList:              nil,
				FiveGProSeAuthorized:                   nil,
				FiveGProSeUEPC5AggregateMaximumBitRate: nil,
				FiveGProSePC5QoSParameters:             nil,
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
