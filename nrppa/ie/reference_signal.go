package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ReferenceSignal struct {
	Choice ReferenceSignalAlt
}

type ReferenceSignalAlt interface {
	ReferenceSignalAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 NZPCSIRSResourceID) ReferenceSignalAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 SSB) ReferenceSignalAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 SRSResourceID) ReferenceSignalAltIndex() int64 {
	return int64(2)
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 SRSPosResourceID) ReferenceSignalAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 DLPRS) ReferenceSignalAltIndex() int64 {
	return int64(4)
}

// Choice type and its Read/Write is defined elsewhere
func (alt5 ProtocolIESingleContainerReferenceSignalExtensionIE) ReferenceSignalAltIndex() int64 {
	return int64(5)
}

// Choice Type Read/Write Functions

func (x *ReferenceSignal) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 5
	var option_idx int64 = x.Choice.ReferenceSignalAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *ReferenceSignal) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 5
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(NZPCSIRSResourceID)
	} else if option_idx == 1 {
		x.Choice = new(SSB)
	} else if option_idx == 2 {
		x.Choice = new(SRSResourceID)
	} else if option_idx == 3 {
		x.Choice = new(SRSPosResourceID)
	} else if option_idx == 4 {
		x.Choice = new(DLPRS)
	} else if option_idx == 5 {
		x.Choice = new(ProtocolIESingleContainerReferenceSignalExtensionIE)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
