package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResourceSetTypeAperiodic struct {
	SRSResourceTrigger *int64                                                    // valueLB:1,valueUB:3
	Slotoffset         *int64                                                    // valueLB:0,valueUB:32
	IEExtensions       *ProtocolExtensionContainerResourceSetTypeAperiodicExtIEs // optional
}

func (x *ResourceSetTypeAperiodic) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypeAperiodicOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceTrigger == nil {
		return errors.Errorf("SRSResourceTrigger is missing")
	}
	// mandatory field
	if x.Slotoffset == nil {
		return errors.Errorf("Slotoffset is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResourceSetTypeAperiodicOptPresentFlag = append(ResourceSetTypeAperiodicOptPresentFlag, true)
	} else {
		ResourceSetTypeAperiodicOptPresentFlag = append(ResourceSetTypeAperiodicOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypeAperiodicOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 1, 3
	err = pd.WriteInteger(*(x.SRSResourceTrigger), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 32
	err = pd.WriteInteger(*(x.Slotoffset), false, vLb, vUb)
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

func (x *ResourceSetTypeAperiodic) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypeAperiodicOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypeAperiodicOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 1, 3
	x.SRSResourceTrigger = new(int64)
	*(x.SRSResourceTrigger), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 32
	x.Slotoffset = new(int64)
	*(x.Slotoffset), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if ResourceSetTypeAperiodicOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResourceSetTypeAperiodicExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
