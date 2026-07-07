package ie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceModifyRequestTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceModifyRequestTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionModification",
			input: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowAddOrModifyRequestList: &QosFlowAddOrModifyRequestList{
								List: []QosFlowAddOrModifyRequestItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{
											Value: 2,
										},
										QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
											QosCharacteristics: &QosCharacteristics{
												Choice: &NonDynamic5QIDescriptor{
													FiveQI: &FiveQI{
														Value: 5,
													},
												},
											},
											AllocationAndRetentionPriority: &AllocationAndRetentionPriority{
												PriorityLevelARP: &PriorityLevelARP{
													Value: 15,
												},
												PreEmptionCapability: &PreEmptionCapability{
													Value: 0, // Enumerated value 0
												},
												PreEmptionVulnerability: &PreEmptionVulnerability{
													Value: 1, // Enumerated value 1
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
			expected: []byte{
				0x00, 0x00, 0x01, 0x00, 0x87, 0x00, 0x07, 0x01, 0x01, 0x00, 0x00, 0x05, 0x38, 0x40,
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

func TestPDUSessionResourceModifyRequestTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceModifyRequestTransfer
	}{
		{
			name: "Case 1: from ueranemu k8s-basic pipeline TestPDUSessionModification",
			input: []byte{
				0x00, 0x00, 0x01, 0x00, 0x87, 0x00, 0x07, 0x01, 0x01, 0x00, 0x00, 0x05, 0x38, 0x40,
			},
			expected: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowAddOrModifyRequestList: &QosFlowAddOrModifyRequestList{
								List: []QosFlowAddOrModifyRequestItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{
											Value: 2,
										},
										QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
											QosCharacteristics: &QosCharacteristics{
												Choice: &NonDynamic5QIDescriptor{
													FiveQI: &FiveQI{
														Value: 5,
													},
												},
											},
											AllocationAndRetentionPriority: &AllocationAndRetentionPriority{
												PriorityLevelARP: &PriorityLevelARP{
													Value: 15,
												},
												PreEmptionCapability: &PreEmptionCapability{
													Value: 0, // Enumerated value 0
												},
												PreEmptionVulnerability: &PreEmptionVulnerability{
													Value: 1, // Enumerated value 1
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ie := new(PDUSessionResourceModifyRequestTransfer)
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
