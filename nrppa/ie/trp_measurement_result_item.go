package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TrpMeasurementResultItem struct {
	MeasuredResultsValue *TrpMeasuredResultsValue                                  // valueLB:0,valueUB:4
	TimeStamp            *TimeStamp                                                // valueExt
	MeasurementQuality   *TrpMeasurementQuality                                    // valueLB:0,valueUB:2,optional
	MeasurementBeamInfo  *MeasurementBeamInfo                                      // valueExt,optional
	IEExtensions         *ProtocolExtensionContainerTrpMeasurementResultItemExtIEs // optional
}

func (x *TrpMeasurementResultItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementResultItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MeasuredResultsValue == nil {
		return errors.Errorf("MeasuredResultsValue is missing")
	}
	// mandatory field
	if x.TimeStamp == nil {
		return errors.Errorf("TimeStamp is missing")
	}
	// optional field
	if x.MeasurementQuality != nil {
		TrpMeasurementResultItemOptPresentFlag = append(TrpMeasurementResultItemOptPresentFlag, true)
	} else {
		TrpMeasurementResultItemOptPresentFlag = append(TrpMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.MeasurementBeamInfo != nil {
		TrpMeasurementResultItemOptPresentFlag = append(TrpMeasurementResultItemOptPresentFlag, true)
	} else {
		TrpMeasurementResultItemOptPresentFlag = append(TrpMeasurementResultItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TrpMeasurementResultItemOptPresentFlag = append(TrpMeasurementResultItemOptPresentFlag, true)
	} else {
		TrpMeasurementResultItemOptPresentFlag = append(TrpMeasurementResultItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementResultItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MeasuredResultsValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MeasuredResultsValue marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TimeStamp.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TimeStamp marshal failed")
	}

	// optional field
	if x.MeasurementQuality != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementQuality.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MeasurementQuality marshal failed")
		}
	}

	// optional field
	if x.MeasurementBeamInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementBeamInfo.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MeasurementBeamInfo marshal failed")
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

func (x *TrpMeasurementResultItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementResultItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementResultItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MeasuredResultsValue = new(TrpMeasuredResultsValue)
	err = x.MeasuredResultsValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MeasuredResultsValue error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TimeStamp = new(TimeStamp)
	err = x.TimeStamp.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TimeStamp error")
	}

	// optional field (optPresentFlag index: 0)
	if TrpMeasurementResultItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementQuality = new(TrpMeasurementQuality)
		err = x.MeasurementQuality.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementQuality error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TrpMeasurementResultItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementBeamInfo = new(MeasurementBeamInfo)
		err = x.MeasurementBeamInfo.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementBeamInfo error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if TrpMeasurementResultItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTrpMeasurementResultItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
