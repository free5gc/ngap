package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ExpectedZoAOnly struct {
	ExpectedZoAOnly *ExpectedZenithAoA                               // valueExt
	IEExtensions    *ProtocolExtensionContainerExpectedZoAOnlyExtIEs // optional
}

func (x *ExpectedZoAOnly) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedZoAOnlyOptPresentFlag := []bool{}
	// mandatory field
	if x.ExpectedZoAOnly == nil {
		return errors.Errorf("ExpectedZoAOnly is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedZoAOnlyOptPresentFlag = append(ExpectedZoAOnlyOptPresentFlag, true)
	} else {
		ExpectedZoAOnlyOptPresentFlag = append(ExpectedZoAOnlyOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedZoAOnlyOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ExpectedZoAOnly.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExpectedZoAOnly marshal failed")
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

func (x *ExpectedZoAOnly) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedZoAOnlyOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ExpectedZoAOnlyOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExpectedZoAOnly = new(ExpectedZenithAoA)
	err = x.ExpectedZoAOnly.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExpectedZoAOnly error")
	}

	// optional field (optPresentFlag index: 0)
	if ExpectedZoAOnlyOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedZoAOnlyExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
