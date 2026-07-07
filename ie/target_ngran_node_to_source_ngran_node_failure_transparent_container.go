package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer struct {
	CellCAGInformation *CellCAGInformation                                                                          // valueExt,optional
	IEExtensions       *ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEs // optional
}

func (x *TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag := []bool{}
	// optional field
	if x.CellCAGInformation != nil {
		TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag = append(TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag, true)
	} else {
		TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag = append(TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag = append(TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag, true)
	} else {
		TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag = append(TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.CellCAGInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CellCAGInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CellCAGInformation marshal failed")
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

func (x *TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CellCAGInformation = new(CellCAGInformation)
		err = x.CellCAGInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CellCAGInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeFailureTransparentContainerExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
