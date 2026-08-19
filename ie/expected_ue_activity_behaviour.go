package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ExpectedUEActivityBehaviour struct {
	ExpectedActivityPeriod                 *ExpectedActivityPeriod                                      // optional
	ExpectedIdlePeriod                     *ExpectedIdlePeriod                                          // optional
	SourceOfUEActivityBehaviourInformation *SourceOfUEActivityBehaviourInformation                      // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions                           *ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs // optional
}

func (x *ExpectedUEActivityBehaviour) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedUEActivityBehaviourOptPresentFlag := []bool{}
	// optional field
	if x.ExpectedActivityPeriod != nil {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.ExpectedIdlePeriod != nil {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.SourceOfUEActivityBehaviourInformation != nil {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEActivityBehaviourOptPresentFlag = append(ExpectedUEActivityBehaviourOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedUEActivityBehaviourOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.ExpectedActivityPeriod != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedActivityPeriod.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedActivityPeriod marshal failed")
		}
	}

	// optional field
	if x.ExpectedIdlePeriod != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedIdlePeriod.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedIdlePeriod marshal failed")
		}
	}

	// optional field
	if x.SourceOfUEActivityBehaviourInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SourceOfUEActivityBehaviourInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SourceOfUEActivityBehaviourInformation marshal failed")
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

func (x *ExpectedUEActivityBehaviour) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedUEActivityBehaviourOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ExpectedUEActivityBehaviourOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if ExpectedUEActivityBehaviourOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedActivityPeriod = new(ExpectedActivityPeriod)
		err = x.ExpectedActivityPeriod.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedActivityPeriod error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExpectedUEActivityBehaviourOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedIdlePeriod = new(ExpectedIdlePeriod)
		err = x.ExpectedIdlePeriod.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedIdlePeriod error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ExpectedUEActivityBehaviourOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SourceOfUEActivityBehaviourInformation = new(SourceOfUEActivityBehaviourInformation)
		err = x.SourceOfUEActivityBehaviourInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SourceOfUEActivityBehaviourInformation error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ExpectedUEActivityBehaviourOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedUEActivityBehaviourExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
