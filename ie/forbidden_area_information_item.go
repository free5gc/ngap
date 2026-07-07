package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  *PLMNIdentity
	ForbiddenTACs *ForbiddenTACs
	IEExtensions  *ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs // optional
}

func (x *ForbiddenAreaInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ForbiddenAreaInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.ForbiddenTACs == nil {
		return errors.Errorf("ForbiddenTACs is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ForbiddenAreaInformationItemOptPresentFlag = append(ForbiddenAreaInformationItemOptPresentFlag, true)
	} else {
		ForbiddenAreaInformationItemOptPresentFlag = append(ForbiddenAreaInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ForbiddenAreaInformationItemOptPresentFlag, true)
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
	err = x.ForbiddenTACs.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ForbiddenTACs marshal failed")
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

func (x *ForbiddenAreaInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ForbiddenAreaInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ForbiddenAreaInformationItemOptPresentFlag, true)
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
	x.ForbiddenTACs = new(ForbiddenTACs)
	err = x.ForbiddenTACs.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ForbiddenTACs error")
	}

	// optional field (optPresentFlag index: 0)
	if ForbiddenAreaInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
