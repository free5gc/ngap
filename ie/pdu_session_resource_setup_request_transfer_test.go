package ie

import (
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/stretchr/testify/require"
)

func TestPDUSessionResourceSetupRequestTransferMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *PDUSessionResourceSetupRequestTransfer
		expected []byte
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover",
			input: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{
									Value: 1000000, // Downlink bitrate
								},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{
									Value: 1000000, // Uplink bitrate
								},
							},
						},
						{
							PDUSessionAggregateMaximumBitRate: nil,
							ULNGUUPTNLInformation: &UPTransportLayerInformation{
								Choice: &GTPTunnel{
									TransportLayerAddress: &TransportLayerAddress{
										Value: aper.BitString{
											Bytes:     []uint8{0xac, 0x10, 0x1f, 0x3d},
											BitLength: 32,
										},
									},
									GTPTEID: &GTPTEID{
										Value: aper.OctetString{0x26, 0x3f, 0x25, 0x91},
									},
								},
							},
						},
						{
							PDUSessionType: &PDUSessionType{
								Value: aper.Enumerated(0),
							},
						},
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{
											Value: 1,
										},
										QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
											QosCharacteristics: &QosCharacteristics{
												Choice: &NonDynamic5QIDescriptor{
													FiveQI: &FiveQI{
														Value: 9,
													},
												},
											},
											AllocationAndRetentionPriority: &AllocationAndRetentionPriority{
												PriorityLevelARP: &PriorityLevelARP{
													Value: 8,
												},
												PreEmptionCapability: &PreEmptionCapability{
													Value: aper.Enumerated(1),
												},
												PreEmptionVulnerability: &PreEmptionVulnerability{
													Value: aper.Enumerated(0),
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
				0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08, 0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00,
				0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f, 0x3d, 0x26, 0x3f, 0x25, 0x91, 0x00, 0x86, 0x00,
				0x01, 0x00, 0x00, 0x88, 0x00, 0x07, 0x00, 0x01, 0x00, 0x00, 0x09, 0x1d, 0x00,
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestPDUSessionEstablishment",
			input: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{
									Value: 1000000,
								},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{
									Value: 1000000,
								},
							},
						},
						{
							PDUSessionAggregateMaximumBitRate: nil,
							ULNGUUPTNLInformation: &UPTransportLayerInformation{
								Choice: &GTPTunnel{
									TransportLayerAddress: &TransportLayerAddress{
										Value: aper.BitString{
											Bytes:     []byte{0xac, 0x10, 0x1f, 0x3d},
											BitLength: 32,
										},
									},
									GTPTEID: &GTPTEID{
										Value: aper.OctetString{0xc4, 0xb2, 0x20, 0x4e},
									},
								},
							},
						},
						{
							PDUSessionType: &PDUSessionType{
								Value: aper.Enumerated(0),
							},
						},
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{
											Value: 1,
										},
										QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
											QosCharacteristics: &QosCharacteristics{
												Choice: &NonDynamic5QIDescriptor{
													FiveQI: &FiveQI{
														Value: 9,
													},
												},
											},
											AllocationAndRetentionPriority: &AllocationAndRetentionPriority{
												PriorityLevelARP: &PriorityLevelARP{
													Value: 8,
												},
												PreEmptionCapability: &PreEmptionCapability{
													Value: aper.Enumerated(1),
												},
												PreEmptionVulnerability: &PreEmptionVulnerability{
													Value: aper.Enumerated(0),
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
				0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08,
				0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00,
				0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f,
				0x3d, 0xc4, 0xb2, 0x20, 0x4e, 0x00, 0x86, 0x00,
				0x01, 0x00, 0x00, 0x88, 0x00, 0x07, 0x00, 0x01,
				0x00, 0x00, 0x09, 0x1d, 0x00,
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

func TestPDUSessionResourceSetupRequestTransferUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *PDUSessionResourceSetupRequestTransfer
	}{
		{
			name: "Case 1: from ueranemu basic-k8s pipeline TestN2Handover",
			input: []byte{
				0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08, 0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00,
				0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f, 0x3d, 0x26, 0x3f, 0x25, 0x91, 0x00, 0x86, 0x00,
				0x01, 0x00, 0x00, 0x88, 0x00, 0x07, 0x00, 0x01, 0x00, 0x00, 0x09, 0x1d, 0x00,
			},
			expected: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{
									Value: 1000000, // Downlink bitrate
								},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{
									Value: 1000000, // Uplink bitrate
								},
							},
						},
						{
							ULNGUUPTNLInformation: &UPTransportLayerInformation{
								Choice: &GTPTunnel{
									TransportLayerAddress: &TransportLayerAddress{
										Value: aper.BitString{
											Bytes:     []uint8{0xac, 0x10, 0x1f, 0x3d},
											BitLength: 32,
										},
									},
									GTPTEID: &GTPTEID{
										Value: aper.OctetString{0x26, 0x3f, 0x25, 0x91},
									},
								},
							},
						},
						{
							PDUSessionType: &PDUSessionType{
								Value: aper.Enumerated(0),
							},
						},
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{
											Value: 1,
										},
										QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
											QosCharacteristics: &QosCharacteristics{
												Choice: &NonDynamic5QIDescriptor{
													FiveQI: &FiveQI{
														Value: 9,
													},
												},
											},
											AllocationAndRetentionPriority: &AllocationAndRetentionPriority{
												PriorityLevelARP: &PriorityLevelARP{
													Value: 8,
												},
												PreEmptionCapability: &PreEmptionCapability{
													Value: aper.Enumerated(1),
												},
												PreEmptionVulnerability: &PreEmptionVulnerability{
													Value: aper.Enumerated(0),
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
			name: "Case 2: from ueranemu k8s-basic pipeline TestPDUSessionEstablishment",
			input: []byte{
				0x00, 0x00, 0x04, 0x00, 0x82, 0x00, 0x08, 0x08,
				0x0f, 0x42, 0x40, 0x20, 0x0f, 0x42, 0x40, 0x00,
				0x8b, 0x00, 0x0a, 0x01, 0xf0, 0xac, 0x10, 0x1f,
				0x3d, 0xc4, 0xb2, 0x20, 0x4e, 0x00, 0x86, 0x00,
				0x01, 0x00, 0x00, 0x88, 0x00, 0x07, 0x00, 0x01,
				0x00, 0x00, 0x09, 0x1d, 0x00,
			},
			expected: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{
									Value: 1000000,
								},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{
									Value: 1000000,
								},
							},
						},
						{
							PDUSessionAggregateMaximumBitRate: nil,
							ULNGUUPTNLInformation: &UPTransportLayerInformation{
								Choice: &GTPTunnel{
									TransportLayerAddress: &TransportLayerAddress{
										Value: aper.BitString{
											Bytes:     []byte{0xac, 0x10, 0x1f, 0x3d},
											BitLength: 32,
										},
									},
									GTPTEID: &GTPTEID{
										Value: aper.OctetString{0xc4, 0xb2, 0x20, 0x4e},
									},
								},
							},
						},
						{
							PDUSessionType: &PDUSessionType{
								Value: aper.Enumerated(0),
							},
						},
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
									{
										QosFlowIdentifier: &QosFlowIdentifier{
											Value: 1,
										},
										QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
											QosCharacteristics: &QosCharacteristics{
												Choice: &NonDynamic5QIDescriptor{
													FiveQI: &FiveQI{
														Value: 9,
													},
												},
											},
											AllocationAndRetentionPriority: &AllocationAndRetentionPriority{
												PriorityLevelARP: &PriorityLevelARP{
													Value: 8,
												},
												PreEmptionCapability: &PreEmptionCapability{
													Value: aper.Enumerated(1),
												},
												PreEmptionVulnerability: &PreEmptionVulnerability{
													Value: aper.Enumerated(0),
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
			ie := new(PDUSessionResourceSetupRequestTransfer)
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
