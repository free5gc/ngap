package message

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

type Paging struct {
	UEPagingIdentity            *ie.UEPagingIdentity            // mandatory
	PagingDRX                   *ie.PagingDRX                   // optional
	TAIListForPaging            *ie.TAIListForPaging            // mandatory
	PagingPriority              *ie.PagingPriority              // optional
	UERadioCapabilityForPaging  *ie.UERadioCapabilityForPaging  // optional
	PagingOrigin                *ie.PagingOrigin                // optional
	AssistanceDataForPaging     *ie.AssistanceDataForPaging     // optional
	NBIoTPagingEDRXInfo         *ie.NBIoTPagingEDRXInfo         // optional
	NBIoTPagingDRX              *ie.NBIoTPagingDRX              // optional
	EnhancedCoverageRestriction *ie.EnhancedCoverageRestriction // optional
	WUSAssistanceInformation    *ie.WUSAssistanceInformation    // optional
	EUTRAPagingeDRXInformation  *ie.EUTRAPagingeDRXInformation  // optional
	CEmodeBrestricted           *ie.CEmodeBrestricted           // optional
	NRPagingeDRXInformation     *ie.NRPagingeDRXInformation     // optional
	PagingCause                 *ie.PagingCause                 // optional
	PEIPSassistanceInformation  *ie.PEIPSassistanceInformation  // optional
}

func (x *Paging) MessageType() int64 {
	return MessageTypeInitiatingMessage
}

func (x *Paging) ProcedureCode() int64 {
	return ProcedureCodePaging
}

func (x *Paging) Criticality() aper.Enumerated {
	return CriticalityIgnore
}

func (x *Paging) MarshalBinary() ([]byte, error) {
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

	// open type value: Paging (SEQUENCE)
	optPresentFlag = []bool{} // no optional field
	err = pdOpenType.WriteSequencePreambleBitMap(optPresentFlag, true)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence element: ProtocolIE-Container (SEQUENCE OF)
	*sLb, *sUb = 0, 65535
	// count number of ies as SEQUENCE OF SIZE
	numIes := uint64(0)
	if x.UEPagingIdentity != nil {
		numIes++
	}
	if x.PagingDRX != nil {
		numIes++
	}
	if x.TAIListForPaging != nil {
		numIes++
	}
	if x.PagingPriority != nil {
		numIes++
	}
	if x.UERadioCapabilityForPaging != nil {
		numIes++
	}
	if x.PagingOrigin != nil {
		numIes++
	}
	if x.AssistanceDataForPaging != nil {
		numIes++
	}
	if x.NBIoTPagingEDRXInfo != nil {
		numIes++
	}
	if x.NBIoTPagingDRX != nil {
		numIes++
	}
	if x.EnhancedCoverageRestriction != nil {
		numIes++
	}
	if x.WUSAssistanceInformation != nil {
		numIes++
	}
	if x.EUTRAPagingeDRXInformation != nil {
		numIes++
	}
	if x.CEmodeBrestricted != nil {
		numIes++
	}
	if x.NRPagingeDRXInformation != nil {
		numIes++
	}
	if x.PagingCause != nil {
		numIes++
	}
	if x.PEIPSassistanceInformation != nil {
		numIes++
	}
	err = pdOpenType.WriteSequenceOfPreambleBitMap(numIes, false, sLb, sUb)
	if err != nil {
		return pd.Bytes(), errors.Wrap(err, "message marshal failed")
	}
	// sequence of element:  ProtocolIE-Field (SEQUENCE)
	// encode if the IE field is present

	// IE Field 1 (mandatory)
	if x.UEPagingIdentity != nil {
		err = x.UEPagingIdentity.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUEPagingIdentity},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field UEPagingIdentity is missing")
	}

	// IE Field 2 (optional)
	if x.PagingDRX != nil {
		err = x.PagingDRX.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPagingDRX},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 3 (mandatory)
	if x.TAIListForPaging != nil {
		err = x.TAIListForPaging.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDTAIListForPaging},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	} else {
		return pd.Bytes(), errors.Errorf("mandatory field TAIListForPaging is missing")
	}

	// IE Field 4 (optional)
	if x.PagingPriority != nil {
		err = x.PagingPriority.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPagingPriority},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 5 (optional)
	if x.UERadioCapabilityForPaging != nil {
		err = x.UERadioCapabilityForPaging.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDUERadioCapabilityForPaging},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 6 (optional)
	if x.PagingOrigin != nil {
		err = x.PagingOrigin.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPagingOrigin},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 7 (optional)
	if x.AssistanceDataForPaging != nil {
		err = x.AssistanceDataForPaging.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDAssistanceDataForPaging},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 8 (optional)
	if x.NBIoTPagingEDRXInfo != nil {
		err = x.NBIoTPagingEDRXInfo.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNBIoTPagingEDRXInfo},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 9 (optional)
	if x.NBIoTPagingDRX != nil {
		err = x.NBIoTPagingDRX.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNBIoTPagingDRX},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 10 (optional)
	if x.EnhancedCoverageRestriction != nil {
		err = x.EnhancedCoverageRestriction.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDEnhancedCoverageRestriction},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 11 (optional)
	if x.WUSAssistanceInformation != nil {
		err = x.WUSAssistanceInformation.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDWUSAssistanceInformation},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 12 (optional)
	if x.EUTRAPagingeDRXInformation != nil {
		err = x.EUTRAPagingeDRXInformation.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDEUTRAPagingeDRXInformation},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 13 (optional)
	if x.CEmodeBrestricted != nil {
		err = x.CEmodeBrestricted.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDCEmodeBrestricted},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 14 (optional)
	if x.NRPagingeDRXInformation != nil {
		err = x.NRPagingeDRXInformation.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDNRPagingeDRXInformation},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 15 (optional)
	if x.PagingCause != nil {
		err = x.PagingCause.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPagingCause},
			ie.ProtocolIECriticality{Value: ie.CriticalityIgnore})
		if err != nil {
			return pd.Bytes(), errors.Wrap(err, "message marshal failed")
		}
	}

	// IE Field 16 (optional)
	if x.PEIPSassistanceInformation != nil {
		err = x.PEIPSassistanceInformation.WriteIE(pdOpenType,
			ie.ProtocolIEID{Value: ie.ProtocolIEIDPEIPSassistanceInformation},
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

func (x *Paging) UnmarshalBinary(marshalled []byte) error {
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

	// open type value: Paging (SEQUENCE)
	optPresentFlag := []bool{} // no optional field
	err = pdOpenType.ReadSequencePreambleBitMap(&optPresentFlag, true)
	if err != nil {
		return ie.BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode message Value (Paging) sequence error"))
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
		if protocolIeId == ie.ProtocolIEIDUEPagingIdentity {
			// check if ie is duplicated
			if x.UEPagingIdentity != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UEPagingIdentity")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UEPagingIdentity = &ie.UEPagingIdentity{}
			err = x.UEPagingIdentity.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 2
		if protocolIeId == ie.ProtocolIEIDPagingDRX {
			// check if ie is duplicated
			if x.PagingDRX != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PagingDRX")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PagingDRX = &ie.PagingDRX{}
			err = x.PagingDRX.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 3
		if protocolIeId == ie.ProtocolIEIDTAIListForPaging {
			// check if ie is duplicated
			if x.TAIListForPaging != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: TAIListForPaging")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.TAIListForPaging = &ie.TAIListForPaging{}
			err = x.TAIListForPaging.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 4
		if protocolIeId == ie.ProtocolIEIDPagingPriority {
			// check if ie is duplicated
			if x.PagingPriority != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PagingPriority")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PagingPriority = &ie.PagingPriority{}
			err = x.PagingPriority.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 5
		if protocolIeId == ie.ProtocolIEIDUERadioCapabilityForPaging {
			// check if ie is duplicated
			if x.UERadioCapabilityForPaging != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: UERadioCapabilityForPaging")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.UERadioCapabilityForPaging = &ie.UERadioCapabilityForPaging{}
			err = x.UERadioCapabilityForPaging.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 6
		if protocolIeId == ie.ProtocolIEIDPagingOrigin {
			// check if ie is duplicated
			if x.PagingOrigin != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PagingOrigin")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PagingOrigin = &ie.PagingOrigin{}
			err = x.PagingOrigin.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 7
		if protocolIeId == ie.ProtocolIEIDAssistanceDataForPaging {
			// check if ie is duplicated
			if x.AssistanceDataForPaging != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: AssistanceDataForPaging")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.AssistanceDataForPaging = &ie.AssistanceDataForPaging{}
			err = x.AssistanceDataForPaging.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 8
		if protocolIeId == ie.ProtocolIEIDNBIoTPagingEDRXInfo {
			// check if ie is duplicated
			if x.NBIoTPagingEDRXInfo != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NBIoTPagingEDRXInfo")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NBIoTPagingEDRXInfo = &ie.NBIoTPagingEDRXInfo{}
			err = x.NBIoTPagingEDRXInfo.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 9
		if protocolIeId == ie.ProtocolIEIDNBIoTPagingDRX {
			// check if ie is duplicated
			if x.NBIoTPagingDRX != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NBIoTPagingDRX")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NBIoTPagingDRX = &ie.NBIoTPagingDRX{}
			err = x.NBIoTPagingDRX.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 10
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

		// IE Field 11
		if protocolIeId == ie.ProtocolIEIDWUSAssistanceInformation {
			// check if ie is duplicated
			if x.WUSAssistanceInformation != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: WUSAssistanceInformation")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.WUSAssistanceInformation = &ie.WUSAssistanceInformation{}
			err = x.WUSAssistanceInformation.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 12
		if protocolIeId == ie.ProtocolIEIDEUTRAPagingeDRXInformation {
			// check if ie is duplicated
			if x.EUTRAPagingeDRXInformation != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: EUTRAPagingeDRXInformation")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.EUTRAPagingeDRXInformation = &ie.EUTRAPagingeDRXInformation{}
			err = x.EUTRAPagingeDRXInformation.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 13
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

		// IE Field 14
		if protocolIeId == ie.ProtocolIEIDNRPagingeDRXInformation {
			// check if ie is duplicated
			if x.NRPagingeDRXInformation != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: NRPagingeDRXInformation")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.NRPagingeDRXInformation = &ie.NRPagingeDRXInformation{}
			err = x.NRPagingeDRXInformation.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 15
		if protocolIeId == ie.ProtocolIEIDPagingCause {
			// check if ie is duplicated
			if x.PagingCause != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PagingCause")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PagingCause = &ie.PagingCause{}
			err = x.PagingCause.ReadIE(pdOpenType)
			if err != nil {
				return errors.Wrap(err, "message unmarshal failed")
			}
			continue
		}

		// IE Field 16
		if protocolIeId == ie.ProtocolIEIDPEIPSassistanceInformation {
			// check if ie is duplicated
			if x.PEIPSassistanceInformation != nil {
				// Build Abstract Syntax Error (Falsely Constructed Message)
				errTrace := errors.Errorf("Duplicated IE: PEIPSassistanceInformation")
				return ie.BuildAbstractSyntaxErr(
					x.ProcedureCode(),
					aper.Enumerated(x.MessageType()),
					x.Criticality(),
					&ie.AbstractSyntaxErrIEWrongOrderOrTooManyOccur{},
					errTrace)
			}
			x.PEIPSassistanceInformation = &ie.PEIPSassistanceInformation{}
			err = x.PEIPSassistanceInformation.ReadIE(pdOpenType)
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
	// UEPagingIdentity is mandatory but may be nil (ignored)

	// TAIListForPaging is mandatory but may be nil (ignored)

	return nil
}
