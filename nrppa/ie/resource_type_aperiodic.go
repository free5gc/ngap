package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceTypeAperiodicAperiodicResourceTypePresentTrue aper.Enumerated = 0
)

type ResourceTypeAperiodic struct {
	AperiodicResourceType *aper.Enumerated                                       // valueExt,valueLB:0,valueUB:0
	IEExtensions          *ProtocolExtensionContainerResourceTypeAperiodicExtIEs // optional
}

func (x *ResourceTypeAperiodic) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeAperiodicOptPresentFlag := []bool{}
	// mandatory field
	if x.AperiodicResourceType == nil {
		return errors.Errorf("AperiodicResourceType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResourceTypeAperiodicOptPresentFlag = append(ResourceTypeAperiodicOptPresentFlag, true)
	} else {
		ResourceTypeAperiodicOptPresentFlag = append(ResourceTypeAperiodicOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeAperiodicOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.AperiodicResourceType), true, vLb, vUb)
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

func (x *ResourceTypeAperiodic) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeAperiodicOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeAperiodicOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.AperiodicResourceType = new(aper.Enumerated)
	*(x.AperiodicResourceType), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceTypeAperiodicOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceTypeAperiodicExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
