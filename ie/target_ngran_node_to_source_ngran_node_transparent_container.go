package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer *RRCContainer
	IEExtensions *ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs // optional
}

func (x *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag := []bool{}
	// mandatory field
	if x.RRCContainer == nil {
		return errors.Errorf("RRCContainer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag = append(TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag, true)
	} else {
		TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag = append(TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RRCContainer.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RRCContainer marshal failed")
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

func (x *TargetNGRANNodeToSourceNGRANNodeTransparentContainer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RRCContainer = new(RRCContainer)
	err = x.RRCContainer.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RRCContainer error")
	}

	// optional field (optPresentFlag index: 0)
	if TargetNGRANNodeToSourceNGRANNodeTransparentContainerOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
