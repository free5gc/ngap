package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type UserLocationInformationTWIF struct {	/* Sequence Type (Extensible) */
	TWAPID	TWAPID
	IPAddress	TransportLayerAddress
	PortNumber	*PortNumber `aper:"optional"`
	IEExtensions	*ProtocolExtensionContainerUserLocationInformationTWIFExtIEs `aper:"optional"`
}

