package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type N3IWFID struct {
	Choice N3IWFIDAlt
}

type N3IWFIDAlt interface {
	N3IWFIDAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type N3IWFIDForN3IWFID struct {
	Value aper.BitString
}

func (alt0 *N3IWFIDForN3IWFID) N3IWFIDAltIndex() int64 {
	return int64(0)
}

func (x *N3IWFIDForN3IWFID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 16, 16
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *N3IWFIDForN3IWFID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 16, 16
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 ProtocolIESingleContainerN3IWFIDExtIEs) N3IWFIDAltIndex() int64 {
	return int64(1)
}

// Choice Type Read/Write Functions

func (x *N3IWFID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 1
	var option_idx int64 = x.Choice.N3IWFIDAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *N3IWFID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 1
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(N3IWFIDForN3IWFID)
	} else if option_idx == 1 {
		x.Choice = new(ProtocolIESingleContainerN3IWFIDExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
