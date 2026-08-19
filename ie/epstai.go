package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EPSTAI struct {
	PLMNIdentity *PLMNIdentity
	EPSTAC       *EPSTAC
	IEExtensions *ProtocolExtensionContainerEPSTAIExtIEs // optional
}

func (x *EPSTAI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EPSTAIOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.EPSTAC == nil {
		return errors.Errorf("EPSTAC is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EPSTAIOptPresentFlag = append(EPSTAIOptPresentFlag, true)
	} else {
		EPSTAIOptPresentFlag = append(EPSTAIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EPSTAIOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.EPSTAC.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EPSTAC marshal failed")
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

func (x *EPSTAI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EPSTAIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EPSTAIOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNIdentity = new(PLMNIdentity)
	err = x.PLMNIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNIdentity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EPSTAC = new(EPSTAC)
	err = x.EPSTAC.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EPSTAC error")
	}

	// optional field (optPresentFlag index: 0)
	if EPSTAIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEPSTAIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
