package ie

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
)

func TestCauseToString(t *testing.T) {
	testCases := []struct {
		name     string
		cause    *Cause
		causeStr string
	}{
		{
			name:     "Nil Cause",
			cause:    nil,
			causeStr: "",
		},
		{
			name:     "Nil Cause Group",
			cause:    &Cause{Choice: nil},
			causeStr: "Invalid: Nil cause group",
		},
		{
			name:     "Radio Network Group",
			cause:    &Cause{Choice: &CauseRadioNetwork{Value: CauseRadioNetworkPresentNotSupported5QIValue}},
			causeStr: "Radio Network Layer: Not supported 5QI value",
		},
		{
			name:     "Transport Layer Group",
			cause:    &Cause{Choice: &CauseTransport{Value: CauseTransportPresentTransportResourceUnavailable}},
			causeStr: "Transport Layer: Transport resource unavailable",
		},
		{
			name:     "NAS Group",
			cause:    &Cause{Choice: &CauseNas{Value: CauseNasPresentNormalRelease}},
			causeStr: "NAS: Normal release",
		},
		{
			name:     "Protocol Group",
			cause:    &Cause{Choice: &CauseProtocol{Value: CauseProtocolPresentTransferSyntaxError}},
			causeStr: "Protocol: Transfer syntax error",
		},
		{
			name:     "Miscellaneous Group",
			cause:    &Cause{Choice: &CauseMisc{CauseMiscPresentNotEnoughUserPlaneProcessingResources}},
			causeStr: "Miscellaneous: Not enough user plane processing resources",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			require.Equal(t, tc.causeStr, tc.cause.String())
		})
	}
}

func TestBriefPDUSessResrcRelCmdXfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceReleaseCommandTransfer
		want string
	}{
		{
			name: "case normal release",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{
						Value: CauseNasPresentNormalRelease,
					},
				},
			},
			want: "PDUSessResrcRelCmdXfer:[NAS: Normal release]",
		},
		{
			name: "case authentication failure",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{
						Value: CauseNasPresentAuthenticationFailure,
					},
				},
			},
			want: "PDUSessResrcRelCmdXfer:[NAS: Authentication failure]",
		},
		{
			name: "case deregister",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{
						Value: CauseNasPresentDeregister,
					},
				},
			},
			want: "PDUSessResrcRelCmdXfer:[NAS: Deregister]",
		},
		{
			name: "case unspecif",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{
						Value: CauseNasPresentUnspecified,
					},
				},
			},
			want: "PDUSessResrcRelCmdXfer:[NAS: Unspecified]",
		},
		{
			name: "case empty CauseNas",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: &CauseNas{},
				},
			},
			// Value 0 mean Normal Release
			want: "PDUSessResrcRelCmdXfer:[NAS: Normal release]",
		},
		{
			name: "case nil cause",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: nil,
			},
			want: "PDUSessResrcRelCmdXfer:[]",
		},
		{
			name: "case nil Choice",
			xfer: &PDUSessionResourceReleaseCommandTransfer{
				Cause: &Cause{
					Choice: nil,
				},
			},
			want: "PDUSessResrcRelCmdXfer:[Invalid: Nil cause group]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefQosFlowSetupRequestList(t *testing.T) {
	testCases := []struct {
		name string
		list *QosFlowSetupRequestList
		want string
	}{
		{
			name: "case empty 1",
			list: &QosFlowSetupRequestList{},
			want: "QosSetup[]",
		},
		{
			name: "case empty 2",
			list: nil,
			want: "",
		},
		{
			name: "case QFI + 5QI",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowIdentifier: &QosFlowIdentifier{
							Value: 1,
						},
						QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
							QosCharacteristics: &QosCharacteristics{
								Choice: &NonDynamic5QIDescriptor{
									FiveQI: &FiveQI{
										Value: 5,
									},
								},
							},
						},
					},
				},
			},
			want: "QosSetup[QFI:1,5QI:5]",
		},
		{
			name: "case QFI + nil 5QI",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowIdentifier: &QosFlowIdentifier{
							Value: 1,
						},
						QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
							QosCharacteristics: &QosCharacteristics{
								Choice: &NonDynamic5QIDescriptor{
									FiveQI: nil,
								},
							},
						},
					},
				},
			},
			want: "QosSetup[QFI:1]",
		},
		{
			name: "case QFI + nil Choice",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowIdentifier: &QosFlowIdentifier{
							Value: 1,
						},
						QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
							QosCharacteristics: &QosCharacteristics{
								Choice: nil,
							},
						},
					},
				},
			},
			want: "QosSetup[QFI:1]",
		},
		{
			name: "case QFI + nil QosCharacteristics",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowIdentifier: &QosFlowIdentifier{
							Value: 1,
						},
						QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
							QosCharacteristics: nil,
						},
					},
				},
			},
			want: "QosSetup[QFI:1]",
		},
		{
			name: "case QFI + nil QosFlowLevelQosParameters",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowIdentifier: &QosFlowIdentifier{
							Value: 1,
						},
						QosFlowLevelQosParameters: nil,
					},
				},
			},
			want: "QosSetup[QFI:1]",
		},
		{
			name: "case nil QFI + nil QosFlowLevelQosParameters",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowIdentifier:         nil,
						QosFlowLevelQosParameters: nil,
					},
				},
			},
			want: "QosSetup[]",
		},
		{
			name: "case 5QI only",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
						QosFlowLevelQosParameters: &QosFlowLevelQosParameters{
							QosCharacteristics: &QosCharacteristics{
								Choice: &NonDynamic5QIDescriptor{
									FiveQI: &FiveQI{
										Value: 5,
									},
								},
							},
						},
					},
				},
			},
			want: "QosSetup[5QI:5]",
		},
		{
			name: "case QFI + 5QI + ARP",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
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
									Value: 6,
								},
								PreEmptionCapability: &PreEmptionCapability{
									Value: 0,
								},
								PreEmptionVulnerability: &PreEmptionVulnerability{
									Value: 1,
								},
							},
						},
					},
				},
			},
			want: "QosSetup[QFI:2,5QI:5,ArpLv:6,CapNoPrem,Prem-able]",
		},
		{
			name: "case +GBR, but no MBR",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
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
									Value: 6,
								},
								PreEmptionCapability: &PreEmptionCapability{
									Value: 0,
								},
								PreEmptionVulnerability: &PreEmptionVulnerability{
									Value: 1,
								},
							},
							GBRQosInformation: &GBRQosInformation{
								GuaranteedFlowBitRateUL: &BitRate{Value: 10},
								GuaranteedFlowBitRateDL: &BitRate{Value: 20},
							},
						},
					},
				},
			},
			want: "QosSetup[5QI:5,ArpLv:6,CapNoPrem,Prem-able,GBR[UL:10,DL:20]]",
		},
		{
			name: "case + GBR + MBR",
			list: &QosFlowSetupRequestList{
				List: []QosFlowSetupRequestItem{
					{
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
									Value: 6,
								},
								PreEmptionCapability: &PreEmptionCapability{
									Value: 0,
								},
								PreEmptionVulnerability: &PreEmptionVulnerability{
									Value: 1,
								},
							},
							GBRQosInformation: &GBRQosInformation{
								GuaranteedFlowBitRateUL: &BitRate{Value: 10},
								GuaranteedFlowBitRateDL: &BitRate{Value: 20},
								MaximumFlowBitRateUL:    &BitRate{Value: 30},
								MaximumFlowBitRateDL:    &BitRate{Value: 40},
							},
						},
					},
				},
			},
			want: "QosSetup[5QI:5,ArpLv:6,CapNoPrem,Prem-able,GBR[UL:10,DL:20],MBR[UL:30,DL:40]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.list.String())
		})
	}
}

func TestBriefPduSessResrcModReqXfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceModifyRequestTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PDUSessionResourceModifyRequestTransfer{},
			want: "PDUSessResrcModReqXfer:[]",
		},
		{
			name: "Empty AMBR",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[SessAMBR:[]]",
		},
		{
			name: "AMBR",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{Value: 20},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{Value: 10},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[SessAMBR:[DL:20,UL:10]]",
		},
		{
			name: "Empty QoS Release",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowToReleaseList: &QosFlowListWithCause{
								List: []QosFlowWithCauseItem{},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[QosToRel[]]",
		},
		{
			name: "QoS Release",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowToReleaseList: &QosFlowListWithCause{
								List: []QosFlowWithCauseItem{
									getQosFlowWithCauseItem(1, CauseNasPresentNormalRelease),
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[QosToRel[QFI:1,CauseNAS: Normal release]]",
		},
		{
			name: "AMBR + QoS Release",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowToReleaseList: &QosFlowListWithCause{
								List: []QosFlowWithCauseItem{
									getQosFlowWithCauseItem(1, CauseNasPresentNormalRelease),
								},
							},
						},
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{Value: 20},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{Value: 10},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[QosToRel[QFI:1,CauseNAS: Normal release],SessAMBR:[DL:20,UL:10]]",
		},
		{
			name: "Empty Qos Add Modify",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowAddOrModifyRequestList: &QosFlowAddOrModifyRequestList{
								List: []QosFlowAddOrModifyRequestItem{
									{},
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[QosAddOrMod[]]",
		},
		{
			name: "Qos Add Modify",
			xfer: &PDUSessionResourceModifyRequestTransfer{
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
										},
									},
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[QosAddOrMod[QFI:2,5QI:5]]",
		},
		{
			name: "AMBR, QosAddOrMod, QosRelease",
			xfer: &PDUSessionResourceModifyRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceModifyRequestTransferIEs{
					List: []PDUSessionResourceModifyRequestTransferIEs{
						{
							QosFlowToReleaseList: &QosFlowListWithCause{
								List: []QosFlowWithCauseItem{
									getQosFlowWithCauseItem(1, CauseNasPresentNormalRelease),
								},
							},
						},
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{Value: 20},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{Value: 10},
							},
						},
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
										},
									},
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcModReqXfer:[QosToRel[QFI:1,CauseNAS: Normal release],SessAMBR:[DL:20,UL:10],QosAddOrMod[QFI:2,5QI:5]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefPduSessResrcSetupReqXfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceSetupRequestTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PDUSessionResourceSetupRequestTransfer{},
			want: "PDUSessResrcSetupReqXfer:[]",
		},
		{
			name: "Empty AMBR",
			xfer: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{},
						},
					},
				},
			},
			want: "PDUSessResrcSetupReqXfer:[SessAMBR:[]]",
		},
		{
			name: "AMBR",
			xfer: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{Value: 20},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{Value: 10},
							},
						},
					},
				},
			},
			want: "PDUSessResrcSetupReqXfer:[SessAMBR:[DL:20,UL:10]]",
		},
		{
			name: "Empty Qos Setup",
			xfer: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
									{},
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcSetupReqXfer:[QosSetup[]]",
		},
		{
			name: "Qos Setup",
			xfer: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
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
										},
									},
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcSetupReqXfer:[QosSetup[QFI:2,5QI:5]]",
		},
		{
			name: "AMBR, QosSetup",
			xfer: &PDUSessionResourceSetupRequestTransfer{
				ProtocolIEs: &ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs{
					List: []PDUSessionResourceSetupRequestTransferIEs{
						{
							PDUSessionAggregateMaximumBitRate: &PDUSessionAggregateMaximumBitRate{
								PDUSessionAggregateMaximumBitRateDL: &BitRate{Value: 20},
								PDUSessionAggregateMaximumBitRateUL: &BitRate{Value: 10},
							},
						},
						{
							QosFlowSetupRequestList: &QosFlowSetupRequestList{
								List: []QosFlowSetupRequestItem{
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
										},
									},
								},
							},
						},
					},
				},
			},
			want: "PDUSessResrcSetupReqXfer:[SessAMBR:[DL:20,UL:10],QosSetup[QFI:2,5QI:5]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefPDUSessResrcSetupRspXfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceSetupResponseTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PDUSessionResourceSetupResponseTransfer{},
			want: "PDUSessResrcSetupRspXfer:[]",
		},
		{
			name: "Empty QosFlowFailedToSetupList",
			xfer: &PDUSessionResourceSetupResponseTransfer{
				QosFlowFailedToSetupList: &QosFlowListWithCause{
					List: []QosFlowWithCauseItem{},
				},
			},
			want: "PDUSessResrcSetupRspXfer:[QosFlowFailedToSetup[]]",
		},
		{
			name: "With QosFlowFailedToSetupList",
			xfer: &PDUSessionResourceSetupResponseTransfer{
				QosFlowFailedToSetupList: &QosFlowListWithCause{
					List: []QosFlowWithCauseItem{
						getQosFlowWithCauseItem(1, CauseNasPresentNormalRelease),
					},
				},
			},
			want: "PDUSessResrcSetupRspXfer:[QosFlowFailedToSetup[QFI:1,CauseNAS: Normal release]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefPDUSessionResourceModifyResponseTransfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceModifyResponseTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PDUSessionResourceModifyResponseTransfer{},
			want: "PDUSessResrcModRspXfer:[]",
		},
		{
			name: "With QosFlowFailedToModifyList",
			xfer: &PDUSessionResourceModifyResponseTransfer{
				QosFlowFailedToAddOrModifyList: &QosFlowListWithCause{
					List: []QosFlowWithCauseItem{
						getQosFlowWithCauseItem(1, CauseNasPresentNormalRelease),
					},
				},
			},
			want: "PDUSessResrcModRspXfer:[QosFailedToMod[QFI:1,CauseNAS: Normal release]]",
		},
		{
			name: "With QosFlowAddOrModRespList",
			xfer: &PDUSessionResourceModifyResponseTransfer{
				QosFlowAddOrModifyResponseList: &QosFlowAddOrModifyResponseList{
					List: []QosFlowAddOrModifyResponseItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
						},
					},
				},
			},
			want: "PDUSessResrcModRspXfer:[QosAddModRsp[QFI:1]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefPDUSessionResourceReleaseResponseTransfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceReleaseResponseTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PDUSessionResourceReleaseResponseTransfer{},
			want: "PDUSessResrcRelRspXfer",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefPathSwitchRequestTransfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PathSwitchRequestTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PathSwitchRequestTransfer{},
			want: "PathSwitchReqXfer:[]",
		},
		{
			name: "With QosFlowAcceptedList",
			xfer: &PathSwitchRequestTransfer{
				QosFlowAcceptedList: &QosFlowAcceptedList{
					List: []QosFlowAcceptedItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
						},
					},
				},
			},
			want: "PathSwitchReqXfer:[QosAccept[QFI:1]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefPDUSessionResourceNotifyTransfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *PDUSessionResourceNotifyTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &PDUSessionResourceNotifyTransfer{},
			want: "PDUSessResrcNotifyXfer:[]",
		},
		{
			name: "With Empty QosFlowNotifyList",
			xfer: &PDUSessionResourceNotifyTransfer{
				QosFlowNotifyList: &QosFlowNotifyList{
					List: []QosFlowNotifyItem{},
				},
			},
			want: "PDUSessResrcNotifyXfer:[QosNotify[]]",
		},
		{
			name: "With QosFlowNotifyList",
			xfer: &PDUSessionResourceNotifyTransfer{
				QosFlowNotifyList: &QosFlowNotifyList{
					List: []QosFlowNotifyItem{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 2,
							},
						},
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
						},
					},
				},
			},
			want: "PDUSessResrcNotifyXfer:[QosNotify[QFI:2,QFI:1]]",
		},
		{
			name: "With Empty QosFlowReleasedList",
			xfer: &PDUSessionResourceNotifyTransfer{
				QosFlowReleasedList: &QosFlowListWithCause{
					List: []QosFlowWithCauseItem{},
				},
			},
			want: "PDUSessResrcNotifyXfer:[QosRel[]]",
		},
		{
			name: "With QosFlowReleasedList",
			xfer: &PDUSessionResourceNotifyTransfer{
				QosFlowReleasedList: &QosFlowListWithCause{
					List: []QosFlowWithCauseItem{
						getQosFlowWithCauseItem(3, CauseNasPresentNormalRelease),
					},
				},
			},
			want: "PDUSessResrcNotifyXfer:[QosRel[QFI:3,CauseNAS: Normal release]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefHandoverRequiredTransfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *HandoverRequiredTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &HandoverRequiredTransfer{},
			want: "HandoverRequiredXfer",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func TestBriefHandoverRequestAcknowledgeTransfer(t *testing.T) {
	testCases := []struct {
		name string
		xfer *HandoverRequestAcknowledgeTransfer
		want string
	}{
		{
			name: "empty",
			xfer: &HandoverRequestAcknowledgeTransfer{},
			want: "HandoverReqAckXfer:[]",
		},
		{
			name: "With Empty QosFlowSetupResponseList",
			xfer: &HandoverRequestAcknowledgeTransfer{
				QosFlowSetupResponseList: &QosFlowListWithDataForwarding{
					List: []QosFlowItemWithDataForwarding{},
				},
			},
			want: "HandoverReqAckXfer:[QosSetup[]]",
		},
		{
			name: "With QosFlowSetupResponseList",
			xfer: &HandoverRequestAcknowledgeTransfer{
				QosFlowSetupResponseList: &QosFlowListWithDataForwarding{
					List: []QosFlowItemWithDataForwarding{
						{
							QosFlowIdentifier: &QosFlowIdentifier{
								Value: 1,
							},
							DataForwardingAccepted: &DataForwardingAccepted{
								Value: DataForwardingAcceptedPresentDataForwardingAccepted,
							},
						},
					},
				},
			},
			want: "HandoverReqAckXfer:[QosSetup[QFI:1,accepted]]",
		},
		{
			name: "With QosFlowFailedToSetupList",
			xfer: &HandoverRequestAcknowledgeTransfer{
				QosFlowFailedToSetupList: &QosFlowListWithCause{
					List: []QosFlowWithCauseItem{
						getQosFlowWithCauseItem(4, CauseNasPresentNormalRelease),
					},
				},
			},
			want: "HandoverReqAckXfer:[QosSetupFailed[QFI:4,CauseNAS: Normal release]]",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
			require.Equal(t, tc.want, tc.xfer.String())
		})
	}
}

func getQosFlowWithCauseItem(qfi uint8, cause aper.Enumerated) QosFlowWithCauseItem {
	return QosFlowWithCauseItem{
		QosFlowIdentifier: &QosFlowIdentifier{
			Value: int64(qfi),
		},
		Cause: &Cause{
			Choice: &CauseNas{
				Value: cause,
			},
		},
	}
}
