package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type BluetoothMeasConfigNameItem struct {
	BluetoothName *BluetoothName
	IEExtensions  *ProtocolExtensionContainerBluetoothMeasConfigNameItemExtIEs // optional
}

func (x *BluetoothMeasConfigNameItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	BluetoothMeasConfigNameItemOptPresentFlag := []bool{}
	// mandatory field
	if x.BluetoothName == nil {
		return errors.Errorf("BluetoothName is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		BluetoothMeasConfigNameItemOptPresentFlag = append(BluetoothMeasConfigNameItemOptPresentFlag, true)
	} else {
		BluetoothMeasConfigNameItemOptPresentFlag = append(BluetoothMeasConfigNameItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(BluetoothMeasConfigNameItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.BluetoothName.Write(pd)
	if err != nil {
		return errors.Wrap(err, "BluetoothName marshal failed")
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

func (x *BluetoothMeasConfigNameItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	BluetoothMeasConfigNameItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&BluetoothMeasConfigNameItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.BluetoothName = new(BluetoothName)
	err = x.BluetoothName.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode BluetoothName error")
	}

	// optional field (optPresentFlag index: 0)
	if BluetoothMeasConfigNameItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerBluetoothMeasConfigNameItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
