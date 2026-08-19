package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AreaScopeOfMDTEUTRA struct {
	Choice AreaScopeOfMDTEUTRAAlt
}

type AreaScopeOfMDTEUTRAAlt interface {
	AreaScopeOfMDTEUTRAAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 CellBasedMDTEUTRA) AreaScopeOfMDTEUTRAAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 TABasedMDT) AreaScopeOfMDTEUTRAAltIndex() int64 {
	return int64(1)
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PLMNWideForAreaScopeOfMDTEUTRA struct {
	Value aper.NULL
}

func (alt2 *PLMNWideForAreaScopeOfMDTEUTRA) AreaScopeOfMDTEUTRAAltIndex() int64 {
	return int64(2)
}

func (x *PLMNWideForAreaScopeOfMDTEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

func (x *PLMNWideForAreaScopeOfMDTEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Value: NULL type has no encoding bytes
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 TAIBasedMDT) AreaScopeOfMDTEUTRAAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 ProtocolIESingleContainerAreaScopeOfMDTEUTRAExtIEs) AreaScopeOfMDTEUTRAAltIndex() int64 {
	return int64(4)
}

// Choice Type Read/Write Functions

func (x *AreaScopeOfMDTEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 4
	var option_idx int64 = x.Choice.AreaScopeOfMDTEUTRAAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *AreaScopeOfMDTEUTRA) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(CellBasedMDTEUTRA)
	} else if option_idx == 1 {
		x.Choice = new(TABasedMDT)
	} else if option_idx == 2 {
		x.Choice = new(PLMNWideForAreaScopeOfMDTEUTRA)
	} else if option_idx == 3 {
		x.Choice = new(TAIBasedMDT)
	} else if option_idx == 4 {
		x.Choice = new(ProtocolIESingleContainerAreaScopeOfMDTEUTRAExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
