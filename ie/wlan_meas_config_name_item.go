package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type WLANMeasConfigNameItem struct {
	WLANName     *WLANName
	IEExtensions *ProtocolExtensionContainerWLANMeasConfigNameItemExtIEs // optional
}

func (x *WLANMeasConfigNameItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	WLANMeasConfigNameItemOptPresentFlag := []bool{}
	// mandatory field
	if x.WLANName == nil {
		return errors.Errorf("WLANName is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		WLANMeasConfigNameItemOptPresentFlag = append(WLANMeasConfigNameItemOptPresentFlag, true)
	} else {
		WLANMeasConfigNameItemOptPresentFlag = append(WLANMeasConfigNameItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(WLANMeasConfigNameItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.WLANName.Write(pd)
	if err != nil {
		return errors.Wrap(err, "WLANName marshal failed")
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

func (x *WLANMeasConfigNameItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	WLANMeasConfigNameItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&WLANMeasConfigNameItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.WLANName = new(WLANName)
	err = x.WLANName.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode WLANName error")
	}

	// optional field (optPresentFlag index: 0)
	if WLANMeasConfigNameItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerWLANMeasConfigNameItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
