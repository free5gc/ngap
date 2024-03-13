package ngapType

import "github.com/free5gc/aper"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

const (
	TNGFIDPresentNothing int = iota /* No components present */
	TNGFIDPresentTNGFID
	TNGFIDPresentChoiceExtensions
)

type TNGFID struct {
	Present          int             /* Choice Type */
	TNGFID           *aper.BitString `aper:"sizeLB:32,sizeUB:32"`
	ChoiceExtensions *ProtocolIESingleContainerTNGFIDExtIEs
}
