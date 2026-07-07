package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PosResourceSetTypeAperiodic struct {
	SRSResourceTrigger *int64                                                       // valueLB:1,valueUB:3
	IEExtensions       *ProtocolExtensionContainerPosResourceSetTypeAperiodicExtIEs // optional
}

func (x *PosResourceSetTypeAperiodic) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosResourceSetTypeAperiodicOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceTrigger == nil {
		return errors.Errorf("SRSResourceTrigger is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PosResourceSetTypeAperiodicOptPresentFlag = append(PosResourceSetTypeAperiodicOptPresentFlag, true)
	} else {
		PosResourceSetTypeAperiodicOptPresentFlag = append(PosResourceSetTypeAperiodicOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PosResourceSetTypeAperiodicOptPresentFlag, true)
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

func (x *PosResourceSetTypeAperiodic) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosResourceSetTypeAperiodicOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PosResourceSetTypeAperiodicOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if PosResourceSetTypeAperiodicOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPosResourceSetTypeAperiodicExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
