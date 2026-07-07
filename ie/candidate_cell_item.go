package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CandidateCellItem struct {
	CandidateCell *CandidateCell                                     // valueLB:0,valueUB:2
	IEExtensions  *ProtocolExtensionContainerCandidateCellItemExtIEs // optional
}

func (x *CandidateCellItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CandidateCellItemOptPresentFlag := []bool{}
	// mandatory field
	if x.CandidateCell == nil {
		return errors.Errorf("CandidateCell is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CandidateCellItemOptPresentFlag = append(CandidateCellItemOptPresentFlag, true)
	} else {
		CandidateCellItemOptPresentFlag = append(CandidateCellItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CandidateCellItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.CandidateCell.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CandidateCell marshal failed")
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

func (x *CandidateCellItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CandidateCellItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CandidateCellItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CandidateCell = new(CandidateCell)
	err = x.CandidateCell.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CandidateCell error")
	}

	// optional field (optPresentFlag index: 0)
	if CandidateCellItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCandidateCellItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
