package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TAINSAGSupportItem struct {
	NSAGID               *NSAGID
	NSAGSliceSupportList *ExtendedSliceSupportList
	IEExtensions         *ProtocolExtensionContainerTAINSAGSupportItemExtIEs // optional
}

func (x *TAINSAGSupportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TAINSAGSupportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NSAGID == nil {
		return errors.Errorf("NSAGID is missing")
	}
	// mandatory field
	if x.NSAGSliceSupportList == nil {
		return errors.Errorf("NSAGSliceSupportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TAINSAGSupportItemOptPresentFlag = append(TAINSAGSupportItemOptPresentFlag, true)
	} else {
		TAINSAGSupportItemOptPresentFlag = append(TAINSAGSupportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TAINSAGSupportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NSAGID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NSAGID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NSAGSliceSupportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NSAGSliceSupportList marshal failed")
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

func (x *TAINSAGSupportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TAINSAGSupportItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TAINSAGSupportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NSAGID = new(NSAGID)
	err = x.NSAGID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NSAGID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NSAGSliceSupportList = new(ExtendedSliceSupportList)
	err = x.NSAGSliceSupportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NSAGSliceSupportList error")
	}

	// optional field (optPresentFlag index: 0)
	if TAINSAGSupportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTAINSAGSupportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
