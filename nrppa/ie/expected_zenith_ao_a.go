package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ExpectedZenithAoA struct {
	ExpectedZenithAoAValue       *ExpectedValueZoA
	ExpectedZenithAoAUncertainty *UncertaintyRangeZoA
	IEExtensions                 *ProtocolExtensionContainerExpectedZenithAoAExtIEs // optional
}

func (x *ExpectedZenithAoA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedZenithAoAOptPresentFlag := []bool{}
	// mandatory field
	if x.ExpectedZenithAoAValue == nil {
		return errors.Errorf("ExpectedZenithAoAValue is missing")
	}
	// mandatory field
	if x.ExpectedZenithAoAUncertainty == nil {
		return errors.Errorf("ExpectedZenithAoAUncertainty is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedZenithAoAOptPresentFlag = append(ExpectedZenithAoAOptPresentFlag, true)
	} else {
		ExpectedZenithAoAOptPresentFlag = append(ExpectedZenithAoAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedZenithAoAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ExpectedZenithAoAValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExpectedZenithAoAValue marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ExpectedZenithAoAUncertainty.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExpectedZenithAoAUncertainty marshal failed")
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

func (x *ExpectedZenithAoA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedZenithAoAOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ExpectedZenithAoAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExpectedZenithAoAValue = new(ExpectedValueZoA)
	err = x.ExpectedZenithAoAValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExpectedZenithAoAValue error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExpectedZenithAoAUncertainty = new(UncertaintyRangeZoA)
	err = x.ExpectedZenithAoAUncertainty.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExpectedZenithAoAUncertainty error")
	}

	// optional field (optPresentFlag index: 0)
	if ExpectedZenithAoAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedZenithAoAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
