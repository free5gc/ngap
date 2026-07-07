package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ExpectedUEMovingTrajectoryItem struct {
	NGRANCGI         *NGRANCGI                                                       // valueLB:0,valueUB:2
	TimeStayedInCell *int64                                                          // valueLB:0,valueUB:4095,optional
	IEExtensions     *ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs // optional
}

func (x *ExpectedUEMovingTrajectoryItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedUEMovingTrajectoryItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCGI == nil {
		return errors.Errorf("NGRANCGI is missing")
	}
	// optional field
	if x.TimeStayedInCell != nil {
		ExpectedUEMovingTrajectoryItemOptPresentFlag = append(ExpectedUEMovingTrajectoryItemOptPresentFlag, true)
	} else {
		ExpectedUEMovingTrajectoryItemOptPresentFlag = append(ExpectedUEMovingTrajectoryItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExpectedUEMovingTrajectoryItemOptPresentFlag = append(ExpectedUEMovingTrajectoryItemOptPresentFlag, true)
	} else {
		ExpectedUEMovingTrajectoryItemOptPresentFlag = append(ExpectedUEMovingTrajectoryItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedUEMovingTrajectoryItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCGI marshal failed")
	}

	// optional field
	if x.TimeStayedInCell != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 4095
		err = pd.WriteInteger(*(x.TimeStayedInCell), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *ExpectedUEMovingTrajectoryItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedUEMovingTrajectoryItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ExpectedUEMovingTrajectoryItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCGI = new(NGRANCGI)
	err = x.NGRANCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCGI error")
	}

	// optional field (optPresentFlag index: 0)
	if ExpectedUEMovingTrajectoryItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 4095
		x.TimeStayedInCell = new(int64)
		*(x.TimeStayedInCell), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExpectedUEMovingTrajectoryItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExpectedUEMovingTrajectoryItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
