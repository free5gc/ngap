package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CompletedCellsInEAINRItem struct {
	NRCGI        *NRCGI                                                     // valueExt
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs // optional
}

func (x *CompletedCellsInEAINRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CompletedCellsInEAINRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NRCGI == nil {
		return errors.Errorf("NRCGI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CompletedCellsInEAINRItemOptPresentFlag = append(CompletedCellsInEAINRItemOptPresentFlag, true)
	} else {
		CompletedCellsInEAINRItemOptPresentFlag = append(CompletedCellsInEAINRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CompletedCellsInEAINRItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRCGI marshal failed")
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

func (x *CompletedCellsInEAINRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CompletedCellsInEAINRItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CompletedCellsInEAINRItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRCGI = new(NRCGI)
	err = x.NRCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRCGI error")
	}

	// optional field (optPresentFlag index: 0)
	if CompletedCellsInEAINRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
