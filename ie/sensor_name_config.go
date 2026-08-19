package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	SensorNameConfigUncompensatedBarometricConfigPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	SensorNameConfigUeSpeedConfigPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	SensorNameConfigUeOrientationConfigPresentTrue aper.Enumerated = 0
)

type SensorNameConfig struct {
	Choice SensorNameConfigAlt
}

type SensorNameConfigAlt interface {
	SensorNameConfigAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type UncompensatedBarometricConfigForSensorNameConfig struct {
	Value aper.Enumerated
}

func (alt0 *UncompensatedBarometricConfigForSensorNameConfig) SensorNameConfigAltIndex() int64 {
	return int64(0)
}

func (x *UncompensatedBarometricConfigForSensorNameConfig) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *UncompensatedBarometricConfigForSensorNameConfig) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 0
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type UeSpeedConfigForSensorNameConfig struct {
	Value aper.Enumerated
}

func (alt1 *UeSpeedConfigForSensorNameConfig) SensorNameConfigAltIndex() int64 {
	return int64(1)
}

func (x *UeSpeedConfigForSensorNameConfig) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *UeSpeedConfigForSensorNameConfig) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 0
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type UeOrientationConfigForSensorNameConfig struct {
	Value aper.Enumerated
}

func (alt2 *UeOrientationConfigForSensorNameConfig) SensorNameConfigAltIndex() int64 {
	return int64(2)
}

func (x *UeOrientationConfigForSensorNameConfig) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *UeOrientationConfigForSensorNameConfig) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 0
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 ProtocolIESingleContainerSensorNameConfigExtIEs) SensorNameConfigAltIndex() int64 {
	return int64(3)
}

// Choice Type Read/Write Functions

func (x *SensorNameConfig) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 3
	var option_idx int64 = x.Choice.SensorNameConfigAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *SensorNameConfig) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 3
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(UeOrientationConfigForSensorNameConfig)
	} else if option_idx == 1 {
		x.Choice = new(UeOrientationConfigForSensorNameConfig)
	} else if option_idx == 2 {
		x.Choice = new(UeOrientationConfigForSensorNameConfig)
	} else if option_idx == 3 {
		x.Choice = new(ProtocolIESingleContainerSensorNameConfigExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
