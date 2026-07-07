package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type StartTimeAndDuration struct {
	StartTime    *RelativeTime1900                                     // optional
	Duration     *int64                                                // valueExt,valueLB:0,valueUB:90060,optional
	IEExtensions *ProtocolExtensionContainerStartTimeAndDurationExtIEs // optional
}

func (x *StartTimeAndDuration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	StartTimeAndDurationOptPresentFlag := []bool{}
	// optional field
	if x.StartTime != nil {
		StartTimeAndDurationOptPresentFlag = append(StartTimeAndDurationOptPresentFlag, true)
	} else {
		StartTimeAndDurationOptPresentFlag = append(StartTimeAndDurationOptPresentFlag, false)
	}
	// optional field
	if x.Duration != nil {
		StartTimeAndDurationOptPresentFlag = append(StartTimeAndDurationOptPresentFlag, true)
	} else {
		StartTimeAndDurationOptPresentFlag = append(StartTimeAndDurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		StartTimeAndDurationOptPresentFlag = append(StartTimeAndDurationOptPresentFlag, true)
	} else {
		StartTimeAndDurationOptPresentFlag = append(StartTimeAndDurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(StartTimeAndDurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.StartTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.StartTime.Write(pd)
		if err != nil {
			return errors.Wrap(err, "StartTime marshal failed")
		}
	}

	// optional field
	if x.Duration != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 90060
		err = pd.WriteInteger(*(x.Duration), true, vLb, vUb)
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

func (x *StartTimeAndDuration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	StartTimeAndDurationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&StartTimeAndDurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if StartTimeAndDurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.StartTime = new(RelativeTime1900)
		err = x.StartTime.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode StartTime error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if StartTimeAndDurationOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 90060
		x.Duration = new(int64)
		*(x.Duration), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if StartTimeAndDurationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerStartTimeAndDurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
