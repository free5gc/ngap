package message

import (
	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/nrppa/ie"
	"github.com/pkg/errors"
)

type ECIDMeasurementInitiationRequest struct {
	NRPPATransactionID            int64                             // mandatory
	LMFUEMeasurementID            *ie.UEMeasurementID               // mandatory
	ReportCharacteristics         *ie.ReportCharacteristics         // mandatory
	MeasurementPeriodicity        *ie.MeasurementPeriodicity        // conditional
	MeasurementQuantities         *ie.MeasurementQuantities         // mandatory
	OtherRATMeasurementQuantities *ie.OtherRATMeasurementQuantities // optional
	WLANMeasurementQuantities     *ie.WLANMeasurementQuantities     // optional
	MeasurementPeriodicityNRAoA   *ie.MeasurementPeriodicityNRAoA   // conditional,
}

func (x *ECIDMeasurementInitiationRequest) MessageType() int64 {
	return MessageTypeInitiatingMessage
}

func (x *ECIDMeasurementInitiationRequest) ProcedureCode() int64 {
	return ProcedureCodeECIDMeasurementInitiation
}

func (x *ECIDMeasurementInitiationRequest) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *ECIDMeasurementInitiationRequest) MarshalBinary() ([]byte, error) {
	pd := aper.NewPerBitData(nil)
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// encode ngappdu (CHOICE)
	*vUb = 2
	err = pd.WriteChoicePreambleBitMap(MessageTypeInitiatingMessage, true, vUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}

	// encode MessageTypeInitiatingMessage (SEQUENCE)
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

	// open type value: ECIDMeasurementInitiationRequest (SEQUENCE)
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
	if x.ReportCharacteristics != nil {
		numIes++
	}
	if x.MeasurementPeriodicity != nil {
		numIes++
	}
	if x.MeasurementQuantities != nil {
		numIes++
	}
	if x.OtherRATMeasurementQuantities != nil {
		numIes++
	}
	if x.WLANMeasurementQuantities != nil {
		numIes++
	}
	if x.MeasurementPeriodicityNRAoA != nil {
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
	if x.ReportCharacteristics != nil {
		err = x.ReportCharacteristics.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDReportCharacteristics},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field ReportCharacteristics is missing")
	}

	// IE Field 3 (conditional)
	if x.MeasurementPeriodicity != nil {
		err = x.MeasurementPeriodicity.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementPeriodicity},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 4 (mandatory)
	if x.MeasurementQuantities != nil {
		err = x.MeasurementQuantities.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementQuantities},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field MeasurementQuantities is missing")
	}

	// IE Field 5 (optional)
	if x.OtherRATMeasurementQuantities != nil {
		err = x.OtherRATMeasurementQuantities.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDOtherRATMeasurementQuantities},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 6 (optional)
	if x.WLANMeasurementQuantities != nil {
		err = x.WLANMeasurementQuantities.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDWLANMeasurementQuantities},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (conditional,)
	if x.MeasurementPeriodicityNRAoA != nil {
		err = x.MeasurementPeriodicityNRAoA.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementPeriodicityNRAoA},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// Finish MessageTypeInitiatingMessage open type Value Encoding (#)
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}

	return pd.Bytes(), nil
}

func (x *ECIDMeasurementInitiationRequest) UnmarshalBinary(marshalled []byte) error {
	pd := aper.NewPerBitData(marshalled)
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)
	foo(err, sLb, sUb, vLb, vUb)

	// Some fields are decoded already:
	// decode ngappdu (CHOICE)
	// decode MessageTypeInitiatingMessage (SEQUENCE)
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

	// open type value: ECIDMeasurementInitiationRequest (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (ECIDMeasurementInitiationRequest) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDReportCharacteristics {
			// check if ie is duplicated
			if x.ReportCharacteristics != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ReportCharacteristics")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ReportCharacteristics = &ie.ReportCharacteristics{}
			err = x.ReportCharacteristics.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
		if protocolIeId == ie.ProtocolIEIDMeasurementPeriodicity {
			// check if ie is duplicated
			if x.MeasurementPeriodicity != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementPeriodicity")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementPeriodicity = &ie.MeasurementPeriodicity{}
			err = x.MeasurementPeriodicity.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
		if protocolIeId == ie.ProtocolIEIDMeasurementQuantities {
			// check if ie is duplicated
			if x.MeasurementQuantities != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementQuantities")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementQuantities = &ie.MeasurementQuantities{}
			err = x.MeasurementQuantities.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDOtherRATMeasurementQuantities {
			// check if ie is duplicated
			if x.OtherRATMeasurementQuantities != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: OtherRATMeasurementQuantities")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.OtherRATMeasurementQuantities = &ie.OtherRATMeasurementQuantities{}
			err = x.OtherRATMeasurementQuantities.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDWLANMeasurementQuantities {
			// check if ie is duplicated
			if x.WLANMeasurementQuantities != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: WLANMeasurementQuantities")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.WLANMeasurementQuantities = &ie.WLANMeasurementQuantities{}
			err = x.WLANMeasurementQuantities.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDMeasurementPeriodicityNRAoA {
			// check if ie is duplicated
			if x.MeasurementPeriodicityNRAoA != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementPeriodicityNRAoA")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementPeriodicityNRAoA = &ie.MeasurementPeriodicityNRAoA{}
			err = x.MeasurementPeriodicityNRAoA.ReadIE(pdOpenType)
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

	if x.ReportCharacteristics == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDReportCharacteristics, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE ReportCharacteristics is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.MeasurementQuantities == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDMeasurementQuantities, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE MeasurementQuantities is missing")
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
