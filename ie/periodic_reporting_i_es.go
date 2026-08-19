package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PeriodicReportingIEs struct {
	ReportingPeriodicity *ReportingPeriodicity                                 // valueExt,valueLB:0,valueUB:5
	IEExtensions         *ProtocolExtensionContainerPeriodicReportingIEsExtIEs // optional
}

func (x *PeriodicReportingIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PeriodicReportingIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.ReportingPeriodicity == nil {
		return errors.Errorf("ReportingPeriodicity is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PeriodicReportingIEsOptPresentFlag = append(PeriodicReportingIEsOptPresentFlag, true)
	} else {
		PeriodicReportingIEsOptPresentFlag = append(PeriodicReportingIEsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PeriodicReportingIEsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ReportingPeriodicity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportingPeriodicity marshal failed")
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

func (x *PeriodicReportingIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PeriodicReportingIEsOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PeriodicReportingIEsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportingPeriodicity = new(ReportingPeriodicity)
	err = x.ReportingPeriodicity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportingPeriodicity error")
	}

	// optional field (optPresentFlag index: 0)
	if PeriodicReportingIEsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPeriodicReportingIEsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
