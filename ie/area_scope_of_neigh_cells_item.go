package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AreaScopeOfNeighCellsItem struct {
	NrFrequencyInfo *NRFrequencyInfo                                           // valueExt
	PciListForMDT   *PCIListForMDT                                             // optional
	IEExtensions    *ProtocolExtensionContainerAreaScopeOfNeighCellsItemExtIEs // optional
}

func (x *AreaScopeOfNeighCellsItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AreaScopeOfNeighCellsItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NrFrequencyInfo == nil {
		return errors.Errorf("NrFrequencyInfo is missing")
	}
	// optional field
	if x.PciListForMDT != nil {
		AreaScopeOfNeighCellsItemOptPresentFlag = append(AreaScopeOfNeighCellsItemOptPresentFlag, true)
	} else {
		AreaScopeOfNeighCellsItemOptPresentFlag = append(AreaScopeOfNeighCellsItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AreaScopeOfNeighCellsItemOptPresentFlag = append(AreaScopeOfNeighCellsItemOptPresentFlag, true)
	} else {
		AreaScopeOfNeighCellsItemOptPresentFlag = append(AreaScopeOfNeighCellsItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AreaScopeOfNeighCellsItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NrFrequencyInfo.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NrFrequencyInfo marshal failed")
	}

	// optional field
	if x.PciListForMDT != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PciListForMDT.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PciListForMDT marshal failed")
		}
	}

	// optional field
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *AreaScopeOfNeighCellsItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AreaScopeOfNeighCellsItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&AreaScopeOfNeighCellsItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NrFrequencyInfo = new(NRFrequencyInfo)
	err = x.NrFrequencyInfo.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NrFrequencyInfo error")
	}

	// optional field (optPresentFlag index: 0)
	if AreaScopeOfNeighCellsItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PciListForMDT = new(PCIListForMDT)
		err = x.PciListForMDT.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PciListForMDT error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AreaScopeOfNeighCellsItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAreaScopeOfNeighCellsItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
