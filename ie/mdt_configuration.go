package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MDTConfiguration struct {
	MdtConfigNR    *MDTConfigurationNR                               // valueExt,optional
	MdtConfigEUTRA *MDTConfigurationEUTRA                            // valueExt,optional
	IEExtensions   *ProtocolExtensionContainerMDTConfigurationExtIEs // optional
}

func (x *MDTConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MDTConfigurationOptPresentFlag := []bool{}
	// optional field
	if x.MdtConfigNR != nil {
		MDTConfigurationOptPresentFlag = append(MDTConfigurationOptPresentFlag, true)
	} else {
		MDTConfigurationOptPresentFlag = append(MDTConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.MdtConfigEUTRA != nil {
		MDTConfigurationOptPresentFlag = append(MDTConfigurationOptPresentFlag, true)
	} else {
		MDTConfigurationOptPresentFlag = append(MDTConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MDTConfigurationOptPresentFlag = append(MDTConfigurationOptPresentFlag, true)
	} else {
		MDTConfigurationOptPresentFlag = append(MDTConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MDTConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.MdtConfigNR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MdtConfigNR.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MdtConfigNR marshal failed")
		}
	}

	// optional field
	if x.MdtConfigEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MdtConfigEUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MdtConfigEUTRA marshal failed")
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

func (x *MDTConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MDTConfigurationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&MDTConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if MDTConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MdtConfigNR = new(MDTConfigurationNR)
		err = x.MdtConfigNR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MdtConfigNR error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MDTConfigurationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MdtConfigEUTRA = new(MDTConfigurationEUTRA)
		err = x.MdtConfigEUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MdtConfigEUTRA error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if MDTConfigurationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMDTConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
