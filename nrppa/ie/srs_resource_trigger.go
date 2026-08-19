package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSResourceTrigger struct {
	AperiodicSRSResourceTriggerList *AperiodicSRSResourceTriggerList
	IEExtensions                    *ProtocolExtensionContainerSRSResourceTriggerExtIEs // optional
}

func (x *SRSResourceTrigger) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceTriggerOptPresentFlag := []bool{}
	// mandatory field
	if x.AperiodicSRSResourceTriggerList == nil {
		return errors.Errorf("AperiodicSRSResourceTriggerList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SRSResourceTriggerOptPresentFlag = append(SRSResourceTriggerOptPresentFlag, true)
	} else {
		SRSResourceTriggerOptPresentFlag = append(SRSResourceTriggerOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceTriggerOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AperiodicSRSResourceTriggerList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AperiodicSRSResourceTriggerList marshal failed")
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

func (x *SRSResourceTrigger) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceTriggerOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceTriggerOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AperiodicSRSResourceTriggerList = new(AperiodicSRSResourceTriggerList)
	err = x.AperiodicSRSResourceTriggerList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AperiodicSRSResourceTriggerList error")
	}

	// optional field (optPresentFlag index: 0)
	if SRSResourceTriggerOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSResourceTriggerExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
