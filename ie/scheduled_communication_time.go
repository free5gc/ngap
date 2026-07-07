package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ScheduledCommunicationTime struct {
	DayofWeek      *aper.BitString                                             // sizeLB:7,sizeUB:7,optional
	TimeofDayStart *int64                                                      // valueExt,valueLB:0,valueUB:86399,optional
	TimeofDayEnd   *int64                                                      // valueExt,valueLB:0,valueUB:86399,optional
	IEExtensions   *ProtocolExtensionContainerScheduledCommunicationTimeExtIEs // optional
}

func (x *ScheduledCommunicationTime) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ScheduledCommunicationTimeOptPresentFlag := []bool{}
	// optional field
	if x.DayofWeek != nil {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, true)
	} else {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, false)
	}
	// optional field
	if x.TimeofDayStart != nil {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, true)
	} else {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, false)
	}
	// optional field
	if x.TimeofDayEnd != nil {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, true)
	} else {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, true)
	} else {
		ScheduledCommunicationTimeOptPresentFlag = append(ScheduledCommunicationTimeOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ScheduledCommunicationTimeOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.DayofWeek != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 7, 7
		err = pd.WriteBitString(*(x.DayofWeek), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
	}

	// optional field
	if x.TimeofDayStart != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 86399
		err = pd.WriteInteger(*(x.TimeofDayStart), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.TimeofDayEnd != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 86399
		err = pd.WriteInteger(*(x.TimeofDayEnd), true, vLb, vUb)
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

func (x *ScheduledCommunicationTime) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ScheduledCommunicationTimeOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ScheduledCommunicationTimeOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if ScheduledCommunicationTimeOptPresentFlag[0] {
		// Read BitString (Pointer)
		*sLb, *sUb = 7, 7
		x.DayofWeek = new(aper.BitString)
		*(x.DayofWeek), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if ScheduledCommunicationTimeOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 86399
		x.TimeofDayStart = new(int64)
		*(x.TimeofDayStart), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if ScheduledCommunicationTimeOptPresentFlag[2] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 86399
		x.TimeofDayEnd = new(int64)
		*(x.TimeofDayEnd), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if ScheduledCommunicationTimeOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerScheduledCommunicationTimeExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
