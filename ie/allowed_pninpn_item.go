package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	AllowedPNINPNItemPNINPNRestrictedPresentRestricted    aper.Enumerated = 0
	AllowedPNINPNItemPNINPNRestrictedPresentNotRestricted aper.Enumerated = 1
)

type AllowedPNINPNItem struct {
	PLMNIdentity          *PLMNIdentity
	PNINPNRestricted      *aper.Enumerated // valueExt,valueLB:0,valueUB:1
	AllowedCAGListPerPLMN *AllowedCAGListPerPLMN
	IEExtensions          *ProtocolExtensionContainerAllowedPNINPNItemExtIEs // optional
}

func (x *AllowedPNINPNItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AllowedPNINPNItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.PNINPNRestricted == nil {
		return errors.Errorf("PNINPNRestricted is missing")
	}
	// mandatory field
	if x.AllowedCAGListPerPLMN == nil {
		return errors.Errorf("AllowedCAGListPerPLMN is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AllowedPNINPNItemOptPresentFlag = append(AllowedPNINPNItemOptPresentFlag, true)
	} else {
		AllowedPNINPNItemOptPresentFlag = append(AllowedPNINPNItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AllowedPNINPNItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.PNINPNRestricted), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AllowedCAGListPerPLMN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AllowedCAGListPerPLMN marshal failed")
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

func (x *AllowedPNINPNItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AllowedPNINPNItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AllowedPNINPNItemOptPresentFlag, true)
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
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.PNINPNRestricted = new(aper.Enumerated)
	*(x.PNINPNRestricted), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AllowedCAGListPerPLMN = new(AllowedCAGListPerPLMN)
	err = x.AllowedCAGListPerPLMN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AllowedCAGListPerPLMN error")
	}

	// optional field (optPresentFlag index: 0)
	if AllowedPNINPNItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAllowedPNINPNItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
