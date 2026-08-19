package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MulticastGroupPagingArea struct {
	MBSAreaTAIList *MBSAreaTAIList
	IEExtensions   *ProtocolExtensionContainerMulticastGroupPagingAreaExtIEs // optional
}

func (x *MulticastGroupPagingArea) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MulticastGroupPagingAreaOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSAreaTAIList == nil {
		return errors.Errorf("MBSAreaTAIList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		MulticastGroupPagingAreaOptPresentFlag = append(MulticastGroupPagingAreaOptPresentFlag, true)
	} else {
		MulticastGroupPagingAreaOptPresentFlag = append(MulticastGroupPagingAreaOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MulticastGroupPagingAreaOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSAreaTAIList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSAreaTAIList marshal failed")
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

func (x *MulticastGroupPagingArea) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MulticastGroupPagingAreaOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&MulticastGroupPagingAreaOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSAreaTAIList = new(MBSAreaTAIList)
	err = x.MBSAreaTAIList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSAreaTAIList error")
	}

	// optional field (optPresentFlag index: 0)
	if MulticastGroupPagingAreaOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMulticastGroupPagingAreaExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
