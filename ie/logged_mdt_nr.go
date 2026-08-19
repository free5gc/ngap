package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LoggedMDTNr struct {
	LoggingInterval                   *LoggingInterval                             // valueExt,valueLB:0,valueUB:10
	LoggingDuration                   *LoggingDuration                             // valueExt,valueLB:0,valueUB:5
	LoggedMDTTrigger                  *LoggedMDTTrigger                            // valueLB:0,valueUB:2
	BluetoothMeasurementConfiguration *BluetoothMeasurementConfiguration           // valueExt,optional
	WLANMeasurementConfiguration      *WLANMeasurementConfiguration                // valueExt,optional
	SensorMeasurementConfiguration    *SensorMeasurementConfiguration              // valueExt,optional
	AreaScopeOfNeighCellsList         *AreaScopeOfNeighCellsList                   // optional
	IEExtensions                      *ProtocolExtensionContainerLoggedMDTNrExtIEs // optional
}

func (x *LoggedMDTNr) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LoggedMDTNrOptPresentFlag := []bool{}
	// mandatory field
	if x.LoggingInterval == nil {
		return errors.Errorf("LoggingInterval is missing")
	}
	// mandatory field
	if x.LoggingDuration == nil {
		return errors.Errorf("LoggingDuration is missing")
	}
	// mandatory field
	if x.LoggedMDTTrigger == nil {
		return errors.Errorf("LoggedMDTTrigger is missing")
	}
	// optional field
	if x.BluetoothMeasurementConfiguration != nil {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, true)
	} else {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.WLANMeasurementConfiguration != nil {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, true)
	} else {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.SensorMeasurementConfiguration != nil {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, true)
	} else {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.AreaScopeOfNeighCellsList != nil {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, true)
	} else {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, true)
	} else {
		LoggedMDTNrOptPresentFlag = append(LoggedMDTNrOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LoggedMDTNrOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.LoggingInterval.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LoggingInterval marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.LoggingDuration.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LoggingDuration marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.LoggedMDTTrigger.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LoggedMDTTrigger marshal failed")
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
	if x.SensorMeasurementConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SensorMeasurementConfiguration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SensorMeasurementConfiguration marshal failed")
		}
	}

	// optional field
	if x.AreaScopeOfNeighCellsList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AreaScopeOfNeighCellsList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AreaScopeOfNeighCellsList marshal failed")
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

func (x *LoggedMDTNr) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LoggedMDTNrOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&LoggedMDTNrOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LoggingInterval = new(LoggingInterval)
	err = x.LoggingInterval.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LoggingInterval error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LoggingDuration = new(LoggingDuration)
	err = x.LoggingDuration.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LoggingDuration error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LoggedMDTTrigger = new(LoggedMDTTrigger)
	err = x.LoggedMDTTrigger.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LoggedMDTTrigger error")
	}

	// optional field (optPresentFlag index: 0)
	if LoggedMDTNrOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.BluetoothMeasurementConfiguration = new(BluetoothMeasurementConfiguration)
		err = x.BluetoothMeasurementConfiguration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BluetoothMeasurementConfiguration error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if LoggedMDTNrOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.WLANMeasurementConfiguration = new(WLANMeasurementConfiguration)
		err = x.WLANMeasurementConfiguration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode WLANMeasurementConfiguration error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if LoggedMDTNrOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SensorMeasurementConfiguration = new(SensorMeasurementConfiguration)
		err = x.SensorMeasurementConfiguration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SensorMeasurementConfiguration error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if LoggedMDTNrOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.AreaScopeOfNeighCellsList = new(AreaScopeOfNeighCellsList)
		err = x.AreaScopeOfNeighCellsList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AreaScopeOfNeighCellsList error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if LoggedMDTNrOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLoggedMDTNrExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
