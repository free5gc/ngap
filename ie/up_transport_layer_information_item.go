package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UPTransportLayerInformationItem struct {
	NGUUPTNLInformation *UPTransportLayerInformation                                     // valueLB:0,valueUB:1
	IEExtensions        *ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs // optional
}

func (x *UPTransportLayerInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UPTransportLayerInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NGUUPTNLInformation == nil {
		return errors.Errorf("NGUUPTNLInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UPTransportLayerInformationItemOptPresentFlag = append(UPTransportLayerInformationItemOptPresentFlag, true)
	} else {
		UPTransportLayerInformationItemOptPresentFlag = append(UPTransportLayerInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UPTransportLayerInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGUUPTNLInformation marshal failed")
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

func (x *UPTransportLayerInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UPTransportLayerInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UPTransportLayerInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.NGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGUUPTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if UPTransportLayerInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUPTransportLayerInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
