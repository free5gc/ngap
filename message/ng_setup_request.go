package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

type NGSetupRequest struct {
	GlobalRANNodeID        *ie.GlobalRANNodeID        // mandatory
	RANNodeName            *ie.RANNodeName            // optional
	SupportedTAList        *ie.SupportedTAList        // mandatory
	DefaultPagingDRX       *ie.PagingDRX              // mandatory
	UERetentionInformation *ie.UERetentionInformation // optional
	NBIoTDefaultPagingDRX  *ie.NBIoTDefaultPagingDRX  // optional
	ExtendedRANNodeName    *ie.ExtendedRANNodeName    // optional
}

func (x *NGSetupRequest) MessageType() int64 {
	return MessageTypeInitiatingMessage
}

func (x *NGSetupRequest) ProcedureCode() int64 {
	return ProcedureCodeNGSetup
}

func (x *NGSetupRequest) Criticality() aper.Enumerated {
	return CriticalityReject
}

func (x *NGSetupRequest) MarshalBinary() ([]byte, error) {
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

	// open type value: NGSetupRequest (SEQUENCE)
	optPresentFlag = []bool{} // no optional field
	err = pdOpenType.WriteSequencePreambleBitMap(optPresentFlag, true)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// count number of ies as SEQUENCE OF SIZE
	numIes := uint64(0)
	if x.GlobalRANNodeID != nil {
		numIes++
	}
	if x.RANNodeName != nil {
		numIes++
	}
	if x.SupportedTAList != nil {
		numIes++
	}
	if x.DefaultPagingDRX != nil {
		numIes++
	}
	if x.UERetentionInformation != nil {
		numIes++
	}
	if x.NBIoTDefaultPagingDRX != nil {
		numIes++
	}
	if x.ExtendedRANNodeName != nil {
		numIes++
	}
	err = pdOpenType.WriteSequenceOfPreambleBitMap(numIes, false, sLb, sUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// encode if the IE field is present

	// IE Field 1 (mandatory)
	if x.GlobalRANNodeID != nil {
		err = x.GlobalRANNodeID.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDGlobalRANNodeID},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field GlobalRANNodeID is missing")
	}

	// IE Field 2 (optional)
	if x.RANNodeName != nil {
		err = x.RANNodeName.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDRANNodeName},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 3 (mandatory)
	if x.SupportedTAList != nil {
		err = x.SupportedTAList.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDSupportedTAList},
			ie.ProtocolIECriticality{Value: ie.CriticalityReject})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field SupportedTAList is missing")
	}

	// IE Field 4 (mandatory)
	if x.DefaultPagingDRX != nil {
		err = x.DefaultPagingDRX.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDDefaultPagingDRX},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field DefaultPagingDRX is missing")
	}

	// IE Field 5 (optional)
	if x.UERetentionInformation != nil {
		err = x.UERetentionInformation.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUERetentionInformation},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 6 (optional)
	if x.NBIoTDefaultPagingDRX != nil {
		err = x.NBIoTDefaultPagingDRX.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNBIoTDefaultPagingDRX},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (optional)
	if x.ExtendedRANNodeName != nil {
		err = x.ExtendedRANNodeName.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDExtendedRANNodeName},
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

func (x *NGSetupRequest) UnmarshalBinary(marshalled []byte) error {
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

	// open type value: NGSetupRequest (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (NGSetupRequest) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDGlobalRANNodeID {
			// check if ie is duplicated
			if x.GlobalRANNodeID != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: GlobalRANNodeID")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.GlobalRANNodeID = &ie.GlobalRANNodeID{}
			err = x.GlobalRANNodeID.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 2
		if protocolIeId == ie.ProtocolIEIDRANNodeName {
			// check if ie is duplicated
			if x.RANNodeName != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: RANNodeName")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.RANNodeName = &ie.RANNodeName{}
			err = x.RANNodeName.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
		if protocolIeId == ie.ProtocolIEIDSupportedTAList {
			// check if ie is duplicated
			if x.SupportedTAList != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: SupportedTAList")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.SupportedTAList = &ie.SupportedTAList{}
			err = x.SupportedTAList.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
		if protocolIeId == ie.ProtocolIEIDDefaultPagingDRX {
			// check if ie is duplicated
			if x.DefaultPagingDRX != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: DefaultPagingDRX")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.DefaultPagingDRX = &ie.PagingDRX{}
			err = x.DefaultPagingDRX.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDUERetentionInformation {
			// check if ie is duplicated
			if x.UERetentionInformation != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UERetentionInformation")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UERetentionInformation = &ie.UERetentionInformation{}
			err = x.UERetentionInformation.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDNBIoTDefaultPagingDRX {
			// check if ie is duplicated
			if x.NBIoTDefaultPagingDRX != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NBIoTDefaultPagingDRX")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NBIoTDefaultPagingDRX = &ie.NBIoTDefaultPagingDRX{}
			err = x.NBIoTDefaultPagingDRX.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDExtendedRANNodeName {
			// check if ie is duplicated
			if x.ExtendedRANNodeName != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: ExtendedRANNodeName")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.ExtendedRANNodeName = &ie.ExtendedRANNodeName{}
			err = x.ExtendedRANNodeName.ReadIE(pdOpenType)
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
	if x.GlobalRANNodeID == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDGlobalRANNodeID, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE GlobalRANNodeID is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	if x.SupportedTAList == nil {
		// Missing Mandatory IE (reject)
		// Build Abstract Syntax Error (missing IE/IE Group)
		// TODO: support reporting more than one abstract syntax error
		reportIe := ie.BuildAbstractSyntaxErrReportIe(ie.ProtocolIEIDSupportedTAList, ie.CriticalityReject)
		errTrace := errors.Errorf("mandatory IE SupportedTAList is missing")
		return ie.BuildAbstractSyntaxErr(
			x.ProcedureCode(),
			aper.Enumerated(x.MessageType()),
			x.Criticality(),
			&ie.AbstractSyntaxErrMissingIE{
				ReportIe: reportIe,
			},
			errTrace)
	}

	// DefaultPagingDRX is mandatory but may be nil (ignored)

	return nil
}
