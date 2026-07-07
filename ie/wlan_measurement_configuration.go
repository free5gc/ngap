package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	WLANMeasurementConfigurationWlanRssiPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	WLANMeasurementConfigurationWlanRttPresentTrue aper.Enumerated = 0
)

type WLANMeasurementConfiguration struct {
	WlanMeasConfig         *WLANMeasConfig                                               // valueExt,valueLB:0,valueUB:0
	WlanMeasConfigNameList *WLANMeasConfigNameList                                       // optional
	WlanRssi               *aper.Enumerated                                              // valueExt,valueLB:0,valueUB:0,optional
	WlanRtt                *aper.Enumerated                                              // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions           *ProtocolExtensionContainerWLANMeasurementConfigurationExtIEs // optional
}

func (x *WLANMeasurementConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	WLANMeasurementConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.WlanMeasConfig == nil {
		return errors.Errorf("WlanMeasConfig is missing")
	}
	// optional field
	if x.WlanMeasConfigNameList != nil {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, true)
	} else {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.WlanRssi != nil {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, true)
	} else {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.WlanRtt != nil {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, true)
	} else {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, true)
	} else {
		WLANMeasurementConfigurationOptPresentFlag = append(WLANMeasurementConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(WLANMeasurementConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.WlanMeasConfig.Write(pd)
	if err != nil {
		return errors.Wrap(err, "WlanMeasConfig marshal failed")
	}

	// optional field
	if x.WlanMeasConfigNameList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WlanMeasConfigNameList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "WlanMeasConfigNameList marshal failed")
		}
	}

	// optional field
	if x.WlanRssi != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.WlanRssi), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.WlanRtt != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.WlanRtt), true, vLb, vUb)
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

func (x *WLANMeasurementConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	WLANMeasurementConfigurationOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&WLANMeasurementConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.WlanMeasConfig = new(WLANMeasConfig)
	err = x.WlanMeasConfig.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode WlanMeasConfig error")
	}

	// optional field (optPresentFlag index: 0)
	if WLANMeasurementConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.WlanMeasConfigNameList = new(WLANMeasConfigNameList)
		err = x.WlanMeasConfigNameList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode WlanMeasConfigNameList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if WLANMeasurementConfigurationOptPresentFlag[1] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.WlanRssi = new(aper.Enumerated)
		*(x.WlanRssi), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if WLANMeasurementConfigurationOptPresentFlag[2] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.WlanRtt = new(aper.Enumerated)
		*(x.WlanRtt), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if WLANMeasurementConfigurationOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerWLANMeasurementConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
