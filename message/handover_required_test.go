package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestHandoverRequiredMarshalBinary(t *testing.T) {
	t.Parallel()

	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}

	// HandoverRequiredTransfer
	handoverRequiredTransfer := ie.HandoverRequiredTransfer{
		DirectForwardingPathAvailability: &ie.DirectForwardingPathAvailability{
			Value: ie.DirectForwardingPathAvailabilityPresentDirectPathAvailable,
		},
	}
	handoverRequiredTransferBytes, err := ie.MarshalBinary(&handoverRequiredTransfer)
	require.NoError(t, err)
	handoverRequiredTransferOS := aper.OctetString(handoverRequiredTransferBytes)

	// SourceToTargetTransparentContainer (Contains sourceNGRANNodeToTargetNGRANNodeTransparentContainer)
	targetCellNrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x0d, 0x90, 0x30},
			BitLength: 36,
		},
	}
	targetCellGnbId := &ie.GNBID{
		Choice: &ie.GNBIDForGNBID{
			Value: aper.BitString{
				Bytes:     []byte{0x00, 0x00, 0x0d, 0x90},
				BitLength: 32,
			},
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
	sourceNGRANNodeToTargetNGRANNodeTransparentContainerOS := aper.OctetString(
		sourceNGRANNodeToTargetNGRANNodeTransparentContainerBytes)

	testCases := []struct {
		name     string
		input    *HandoverRequired
		expected []byte
	}{
		{
			name: "Case 1",
			input: &HandoverRequired{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3405774848,
				},
				HandoverType: &ie.HandoverType{
					Value: ie.HandoverTypePresentIntra5gs,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentNgIntraSystemHandoverTriggered,
					},
				},
				TargetID: &ie.TargetID{
					Choice: &ie.TargetRANNodeID{
						GlobalRANNodeID: &ie.GlobalRANNodeID{
							Choice: &ie.GlobalGNBID{
								PLMNIdentity: plmnIdentity_case1,
								GNBID:        targetCellGnbId,
							},
						},
						SelectedTAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
				DirectForwardingPathAvailability: &ie.DirectForwardingPathAvailability{
					Value: ie.DirectForwardingPathAvailabilityPresentDirectPathAvailable,
				},
				PDUSessionResourceListHORqd: &ie.PDUSessionResourceListHORqd{
					List: []ie.PDUSessionResourceItemHORqd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							HandoverRequiredTransfer: &handoverRequiredTransferOS,
						},
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: sourceNGRANNodeToTargetNGRANNodeTransparentContainerOS,
				},
			},
			expected: []byte{
				0x00, 0x0c, 0x00, 0x73, 0x00, 0x00, 0x08, 0x00,
				0x0a, 0x00, 0x02, 0x00, 0x01, 0x00, 0x55, 0x00, 0x05, 0xc0, 0xcb,
				0x00, 0x00, 0x00, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00, 0x0f, 0x40,
				0x02, 0x07, 0xc0, 0x00, 0x69, 0x00, 0x10, 0x00, 0x13, 0x30, 0x01,
				0x50, 0x00, 0x00, 0x0d, 0x90, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00,
				0x01, 0x00, 0x16, 0x40, 0x01, 0x00, 0x00, 0x3d, 0x00, 0x05, 0x00,
				0x00, 0x05, 0x01, 0x40, 0x00, 0x65, 0x00, 0x30, 0x2f, 0x40, 0x12,
				0x10, 0x00, 0x09, 0x33, 0x31, 0x30, 0x33, 0x31, 0x30, 0x31, 0x34,
				0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x30, 0x00, 0x00, 0x05, 0x00,
				0x09, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x0d, 0x90, 0x30, 0x00,
				0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x02, 0xb6, 0x71, 0x00, 0x00,
				0x64,
			},
		},
		{
			name: "Case 2: from ueranemu basic-k8s pipeline TestN2Handover",
			input: &HandoverRequired{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 910057611806,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				HandoverType: &ie.HandoverType{
					Value: 0, // aper.Enumerated value
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHandoverDesirableForRadioReason,
					},
				},
				TargetID: &ie.TargetID{
					Choice: &ie.TargetRANNodeID{
						GlobalRANNodeID: &ie.GlobalRANNodeID{
							Choice: &ie.GlobalGNBID{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: []byte{0x02, 0xf8, 0x39},
								},
								GNBID: &ie.GNBID{
									Choice: &ie.GNBIDForGNBID{
										Value: aper.BitString{
											Bytes:     []uint8{0x00, 0x01, 0x02},
											BitLength: 24,
										},
									},
								},
								IEExtensions: nil,
							},
						},
						SelectedTAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: []byte{0x30, 0x33, 0x99},
							},
							IEExtensions: nil,
						},
						IEExtensions: nil,
					},
				},
				DirectForwardingPathAvailability: nil,
				PDUSessionResourceListHORqd: &ie.PDUSessionResourceListHORqd{
					List: []ie.PDUSessionResourceItemHORqd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							HandoverRequiredTransfer: &aper.OctetString{0x40},
							IEExtensions:             nil,
						},
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: []byte{
						0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01,
						0x00, 0x01, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x01,
						0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39,
						0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x0a,
					},
				},
			},
			expected: []byte{
				0x00, 0x0c, 0x00, 0x5f, 0x00, 0x00, 0x07, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xd3, 0xe3, 0xa9, 0x22,
				0x1e, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00, 0x0f, 0x40, 0x02,
				0x04, 0x00, 0x00, 0x69, 0x00, 0x0f, 0x00, 0x02, 0xf8, 0x39, 0x10, 0x00, 0x01, 0x02, 0x00, 0x02,
				0xf8, 0x39, 0x30, 0x33, 0x99, 0x00, 0x3d, 0x00, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x40, 0x00, 0x65,
				0x00, 0x21, 0x20, 0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02, 0xf8,
				0x39, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00, 0x10,
				0x00, 0x00, 0x0a,
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

func TestHandoverRequiredUnmarshalBinary(t *testing.T) {
	t.Parallel()

	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}

	// HandoverRequiredTransfer
	handoverRequiredTransfer := ie.HandoverRequiredTransfer{
		DirectForwardingPathAvailability: &ie.DirectForwardingPathAvailability{
			Value: ie.DirectForwardingPathAvailabilityPresentDirectPathAvailable,
		},
	}
	handoverRequiredTransferBytes, err := ie.MarshalBinary(&handoverRequiredTransfer)
	require.NoError(t, err)
	handoverRequiredTransferOS := aper.OctetString(handoverRequiredTransferBytes)

	// SourceToTargetTransparentContainer (Contains sourceNGRANNodeToTargetNGRANNodeTransparentContainer)
	targetCellNrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00, 0x0d, 0x90, 0x30},
			BitLength: 36,
		},
	}
	targetCellGnbId := &ie.GNBID{
		Choice: &ie.GNBIDForGNBID{
			Value: aper.BitString{
				Bytes:     []byte{0x00, 0x00, 0x0d, 0x90},
				BitLength: 32,
			},
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
	sourceNGRANNodeToTargetNGRANNodeTransparentContainerOS := aper.OctetString(
		sourceNGRANNodeToTargetNGRANNodeTransparentContainerBytes)

	testCases := []struct {
		name     string
		input    []byte
		expected *HandoverRequired
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x0c, 0x00, 0x73, 0x00, 0x00, 0x08, 0x00,
				0x0a, 0x00, 0x02, 0x00, 0x01, 0x00, 0x55, 0x00, 0x05, 0xc0, 0xcb,
				0x00, 0x00, 0x00, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00, 0x0f, 0x40,
				0x02, 0x07, 0xc0, 0x00, 0x69, 0x00, 0x10, 0x00, 0x13, 0x30, 0x01,
				0x50, 0x00, 0x00, 0x0d, 0x90, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00,
				0x01, 0x00, 0x16, 0x40, 0x01, 0x00, 0x00, 0x3d, 0x00, 0x05, 0x00,
				0x00, 0x05, 0x01, 0x40, 0x00, 0x65, 0x00, 0x30, 0x2f, 0x40, 0x12,
				0x10, 0x00, 0x09, 0x33, 0x31, 0x30, 0x33, 0x31, 0x30, 0x31, 0x34,
				0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x30, 0x00, 0x00, 0x05, 0x00,
				0x09, 0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x0d, 0x90, 0x30, 0x00,
				0x00, 0x13, 0x30, 0x01, 0x00, 0x00, 0x02, 0xb6, 0x71, 0x00, 0x00,
				0x64,
			},
			expected: &HandoverRequired{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 3405774848,
				},
				HandoverType: &ie.HandoverType{
					Value: ie.HandoverTypePresentIntra5gs,
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentNgIntraSystemHandoverTriggered,
					},
				},
				TargetID: &ie.TargetID{
					Choice: &ie.TargetRANNodeID{
						GlobalRANNodeID: &ie.GlobalRANNodeID{
							Choice: &ie.GlobalGNBID{
								PLMNIdentity: plmnIdentity_case1,
								GNBID:        targetCellGnbId,
							},
						},
						SelectedTAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
				DirectForwardingPathAvailability: &ie.DirectForwardingPathAvailability{
					Value: ie.DirectForwardingPathAvailabilityPresentDirectPathAvailable,
				},
				PDUSessionResourceListHORqd: &ie.PDUSessionResourceListHORqd{
					List: []ie.PDUSessionResourceItemHORqd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							HandoverRequiredTransfer: &handoverRequiredTransferOS,
						},
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: sourceNGRANNodeToTargetNGRANNodeTransparentContainerOS,
				},
			},
		},
		{
			name: "Case 2: from ueranemu basic-k8s pipeline TestN2Handover",
			input: []byte{
				0x00, 0x0c, 0x00, 0x5f, 0x00, 0x00, 0x07, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xd3, 0xe3, 0xa9, 0x22,
				0x1e, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x1d, 0x00, 0x01, 0x00, 0x00, 0x0f, 0x40, 0x02,
				0x04, 0x00, 0x00, 0x69, 0x00, 0x0f, 0x00, 0x02, 0xf8, 0x39, 0x10, 0x00, 0x01, 0x02, 0x00, 0x02,
				0xf8, 0x39, 0x30, 0x33, 0x99, 0x00, 0x3d, 0x00, 0x05, 0x00, 0x00, 0x0a, 0x01, 0x40, 0x00, 0x65,
				0x00, 0x21, 0x20, 0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02, 0xf8,
				0x39, 0x00, 0x01, 0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00, 0x10,
				0x00, 0x00, 0x0a,
			},
			expected: &HandoverRequired{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 910057611806,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				HandoverType: &ie.HandoverType{
					Value: 0, // aper.Enumerated value
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentHandoverDesirableForRadioReason,
					},
				},
				TargetID: &ie.TargetID{
					Choice: &ie.TargetRANNodeID{
						GlobalRANNodeID: &ie.GlobalRANNodeID{
							Choice: &ie.GlobalGNBID{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: []byte{0x02, 0xf8, 0x39},
								},
								GNBID: &ie.GNBID{
									Choice: &ie.GNBIDForGNBID{
										Value: aper.BitString{
											Bytes:     []uint8{0x00, 0x01, 0x02},
											BitLength: 24,
										},
									},
								},
								IEExtensions: nil,
							},
						},
						SelectedTAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: []byte{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: []byte{0x30, 0x33, 0x99},
							},
							IEExtensions: nil,
						},
						IEExtensions: nil,
					},
				},
				DirectForwardingPathAvailability: nil,
				PDUSessionResourceListHORqd: &ie.PDUSessionResourceListHORqd{
					List: []ie.PDUSessionResourceItemHORqd{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							HandoverRequiredTransfer: &aper.OctetString{0x40},
							IEExtensions:             nil,
						},
					},
				},
				SourceToTargetTransparentContainer: &ie.SourceToTargetTransparentContainer{
					Value: []byte{
						0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01,
						0x00, 0x01, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x01,
						0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39,
						0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x0a,
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
