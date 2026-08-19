package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MDTLocationInfo struct {
	MDTLocationInformation *MDTLocationInformation
	IEExtensions           *ProtocolExtensionContainerMDTLocationInfoExtIEs // optional
}

func (x *MDTLocationInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MDTLocationInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.MDTLocationInformation == nil {
		return errors.Errorf("MDTLocationInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		MDTLocationInfoOptPresentFlag = append(MDTLocationInfoOptPresentFlag, true)
	} else {
		MDTLocationInfoOptPresentFlag = append(MDTLocationInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MDTLocationInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MDTLocationInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MDTLocationInformation marshal failed")
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

func (x *MDTLocationInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MDTLocationInfoOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&MDTLocationInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MDTLocationInformation = new(MDTLocationInformation)
	err = x.MDTLocationInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MDTLocationInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if MDTLocationInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMDTLocationInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
