package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AreaScopeOfMDTNR struct {
	Choice AreaScopeOfMDTNRAlt
}

type AreaScopeOfMDTNRAlt interface {
	AreaScopeOfMDTNRAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 CellBasedMDTNR) AreaScopeOfMDTNRAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 TABasedMDT) AreaScopeOfMDTNRAltIndex() int64 {
	return int64(1)
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PLMNWideForAreaScopeOfMDTNR struct {
	Value aper.NULL
}

func (alt2 *PLMNWideForAreaScopeOfMDTNR) AreaScopeOfMDTNRAltIndex() int64 {
	return int64(2)
}

func (x *PLMNWideForAreaScopeOfMDTNR) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

func (x *PLMNWideForAreaScopeOfMDTNR) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 TAIBasedMDT) AreaScopeOfMDTNRAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 ProtocolIESingleContainerAreaScopeOfMDTNRExtIEs) AreaScopeOfMDTNRAltIndex() int64 {
	return int64(4)
}

// Choice Type Read/Write Functions

func (x *AreaScopeOfMDTNR) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 4
	var option_idx int64 = x.Choice.AreaScopeOfMDTNRAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *AreaScopeOfMDTNR) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 4
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(CellBasedMDTNR)
	} else if option_idx == 1 {
		x.Choice = new(TABasedMDT)
	} else if option_idx == 2 {
		x.Choice = new(PLMNWideForAreaScopeOfMDTNR)
	} else if option_idx == 3 {
		x.Choice = new(TAIBasedMDT)
	} else if option_idx == 4 {
		x.Choice = new(ProtocolIESingleContainerAreaScopeOfMDTNRExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
