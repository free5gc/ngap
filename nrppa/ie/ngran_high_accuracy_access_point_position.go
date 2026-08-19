package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type NGRANHighAccuracyAccessPointPosition struct {
	Latitude               *int64                                                                // valueLB:-2147483648,valueUB:2147483647
	Longitude              *int64                                                                // valueLB:-2147483648,valueUB:2147483647
	Altitude               *int64                                                                // valueLB:-64000,valueUB:1280000
	UncertaintySemiMajor   *int64                                                                // valueLB:0,valueUB:255
	UncertaintySemiMinor   *int64                                                                // valueLB:0,valueUB:255
	OrientationOfMajorAxis *int64                                                                // valueLB:0,valueUB:179
	HorizontalConfidence   *int64                                                                // valueLB:0,valueUB:100
	UncertaintyAltitude    *int64                                                                // valueLB:0,valueUB:255
	VerticalConfidence     *int64                                                                // valueLB:0,valueUB:100
	IEExtensions           *ProtocolExtensionContainerNGRANHighAccuracyAccessPointPositionExtIEs // optional
}

func (x *NGRANHighAccuracyAccessPointPosition) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANHighAccuracyAccessPointPositionOptPresentFlag := []bool{}
	// mandatory field
	if x.Latitude == nil {
		return errors.Errorf("Latitude is missing")
	}
	// mandatory field
	if x.Longitude == nil {
		return errors.Errorf("Longitude is missing")
	}
	// mandatory field
	if x.Altitude == nil {
		return errors.Errorf("Altitude is missing")
	}
	// mandatory field
	if x.UncertaintySemiMajor == nil {
		return errors.Errorf("UncertaintySemiMajor is missing")
	}
	// mandatory field
	if x.UncertaintySemiMinor == nil {
		return errors.Errorf("UncertaintySemiMinor is missing")
	}
	// mandatory field
	if x.OrientationOfMajorAxis == nil {
		return errors.Errorf("OrientationOfMajorAxis is missing")
	}
	// mandatory field
	if x.HorizontalConfidence == nil {
		return errors.Errorf("HorizontalConfidence is missing")
	}
	// mandatory field
	if x.UncertaintyAltitude == nil {
		return errors.Errorf("UncertaintyAltitude is missing")
	}
	// mandatory field
	if x.VerticalConfidence == nil {
		return errors.Errorf("VerticalConfidence is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANHighAccuracyAccessPointPositionOptPresentFlag = append(NGRANHighAccuracyAccessPointPositionOptPresentFlag, true)
	} else {
		NGRANHighAccuracyAccessPointPositionOptPresentFlag = append(NGRANHighAccuracyAccessPointPositionOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANHighAccuracyAccessPointPositionOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = -2147483648, 2147483647
	err = pd.WriteInteger(*(x.Latitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -2147483648, 2147483647
	err = pd.WriteInteger(*(x.Longitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -64000, 1280000
	err = pd.WriteInteger(*(x.Altitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.UncertaintySemiMajor), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.UncertaintySemiMinor), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 179
	err = pd.WriteInteger(*(x.OrientationOfMajorAxis), false, vLb, vUb)
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
	err = pd.WriteInteger(*(x.UncertaintyAltitude), false, vLb, vUb)
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

func (x *NGRANHighAccuracyAccessPointPosition) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANHighAccuracyAccessPointPositionOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANHighAccuracyAccessPointPositionOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -2147483648, 2147483647
	x.Latitude = new(int64)
	*(x.Latitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -2147483648, 2147483647
	x.Longitude = new(int64)
	*(x.Longitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -64000, 1280000
	x.Altitude = new(int64)
	*(x.Altitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.UncertaintySemiMajor = new(int64)
	*(x.UncertaintySemiMajor), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.UncertaintySemiMinor = new(int64)
	*(x.UncertaintySemiMinor), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 179
	x.OrientationOfMajorAxis = new(int64)
	*(x.OrientationOfMajorAxis), err = pd.ReadInteger(false, vLb, vUb)
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
	x.UncertaintyAltitude = new(int64)
	*(x.UncertaintyAltitude), err = pd.ReadInteger(false, vLb, vUb)
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
	if NGRANHighAccuracyAccessPointPositionOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANHighAccuracyAccessPointPositionExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
