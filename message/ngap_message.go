package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

// TS 38.412
const PPID uint32 = 0x3c000000

type Message interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
	MessageType() int64
	ProcedureCode() int64
	Criticality() aper.Enumerated
}

func ParseMessageType(marshalled []byte) (Message, []byte, error) {
	pd := aper.NewPerBitData(marshalled)
	var err error
	var vLb, vUb *int64 = new(int64), new(int64)
	var target_msg Message

	// decode ngappdu (CHOICE)
	*vUb = 2
	ngappduType, err := pd.ReadChoicePreambleBitMap(true, vUb)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// decode MessageTypeInitiatingMessage (SEQUENCE)
	OptPresentFlag := []bool{} // no optional field
	err = pd.ReadSequencePreambleBitMap(&OptPresentFlag, false)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode error"))
	}
	// sequence element: ProcedureCode
	*vLb, *vUb = 0, 255
	ngappduProcedureCode, err := pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	// sequence element: Criticality
	*vLb, *vUb = 0, 2
	ngappduCrit, err := pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// use post-process/gen_parse_message_type.py to generate the codes in swich ngappduType{}
	switch ngappduType {
	// REPLACE START
	case MessageTypeInitiatingMessage:
		switch ngappduProcedureCode {
		case ProcedureCodeAMFConfigurationUpdate:
			target_msg = &AMFConfigurationUpdate{}
		case ProcedureCodeAMFCPRelocationIndication:
			target_msg = &AMFCPRelocationIndication{}
		case ProcedureCodeAMFStatusIndication:
			target_msg = &AMFStatusIndication{}
		case ProcedureCodeBroadcastSessionModification:
			target_msg = &BroadcastSessionModificationRequest{}
		case ProcedureCodeBroadcastSessionRelease:
			target_msg = &BroadcastSessionReleaseRequest{}
		case ProcedureCodeBroadcastSessionReleaseRequired:
			target_msg = &BroadcastSessionReleaseRequired{}
		case ProcedureCodeBroadcastSessionSetup:
			target_msg = &BroadcastSessionSetupRequest{}
		case ProcedureCodeCellTrafficTrace:
			target_msg = &CellTrafficTrace{}
		case ProcedureCodeConnectionEstablishmentIndication:
			target_msg = &ConnectionEstablishmentIndication{}
		case ProcedureCodeDeactivateTrace:
			target_msg = &DeactivateTrace{}
		case ProcedureCodeDistributionSetup:
			target_msg = &DistributionSetupRequest{}
		case ProcedureCodeDistributionRelease:
			target_msg = &DistributionReleaseRequest{}
		case ProcedureCodeDownlinkNASTransport:
			target_msg = &DownlinkNASTransport{}
		case ProcedureCodeDownlinkNonUEAssociatedNRPPaTransport:
			target_msg = &DownlinkNonUEAssociatedNRPPaTransport{}
		case ProcedureCodeDownlinkRANConfigurationTransfer:
			target_msg = &DownlinkRANConfigurationTransfer{}
		case ProcedureCodeDownlinkRANEarlyStatusTransfer:
			target_msg = &DownlinkRANEarlyStatusTransfer{}
		case ProcedureCodeDownlinkRANStatusTransfer:
			target_msg = &DownlinkRANStatusTransfer{}
		case ProcedureCodeDownlinkUEAssociatedNRPPaTransport:
			target_msg = &DownlinkUEAssociatedNRPPaTransport{}
		case ProcedureCodeErrorIndication:
			target_msg = &ErrorIndication{}
		case ProcedureCodeHandoverCancel:
			target_msg = &HandoverCancel{}
		case ProcedureCodeHandoverNotification:
			target_msg = &HandoverNotify{}
		case ProcedureCodeHandoverPreparation:
			target_msg = &HandoverRequired{}
		case ProcedureCodeHandoverResourceAllocation:
			target_msg = &HandoverRequest{}
		case ProcedureCodeHandoverSuccess:
			target_msg = &HandoverSuccess{}
		case ProcedureCodeInitialContextSetup:
			target_msg = &InitialContextSetupRequest{}
		case ProcedureCodeInitialUEMessage:
			target_msg = &InitialUEMessage{}
		case ProcedureCodeLocationReport:
			target_msg = &LocationReport{}
		case ProcedureCodeLocationReportingControl:
			target_msg = &LocationReportingControl{}
		case ProcedureCodeLocationReportingFailureIndication:
			target_msg = &LocationReportingFailureIndication{}
		case ProcedureCodeMulticastSessionActivation:
			target_msg = &MulticastSessionActivationRequest{}
		case ProcedureCodeMulticastSessionDeactivation:
			target_msg = &MulticastSessionDeactivationRequest{}
		case ProcedureCodeMulticastSessionUpdate:
			target_msg = &MulticastSessionUpdateRequest{}
		case ProcedureCodeMulticastGroupPaging:
			target_msg = &MulticastGroupPaging{}
		case ProcedureCodeNASNonDeliveryIndication:
			target_msg = &NASNonDeliveryIndication{}
		case ProcedureCodeNGReset:
			target_msg = &NGReset{}
		case ProcedureCodeNGSetup:
			target_msg = &NGSetupRequest{}
		case ProcedureCodeOverloadStart:
			target_msg = &OverloadStart{}
		case ProcedureCodeOverloadStop:
			target_msg = &OverloadStop{}
		case ProcedureCodePaging:
			target_msg = &Paging{}
		case ProcedureCodePathSwitchRequest:
			target_msg = &PathSwitchRequest{}
		case ProcedureCodePDUSessionResourceModify:
			target_msg = &PDUSessionResourceModifyRequest{}
		case ProcedureCodePDUSessionResourceModifyIndication:
			target_msg = &PDUSessionResourceModifyIndication{}
		case ProcedureCodePDUSessionResourceNotify:
			target_msg = &PDUSessionResourceNotify{}
		case ProcedureCodePDUSessionResourceRelease:
			target_msg = &PDUSessionResourceReleaseCommand{}
		case ProcedureCodePDUSessionResourceSetup:
			target_msg = &PDUSessionResourceSetupRequest{}
		case ProcedureCodePrivateMessage:
			target_msg = &PrivateMessage{}
		case ProcedureCodePWSCancel:
			target_msg = &PWSCancelRequest{}
		case ProcedureCodePWSFailureIndication:
			target_msg = &PWSFailureIndication{}
		case ProcedureCodePWSRestartIndication:
			target_msg = &PWSRestartIndication{}
		case ProcedureCodeRANConfigurationUpdate:
			target_msg = &RANConfigurationUpdate{}
		case ProcedureCodeRANCPRelocationIndication:
			target_msg = &RANCPRelocationIndication{}
		case ProcedureCodeRerouteNASRequest:
			target_msg = &RerouteNASRequest{}
		case ProcedureCodeRetrieveUEInformation:
			target_msg = &RetrieveUEInformation{}
		case ProcedureCodeRRCInactiveTransitionReport:
			target_msg = &RRCInactiveTransitionReport{}
		case ProcedureCodeSecondaryRATDataUsageReport:
			target_msg = &SecondaryRATDataUsageReport{}
		case ProcedureCodeTraceFailureIndication:
			target_msg = &TraceFailureIndication{}
		case ProcedureCodeTraceStart:
			target_msg = &TraceStart{}
		case ProcedureCodeUEContextModification:
			target_msg = &UEContextModificationRequest{}
		case ProcedureCodeUEContextRelease:
			target_msg = &UEContextReleaseCommand{}
		case ProcedureCodeUEContextReleaseRequest:
			target_msg = &UEContextReleaseRequest{}
		case ProcedureCodeUEContextResume:
			target_msg = &UEContextResumeRequest{}
		case ProcedureCodeUEContextSuspend:
			target_msg = &UEContextSuspendRequest{}
		case ProcedureCodeUEInformationTransfer:
			target_msg = &UEInformationTransfer{}
		case ProcedureCodeUERadioCapabilityCheck:
			target_msg = &UERadioCapabilityCheckRequest{}
		case ProcedureCodeUERadioCapabilityIDMapping:
			target_msg = &UERadioCapabilityIDMappingRequest{}
		case ProcedureCodeUERadioCapabilityInfoIndication:
			target_msg = &UERadioCapabilityInfoIndication{}
		case ProcedureCodeUETNLABindingRelease:
			target_msg = &UETNLABindingReleaseRequest{}
		case ProcedureCodeUplinkNASTransport:
			target_msg = &UplinkNASTransport{}
		case ProcedureCodeUplinkNonUEAssociatedNRPPaTransport:
			target_msg = &UplinkNonUEAssociatedNRPPaTransport{}
		case ProcedureCodeUplinkRANConfigurationTransfer:
			target_msg = &UplinkRANConfigurationTransfer{}
		case ProcedureCodeUplinkRANEarlyStatusTransfer:
			target_msg = &UplinkRANEarlyStatusTransfer{}
		case ProcedureCodeUplinkRANStatusTransfer:
			target_msg = &UplinkRANStatusTransfer{}
		case ProcedureCodeUplinkUEAssociatedNRPPaTransport:
			target_msg = &UplinkUEAssociatedNRPPaTransport{}
		case ProcedureCodeWriteReplaceWarning:
			target_msg = &WriteReplaceWarningRequest{}
		case ProcedureCodeUplinkRIMInformationTransfer:
			target_msg = &UplinkRIMInformationTransfer{}
		case ProcedureCodeDownlinkRIMInformationTransfer:
			target_msg = &DownlinkRIMInformationTransfer{}

		default:
			// ProcedureCode IE is in the ie group MessageType (use MessageType criticality)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdProcedureCode, ngappduCrit)
			errTrace := errors.Errorf("Unknown Procedure Code: %d", ngappduProcedureCode)
			err = ie.BuildAbstractSyntaxErr(
				ngappduProcedureCode,
				aper.Enumerated(MessageTypeInitiatingMessage),
				ngappduCrit,
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	case MessageTypeSuccessfulOutcome:
		switch ngappduProcedureCode {
		case ProcedureCodeAMFConfigurationUpdate:
			target_msg = &AMFConfigurationUpdateAcknowledge{}
		case ProcedureCodeBroadcastSessionModification:
			target_msg = &BroadcastSessionModificationResponse{}
		case ProcedureCodeBroadcastSessionRelease:
			target_msg = &BroadcastSessionReleaseResponse{}
		case ProcedureCodeBroadcastSessionSetup:
			target_msg = &BroadcastSessionSetupResponse{}
		case ProcedureCodeDistributionSetup:
			target_msg = &DistributionSetupResponse{}
		case ProcedureCodeDistributionRelease:
			target_msg = &DistributionReleaseResponse{}
		case ProcedureCodeHandoverCancel:
			target_msg = &HandoverCancelAcknowledge{}
		case ProcedureCodeHandoverPreparation:
			target_msg = &HandoverCommand{}
		case ProcedureCodeHandoverResourceAllocation:
			target_msg = &HandoverRequestAcknowledge{}
		case ProcedureCodeInitialContextSetup:
			target_msg = &InitialContextSetupResponse{}
		case ProcedureCodeMulticastSessionActivation:
			target_msg = &MulticastSessionActivationResponse{}
		case ProcedureCodeMulticastSessionDeactivation:
			target_msg = &MulticastSessionDeactivationResponse{}
		case ProcedureCodeMulticastSessionUpdate:
			target_msg = &MulticastSessionUpdateResponse{}
		case ProcedureCodeNGReset:
			target_msg = &NGResetAcknowledge{}
		case ProcedureCodeNGSetup:
			target_msg = &NGSetupResponse{}
		case ProcedureCodePathSwitchRequest:
			target_msg = &PathSwitchRequestAcknowledge{}
		case ProcedureCodePDUSessionResourceModify:
			target_msg = &PDUSessionResourceModifyResponse{}
		case ProcedureCodePDUSessionResourceModifyIndication:
			target_msg = &PDUSessionResourceModifyConfirm{}
		case ProcedureCodePDUSessionResourceRelease:
			target_msg = &PDUSessionResourceReleaseResponse{}
		case ProcedureCodePDUSessionResourceSetup:
			target_msg = &PDUSessionResourceSetupResponse{}
		case ProcedureCodePWSCancel:
			target_msg = &PWSCancelResponse{}
		case ProcedureCodeRANConfigurationUpdate:
			target_msg = &RANConfigurationUpdateAcknowledge{}
		case ProcedureCodeUEContextModification:
			target_msg = &UEContextModificationResponse{}
		case ProcedureCodeUEContextRelease:
			target_msg = &UEContextReleaseComplete{}
		case ProcedureCodeUEContextResume:
			target_msg = &UEContextResumeResponse{}
		case ProcedureCodeUEContextSuspend:
			target_msg = &UEContextSuspendResponse{}
		case ProcedureCodeUERadioCapabilityCheck:
			target_msg = &UERadioCapabilityCheckResponse{}
		case ProcedureCodeUERadioCapabilityIDMapping:
			target_msg = &UERadioCapabilityIDMappingResponse{}
		case ProcedureCodeWriteReplaceWarning:
			target_msg = &WriteReplaceWarningResponse{}

		default:
			// ProcedureCode IE is in the ie group MessageType (use MessageType criticality)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdProcedureCode, ngappduCrit)
			errTrace := errors.Errorf("Unknown Procedure Code: %d", ngappduProcedureCode)
			err = ie.BuildAbstractSyntaxErr(
				ngappduProcedureCode,
				aper.Enumerated(MessageTypeSuccessfulOutcome),
				ngappduCrit,
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	case MessageTypeUnsuccessfulOutcome:
		switch ngappduProcedureCode {
		case ProcedureCodeAMFConfigurationUpdate:
			target_msg = &AMFConfigurationUpdateFailure{}
		case ProcedureCodeBroadcastSessionModification:
			target_msg = &BroadcastSessionModificationFailure{}
		case ProcedureCodeBroadcastSessionSetup:
			target_msg = &BroadcastSessionSetupFailure{}
		case ProcedureCodeDistributionSetup:
			target_msg = &DistributionSetupFailure{}
		case ProcedureCodeHandoverPreparation:
			target_msg = &HandoverPreparationFailure{}
		case ProcedureCodeHandoverResourceAllocation:
			target_msg = &HandoverFailure{}
		case ProcedureCodeInitialContextSetup:
			target_msg = &InitialContextSetupFailure{}
		case ProcedureCodeMulticastSessionActivation:
			target_msg = &MulticastSessionActivationFailure{}
		case ProcedureCodeMulticastSessionUpdate:
			target_msg = &MulticastSessionUpdateFailure{}
		case ProcedureCodeNGSetup:
			target_msg = &NGSetupFailure{}
		case ProcedureCodePathSwitchRequest:
			target_msg = &PathSwitchRequestFailure{}
		case ProcedureCodeRANConfigurationUpdate:
			target_msg = &RANConfigurationUpdateFailure{}
		case ProcedureCodeUEContextModification:
			target_msg = &UEContextModificationFailure{}
		case ProcedureCodeUEContextResume:
			target_msg = &UEContextResumeFailure{}
		case ProcedureCodeUEContextSuspend:
			target_msg = &UEContextSuspendFailure{}

		default:
			// ProcedureCode IE is in the ie group MessageType (use MessageType criticality)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdProcedureCode, ngappduCrit)
			errTrace := errors.Errorf("Unknown Procedure Code: %d", ngappduProcedureCode)
			err = ie.BuildAbstractSyntaxErr(
				ngappduProcedureCode,
				aper.Enumerated(MessageTypeUnsuccessfulOutcome),
				ngappduCrit,
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	// REPLACE END
	default:
		// TypeOfMessage IE is in the ie group MessageType (use MessageType criticality)
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdMessageType, ngappduCrit)
		errTrace := errors.Errorf("Unknown Message Type: %d", ngappduType)
		err = ie.BuildAbstractSyntaxErr(
			ngappduProcedureCode,
			aper.Enumerated(ngappduType),
			ngappduCrit,
			&ie.AbstractSyntaxErrNotComprehendedIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if err != nil {
		return target_msg, marshalled, errors.Wrap(err, "Parse Message Type error")
	}

	// get left bytes from processed pd
	pdBitsOffset := pd.BitOffset()
	pdByteOffset := pd.ByteOffset()
	// left bytes should only contain encoded open type value
	// get the marshalled open type value in bytes (should start from new byte)
	if pdBitsOffset != 0 {
		pdByteOffset += 1
	}
	leftBytes := pd.Bytes()[pdByteOffset:]

	return target_msg, leftBytes, nil
}

func Parse(marshalled []byte) (Message, error) {
	target_msg, left_bytes, err := ParseMessageType(marshalled)
	if err != nil {
		return target_msg, err
	}

	err = target_msg.UnmarshalBinary(left_bytes)
	return target_msg, err
}
