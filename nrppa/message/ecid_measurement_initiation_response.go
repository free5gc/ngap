package message

import (
	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/nrppa/ie"
	"github.com/pkg/errors"
)

type ECIDMeasurementInitiationResponse struct {
	NRPPATransactionID        int64                         // mandatory
	LMFUEMeasurementID        *ie.UEMeasurementID           // mandatory
	RANUEMeasurementID        *ie.UEMeasurementID           // mandatory
	ECIDMeasurementResult     *ie.ECIDMeasurementResult     // optional
	CriticalityDiagnostics    *ie.CriticalityDiagnostics    // optional
	CellPortionID             *ie.CellPortionID             // optional
	OtherRATMeasurementResult *ie.OtherRATMeasurementResult // optional
	WLANMeasurementResult     *ie.WLANMeasurementResult     // optional,
}

func (x *ECIDMeasurementInitiationResponse) MessageType() int64 {
	return MessageTypeSuccessfulOutcome
}

func (x *ECIDMeasurementInitiationResponse) ProcedureCode() int64 {
	return ProcedureCodeECIDMeasurementInitiation
}

func (x *ECIDMeasurementInitiationResponse) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *ECIDMeasurementInitiationResponse) MarshalBinary() ([]byte, error) {
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
	// sequence element: NRPPATransactionID
	*vLb, *vUb = 0, 32767
	err = pd.WriteInteger(x.NRPPATransactionID, false, vLb, vUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// (#) sequence element: Value (Open Type)
	pdOpenType := aper.NewPerBitData(nil)

	// open type value: ECIDMeasurementInitiationResponse (SEQUENCE)
	optPresentFlag = []bool{} // no optional field
	err = pdOpenType.WriteSequencePreambleBitMap(optPresentFlag, true)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// count number of ies as SEQUENCE OF SIZE
	numIes := uint64(0)
	if x.LMFUEMeasurementID != nil {
		numIes++
	}
	if x.RANUEMeasurementID != nil {
		numIes++
	}
	if x.ECIDMeasurementResult != nil {
		numIes++
	}
	if x.CriticalityDiagnostics != nil {
		numIes++
	}
	if x.CellPortionID != nil {
		numIes++
	}
	if x.OtherRATMeasurementResult != nil {
		numIes++
	}
	if x.WLANMeasurementResult != nil {
		numIes++
	}
	err = pdOpenType.WriteSequenceOfPreambleBitMap(numIes, false, sLb, sUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// encode if the IE field is present

	// IE Field 1 (mandatory)
	if x.LMFUEMeasurementID != nil {
		err = x.LMFUEMeasurementID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDLMFUEMeasurementID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field LMFUEMeasurementID is missing")
	}

	// IE Field 2 (mandatory)
	if x.RANUEMeasurementID != nil {
		err = x.RANUEMeasurementID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRANUEMeasurementID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field RANUEMeasurementID is missing")
	}

	// IE Field 3 (optional)
	if x.ECIDMeasurementResult != nil {
		err = x.ECIDMeasurementResult.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDECIDMeasurementResult},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 4 (optional)
	if x.CriticalityDiagnostics != nil {
		err = x.CriticalityDiagnostics.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCriticalityDiagnostics},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 5 (optional)
	if x.CellPortionID != nil {
		err = x.CellPortionID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCellPortionID},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 6 (optional)
	if x.OtherRATMeasurementResult != nil {
		err = x.OtherRATMeasurementResult.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDOtherRATMeasurementResult},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (optional,)
	if x.WLANMeasurementResult != nil {
		err = x.WLANMeasurementResult.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDWLANMeasurementResult},
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

func (x *ECIDMeasurementInitiationResponse) UnmarshalBinary(marshalled []byte) error {
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
	// sequence element: NRPPATransactionID

	// (#) sequence element: Value (Open Type)
	// Read Open Type byte
	var bytes []byte
	bytes, err = pd.ReadOpenType()
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode message Value (open-type) error"))
	}
	pdOpenType := aper.NewPerBitData(bytes)

	// open type value: ECIDMeasurementInitiationResponse (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (ECIDMeasurementInitiationResponse) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDLMFUEMeasurementID {
			// check if ie is duplicated
			if x.LMFUEMeasurementID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: LMFUEMeasurementID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.LMFUEMeasurementID = &ie.UEMeasurementID{}
			err = x.LMFUEMeasurementID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 2
		if protocolIeId == ie.ProtocolIEIDRANUEMeasurementID {
			// check if ie is duplicated
			if x.RANUEMeasurementID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RANUEMeasurementID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RANUEMeasurementID = &ie.UEMeasurementID{}
			err = x.RANUEMeasurementID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
		if protocolIeId == ie.ProtocolIEIDECIDMeasurementResult {
			// check if ie is duplicated
			if x.ECIDMeasurementResult != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ECIDMeasurementResult")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ECIDMeasurementResult = &ie.ECIDMeasurementResult{}
			err = x.ECIDMeasurementResult.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
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

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDCellPortionID {
			// check if ie is duplicated
			if x.CellPortionID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: CellPortionID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.CellPortionID = &ie.CellPortionID{}
			err = x.CellPortionID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDOtherRATMeasurementResult {
			// check if ie is duplicated
			if x.OtherRATMeasurementResult != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: OtherRATMeasurementResult")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.OtherRATMeasurementResult = &ie.OtherRATMeasurementResult{}
			err = x.OtherRATMeasurementResult.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDWLANMeasurementResult {
			// check if ie is duplicated
			if x.WLANMeasurementResult != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: WLANMeasurementResult")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.WLANMeasurementResult = &ie.WLANMeasurementResult{}
			err = x.WLANMeasurementResult.ReadIE(pdOpenType)
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
	if x.LMFUEMeasurementID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDLMFUEMeasurementID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE LMFUEMeasurementID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.RANUEMeasurementID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDRANUEMeasurementID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE RANUEMeasurementID is missing")
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
