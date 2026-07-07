package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GlobalN3IWFID struct {
	PLMNIdentity *PLMNIdentity
	N3IWFID      *N3IWFID                                       // valueLB:0,valueUB:1
	IEExtensions *ProtocolExtensionContainerGlobalN3IWFIDExtIEs // optional
}

func (x *GlobalN3IWFID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GlobalN3IWFIDOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.N3IWFID == nil {
		return errors.Errorf("N3IWFID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		GlobalN3IWFIDOptPresentFlag = append(GlobalN3IWFIDOptPresentFlag, true)
	} else {
		GlobalN3IWFIDOptPresentFlag = append(GlobalN3IWFIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GlobalN3IWFIDOptPresentFlag, true)
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
	err = x.N3IWFID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "N3IWFID marshal failed")
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

func (x *GlobalN3IWFID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GlobalN3IWFIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&GlobalN3IWFIDOptPresentFlag, true)
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
	x.N3IWFID = new(N3IWFID)
	err = x.N3IWFID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode N3IWFID error")
	}

	// optional field (optPresentFlag index: 0)
	if GlobalN3IWFIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGlobalN3IWFIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
