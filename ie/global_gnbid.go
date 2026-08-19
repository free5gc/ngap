package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GlobalGNBID struct {
	PLMNIdentity *PLMNIdentity
	GNBID        *GNBID                                       // valueLB:0,valueUB:1
	IEExtensions *ProtocolExtensionContainerGlobalGNBIDExtIEs // optional
}

func (x *GlobalGNBID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GlobalGNBIDOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.GNBID == nil {
		return errors.Errorf("GNBID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		GlobalGNBIDOptPresentFlag = append(GlobalGNBIDOptPresentFlag, true)
	} else {
		GlobalGNBIDOptPresentFlag = append(GlobalGNBIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GlobalGNBIDOptPresentFlag, true)
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
	err = x.GNBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GNBID marshal failed")
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

func (x *GlobalGNBID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GlobalGNBIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&GlobalGNBIDOptPresentFlag, true)
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
	x.GNBID = new(GNBID)
	err = x.GNBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GNBID error")
	}

	// optional field (optPresentFlag index: 0)
	if GlobalGNBIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGlobalGNBIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
