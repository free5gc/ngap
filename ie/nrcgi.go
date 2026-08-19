package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NRCGI struct {
	PLMNIdentity   *PLMNIdentity
	NRCellIdentity *NRCellIdentity
	IEExtensions   *ProtocolExtensionContainerNRCGIExtIEs // optional
}

func (x *NRCGI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRCGIOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.NRCellIdentity == nil {
		return errors.Errorf("NRCellIdentity is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NRCGIOptPresentFlag = append(NRCGIOptPresentFlag, true)
	} else {
		NRCGIOptPresentFlag = append(NRCGIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRCGIOptPresentFlag, true)
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
	err = x.NRCellIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRCellIdentity marshal failed")
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

func (x *NRCGI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRCGIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NRCGIOptPresentFlag, true)
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
	x.NRCellIdentity = new(NRCellIdentity)
	err = x.NRCellIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRCellIdentity error")
	}

	// optional field (optPresentFlag index: 0)
	if NRCGIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNRCGIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
