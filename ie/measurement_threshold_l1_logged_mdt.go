package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MeasurementThresholdL1LoggedMDT struct {
	Choice MeasurementThresholdL1LoggedMDTAlt
}

type MeasurementThresholdL1LoggedMDTAlt interface {
	MeasurementThresholdL1LoggedMDTAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 ThresholdRSRP) MeasurementThresholdL1LoggedMDTAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 ThresholdRSRQ) MeasurementThresholdL1LoggedMDTAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 ProtocolIESingleContainerMeasurementThresholdL1LoggedMDTExtIEs) MeasurementThresholdL1LoggedMDTAltIndex() int64 {
	return int64(2)
}

// Choice Type Read/Write Functions

func (x *MeasurementThresholdL1LoggedMDT) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64 = x.Choice.MeasurementThresholdL1LoggedMDTAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *MeasurementThresholdL1LoggedMDT) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(ThresholdRSRP)
	} else if option_idx == 1 {
		x.Choice = new(ThresholdRSRQ)
	} else if option_idx == 2 {
		x.Choice = new(ProtocolIESingleContainerMeasurementThresholdL1LoggedMDTExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
