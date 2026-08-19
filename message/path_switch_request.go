package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

type PathSwitchRequest struct {
	RANUENGAPID                              *ie.RANUENGAPID                              // mandatory
	SourceAMFUENGAPID                        *ie.AMFUENGAPID                              // mandatory
	UserLocationInformation                  *ie.UserLocationInformation                  // mandatory
	UESecurityCapabilities                   *ie.UESecurityCapabilities                   // mandatory
	PDUSessionResourceToBeSwitchedDLList     *ie.PDUSessionResourceToBeSwitchedDLList     // mandatory
	PDUSessionResourceFailedToSetupListPSReq *ie.PDUSessionResourceFailedToSetupListPSReq // optional
	RRCResumeCause                           *ie.RRCEstablishmentCause                    // optional
	RedCapIndication                         *ie.RedCapIndication                         // optional
}

func (x *PathSwitchRequest) MessageType() int64 {
	return MessageTypeInitiatingMessage
}

func (x *PathSwitchRequest) ProcedureCode() int64 {
	return ProcedureCodePathSwitchRequest
}

func (x *PathSwitchRequest) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *PathSwitchRequest) MarshalBinary() ([]byte, error) {
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

	// open type value: PathSwitchRequest (SEQUENCE)
	optPresentFlag = []bool{} // no optional field
	err = pdOpenType.WriteSequencePreambleBitMap(optPresentFlag, true)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// count number of ies as SEQUENCE OF SIZE
	numIes := uint64(0)
	if x.RANUENGAPID != nil {
		numIes++
	}
	if x.SourceAMFUENGAPID != nil {
		numIes++
	}
	if x.UserLocationInformation != nil {
		numIes++
	}
	if x.UESecurityCapabilities != nil {
		numIes++
	}
	if x.PDUSessionResourceToBeSwitchedDLList != nil {
		numIes++
	}
	if x.PDUSessionResourceFailedToSetupListPSReq != nil {
		numIes++
	}
	if x.RRCResumeCause != nil {
		numIes++
	}
	if x.RedCapIndication != nil {
		numIes++
	}
	err = pdOpenType.WriteSequenceOfPreambleBitMap(numIes, false, sLb, sUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// encode if the IE field is present

	// IE Field 1 (mandatory)
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

	// IE Field 2 (mandatory)
	if x.SourceAMFUENGAPID != nil {
		err = x.SourceAMFUENGAPID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSourceAMFUENGAPID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field SourceAMFUENGAPID is missing")
	}

	// IE Field 3 (mandatory)
	if x.UserLocationInformation != nil {
		err = x.UserLocationInformation.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUserLocationInformation},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field UserLocationInformation is missing")
	}

	// IE Field 4 (mandatory)
	if x.UESecurityCapabilities != nil {
		err = x.UESecurityCapabilities.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUESecurityCapabilities},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field UESecurityCapabilities is missing")
	}

	// IE Field 5 (mandatory)
	if x.PDUSessionResourceToBeSwitchedDLList != nil {
		err = x.PDUSessionResourceToBeSwitchedDLList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field PDUSessionResourceToBeSwitchedDLList is missing")
	}

	// IE Field 6 (optional)
	if x.PDUSessionResourceFailedToSetupListPSReq != nil {
		err = x.PDUSessionResourceFailedToSetupListPSReq.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (optional)
	if x.RRCResumeCause != nil {
		err = x.RRCResumeCause.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRRCResumeCause},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 8 (optional)
	if x.RedCapIndication != nil {
		err = x.RedCapIndication.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRedCapIndication},
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

func (x *PathSwitchRequest) UnmarshalBinary(marshalled []byte) error {
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

	// open type value: PathSwitchRequest (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (PathSwitchRequest) sequence error"))
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

		// IE Field 2
		if protocolIeId == ie.ProtocolIEIDSourceAMFUENGAPID {
			// check if ie is duplicated
			if x.SourceAMFUENGAPID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SourceAMFUENGAPID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SourceAMFUENGAPID = &ie.AMFUENGAPID{}
			err = x.SourceAMFUENGAPID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
		if protocolIeId == ie.ProtocolIEIDUserLocationInformation {
			// check if ie is duplicated
			if x.UserLocationInformation != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UserLocationInformation")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UserLocationInformation = &ie.UserLocationInformation{}
			err = x.UserLocationInformation.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
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

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList {
			// check if ie is duplicated
			if x.PDUSessionResourceToBeSwitchedDLList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PDUSessionResourceToBeSwitchedDLList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PDUSessionResourceToBeSwitchedDLList = &ie.PDUSessionResourceToBeSwitchedDLList{}
			err = x.PDUSessionResourceToBeSwitchedDLList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq {
			// check if ie is duplicated
			if x.PDUSessionResourceFailedToSetupListPSReq != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PDUSessionResourceFailedToSetupListPSReq")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PDUSessionResourceFailedToSetupListPSReq = &ie.PDUSessionResourceFailedToSetupListPSReq{}
			err = x.PDUSessionResourceFailedToSetupListPSReq.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDRRCResumeCause {
			// check if ie is duplicated
			if x.RRCResumeCause != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RRCResumeCause")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RRCResumeCause = &ie.RRCEstablishmentCause{}
			err = x.RRCResumeCause.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 8
		if protocolIeId == ie.ProtocolIEIDRedCapIndication {
			// check if ie is duplicated
			if x.RedCapIndication != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RedCapIndication")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RedCapIndication = &ie.RedCapIndication{}
			err = x.RedCapIndication.ReadIE(pdOpenType)
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

	if x.SourceAMFUENGAPID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDSourceAMFUENGAPID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE SourceAMFUENGAPID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	// UserLocationInformation is mandatory but may be nil (ignored)

	// UESecurityCapabilities is mandatory but may be nil (ignored)

	if x.PDUSessionResourceToBeSwitchedDLList == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDPDUSessionResourceToBeSwitchedDLList, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE PDUSessionResourceToBeSwitchedDLList is missing")
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
