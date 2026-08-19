package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TimeStamp struct {
	SystemFrameNumber *SystemFrameNumber
	SlotIndex         *TimeStampSlotIndex                        // valueLB:0,valueUB:4
	MeasurementTime   *RelativeTime1900                          // optional
	IEExtension       *ProtocolExtensionContainerTimeStampExtIEs // optional
}

func (x *TimeStamp) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TimeStampOptPresentFlag := []bool{}
	// mandatory field
	if x.SystemFrameNumber == nil {
		return errors.Errorf("SystemFrameNumber is missing")
	}
	// mandatory field
	if x.SlotIndex == nil {
		return errors.Errorf("SlotIndex is missing")
	}
	// optional field
	if x.MeasurementTime != nil {
		TimeStampOptPresentFlag = append(TimeStampOptPresentFlag, true)
	} else {
		TimeStampOptPresentFlag = append(TimeStampOptPresentFlag, false)
	}
	// optional field
	if x.IEExtension != nil {
		TimeStampOptPresentFlag = append(TimeStampOptPresentFlag, true)
	} else {
		TimeStampOptPresentFlag = append(TimeStampOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TimeStampOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SystemFrameNumber.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SystemFrameNumber marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SlotIndex.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SlotIndex marshal failed")
	}

	// optional field
	if x.MeasurementTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementTime.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MeasurementTime marshal failed")
		}
	}

	// optional field
	if x.IEExtension != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtension.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtension marshal failed")
		}
	}

	return nil
}

func (x *TimeStamp) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TimeStampOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TimeStampOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SystemFrameNumber = new(SystemFrameNumber)
	err = x.SystemFrameNumber.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SystemFrameNumber error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SlotIndex = new(TimeStampSlotIndex)
	err = x.SlotIndex.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SlotIndex error")
	}

	// optional field (optPresentFlag index: 0)
	if TimeStampOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementTime = new(RelativeTime1900)
		err = x.MeasurementTime.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementTime error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TimeStampOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerTimeStampExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
