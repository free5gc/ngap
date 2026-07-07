package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type FailureIndication struct {
	UERLFReportContainer *UERLFReportContainer                              // valueLB:0,valueUB:2
	IEExtensions         *ProtocolExtensionContainerFailureIndicationExtIEs // optional
}

func (x *FailureIndication) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FailureIndicationOptPresentFlag := []bool{}
	// mandatory field
	if x.UERLFReportContainer == nil {
		return errors.Errorf("UERLFReportContainer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		FailureIndicationOptPresentFlag = append(FailureIndicationOptPresentFlag, true)
	} else {
		FailureIndicationOptPresentFlag = append(FailureIndicationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FailureIndicationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.UERLFReportContainer.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UERLFReportContainer marshal failed")
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

func (x *FailureIndication) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FailureIndicationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&FailureIndicationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UERLFReportContainer = new(UERLFReportContainer)
	err = x.UERLFReportContainer.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UERLFReportContainer error")
	}

	// optional field (optPresentFlag index: 0)
	if FailureIndicationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFailureIndicationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
