package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &CriticalityDiagnostics{}

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                                          // optional
	TriggeringMessage         *TriggeringMessage                                      // valueLB:0,valueUB:2,optional
	ProcedureCriticality      *Criticality                                            // valueLB:0,valueUB:2,optional
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList                           // optional
	IEExtensions              *ProtocolExtensionContainerCriticalityDiagnosticsExtIEs // optional
}

func (x *CriticalityDiagnostics) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CriticalityDiagnosticsOptPresentFlag := []bool{}
	// optional field
	if x.ProcedureCode != nil {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, true)
	} else {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, false)
	}
	// optional field
	if x.TriggeringMessage != nil {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, true)
	} else {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, false)
	}
	// optional field
	if x.ProcedureCriticality != nil {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, true)
	} else {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, false)
	}
	// optional field
	if x.IEsCriticalityDiagnostics != nil {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, true)
	} else {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, true)
	} else {
		CriticalityDiagnosticsOptPresentFlag = append(CriticalityDiagnosticsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CriticalityDiagnosticsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.ProcedureCode != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ProcedureCode.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ProcedureCode marshal failed")
		}
	}

	// optional field
	if x.TriggeringMessage != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TriggeringMessage.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TriggeringMessage marshal failed")
		}
	}

	// optional field
	if x.ProcedureCriticality != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ProcedureCriticality.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ProcedureCriticality marshal failed")
		}
	}

	// optional field
	if x.IEsCriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEsCriticalityDiagnostics.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEsCriticalityDiagnostics marshal failed")
		}
	}

	// optional field
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *CriticalityDiagnostics) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CriticalityDiagnosticsOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&CriticalityDiagnosticsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if CriticalityDiagnosticsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ProcedureCode = new(ProcedureCode)
		err = x.ProcedureCode.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ProcedureCode error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if CriticalityDiagnosticsOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.TriggeringMessage = new(TriggeringMessage)
		err = x.TriggeringMessage.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TriggeringMessage error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if CriticalityDiagnosticsOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.ProcedureCriticality = new(Criticality)
		err = x.ProcedureCriticality.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ProcedureCriticality error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if CriticalityDiagnosticsOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEsCriticalityDiagnostics = new(CriticalityDiagnosticsIEList)
		err = x.IEsCriticalityDiagnostics.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEsCriticalityDiagnostics error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if CriticalityDiagnosticsOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCriticalityDiagnosticsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *CriticalityDiagnostics) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.WriteSequencePreambleBitMap(optPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: id
	err = id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: criticality
	err = criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: value (open type)
	pdOpenType := aper.NewPerBitData(nil)
	err = x.Write(pdOpenType)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}

	return nil
}

func (x *CriticalityDiagnostics) ReadIE(pd *aper.PerBitData) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.ReadSequencePreambleBitMap(&optPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err = protocolIECriticality.Read(pd)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	// sequence element: value (open type)
	var pdOpenTypeBytes []byte
	pdOpenTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	pdOpenType := aper.NewPerBitData(pdOpenTypeBytes)
	err = x.Read(pdOpenType)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	return nil
}
