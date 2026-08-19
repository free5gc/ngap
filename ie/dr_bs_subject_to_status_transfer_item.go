package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type DRBsSubjectToStatusTransferItem struct {
	DRBID       *DRBID
	DRBStatusUL *DRBStatusUL                                                     // valueLB:0,valueUB:2
	DRBStatusDL *DRBStatusDL                                                     // valueLB:0,valueUB:2
	IEExtension *ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs // optional
}

func (x *DRBsSubjectToStatusTransferItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DRBsSubjectToStatusTransferItemOptPresentFlag := []bool{}
	// mandatory field
	if x.DRBID == nil {
		return errors.Errorf("DRBID is missing")
	}
	// mandatory field
	if x.DRBStatusUL == nil {
		return errors.Errorf("DRBStatusUL is missing")
	}
	// mandatory field
	if x.DRBStatusDL == nil {
		return errors.Errorf("DRBStatusDL is missing")
	}
	// optional field
	if x.IEExtension != nil {
		DRBsSubjectToStatusTransferItemOptPresentFlag = append(DRBsSubjectToStatusTransferItemOptPresentFlag, true)
	} else {
		DRBsSubjectToStatusTransferItemOptPresentFlag = append(DRBsSubjectToStatusTransferItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DRBsSubjectToStatusTransferItemOptPresentFlag, true)
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
	err = x.DRBStatusUL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DRBStatusUL marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DRBStatusDL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DRBStatusDL marshal failed")
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

func (x *DRBsSubjectToStatusTransferItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DRBsSubjectToStatusTransferItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DRBsSubjectToStatusTransferItemOptPresentFlag, true)
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
	x.DRBStatusUL = new(DRBStatusUL)
	err = x.DRBStatusUL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DRBStatusUL error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DRBStatusDL = new(DRBStatusDL)
	err = x.DRBStatusDL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DRBStatusDL error")
	}

	// optional field (optPresentFlag index: 0)
	if DRBsSubjectToStatusTransferItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerDRBsSubjectToStatusTransferItemExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
