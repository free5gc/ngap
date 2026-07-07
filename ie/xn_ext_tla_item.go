package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type XnExtTLAItem struct {
	IPsecTLA     *TransportLayerAddress                        // optional
	GTPTLAs      *XnGTPTLAs                                    // optional
	IEExtensions *ProtocolExtensionContainerXnExtTLAItemExtIEs // optional
}

func (x *XnExtTLAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	XnExtTLAItemOptPresentFlag := []bool{}
	// optional field
	if x.IPsecTLA != nil {
		XnExtTLAItemOptPresentFlag = append(XnExtTLAItemOptPresentFlag, true)
	} else {
		XnExtTLAItemOptPresentFlag = append(XnExtTLAItemOptPresentFlag, false)
	}
	// optional field
	if x.GTPTLAs != nil {
		XnExtTLAItemOptPresentFlag = append(XnExtTLAItemOptPresentFlag, true)
	} else {
		XnExtTLAItemOptPresentFlag = append(XnExtTLAItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		XnExtTLAItemOptPresentFlag = append(XnExtTLAItemOptPresentFlag, true)
	} else {
		XnExtTLAItemOptPresentFlag = append(XnExtTLAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(XnExtTLAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.IPsecTLA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IPsecTLA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IPsecTLA marshal failed")
		}
	}

	// optional field
	if x.GTPTLAs != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.GTPTLAs.Write(pd)
		if err != nil {
			return errors.Wrap(err, "GTPTLAs marshal failed")
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

func (x *XnExtTLAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	XnExtTLAItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&XnExtTLAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if XnExtTLAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IPsecTLA = new(TransportLayerAddress)
		err = x.IPsecTLA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IPsecTLA error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if XnExtTLAItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.GTPTLAs = new(XnGTPTLAs)
		err = x.GTPTLAs.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode GTPTLAs error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if XnExtTLAItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerXnExtTLAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
