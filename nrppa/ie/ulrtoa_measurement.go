package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ULRTOAMeasurement struct {
	ULRTOAmeas         *ULRTOAMeas                                        // valueLB:0,valueUB:6
	AdditionalPathList *AdditionalPathList                                // optional
	IEExtensions       *ProtocolExtensionContainerULRTOAMeasurementExtIEs // optional
}

func (x *ULRTOAMeasurement) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULRTOAMeasurementOptPresentFlag := []bool{}
	// mandatory field
	if x.ULRTOAmeas == nil {
		return errors.Errorf("ULRTOAmeas is missing")
	}
	// optional field
	if x.AdditionalPathList != nil {
		ULRTOAMeasurementOptPresentFlag = append(ULRTOAMeasurementOptPresentFlag, true)
	} else {
		ULRTOAMeasurementOptPresentFlag = append(ULRTOAMeasurementOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ULRTOAMeasurementOptPresentFlag = append(ULRTOAMeasurementOptPresentFlag, true)
	} else {
		ULRTOAMeasurementOptPresentFlag = append(ULRTOAMeasurementOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ULRTOAMeasurementOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ULRTOAmeas.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ULRTOAmeas marshal failed")
	}

	// optional field
	if x.AdditionalPathList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AdditionalPathList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AdditionalPathList marshal failed")
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

func (x *ULRTOAMeasurement) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULRTOAMeasurementOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ULRTOAMeasurementOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ULRTOAmeas = new(ULRTOAMeas)
	err = x.ULRTOAmeas.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ULRTOAmeas error")
	}

	// optional field (optPresentFlag index: 0)
	if ULRTOAMeasurementOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalPathList = new(AdditionalPathList)
		err = x.AdditionalPathList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalPathList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ULRTOAMeasurementOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerULRTOAMeasurementExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
