package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type IntersystemResourceStatusRequest struct {
	ReportingSystem       *ReportingSystem // valueLB:0,valueUB:3
	ReportCharacteristics *ReportCharacteristics
	ReportType            *ReportType                                                       // valueLB:0,valueUB:2
	IEExtensions          *ProtocolExtensionContainerIntersystemResourceStatusRequestExtIEs // optional
}

func (x *IntersystemResourceStatusRequest) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemResourceStatusRequestOptPresentFlag := []bool{}
	// mandatory field
	if x.ReportingSystem == nil {
		return errors.Errorf("ReportingSystem is missing")
	}
	// mandatory field
	if x.ReportCharacteristics == nil {
		return errors.Errorf("ReportCharacteristics is missing")
	}
	// mandatory field
	if x.ReportType == nil {
		return errors.Errorf("ReportType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemResourceStatusRequestOptPresentFlag = append(IntersystemResourceStatusRequestOptPresentFlag, true)
	} else {
		IntersystemResourceStatusRequestOptPresentFlag = append(IntersystemResourceStatusRequestOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemResourceStatusRequestOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ReportingSystem.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportingSystem marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ReportCharacteristics.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportCharacteristics marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ReportType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportType marshal failed")
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

func (x *IntersystemResourceStatusRequest) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemResourceStatusRequestOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemResourceStatusRequestOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportingSystem = new(ReportingSystem)
	err = x.ReportingSystem.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportingSystem error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportCharacteristics = new(ReportCharacteristics)
	err = x.ReportCharacteristics.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportCharacteristics error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportType = new(ReportType)
	err = x.ReportType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportType error")
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemResourceStatusRequestOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemResourceStatusRequestExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
