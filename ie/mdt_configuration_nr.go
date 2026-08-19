package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MDTConfigurationNR struct {
	MdtActivation              *MDTActivation                                      // valueExt,valueLB:0,valueUB:2
	AreaScopeOfMDT             *AreaScopeOfMDTNR                                   // valueLB:0,valueUB:4
	MDTModeNr                  *MDTModeNr                                          // valueLB:0,valueUB:2
	SignallingBasedMDTPLMNList *MDTPLMNList                                        // optional
	IEExtensions               *ProtocolExtensionContainerMDTConfigurationNRExtIEs // optional
}

func (x *MDTConfigurationNR) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MDTConfigurationNROptPresentFlag := []bool{}
	// mandatory field
	if x.MdtActivation == nil {
		return errors.Errorf("MdtActivation is missing")
	}
	// mandatory field
	if x.AreaScopeOfMDT == nil {
		return errors.Errorf("AreaScopeOfMDT is missing")
	}
	// mandatory field
	if x.MDTModeNr == nil {
		return errors.Errorf("MDTModeNr is missing")
	}
	// optional field
	if x.SignallingBasedMDTPLMNList != nil {
		MDTConfigurationNROptPresentFlag = append(MDTConfigurationNROptPresentFlag, true)
	} else {
		MDTConfigurationNROptPresentFlag = append(MDTConfigurationNROptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MDTConfigurationNROptPresentFlag = append(MDTConfigurationNROptPresentFlag, true)
	} else {
		MDTConfigurationNROptPresentFlag = append(MDTConfigurationNROptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MDTConfigurationNROptPresentFlag, true)
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
	err = x.MDTModeNr.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MDTModeNr marshal failed")
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

func (x *MDTConfigurationNR) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MDTConfigurationNROptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MDTConfigurationNROptPresentFlag, true)
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
	x.AreaScopeOfMDT = new(AreaScopeOfMDTNR)
	err = x.AreaScopeOfMDT.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AreaScopeOfMDT error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MDTModeNr = new(MDTModeNr)
	err = x.MDTModeNr.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MDTModeNr error")
	}

	// optional field (optPresentFlag index: 0)
	if MDTConfigurationNROptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SignallingBasedMDTPLMNList = new(MDTPLMNList)
		err = x.SignallingBasedMDTPLMNList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SignallingBasedMDTPLMNList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MDTConfigurationNROptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMDTConfigurationNRExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
