package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TAIListForPagingItem struct {
	TAI          *TAI                                                  // valueExt
	IEExtensions *ProtocolExtensionContainerTAIListForPagingItemExtIEs // optional
}

func (x *TAIListForPagingItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TAIListForPagingItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TAI == nil {
		return errors.Errorf("TAI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TAIListForPagingItemOptPresentFlag = append(TAIListForPagingItemOptPresentFlag, true)
	} else {
		TAIListForPagingItemOptPresentFlag = append(TAIListForPagingItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TAIListForPagingItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAI marshal failed")
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

func (x *TAIListForPagingItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TAIListForPagingItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TAIListForPagingItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAI = new(TAI)
	err = x.TAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAI error")
	}

	// optional field (optPresentFlag index: 0)
	if TAIListForPagingItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTAIListForPagingItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
