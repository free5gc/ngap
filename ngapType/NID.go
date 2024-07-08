package ngapType

import "github.com/free5gc/aper"

type NID struct {
	Value aper.BitString `aper:"sizeLB:44,sizeUB:44"`
}
