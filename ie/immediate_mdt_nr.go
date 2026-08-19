package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ImmediateMDTNr struct {
	MeasurementsToActivate            *MeasurementsToActivate
	M1Configuration                   *M1Configuration                                // valueExt,optional
	M4Configuration                   *M4Configuration                                // valueExt,optional
	M5Configuration                   *M5Configuration                                // valueExt,optional
	M6Configuration                   *M6Configuration                                // valueExt,optional
	M7Configuration                   *M7Configuration                                // valueExt,optional
	BluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration              // valueExt,optional
	WLANMeasurementConfiguration      *WLANMeasurementConfiguration                   // valueExt,optional
	MDTLocationInfo                   *MDTLocationInfo                                // valueExt,optional
	SensorMeasurementConfiguration    *SensorMeasurementConfiguration                 // valueExt,optional
	IEExtensions                      *ProtocolExtensionContainerImmediateMDTNrExtIEs // optional
}

func (x *ImmediateMDTNr) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ImmediateMDTNrOptPresentFlag := []bool{}
	// mandatory field
	if x.MeasurementsToActivate == nil {
		return errors.Errorf("MeasurementsToActivate is missing")
	}
	// optional field
	if x.M1Configuration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.M4Configuration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.M5Configuration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.M6Configuration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.M7Configuration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.BluetoothMeasurementConfiguration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.WLANMeasurementConfiguration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.MDTLocationInfo != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.SensorMeasurementConfiguration != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, true)
	} else {
		ImmediateMDTNrOptPresentFlag = append(ImmediateMDTNrOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ImmediateMDTNrOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MeasurementsToActivate.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MeasurementsToActivate marshal failed")
	}

	// optional field
	if x.M1Configuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M1Configuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M1Configuration marshal failed")
		}
	}

	// optional field
	if x.M4Configuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M4Configuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M4Configuration marshal failed")
		}
	}

	// optional field
	if x.M5Configuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M5Configuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M5Configuration marshal failed")
		}
	}

	// optional field
	if x.M6Configuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M6Configuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M6Configuration marshal failed")
		}
	}

	// optional field
	if x.M7Configuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M7Configuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M7Configuration marshal failed")
		}
	}

	// optional field
	if x.BluetoothMeasurementConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.BluetoothMeasurementConfiguration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "BluetoothMeasurementConfiguration marshal failed")
		}
	}

	// optional field
	if x.WLANMeasurementConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WLANMeasurementConfiguration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "WLANMeasurementConfiguration marshal failed")
		}
	}

	// optional field
	if x.MDTLocationInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MDTLocationInfo.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MDTLocationInfo marshal failed")
		}
	}

	// optional field
	if x.SensorMeasurementConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SensorMeasurementConfiguration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SensorMeasurementConfiguration marshal failed")
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

func (x *ImmediateMDTNr) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ImmediateMDTNrOptPresentFlag := make([]bool, 10)
	err = pd.ReadSequencePreambleBitMap(&ImmediateMDTNrOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MeasurementsToActivate = new(MeasurementsToActivate)
	err = x.MeasurementsToActivate.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MeasurementsToActivate error")
	}

	// optional field (optPresentFlag index: 0)
	if ImmediateMDTNrOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.M1Configuration = new(M1Configuration)
		err = x.M1Configuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M1Configuration error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ImmediateMDTNrOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.M4Configuration = new(M4Configuration)
		err = x.M4Configuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M4Configuration error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ImmediateMDTNrOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.M5Configuration = new(M5Configuration)
		err = x.M5Configuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M5Configuration error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ImmediateMDTNrOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.M6Configuration = new(M6Configuration)
		err = x.M6Configuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M6Configuration error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if ImmediateMDTNrOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.M7Configuration = new(M7Configuration)
		err = x.M7Configuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M7Configuration error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if ImmediateMDTNrOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.BluetoothMeasurementConfiguration = new(BluetoothMeasurementConfiguration)
		err = x.BluetoothMeasurementConfiguration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BluetoothMeasurementConfiguration error")
		}
	}

	// optional field (optPresentFlag index: 6)
	if ImmediateMDTNrOptPresentFlag[6] {
		// Read struct defined elsewhere (Pointer)
		x.WLANMeasurementConfiguration = new(WLANMeasurementConfiguration)
		err = x.WLANMeasurementConfiguration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode WLANMeasurementConfiguration error")
		}
	}

	// optional field (optPresentFlag index: 7)
	if ImmediateMDTNrOptPresentFlag[7] {
		// Read struct defined elsewhere (Pointer)
		x.MDTLocationInfo = new(MDTLocationInfo)
		err = x.MDTLocationInfo.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MDTLocationInfo error")
		}
	}

	// optional field (optPresentFlag index: 8)
	if ImmediateMDTNrOptPresentFlag[8] {
		// Read struct defined elsewhere (Pointer)
		x.SensorMeasurementConfiguration = new(SensorMeasurementConfiguration)
		err = x.SensorMeasurementConfiguration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SensorMeasurementConfiguration error")
		}
	}

	// optional field (optPresentFlag index: 9)
	if ImmediateMDTNrOptPresentFlag[9] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerImmediateMDTNrExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
