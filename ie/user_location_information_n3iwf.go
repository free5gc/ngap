package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UserLocationInformationN3IWF struct {
	IPAddress    *TransportLayerAddress
	PortNumber   *PortNumber
	IEExtensions *ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs // optional
}

func (x *UserLocationInformationN3IWF) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UserLocationInformationN3IWFOptPresentFlag := []bool{}
	// mandatory field
	if x.IPAddress == nil {
		return errors.Errorf("IPAddress is missing")
	}
	// mandatory field
	if x.PortNumber == nil {
		return errors.Errorf("PortNumber is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UserLocationInformationN3IWFOptPresentFlag = append(UserLocationInformationN3IWFOptPresentFlag, true)
	} else {
		UserLocationInformationN3IWFOptPresentFlag = append(UserLocationInformationN3IWFOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UserLocationInformationN3IWFOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.IPAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IPAddress marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PortNumber.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PortNumber marshal failed")
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

func (x *UserLocationInformationN3IWF) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UserLocationInformationN3IWFOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UserLocationInformationN3IWFOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IPAddress = new(TransportLayerAddress)
	err = x.IPAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IPAddress error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PortNumber = new(PortNumber)
	err = x.PortNumber.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PortNumber error")
	}

	// optional field (optPresentFlag index: 0)
	if UserLocationInformationN3IWFOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
