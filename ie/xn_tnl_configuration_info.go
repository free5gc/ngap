package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type XnTNLConfigurationInfo struct {
	XnTransportLayerAddresses         *XnTLAs
	XnExtendedTransportLayerAddresses *XnExtTLAs                                              // optional
	IEExtensions                      *ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs // optional
}

func (x *XnTNLConfigurationInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	XnTNLConfigurationInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.XnTransportLayerAddresses == nil {
		return errors.Errorf("XnTransportLayerAddresses is missing")
	}
	// optional field
	if x.XnExtendedTransportLayerAddresses != nil {
		XnTNLConfigurationInfoOptPresentFlag = append(XnTNLConfigurationInfoOptPresentFlag, true)
	} else {
		XnTNLConfigurationInfoOptPresentFlag = append(XnTNLConfigurationInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		XnTNLConfigurationInfoOptPresentFlag = append(XnTNLConfigurationInfoOptPresentFlag, true)
	} else {
		XnTNLConfigurationInfoOptPresentFlag = append(XnTNLConfigurationInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(XnTNLConfigurationInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.XnTransportLayerAddresses.Write(pd)
	if err != nil {
		return errors.Wrap(err, "XnTransportLayerAddresses marshal failed")
	}

	// optional field
	if x.XnExtendedTransportLayerAddresses != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.XnExtendedTransportLayerAddresses.Write(pd)
		if err != nil {
			return errors.Wrap(err, "XnExtendedTransportLayerAddresses marshal failed")
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

func (x *XnTNLConfigurationInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	XnTNLConfigurationInfoOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&XnTNLConfigurationInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.XnTransportLayerAddresses = new(XnTLAs)
	err = x.XnTransportLayerAddresses.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode XnTransportLayerAddresses error")
	}

	// optional field (optPresentFlag index: 0)
	if XnTNLConfigurationInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.XnExtendedTransportLayerAddresses = new(XnExtTLAs)
		err = x.XnExtendedTransportLayerAddresses.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode XnExtendedTransportLayerAddresses error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if XnTNLConfigurationInfoOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerXnTNLConfigurationInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
