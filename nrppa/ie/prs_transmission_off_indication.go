package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSTransmissionOffIndication struct {
	Choice PRSTransmissionOffIndicationAlt
}

type PRSTransmissionOffIndicationAlt interface {
	PRSTransmissionOffIndicationAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PRSTransmissionOffPerTRPForPRSTransmissionOffIndication struct {
	Value aper.NULL
}

func (alt0 *PRSTransmissionOffPerTRPForPRSTransmissionOffIndication) PRSTransmissionOffIndicationAltIndex() int64 {
	return int64(0)
}

func (x *PRSTransmissionOffPerTRPForPRSTransmissionOffIndication) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

func (x *PRSTransmissionOffPerTRPForPRSTransmissionOffIndication) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 PRSTransmissionOffPerResourceSet) PRSTransmissionOffIndicationAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 PRSTransmissionOffPerResource) PRSTransmissionOffIndicationAltIndex() int64 {
	return int64(2)
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 ProtocolIESingleContainerPRSTransmissionOffIndicationExtIEs) PRSTransmissionOffIndicationAltIndex() int64 {
	return int64(3)
}

// Choice Type Read/Write Functions

func (x *PRSTransmissionOffIndication) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 3
	var option_idx int64 = x.Choice.PRSTransmissionOffIndicationAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *PRSTransmissionOffIndication) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(PRSTransmissionOffPerTRPForPRSTransmissionOffIndication)
	} else if option_idx == 1 {
		x.Choice = new(PRSTransmissionOffPerResourceSet)
	} else if option_idx == 2 {
		x.Choice = new(PRSTransmissionOffPerResource)
	} else if option_idx == 3 {
		x.Choice = new(ProtocolIESingleContainerPRSTransmissionOffIndicationExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
