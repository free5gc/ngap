package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	AperiodicSRSAperiodicPresentTrue aper.Enumerated = 0
)

type AperiodicSRS struct {
	Aperiodic          *aper.Enumerated                              // valueExt,valueLB:0,valueUB:0
	SRSResourceTrigger *SRSResourceTrigger                           // valueExt,optional
	IEExtensions       *ProtocolExtensionContainerAperiodicSRSExtIEs // optional
}

func (x *AperiodicSRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AperiodicSRSOptPresentFlag := []bool{}
	// mandatory field
	if x.Aperiodic == nil {
		return errors.Errorf("Aperiodic is missing")
	}
	// optional field
	if x.SRSResourceTrigger != nil {
		AperiodicSRSOptPresentFlag = append(AperiodicSRSOptPresentFlag, true)
	} else {
		AperiodicSRSOptPresentFlag = append(AperiodicSRSOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AperiodicSRSOptPresentFlag = append(AperiodicSRSOptPresentFlag, true)
	} else {
		AperiodicSRSOptPresentFlag = append(AperiodicSRSOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AperiodicSRSOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.Aperiodic), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// optional field
	if x.SRSResourceTrigger != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSResourceTrigger.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SRSResourceTrigger marshal failed")
		}
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

func (x *AperiodicSRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AperiodicSRSOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&AperiodicSRSOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.Aperiodic = new(aper.Enumerated)
	*(x.Aperiodic), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if AperiodicSRSOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SRSResourceTrigger = new(SRSResourceTrigger)
		err = x.SRSResourceTrigger.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SRSResourceTrigger error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AperiodicSRSOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAperiodicSRSExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
