package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CellBasedMDTEUTRA struct {
	CellIdListforMDT *CellIdListforMDTEUTRA
	IEExtensions     *ProtocolExtensionContainerCellBasedMDTEUTRAExtIEs // optional
}

func (x *CellBasedMDTEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CellBasedMDTEUTRAOptPresentFlag := []bool{}
	// mandatory field
	if x.CellIdListforMDT == nil {
		return errors.Errorf("CellIdListforMDT is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CellBasedMDTEUTRAOptPresentFlag = append(CellBasedMDTEUTRAOptPresentFlag, true)
	} else {
		CellBasedMDTEUTRAOptPresentFlag = append(CellBasedMDTEUTRAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CellBasedMDTEUTRAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.CellIdListforMDT.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CellIdListforMDT marshal failed")
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

func (x *CellBasedMDTEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CellBasedMDTEUTRAOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CellBasedMDTEUTRAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CellIdListforMDT = new(CellIdListforMDTEUTRA)
	err = x.CellIdListforMDT.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CellIdListforMDT error")
	}

	// optional field (optPresentFlag index: 0)
	if CellBasedMDTEUTRAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCellBasedMDTEUTRAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
