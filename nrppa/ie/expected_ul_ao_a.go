package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ExpectedULAoA struct {
	ExpectedAzimuthAoA *ExpectedAzimuthAoA                            // valueExt
	ExpectedZenithAoA  *ExpectedZenithAoA                             // valueExt,optional
	IEExtensions       *ProtocolExtensionContainerExpectedULAoAExtIEs // optional
}

func (x *ExpectedULAoA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedULAoAOptPresentFlag := []bool{}
	// mandatory field
	if x.ExpectedAzimuthAoA == nil {
		return errors.Errorf("ExpectedAzimuthAoA is missing")
	}
	// optional field
	if x.ExpectedZenithAoA != nil {
		ExpectedULAoAOptPresentFlag = append(ExpectedULAoAOptPresentFlag, true)
	} else {
		ExpectedULAoAOptPresentFlag = append(ExpectedULAoAOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedULAoAOptPresentFlag = append(ExpectedULAoAOptPresentFlag, true)
	} else {
		ExpectedULAoAOptPresentFlag = append(ExpectedULAoAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedULAoAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ExpectedAzimuthAoA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExpectedAzimuthAoA marshal failed")
	}

	// optional field
	if x.ExpectedZenithAoA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedZenithAoA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedZenithAoA marshal failed")
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

func (x *ExpectedULAoA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedULAoAOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ExpectedULAoAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExpectedAzimuthAoA = new(ExpectedAzimuthAoA)
	err = x.ExpectedAzimuthAoA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExpectedAzimuthAoA error")
	}

	// optional field (optPresentFlag index: 0)
	if ExpectedULAoAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedZenithAoA = new(ExpectedZenithAoA)
		err = x.ExpectedZenithAoA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedZenithAoA error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExpectedULAoAOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedULAoAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
