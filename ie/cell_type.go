package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CellType struct {
	CellSize     *CellSize                                 // valueExt,valueLB:0,valueUB:3
	IEExtensions *ProtocolExtensionContainerCellTypeExtIEs // optional
}

func (x *CellType) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CellTypeOptPresentFlag := []bool{}
	// mandatory field
	if x.CellSize == nil {
		return errors.Errorf("CellSize is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CellTypeOptPresentFlag = append(CellTypeOptPresentFlag, true)
	} else {
		CellTypeOptPresentFlag = append(CellTypeOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CellTypeOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.CellSize.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CellSize marshal failed")
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

func (x *CellType) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CellTypeOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CellTypeOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CellSize = new(CellSize)
	err = x.CellSize.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CellSize error")
	}

	// optional field (optPresentFlag index: 0)
	if CellTypeOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCellTypeExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
