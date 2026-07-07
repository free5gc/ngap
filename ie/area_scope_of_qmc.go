package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AreaScopeOfQMC struct {
	Choice AreaScopeOfQMCAlt
}

type AreaScopeOfQMCAlt interface {
	AreaScopeOfQMCAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 CellBasedQMC) AreaScopeOfQMCAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 TABasedQMC) AreaScopeOfQMCAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 TAIBasedQMC) AreaScopeOfQMCAltIndex() int64 {
	return int64(2)
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 PLMNAreaBasedQMC) AreaScopeOfQMCAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 ProtocolIESingleContainerAreaScopeOfQMCExtIEs) AreaScopeOfQMCAltIndex() int64 {
	return int64(4)
}

// Choice Type Read/Write Functions

func (x *AreaScopeOfQMC) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 4
	var option_idx int64 = x.Choice.AreaScopeOfQMCAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *AreaScopeOfQMC) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(CellBasedQMC)
	} else if option_idx == 1 {
		x.Choice = new(TABasedQMC)
	} else if option_idx == 2 {
		x.Choice = new(TAIBasedQMC)
	} else if option_idx == 3 {
		x.Choice = new(PLMNAreaBasedQMC)
	} else if option_idx == 4 {
		x.Choice = new(ProtocolIESingleContainerAreaScopeOfQMCExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
