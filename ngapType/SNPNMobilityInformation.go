package ngapType

type SNPNMobilityInformation struct {
	ServingNID   NID
	IEExtensions *ProtocolExtensionContainerSNPNMobilityInformationExtIEs `aper:"optional"`
}
