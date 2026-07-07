package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

type HandoverCommand struct {
	AMFUENGAPID                          *ie.AMFUENGAPID                          // mandatory
	RANUENGAPID                          *ie.RANUENGAPID                          // mandatory
	HandoverType                         *ie.HandoverType                         // mandatory
	NASSecurityParametersFromNGRAN       *ie.NASSecurityParametersFromNGRAN       // conditional
	PDUSessionResourceHandoverList       *ie.PDUSessionResourceHandoverList       // optional
	PDUSessionResourceToReleaseListHOCmd *ie.PDUSessionResourceToReleaseListHOCmd // optional
	TargetToSourceTransparentContainer   *ie.TargetToSourceTransparentContainer   // mandatory
	CriticalityDiagnostics               *ie.CriticalityDiagnostics               // optional
}

func (x *HandoverCommand) MessageType() int64 {
	return MessageTypeSuccessfulOutcome
}

func (x *HandoverCommand) ProcedureCode() int64 {
	return ProcedureCodeHandoverPreparation
}

func (x *HandoverCommand) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *HandoverCommand) MarshalBinary() ([]byte, error) {
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

	// open type value: HandoverCommand (SEQUENCE)
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
	if x.HandoverType != nil {
		numIes++
	}
	if x.NASSecurityParametersFromNGRAN != nil {
		numIes++
	}
	if x.PDUSessionResourceHandoverList != nil {
		numIes++
	}
	if x.PDUSessionResourceToReleaseListHOCmd != nil {
		numIes++
	}
	if x.TargetToSourceTransparentContainer != nil {
		numIes++
	}
	if x.CriticalityDiagnostics != nil {
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
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
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
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field RANUENGAPID is missing")
	}

	// IE Field 3 (mandatory)
	if x.HandoverType != nil {
		err = x.HandoverType.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDHandoverType},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field HandoverType is missing")
	}

	// IE Field 4 (conditional)
	// This IE shall be present if HandoverType IE is set to value "5GStoEPPS" or “5GStoUTRAN”
	if x.HandoverType.Value == ie.HandoverTypePresentFivegsToEps || x.HandoverType.Value == ie.HandoverTypePresentFivegsToUtran {
		if x.NASSecurityParametersFromNGRAN != nil {
			err = x.NASSecurityParametersFromNGRAN.WriteIE(pdOpenType,
				ie.ProtocolIEID{Value: ie.ProtocolIEIDNASSecurityParametersFromNGRAN},
				ie.ProtocolIECriticality{Value: ie.CriticalityReject})
			if err != nil {
				return pd.Bytes(), errors.Wrap(err, "message marshal failed")
			}
		} else {
			return pd.Bytes(), errors.Errorf("NASSecurityParametersFromNGRAN shall be present if HandoverType IE is set to value \"5GStoEPPS\" or \"5GStoUTRAN\"")
		}
	} else {
		if x.NASSecurityParametersFromNGRAN != nil {
			return pd.Bytes(), errors.Errorf("NASSecurityParametersFromNGRAN shall not be present if HandoverType IE is not set to value \"5GStoEPPS\" nor \"5GStoUTRAN\"")
		}
	}

	// IE Field 5 (optional)
	if x.PDUSessionResourceHandoverList != nil {
		err = x.PDUSessionResourceHandoverList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPDUSessionResourceHandoverList},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 6 (optional)
	if x.PDUSessionResourceToReleaseListHOCmd != nil {
		err = x.PDUSessionResourceToReleaseListHOCmd.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (mandatory)
	if x.TargetToSourceTransparentContainer != nil {
		err = x.TargetToSourceTransparentContainer.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDTargetToSourceTransparentContainer},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field TargetToSourceTransparentContainer is missing")
	}

	// IE Field 8 (optional)
	if x.CriticalityDiagnostics != nil {
		err = x.CriticalityDiagnostics.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCriticalityDiagnostics},
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

func (x *HandoverCommand) UnmarshalBinary(marshalled []byte) error {
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

	// open type value: HandoverCommand (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (HandoverCommand) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDHandoverType {
			// check if ie is duplicated
			if x.HandoverType != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: HandoverType")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.HandoverType = &ie.HandoverType{}
			err = x.HandoverType.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
		if protocolIeId == ie.ProtocolIEIDNASSecurityParametersFromNGRAN {
			// check if ie is duplicated
			if x.NASSecurityParametersFromNGRAN != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NASSecurityParametersFromNGRAN")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NASSecurityParametersFromNGRAN = &ie.NASSecurityParametersFromNGRAN{}
			err = x.NASSecurityParametersFromNGRAN.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDPDUSessionResourceHandoverList {
			// check if ie is duplicated
			if x.PDUSessionResourceHandoverList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PDUSessionResourceHandoverList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PDUSessionResourceHandoverList = &ie.PDUSessionResourceHandoverList{}
			err = x.PDUSessionResourceHandoverList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDPDUSessionResourceToReleaseListHOCmd {
			// check if ie is duplicated
			if x.PDUSessionResourceToReleaseListHOCmd != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PDUSessionResourceToReleaseListHOCmd")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PDUSessionResourceToReleaseListHOCmd = &ie.PDUSessionResourceToReleaseListHOCmd{}
			err = x.PDUSessionResourceToReleaseListHOCmd.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDTargetToSourceTransparentContainer {
			// check if ie is duplicated
			if x.TargetToSourceTransparentContainer != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: TargetToSourceTransparentContainer")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.TargetToSourceTransparentContainer = &ie.TargetToSourceTransparentContainer{}
			err = x.TargetToSourceTransparentContainer.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 8
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
	if x.AMFUENGAPID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDAMFUENGAPID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE AMFUENGAPID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.RANUENGAPID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDRANUENGAPID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE RANUENGAPID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.HandoverType == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDHandoverType, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE HandoverType is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.TargetToSourceTransparentContainer == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDTargetToSourceTransparentContainer, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE TargetToSourceTransparentContainer is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	// Conditional field
	if x.HandoverType.Value == ie.HandoverTypePresentFivegsToEps || x.HandoverType.Value == ie.HandoverTypePresentFivegsToUtran {
		if x.NASSecurityParametersFromNGRAN == nil {
			// Build Abstract Syntax Error (missing IE/IE Group)
			reportIe := ie.BuildAbstractSyntaxErrReportIe(
				ie.ProtocolIEIDNASSecurityParametersFromNGRAN, ie.CriticalityReject)
			errTrace := errors.Errorf("NASSecurityParametersFromNGRAN shall be present if HandoverType IE is set to value \"5GStoEPPS\" or \"5GStoUTRAN\"")
			return ie.BuildAbstractSyntaxErr(
				x.ProcedureCode(),
				aper.Enumerated(x.MessageType()),
				x.Criticality(),
				&ie.AbstractSyntaxErrMissingIE{
					ReportIe: reportIe,
				},
				errTrace)
		}
	} else {
		if x.NASSecurityParametersFromNGRAN != nil {
			// Build Abstract Syntax Error (erroneously present of conditional IEs or IE groups)
			errTrace := errors.Errorf("NASSecurityParametersFromNGRAN shall not be present if HandoverType IE is set to neither \"5GStoEPPS\" nor \"5GStoUTRAN\"")
			return ie.BuildAbstractSyntaxErr(
				x.ProcedureCode(),
				aper.Enumerated(x.MessageType()),
				x.Criticality(),
				&ie.AbstractSyntaxErrIEErrPresent{},
				errTrace)
		}
	}

	return nil
}
