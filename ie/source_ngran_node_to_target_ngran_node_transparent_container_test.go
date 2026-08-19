package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestSourceNGRANNodeToTargetNGRANNodeTransparentContainerMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *SourceNGRANNodeToTargetNGRANNodeTransparentContainer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover HandoverRequired",
			input: &SourceNGRANNodeToTargetNGRANNodeTransparentContainer{
				RRCContainer: &RRCContainer{
					Value: []byte{0x00, 0x00, 0x11},
				},
				PDUSessionResourceInformationList: &PDUSessionResourceInformationList{
					List: []PDUSessionResourceInformationItem{
						{
							PDUSessionID: &PDUSessionID{
								Value: 1,
							},
							QosFlowInformationList: &QosFlowInformationList{
								List: []QosFlowInformationItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{Value: 1},
										DLForwarding:      nil,
										IEExtensions:      nil,
									},
								},
							},
							DRBsToQosFlowsMappingList: nil,
							IEExtensions:              nil,
						},
					},
				},
				ERABInformationList: nil,
				TargetCellID: &NGRANCGI{
					Choice: &NRCGI{
						PLMNIdentity: &PLMNIdentity{
							Value: []byte{0x02, 0xf8, 0x39},
						},
						NRCellIdentity: &NRCellIdentity{
							Value: aper.BitString{
								Bytes:     []uint8{0x00, 0x01, 0x02, 0x00, 0x10},
								BitLength: 36,
							},
						},
						IEExtensions: nil,
					},
				},
				IndexToRFSP: nil,
				UEHistoryInformation: &UEHistoryInformation{
					List: []LastVisitedCellItem{
						{
							LastVisitedCellInformation: &LastVisitedCellInformation{
								Choice: &LastVisitedNGRANCellInformation{
									GlobalCellID: &NGRANCGI{
										Choice: &NRCGI{
											PLMNIdentity: &PLMNIdentity{
												Value: []byte{0x02, 0xf8, 0x39},
											},
											NRCellIdentity: &NRCellIdentity{
												Value: aper.BitString{
													Bytes:     []uint8{0x00, 0x00, 0x00, 0x00, 0x10},
													BitLength: 36,
												},
											},
											IEExtensions: nil,
										},
									},
									CellType: &CellType{
										CellSize: &CellSize{
											Value: 0,
										},
										IEExtensions: nil,
									},
									TimeUEStayedInCell: &TimeUEStayedInCell{
										Value: 10,
									},
									TimeUEStayedInCellEnhancedGranularity: nil,
									HOCauseValue:                          nil,
									IEExtensions:                          nil,
								},
							},
							IEExtensions: nil,
						},
					},
				},
				IEExtensions: nil,
			},
			expected: []byte{
				0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x01,
				0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x0a,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := MarshalBinary(tc.input)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, b)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestSourceNGRANNodeToTargetNGRANNodeTransparentContainerUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *SourceNGRANNodeToTargetNGRANNodeTransparentContainer
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover HandoverRequired",
			input: []byte{
				0x40, 0x03, 0x00, 0x00, 0x11, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x01,
				0x02, 0x00, 0x10, 0x00, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x0a,
			},
			expected: &SourceNGRANNodeToTargetNGRANNodeTransparentContainer{
				RRCContainer: &RRCContainer{
					Value: []byte{0x00, 0x00, 0x11},
				},
				PDUSessionResourceInformationList: &PDUSessionResourceInformationList{
					List: []PDUSessionResourceInformationItem{
						{
							PDUSessionID: &PDUSessionID{
								Value: 1,
							},
							QosFlowInformationList: &QosFlowInformationList{
								List: []QosFlowInformationItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{Value: 1},
										DLForwarding:      nil,
										IEExtensions:      nil,
									},
								},
							},
							DRBsToQosFlowsMappingList: nil,
							IEExtensions:              nil,
						},
					},
				},
				ERABInformationList: nil,
				TargetCellID: &NGRANCGI{
					Choice: &NRCGI{
						PLMNIdentity: &PLMNIdentity{
							Value: []byte{0x02, 0xf8, 0x39},
						},
						NRCellIdentity: &NRCellIdentity{
							Value: aper.BitString{
								Bytes:     []uint8{0x00, 0x01, 0x02, 0x00, 0x10},
								BitLength: 36,
							},
						},
						IEExtensions: nil,
					},
				},
				IndexToRFSP: nil,
				UEHistoryInformation: &UEHistoryInformation{
					List: []LastVisitedCellItem{
						{
							LastVisitedCellInformation: &LastVisitedCellInformation{
								Choice: &LastVisitedNGRANCellInformation{
									GlobalCellID: &NGRANCGI{
										Choice: &NRCGI{
											PLMNIdentity: &PLMNIdentity{
												Value: []byte{0x02, 0xf8, 0x39},
											},
											NRCellIdentity: &NRCellIdentity{
												Value: aper.BitString{
													Bytes:     []uint8{0x00, 0x00, 0x00, 0x00, 0x10},
													BitLength: 36,
												},
											},
											IEExtensions: nil,
										},
									},
									CellType: &CellType{
										CellSize: &CellSize{
											Value: 0,
										},
										IEExtensions: nil,
									},
									TimeUEStayedInCell: &TimeUEStayedInCell{
										Value: 10,
									},
									TimeUEStayedInCellEnhancedGranularity: nil,
									HOCauseValue:                          nil,
									IEExtensions:                          nil,
								},
							},
							IEExtensions: nil,
						},
					},
				},
				IEExtensions: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(SourceNGRANNodeToTargetNGRANNodeTransparentContainer)
			err := UnmarshalBinary(tc.input, ie)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, ie)
			} else {
				require.Error(t, err)
			}
		})
	}
}
