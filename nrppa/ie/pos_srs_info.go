package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PosSRSInfo struct {
	PosSRSResourceID *SRSPosResourceID
}

func (x *PosSRSInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSRSInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.PosSRSResourceID == nil {
		return errors.Errorf("PosSRSResourceID is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosSRSInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PosSRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PosSRSResourceID marshal failed")
	}

	return nil
}

func (x *PosSRSInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSRSInfoOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosSRSInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PosSRSResourceID = new(SRSPosResourceID)
	err = x.PosSRSResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PosSRSResourceID error")
	}

	return nil
}
