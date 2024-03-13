package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type GlobalLineID struct { /* Sequence Type (Extensible) */
	GlobalLineIdentity GlobalLineIdentity
	LineType           *LineType                                     `aper:"valueExt,valueLB:0,valueUB:1,optional"`
	IEExtensions       *ProtocolExtensionContainerGlobalLineIDExtIEs `aper:"optional"`
}
