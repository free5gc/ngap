package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type CGIEUTRA struct {
	PLMNIdentity        *PLMNIdentity
	EUTRAcellIdentifier *EUTRACellIdentifier
	IEExtensions        *ProtocolExtensionContainerCGIEUTRAExtIEs // optional
}

func (x *CGIEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CGIEUTRAOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.EUTRAcellIdentifier == nil {
		return errors.Errorf("EUTRAcellIdentifier is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CGIEUTRAOptPresentFlag = append(CGIEUTRAOptPresentFlag, true)
	} else {
		CGIEUTRAOptPresentFlag = append(CGIEUTRAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CGIEUTRAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRAcellIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRAcellIdentifier marshal failed")
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

func (x *CGIEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CGIEUTRAOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CGIEUTRAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNIdentity = new(PLMNIdentity)
	err = x.PLMNIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNIdentity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRAcellIdentifier = new(EUTRACellIdentifier)
	err = x.EUTRAcellIdentifier.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRAcellIdentifier error")
	}

	// optional field (optPresentFlag index: 0)
	if CGIEUTRAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCGIEUTRAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
