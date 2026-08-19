package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceTypeSemipersistentPosPeriodicityPresentSlot1     aper.Enumerated = 0
	ResourceTypeSemipersistentPosPeriodicityPresentSlot2     aper.Enumerated = 1
	ResourceTypeSemipersistentPosPeriodicityPresentSlot4     aper.Enumerated = 2
	ResourceTypeSemipersistentPosPeriodicityPresentSlot5     aper.Enumerated = 3
	ResourceTypeSemipersistentPosPeriodicityPresentSlot8     aper.Enumerated = 4
	ResourceTypeSemipersistentPosPeriodicityPresentSlot10    aper.Enumerated = 5
	ResourceTypeSemipersistentPosPeriodicityPresentSlot16    aper.Enumerated = 6
	ResourceTypeSemipersistentPosPeriodicityPresentSlot20    aper.Enumerated = 7
	ResourceTypeSemipersistentPosPeriodicityPresentSlot32    aper.Enumerated = 8
	ResourceTypeSemipersistentPosPeriodicityPresentSlot40    aper.Enumerated = 9
	ResourceTypeSemipersistentPosPeriodicityPresentSlot64    aper.Enumerated = 10
	ResourceTypeSemipersistentPosPeriodicityPresentSlot80    aper.Enumerated = 11
	ResourceTypeSemipersistentPosPeriodicityPresentSlot160   aper.Enumerated = 12
	ResourceTypeSemipersistentPosPeriodicityPresentSlot320   aper.Enumerated = 13
	ResourceTypeSemipersistentPosPeriodicityPresentSlot640   aper.Enumerated = 14
	ResourceTypeSemipersistentPosPeriodicityPresentSlot1280  aper.Enumerated = 15
	ResourceTypeSemipersistentPosPeriodicityPresentSlot2560  aper.Enumerated = 16
	ResourceTypeSemipersistentPosPeriodicityPresentSlot5120  aper.Enumerated = 17
	ResourceTypeSemipersistentPosPeriodicityPresentSlot10240 aper.Enumerated = 18
	ResourceTypeSemipersistentPosPeriodicityPresentSlot40960 aper.Enumerated = 19
	ResourceTypeSemipersistentPosPeriodicityPresentSlot81920 aper.Enumerated = 20
	ResourceTypeSemipersistentPosPeriodicityPresentSlot128   aper.Enumerated = 21
	ResourceTypeSemipersistentPosPeriodicityPresentSlot256   aper.Enumerated = 22
	ResourceTypeSemipersistentPosPeriodicityPresentSlot512   aper.Enumerated = 23
	ResourceTypeSemipersistentPosPeriodicityPresentSlot20480 aper.Enumerated = 24
)

type ResourceTypeSemiPersistentPos struct {
	Periodicity  *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:20
	Offset       *int64                                                         // valueExt,valueLB:0,valueUB:81919
	IEExtensions *ProtocolExtensionContainerResourceTypeSemiPersistentPosExtIEs // optional
}

func (x *ResourceTypeSemiPersistentPos) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeSemiPersistentPosOptPresentFlag := []bool{}
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
		ResourceTypeSemiPersistentPosOptPresentFlag = append(ResourceTypeSemiPersistentPosOptPresentFlag, true)
	} else {
		ResourceTypeSemiPersistentPosOptPresentFlag = append(ResourceTypeSemiPersistentPosOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeSemiPersistentPosOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 20
	err = pd.WriteEnumerated(*(x.Periodicity), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 81919
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

func (x *ResourceTypeSemiPersistentPos) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeSemiPersistentPosOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeSemiPersistentPosOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 20
	x.Periodicity = new(aper.Enumerated)
	*(x.Periodicity), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 81919
	x.Offset = new(int64)
	*(x.Offset), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceTypeSemiPersistentPosOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceTypeSemiPersistentPosExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
