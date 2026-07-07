package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GTPTunnel struct {
	TransportLayerAddress *TransportLayerAddress
	GTPTEID               *GTPTEID
	IEExtensions          *ProtocolExtensionContainerGTPTunnelExtIEs // optional
}

func (x *GTPTunnel) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GTPTunnelOptPresentFlag := []bool{}
	// mandatory field
	if x.TransportLayerAddress == nil {
		return errors.Errorf("TransportLayerAddress is missing")
	}
	// mandatory field
	if x.GTPTEID == nil {
		return errors.Errorf("GTPTEID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		GTPTunnelOptPresentFlag = append(GTPTunnelOptPresentFlag, true)
	} else {
		GTPTunnelOptPresentFlag = append(GTPTunnelOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GTPTunnelOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TransportLayerAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TransportLayerAddress marshal failed")
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

func (x *GTPTunnel) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GTPTunnelOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&GTPTunnelOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TransportLayerAddress = new(TransportLayerAddress)
	err = x.TransportLayerAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TransportLayerAddress error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GTPTEID = new(GTPTEID)
	err = x.GTPTEID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GTPTEID error")
	}

	// optional field (optPresentFlag index: 0)
	if GTPTunnelOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGTPTunnelExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
