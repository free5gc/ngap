package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type DRBStatusDL12 struct {
	DLCOUNTValue *COUNTValueForPDCPSN12                         // valueExt
	IEExtension  *ProtocolExtensionContainerDRBStatusDL12ExtIEs // optional
}

func (x *DRBStatusDL12) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DRBStatusDL12OptPresentFlag := []bool{}
	// mandatory field
	if x.DLCOUNTValue == nil {
		return errors.Errorf("DLCOUNTValue is missing")
	}
	// optional field
	if x.IEExtension != nil {
		DRBStatusDL12OptPresentFlag = append(DRBStatusDL12OptPresentFlag, true)
	} else {
		DRBStatusDL12OptPresentFlag = append(DRBStatusDL12OptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DRBStatusDL12OptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DLCOUNTValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLCOUNTValue marshal failed")
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

func (x *DRBStatusDL12) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DRBStatusDL12OptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DRBStatusDL12OptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLCOUNTValue = new(COUNTValueForPDCPSN12)
	err = x.DLCOUNTValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLCOUNTValue error")
	}

	// optional field (optPresentFlag index: 0)
	if DRBStatusDL12OptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerDRBStatusDL12ExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
