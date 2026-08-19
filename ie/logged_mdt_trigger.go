package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LoggedMDTTrigger struct {
	Choice LoggedMDTTriggerAlt
}

type LoggedMDTTriggerAlt interface {
	LoggedMDTTriggerAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PeriodicalForLoggedMDTTrigger struct {
	Value aper.NULL
}

func (alt0 *PeriodicalForLoggedMDTTrigger) LoggedMDTTriggerAltIndex() int64 {
	return int64(0)
}

func (x *PeriodicalForLoggedMDTTrigger) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

func (x *PeriodicalForLoggedMDTTrigger) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 EventTrigger) LoggedMDTTriggerAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 ProtocolIESingleContainerLoggedMDTTriggerExtIEs) LoggedMDTTriggerAltIndex() int64 {
	return int64(2)
}

// Choice Type Read/Write Functions

func (x *LoggedMDTTrigger) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64 = x.Choice.LoggedMDTTriggerAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *LoggedMDTTrigger) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(PeriodicalForLoggedMDTTrigger)
	} else if option_idx == 1 {
		x.Choice = new(EventTrigger)
	} else if option_idx == 2 {
		x.Choice = new(ProtocolIESingleContainerLoggedMDTTriggerExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
