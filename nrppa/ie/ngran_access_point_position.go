package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	NGRANAccessPointPositionLatitudeSignPresentNorth aper.Enumerated = 0
	NGRANAccessPointPositionLatitudeSignPresentSouth aper.Enumerated = 1
)

const ( /* Enum Type */
	NGRANAccessPointPositionDirectionOfAltitudePresentHeight aper.Enumerated = 0
	NGRANAccessPointPositionDirectionOfAltitudePresentDepth  aper.Enumerated = 1
)

type NGRANAccessPointPosition struct {
	LatitudeSign           *aper.Enumerated                                          // valueLB:0,valueUB:1
	Latitude               *int64                                                    // valueLB:0,valueUB:8388607
	Longitude              *int64                                                    // valueLB:-8388608,valueUB:8388607
	DirectionOfAltitude    *aper.Enumerated                                          // valueLB:0,valueUB:1
	Altitude               *int64                                                    // valueLB:0,valueUB:32767
	UncertaintySemiMajor   *int64                                                    // valueLB:0,valueUB:127
	UncertaintySemiMinor   *int64                                                    // valueLB:0,valueUB:127
	OrientationOfMajorAxis *int64                                                    // valueLB:0,valueUB:179
	UncertaintyAltitude    *int64                                                    // valueLB:0,valueUB:127
	Confidence             *int64                                                    // valueLB:0,valueUB:100
	IEExtensions           *ProtocolExtensionContainerNGRANAccessPointPositionExtIEs // optional
}

func (x *NGRANAccessPointPosition) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANAccessPointPositionOptPresentFlag := []bool{}
	// mandatory field
	if x.LatitudeSign == nil {
		return errors.Errorf("LatitudeSign is missing")
	}
	// mandatory field
	if x.Latitude == nil {
		return errors.Errorf("Latitude is missing")
	}
	// mandatory field
	if x.Longitude == nil {
		return errors.Errorf("Longitude is missing")
	}
	// mandatory field
	if x.DirectionOfAltitude == nil {
		return errors.Errorf("DirectionOfAltitude is missing")
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
	if x.UncertaintyAltitude == nil {
		return errors.Errorf("UncertaintyAltitude is missing")
	}
	// mandatory field
	if x.Confidence == nil {
		return errors.Errorf("Confidence is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANAccessPointPositionOptPresentFlag = append(NGRANAccessPointPositionOptPresentFlag, true)
	} else {
		NGRANAccessPointPositionOptPresentFlag = append(NGRANAccessPointPositionOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANAccessPointPositionOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.LatitudeSign), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 8388607
	err = pd.WriteInteger(*(x.Latitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -8388608, 8388607
	err = pd.WriteInteger(*(x.Longitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.DirectionOfAltitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 32767
	err = pd.WriteInteger(*(x.Altitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 127
	err = pd.WriteInteger(*(x.UncertaintySemiMajor), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 127
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
	*vLb, *vUb = 0, 127
	err = pd.WriteInteger(*(x.UncertaintyAltitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.Confidence), false, vLb, vUb)
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

func (x *NGRANAccessPointPosition) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANAccessPointPositionOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANAccessPointPositionOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.LatitudeSign = new(aper.Enumerated)
	*(x.LatitudeSign), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 8388607
	x.Latitude = new(int64)
	*(x.Latitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -8388608, 8388607
	x.Longitude = new(int64)
	*(x.Longitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.DirectionOfAltitude = new(aper.Enumerated)
	*(x.DirectionOfAltitude), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 32767
	x.Altitude = new(int64)
	*(x.Altitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 127
	x.UncertaintySemiMajor = new(int64)
	*(x.UncertaintySemiMajor), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 127
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
	*vLb, *vUb = 0, 127
	x.UncertaintyAltitude = new(int64)
	*(x.UncertaintyAltitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.Confidence = new(int64)
	*(x.Confidence), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if NGRANAccessPointPositionOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANAccessPointPositionExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
