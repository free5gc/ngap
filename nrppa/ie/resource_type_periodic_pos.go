package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceTypePeriodicPosPeriodicityPresentSlot1     aper.Enumerated = 0
	ResourceTypePeriodicPosPeriodicityPresentSlot2     aper.Enumerated = 1
	ResourceTypePeriodicPosPeriodicityPresentSlot4     aper.Enumerated = 2
	ResourceTypePeriodicPosPeriodicityPresentSlot5     aper.Enumerated = 3
	ResourceTypePeriodicPosPeriodicityPresentSlot8     aper.Enumerated = 4
	ResourceTypePeriodicPosPeriodicityPresentSlot10    aper.Enumerated = 5
	ResourceTypePeriodicPosPeriodicityPresentSlot16    aper.Enumerated = 6
	ResourceTypePeriodicPosPeriodicityPresentSlot20    aper.Enumerated = 7
	ResourceTypePeriodicPosPeriodicityPresentSlot32    aper.Enumerated = 8
	ResourceTypePeriodicPosPeriodicityPresentSlot40    aper.Enumerated = 9
	ResourceTypePeriodicPosPeriodicityPresentSlot64    aper.Enumerated = 10
	ResourceTypePeriodicPosPeriodicityPresentSlot80    aper.Enumerated = 11
	ResourceTypePeriodicPosPeriodicityPresentSlot160   aper.Enumerated = 12
	ResourceTypePeriodicPosPeriodicityPresentSlot320   aper.Enumerated = 13
	ResourceTypePeriodicPosPeriodicityPresentSlot640   aper.Enumerated = 14
	ResourceTypePeriodicPosPeriodicityPresentSlot1280  aper.Enumerated = 15
	ResourceTypePeriodicPosPeriodicityPresentSlot2560  aper.Enumerated = 16
	ResourceTypePeriodicPosPeriodicityPresentSlot5120  aper.Enumerated = 17
	ResourceTypePeriodicPosPeriodicityPresentSlot10240 aper.Enumerated = 18
	ResourceTypePeriodicPosPeriodicityPresentSlot40960 aper.Enumerated = 19
	ResourceTypePeriodicPosPeriodicityPresentSlot81920 aper.Enumerated = 20
	ResourceTypePeriodicPosPeriodicityPresentSlot128   aper.Enumerated = 21
	ResourceTypePeriodicPosPeriodicityPresentSlot256   aper.Enumerated = 22
	ResourceTypePeriodicPosPeriodicityPresentSlot512   aper.Enumerated = 23
	ResourceTypePeriodicPosPeriodicityPresentSlot20480 aper.Enumerated = 24
)

type ResourceTypePeriodicPos struct {
	Periodicity  *aper.Enumerated                                         // valueExt,valueLB:0,valueUB:20
	Offset       *int64                                                   // valueExt,valueLB:0,valueUB:81919
	IEExtensions *ProtocolExtensionContainerResourceTypePeriodicPosExtIEs // optional
}

func (x *ResourceTypePeriodicPos) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypePeriodicPosOptPresentFlag := []bool{}
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
		ResourceTypePeriodicPosOptPresentFlag = append(ResourceTypePeriodicPosOptPresentFlag, true)
	} else {
		ResourceTypePeriodicPosOptPresentFlag = append(ResourceTypePeriodicPosOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypePeriodicPosOptPresentFlag, true)
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

func (x *ResourceTypePeriodicPos) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypePeriodicPosOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypePeriodicPosOptPresentFlag, true)
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
	if ResourceTypePeriodicPosOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceTypePeriodicPosExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
