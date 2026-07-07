package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	CauseRadioNetworkPresentUnspecified                                              aper.Enumerated = 0
	CauseRadioNetworkPresentTxnrelocoverallExpiry                                    aper.Enumerated = 1
	CauseRadioNetworkPresentSuccessfulHandover                                       aper.Enumerated = 2
	CauseRadioNetworkPresentReleaseDueToNgranGeneratedReason                         aper.Enumerated = 3
	CauseRadioNetworkPresentReleaseDueTo5gcGeneratedReason                           aper.Enumerated = 4
	CauseRadioNetworkPresentHandoverCancelled                                        aper.Enumerated = 5
	CauseRadioNetworkPresentPartialHandover                                          aper.Enumerated = 6
	CauseRadioNetworkPresentHoFailureInTarget5GCNgranNodeOrTargetSystem              aper.Enumerated = 7
	CauseRadioNetworkPresentHoTargetNotAllowed                                       aper.Enumerated = 8
	CauseRadioNetworkPresentTngrelocoverallExpiry                                    aper.Enumerated = 9
	CauseRadioNetworkPresentTngrelocprepExpiry                                       aper.Enumerated = 10
	CauseRadioNetworkPresentCellNotAvailable                                         aper.Enumerated = 11
	CauseRadioNetworkPresentUnknownTargetID                                          aper.Enumerated = 12
	CauseRadioNetworkPresentNoRadioResourcesAvailableInTargetCell                    aper.Enumerated = 13
	CauseRadioNetworkPresentUnknownLocalUENGAPID                                     aper.Enumerated = 14
	CauseRadioNetworkPresentInconsistentRemoteUENGAPID                               aper.Enumerated = 15
	CauseRadioNetworkPresentHandoverDesirableForRadioReason                          aper.Enumerated = 16
	CauseRadioNetworkPresentTimeCriticalHandover                                     aper.Enumerated = 17
	CauseRadioNetworkPresentResourceOptimisationHandover                             aper.Enumerated = 18
	CauseRadioNetworkPresentReduceLoadInServingCell                                  aper.Enumerated = 19
	CauseRadioNetworkPresentUserInactivity                                           aper.Enumerated = 20
	CauseRadioNetworkPresentRadioConnectionWithUeLost                                aper.Enumerated = 21
	CauseRadioNetworkPresentRadioResourcesNotAvailable                               aper.Enumerated = 22
	CauseRadioNetworkPresentInvalidQosCombination                                    aper.Enumerated = 23
	CauseRadioNetworkPresentFailureInRadioInterfaceProcedure                         aper.Enumerated = 24
	CauseRadioNetworkPresentInteractionWithOtherProcedure                            aper.Enumerated = 25
	CauseRadioNetworkPresentUnknownPDUSessionID                                      aper.Enumerated = 26
	CauseRadioNetworkPresentUnkownQosFlowID                                          aper.Enumerated = 27
	CauseRadioNetworkPresentMultiplePDUSessionIDInstances                            aper.Enumerated = 28
	CauseRadioNetworkPresentMultipleQosFlowIDInstances                               aper.Enumerated = 29
	CauseRadioNetworkPresentEncryptionAndOrIntegrityProtectionAlgorithmsNotSupported aper.Enumerated = 30
	CauseRadioNetworkPresentNgIntraSystemHandoverTriggered                           aper.Enumerated = 31
	CauseRadioNetworkPresentNgInterSystemHandoverTriggered                           aper.Enumerated = 32
	CauseRadioNetworkPresentXnHandoverTriggered                                      aper.Enumerated = 33
	CauseRadioNetworkPresentNotSupported5QIValue                                     aper.Enumerated = 34
	CauseRadioNetworkPresentUeContextTransfer                                        aper.Enumerated = 35
	CauseRadioNetworkPresentImsVoiceEpsFallbackOrRatFallbackTriggered                aper.Enumerated = 36
	CauseRadioNetworkPresentUpIntegrityProtectionNotPossible                         aper.Enumerated = 37
	CauseRadioNetworkPresentUpConfidentialityProtectionNotPossible                   aper.Enumerated = 38
	CauseRadioNetworkPresentSliceNotSupported                                        aper.Enumerated = 39
	CauseRadioNetworkPresentUeInRrcInactiveStateNotReachable                         aper.Enumerated = 40
	CauseRadioNetworkPresentRedirection                                              aper.Enumerated = 41
	CauseRadioNetworkPresentResourcesNotAvailableForTheSlice                         aper.Enumerated = 42
	CauseRadioNetworkPresentUeMaxIntegrityProtectedDataRateReason                    aper.Enumerated = 43
	CauseRadioNetworkPresentReleaseDueToCnDetectedMobility                           aper.Enumerated = 44
	CauseRadioNetworkPresentN26InterfaceNotAvailable                                 aper.Enumerated = 45
	CauseRadioNetworkPresentReleaseDueToPreEmption                                   aper.Enumerated = 46
	CauseRadioNetworkPresentMultipleLocationReportingReferenceIDInstances            aper.Enumerated = 47
	CauseRadioNetworkPresentRsnNotAvailableForTheUp                                  aper.Enumerated = 48
	CauseRadioNetworkPresentNpnAccessDenied                                          aper.Enumerated = 49
	CauseRadioNetworkPresentCagOnlyAccessDenied                                      aper.Enumerated = 50
	CauseRadioNetworkPresentInsufficientUeCapabilities                               aper.Enumerated = 51
	CauseRadioNetworkPresentRedcapUeNotSupported                                     aper.Enumerated = 52
	CauseRadioNetworkPresentUnknownMBSSessionID                                      aper.Enumerated = 53
	CauseRadioNetworkPresentIndicatedMBSSessionAreaInformationNotServedByTheGNB      aper.Enumerated = 54
	CauseRadioNetworkPresentInconsistentSliceInfoForTheSession                       aper.Enumerated = 55
	CauseRadioNetworkPresentMisalignedAssociationForMulticastUnicast                 aper.Enumerated = 56
)

type CauseRadioNetwork struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:44
}

func (x *CauseRadioNetwork) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 44
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *CauseRadioNetwork) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 44
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
