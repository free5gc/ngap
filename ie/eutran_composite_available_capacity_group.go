package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EUTRANCompositeAvailableCapacityGroup struct {
	DLCompositeAvailableCapacity *CompositeAvailableCapacity                                            // valueExt
	ULCompositeAvailableCapacity *CompositeAvailableCapacity                                            // valueExt
	IEExtensions                 *ProtocolExtensionContainerEUTRANCompositeAvailableCapacityGroupExtIEs // optional
}

func (x *EUTRANCompositeAvailableCapacityGroup) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EUTRANCompositeAvailableCapacityGroupOptPresentFlag := []bool{}
	// mandatory field
	if x.DLCompositeAvailableCapacity == nil {
		return errors.Errorf("DLCompositeAvailableCapacity is missing")
	}
	// mandatory field
	if x.ULCompositeAvailableCapacity == nil {
		return errors.Errorf("ULCompositeAvailableCapacity is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EUTRANCompositeAvailableCapacityGroupOptPresentFlag = append(EUTRANCompositeAvailableCapacityGroupOptPresentFlag, true)
	} else {
		EUTRANCompositeAvailableCapacityGroupOptPresentFlag = append(EUTRANCompositeAvailableCapacityGroupOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EUTRANCompositeAvailableCapacityGroupOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DLCompositeAvailableCapacity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLCompositeAvailableCapacity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ULCompositeAvailableCapacity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ULCompositeAvailableCapacity marshal failed")
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

func (x *EUTRANCompositeAvailableCapacityGroup) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EUTRANCompositeAvailableCapacityGroupOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EUTRANCompositeAvailableCapacityGroupOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLCompositeAvailableCapacity = new(CompositeAvailableCapacity)
	err = x.DLCompositeAvailableCapacity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLCompositeAvailableCapacity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ULCompositeAvailableCapacity = new(CompositeAvailableCapacity)
	err = x.ULCompositeAvailableCapacity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ULCompositeAvailableCapacity error")
	}

	// optional field (optPresentFlag index: 0)
	if EUTRANCompositeAvailableCapacityGroupOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEUTRANCompositeAvailableCapacityGroupExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
