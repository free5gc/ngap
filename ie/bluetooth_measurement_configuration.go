package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	BluetoothMeasurementConfigurationBtRssiPresentTrue aper.Enumerated = 0
)

type BluetoothMeasurementConfiguration struct {
	BluetoothMeasConfig         *BluetoothMeasConfig                                               // valueExt,valueLB:0,valueUB:0
	BluetoothMeasConfigNameList *BluetoothMeasConfigNameList                                       // optional
	BtRssi                      *aper.Enumerated                                                   // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions                *ProtocolExtensionContainerBluetoothMeasurementConfigurationExtIEs // optional
}

func (x *BluetoothMeasurementConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	BluetoothMeasurementConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.BluetoothMeasConfig == nil {
		return errors.Errorf("BluetoothMeasConfig is missing")
	}
	// optional field
	if x.BluetoothMeasConfigNameList != nil {
		BluetoothMeasurementConfigurationOptPresentFlag = append(BluetoothMeasurementConfigurationOptPresentFlag, true)
	} else {
		BluetoothMeasurementConfigurationOptPresentFlag = append(BluetoothMeasurementConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.BtRssi != nil {
		BluetoothMeasurementConfigurationOptPresentFlag = append(BluetoothMeasurementConfigurationOptPresentFlag, true)
	} else {
		BluetoothMeasurementConfigurationOptPresentFlag = append(BluetoothMeasurementConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		BluetoothMeasurementConfigurationOptPresentFlag = append(BluetoothMeasurementConfigurationOptPresentFlag, true)
	} else {
		BluetoothMeasurementConfigurationOptPresentFlag = append(BluetoothMeasurementConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(BluetoothMeasurementConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.BluetoothMeasConfig.Write(pd)
	if err != nil {
		return errors.Wrap(err, "BluetoothMeasConfig marshal failed")
	}

	// optional field
	if x.BluetoothMeasConfigNameList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.BluetoothMeasConfigNameList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "BluetoothMeasConfigNameList marshal failed")
		}
	}

	// optional field
	if x.BtRssi != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.BtRssi), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
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

func (x *BluetoothMeasurementConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	BluetoothMeasurementConfigurationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&BluetoothMeasurementConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.BluetoothMeasConfig = new(BluetoothMeasConfig)
	err = x.BluetoothMeasConfig.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode BluetoothMeasConfig error")
	}

	// optional field (optPresentFlag index: 0)
	if BluetoothMeasurementConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.BluetoothMeasConfigNameList = new(BluetoothMeasConfigNameList)
		err = x.BluetoothMeasConfigNameList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BluetoothMeasConfigNameList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if BluetoothMeasurementConfigurationOptPresentFlag[1] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.BtRssi = new(aper.Enumerated)
		*(x.BtRssi), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if BluetoothMeasurementConfigurationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerBluetoothMeasurementConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
