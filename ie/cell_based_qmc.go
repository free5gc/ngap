package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CellBasedQMC struct {
	CellIdListforQMC *CellIdListforQMC
	IEExtensions     *ProtocolExtensionContainerCellBasedQMCExtIEs // optional
}

func (x *CellBasedQMC) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CellBasedQMCOptPresentFlag := []bool{}
	// mandatory field
	if x.CellIdListforQMC == nil {
		return errors.Errorf("CellIdListforQMC is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CellBasedQMCOptPresentFlag = append(CellBasedQMCOptPresentFlag, true)
	} else {
		CellBasedQMCOptPresentFlag = append(CellBasedQMCOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CellBasedQMCOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.CellIdListforQMC.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CellIdListforQMC marshal failed")
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

func (x *CellBasedQMC) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CellBasedQMCOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CellBasedQMCOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CellIdListforQMC = new(CellIdListforQMC)
	err = x.CellIdListforQMC.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CellIdListforQMC error")
	}

	// optional field (optPresentFlag index: 0)
	if CellBasedQMCOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCellBasedQMCExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
