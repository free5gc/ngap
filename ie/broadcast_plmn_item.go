package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type BroadcastPLMNItem struct {
	PLMNIdentity        *PLMNIdentity
	TAISliceSupportList *SliceSupportList
	IEExtensions        *ProtocolExtensionContainerBroadcastPLMNItemExtIEs // optional
}

func (x *BroadcastPLMNItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	BroadcastPLMNItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.TAISliceSupportList == nil {
		return errors.Errorf("TAISliceSupportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		BroadcastPLMNItemOptPresentFlag = append(BroadcastPLMNItemOptPresentFlag, true)
	} else {
		BroadcastPLMNItemOptPresentFlag = append(BroadcastPLMNItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(BroadcastPLMNItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TAISliceSupportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAISliceSupportList marshal failed")
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

func (x *BroadcastPLMNItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	BroadcastPLMNItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&BroadcastPLMNItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNIdentity = new(PLMNIdentity)
	err = x.PLMNIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNIdentity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAISliceSupportList = new(SliceSupportList)
	err = x.TAISliceSupportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAISliceSupportList error")
	}

	// optional field (optPresentFlag index: 0)
	if BroadcastPLMNItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerBroadcastPLMNItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
