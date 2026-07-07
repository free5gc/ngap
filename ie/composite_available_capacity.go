package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CompositeAvailableCapacity struct {
	CellCapacityClassValue *int64                                                      // valueExt,valueLB:1,valueUB:100,optional
	CapacityValue          *int64                                                      // valueLB:0,valueUB:100
	IEExtensions           *ProtocolExtensionContainerCompositeAvailableCapacityExtIEs // optional
}

func (x *CompositeAvailableCapacity) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CompositeAvailableCapacityOptPresentFlag := []bool{}
	// optional field
	if x.CellCapacityClassValue != nil {
		CompositeAvailableCapacityOptPresentFlag = append(CompositeAvailableCapacityOptPresentFlag, true)
	} else {
		CompositeAvailableCapacityOptPresentFlag = append(CompositeAvailableCapacityOptPresentFlag, false)
	}
	// mandatory field
	if x.CapacityValue == nil {
		return errors.Errorf("CapacityValue is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CompositeAvailableCapacityOptPresentFlag = append(CompositeAvailableCapacityOptPresentFlag, true)
	} else {
		CompositeAvailableCapacityOptPresentFlag = append(CompositeAvailableCapacityOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CompositeAvailableCapacityOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.CellCapacityClassValue != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 100
		err = pd.WriteInteger(*(x.CellCapacityClassValue), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.CapacityValue), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *CompositeAvailableCapacity) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CompositeAvailableCapacityOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&CompositeAvailableCapacityOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if CompositeAvailableCapacityOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 100
		x.CellCapacityClassValue = new(int64)
		*(x.CellCapacityClassValue), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.CapacityValue = new(int64)
	*(x.CapacityValue), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 1)
	if CompositeAvailableCapacityOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCompositeAvailableCapacityExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
