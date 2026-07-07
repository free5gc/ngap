package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSResourceIDItem struct {
	SRSResourceID *SRSResourceID
	IEExtensions  *ProtocolExtensionContainerSRSResourceIDItemExtIEs // optional
}

func (x *SRSResourceIDItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceIDItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceID == nil {
		return errors.Errorf("SRSResourceID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SRSResourceIDItemOptPresentFlag = append(SRSResourceIDItemOptPresentFlag, true)
	} else {
		SRSResourceIDItemOptPresentFlag = append(SRSResourceIDItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceIDItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSResourceID marshal failed")
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

func (x *SRSResourceIDItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceIDItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceIDItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSResourceID = new(SRSResourceID)
	err = x.SRSResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSResourceID error")
	}

	// optional field (optPresentFlag index: 0)
	if SRSResourceIDItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSResourceIDItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
