package message

import (
	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/nrppa/ie"
	"github.com/pkg/errors"
)

type MeasurementRequest struct {
	NRPPATransactionID                         int64                                          // mandatory
	LMFMeasurementID                           *ie.MeasurementID                              // mandatory
	TRPMeasurementRequestList                  *ie.TRPMeasurementRequestList                  // mandatory
	ReportCharacteristics                      *ie.ReportCharacteristics                      // mandatory
	MeasurementPeriodicity                     *ie.MeasurementPeriodicity                     // conditional
	TRPMeasurementQuantities                   *ie.TRPMeasurementQuantities                   // mandatory
	SFNInitialisationTime                      *ie.RelativeTime1900                           // optional
	SRSConfiguration                           *ie.SRSConfiguration                           // optional
	MeasurementBeamInfoRequest                 *ie.MeasurementBeamInfoRequest                 // optional
	SystemFrameNumber                          *ie.SystemFrameNumber                          // optional
	SlotNumber                                 *ie.SlotNumber                                 // optional
	MeasurementPeriodicityExtended             *ie.MeasurementPeriodicityExtended             // conditional
	ResponseTime                               *ie.ResponseTime                               // optional
	MeasurementCharacteristicsRequestIndicator *ie.MeasurementCharacteristicsRequestIndicator // optional
	MeasurementTimeOccasion                    *ie.MeasurementTimeOccasion                    // optional
	MeasurementAmount                          *ie.MeasurementAmount                          // optional,
}

func (x *MeasurementRequest) MessageType() int64 {
	return MessageTypeInitiatingMessage
}

func (x *MeasurementRequest) ProcedureCode() int64 {
	return ProcedureCodeMeasurement
}

func (x *MeasurementRequest) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *MeasurementRequest) MarshalBinary() ([]byte, error) {
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

	// open type value: MeasurementRequest (SEQUENCE)
	optPresentFlag = []bool{} // no optional field
	err = pdOpenType.WriteSequencePreambleBitMap(optPresentFlag, true)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// count number of ies as SEQUENCE OF SIZE
	numIes := uint64(0)
	if x.LMFMeasurementID != nil {
		numIes++
	}
	if x.TRPMeasurementRequestList != nil {
		numIes++
	}
	if x.ReportCharacteristics != nil {
		numIes++
	}
	if x.MeasurementPeriodicity != nil {
		numIes++
	}
	if x.TRPMeasurementQuantities != nil {
		numIes++
	}
	if x.SFNInitialisationTime != nil {
		numIes++
	}
	if x.SRSConfiguration != nil {
		numIes++
	}
	if x.MeasurementBeamInfoRequest != nil {
		numIes++
	}
	if x.SystemFrameNumber != nil {
		numIes++
	}
	if x.SlotNumber != nil {
		numIes++
	}
	if x.MeasurementPeriodicityExtended != nil {
		numIes++
	}
	if x.ResponseTime != nil {
		numIes++
	}
	if x.MeasurementCharacteristicsRequestIndicator != nil {
		numIes++
	}
	if x.MeasurementTimeOccasion != nil {
		numIes++
	}
	if x.MeasurementAmount != nil {
		numIes++
	}
	err = pdOpenType.WriteSequenceOfPreambleBitMap(numIes, false, sLb, sUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// encode if the IE field is present

	// IE Field 1 (mandatory)
	if x.LMFMeasurementID != nil {
		err = x.LMFMeasurementID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDLMFMeasurementID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field LMFMeasurementID is missing")
	}

	// IE Field 2 (mandatory)
	if x.TRPMeasurementRequestList != nil {
		err = x.TRPMeasurementRequestList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDTRPMeasurementRequestList},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field TRPMeasurementRequestList is missing")
	}

	// IE Field 3 (mandatory)
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

	// IE Field 4 (conditional)
	if x.MeasurementPeriodicity != nil {
		err = x.MeasurementPeriodicity.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementPeriodicity},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 5 (mandatory)
	if x.TRPMeasurementQuantities != nil {
		err = x.TRPMeasurementQuantities.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDTRPMeasurementQuantities},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field TRPMeasurementQuantities is missing")
	}

	// IE Field 6 (optional)
	if x.SFNInitialisationTime != nil {
		err = x.SFNInitialisationTime.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSFNInitialisationTime},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (optional)
	if x.SRSConfiguration != nil {
		err = x.SRSConfiguration.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSRSConfiguration},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 8 (optional)
	if x.MeasurementBeamInfoRequest != nil {
		err = x.MeasurementBeamInfoRequest.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementBeamInfoRequest},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 9 (optional)
	if x.SystemFrameNumber != nil {
		err = x.SystemFrameNumber.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSystemFrameNumber},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 10 (optional)
	if x.SlotNumber != nil {
		err = x.SlotNumber.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSlotNumber},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 11 (conditional)
	if x.MeasurementPeriodicityExtended != nil {
		err = x.MeasurementPeriodicityExtended.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementPeriodicityExtended},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 12 (optional)
	if x.ResponseTime != nil {
		err = x.ResponseTime.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDResponseTime},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 13 (optional)
	if x.MeasurementCharacteristicsRequestIndicator != nil {
		err = x.MeasurementCharacteristicsRequestIndicator.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementCharacteristicsRequestIndicator},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 14 (optional)
	if x.MeasurementTimeOccasion != nil {
		err = x.MeasurementTimeOccasion.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementTimeOccasion},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 15 (optional,)
	if x.MeasurementAmount != nil {
		err = x.MeasurementAmount.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDMeasurementAmount},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
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

func (x *MeasurementRequest) UnmarshalBinary(marshalled []byte) error {
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

	// open type value: MeasurementRequest (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (MeasurementRequest) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDLMFMeasurementID {
			// check if ie is duplicated
			if x.LMFMeasurementID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: LMFMeasurementID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.LMFMeasurementID = &ie.MeasurementID{}
			err = x.LMFMeasurementID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 2
		if protocolIeId == ie.ProtocolIEIDTRPMeasurementRequestList {
			// check if ie is duplicated
			if x.TRPMeasurementRequestList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: TRPMeasurementRequestList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.TRPMeasurementRequestList = &ie.TRPMeasurementRequestList{}
			err = x.TRPMeasurementRequestList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
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

		// IE Field 4
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

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDTRPMeasurementQuantities {
			// check if ie is duplicated
			if x.TRPMeasurementQuantities != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: TRPMeasurementQuantities")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.TRPMeasurementQuantities = &ie.TRPMeasurementQuantities{}
			err = x.TRPMeasurementQuantities.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDSFNInitialisationTime {
			// check if ie is duplicated
			if x.SFNInitialisationTime != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SFNInitialisationTime")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SFNInitialisationTime = &ie.RelativeTime1900{}
			err = x.SFNInitialisationTime.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDSRSConfiguration {
			// check if ie is duplicated
			if x.SRSConfiguration != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SRSConfiguration")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SRSConfiguration = &ie.SRSConfiguration{}
			err = x.SRSConfiguration.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 8
		if protocolIeId == ie.ProtocolIEIDMeasurementBeamInfoRequest {
			// check if ie is duplicated
			if x.MeasurementBeamInfoRequest != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementBeamInfoRequest")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementBeamInfoRequest = &ie.MeasurementBeamInfoRequest{}
			err = x.MeasurementBeamInfoRequest.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 9
		if protocolIeId == ie.ProtocolIEIDSystemFrameNumber {
			// check if ie is duplicated
			if x.SystemFrameNumber != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SystemFrameNumber")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SystemFrameNumber = &ie.SystemFrameNumber{}
			err = x.SystemFrameNumber.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 10
		if protocolIeId == ie.ProtocolIEIDSlotNumber {
			// check if ie is duplicated
			if x.SlotNumber != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SlotNumber")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SlotNumber = &ie.SlotNumber{}
			err = x.SlotNumber.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 11
		if protocolIeId == ie.ProtocolIEIDMeasurementPeriodicityExtended {
			// check if ie is duplicated
			if x.MeasurementPeriodicityExtended != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementPeriodicityExtended")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementPeriodicityExtended = &ie.MeasurementPeriodicityExtended{}
			err = x.MeasurementPeriodicityExtended.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 12
		if protocolIeId == ie.ProtocolIEIDResponseTime {
			// check if ie is duplicated
			if x.ResponseTime != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ResponseTime")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ResponseTime = &ie.ResponseTime{}
			err = x.ResponseTime.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 13
		if protocolIeId == ie.ProtocolIEIDMeasurementCharacteristicsRequestIndicator {
			// check if ie is duplicated
			if x.MeasurementCharacteristicsRequestIndicator != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementCharacteristicsRequestIndicator")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementCharacteristicsRequestIndicator = &ie.MeasurementCharacteristicsRequestIndicator{}
			err = x.MeasurementCharacteristicsRequestIndicator.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 14
		if protocolIeId == ie.ProtocolIEIDMeasurementTimeOccasion {
			// check if ie is duplicated
			if x.MeasurementTimeOccasion != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementTimeOccasion")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementTimeOccasion = &ie.MeasurementTimeOccasion{}
			err = x.MeasurementTimeOccasion.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 15
		if protocolIeId == ie.ProtocolIEIDMeasurementAmount {
			// check if ie is duplicated
			if x.MeasurementAmount != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: MeasurementAmount")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.MeasurementAmount = &ie.MeasurementAmount{}
			err = x.MeasurementAmount.ReadIE(pdOpenType)
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
	if x.LMFMeasurementID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDLMFMeasurementID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE LMFMeasurementID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.TRPMeasurementRequestList == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDTRPMeasurementRequestList, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE TRPMeasurementRequestList is missing")
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

	if x.TRPMeasurementQuantities == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDTRPMeasurementQuantities, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE TRPMeasurementQuantities is missing")
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
