package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	TrpMeasurementTimingQualityResolutionPresentM0dot1 aper.Enumerated = 0
	TrpMeasurementTimingQualityResolutionPresentM1     aper.Enumerated = 1
	TrpMeasurementTimingQualityResolutionPresentM10    aper.Enumerated = 2
	TrpMeasurementTimingQualityResolutionPresentM30    aper.Enumerated = 3
)

type TrpMeasurementTimingQuality struct {
	MeasurementQuality *int64                                                       // valueLB:0,valueUB:31
	Resolution         *aper.Enumerated                                             // valueExt,valueLB:0,valueUB:3
	IEExtensions       *ProtocolExtensionContainerTrpMeasurementTimingQualityExtIEs // optional
}

func (x *TrpMeasurementTimingQuality) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementTimingQualityOptPresentFlag := []bool{}
	// mandatory field
	if x.MeasurementQuality == nil {
		return errors.Errorf("MeasurementQuality is missing")
	}
	// mandatory field
	if x.Resolution == nil {
		return errors.Errorf("Resolution is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TrpMeasurementTimingQualityOptPresentFlag = append(TrpMeasurementTimingQualityOptPresentFlag, true)
	} else {
		TrpMeasurementTimingQualityOptPresentFlag = append(TrpMeasurementTimingQualityOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementTimingQualityOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 31
	err = pd.WriteInteger(*(x.MeasurementQuality), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
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

func (x *TrpMeasurementTimingQuality) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementTimingQualityOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementTimingQualityOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 31
	x.MeasurementQuality = new(int64)
	*(x.MeasurementQuality), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.Resolution = new(aper.Enumerated)
	*(x.Resolution), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if TrpMeasurementTimingQualityOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTrpMeasurementTimingQualityExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
