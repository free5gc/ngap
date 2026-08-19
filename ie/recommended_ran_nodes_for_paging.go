package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList *RecommendedRANNodeList
	IEExtensions           *ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs // optional
}

func (x *RecommendedRANNodesForPaging) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RecommendedRANNodesForPagingOptPresentFlag := []bool{}
	// mandatory field
	if x.RecommendedRANNodeList == nil {
		return errors.Errorf("RecommendedRANNodeList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RecommendedRANNodesForPagingOptPresentFlag = append(RecommendedRANNodesForPagingOptPresentFlag, true)
	} else {
		RecommendedRANNodesForPagingOptPresentFlag = append(RecommendedRANNodesForPagingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RecommendedRANNodesForPagingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RecommendedRANNodeList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RecommendedRANNodeList marshal failed")
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

func (x *RecommendedRANNodesForPaging) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RecommendedRANNodesForPagingOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RecommendedRANNodesForPagingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RecommendedRANNodeList = new(RecommendedRANNodeList)
	err = x.RecommendedRANNodeList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RecommendedRANNodeList error")
	}

	// optional field (optPresentFlag index: 0)
	if RecommendedRANNodesForPagingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
