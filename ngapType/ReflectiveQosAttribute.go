package ngapType

import "github.com/free5gc/aper"

// Need to import "github.com/free5gc/aper" if it uses "aper"

const (
	ReflectiveQosAttributePresentSubjectTo aper.Enumerated = 0
)

type ReflectiveQosAttribute struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
