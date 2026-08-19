package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GlobalNgENBID struct {
	PLMNIdentity *PLMNIdentity
	NgENBID      *NgENBID                                       // valueLB:0,valueUB:3
	IEExtensions *ProtocolExtensionContainerGlobalNgENBIDExtIEs // optional
}

func (x *GlobalNgENBID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GlobalNgENBIDOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.NgENBID == nil {
		return errors.Errorf("NgENBID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		GlobalNgENBIDOptPresentFlag = append(GlobalNgENBIDOptPresentFlag, true)
	} else {
		GlobalNgENBIDOptPresentFlag = append(GlobalNgENBIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GlobalNgENBIDOptPresentFlag, true)
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
	err = x.NgENBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NgENBID marshal failed")
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

func (x *GlobalNgENBID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GlobalNgENBIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&GlobalNgENBIDOptPresentFlag, true)
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
	x.NgENBID = new(NgENBID)
	err = x.NgENBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NgENBID error")
	}

	// optional field (optPresentFlag index: 0)
	if GlobalNgENBIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGlobalNgENBIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
