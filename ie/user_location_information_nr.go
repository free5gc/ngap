package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UserLocationInformationNR struct {
	NRCGI        *NRCGI                                                     // valueExt
	TAI          *TAI                                                       // valueExt
	TimeStamp    *TimeStamp                                                 // optional
	IEExtensions *ProtocolExtensionContainerUserLocationInformationNRExtIEs // optional
}

func (x *UserLocationInformationNR) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UserLocationInformationNROptPresentFlag := []bool{}
	// mandatory field
	if x.NRCGI == nil {
		return errors.Errorf("NRCGI is missing")
	}
	// mandatory field
	if x.TAI == nil {
		return errors.Errorf("TAI is missing")
	}
	// optional field
	if x.TimeStamp != nil {
		UserLocationInformationNROptPresentFlag = append(UserLocationInformationNROptPresentFlag, true)
	} else {
		UserLocationInformationNROptPresentFlag = append(UserLocationInformationNROptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UserLocationInformationNROptPresentFlag = append(UserLocationInformationNROptPresentFlag, true)
	} else {
		UserLocationInformationNROptPresentFlag = append(UserLocationInformationNROptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UserLocationInformationNROptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRCGI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAI marshal failed")
	}

	// optional field
	if x.TimeStamp != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TimeStamp.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TimeStamp marshal failed")
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

func (x *UserLocationInformationNR) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UserLocationInformationNROptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UserLocationInformationNROptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRCGI = new(NRCGI)
	err = x.NRCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRCGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAI = new(TAI)
	err = x.TAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAI error")
	}

	// optional field (optPresentFlag index: 0)
	if UserLocationInformationNROptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TimeStamp = new(TimeStamp)
		err = x.TimeStamp.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TimeStamp error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UserLocationInformationNROptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUserLocationInformationNRExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
