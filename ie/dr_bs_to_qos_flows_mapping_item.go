package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type DRBsToQosFlowsMappingItem struct {
	DRBID                 *DRBID
	AssociatedQosFlowList *AssociatedQosFlowList
	IEExtensions          *ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs // optional
}

func (x *DRBsToQosFlowsMappingItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DRBsToQosFlowsMappingItemOptPresentFlag := []bool{}
	// mandatory field
	if x.DRBID == nil {
		return errors.Errorf("DRBID is missing")
	}
	// mandatory field
	if x.AssociatedQosFlowList == nil {
		return errors.Errorf("AssociatedQosFlowList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		DRBsToQosFlowsMappingItemOptPresentFlag = append(DRBsToQosFlowsMappingItemOptPresentFlag, true)
	} else {
		DRBsToQosFlowsMappingItemOptPresentFlag = append(DRBsToQosFlowsMappingItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DRBsToQosFlowsMappingItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DRBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DRBID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AssociatedQosFlowList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AssociatedQosFlowList marshal failed")
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

func (x *DRBsToQosFlowsMappingItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DRBsToQosFlowsMappingItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DRBsToQosFlowsMappingItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DRBID = new(DRBID)
	err = x.DRBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DRBID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AssociatedQosFlowList = new(AssociatedQosFlowList)
	err = x.AssociatedQosFlowList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AssociatedQosFlowList error")
	}

	// optional field (optPresentFlag index: 0)
	if DRBsToQosFlowsMappingItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
