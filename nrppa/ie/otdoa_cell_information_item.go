package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type OTDOACellInformationItem struct {
	Choice OTDOACellInformationItemAlt
}

type OTDOACellInformationItemAlt interface {
	OTDOACellInformationItemAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 PCIEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 CGIEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 TAC) OTDOACellInformationItemAltIndex() int64 {
	return int64(2)
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 EARFCN) OTDOACellInformationItemAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 PRSBandwidthEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(4)
}

// Choice type and its Read/Write is defined elsewhere
func (alt5 PRSConfigurationIndexEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(5)
}

// Choice type and its Read/Write is defined elsewhere
func (alt6 CPLengthEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(6)
}

// Choice type and its Read/Write is defined elsewhere
func (alt7 NumberOfDlFramesEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(7)
}

// Choice type and its Read/Write is defined elsewhere
func (alt8 NumberOfAntennaPortsEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(8)
}

// Choice type and its Read/Write is defined elsewhere
func (alt9 SFNInitialisationTimeEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(9)
}

// Choice type and its Read/Write is defined elsewhere
func (alt10 NGRANAccessPointPosition) OTDOACellInformationItemAltIndex() int64 {
	return int64(10)
}

// Choice type and its Read/Write is defined elsewhere
func (alt11 PRSMutingConfigurationEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(11)
}

// Choice type and its Read/Write is defined elsewhere
func (alt12 PRSIDEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(12)
}

// Choice type and its Read/Write is defined elsewhere
func (alt13 TPIDEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(13)
}

// Choice type and its Read/Write is defined elsewhere
func (alt14 TPTypeEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(14)
}

// Choice type and its Read/Write is defined elsewhere
func (alt15 NumberOfDlFramesExtendedEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(15)
}

// Choice type and its Read/Write is defined elsewhere
type CrsCPLengthEUTRA CPLengthEUTRA

func (alt16 CrsCPLengthEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(16)
}

// Choice type and its Read/Write is defined elsewhere
func (alt17 DLBandwidthEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(17)
}

// Choice type and its Read/Write is defined elsewhere
func (alt18 PRSOccasionGroupEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(18)
}

// Choice type and its Read/Write is defined elsewhere
func (alt19 PRSFrequencyHoppingConfigurationEUTRA) OTDOACellInformationItemAltIndex() int64 {
	return int64(19)
}

// Choice type and its Read/Write is defined elsewhere
func (alt20 ProtocolIESingleContainerOTDOACellInformationItemExtensionIE) OTDOACellInformationItemAltIndex() int64 {
	return int64(20)
}

// Choice Type Read/Write Functions

func (x *OTDOACellInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 20
	var option_idx int64 = x.Choice.OTDOACellInformationItemAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *OTDOACellInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 20
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(PCIEUTRA)
	} else if option_idx == 1 {
		x.Choice = new(CGIEUTRA)
	} else if option_idx == 2 {
		x.Choice = new(TAC)
	} else if option_idx == 3 {
		x.Choice = new(EARFCN)
	} else if option_idx == 4 {
		x.Choice = new(PRSBandwidthEUTRA)
	} else if option_idx == 5 {
		x.Choice = new(PRSConfigurationIndexEUTRA)
	} else if option_idx == 6 {
		x.Choice = new(CPLengthEUTRA)
	} else if option_idx == 7 {
		x.Choice = new(NumberOfDlFramesEUTRA)
	} else if option_idx == 8 {
		x.Choice = new(NumberOfAntennaPortsEUTRA)
	} else if option_idx == 9 {
		x.Choice = new(SFNInitialisationTimeEUTRA)
	} else if option_idx == 10 {
		x.Choice = new(NGRANAccessPointPosition)
	} else if option_idx == 11 {
		x.Choice = new(PRSMutingConfigurationEUTRA)
	} else if option_idx == 12 {
		x.Choice = new(PRSIDEUTRA)
	} else if option_idx == 13 {
		x.Choice = new(TPIDEUTRA)
	} else if option_idx == 14 {
		x.Choice = new(TPTypeEUTRA)
	} else if option_idx == 15 {
		x.Choice = new(NumberOfDlFramesExtendedEUTRA)
	} else if option_idx == 16 {
		x.Choice = new(CPLengthEUTRA)
	} else if option_idx == 17 {
		x.Choice = new(DLBandwidthEUTRA)
	} else if option_idx == 18 {
		x.Choice = new(PRSOccasionGroupEUTRA)
	} else if option_idx == 19 {
		x.Choice = new(PRSFrequencyHoppingConfigurationEUTRA)
	} else if option_idx == 20 {
		x.Choice = new(ProtocolIESingleContainerOTDOACellInformationItemExtensionIE)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
