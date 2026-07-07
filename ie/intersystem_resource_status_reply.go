package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type IntersystemResourceStatusReply struct {
	Reportingsystem *ReportingSystem                                                // valueLB:0,valueUB:3
	IEExtensions    *ProtocolExtensionContainerIntersystemResourceStatusReplyExtIEs // optional
}

func (x *IntersystemResourceStatusReply) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemResourceStatusReplyOptPresentFlag := []bool{}
	// mandatory field
	if x.Reportingsystem == nil {
		return errors.Errorf("Reportingsystem is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemResourceStatusReplyOptPresentFlag = append(IntersystemResourceStatusReplyOptPresentFlag, true)
	} else {
		IntersystemResourceStatusReplyOptPresentFlag = append(IntersystemResourceStatusReplyOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemResourceStatusReplyOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Reportingsystem.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Reportingsystem marshal failed")
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

func (x *IntersystemResourceStatusReply) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemResourceStatusReplyOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemResourceStatusReplyOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Reportingsystem = new(ReportingSystem)
	err = x.Reportingsystem.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Reportingsystem error")
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemResourceStatusReplyOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemResourceStatusReplyExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
