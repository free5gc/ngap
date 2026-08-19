package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type AoAAssistanceInfo struct {
	AngleMeasurement    *AngleMeasurementType                              // valueLB:0,valueUB:2
	LCSToGCSTranslation *LCSToGCSTranslation                               // valueExt,optional
	IEExtensions        *ProtocolExtensionContainerAoAAssistanceInfoExtIEs // optional
}

func (x *AoAAssistanceInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AoAAssistanceInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.AngleMeasurement == nil {
		return errors.Errorf("AngleMeasurement is missing")
	}
	// optional field
	if x.LCSToGCSTranslation != nil {
		AoAAssistanceInfoOptPresentFlag = append(AoAAssistanceInfoOptPresentFlag, true)
	} else {
		AoAAssistanceInfoOptPresentFlag = append(AoAAssistanceInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AoAAssistanceInfoOptPresentFlag = append(AoAAssistanceInfoOptPresentFlag, true)
	} else {
		AoAAssistanceInfoOptPresentFlag = append(AoAAssistanceInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AoAAssistanceInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AngleMeasurement.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AngleMeasurement marshal failed")
	}

	// optional field
	if x.LCSToGCSTranslation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LCSToGCSTranslation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "LCSToGCSTranslation marshal failed")
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

func (x *AoAAssistanceInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AoAAssistanceInfoOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&AoAAssistanceInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AngleMeasurement = new(AngleMeasurementType)
	err = x.AngleMeasurement.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AngleMeasurement error")
	}

	// optional field (optPresentFlag index: 0)
	if AoAAssistanceInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.LCSToGCSTranslation = new(LCSToGCSTranslation)
		err = x.LCSToGCSTranslation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode LCSToGCSTranslation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AoAAssistanceInfoOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAoAAssistanceInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
