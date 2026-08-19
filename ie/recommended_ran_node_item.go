package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type RecommendedRANNodeItem struct {
	AMFPagingTarget *AMFPagingTarget                                        // valueLB:0,valueUB:2
	IEExtensions    *ProtocolExtensionContainerRecommendedRANNodeItemExtIEs // optional
}

func (x *RecommendedRANNodeItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RecommendedRANNodeItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AMFPagingTarget == nil {
		return errors.Errorf("AMFPagingTarget is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RecommendedRANNodeItemOptPresentFlag = append(RecommendedRANNodeItemOptPresentFlag, true)
	} else {
		RecommendedRANNodeItemOptPresentFlag = append(RecommendedRANNodeItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RecommendedRANNodeItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AMFPagingTarget.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AMFPagingTarget marshal failed")
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

func (x *RecommendedRANNodeItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RecommendedRANNodeItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RecommendedRANNodeItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AMFPagingTarget = new(AMFPagingTarget)
	err = x.AMFPagingTarget.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AMFPagingTarget error")
	}

	// optional field (optPresentFlag index: 0)
	if RecommendedRANNodeItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRecommendedRANNodeItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
