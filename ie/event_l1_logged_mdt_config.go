package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EventL1LoggedMDTConfig struct {
	L1Threshold   *MeasurementThresholdL1LoggedMDT // valueLB:0,valueUB:2
	Hysteresis    *Hysteresis
	TimeToTrigger *TimeToTrigger                                          // valueLB:0,valueUB:15
	IEExtensions  *ProtocolExtensionContainerEventL1LoggedMDTConfigExtIEs // optional
}

func (x *EventL1LoggedMDTConfig) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EventL1LoggedMDTConfigOptPresentFlag := []bool{}
	// mandatory field
	if x.L1Threshold == nil {
		return errors.Errorf("L1Threshold is missing")
	}
	// mandatory field
	if x.Hysteresis == nil {
		return errors.Errorf("Hysteresis is missing")
	}
	// mandatory field
	if x.TimeToTrigger == nil {
		return errors.Errorf("TimeToTrigger is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EventL1LoggedMDTConfigOptPresentFlag = append(EventL1LoggedMDTConfigOptPresentFlag, true)
	} else {
		EventL1LoggedMDTConfigOptPresentFlag = append(EventL1LoggedMDTConfigOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EventL1LoggedMDTConfigOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.L1Threshold.Write(pd)
	if err != nil {
		return errors.Wrap(err, "L1Threshold marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Hysteresis.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Hysteresis marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TimeToTrigger.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TimeToTrigger marshal failed")
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

func (x *EventL1LoggedMDTConfig) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EventL1LoggedMDTConfigOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EventL1LoggedMDTConfigOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.L1Threshold = new(MeasurementThresholdL1LoggedMDT)
	err = x.L1Threshold.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode L1Threshold error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Hysteresis = new(Hysteresis)
	err = x.Hysteresis.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Hysteresis error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TimeToTrigger = new(TimeToTrigger)
	err = x.TimeToTrigger.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TimeToTrigger error")
	}

	// optional field (optPresentFlag index: 0)
	if EventL1LoggedMDTConfigOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEventL1LoggedMDTConfigExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
