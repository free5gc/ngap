package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type DRBStatusUL12 struct {
	ULCOUNTValue              *COUNTValueForPDCPSN12                         // valueExt
	ReceiveStatusOfULPDCPSDUs *aper.BitString                                // sizeLB:1,sizeUB:2048,optional
	IEExtension               *ProtocolExtensionContainerDRBStatusUL12ExtIEs // optional
}

func (x *DRBStatusUL12) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DRBStatusUL12OptPresentFlag := []bool{}
	// mandatory field
	if x.ULCOUNTValue == nil {
		return errors.Errorf("ULCOUNTValue is missing")
	}
	// optional field
	if x.ReceiveStatusOfULPDCPSDUs != nil {
		DRBStatusUL12OptPresentFlag = append(DRBStatusUL12OptPresentFlag, true)
	} else {
		DRBStatusUL12OptPresentFlag = append(DRBStatusUL12OptPresentFlag, false)
	}
	// optional field
	if x.IEExtension != nil {
		DRBStatusUL12OptPresentFlag = append(DRBStatusUL12OptPresentFlag, true)
	} else {
		DRBStatusUL12OptPresentFlag = append(DRBStatusUL12OptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DRBStatusUL12OptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ULCOUNTValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ULCOUNTValue marshal failed")
	}

	// optional field
	if x.ReceiveStatusOfULPDCPSDUs != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 1, 2048
		err = pd.WriteBitString(*(x.ReceiveStatusOfULPDCPSDUs), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
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

func (x *DRBStatusUL12) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DRBStatusUL12OptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&DRBStatusUL12OptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ULCOUNTValue = new(COUNTValueForPDCPSN12)
	err = x.ULCOUNTValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ULCOUNTValue error")
	}

	// optional field (optPresentFlag index: 0)
	if DRBStatusUL12OptPresentFlag[0] {
		// Read BitString (Pointer)
		*sLb, *sUb = 1, 2048
		x.ReceiveStatusOfULPDCPSDUs = new(aper.BitString)
		*(x.ReceiveStatusOfULPDCPSDUs), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if DRBStatusUL12OptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerDRBStatusUL12ExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
