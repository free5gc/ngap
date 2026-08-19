package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	TrpMeasurementAngleQualityResolutionPresentDeg0dot1 aper.Enumerated = 0
)

type TrpMeasurementAngleQuality struct {
	AzimuthQuality *int64                                                      // valueLB:0,valueUB:255
	ZenithQuality  *int64                                                      // valueLB:0,valueUB:255,optional
	Resolution     *aper.Enumerated                                            // valueExt,valueLB:0,valueUB:0
	IEExtensions   *ProtocolExtensionContainerTrpMeasurementAngleQualityExtIEs // optional
}

func (x *TrpMeasurementAngleQuality) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementAngleQualityOptPresentFlag := []bool{}
	// mandatory field
	if x.AzimuthQuality == nil {
		return errors.Errorf("AzimuthQuality is missing")
	}
	// optional field
	if x.ZenithQuality != nil {
		TrpMeasurementAngleQualityOptPresentFlag = append(TrpMeasurementAngleQualityOptPresentFlag, true)
	} else {
		TrpMeasurementAngleQualityOptPresentFlag = append(TrpMeasurementAngleQualityOptPresentFlag, false)
	}
	// mandatory field
	if x.Resolution == nil {
		return errors.Errorf("Resolution is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TrpMeasurementAngleQualityOptPresentFlag = append(TrpMeasurementAngleQualityOptPresentFlag, true)
	} else {
		TrpMeasurementAngleQualityOptPresentFlag = append(TrpMeasurementAngleQualityOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementAngleQualityOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.AzimuthQuality), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.ZenithQuality != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 255
		err = pd.WriteInteger(*(x.ZenithQuality), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.Resolution), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *TrpMeasurementAngleQuality) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementAngleQualityOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementAngleQualityOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.AzimuthQuality = new(int64)
	*(x.AzimuthQuality), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if TrpMeasurementAngleQualityOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 255
		x.ZenithQuality = new(int64)
		*(x.ZenithQuality), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.Resolution = new(aper.Enumerated)
	*(x.Resolution), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 1)
	if TrpMeasurementAngleQualityOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTrpMeasurementAngleQualityExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
