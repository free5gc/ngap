package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PLMNSupportItem struct {
	PLMNIdentity     *PLMNIdentity
	SliceSupportList *SliceSupportList
	IEExtensions     *ProtocolExtensionContainerPLMNSupportItemExtIEs // optional
}

func (x *PLMNSupportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PLMNSupportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.SliceSupportList == nil {
		return errors.Errorf("SliceSupportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PLMNSupportItemOptPresentFlag = append(PLMNSupportItemOptPresentFlag, true)
	} else {
		PLMNSupportItemOptPresentFlag = append(PLMNSupportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PLMNSupportItemOptPresentFlag, true)
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
	err = x.SliceSupportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SliceSupportList marshal failed")
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

func (x *PLMNSupportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PLMNSupportItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PLMNSupportItemOptPresentFlag, true)
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
	x.SliceSupportList = new(SliceSupportList)
	err = x.SliceSupportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SliceSupportList error")
	}

	// optional field (optPresentFlag index: 0)
	if PLMNSupportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPLMNSupportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
