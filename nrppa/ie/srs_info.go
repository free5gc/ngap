package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSInfo struct {
	SRSResource *SRSResourceID
}

func (x *SRSInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResource == nil {
		return errors.Errorf("SRSResource is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SRSResource.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSResource marshal failed")
	}

	return nil
}

func (x *SRSInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSInfoOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSResource = new(SRSResourceID)
	err = x.SRSResource.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSResource error")
	}

	return nil
}
