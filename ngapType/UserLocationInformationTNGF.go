package ngapType

// Need to import "github.com/free5gc/aper" if it uses "aper"

type UserLocationInformationTNGF struct {
	TNAPID       TNAPID
	IPAddress    TransportLayerAddress
	PortNumber   *PortNumber `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationTNGFExtIEs `aper:"optional"`
}
