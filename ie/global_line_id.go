package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GlobalLineID struct {
	GlobalLineIdentity *GlobalLineIdentity
	LineType           *LineType                                     // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions       *ProtocolExtensionContainerGlobalLineIDExtIEs // optional
}

func (x *GlobalLineID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GlobalLineIDOptPresentFlag := []bool{}
	// mandatory field
	if x.GlobalLineIdentity == nil {
		return errors.Errorf("GlobalLineIdentity is missing")
	}
	// optional field
	if x.LineType != nil {
		GlobalLineIDOptPresentFlag = append(GlobalLineIDOptPresentFlag, true)
	} else {
		GlobalLineIDOptPresentFlag = append(GlobalLineIDOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		GlobalLineIDOptPresentFlag = append(GlobalLineIDOptPresentFlag, true)
	} else {
		GlobalLineIDOptPresentFlag = append(GlobalLineIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GlobalLineIDOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.GlobalLineIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GlobalLineIdentity marshal failed")
	}

	// optional field
	if x.LineType != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LineType.Write(pd)
		if err != nil {
			return errors.Wrap(err, "LineType marshal failed")
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

func (x *GlobalLineID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GlobalLineIDOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&GlobalLineIDOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GlobalLineIdentity = new(GlobalLineIdentity)
	err = x.GlobalLineIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GlobalLineIdentity error")
	}

	// optional field (optPresentFlag index: 0)
	if GlobalLineIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.LineType = new(LineType)
		err = x.LineType.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode LineType error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if GlobalLineIDOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGlobalLineIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
