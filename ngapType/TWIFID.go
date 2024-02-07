package ngapType

import "github.com/free5gc/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	TWIFIDPresentNothing	int = iota	/* No components present */
	TWIFIDPresentTWIFID
	TWIFIDPresentChoiceExtensions
)

type TWIFID struct {
	Present	int	/* Choice Type */
	TWIFID	*aper.BitString `aper:"sizeExt,sizeLB:32,sizeUB:32"`
	ChoiceExtensions	*ProtocolIESingleContainerTWIFIDExtIEs
}

