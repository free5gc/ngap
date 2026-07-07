package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEPagingItem struct {
	UEIdentityIndexValue *UEIdentityIndexValue                         // valueLB:0,valueUB:1
	PagingDRX            *PagingDRX                                    // valueExt,valueLB:0,valueUB:3,optional
	IEExtensions         *ProtocolExtensionContainerUEPagingItemExtIEs // optional
}

func (x *UEPagingItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEPagingItemOptPresentFlag := []bool{}
	// mandatory field
	if x.UEIdentityIndexValue == nil {
		return errors.Errorf("UEIdentityIndexValue is missing")
	}
	// optional field
	if x.PagingDRX != nil {
		UEPagingItemOptPresentFlag = append(UEPagingItemOptPresentFlag, true)
	} else {
		UEPagingItemOptPresentFlag = append(UEPagingItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEPagingItemOptPresentFlag = append(UEPagingItemOptPresentFlag, true)
	} else {
		UEPagingItemOptPresentFlag = append(UEPagingItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEPagingItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.UEIdentityIndexValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UEIdentityIndexValue marshal failed")
	}

	// optional field
	if x.PagingDRX != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PagingDRX.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PagingDRX marshal failed")
		}
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

func (x *UEPagingItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEPagingItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UEPagingItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UEIdentityIndexValue = new(UEIdentityIndexValue)
	err = x.UEIdentityIndexValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UEIdentityIndexValue error")
	}

	// optional field (optPresentFlag index: 0)
	if UEPagingItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PagingDRX = new(PagingDRX)
		err = x.PagingDRX.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PagingDRX error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UEPagingItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEPagingItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
