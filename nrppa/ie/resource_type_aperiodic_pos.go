package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResourceTypeAperiodicPos struct {
	SlotOffset   *int64                                                    // valueLB:0,valueUB:32
	IEExtensions *ProtocolExtensionContainerResourceTypeAperiodicPosExtIEs // optional
}

func (x *ResourceTypeAperiodicPos) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeAperiodicPosOptPresentFlag := []bool{}
	// mandatory field
	if x.SlotOffset == nil {
		return errors.Errorf("SlotOffset is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResourceTypeAperiodicPosOptPresentFlag = append(ResourceTypeAperiodicPosOptPresentFlag, true)
	} else {
		ResourceTypeAperiodicPosOptPresentFlag = append(ResourceTypeAperiodicPosOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeAperiodicPosOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 32
	err = pd.WriteInteger(*(x.SlotOffset), false, vLb, vUb)
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

func (x *ResourceTypeAperiodicPos) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeAperiodicPosOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeAperiodicPosOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 32
	x.SlotOffset = new(int64)
	*(x.SlotOffset), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceTypeAperiodicPosOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceTypeAperiodicPosExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
