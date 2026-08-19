package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type M1PeriodicReporting struct {
	ReportInterval *ReportIntervalMDT                                   // valueLB:0,valueUB:12
	ReportAmount   *ReportAmountMDT                                     // valueLB:0,valueUB:7
	IEExtensions   *ProtocolExtensionContainerM1PeriodicReportingExtIEs // optional
}

func (x *M1PeriodicReporting) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	M1PeriodicReportingOptPresentFlag := []bool{}
	// mandatory field
	if x.ReportInterval == nil {
		return errors.Errorf("ReportInterval is missing")
	}
	// mandatory field
	if x.ReportAmount == nil {
		return errors.Errorf("ReportAmount is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		M1PeriodicReportingOptPresentFlag = append(M1PeriodicReportingOptPresentFlag, true)
	} else {
		M1PeriodicReportingOptPresentFlag = append(M1PeriodicReportingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(M1PeriodicReportingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ReportInterval.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportInterval marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ReportAmount.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportAmount marshal failed")
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

func (x *M1PeriodicReporting) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	M1PeriodicReportingOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&M1PeriodicReportingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportInterval = new(ReportIntervalMDT)
	err = x.ReportInterval.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportInterval error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportAmount = new(ReportAmountMDT)
	err = x.ReportAmount.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportAmount error")
	}

	// optional field (optPresentFlag index: 0)
	if M1PeriodicReportingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerM1PeriodicReportingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
