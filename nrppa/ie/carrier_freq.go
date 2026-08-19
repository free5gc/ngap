package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type CarrierFreq struct {
	PointA          *int64                                       // valueLB:0,valueUB:3279165
	OffsetToCarrier *int64                                       // valueExt,valueLB:0,valueUB:2199
	IEExtensions    *ProtocolExtensionContainerCarrierFreqExtIEs // optional
}

func (x *CarrierFreq) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CarrierFreqOptPresentFlag := []bool{}
	// mandatory field
	if x.PointA == nil {
		return errors.Errorf("PointA is missing")
	}
	// mandatory field
	if x.OffsetToCarrier == nil {
		return errors.Errorf("OffsetToCarrier is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CarrierFreqOptPresentFlag = append(CarrierFreqOptPresentFlag, true)
	} else {
		CarrierFreqOptPresentFlag = append(CarrierFreqOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CarrierFreqOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(*(x.PointA), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 2199
	err = pd.WriteInteger(*(x.OffsetToCarrier), true, vLb, vUb)
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

func (x *CarrierFreq) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CarrierFreqOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CarrierFreqOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	x.PointA = new(int64)
	*(x.PointA), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 2199
	x.OffsetToCarrier = new(int64)
	*(x.OffsetToCarrier), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if CarrierFreqOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCarrierFreqExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
