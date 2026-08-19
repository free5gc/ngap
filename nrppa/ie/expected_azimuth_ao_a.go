package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ExpectedAzimuthAoA struct {
	ExpectedAzimuthAoAValue       *ExpectedValueAoA
	ExpectedAzimuthAoAUncertainty *UncertaintyRangeAoA
	IEExtensions                  *ProtocolExtensionContainerExpectedAzimuthAoAExtIEs // optional
}

func (x *ExpectedAzimuthAoA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedAzimuthAoAOptPresentFlag := []bool{}
	// mandatory field
	if x.ExpectedAzimuthAoAValue == nil {
		return errors.Errorf("ExpectedAzimuthAoAValue is missing")
	}
	// mandatory field
	if x.ExpectedAzimuthAoAUncertainty == nil {
		return errors.Errorf("ExpectedAzimuthAoAUncertainty is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedAzimuthAoAOptPresentFlag = append(ExpectedAzimuthAoAOptPresentFlag, true)
	} else {
		ExpectedAzimuthAoAOptPresentFlag = append(ExpectedAzimuthAoAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedAzimuthAoAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ExpectedAzimuthAoAValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExpectedAzimuthAoAValue marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ExpectedAzimuthAoAUncertainty.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExpectedAzimuthAoAUncertainty marshal failed")
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

func (x *ExpectedAzimuthAoA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedAzimuthAoAOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ExpectedAzimuthAoAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExpectedAzimuthAoAValue = new(ExpectedValueAoA)
	err = x.ExpectedAzimuthAoAValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExpectedAzimuthAoAValue error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExpectedAzimuthAoAUncertainty = new(UncertaintyRangeAoA)
	err = x.ExpectedAzimuthAoAUncertainty.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExpectedAzimuthAoAUncertainty error")
	}

	// optional field (optPresentFlag index: 0)
	if ExpectedAzimuthAoAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedAzimuthAoAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
