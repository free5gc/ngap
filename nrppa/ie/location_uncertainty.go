package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type LocationUncertainty struct {
	HorizontalUncertainty *int64                                               // valueLB:0,valueUB:255
	HorizontalConfidence  *int64                                               // valueLB:0,valueUB:100
	VerticalUncertainty   *int64                                               // valueLB:0,valueUB:255
	VerticalConfidence    *int64                                               // valueLB:0,valueUB:100
	IEExtensions          *ProtocolExtensionContainerLocationUncertaintyExtIEs // optional
}

func (x *LocationUncertainty) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LocationUncertaintyOptPresentFlag := []bool{}
	// mandatory field
	if x.HorizontalUncertainty == nil {
		return errors.Errorf("HorizontalUncertainty is missing")
	}
	// mandatory field
	if x.HorizontalConfidence == nil {
		return errors.Errorf("HorizontalConfidence is missing")
	}
	// mandatory field
	if x.VerticalUncertainty == nil {
		return errors.Errorf("VerticalUncertainty is missing")
	}
	// mandatory field
	if x.VerticalConfidence == nil {
		return errors.Errorf("VerticalConfidence is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		LocationUncertaintyOptPresentFlag = append(LocationUncertaintyOptPresentFlag, true)
	} else {
		LocationUncertaintyOptPresentFlag = append(LocationUncertaintyOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LocationUncertaintyOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.HorizontalUncertainty), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.HorizontalConfidence), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.VerticalUncertainty), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.VerticalConfidence), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *LocationUncertainty) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LocationUncertaintyOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&LocationUncertaintyOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.HorizontalUncertainty = new(int64)
	*(x.HorizontalUncertainty), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.HorizontalConfidence = new(int64)
	*(x.HorizontalConfidence), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.VerticalUncertainty = new(int64)
	*(x.VerticalUncertainty), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.VerticalConfidence = new(int64)
	*(x.VerticalConfidence), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if LocationUncertaintyOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLocationUncertaintyExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
