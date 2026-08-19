package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	RelativeGeodeticLocationMilliArcSecondUnitsPresentZerodot03 aper.Enumerated = 0
	RelativeGeodeticLocationMilliArcSecondUnitsPresentZerodot3  aper.Enumerated = 1
	RelativeGeodeticLocationMilliArcSecondUnitsPresentThree     aper.Enumerated = 2
)

const ( /* Enum Type */
	RelativeGeodeticLocationHeightUnitsPresentMm aper.Enumerated = 0
	RelativeGeodeticLocationHeightUnitsPresentCm aper.Enumerated = 1
	RelativeGeodeticLocationHeightUnitsPresentM  aper.Enumerated = 2
)

type RelativeGeodeticLocation struct {
	MilliArcSecondUnits *aper.Enumerated                                          // valueExt,valueLB:0,valueUB:2
	HeightUnits         *aper.Enumerated                                          // valueExt,valueLB:0,valueUB:2
	DeltaLatitude       *int64                                                    // valueLB:-1024,valueUB:1023
	DeltaLongitude      *int64                                                    // valueLB:-1024,valueUB:1023
	DeltaHeight         *int64                                                    // valueLB:-1024,valueUB:1023
	LocationUncertainty *LocationUncertainty                                      // valueExt
	IEExtensions        *ProtocolExtensionContainerRelativeGeodeticLocationExtIEs // optional
}

func (x *RelativeGeodeticLocation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RelativeGeodeticLocationOptPresentFlag := []bool{}
	// mandatory field
	if x.MilliArcSecondUnits == nil {
		return errors.Errorf("MilliArcSecondUnits is missing")
	}
	// mandatory field
	if x.HeightUnits == nil {
		return errors.Errorf("HeightUnits is missing")
	}
	// mandatory field
	if x.DeltaLatitude == nil {
		return errors.Errorf("DeltaLatitude is missing")
	}
	// mandatory field
	if x.DeltaLongitude == nil {
		return errors.Errorf("DeltaLongitude is missing")
	}
	// mandatory field
	if x.DeltaHeight == nil {
		return errors.Errorf("DeltaHeight is missing")
	}
	// mandatory field
	if x.LocationUncertainty == nil {
		return errors.Errorf("LocationUncertainty is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RelativeGeodeticLocationOptPresentFlag = append(RelativeGeodeticLocationOptPresentFlag, true)
	} else {
		RelativeGeodeticLocationOptPresentFlag = append(RelativeGeodeticLocationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RelativeGeodeticLocationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.MilliArcSecondUnits), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.HeightUnits), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -1024, 1023
	err = pd.WriteInteger(*(x.DeltaLatitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -1024, 1023
	err = pd.WriteInteger(*(x.DeltaLongitude), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -1024, 1023
	err = pd.WriteInteger(*(x.DeltaHeight), false, vLb, vUb)
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

func (x *RelativeGeodeticLocation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RelativeGeodeticLocationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RelativeGeodeticLocationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.MilliArcSecondUnits = new(aper.Enumerated)
	*(x.MilliArcSecondUnits), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.HeightUnits = new(aper.Enumerated)
	*(x.HeightUnits), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -1024, 1023
	x.DeltaLatitude = new(int64)
	*(x.DeltaLatitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -1024, 1023
	x.DeltaLongitude = new(int64)
	*(x.DeltaLongitude), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -1024, 1023
	x.DeltaHeight = new(int64)
	*(x.DeltaHeight), err = pd.ReadInteger(false, vLb, vUb)
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
	if RelativeGeodeticLocationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRelativeGeodeticLocationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
