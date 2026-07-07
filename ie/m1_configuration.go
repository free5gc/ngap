package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type M1Configuration struct {
	M1reportingTrigger  *M1ReportingTrigger                              // valueExt,valueLB:0,valueUB:2
	M1thresholdEventA2  *M1ThresholdEventA2                              // valueExt,optional
	M1periodicReporting *M1PeriodicReporting                             // valueExt,optional
	IEExtensions        *ProtocolExtensionContainerM1ConfigurationExtIEs // optional
}

func (x *M1Configuration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	M1ConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.M1reportingTrigger == nil {
		return errors.Errorf("M1reportingTrigger is missing")
	}
	// optional field
	if x.M1thresholdEventA2 != nil {
		M1ConfigurationOptPresentFlag = append(M1ConfigurationOptPresentFlag, true)
	} else {
		M1ConfigurationOptPresentFlag = append(M1ConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.M1periodicReporting != nil {
		M1ConfigurationOptPresentFlag = append(M1ConfigurationOptPresentFlag, true)
	} else {
		M1ConfigurationOptPresentFlag = append(M1ConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		M1ConfigurationOptPresentFlag = append(M1ConfigurationOptPresentFlag, true)
	} else {
		M1ConfigurationOptPresentFlag = append(M1ConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(M1ConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.M1reportingTrigger.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M1reportingTrigger marshal failed")
	}

	// optional field
	if x.M1thresholdEventA2 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M1thresholdEventA2.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M1thresholdEventA2 marshal failed")
		}
	}

	// optional field
	if x.M1periodicReporting != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.M1periodicReporting.Write(pd)
		if err != nil {
			return errors.Wrap(err, "M1periodicReporting marshal failed")
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

func (x *M1Configuration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	M1ConfigurationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&M1ConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M1reportingTrigger = new(M1ReportingTrigger)
	err = x.M1reportingTrigger.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M1reportingTrigger error")
	}

	// optional field (optPresentFlag index: 0)
	if M1ConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.M1thresholdEventA2 = new(M1ThresholdEventA2)
		err = x.M1thresholdEventA2.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M1thresholdEventA2 error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if M1ConfigurationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.M1periodicReporting = new(M1PeriodicReporting)
		err = x.M1periodicReporting.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode M1periodicReporting error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if M1ConfigurationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerM1ConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
