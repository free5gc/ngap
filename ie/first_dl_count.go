package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type FirstDLCount struct {
	DRBsSubjectToEarlyStatusTransfer *DRBsSubjectToEarlyStatusTransferList
	IEExtension                      *ProtocolExtensionContainerFirstDLCountExtIEs // optional
}

func (x *FirstDLCount) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FirstDLCountOptPresentFlag := []bool{}
	// mandatory field
	if x.DRBsSubjectToEarlyStatusTransfer == nil {
		return errors.Errorf("DRBsSubjectToEarlyStatusTransfer is missing")
	}
	// optional field
	if x.IEExtension != nil {
		FirstDLCountOptPresentFlag = append(FirstDLCountOptPresentFlag, true)
	} else {
		FirstDLCountOptPresentFlag = append(FirstDLCountOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FirstDLCountOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DRBsSubjectToEarlyStatusTransfer.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DRBsSubjectToEarlyStatusTransfer marshal failed")
	}

	// optional field
	if x.IEExtension != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtension.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtension marshal failed")
		}
	}

	return nil
}

func (x *FirstDLCount) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FirstDLCountOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&FirstDLCountOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DRBsSubjectToEarlyStatusTransfer = new(DRBsSubjectToEarlyStatusTransferList)
	err = x.DRBsSubjectToEarlyStatusTransfer.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DRBsSubjectToEarlyStatusTransfer error")
	}

	// optional field (optPresentFlag index: 0)
	if FirstDLCountOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerFirstDLCountExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
