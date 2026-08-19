package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/nrppa/ie"
)

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

	// decode nrppapdu (CHOICE)
	*vUb = 2
	nrppapduType, err := pd.ReadChoicePreambleBitMap(true, vUb)
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
	nrppapduProcedureCode, err := pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	// sequence element: Criticality
	*vLb, *vUb = 0, 2
	nrppapduCrit, err := pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	// sequence element: NRPPATransactionID
	*vLb, *vUb = 0, 32767
	nrppapduTransactionId, err := pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return target_msg, marshalled,
			ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// use post-process/gen_parse_message_type.py to generate the codes in swich nrppapduType{}
	switch nrppapduType {
	// REPLACE START
	case MessageTypeInitiatingMessage:
		switch nrppapduProcedureCode {
		case ProcedureCodeECIDMeasurementInitiation:
			target_msg = &ECIDMeasurementInitiationRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeECIDMeasurementFailureIndication:
			target_msg = &ECIDMeasurementFailureIndication{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeECIDMeasurementReport:
			target_msg = &ECIDMeasurementReport{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeECIDMeasurementTermination:
			target_msg = &ECIDMeasurementTerminationCommand{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeOTDOAInformationExchange:
			target_msg = &OTDOAInformationRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeAssistanceInformationControl:
			target_msg = &AssistanceInformationControl{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeAssistanceInformationFeedback:
			target_msg = &AssistanceInformationFeedback{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeErrorIndication:
			target_msg = &ErrorIndication{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePrivateMessage:
			target_msg = &PrivateMessage{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningInformationExchange:
			target_msg = &PositioningInformationRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningInformationUpdate:
			target_msg = &PositioningInformationUpdate{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurement:
			target_msg = &MeasurementRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementReport:
			target_msg = &MeasurementReport{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementUpdate:
			target_msg = &MeasurementUpdate{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementAbort:
			target_msg = &MeasurementAbort{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementFailureIndication:
			target_msg = &MeasurementFailureIndication{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeTRPInformationExchange:
			target_msg = &TRPInformationRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningActivation:
			target_msg = &PositioningActivationRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningDeactivation:
			target_msg = &PositioningDeactivation{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePRSConfigurationExchange:
			target_msg = &PRSConfigurationRequest{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementPreconfiguration:
			target_msg = &MeasurementPreconfigurationRequired{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementActivation:
			target_msg = &MeasurementActivation{
				NRPPATransactionID: nrppapduTransactionId,
			}

		default:
			// ProcedureCode IE is in the ie group MessageType (use MessageType criticality)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdProcedureCode, nrppapduCrit)
			errTrace := errors.Errorf("Unknown Procedure Code: %d", nrppapduProcedureCode)
			err = ie.BuildAbstractSyntaxErr(
				nrppapduProcedureCode,
				aper.Enumerated(MessageTypeInitiatingMessage),
				nrppapduCrit,
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	case MessageTypeSuccessfulOutcome:
		switch nrppapduProcedureCode {
		case ProcedureCodeECIDMeasurementInitiation:
			target_msg = &ECIDMeasurementInitiationResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeOTDOAInformationExchange:
			target_msg = &OTDOAInformationResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningInformationExchange:
			target_msg = &PositioningInformationResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurement:
			target_msg = &MeasurementResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeTRPInformationExchange:
			target_msg = &TRPInformationResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningActivation:
			target_msg = &PositioningActivationResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePRSConfigurationExchange:
			target_msg = &PRSConfigurationResponse{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementPreconfiguration:
			target_msg = &MeasurementPreconfigurationConfirm{
				NRPPATransactionID: nrppapduTransactionId,
			}

		default:
			// ProcedureCode IE is in the ie group MessageType (use MessageType criticality)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdProcedureCode, nrppapduCrit)
			errTrace := errors.Errorf("Unknown Procedure Code: %d", nrppapduProcedureCode)
			err = ie.BuildAbstractSyntaxErr(
				nrppapduProcedureCode,
				aper.Enumerated(MessageTypeSuccessfulOutcome),
				nrppapduCrit,
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	case MessageTypeUnsuccessfulOutcome:
		switch nrppapduProcedureCode {
		case ProcedureCodeECIDMeasurementInitiation:
			target_msg = &ECIDMeasurementInitiationFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeOTDOAInformationExchange:
			target_msg = &OTDOAInformationFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningInformationExchange:
			target_msg = &PositioningInformationFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurement:
			target_msg = &MeasurementFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeTRPInformationExchange:
			target_msg = &TRPInformationFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePositioningActivation:
			target_msg = &PositioningActivationFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodePRSConfigurationExchange:
			target_msg = &PRSConfigurationFailure{
				NRPPATransactionID: nrppapduTransactionId,
			}
		case ProcedureCodeMeasurementPreconfiguration:
			target_msg = &MeasurementPreconfigurationRefuse{
				NRPPATransactionID: nrppapduTransactionId,
			}

		default:
			// ProcedureCode IE is in the ie group MessageType (use MessageType criticality)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdProcedureCode, nrppapduCrit)
			errTrace := errors.Errorf("Unknown Procedure Code: %d", nrppapduProcedureCode)
			err = ie.BuildAbstractSyntaxErr(
				nrppapduProcedureCode,
				aper.Enumerated(MessageTypeUnsuccessfulOutcome),
				nrppapduCrit,
				&ie.AbstractSyntaxErrNotComprehendedIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	// REPLACE END
	default:
		// TypeOfMessage IE is in the ie group MessageType (use MessageType criticality)
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.AbstractSyntaxErrReportIeIdMessageType, nrppapduCrit)
		errTrace := errors.Errorf("Unknown Message Type: %d", nrppapduType)
		err = ie.BuildAbstractSyntaxErr(
			nrppapduProcedureCode,
			aper.Enumerated(nrppapduType),
			nrppapduCrit,
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
