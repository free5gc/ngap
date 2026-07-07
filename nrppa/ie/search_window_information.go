package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SearchWindowInformation struct {
	ExpectedPropagationDelay *int64                                                   // valueExt,valueLB:-3841,valueUB:3841
	DelayUncertainty         *int64                                                   // valueExt,valueLB:1,valueUB:246
	IEExtensions             *ProtocolExtensionContainerSearchWindowInformationExtIEs // optional
}

func (x *SearchWindowInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SearchWindowInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.ExpectedPropagationDelay == nil {
		return errors.Errorf("ExpectedPropagationDelay is missing")
	}
	// mandatory field
	if x.DelayUncertainty == nil {
		return errors.Errorf("DelayUncertainty is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SearchWindowInformationOptPresentFlag = append(SearchWindowInformationOptPresentFlag, true)
	} else {
		SearchWindowInformationOptPresentFlag = append(SearchWindowInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SearchWindowInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = -3841, 3841
	err = pd.WriteInteger(*(x.ExpectedPropagationDelay), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 1, 246
	err = pd.WriteInteger(*(x.DelayUncertainty), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *SearchWindowInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SearchWindowInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SearchWindowInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -3841, 3841
	x.ExpectedPropagationDelay = new(int64)
	*(x.ExpectedPropagationDelay), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 1, 246
	x.DelayUncertainty = new(int64)
	*(x.DelayUncertainty), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if SearchWindowInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSearchWindowInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
