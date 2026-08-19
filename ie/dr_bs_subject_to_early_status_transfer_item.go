package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type DRBsSubjectToEarlyStatusTransferItem struct {
	DRBID        *DRBID
	FirstDLCOUNT *DRBStatusDL                                                          // valueLB:0,valueUB:2
	IEExtension  *ProtocolExtensionContainerDRBsSubjectToEarlyStatusTransferItemExtIEs // optional
}

func (x *DRBsSubjectToEarlyStatusTransferItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DRBsSubjectToEarlyStatusTransferItemOptPresentFlag := []bool{}
	// mandatory field
	if x.DRBID == nil {
		return errors.Errorf("DRBID is missing")
	}
	// mandatory field
	if x.FirstDLCOUNT == nil {
		return errors.Errorf("FirstDLCOUNT is missing")
	}
	// optional field
	if x.IEExtension != nil {
		DRBsSubjectToEarlyStatusTransferItemOptPresentFlag = append(DRBsSubjectToEarlyStatusTransferItemOptPresentFlag, true)
	} else {
		DRBsSubjectToEarlyStatusTransferItemOptPresentFlag = append(DRBsSubjectToEarlyStatusTransferItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DRBsSubjectToEarlyStatusTransferItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DRBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DRBID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.FirstDLCOUNT.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FirstDLCOUNT marshal failed")
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

func (x *DRBsSubjectToEarlyStatusTransferItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DRBsSubjectToEarlyStatusTransferItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DRBsSubjectToEarlyStatusTransferItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DRBID = new(DRBID)
	err = x.DRBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DRBID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FirstDLCOUNT = new(DRBStatusDL)
	err = x.FirstDLCOUNT.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FirstDLCOUNT error")
	}

	// optional field (optPresentFlag index: 0)
	if DRBsSubjectToEarlyStatusTransferItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerDRBsSubjectToEarlyStatusTransferItemExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
