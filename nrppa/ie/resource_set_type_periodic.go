package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceSetTypePeriodicPeriodicSetPresentTrue aper.Enumerated = 0
)

type ResourceSetTypePeriodic struct {
	PeriodicSet  *aper.Enumerated                                         // valueExt,valueLB:0,valueUB:0
	IEExtensions *ProtocolExtensionContainerResourceSetTypePeriodicExtIEs // optional
}

func (x *ResourceSetTypePeriodic) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypePeriodicOptPresentFlag := []bool{}
	// mandatory field
	if x.PeriodicSet == nil {
		return errors.Errorf("PeriodicSet is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResourceSetTypePeriodicOptPresentFlag = append(ResourceSetTypePeriodicOptPresentFlag, true)
	} else {
		ResourceSetTypePeriodicOptPresentFlag = append(ResourceSetTypePeriodicOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypePeriodicOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.PeriodicSet), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *ResourceSetTypePeriodic) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypePeriodicOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypePeriodicOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.PeriodicSet = new(aper.Enumerated)
	*(x.PeriodicSet), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceSetTypePeriodicOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceSetTypePeriodicExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
