package ngapType

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type GlobalWAGFID struct {	/* Sequence Type (Extensible) */
	PLMNIdentity	PLMNIdentity
	WAGFID	WAGFID `aper:"valueLB:0,valueUB:1"`
	IEExtensions	*ProtocolExtensionContainerGlobalWAGFIDExtIEs `aper:"optional"`
}

