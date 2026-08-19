package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSResourcetype struct {
	SRSResourceTypeChoice *SRSResourceTypeChoice                           // valueExt,valueLB:0,valueUB:1
	IEExtensions          *ProtocolExtensionContainerSRSResourcetypeExtIEs // optional
}

func (x *SRSResourcetype) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourcetypeOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceTypeChoice == nil {
		return errors.Errorf("SRSResourceTypeChoice is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SRSResourcetypeOptPresentFlag = append(SRSResourcetypeOptPresentFlag, true)
	} else {
		SRSResourcetypeOptPresentFlag = append(SRSResourcetypeOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourcetypeOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SRSResourceTypeChoice.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSResourceTypeChoice marshal failed")
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

func (x *SRSResourcetype) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourcetypeOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SRSResourcetypeOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSResourceTypeChoice = new(SRSResourceTypeChoice)
	err = x.SRSResourceTypeChoice.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSResourceTypeChoice error")
	}

	// optional field (optPresentFlag index: 0)
	if SRSResourcetypeOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSResourcetypeExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
