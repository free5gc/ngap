package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UESliceMaximumBitRateItem struct {
	SNSSAI                  *SNSSAI // valueExt
	UESliceMaximumBitRateDL *BitRate
	UESliceMaximumBitRateUL *BitRate
	IEExtensions            *ProtocolExtensionContainerUESliceMaximumBitRateItemExtIEs // optional
}

func (x *UESliceMaximumBitRateItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UESliceMaximumBitRateItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SNSSAI == nil {
		return errors.Errorf("SNSSAI is missing")
	}
	// mandatory field
	if x.UESliceMaximumBitRateDL == nil {
		return errors.Errorf("UESliceMaximumBitRateDL is missing")
	}
	// mandatory field
	if x.UESliceMaximumBitRateUL == nil {
		return errors.Errorf("UESliceMaximumBitRateUL is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UESliceMaximumBitRateItemOptPresentFlag = append(UESliceMaximumBitRateItemOptPresentFlag, true)
	} else {
		UESliceMaximumBitRateItemOptPresentFlag = append(UESliceMaximumBitRateItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UESliceMaximumBitRateItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SNSSAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SNSSAI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.UESliceMaximumBitRateDL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UESliceMaximumBitRateDL marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.UESliceMaximumBitRateUL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UESliceMaximumBitRateUL marshal failed")
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

func (x *UESliceMaximumBitRateItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UESliceMaximumBitRateItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UESliceMaximumBitRateItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SNSSAI = new(SNSSAI)
	err = x.SNSSAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SNSSAI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UESliceMaximumBitRateDL = new(BitRate)
	err = x.UESliceMaximumBitRateDL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UESliceMaximumBitRateDL error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UESliceMaximumBitRateUL = new(BitRate)
	err = x.UESliceMaximumBitRateUL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UESliceMaximumBitRateUL error")
	}

	// optional field (optPresentFlag index: 0)
	if UESliceMaximumBitRateItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUESliceMaximumBitRateItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
