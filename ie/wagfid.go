package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type WAGFID struct {
	Choice WAGFIDAlt
}

type WAGFIDAlt interface {
	WAGFIDAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type WAGFIDForWAGFID struct {
	Value aper.BitString
}

func (alt0 *WAGFIDForWAGFID) WAGFIDAltIndex() int64 {
	return int64(0)
}

func (x *WAGFIDForWAGFID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 16, 16
	err = pd.WriteBitString(x.Value, true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *WAGFIDForWAGFID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 16, 16
	x.Value, err = pd.ReadBitString(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 ProtocolIESingleContainerWAGFIDExtIEs) WAGFIDAltIndex() int64 {
	return int64(1)
}

// Choice Type Read/Write Functions

func (x *WAGFID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 1
	var option_idx int64 = x.Choice.WAGFIDAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *WAGFID) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(WAGFIDForWAGFID)
	} else if option_idx == 1 {
		x.Choice = new(ProtocolIESingleContainerWAGFIDExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
