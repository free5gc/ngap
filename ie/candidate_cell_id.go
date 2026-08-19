package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CandidateCellID struct {
	CandidateCellID *NRCGI                                           // valueExt
	IEExtensions    *ProtocolExtensionContainerCandidateCellIDExtIEs // optional
}

func (x *CandidateCellID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CandidateCellIDOptPresentFlag := []bool{}
	// mandatory field
	if x.CandidateCellID == nil {
		return errors.Errorf("CandidateCellID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CandidateCellIDOptPresentFlag = append(CandidateCellIDOptPresentFlag, true)
	} else {
		CandidateCellIDOptPresentFlag = append(CandidateCellIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CandidateCellIDOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.CandidateCellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CandidateCellID marshal failed")
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

func (x *CandidateCellID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CandidateCellIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CandidateCellIDOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CandidateCellID = new(NRCGI)
	err = x.CandidateCellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CandidateCellID error")
	}

	// optional field (optPresentFlag index: 0)
	if CandidateCellIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCandidateCellIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
