package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceSetTypeSemipersistentSemiPersistentSetPresentTrue aper.Enumerated = 0
)

type ResourceSetTypeSemiPersistent struct {
	SemiPersistentSet *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:0
	IEExtensions      *ProtocolExtensionContainerResourceSetTypeSemiPersistentExtIEs // optional
}

func (x *ResourceSetTypeSemiPersistent) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypeSemiPersistentOptPresentFlag := []bool{}
	// mandatory field
	if x.SemiPersistentSet == nil {
		return errors.Errorf("SemiPersistentSet is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResourceSetTypeSemiPersistentOptPresentFlag = append(ResourceSetTypeSemiPersistentOptPresentFlag, true)
	} else {
		ResourceSetTypeSemiPersistentOptPresentFlag = append(ResourceSetTypeSemiPersistentOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypeSemiPersistentOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.SemiPersistentSet), true, vLb, vUb)
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

func (x *ResourceSetTypeSemiPersistent) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypeSemiPersistentOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypeSemiPersistentOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.SemiPersistentSet = new(aper.Enumerated)
	*(x.SemiPersistentSet), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceSetTypeSemiPersistentOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceSetTypeSemiPersistentExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
