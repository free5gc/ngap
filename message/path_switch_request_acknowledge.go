package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

type PathSwitchRequestAcknowledge struct {
	AMFUENGAPID                                 *ie.AMFUENGAPID                                 // mandatory
	RANUENGAPID                                 *ie.RANUENGAPID                                 // mandatory
	UESecurityCapabilities                      *ie.UESecurityCapabilities                      // optional
	SecurityContext                             *ie.SecurityContext                             // mandatory
	NewSecurityContextInd                       *ie.NewSecurityContextInd                       // optional
	PDUSessionResourceSwitchedList              *ie.PDUSessionResourceSwitchedList              // mandatory
	PDUSessionResourceReleasedListPSAck         *ie.PDUSessionResourceReleasedListPSAck         // optional
	AllowedNSSAI                                *ie.AllowedNSSAI                                // mandatory
	CoreNetworkAssistanceInformationForInactive *ie.CoreNetworkAssistanceInformationForInactive // optional
	RRCInactiveTransitionReportRequest          *ie.RRCInactiveTransitionReportRequest          // optional
	CriticalityDiagnostics                      *ie.CriticalityDiagnostics                      // optional
	RedirectionVoiceFallback                    *ie.RedirectionVoiceFallback                    // optional
	CNAssistedRANTuning                         *ie.CNAssistedRANTuning                         // optional
	SRVCCOperationPossible                      *ie.SRVCCOperationPossible                      // optional
	EnhancedCoverageRestriction                 *ie.EnhancedCoverageRestriction                 // optional
	ExtendedConnectedTime                       *ie.ExtendedConnectedTime                       // optional
	UEDifferentiationInfo                       *ie.UEDifferentiationInfo                       // optional
	NRV2XServicesAuthorized                     *ie.NRV2XServicesAuthorized                     // optional
	LTEV2XServicesAuthorized                    *ie.LTEV2XServicesAuthorized                    // optional
	NRUESidelinkAggregateMaximumBitrate         *ie.NRUESidelinkAggregateMaximumBitrate         // optional
	LTEUESidelinkAggregateMaximumBitrate        *ie.LTEUESidelinkAggregateMaximumBitrate        // optional
	PC5QoSParameters                            *ie.PC5QoSParameters                            // optional
	CEmodeBrestricted                           *ie.CEmodeBrestricted                           // optional
	UEUPCIoTSupport                             *ie.UEUPCIoTSupport                             // optional
	UERadioCapabilityID                         *ie.UERadioCapabilityID                         // optional
	ManagementBasedMDTPLMNList                  *ie.MDTPLMNList                                 // optional
	TimeSyncAssistanceInfo                      *ie.TimeSyncAssistanceInfo                      // optional
	FiveGProSeAuthorized                        *ie.FiveGProSeAuthorized                        // optional
	FiveGProSeUEPC5AggregateMaximumBitRate      *ie.NRUESidelinkAggregateMaximumBitrate         // optional
	FiveGProSePC5QoSParameters                  *ie.FiveGProSePC5QoSParameters                  // optional
	ManagementBasedMDTPLMNModificationList      *ie.MDTPLMNModificationList                     // optional
	IABAuthorized                               *ie.IABAuthorized                               // optional
}

func (x *PathSwitchRequestAcknowledge) MessageType() int64 {
	return MessageTypeSuccessfulOutcome
}

func (x *PathSwitchRequestAcknowledge) ProcedureCode() int64 {
	return ProcedureCodePathSwitchRequest
}

func (x *PathSwitchRequestAcknowledge) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *PathSwitchRequestAcknowledge) MarshalBinary() ([]byte, error) {
	pd := aper.NewPerBitData(nil)
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// encode ngappdu (CHOICE)
	*vUb = 2
	err = pd.WriteChoicePreambleBitMap(MessageTypeSuccessfulOutcome, true, vUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}

	// encode MessageTypeSuccessfulOutcome (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pd.WriteSequencePreambleBitMap(optPresentFlag, false)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProcedureCode
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(x.ProcedureCode(), false, vLb, vUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: Criticality
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(x.Criticality(), false, vLb, vUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// (#) sequence element: Value (Open Type)
	pdOpenType := aper.NewPerBitData(nil)

	// open type value: PathSwitchRequestAcknowledge (SEQUENCE)
	optPresentFlag = []bool{} // no optional field
	err = pdOpenType.WriteSequencePreambleBitMap(optPresentFlag, true)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// count number of ies as SEQUENCE OF SIZE
	numIes := uint64(0)
	if x.AMFUENGAPID != nil {
		numIes++
	}
	if x.RANUENGAPID != nil {
		numIes++
	}
	if x.UESecurityCapabilities != nil {
		numIes++
	}
	if x.SecurityContext != nil {
		numIes++
	}
	if x.NewSecurityContextInd != nil {
		numIes++
	}
	if x.PDUSessionResourceSwitchedList != nil {
		numIes++
	}
	if x.PDUSessionResourceReleasedListPSAck != nil {
		numIes++
	}
	if x.AllowedNSSAI != nil {
		numIes++
	}
	if x.CoreNetworkAssistanceInformationForInactive != nil {
		numIes++
	}
	if x.RRCInactiveTransitionReportRequest != nil {
		numIes++
	}
	if x.CriticalityDiagnostics != nil {
		numIes++
	}
	if x.RedirectionVoiceFallback != nil {
		numIes++
	}
	if x.CNAssistedRANTuning != nil {
		numIes++
	}
	if x.SRVCCOperationPossible != nil {
		numIes++
	}
	if x.EnhancedCoverageRestriction != nil {
		numIes++
	}
	if x.ExtendedConnectedTime != nil {
		numIes++
	}
	if x.UEDifferentiationInfo != nil {
		numIes++
	}
	if x.NRV2XServicesAuthorized != nil {
		numIes++
	}
	if x.LTEV2XServicesAuthorized != nil {
		numIes++
	}
	if x.NRUESidelinkAggregateMaximumBitrate != nil {
		numIes++
	}
	if x.LTEUESidelinkAggregateMaximumBitrate != nil {
		numIes++
	}
	if x.PC5QoSParameters != nil {
		numIes++
	}
	if x.CEmodeBrestricted != nil {
		numIes++
	}
	if x.UEUPCIoTSupport != nil {
		numIes++
	}
	if x.UERadioCapabilityID != nil {
		numIes++
	}
	if x.ManagementBasedMDTPLMNList != nil {
		numIes++
	}
	if x.TimeSyncAssistanceInfo != nil {
		numIes++
	}
	if x.FiveGProSeAuthorized != nil {
		numIes++
	}
	if x.FiveGProSeUEPC5AggregateMaximumBitRate != nil {
		numIes++
	}
	if x.FiveGProSePC5QoSParameters != nil {
		numIes++
	}
	if x.ManagementBasedMDTPLMNModificationList != nil {
		numIes++
	}
	if x.IABAuthorized != nil {
		numIes++
	}
	err = pdOpenType.WriteSequenceOfPreambleBitMap(numIes, false, sLb, sUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// encode if the IE field is present

	// IE Field 1 (mandatory)
	if x.AMFUENGAPID != nil {
		err = x.AMFUENGAPID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDAMFUENGAPID},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field AMFUENGAPID is missing")
	}

	// IE Field 2 (mandatory)
	if x.RANUENGAPID != nil {
		err = x.RANUENGAPID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRANUENGAPID},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field RANUENGAPID is missing")
	}

	// IE Field 3 (optional)
	if x.UESecurityCapabilities != nil {
		err = x.UESecurityCapabilities.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUESecurityCapabilities},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 4 (mandatory)
	if x.SecurityContext != nil {
		err = x.SecurityContext.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSecurityContext},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field SecurityContext is missing")
	}

	// IE Field 5 (optional)
	if x.NewSecurityContextInd != nil {
		err = x.NewSecurityContextInd.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNewSecurityContextInd},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 6 (mandatory)
	if x.PDUSessionResourceSwitchedList != nil {
		err = x.PDUSessionResourceSwitchedList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPDUSessionResourceSwitchedList},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field PDUSessionResourceSwitchedList is missing")
	}

	// IE Field 7 (optional)
	if x.PDUSessionResourceReleasedListPSAck != nil {
		err = x.PDUSessionResourceReleasedListPSAck.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPDUSessionResourceReleasedListPSAck},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 8 (mandatory)
	if x.AllowedNSSAI != nil {
		err = x.AllowedNSSAI.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDAllowedNSSAI},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field AllowedNSSAI is missing")
	}

	// IE Field 9 (optional)
	if x.CoreNetworkAssistanceInformationForInactive != nil {
		err = x.CoreNetworkAssistanceInformationForInactive.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCoreNetworkAssistanceInformationForInactive},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 10 (optional)
	if x.RRCInactiveTransitionReportRequest != nil {
		err = x.RRCInactiveTransitionReportRequest.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRRCInactiveTransitionReportRequest},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 11 (optional)
	if x.CriticalityDiagnostics != nil {
		err = x.CriticalityDiagnostics.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCriticalityDiagnostics},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 12 (optional)
	if x.RedirectionVoiceFallback != nil {
		err = x.RedirectionVoiceFallback.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRedirectionVoiceFallback},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 13 (optional)
	if x.CNAssistedRANTuning != nil {
		err = x.CNAssistedRANTuning.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCNAssistedRANTuning},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 14 (optional)
	if x.SRVCCOperationPossible != nil {
		err = x.SRVCCOperationPossible.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSRVCCOperationPossible},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 15 (optional)
	if x.EnhancedCoverageRestriction != nil {
		err = x.EnhancedCoverageRestriction.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDEnhancedCoverageRestriction},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 16 (optional)
	if x.ExtendedConnectedTime != nil {
		err = x.ExtendedConnectedTime.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDExtendedConnectedTime},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 17 (optional)
	if x.UEDifferentiationInfo != nil {
		err = x.UEDifferentiationInfo.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUEDifferentiationInfo},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 18 (optional)
	if x.NRV2XServicesAuthorized != nil {
		err = x.NRV2XServicesAuthorized.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNRV2XServicesAuthorized},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 19 (optional)
	if x.LTEV2XServicesAuthorized != nil {
		err = x.LTEV2XServicesAuthorized.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDLTEV2XServicesAuthorized},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 20 (optional)
	if x.NRUESidelinkAggregateMaximumBitrate != nil {
		err = x.NRUESidelinkAggregateMaximumBitrate.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNRUESidelinkAggregateMaximumBitrate},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 21 (optional)
	if x.LTEUESidelinkAggregateMaximumBitrate != nil {
		err = x.LTEUESidelinkAggregateMaximumBitrate.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 22 (optional)
	if x.PC5QoSParameters != nil {
		err = x.PC5QoSParameters.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPC5QoSParameters},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 23 (optional)
	if x.CEmodeBrestricted != nil {
		err = x.CEmodeBrestricted.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCEmodeBrestricted},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 24 (optional)
	if x.UEUPCIoTSupport != nil {
		err = x.UEUPCIoTSupport.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUEUPCIoTSupport},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 25 (optional)
	if x.UERadioCapabilityID != nil {
		err = x.UERadioCapabilityID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUERadioCapabilityID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 26 (optional)
	if x.ManagementBasedMDTPLMNList != nil {
		err = x.ManagementBasedMDTPLMNList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDManagementBasedMDTPLMNList},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 27 (optional)
	if x.TimeSyncAssistanceInfo != nil {
		err = x.TimeSyncAssistanceInfo.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDTimeSyncAssistanceInfo},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 28 (optional)
	if x.FiveGProSeAuthorized != nil {
		err = x.FiveGProSeAuthorized.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDFiveGProSeAuthorized},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 29 (optional)
	if x.FiveGProSeUEPC5AggregateMaximumBitRate != nil {
		err = x.FiveGProSeUEPC5AggregateMaximumBitRate.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 30 (optional)
	if x.FiveGProSePC5QoSParameters != nil {
		err = x.FiveGProSePC5QoSParameters.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDFiveGProSePC5QoSParameters},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 31 (optional)
	if x.ManagementBasedMDTPLMNModificationList != nil {
		err = x.ManagementBasedMDTPLMNModificationList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDManagementBasedMDTPLMNModificationList},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 32 (optional)
	if x.IABAuthorized != nil {
		err = x.IABAuthorized.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDIABAuthorized},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// Finish MessageTypeSuccessfulOutcome open type Value Encoding (#)
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}

	return pd.Bytes(), nil
}

func (x *PathSwitchRequestAcknowledge) UnmarshalBinary(marshalled []byte) error {
	pd := aper.NewPerBitData(marshalled)
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)
	foo(err, sLb, sUb, vLb, vUb)

	// Some fields are decoded already:
	// decode ngappdu (CHOICE)
	// decode MessageTypeSuccessfulOutcome (SEQUENCE)
	// sequence element: ProcedureCode
	// sequence element: Criticality

	// (#) sequence element: Value (Open Type)
	// Read Open Type byte
	var bytes []byte
	bytes, err = pd.ReadOpenType()
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode message Value (open-type) error"))
	}
	pdOpenType := aper.NewPerBitData(bytes)

	// open type value: PathSwitchRequestAcknowledge (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (PathSwitchRequestAcknowledge) sequence error"))
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// get number of ies
	var numIes uint64
	numIes, err = pdOpenType.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode ProtocolIE-Container (seqof) error"))
	}

	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// read IEs
	foo(numIes)

	for i := 0; i < int(numIes); i++ {
		var protocolIeId int64
		protocolIeId, err = ie.ReadProtocolIEID(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "message unmarshal failed")
		}
		// IE Field 1
		if protocolIeId == ie.ProtocolIEIDAMFUENGAPID {
			// check if ie is duplicated
			if x.AMFUENGAPID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: AMFUENGAPID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.AMFUENGAPID = &ie.AMFUENGAPID{}
			err = x.AMFUENGAPID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 2
		if protocolIeId == ie.ProtocolIEIDRANUENGAPID {
			// check if ie is duplicated
			if x.RANUENGAPID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RANUENGAPID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RANUENGAPID = &ie.RANUENGAPID{}
			err = x.RANUENGAPID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
		if protocolIeId == ie.ProtocolIEIDUESecurityCapabilities {
			// check if ie is duplicated
			if x.UESecurityCapabilities != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UESecurityCapabilities")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UESecurityCapabilities = &ie.UESecurityCapabilities{}
			err = x.UESecurityCapabilities.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
		if protocolIeId == ie.ProtocolIEIDSecurityContext {
			// check if ie is duplicated
			if x.SecurityContext != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SecurityContext")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SecurityContext = &ie.SecurityContext{}
			err = x.SecurityContext.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDNewSecurityContextInd {
			// check if ie is duplicated
			if x.NewSecurityContextInd != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NewSecurityContextInd")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NewSecurityContextInd = &ie.NewSecurityContextInd{}
			err = x.NewSecurityContextInd.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDPDUSessionResourceSwitchedList {
			// check if ie is duplicated
			if x.PDUSessionResourceSwitchedList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PDUSessionResourceSwitchedList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PDUSessionResourceSwitchedList = &ie.PDUSessionResourceSwitchedList{}
			err = x.PDUSessionResourceSwitchedList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDPDUSessionResourceReleasedListPSAck {
			// check if ie is duplicated
			if x.PDUSessionResourceReleasedListPSAck != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PDUSessionResourceReleasedListPSAck")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PDUSessionResourceReleasedListPSAck = &ie.PDUSessionResourceReleasedListPSAck{}
			err = x.PDUSessionResourceReleasedListPSAck.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 8
		if protocolIeId == ie.ProtocolIEIDAllowedNSSAI {
			// check if ie is duplicated
			if x.AllowedNSSAI != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: AllowedNSSAI")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.AllowedNSSAI = &ie.AllowedNSSAI{}
			err = x.AllowedNSSAI.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 9
		if protocolIeId == ie.ProtocolIEIDCoreNetworkAssistanceInformationForInactive {
			// check if ie is duplicated
			if x.CoreNetworkAssistanceInformationForInactive != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: CoreNetworkAssistanceInformationForInactive")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.CoreNetworkAssistanceInformationForInactive = &ie.CoreNetworkAssistanceInformationForInactive{}
			err = x.CoreNetworkAssistanceInformationForInactive.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 10
		if protocolIeId == ie.ProtocolIEIDRRCInactiveTransitionReportRequest {
			// check if ie is duplicated
			if x.RRCInactiveTransitionReportRequest != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RRCInactiveTransitionReportRequest")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RRCInactiveTransitionReportRequest = &ie.RRCInactiveTransitionReportRequest{}
			err = x.RRCInactiveTransitionReportRequest.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 11
		if protocolIeId == ie.ProtocolIEIDCriticalityDiagnostics {
			// check if ie is duplicated
			if x.CriticalityDiagnostics != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: CriticalityDiagnostics")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.CriticalityDiagnostics = &ie.CriticalityDiagnostics{}
			err = x.CriticalityDiagnostics.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 12
		if protocolIeId == ie.ProtocolIEIDRedirectionVoiceFallback {
			// check if ie is duplicated
			if x.RedirectionVoiceFallback != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RedirectionVoiceFallback")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RedirectionVoiceFallback = &ie.RedirectionVoiceFallback{}
			err = x.RedirectionVoiceFallback.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 13
		if protocolIeId == ie.ProtocolIEIDCNAssistedRANTuning {
			// check if ie is duplicated
			if x.CNAssistedRANTuning != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: CNAssistedRANTuning")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.CNAssistedRANTuning = &ie.CNAssistedRANTuning{}
			err = x.CNAssistedRANTuning.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 14
		if protocolIeId == ie.ProtocolIEIDSRVCCOperationPossible {
			// check if ie is duplicated
			if x.SRVCCOperationPossible != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SRVCCOperationPossible")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SRVCCOperationPossible = &ie.SRVCCOperationPossible{}
			err = x.SRVCCOperationPossible.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 15
		if protocolIeId == ie.ProtocolIEIDEnhancedCoverageRestriction {
			// check if ie is duplicated
			if x.EnhancedCoverageRestriction != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: EnhancedCoverageRestriction")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.EnhancedCoverageRestriction = &ie.EnhancedCoverageRestriction{}
			err = x.EnhancedCoverageRestriction.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 16
		if protocolIeId == ie.ProtocolIEIDExtendedConnectedTime {
			// check if ie is duplicated
			if x.ExtendedConnectedTime != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ExtendedConnectedTime")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ExtendedConnectedTime = &ie.ExtendedConnectedTime{}
			err = x.ExtendedConnectedTime.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 17
		if protocolIeId == ie.ProtocolIEIDUEDifferentiationInfo {
			// check if ie is duplicated
			if x.UEDifferentiationInfo != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UEDifferentiationInfo")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UEDifferentiationInfo = &ie.UEDifferentiationInfo{}
			err = x.UEDifferentiationInfo.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 18
		if protocolIeId == ie.ProtocolIEIDNRV2XServicesAuthorized {
			// check if ie is duplicated
			if x.NRV2XServicesAuthorized != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NRV2XServicesAuthorized")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NRV2XServicesAuthorized = &ie.NRV2XServicesAuthorized{}
			err = x.NRV2XServicesAuthorized.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 19
		if protocolIeId == ie.ProtocolIEIDLTEV2XServicesAuthorized {
			// check if ie is duplicated
			if x.LTEV2XServicesAuthorized != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: LTEV2XServicesAuthorized")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.LTEV2XServicesAuthorized = &ie.LTEV2XServicesAuthorized{}
			err = x.LTEV2XServicesAuthorized.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 20
		if protocolIeId == ie.ProtocolIEIDNRUESidelinkAggregateMaximumBitrate {
			// check if ie is duplicated
			if x.NRUESidelinkAggregateMaximumBitrate != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NRUESidelinkAggregateMaximumBitrate")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NRUESidelinkAggregateMaximumBitrate = &ie.NRUESidelinkAggregateMaximumBitrate{}
			err = x.NRUESidelinkAggregateMaximumBitrate.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 21
		if protocolIeId == ie.ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate {
			// check if ie is duplicated
			if x.LTEUESidelinkAggregateMaximumBitrate != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: LTEUESidelinkAggregateMaximumBitrate")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.LTEUESidelinkAggregateMaximumBitrate = &ie.LTEUESidelinkAggregateMaximumBitrate{}
			err = x.LTEUESidelinkAggregateMaximumBitrate.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 22
		if protocolIeId == ie.ProtocolIEIDPC5QoSParameters {
			// check if ie is duplicated
			if x.PC5QoSParameters != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PC5QoSParameters")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PC5QoSParameters = &ie.PC5QoSParameters{}
			err = x.PC5QoSParameters.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 23
		if protocolIeId == ie.ProtocolIEIDCEmodeBrestricted {
			// check if ie is duplicated
			if x.CEmodeBrestricted != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: CEmodeBrestricted")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.CEmodeBrestricted = &ie.CEmodeBrestricted{}
			err = x.CEmodeBrestricted.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 24
		if protocolIeId == ie.ProtocolIEIDUEUPCIoTSupport {
			// check if ie is duplicated
			if x.UEUPCIoTSupport != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UEUPCIoTSupport")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UEUPCIoTSupport = &ie.UEUPCIoTSupport{}
			err = x.UEUPCIoTSupport.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 25
		if protocolIeId == ie.ProtocolIEIDUERadioCapabilityID {
			// check if ie is duplicated
			if x.UERadioCapabilityID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UERadioCapabilityID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UERadioCapabilityID = &ie.UERadioCapabilityID{}
			err = x.UERadioCapabilityID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 26
		if protocolIeId == ie.ProtocolIEIDManagementBasedMDTPLMNList {
			// check if ie is duplicated
			if x.ManagementBasedMDTPLMNList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ManagementBasedMDTPLMNList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ManagementBasedMDTPLMNList = &ie.MDTPLMNList{}
			err = x.ManagementBasedMDTPLMNList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 27
		if protocolIeId == ie.ProtocolIEIDTimeSyncAssistanceInfo {
			// check if ie is duplicated
			if x.TimeSyncAssistanceInfo != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: TimeSyncAssistanceInfo")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.TimeSyncAssistanceInfo = &ie.TimeSyncAssistanceInfo{}
			err = x.TimeSyncAssistanceInfo.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 28
		if protocolIeId == ie.ProtocolIEIDFiveGProSeAuthorized {
			// check if ie is duplicated
			if x.FiveGProSeAuthorized != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: FiveGProSeAuthorized")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.FiveGProSeAuthorized = &ie.FiveGProSeAuthorized{}
			err = x.FiveGProSeAuthorized.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 29
		if protocolIeId == ie.ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate {
			// check if ie is duplicated
			if x.FiveGProSeUEPC5AggregateMaximumBitRate != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: FiveGProSeUEPC5AggregateMaximumBitRate")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.FiveGProSeUEPC5AggregateMaximumBitRate = &ie.NRUESidelinkAggregateMaximumBitrate{}
			err = x.FiveGProSeUEPC5AggregateMaximumBitRate.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 30
		if protocolIeId == ie.ProtocolIEIDFiveGProSePC5QoSParameters {
			// check if ie is duplicated
			if x.FiveGProSePC5QoSParameters != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: FiveGProSePC5QoSParameters")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.FiveGProSePC5QoSParameters = &ie.FiveGProSePC5QoSParameters{}
			err = x.FiveGProSePC5QoSParameters.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 31
		if protocolIeId == ie.ProtocolIEIDManagementBasedMDTPLMNModificationList {
			// check if ie is duplicated
			if x.ManagementBasedMDTPLMNModificationList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ManagementBasedMDTPLMNModificationList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ManagementBasedMDTPLMNModificationList = &ie.MDTPLMNModificationList{}
			err = x.ManagementBasedMDTPLMNModificationList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 32
		if protocolIeId == ie.ProtocolIEIDIABAuthorized {
			// check if ie is duplicated
			if x.IABAuthorized != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: IABAuthorized")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.IABAuthorized = &ie.IABAuthorized{}
			err = x.IABAuthorized.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// Unknown IE ID
		crit, err := ie.UnmarshalUnknownIE(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "message unmarshal failed")
		}
		// If the unknown IE's criticality is reject, raise an Abstract Syntax Error
		if crit.Value == CriticalityReject {
			reportIe := ie.BuildAbstractSyntaxErrReportIe(protocolIeId, crit.Value)
			errTrace := errors.Errorf("Unknown IE ID: %d [Criticality: rej]", protocolIeId)
			return ie.BuildAbstractSyntaxErr(
				x.ProcedureCode(),
				aper.Enumerated(x.MessageType()),
				x.Criticality(),
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}

	}

	// Check if mandatory field(s) is present
	// AMFUENGAPID is mandatory but may be nil (ignored)

	// RANUENGAPID is mandatory but may be nil (ignored)

	if x.SecurityContext == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDSecurityContext, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE SecurityContext is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	// PDUSessionResourceSwitchedList is mandatory but may be nil (ignored)

	if x.AllowedNSSAI == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDAllowedNSSAI, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE AllowedNSSAI is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	return nil
}
