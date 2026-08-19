package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI     *EUTRACGI                                                     // valueExt
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs // optional
}

func (x *CompletedCellsInEAIEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CompletedCellsInEAIEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.EUTRACGI == nil {
		return errors.Errorf("EUTRACGI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CompletedCellsInEAIEUTRAItemOptPresentFlag = append(CompletedCellsInEAIEUTRAItemOptPresentFlag, true)
	} else {
		CompletedCellsInEAIEUTRAItemOptPresentFlag = append(CompletedCellsInEAIEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CompletedCellsInEAIEUTRAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRACGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRACGI marshal failed")
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

func (x *CompletedCellsInEAIEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CompletedCellsInEAIEUTRAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CompletedCellsInEAIEUTRAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRACGI = new(EUTRACGI)
	err = x.EUTRACGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRACGI error")
	}

	// optional field (optPresentFlag index: 0)
	if CompletedCellsInEAIEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
