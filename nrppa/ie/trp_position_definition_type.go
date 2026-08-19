package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPPositionDefinitionType struct {
	Choice TRPPositionDefinitionTypeAlt
}

type TRPPositionDefinitionTypeAlt interface {
	TRPPositionDefinitionTypeAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 TRPPositionDirect) TRPPositionDefinitionTypeAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 TRPPositionReferenced) TRPPositionDefinitionTypeAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 ProtocolIESingleContainerTRPPositionDefinitionTypeExtIEs) TRPPositionDefinitionTypeAltIndex() int64 {
	return int64(2)
}

// Choice Type Read/Write Functions

func (x *TRPPositionDefinitionType) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64 = x.Choice.TRPPositionDefinitionTypeAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *TRPPositionDefinitionType) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(TRPPositionDirect)
	} else if option_idx == 1 {
		x.Choice = new(TRPPositionReferenced)
	} else if option_idx == 2 {
		x.Choice = new(ProtocolIESingleContainerTRPPositionDefinitionTypeExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
