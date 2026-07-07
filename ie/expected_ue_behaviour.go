package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ExpectedUEBehaviour struct {
	ExpectedUEActivityBehaviour *ExpectedUEActivityBehaviour                         // valueExt,optional
	ExpectedHOInterval          *ExpectedHOInterval                                  // valueExt,valueLB:0,valueUB:6,optional
	ExpectedUEMobility          *ExpectedUEMobility                                  // valueExt,valueLB:0,valueUB:1,optional
	ExpectedUEMovingTrajectory  *ExpectedUEMovingTrajectory                          // optional
	IEExtensions                *ProtocolExtensionContainerExpectedUEBehaviourExtIEs // optional
}

func (x *ExpectedUEBehaviour) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedUEBehaviourOptPresentFlag := []bool{}
	// optional field
	if x.ExpectedUEActivityBehaviour != nil {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.ExpectedHOInterval != nil {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.ExpectedUEMobility != nil {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.ExpectedUEMovingTrajectory != nil {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, true)
	} else {
		ExpectedUEBehaviourOptPresentFlag = append(ExpectedUEBehaviourOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedUEBehaviourOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.ExpectedUEActivityBehaviour != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedUEActivityBehaviour.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedUEActivityBehaviour marshal failed")
		}
	}

	// optional field
	if x.ExpectedHOInterval != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedHOInterval.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedHOInterval marshal failed")
		}
	}

	// optional field
	if x.ExpectedUEMobility != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedUEMobility.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedUEMobility marshal failed")
		}
	}

	// optional field
	if x.ExpectedUEMovingTrajectory != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedUEMovingTrajectory.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedUEMovingTrajectory marshal failed")
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

func (x *ExpectedUEBehaviour) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedUEBehaviourOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&ExpectedUEBehaviourOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if ExpectedUEBehaviourOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedUEActivityBehaviour = new(ExpectedUEActivityBehaviour)
		err = x.ExpectedUEActivityBehaviour.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedUEActivityBehaviour error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExpectedUEBehaviourOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedHOInterval = new(ExpectedHOInterval)
		err = x.ExpectedHOInterval.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedHOInterval error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ExpectedUEBehaviourOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedUEMobility = new(ExpectedUEMobility)
		err = x.ExpectedUEMobility.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedUEMobility error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ExpectedUEBehaviourOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedUEMovingTrajectory = new(ExpectedUEMovingTrajectory)
		err = x.ExpectedUEMovingTrajectory.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedUEMovingTrajectory error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if ExpectedUEBehaviourOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedUEBehaviourExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
