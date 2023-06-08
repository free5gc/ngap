package ngapType

// Need to import "github.com/free5gc/aper" if it uses "aper"

type ServedGUAMIItem struct {
	GUAMI         GUAMI                                            `aper:"valueExt"`
	BackupAMFName *AMFName                                         `aper:"optional"`
	IEExtensions  *ProtocolExtensionContainerServedGUAMIItemExtIEs `aper:"optional"`
}
