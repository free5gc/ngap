package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SensorMeasConfigNameItem struct {
	SensorNameConfig *SensorNameConfig                                         // valueLB:0,valueUB:3
	IEExtensions     *ProtocolExtensionContainerSensorMeasConfigNameItemExtIEs // optional
}

func (x *SensorMeasConfigNameItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SensorMeasConfigNameItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SensorNameConfig == nil {
		return errors.Errorf("SensorNameConfig is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SensorMeasConfigNameItemOptPresentFlag = append(SensorMeasConfigNameItemOptPresentFlag, true)
	} else {
		SensorMeasConfigNameItemOptPresentFlag = append(SensorMeasConfigNameItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SensorMeasConfigNameItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SensorNameConfig.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SensorNameConfig marshal failed")
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

func (x *SensorMeasConfigNameItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SensorMeasConfigNameItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SensorMeasConfigNameItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SensorNameConfig = new(SensorNameConfig)
	err = x.SensorNameConfig.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SensorNameConfig error")
	}

	// optional field (optPresentFlag index: 0)
	if SensorMeasConfigNameItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSensorMeasConfigNameItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
