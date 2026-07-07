package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MDTConfigurationEUTRA struct {
	MdtActivation              *MDTActivation       // valueExt,valueLB:0,valueUB:2
	AreaScopeOfMDT             *AreaScopeOfMDTEUTRA // valueLB:0,valueUB:4
	MDTMode                    *MDTModeEutra
	SignallingBasedMDTPLMNList *MDTPLMNList                                           // optional
	IEExtensions               *ProtocolExtensionContainerMDTConfigurationEUTRAExtIEs // optional
}

func (x *MDTConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MDTConfigurationEUTRAOptPresentFlag := []bool{}
	// mandatory field
	if x.MdtActivation == nil {
		return errors.Errorf("MdtActivation is missing")
	}
	// mandatory field
	if x.AreaScopeOfMDT == nil {
		return errors.Errorf("AreaScopeOfMDT is missing")
	}
	// mandatory field
	if x.MDTMode == nil {
		return errors.Errorf("MDTMode is missing")
	}
	// optional field
	if x.SignallingBasedMDTPLMNList != nil {
		MDTConfigurationEUTRAOptPresentFlag = append(MDTConfigurationEUTRAOptPresentFlag, true)
	} else {
		MDTConfigurationEUTRAOptPresentFlag = append(MDTConfigurationEUTRAOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MDTConfigurationEUTRAOptPresentFlag = append(MDTConfigurationEUTRAOptPresentFlag, true)
	} else {
		MDTConfigurationEUTRAOptPresentFlag = append(MDTConfigurationEUTRAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MDTConfigurationEUTRAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MdtActivation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MdtActivation marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AreaScopeOfMDT.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AreaScopeOfMDT marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MDTMode.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MDTMode marshal failed")
	}

	// optional field
	if x.SignallingBasedMDTPLMNList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SignallingBasedMDTPLMNList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SignallingBasedMDTPLMNList marshal failed")
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

func (x *MDTConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MDTConfigurationEUTRAOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MDTConfigurationEUTRAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MdtActivation = new(MDTActivation)
	err = x.MdtActivation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MdtActivation error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AreaScopeOfMDT = new(AreaScopeOfMDTEUTRA)
	err = x.AreaScopeOfMDT.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AreaScopeOfMDT error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MDTMode = new(MDTModeEutra)
	err = x.MDTMode.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MDTMode error")
	}

	// optional field (optPresentFlag index: 0)
	if MDTConfigurationEUTRAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SignallingBasedMDTPLMNList = new(MDTPLMNList)
		err = x.SignallingBasedMDTPLMNList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SignallingBasedMDTPLMNList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MDTConfigurationEUTRAOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMDTConfigurationEUTRAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
