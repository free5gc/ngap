package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceTypeSemipersistentPeriodicityPresentSlot1    aper.Enumerated = 0
	ResourceTypeSemipersistentPeriodicityPresentSlot2    aper.Enumerated = 1
	ResourceTypeSemipersistentPeriodicityPresentSlot4    aper.Enumerated = 2
	ResourceTypeSemipersistentPeriodicityPresentSlot5    aper.Enumerated = 3
	ResourceTypeSemipersistentPeriodicityPresentSlot8    aper.Enumerated = 4
	ResourceTypeSemipersistentPeriodicityPresentSlot10   aper.Enumerated = 5
	ResourceTypeSemipersistentPeriodicityPresentSlot16   aper.Enumerated = 6
	ResourceTypeSemipersistentPeriodicityPresentSlot20   aper.Enumerated = 7
	ResourceTypeSemipersistentPeriodicityPresentSlot32   aper.Enumerated = 8
	ResourceTypeSemipersistentPeriodicityPresentSlot40   aper.Enumerated = 9
	ResourceTypeSemipersistentPeriodicityPresentSlot64   aper.Enumerated = 10
	ResourceTypeSemipersistentPeriodicityPresentSlot80   aper.Enumerated = 11
	ResourceTypeSemipersistentPeriodicityPresentSlot160  aper.Enumerated = 12
	ResourceTypeSemipersistentPeriodicityPresentSlot320  aper.Enumerated = 13
	ResourceTypeSemipersistentPeriodicityPresentSlot640  aper.Enumerated = 14
	ResourceTypeSemipersistentPeriodicityPresentSlot1280 aper.Enumerated = 15
	ResourceTypeSemipersistentPeriodicityPresentSlot2560 aper.Enumerated = 16
)

type ResourceTypeSemiPersistent struct {
	Periodicity  *aper.Enumerated                                            // valueExt,valueLB:0,valueUB:16
	Offset       *int64                                                      // valueExt,valueLB:0,valueUB:2559
	IEExtensions *ProtocolExtensionContainerResourceTypeSemiPersistentExtIEs // optional
}

func (x *ResourceTypeSemiPersistent) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeSemiPersistentOptPresentFlag := []bool{}
	// mandatory field
	if x.Periodicity == nil {
		return errors.Errorf("Periodicity is missing")
	}
	// mandatory field
	if x.Offset == nil {
		return errors.Errorf("Offset is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResourceTypeSemiPersistentOptPresentFlag = append(ResourceTypeSemiPersistentOptPresentFlag, true)
	} else {
		ResourceTypeSemiPersistentOptPresentFlag = append(ResourceTypeSemiPersistentOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeSemiPersistentOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 16
	err = pd.WriteEnumerated(*(x.Periodicity), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 2559
	err = pd.WriteInteger(*(x.Offset), true, vLb, vUb)
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

func (x *ResourceTypeSemiPersistent) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeSemiPersistentOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeSemiPersistentOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 16
	x.Periodicity = new(aper.Enumerated)
	*(x.Periodicity), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 2559
	x.Offset = new(int64)
	*(x.Offset), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceTypeSemiPersistentOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceTypeSemiPersistentExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
