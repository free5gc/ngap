package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TargetRANNodeIDSON struct {
	GlobalRANNodeID *GlobalRANNodeID                                    // valueLB:0,valueUB:3
	SelectedTAI     *TAI                                                // valueExt
	IEExtensions    *ProtocolExtensionContainerTargetRANNodeIDSONExtIEs // optional
}

func (x *TargetRANNodeIDSON) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TargetRANNodeIDSONOptPresentFlag := []bool{}
	// mandatory field
	if x.GlobalRANNodeID == nil {
		return errors.Errorf("GlobalRANNodeID is missing")
	}
	// mandatory field
	if x.SelectedTAI == nil {
		return errors.Errorf("SelectedTAI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TargetRANNodeIDSONOptPresentFlag = append(TargetRANNodeIDSONOptPresentFlag, true)
	} else {
		TargetRANNodeIDSONOptPresentFlag = append(TargetRANNodeIDSONOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TargetRANNodeIDSONOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.GlobalRANNodeID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GlobalRANNodeID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SelectedTAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SelectedTAI marshal failed")
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

func (x *TargetRANNodeIDSON) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TargetRANNodeIDSONOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TargetRANNodeIDSONOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GlobalRANNodeID = new(GlobalRANNodeID)
	err = x.GlobalRANNodeID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GlobalRANNodeID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SelectedTAI = new(TAI)
	err = x.SelectedTAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SelectedTAI error")
	}

	// optional field (optPresentFlag index: 0)
	if TargetRANNodeIDSONOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTargetRANNodeIDSONExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
