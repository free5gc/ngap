package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LastVisitedCellInformation struct {
	Choice LastVisitedCellInformationAlt
}

type LastVisitedCellInformationAlt interface {
	LastVisitedCellInformationAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 LastVisitedNGRANCellInformation) LastVisitedCellInformationAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 LastVisitedEUTRANCellInformation) LastVisitedCellInformationAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 LastVisitedUTRANCellInformation) LastVisitedCellInformationAltIndex() int64 {
	return int64(2)
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 LastVisitedGERANCellInformation) LastVisitedCellInformationAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 ProtocolIESingleContainerLastVisitedCellInformationExtIEs) LastVisitedCellInformationAltIndex() int64 {
	return int64(4)
}

// Choice Type Read/Write Functions

func (x *LastVisitedCellInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 4
	var option_idx int64 = x.Choice.LastVisitedCellInformationAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *LastVisitedCellInformation) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(LastVisitedNGRANCellInformation)
	} else if option_idx == 1 {
		x.Choice = new(LastVisitedEUTRANCellInformation)
	} else if option_idx == 2 {
		x.Choice = new(LastVisitedUTRANCellInformation)
	} else if option_idx == 3 {
		x.Choice = new(LastVisitedGERANCellInformation)
	} else if option_idx == 4 {
		x.Choice = new(ProtocolIESingleContainerLastVisitedCellInformationExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
