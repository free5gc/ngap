package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type IntersystemSONeNBID struct {
	GlobaleNBID    *GlobalENBID                                         // valueExt
	SelectedEPSTAI *EPSTAI                                              // valueExt
	IEExtensions   *ProtocolExtensionContainerIntersystemSONeNBIDExtIEs // optional
}

func (x *IntersystemSONeNBID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemSONeNBIDOptPresentFlag := []bool{}
	// mandatory field
	if x.GlobaleNBID == nil {
		return errors.Errorf("GlobaleNBID is missing")
	}
	// mandatory field
	if x.SelectedEPSTAI == nil {
		return errors.Errorf("SelectedEPSTAI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemSONeNBIDOptPresentFlag = append(IntersystemSONeNBIDOptPresentFlag, true)
	} else {
		IntersystemSONeNBIDOptPresentFlag = append(IntersystemSONeNBIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemSONeNBIDOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.GlobaleNBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GlobaleNBID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SelectedEPSTAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SelectedEPSTAI marshal failed")
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

func (x *IntersystemSONeNBID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemSONeNBIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemSONeNBIDOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GlobaleNBID = new(GlobalENBID)
	err = x.GlobaleNBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GlobaleNBID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SelectedEPSTAI = new(EPSTAI)
	err = x.SelectedEPSTAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SelectedEPSTAI error")
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemSONeNBIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemSONeNBIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
