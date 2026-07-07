package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SensorMeasurementConfiguration struct {
	SensorMeasConfig         *SensorMeasConfig                                               // valueExt,valueLB:0,valueUB:0
	SensorMeasConfigNameList *SensorMeasConfigNameList                                       // optional
	IEExtensions             *ProtocolExtensionContainerSensorMeasurementConfigurationExtIEs // optional
}

func (x *SensorMeasurementConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SensorMeasurementConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.SensorMeasConfig == nil {
		return errors.Errorf("SensorMeasConfig is missing")
	}
	// optional field
	if x.SensorMeasConfigNameList != nil {
		SensorMeasurementConfigurationOptPresentFlag = append(SensorMeasurementConfigurationOptPresentFlag, true)
	} else {
		SensorMeasurementConfigurationOptPresentFlag = append(SensorMeasurementConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SensorMeasurementConfigurationOptPresentFlag = append(SensorMeasurementConfigurationOptPresentFlag, true)
	} else {
		SensorMeasurementConfigurationOptPresentFlag = append(SensorMeasurementConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SensorMeasurementConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SensorMeasConfig.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SensorMeasConfig marshal failed")
	}

	// optional field
	if x.SensorMeasConfigNameList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SensorMeasConfigNameList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SensorMeasConfigNameList marshal failed")
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

func (x *SensorMeasurementConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SensorMeasurementConfigurationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SensorMeasurementConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SensorMeasConfig = new(SensorMeasConfig)
	err = x.SensorMeasConfig.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SensorMeasConfig error")
	}

	// optional field (optPresentFlag index: 0)
	if SensorMeasurementConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SensorMeasConfigNameList = new(SensorMeasConfigNameList)
		err = x.SensorMeasConfigNameList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SensorMeasConfigNameList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SensorMeasurementConfigurationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSensorMeasurementConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
