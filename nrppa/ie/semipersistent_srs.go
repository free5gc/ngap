package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SemipersistentSRS struct {
	SRSResourceSetID *SRSResourceSetID
	IEExtensions     *ProtocolExtensionContainerSemipersistentSRSExtIEs // optional
}

func (x *SemipersistentSRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SemipersistentSRSOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceSetID == nil {
		return errors.Errorf("SRSResourceSetID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SemipersistentSRSOptPresentFlag = append(SemipersistentSRSOptPresentFlag, true)
	} else {
		SemipersistentSRSOptPresentFlag = append(SemipersistentSRSOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SemipersistentSRSOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSResourceSetID marshal failed")
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

func (x *SemipersistentSRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SemipersistentSRSOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SemipersistentSRSOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSResourceSetID = new(SRSResourceSetID)
	err = x.SRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSResourceSetID error")
	}

	// optional field (optPresentFlag index: 0)
	if SemipersistentSRSOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSemipersistentSRSExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
