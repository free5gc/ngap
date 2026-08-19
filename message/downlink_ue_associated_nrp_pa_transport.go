package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

type DownlinkUEAssociatedNRPPaTransport struct {
	AMFUENGAPID *ie.AMFUENGAPID // mandatory
	RANUENGAPID *ie.RANUENGAPID // mandatory
	RoutingID   *ie.RoutingID   // mandatory
	NRPPaPDU    *ie.NRPPaPDU    // mandatory
}

func (x *DownlinkUEAssociatedNRPPaTransport) MessageType() int64 {
	return MessageTypeInitiatingMessage
}

func (x *DownlinkUEAssociatedNRPPaTransport) ProcedureCode() int64 {
	return ProcedureCodeDownlinkUEAssociatedNRPPaTransport
}

func (x *DownlinkUEAssociatedNRPPaTransport) Criticality() aper.Enumerated {
	return CriticalityIgnore
}

func (x *DownlinkUEAssociatedNRPPaTransport) MarshalBinary() ([]byte, error) {
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
	// (#) sequence element: Value (Open Type)
	pdOpenType := aper.NewPerBitData(nil)

	// open type value: DownlinkUEAssociatedNRPPaTransport (SEQUENCE)
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
	if x.RoutingID != nil {
		numIes++
	}
	if x.NRPPaPDU != nil {
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
	if x.RoutingID != nil {
		err = x.RoutingID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRoutingID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field RoutingID is missing")
	}

	// IE Field 4 (mandatory)
	if x.NRPPaPDU != nil {
		err = x.NRPPaPDU.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNRPPaPDU},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field NRPPaPDU is missing")
	}

	// Finish MessageTypeInitiatingMessage open type Value Encoding (#)
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}

	return pd.Bytes(), nil
}

func (x *DownlinkUEAssociatedNRPPaTransport) UnmarshalBinary(marshalled []byte) error {
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

	// (#) sequence element: Value (Open Type)
	// Read Open Type byte
	var bytes []byte
	bytes, err = pd.ReadOpenType()
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode message Value (open-type) error"))
	}
	pdOpenType := aper.NewPerBitData(bytes)

	// open type value: DownlinkUEAssociatedNRPPaTransport (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (DownlinkUEAssociatedNRPPaTransport) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDRoutingID {
			// check if ie is duplicated
			if x.RoutingID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RoutingID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RoutingID = &ie.RoutingID{}
			err = x.RoutingID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
		if protocolIeId == ie.ProtocolIEIDNRPPaPDU {
			// check if ie is duplicated
			if x.NRPPaPDU != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NRPPaPDU")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NRPPaPDU = &ie.NRPPaPDU{}
			err = x.NRPPaPDU.ReadIE(pdOpenType)
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

	if x.RoutingID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDRoutingID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE RoutingID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.NRPPaPDU == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDNRPPaPDU, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE NRPPaPDU is missing")
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
