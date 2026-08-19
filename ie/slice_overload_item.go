package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SliceOverloadItem struct {
	SNSSAI       *SNSSAI                                            // valueExt
	IEExtensions *ProtocolExtensionContainerSliceOverloadItemExtIEs // optional
}

func (x *SliceOverloadItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SliceOverloadItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SNSSAI == nil {
		return errors.Errorf("SNSSAI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SliceOverloadItemOptPresentFlag = append(SliceOverloadItemOptPresentFlag, true)
	} else {
		SliceOverloadItemOptPresentFlag = append(SliceOverloadItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SliceOverloadItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SNSSAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SNSSAI marshal failed")
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

func (x *SliceOverloadItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SliceOverloadItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SliceOverloadItemOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if SliceOverloadItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSliceOverloadItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
