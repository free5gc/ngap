package ie

import (
	"fmt"
	"strconv"

	"github.com/free5gc/ngap/aper"
)

/* This file implements stringify functions for NGAP IE types.
   Since the go files in this directory are auto-generated,
   use a separated file to add the String() functions.
   Only fields of interest of IEs of interest are implemented.
*/

var CauseRadioNetworkStrings = map[aper.Enumerated]string{
	CauseRadioNetworkPresentUnspecified:                                   "Unspecified",
	CauseRadioNetworkPresentTxnrelocoverallExpiry:                         "TXnRELOCOverall expiry",
	CauseRadioNetworkPresentSuccessfulHandover:                            "Successful handover",
	CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason:              "Release due to NG-RAN generated reason",
	CauseRadioNetworkPresentReleaseDueTo5gcGeneratedReason:                "Release due to 5GC generated reason",
	CauseRadioNetworkPresentHandoverCancelled:                             "Handover cancelled",
	CauseRadioNetworkPresentPartialHandover:                               "Partial handover",
	CauseRadioNetworkPresentHoTargetNotAllowed:                            "Handover target not allowed",
	CauseRadioNetworkPresentTngrelocoverallExpiry:                         "TNGRELOCoverall expiry",
	CauseRadioNetworkPresentTngrelocprepExpiry:                            "TNGRELOCprep expiry",
	CauseRadioNetworkPresentCellNotAvailable:                              "Cell not available",
	CauseRadioNetworkPresentUnknownTargetID:                               "Unknown target ID",
	CauseRadioNetworkPresentNoRadioResourcesAvailableInTargetCell:         "No radio resources available in target cell",
	CauseRadioNetworkPresentUnknownLocalUENGAPID:                          "Unknown local UE NGAP ID",
	CauseRadioNetworkPresentInconsistentRemoteUENGAPID:                    "Inconsistent remote UE NGAP ID",
	CauseRadioNetworkPresentHandoverDesirableForRadioReason:               "Handover desirable for radio reasons",
	CauseRadioNetworkPresentTimeCriticalHandover:                          "Time critical handover",
	CauseRadioNetworkPresentResourceOptimisationHandover:                  "Resource optimisation handover",
	CauseRadioNetworkPresentReduceLoadInServingCell:                       "Reduce load in serving cell",
	CauseRadioNetworkPresentUserInactivity:                                "User inactivity",
	CauseRadioNetworkPresentRadioConnectionWithUeLost:                     "Radio connection with UE lost",
	CauseRadioNetworkPresentRadioResourcesNotAvailable:                    "Radio resources not available",
	CauseRadioNetworkPresentInvalidQosCombination:                         "Invalid QoS combination",
	CauseRadioNetworkPresentFailureInRadioInterfaceProcedure:              "Failure in the radio interface procedure",
	CauseRadioNetworkPresentInteractionWithOtherProcedure:                 "Interaction with other procedure",
	CauseRadioNetworkPresentUnknownPDUSessionID:                           "Unknown PDU Session ID",
	CauseRadioNetworkPresentUnkownQosFlowID:                               "Unknown QoS Flow ID",
	CauseRadioNetworkPresentMultiplePDUSessionIDInstances:                 "Multiple PDU Session ID Instances",
	CauseRadioNetworkPresentMultipleQosFlowIDInstances:                    "Multiple QoS Flow ID Instances",
	CauseRadioNetworkPresentNgIntraSystemHandoverTriggered:                "NG intra-system handover triggered",
	CauseRadioNetworkPresentNgInterSystemHandoverTriggered:                "NG inter-system handover triggered",
	CauseRadioNetworkPresentXnHandoverTriggered:                           "Xn handover triggered",
	CauseRadioNetworkPresentNotSupported5QIValue:                          "Not supported 5QI value",
	CauseRadioNetworkPresentUeContextTransfer:                             "UE context transfer",
	CauseRadioNetworkPresentImsVoiceEpsFallbackOrRatFallbackTriggered:     "IMS voice EPS fallback or RAT fallback triggered",
	CauseRadioNetworkPresentUpIntegrityProtectionNotPossible:              "UP integrity protection not possible",
	CauseRadioNetworkPresentUpConfidentialityProtectionNotPossible:        "UP confidentiality protection not possible",
	CauseRadioNetworkPresentSliceNotSupported:                             "Slice(s) not supported",
	CauseRadioNetworkPresentUeInRrcInactiveStateNotReachable:              "UE in RRC_INACTIVE state not reachable",
	CauseRadioNetworkPresentRedirection:                                   "Redirection",
	CauseRadioNetworkPresentResourcesNotAvailableForTheSlice:              "Resources not available for the slice(s)",
	CauseRadioNetworkPresentUeMaxIntegrityProtectedDataRateReason:         "UE maximum integrity protected data rate reason",
	CauseRadioNetworkPresentReleaseDueToCnDetectedMobility:                "Release due to CN-detected mobility",
	CauseRadioNetworkPresentN26InterfaceNotAvailable:                      "N26 interface not available",
	CauseRadioNetworkPresentReleaseDueToPreEmption:                        "Release due to pre-emption",
	CauseRadioNetworkPresentMultipleLocationReportingReferenceIDInstances: "Multiple Location Reporting Reference ID Instances",
	CauseRadioNetworkPresentRsnNotAvailableForTheUp:                       "RSN not available for the UP",
	CauseRadioNetworkPresentNpnAccessDenied:                               "NPN access denied",
	CauseRadioNetworkPresentCagOnlyAccessDenied:                           "CAG only access denied",
	CauseRadioNetworkPresentInsufficientUeCapabilities:                    "Insufficient UE capabilities",
	CauseRadioNetworkPresentRedcapUeNotSupported:                          "Redcap ue not supported",
	CauseRadioNetworkPresentUnknownMBSSessionID:                           "Unknown MBS session ID",
	CauseRadioNetworkPresentInconsistentSliceInfoForTheSession:            "Inconsistent slice info for the session",
	CauseRadioNetworkPresentMisalignedAssociationForMulticastUnicast:      "Misaligned association for multicast unicast",

	// avoids line too long
	CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem: "Handover failure in target 5GC/NG-RAN node " +
		"or target system",
	CauseRadioNetworkPresentEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported: "Encryption and/or integrity " +
		"protection algorithms not supported",
	CauseRadioNetworkPresentIndicatedMBSSessionAreaInformationNotServedByTheGNB: "Indicated MBS session area " +
		"information not served by the GNB",
}

var CauseTransportStrings = map[aper.Enumerated]string{
	CauseTransportPresentTransportResourceUnavailable: "Transport resource unavailable",
	CauseTransportPresentUnspecified:                  "Unspecified",
}

var CauseNasStrings = map[aper.Enumerated]string{
	CauseNasPresentNormalRelease:          "Normal release",
	CauseNasPresentAuthenticationFailure:  "Authentication failure",
	CauseNasPresentDeregister:             "Deregister",
	CauseNasPresentUnspecified:            "Unspecified",
	CauseNasPresentUENotInPLMNServingArea: "UE not in plmn serving area",
}

var CauseProtocolStrings = map[aper.Enumerated]string{
	CauseProtocolPresentTransferSyntaxError:                          "Transfer syntax error",
	CauseProtocolPresentAbstractSyntaxErrorReject:                    "Abstract syntax error (reject)",
	CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify:           "Abstract syntax error (ignore and notify)",
	CauseProtocolPresentMessageNotCompatibleWithReceiverState:        "Message not compatible with receiver state",
	CauseProtocolPresentSemanticError:                                "Semantic error",
	CauseProtocolPresentAbstractSyntaxErrorFalselyConstructedMessage: "Abstract syntax error (falsely constructed message)",
	CauseProtocolPresentUnspecified:                                  "Unspecified",
}

var CauseMiscStrings = map[aper.Enumerated]string{
	CauseMiscPresentControlProcessingOverload:             "Control processing overload",
	CauseMiscPresentNotEnoughUserPlaneProcessingResources: "Not enough user plane processing resources",
	CauseMiscPresentHardwareFailure:                       "Hardware failure",
	CauseMiscPresentOmIntervention:                        "O&M intervention",
	CauseMiscPresentUnknownPLMNOrSNPN:                     "Unknown PLMN or SNPN",
	CauseMiscPresentUnspecified:                           "Unspecified",
}

var NotificationCauseStrings = map[aper.Enumerated]string{
	NotificationCausePresentFulfilled:    "Fulfilled",
	NotificationCausePresentNotFulfilled: "Not fulfilled",
}

var DataForwardingAcceptedStrings = map[aper.Enumerated]string{
	DataForwardingAcceptedPresentDataForwardingAccepted: "accepted",
}

func concat(s1, s2 string) string {
	if s1 != "" && s2 != "" {
		s1 += ","
	}
	return s1 + s2
}

func (cause *Cause) String() string {
	if cause == nil {
		return ""
	}
	switch cg := cause.Choice.(type) {
	case *CauseRadioNetwork:
		return "Radio Network Layer: " + cg.String()
	case *CauseTransport:
		return "Transport Layer: " + cg.String()
	case *CauseNas:
		return "NAS: " + cg.String()
	case *CauseProtocol:
		return "Protocol: " + cg.String()
	case *CauseMisc:
		return "Miscellaneous: " + cg.String()
	case nil:
		return "Invalid: Nil cause group"
	}
	return fmt.Sprintf("Unregconized cause group(%d)", cause.Choice.CauseAltIndex())
}

func (cause *CauseRadioNetwork) String() string {
	if cause == nil {
		return ""
	}
	if str, ok := CauseRadioNetworkStrings[cause.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized RadioNetwork cause value(%d)", cause.Value)
}

func (cause *CauseTransport) String() string {
	if cause == nil {
		return ""
	}

	if str, ok := CauseTransportStrings[cause.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized Transport cause value(%d)", cause.Value)
}

func (cause *CauseNas) String() string {
	if cause == nil {
		return ""
	}
	if str, ok := CauseNasStrings[cause.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized NAS cause value(%d)", cause.Value)
}

func (cause *CauseProtocol) String() string {
	if cause == nil {
		return ""
	}
	if str, ok := CauseProtocolStrings[cause.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized Protocol cause value(%d)", cause.Value)
}

func (cause *CauseMisc) String() string {
	if cause == nil {
		return ""
	}
	if str, ok := CauseMiscStrings[cause.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized Miscellaneous cause value(%d)", cause.Value)
}

func (cause *NotificationCause) String() string {
	if cause == nil {
		return ""
	}
	if str, ok := NotificationCauseStrings[cause.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized Notification cause value(%d)", cause.Value)
}

func (dfa *DataForwardingAccepted) String() string {
	if dfa == nil {
		return ""
	}
	if str, ok := DataForwardingAcceptedStrings[dfa.Value]; ok {
		return str
	}
	return fmt.Sprintf("Unregconized Data Forwarding Accepted value(%d)", dfa.Value)
}

func (xfer *PDUSessionResourceReleaseCommandTransfer) String() string {
	if xfer == nil {
		return ""
	}
	return "PDUSessResrcRelCmdXfer:[" + xfer.Cause.String() + "]"
}

func (ambr *PDUSessionAggregateMaximumBitRate) String() string {
	if ambr == nil {
		return ""
	}

	var dl, ul string
	if bitRate := ambr.PDUSessionAggregateMaximumBitRateDL; bitRate != nil {
		dl = "DL:" + strconv.Itoa(int(bitRate.Value))
	}
	if bitRate := ambr.PDUSessionAggregateMaximumBitRateUL; bitRate != nil {
		ul = "UL:" + strconv.Itoa(int(bitRate.Value))
	}
	return "SessAMBR:[" + concat(dl, ul) + "]"
}

func (item *QosFlowWithCauseItem) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	cause := "Cause" + item.Cause.String()
	return concat(qfi, cause)
}

func (list *QosFlowListWithCause) String() string {
	if list == nil || len(list.List) == 0 {
		return ""
	}

	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return strs
}

func (xfer *PDUSessionResourceModifyRequestTransfer) String() string {
	if xfer == nil || xfer.ProtocolIEs == nil || len(xfer.ProtocolIEs.List) == 0 {
		return "PDUSessResrcModReqXfer:[]"
	}
	strs := ""
	for _, ie := range xfer.ProtocolIEs.List {
		if ie := ie.PDUSessionAggregateMaximumBitRate; ie != nil {
			strs = concat(strs, ie.String())
		}
		if list := ie.QosFlowAddOrModifyRequestList; list != nil {
			strs = concat(strs, list.String())
		}
		if ie.QosFlowToReleaseList != nil {
			// this list shares the same struct with QosFlowFailedToSetupList,
			// QosFlowFailedToResumeList. So print the list type here.
			toRel := "QosToRel[" + ie.QosFlowToReleaseList.String() + "]"
			strs = concat(strs, toRel)
		}
	}
	return "PDUSessResrcModReqXfer:[" + strs + "]"
}

// TODO: need test function
func (xfer *PDUSessionResourceSetupRequestTransfer) String() string {
	if xfer == nil || xfer.ProtocolIEs == nil || len(xfer.ProtocolIEs.List) == 0 {
		return "PDUSessResrcSetupReqXfer:[]"
	}

	strs := ""
	for _, ie := range xfer.ProtocolIEs.List {
		if ie := ie.PDUSessionAggregateMaximumBitRate; ie != nil {
			strs = concat(strs, ie.String())
		}
		if list := ie.QosFlowSetupRequestList; list != nil {
			strs = concat(strs, list.String())
		}
	}
	return "PDUSessResrcSetupReqXfer:[" + strs + "]"
}

func (list *QosFlowAddOrModifyRequestList) String() string {
	if list == nil {
		return ""
	}
	if len(list.List) == 0 {
		return "QosAddOrMod[]"
	}
	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return "QosAddOrMod[" + strs + "]"
}

func (list *QosFlowSetupRequestList) String() string {
	if list == nil {
		return ""
	}
	if len(list.List) == 0 {
		return "QosSetup[]"
	}
	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return "QosSetup[" + strs + "]"
}

func (item *QosFlowSetupRequestItem) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	param := item.QosFlowLevelQosParameters.String()
	return concat(qfi, param)
}

func (item *QosFlowAddOrModifyRequestItem) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	param := item.QosFlowLevelQosParameters.String()
	return concat(qfi, param)
}

func (list *QosFlowAddOrModifyResponseList) String() string {
	if list == nil {
		return ""
	}
	if len(list.List) == 0 {
		return ""
	}
	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return strs
}

func (item *QosFlowAddOrModifyResponseItem) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	return qfi
}

func (list *QosFlowAcceptedList) String() string {
	if list == nil {
		return ""
	}
	if len(list.List) == 0 {
		return ""
	}
	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return strs
}

func (item *QosFlowAcceptedItem) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	return qfi
}

func (list *QosFlowNotifyList) String() string {
	if list == nil {
		return ""
	}
	if len(list.List) == 0 {
		return ""
	}
	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return strs
}

func (item *QosFlowNotifyItem) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	notifyCause := item.NotificationCause.String()
	return concat(qfi, notifyCause)
}

func (list *QosFlowListWithDataForwarding) String() string {
	if list == nil {
		return ""
	}
	if len(list.List) == 0 {
		return ""
	}
	strs := ""
	for i := range list.List {
		item := &list.List[i]
		strs = concat(strs, item.String())
	}
	return strs
}

func (item *QosFlowItemWithDataForwarding) String() string {
	if item == nil {
		return ""
	}
	qfi := item.QosFlowIdentifier.String()
	notifyCause := item.DataForwardingAccepted.String()
	return concat(qfi, notifyCause)
}

func (qfi *QosFlowIdentifier) String() string {
	if qfi == nil {
		return ""
	}
	return "QFI:" + strconv.Itoa(int(qfi.Value))
}

func (param *QosFlowLevelQosParameters) String() string {
	if param == nil {
		return ""
	}

	fiveQi := ""
	if param.QosCharacteristics != nil && param.QosCharacteristics.Choice != nil {
		switch v := param.QosCharacteristics.Choice.(type) {
		case *NonDynamic5QIDescriptor:
			if v.FiveQI != nil {
				fiveQi += "5QI:" + strconv.Itoa(int(v.FiveQI.Value))
			}
		case *Dynamic5QIDescriptor:
			if v.FiveQI != nil {
				fiveQi += "5QI:" + strconv.Itoa(int(v.FiveQI.Value))
			}
		}
	}

	arp := ""
	if p := param.AllocationAndRetentionPriority; p != nil {
		if p.PriorityLevelARP != nil {
			arp = concat(arp, "ArpLv:"+strconv.Itoa(int(p.PriorityLevelARP.Value)))
		}
		if p.PreEmptionCapability != nil && p.PreEmptionVulnerability != nil {
			switch p.PreEmptionCapability.Value {
			case PreEmptionCapabilityPresentShallNotTriggerPreEmption:
				arp = concat(arp, "CapNoPrem")
			case PreEmptionCapabilityPresentMayTriggerPreEmption:
				arp = concat(arp, "CapMayPrem")
			}
			switch p.PreEmptionVulnerability.Value {
			case PreEmptionVulnerabilityPresentNotPreEmptable:
				arp = concat(arp, "nonPrem-able")
			case PreEmptionVulnerabilityPresentPreEmptable:
				arp = concat(arp, "Prem-able")
			}
		}
	}

	gbrStr := ""
	if gbr := param.GBRQosInformation; gbr != nil {
		tmp := ""
		if gbr.GuaranteedFlowBitRateUL != nil && gbr.GuaranteedFlowBitRateDL != nil {
			tmp += "GBR[UL:" + strconv.Itoa(int(gbr.GuaranteedFlowBitRateUL.Value)) + ","
			tmp += "DL:" + strconv.Itoa(int(gbr.GuaranteedFlowBitRateDL.Value)) + "]"
		}
		gbrStr = concat(gbrStr, tmp)
		tmp = ""
		if gbr.MaximumFlowBitRateUL != nil && gbr.MaximumFlowBitRateDL != nil {
			tmp += "MBR[UL:" + strconv.Itoa(int(gbr.MaximumFlowBitRateUL.Value)) + ","
			tmp += "DL:" + strconv.Itoa(int(gbr.MaximumFlowBitRateDL.Value)) + "]"
		}
		gbrStr = concat(gbrStr, tmp)
	}

	brief := ""
	if fiveQi != "" || arp != "" || gbrStr != "" {
		brief = concat(brief, fiveQi)
		brief = concat(brief, arp)
		brief = concat(brief, gbrStr)
	}
	return brief
}

func (xfer *PDUSessionResourceSetupResponseTransfer) String() string {
	strs := ""
	if list := xfer.QosFlowFailedToSetupList; list != nil {
		// this list shares the same struct with QosFlowToReleaseList,
		// QosFlowFailedToResumeList. So print the list type here.
		tmp := "QosFlowFailedToSetup[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	return "PDUSessResrcSetupRspXfer:[" + strs + "]"
}

func (xfer *PDUSessionResourceModifyResponseTransfer) String() string {
	strs := ""
	if list := xfer.QosFlowAddOrModifyResponseList; list != nil {
		tmp := "QosAddModRsp[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	if list := xfer.QosFlowFailedToAddOrModifyList; list != nil {
		tmp := "QosFailedToMod[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	return "PDUSessResrcModRspXfer:[" + strs + "]"
}

func (xfer *PDUSessionResourceReleaseResponseTransfer) String() string {
	return "PDUSessResrcRelRspXfer"
}

func (xfer *PathSwitchRequestTransfer) String() string {
	strs := ""
	if list := xfer.QosFlowAcceptedList; list != nil {
		tmp := "QosAccept[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	return "PathSwitchReqXfer:[" + strs + "]"
}

func (xfer *PDUSessionResourceNotifyTransfer) String() string {
	strs := ""
	if list := xfer.QosFlowNotifyList; list != nil {
		tmp := "QosNotify[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	if list := xfer.QosFlowReleasedList; list != nil {
		tmp := "QosRel[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	return "PDUSessResrcNotifyXfer:[" + strs + "]"
}

func (xfer *HandoverRequiredTransfer) String() string {
	return "HandoverRequiredXfer"
}

func (xfer *HandoverRequestAcknowledgeTransfer) String() string {
	strs := ""
	if list := xfer.QosFlowSetupResponseList; list != nil {
		tmp := "QosSetup[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	if list := xfer.QosFlowFailedToSetupList; list != nil {
		tmp := "QosSetupFailed[" + list.String() + "]"
		strs = concat(strs, tmp)
	}
	return "HandoverReqAckXfer:[" + strs + "]"
}
