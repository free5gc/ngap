package ngapType

const (
	NPNSupportPresentNothing int = iota /* No components present */
	NPNSupportPresentSNPN
	NPNSupportPresentChoiceExtensions
)

type NPNSupport struct {
	Present          int
	SNPN             *NID
	ChoiceExtensions *ProtocolIESingleContainerNPNSupportExtIEs
}
