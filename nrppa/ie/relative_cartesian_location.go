package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	RelativeCartesianLocationXYZunitPresentMm aper.Enumerated = 0
	RelativeCartesianLocationXYZunitPresentCm aper.Enumerated = 1
	RelativeCartesianLocationXYZunitPresentDm aper.Enumerated = 2
)

type RelativeCartesianLocation struct {
	XYZunit             *aper.Enumerated                                           // valueExt,valueLB:0,valueUB:2
	Xvalue              *int64                                                     // valueLB:-65536,valueUB:65535
	Yvalue              *int64                                                     // valueLB:-65536,valueUB:65535
	Zvalue              *int64                                                     // valueLB:-32768,valueUB:32767
	LocationUncertainty *LocationUncertainty                                       // valueExt
	IEExtensions        *ProtocolExtensionContainerRelativeCartesianLocationExtIEs // optional
}

func (x *RelativeCartesianLocation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RelativeCartesianLocationOptPresentFlag := []bool{}
	// mandatory field
	if x.XYZunit == nil {
		return errors.Errorf("XYZunit is missing")
	}
	// mandatory field
	if x.Xvalue == nil {
		return errors.Errorf("Xvalue is missing")
	}
	// mandatory field
	if x.Yvalue == nil {
		return errors.Errorf("Yvalue is missing")
	}
	// mandatory field
	if x.Zvalue == nil {
		return errors.Errorf("Zvalue is missing")
	}
	// mandatory field
	if x.LocationUncertainty == nil {
		return errors.Errorf("LocationUncertainty is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RelativeCartesianLocationOptPresentFlag = append(RelativeCartesianLocationOptPresentFlag, true)
	} else {
		RelativeCartesianLocationOptPresentFlag = append(RelativeCartesianLocationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RelativeCartesianLocationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.XYZunit), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -65536, 65535
	err = pd.WriteInteger(*(x.Xvalue), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -65536, 65535
	err = pd.WriteInteger(*(x.Yvalue), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -32768, 32767
	err = pd.WriteInteger(*(x.Zvalue), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.LocationUncertainty.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LocationUncertainty marshal failed")
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

func (x *RelativeCartesianLocation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RelativeCartesianLocationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RelativeCartesianLocationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.XYZunit = new(aper.Enumerated)
	*(x.XYZunit), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -65536, 65535
	x.Xvalue = new(int64)
	*(x.Xvalue), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -65536, 65535
	x.Yvalue = new(int64)
	*(x.Yvalue), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -32768, 32767
	x.Zvalue = new(int64)
	*(x.Zvalue), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LocationUncertainty = new(LocationUncertainty)
	err = x.LocationUncertainty.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LocationUncertainty error")
	}

	// optional field (optPresentFlag index: 0)
	if RelativeCartesianLocationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRelativeCartesianLocationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
