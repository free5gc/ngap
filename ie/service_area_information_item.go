package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ServiceAreaInformationItem struct {
	PLMNIdentity   *PLMNIdentity
	AllowedTACs    *AllowedTACs                                                // optional
	NotAllowedTACs *NotAllowedTACs                                             // optional
	IEExtensions   *ProtocolExtensionContainerServiceAreaInformationItemExtIEs // optional
}

func (x *ServiceAreaInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ServiceAreaInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// optional field
	if x.AllowedTACs != nil {
		ServiceAreaInformationItemOptPresentFlag = append(ServiceAreaInformationItemOptPresentFlag, true)
	} else {
		ServiceAreaInformationItemOptPresentFlag = append(ServiceAreaInformationItemOptPresentFlag, false)
	}
	// optional field
	if x.NotAllowedTACs != nil {
		ServiceAreaInformationItemOptPresentFlag = append(ServiceAreaInformationItemOptPresentFlag, true)
	} else {
		ServiceAreaInformationItemOptPresentFlag = append(ServiceAreaInformationItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ServiceAreaInformationItemOptPresentFlag = append(ServiceAreaInformationItemOptPresentFlag, true)
	} else {
		ServiceAreaInformationItemOptPresentFlag = append(ServiceAreaInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ServiceAreaInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// optional field
	if x.AllowedTACs != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AllowedTACs.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AllowedTACs marshal failed")
		}
	}

	// optional field
	if x.NotAllowedTACs != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NotAllowedTACs.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NotAllowedTACs marshal failed")
		}
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

func (x *ServiceAreaInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ServiceAreaInformationItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&ServiceAreaInformationItemOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if ServiceAreaInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AllowedTACs = new(AllowedTACs)
		err = x.AllowedTACs.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AllowedTACs error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ServiceAreaInformationItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.NotAllowedTACs = new(NotAllowedTACs)
		err = x.NotAllowedTACs.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NotAllowedTACs error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ServiceAreaInformationItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerServiceAreaInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
