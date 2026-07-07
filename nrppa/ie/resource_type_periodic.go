package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ResourceTypePeriodicPeriodicityPresentSlot1    aper.Enumerated = 0
	ResourceTypePeriodicPeriodicityPresentSlot2    aper.Enumerated = 1
	ResourceTypePeriodicPeriodicityPresentSlot4    aper.Enumerated = 2
	ResourceTypePeriodicPeriodicityPresentSlot5    aper.Enumerated = 3
	ResourceTypePeriodicPeriodicityPresentSlot8    aper.Enumerated = 4
	ResourceTypePeriodicPeriodicityPresentSlot10   aper.Enumerated = 5
	ResourceTypePeriodicPeriodicityPresentSlot16   aper.Enumerated = 6
	ResourceTypePeriodicPeriodicityPresentSlot20   aper.Enumerated = 7
	ResourceTypePeriodicPeriodicityPresentSlot32   aper.Enumerated = 8
	ResourceTypePeriodicPeriodicityPresentSlot40   aper.Enumerated = 9
	ResourceTypePeriodicPeriodicityPresentSlot64   aper.Enumerated = 10
	ResourceTypePeriodicPeriodicityPresentSlot80   aper.Enumerated = 11
	ResourceTypePeriodicPeriodicityPresentSlot160  aper.Enumerated = 12
	ResourceTypePeriodicPeriodicityPresentSlot320  aper.Enumerated = 13
	ResourceTypePeriodicPeriodicityPresentSlot640  aper.Enumerated = 14
	ResourceTypePeriodicPeriodicityPresentSlot1280 aper.Enumerated = 15
	ResourceTypePeriodicPeriodicityPresentSlot2560 aper.Enumerated = 16
)

type ResourceTypePeriodic struct {
	Periodicity  *aper.Enumerated                                      // valueExt,valueLB:0,valueUB:16
	Offset       *int64                                                // valueExt,valueLB:0,valueUB:2559
	IEExtensions *ProtocolExtensionContainerResourceTypePeriodicExtIEs // optional
}

func (x *ResourceTypePeriodic) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypePeriodicOptPresentFlag := []bool{}
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
		ResourceTypePeriodicOptPresentFlag = append(ResourceTypePeriodicOptPresentFlag, true)
	} else {
		ResourceTypePeriodicOptPresentFlag = append(ResourceTypePeriodicOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypePeriodicOptPresentFlag, true)
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

func (x *ResourceTypePeriodic) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypePeriodicOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypePeriodicOptPresentFlag, true)
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
	if ResourceTypePeriodicOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceTypePeriodicExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
