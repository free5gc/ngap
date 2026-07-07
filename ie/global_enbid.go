package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GlobalENBID struct {
	PLMNidentity *PLMNIdentity
	ENBID        *ENBID                                       // valueLB:0,valueUB:4
	IEExtensions *ProtocolExtensionContainerGlobalENBIDExtIEs // optional
}

func (x *GlobalENBID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GlobalENBIDOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNidentity == nil {
		return errors.Errorf("PLMNidentity is missing")
	}
	// mandatory field
	if x.ENBID == nil {
		return errors.Errorf("ENBID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		GlobalENBIDOptPresentFlag = append(GlobalENBIDOptPresentFlag, true)
	} else {
		GlobalENBIDOptPresentFlag = append(GlobalENBIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GlobalENBIDOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNidentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNidentity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ENBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ENBID marshal failed")
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

func (x *GlobalENBID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GlobalENBIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&GlobalENBIDOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNidentity = new(PLMNIdentity)
	err = x.PLMNidentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNidentity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ENBID = new(ENBID)
	err = x.ENBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ENBID error")
	}

	// optional field (optPresentFlag index: 0)
	if GlobalENBIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGlobalENBIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
