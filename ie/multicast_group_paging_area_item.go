package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MulticastGroupPagingAreaItem struct {
	MulticastGroupPagingArea *MulticastGroupPagingArea                                     // valueExt
	UEPagingList             *UEPagingList                                                 // optional
	IEExtensions             *ProtocolExtensionContainerMulticastGroupPagingAreaItemExtIEs // optional
}

func (x *MulticastGroupPagingAreaItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MulticastGroupPagingAreaItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MulticastGroupPagingArea == nil {
		return errors.Errorf("MulticastGroupPagingArea is missing")
	}
	// optional field
	if x.UEPagingList != nil {
		MulticastGroupPagingAreaItemOptPresentFlag = append(MulticastGroupPagingAreaItemOptPresentFlag, true)
	} else {
		MulticastGroupPagingAreaItemOptPresentFlag = append(MulticastGroupPagingAreaItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MulticastGroupPagingAreaItemOptPresentFlag = append(MulticastGroupPagingAreaItemOptPresentFlag, true)
	} else {
		MulticastGroupPagingAreaItemOptPresentFlag = append(MulticastGroupPagingAreaItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MulticastGroupPagingAreaItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MulticastGroupPagingArea.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MulticastGroupPagingArea marshal failed")
	}

	// optional field
	if x.UEPagingList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UEPagingList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UEPagingList marshal failed")
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

func (x *MulticastGroupPagingAreaItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MulticastGroupPagingAreaItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MulticastGroupPagingAreaItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MulticastGroupPagingArea = new(MulticastGroupPagingArea)
	err = x.MulticastGroupPagingArea.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MulticastGroupPagingArea error")
	}

	// optional field (optPresentFlag index: 0)
	if MulticastGroupPagingAreaItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UEPagingList = new(UEPagingList)
		err = x.UEPagingList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UEPagingList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MulticastGroupPagingAreaItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMulticastGroupPagingAreaItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
