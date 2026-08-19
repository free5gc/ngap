package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type InterSystemHOReport struct {
	HandoverReportType *InterSystemHandoverReportType                       // valueLB:0,valueUB:2
	IEExtensions       *ProtocolExtensionContainerInterSystemHOReportExtIEs // optional
}

func (x *InterSystemHOReport) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	InterSystemHOReportOptPresentFlag := []bool{}
	// mandatory field
	if x.HandoverReportType == nil {
		return errors.Errorf("HandoverReportType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		InterSystemHOReportOptPresentFlag = append(InterSystemHOReportOptPresentFlag, true)
	} else {
		InterSystemHOReportOptPresentFlag = append(InterSystemHOReportOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(InterSystemHOReportOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.HandoverReportType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "HandoverReportType marshal failed")
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

func (x *InterSystemHOReport) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	InterSystemHOReportOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&InterSystemHOReportOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.HandoverReportType = new(InterSystemHandoverReportType)
	err = x.HandoverReportType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode HandoverReportType error")
	}

	// optional field (optPresentFlag index: 0)
	if InterSystemHOReportOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerInterSystemHOReportExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
