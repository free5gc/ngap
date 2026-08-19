package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LAI struct {
	PLMNidentity *PLMNIdentity
	LAC          *LAC
	IEExtensions *ProtocolExtensionContainerLAIExtIEs // optional
}

func (x *LAI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LAIOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNidentity == nil {
		return errors.Errorf("PLMNidentity is missing")
	}
	// mandatory field
	if x.LAC == nil {
		return errors.Errorf("LAC is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		LAIOptPresentFlag = append(LAIOptPresentFlag, true)
	} else {
		LAIOptPresentFlag = append(LAIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LAIOptPresentFlag, true)
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
	err = x.LAC.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LAC marshal failed")
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

func (x *LAI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LAIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&LAIOptPresentFlag, true)
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
	x.LAC = new(LAC)
	err = x.LAC.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LAC error")
	}

	// optional field (optPresentFlag index: 0)
	if LAIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLAIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
