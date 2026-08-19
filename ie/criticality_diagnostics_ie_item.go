package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CriticalityDiagnosticsIEItem struct {
	IECriticality *Criticality // valueLB:0,valueUB:2
	IEID          *ProtocolIEID
	TypeOfError   *TypeOfError                                                  // valueExt,valueLB:0,valueUB:1
	IEExtensions  *ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs // optional
}

func (x *CriticalityDiagnosticsIEItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CriticalityDiagnosticsIEItemOptPresentFlag := []bool{}
	// mandatory field
	if x.IECriticality == nil {
		return errors.Errorf("IECriticality is missing")
	}
	// mandatory field
	if x.IEID == nil {
		return errors.Errorf("IEID is missing")
	}
	// mandatory field
	if x.TypeOfError == nil {
		return errors.Errorf("TypeOfError is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CriticalityDiagnosticsIEItemOptPresentFlag = append(CriticalityDiagnosticsIEItemOptPresentFlag, true)
	} else {
		CriticalityDiagnosticsIEItemOptPresentFlag = append(CriticalityDiagnosticsIEItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CriticalityDiagnosticsIEItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.IECriticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IECriticality marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.IEID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IEID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TypeOfError.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TypeOfError marshal failed")
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

func (x *CriticalityDiagnosticsIEItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CriticalityDiagnosticsIEItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CriticalityDiagnosticsIEItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IECriticality = new(Criticality)
	err = x.IECriticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IECriticality error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IEID = new(ProtocolIEID)
	err = x.IEID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IEID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TypeOfError = new(TypeOfError)
	err = x.TypeOfError.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TypeOfError error")
	}

	// optional field (optPresentFlag index: 0)
	if CriticalityDiagnosticsIEItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
