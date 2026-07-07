package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEAppLayerMeasInfoItem struct {
	UEAppLayerMeasConfigInfo *UEAppLayerMeasConfigInfo                               // valueExt
	IEExtensions             *ProtocolExtensionContainerUEAppLayerMeasInfoItemExtIEs // optional
}

func (x *UEAppLayerMeasInfoItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEAppLayerMeasInfoItemOptPresentFlag := []bool{}
	// mandatory field
	if x.UEAppLayerMeasConfigInfo == nil {
		return errors.Errorf("UEAppLayerMeasConfigInfo is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UEAppLayerMeasInfoItemOptPresentFlag = append(UEAppLayerMeasInfoItemOptPresentFlag, true)
	} else {
		UEAppLayerMeasInfoItemOptPresentFlag = append(UEAppLayerMeasInfoItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEAppLayerMeasInfoItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.UEAppLayerMeasConfigInfo.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UEAppLayerMeasConfigInfo marshal failed")
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

func (x *UEAppLayerMeasInfoItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEAppLayerMeasInfoItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UEAppLayerMeasInfoItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UEAppLayerMeasConfigInfo = new(UEAppLayerMeasConfigInfo)
	err = x.UEAppLayerMeasConfigInfo.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UEAppLayerMeasConfigInfo error")
	}

	// optional field (optPresentFlag index: 0)
	if UEAppLayerMeasInfoItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEAppLayerMeasInfoItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
