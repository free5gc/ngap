package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EventBasedReportingIEs struct {
	IntersystemResourceThresholdLow    *IntersystemResourceThreshold
	IntersystemResourceThresholdHigh   *IntersystemResourceThreshold
	NumberOfMeasurementReportingLevels *NumberOfMeasurementReportingLevels                     // valueExt,valueLB:0,valueUB:4
	IEExtensions                       *ProtocolExtensionContainerEventBasedReportingIEsExtIEs // optional
}

func (x *EventBasedReportingIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EventBasedReportingIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.IntersystemResourceThresholdLow == nil {
		return errors.Errorf("IntersystemResourceThresholdLow is missing")
	}
	// mandatory field
	if x.IntersystemResourceThresholdHigh == nil {
		return errors.Errorf("IntersystemResourceThresholdHigh is missing")
	}
	// mandatory field
	if x.NumberOfMeasurementReportingLevels == nil {
		return errors.Errorf("NumberOfMeasurementReportingLevels is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EventBasedReportingIEsOptPresentFlag = append(EventBasedReportingIEsOptPresentFlag, true)
	} else {
		EventBasedReportingIEsOptPresentFlag = append(EventBasedReportingIEsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EventBasedReportingIEsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.IntersystemResourceThresholdLow.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IntersystemResourceThresholdLow marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.IntersystemResourceThresholdHigh.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IntersystemResourceThresholdHigh marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NumberOfMeasurementReportingLevels.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NumberOfMeasurementReportingLevels marshal failed")
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

func (x *EventBasedReportingIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EventBasedReportingIEsOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EventBasedReportingIEsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IntersystemResourceThresholdLow = new(IntersystemResourceThreshold)
	err = x.IntersystemResourceThresholdLow.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IntersystemResourceThresholdLow error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IntersystemResourceThresholdHigh = new(IntersystemResourceThreshold)
	err = x.IntersystemResourceThresholdHigh.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IntersystemResourceThresholdHigh error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NumberOfMeasurementReportingLevels = new(NumberOfMeasurementReportingLevels)
	err = x.NumberOfMeasurementReportingLevels.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NumberOfMeasurementReportingLevels error")
	}

	// optional field (optPresentFlag index: 0)
	if EventBasedReportingIEsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEventBasedReportingIEsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
