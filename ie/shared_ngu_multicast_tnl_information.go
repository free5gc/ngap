package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SharedNGUMulticastTNLInformation struct {
	IPMulticastAddress *TransportLayerAddress
	IPSourceAddress    *TransportLayerAddress
	GTPTEID            *GTPTEID
	IEExtensions       *ProtocolExtensionContainerSharedNGUMulticastTNLInformationExtIEs // optional
}

func (x *SharedNGUMulticastTNLInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SharedNGUMulticastTNLInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.IPMulticastAddress == nil {
		return errors.Errorf("IPMulticastAddress is missing")
	}
	// mandatory field
	if x.IPSourceAddress == nil {
		return errors.Errorf("IPSourceAddress is missing")
	}
	// mandatory field
	if x.GTPTEID == nil {
		return errors.Errorf("GTPTEID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SharedNGUMulticastTNLInformationOptPresentFlag = append(SharedNGUMulticastTNLInformationOptPresentFlag, true)
	} else {
		SharedNGUMulticastTNLInformationOptPresentFlag = append(SharedNGUMulticastTNLInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SharedNGUMulticastTNLInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.IPMulticastAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IPMulticastAddress marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.IPSourceAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IPSourceAddress marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.GTPTEID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GTPTEID marshal failed")
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

func (x *SharedNGUMulticastTNLInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SharedNGUMulticastTNLInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SharedNGUMulticastTNLInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IPMulticastAddress = new(TransportLayerAddress)
	err = x.IPMulticastAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IPMulticastAddress error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IPSourceAddress = new(TransportLayerAddress)
	err = x.IPSourceAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IPSourceAddress error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GTPTEID = new(GTPTEID)
	err = x.GTPTEID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GTPTEID error")
	}

	// optional field (optPresentFlag index: 0)
	if SharedNGUMulticastTNLInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSharedNGUMulticastTNLInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
