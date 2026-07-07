package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type RecommendedCellsForPaging struct {
	RecommendedCellList *RecommendedCellList
	IEExtensions        *ProtocolExtensionContainerRecommendedCellsForPagingExtIEs // optional
}

func (x *RecommendedCellsForPaging) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RecommendedCellsForPagingOptPresentFlag := []bool{}
	// mandatory field
	if x.RecommendedCellList == nil {
		return errors.Errorf("RecommendedCellList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RecommendedCellsForPagingOptPresentFlag = append(RecommendedCellsForPagingOptPresentFlag, true)
	} else {
		RecommendedCellsForPagingOptPresentFlag = append(RecommendedCellsForPagingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RecommendedCellsForPagingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RecommendedCellList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RecommendedCellList marshal failed")
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

func (x *RecommendedCellsForPaging) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RecommendedCellsForPagingOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RecommendedCellsForPagingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RecommendedCellList = new(RecommendedCellList)
	err = x.RecommendedCellList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RecommendedCellList error")
	}

	// optional field (optPresentFlag index: 0)
	if RecommendedCellsForPagingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRecommendedCellsForPagingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
