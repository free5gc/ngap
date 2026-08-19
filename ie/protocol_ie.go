package ie

import "github.com/free5gc/ngap/aper"

// dummy function to avoid unused error
func foo(args ...interface{}) {}

type ProtocolIE interface {
	// Elements of protocolIE-Field (SEQUENCE)
	// protocolIEID() *protocolIEID # depends on message type so passed as parameter in ie Write function
	// protocolIECriticality() *protocolIECriticality # depends on message type so passed as parameter in ie Write function
	// open type Value
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error

	// Complete IE encode/decode function
	WriteIE(*aper.PerBitData, ProtocolIEID, ProtocolIECriticality) error
	ReadIE(*aper.PerBitData) error
}

type ProtocolIECriticality struct {
	Value aper.Enumerated
}

const (
	CriticalityReject aper.Enumerated = 0
	CriticalityIgnore aper.Enumerated = 1
	CriticalityNotify aper.Enumerated = 2
)

func (c *ProtocolIECriticality) Write(pd *aper.PerBitData) error {
	var lb, ub int64 = 0, 2
	err := pd.WriteEnumerated(aper.Enumerated(c.Value), false, &lb, &ub)
	return err
}

func (c *ProtocolIECriticality) Read(pd *aper.PerBitData) error {
	var lb, ub int64 = 0, 2
	var err error
	c.Value, err = pd.ReadEnumerated(false, &lb, &ub)
	return err
}

type ProtocolIEID struct {
	Value int64
}

func (i *ProtocolIEID) Write(pd *aper.PerBitData) error {
	var lb, ub int64 = 0, 65535
	err := pd.WriteInteger(i.Value, false, &lb, &ub)
	return err
}

// Some Ies, e.g. CriticalityDiagnosticsIEItem, stil required generic Read Function for ProtocolIEID
func (i *ProtocolIEID) Read(pd *aper.PerBitData) error {
	var lb, ub int64 = 0, 65535
	var err error
	val, err := pd.ReadInteger(false, &lb, &ub)
	i.Value = int64(val)
	return err
}

// ProtocolIEID is the reference field for ProtocolIE-Field
func ReadProtocolIEID(pd *aper.PerBitData) (int64, error) {
	var lb, ub int64 = 0, 65535
	val, err := pd.ReadInteger(false, &lb, &ub)
	return val, err
}

const (
	ProtocolIEIDAllowedNSSAI                                  int64 = 0
	ProtocolIEIDAMFName                                       int64 = 1
	ProtocolIEIDAMFOverloadResponse                           int64 = 2
	ProtocolIEIDAMFSetID                                      int64 = 3
	ProtocolIEIDAMFTNLAssociationFailedToSetupList            int64 = 4
	ProtocolIEIDAMFTNLAssociationSetupList                    int64 = 5
	ProtocolIEIDAMFTNLAssociationToAddList                    int64 = 6
	ProtocolIEIDAMFTNLAssociationToRemoveList                 int64 = 7
	ProtocolIEIDAMFTNLAssociationToUpdateList                 int64 = 8
	ProtocolIEIDAMFTrafficLoadReductionIndication             int64 = 9
	ProtocolIEIDAMFUENGAPID                                   int64 = 10
	ProtocolIEIDAssistanceDataForPaging                       int64 = 11
	ProtocolIEIDBroadcastCancelledAreaList                    int64 = 12
	ProtocolIEIDBroadcastCompletedAreaList                    int64 = 13
	ProtocolIEIDCancelAllWarningMessages                      int64 = 14
	ProtocolIEIDCause                                         int64 = 15
	ProtocolIEIDCellIDListForRestart                          int64 = 16
	ProtocolIEIDConcurrentWarningMessageInd                   int64 = 17
	ProtocolIEIDCoreNetworkAssistanceInformationForInactive   int64 = 18
	ProtocolIEIDCriticalityDiagnostics                        int64 = 19
	ProtocolIEIDDataCodingScheme                              int64 = 20
	ProtocolIEIDDefaultPagingDRX                              int64 = 21
	ProtocolIEIDDirectForwardingPathAvailability              int64 = 22
	ProtocolIEIDEmergencyAreaIDListForRestart                 int64 = 23
	ProtocolIEIDEmergencyFallbackIndicator                    int64 = 24
	ProtocolIEIDEUTRACGI                                      int64 = 25
	ProtocolIEIDFiveGSTMSI                                    int64 = 26
	ProtocolIEIDGlobalRANNodeID                               int64 = 27
	ProtocolIEIDGUAMI                                         int64 = 28
	ProtocolIEIDHandoverType                                  int64 = 29
	ProtocolIEIDIMSVoiceSupportIndicator                      int64 = 30
	ProtocolIEIDIndexToRFSP                                   int64 = 31
	ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging    int64 = 32
	ProtocolIEIDLocationReportingRequestType                  int64 = 33
	ProtocolIEIDMaskedIMEISV                                  int64 = 34
	ProtocolIEIDMessageIdentifier                             int64 = 35
	ProtocolIEIDMobilityRestrictionList                       int64 = 36
	ProtocolIEIDNASC                                          int64 = 37
	ProtocolIEIDNASPDU                                        int64 = 38
	ProtocolIEIDNASSecurityParametersFromNGRAN                int64 = 39
	ProtocolIEIDNewAMFUENGAPID                                int64 = 40
	ProtocolIEIDNewSecurityContextInd                         int64 = 41
	ProtocolIEIDNGAPMessage                                   int64 = 42
	ProtocolIEIDNGRANCGI                                      int64 = 43
	ProtocolIEIDNGRANTraceID                                  int64 = 44
	ProtocolIEIDNRCGI                                         int64 = 45
	ProtocolIEIDNRPPaPDU                                      int64 = 46
	ProtocolIEIDNumberOfBroadcastsRequested                   int64 = 47
	ProtocolIEIDOldAMF                                        int64 = 48
	ProtocolIEIDOverloadStartNSSAIList                        int64 = 49
	ProtocolIEIDPagingDRX                                     int64 = 50
	ProtocolIEIDPagingOrigin                                  int64 = 51
	ProtocolIEIDPagingPriority                                int64 = 52
	ProtocolIEIDPDUSessionResourceAdmittedList                int64 = 53
	ProtocolIEIDPDUSessionResourceFailedToModifyListModRes    int64 = 54
	ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes     int64 = 55
	ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck      int64 = 56
	ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq      int64 = 57
	ProtocolIEIDPDUSessionResourceFailedToSetupListSURes      int64 = 58
	ProtocolIEIDPDUSessionResourceHandoverList                int64 = 59
	ProtocolIEIDPDUSessionResourceListCxtRelCpl               int64 = 60
	ProtocolIEIDPDUSessionResourceListHORqd                   int64 = 61
	ProtocolIEIDPDUSessionResourceModifyListModCfm            int64 = 62
	ProtocolIEIDPDUSessionResourceModifyListModInd            int64 = 63
	ProtocolIEIDPDUSessionResourceModifyListModReq            int64 = 64
	ProtocolIEIDPDUSessionResourceModifyListModRes            int64 = 65
	ProtocolIEIDPDUSessionResourceNotifyList                  int64 = 66
	ProtocolIEIDPDUSessionResourceReleasedListNot             int64 = 67
	ProtocolIEIDPDUSessionResourceReleasedListPSAck           int64 = 68
	ProtocolIEIDPDUSessionResourceReleasedListPSFail          int64 = 69
	ProtocolIEIDPDUSessionResourceReleasedListRelRes          int64 = 70
	ProtocolIEIDPDUSessionResourceSetupListCxtReq             int64 = 71
	ProtocolIEIDPDUSessionResourceSetupListCxtRes             int64 = 72
	ProtocolIEIDPDUSessionResourceSetupListHOReq              int64 = 73
	ProtocolIEIDPDUSessionResourceSetupListSUReq              int64 = 74
	ProtocolIEIDPDUSessionResourceSetupListSURes              int64 = 75
	ProtocolIEIDPDUSessionResourceToBeSwitchedDLList          int64 = 76
	ProtocolIEIDPDUSessionResourceSwitchedList                int64 = 77
	ProtocolIEIDPDUSessionResourceToReleaseListHOCmd          int64 = 78
	ProtocolIEIDPDUSessionResourceToReleaseListRelCmd         int64 = 79
	ProtocolIEIDPLMNSupportList                               int64 = 80
	ProtocolIEIDPWSFailedCellIDList                           int64 = 81
	ProtocolIEIDRANNodeName                                   int64 = 82
	ProtocolIEIDRANPagingPriority                             int64 = 83
	ProtocolIEIDRANStatusTransferTransparentContainer         int64 = 84
	ProtocolIEIDRANUENGAPID                                   int64 = 85
	ProtocolIEIDRelativeAMFCapacity                           int64 = 86
	ProtocolIEIDRepetitionPeriod                              int64 = 87
	ProtocolIEIDResetType                                     int64 = 88
	ProtocolIEIDRoutingID                                     int64 = 89
	ProtocolIEIDRRCEstablishmentCause                         int64 = 90
	ProtocolIEIDRRCInactiveTransitionReportRequest            int64 = 91
	ProtocolIEIDRRCState                                      int64 = 92
	ProtocolIEIDSecurityContext                               int64 = 93
	ProtocolIEIDSecurityKey                                   int64 = 94
	ProtocolIEIDSerialNumber                                  int64 = 95
	ProtocolIEIDServedGUAMIList                               int64 = 96
	ProtocolIEIDSliceSupportList                              int64 = 97
	ProtocolIEIDSONConfigurationTransferDL                    int64 = 98
	ProtocolIEIDSONConfigurationTransferUL                    int64 = 99
	ProtocolIEIDSourceAMFUENGAPID                             int64 = 100
	ProtocolIEIDSourceToTargetTransparentContainer            int64 = 101
	ProtocolIEIDSupportedTAList                               int64 = 102
	ProtocolIEIDTAIListForPaging                              int64 = 103
	ProtocolIEIDTAIListForRestart                             int64 = 104
	ProtocolIEIDTargetID                                      int64 = 105
	ProtocolIEIDTargetToSourceTransparentContainer            int64 = 106
	ProtocolIEIDTimeToWait                                    int64 = 107
	ProtocolIEIDTraceActivation                               int64 = 108
	ProtocolIEIDTraceCollectionEntityIPAddress                int64 = 109
	ProtocolIEIDUEAggregateMaximumBitRate                     int64 = 110
	ProtocolIEIDUEAssociatedLogicalNGConnectionList           int64 = 111
	ProtocolIEIDUEContextRequest                              int64 = 112
	ProtocolIEIDUENGAPIDs                                     int64 = 114
	ProtocolIEIDUEPagingIdentity                              int64 = 115
	ProtocolIEIDUEPresenceInAreaOfInterestList                int64 = 116
	ProtocolIEIDUERadioCapability                             int64 = 117
	ProtocolIEIDUERadioCapabilityForPaging                    int64 = 118
	ProtocolIEIDUESecurityCapabilities                        int64 = 119
	ProtocolIEIDUnavailableGUAMIList                          int64 = 120
	ProtocolIEIDUserLocationInformation                       int64 = 121
	ProtocolIEIDWarningAreaList                               int64 = 122
	ProtocolIEIDWarningMessageContents                        int64 = 123
	ProtocolIEIDWarningSecurityInfo                           int64 = 124
	ProtocolIEIDWarningType                                   int64 = 125
	ProtocolIEIDAdditionalULNGUUPTNLInformation               int64 = 126
	ProtocolIEIDDataForwardingNotPossible                     int64 = 127
	ProtocolIEIDDLNGUUPTNLInformation                         int64 = 128
	ProtocolIEIDNetworkInstance                               int64 = 129
	ProtocolIEIDPDUSessionAggregateMaximumBitRate             int64 = 130
	ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm    int64 = 131
	ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail    int64 = 132
	ProtocolIEIDPDUSessionResourceListCxtRelReq               int64 = 133
	ProtocolIEIDPDUSessionType                                int64 = 134
	ProtocolIEIDQosFlowAddOrModifyRequestList                 int64 = 135
	ProtocolIEIDQosFlowSetupRequestList                       int64 = 136
	ProtocolIEIDQosFlowToReleaseList                          int64 = 137
	ProtocolIEIDSecurityIndication                            int64 = 138
	ProtocolIEIDULNGUUPTNLInformation                         int64 = 139
	ProtocolIEIDULNGUUPTNLModifyList                          int64 = 140
	ProtocolIEIDWarningAreaCoordinates                        int64 = 141
	ProtocolIEIDPDUSessionResourceSecondaryRATUsageList       int64 = 142
	ProtocolIEIDHandoverFlag                                  int64 = 143
	ProtocolIEIDSecondaryRATUsageInformation                  int64 = 144
	ProtocolIEIDPDUSessionResourceReleaseResponseTransfer     int64 = 145
	ProtocolIEIDRedirectionVoiceFallback                      int64 = 146
	ProtocolIEIDUERetentionInformation                        int64 = 147
	ProtocolIEIDSNSSAI                                        int64 = 148
	ProtocolIEIDPSCellInformation                             int64 = 149
	ProtocolIEIDLastEUTRANPLMNIdentity                        int64 = 150
	ProtocolIEIDMaximumIntegrityProtectedDataRateDL           int64 = 151
	ProtocolIEIDAdditionalDLForwardingUPTNLInformation        int64 = 152
	ProtocolIEIDAdditionalDLUPTNLInformationForHOList         int64 = 153
	ProtocolIEIDAdditionalNGUUPTNLInformation                 int64 = 154
	ProtocolIEIDAdditionalDLQosFlowPerTNLInformation          int64 = 155
	ProtocolIEIDSecurityResult                                int64 = 156
	ProtocolIEIDENDCSONConfigurationTransferDL                int64 = 157
	ProtocolIEIDENDCSONConfigurationTransferUL                int64 = 158
	ProtocolIEIDOldAssociatedQosFlowListULendmarkerexpected   int64 = 159
	ProtocolIEIDCNTypeRestrictionsForEquivalent               int64 = 160
	ProtocolIEIDCNTypeRestrictionsForServing                  int64 = 161
	ProtocolIEIDNewGUAMI                                      int64 = 162
	ProtocolIEIDULForwarding                                  int64 = 163
	ProtocolIEIDULForwardingUPTNLInformation                  int64 = 164
	ProtocolIEIDCNAssistedRANTuning                           int64 = 165
	ProtocolIEIDCommonNetworkInstance                         int64 = 166
	ProtocolIEIDNGRANTNLAssociationToRemoveList               int64 = 167
	ProtocolIEIDTNLAssociationTransportLayerAddressNGRAN      int64 = 168
	ProtocolIEIDEndpointIPAddressAndPort                      int64 = 169
	ProtocolIEIDLocationReportingAdditionalInfo               int64 = 170
	ProtocolIEIDSourceToTargetAMFInformationReroute           int64 = 171
	ProtocolIEIDAdditionalULForwardingUPTNLInformation        int64 = 172
	ProtocolIEIDSCTPTLAs                                      int64 = 173
	ProtocolIEIDSelectedPLMNIdentity                          int64 = 174
	ProtocolIEIDRIMInformationTransfer                        int64 = 175
	ProtocolIEIDGUAMIType                                     int64 = 176
	ProtocolIEIDSRVCCOperationPossible                        int64 = 177
	ProtocolIEIDTargetRNCID                                   int64 = 178
	ProtocolIEIDRATInformation                                int64 = 179
	ProtocolIEIDExtendedRATRestrictionInformation             int64 = 180
	ProtocolIEIDQosMonitoringRequest                          int64 = 181
	ProtocolIEIDSgNBUEX2APID                                  int64 = 182
	ProtocolIEIDAdditionalRedundantDLNGUUPTNLInformation      int64 = 183
	ProtocolIEIDAdditionalRedundantDLQosFlowPerTNLInformation int64 = 184
	ProtocolIEIDAdditionalRedundantNGUUPTNLInformation        int64 = 185
	ProtocolIEIDAdditionalRedundantULNGUUPTNLInformation      int64 = 186
	ProtocolIEIDCNPacketDelayBudgetDL                         int64 = 187
	ProtocolIEIDCNPacketDelayBudgetUL                         int64 = 188
	ProtocolIEIDExtendedPacketDelayBudget                     int64 = 189
	ProtocolIEIDRedundantCommonNetworkInstance                int64 = 190
	ProtocolIEIDRedundantDLNGUTNLInformationReused            int64 = 191
	ProtocolIEIDRedundantDLNGUUPTNLInformation                int64 = 192
	ProtocolIEIDRedundantDLQosFlowPerTNLInformation           int64 = 193
	ProtocolIEIDRedundantQosFlowIndicator                     int64 = 194
	ProtocolIEIDRedundantULNGUUPTNLInformation                int64 = 195
	ProtocolIEIDTSCTrafficCharacteristics                     int64 = 196
	ProtocolIEIDRedundantPDUSessionInformation                int64 = 197
	ProtocolIEIDUsedRSNInformation                            int64 = 198
	ProtocolIEIDIABAuthorized                                 int64 = 199
	ProtocolIEIDIABSupported                                  int64 = 200
	ProtocolIEIDIABNodeIndication                             int64 = 201
	ProtocolIEIDNBIoTPagingDRX                                int64 = 202
	ProtocolIEIDNBIoTPagingEDRXInfo                           int64 = 203
	ProtocolIEIDNBIoTDefaultPagingDRX                         int64 = 204
	ProtocolIEIDEnhancedCoverageRestriction                   int64 = 205
	ProtocolIEIDExtendedConnectedTime                         int64 = 206
	ProtocolIEIDPagingAssisDataforCEcapabUE                   int64 = 207
	ProtocolIEIDWUSAssistanceInformation                      int64 = 208
	ProtocolIEIDUEDifferentiationInfo                         int64 = 209
	ProtocolIEIDNBIoTUEPriority                               int64 = 210
	ProtocolIEIDULCPSecurityInformation                       int64 = 211
	ProtocolIEIDDLCPSecurityInformation                       int64 = 212
	ProtocolIEIDTAI                                           int64 = 213
	ProtocolIEIDUERadioCapabilityForPagingOfNBIoT             int64 = 214
	ProtocolIEIDLTEV2XServicesAuthorized                      int64 = 215
	ProtocolIEIDNRV2XServicesAuthorized                       int64 = 216
	ProtocolIEIDLTEUESidelinkAggregateMaximumBitrate          int64 = 217
	ProtocolIEIDNRUESidelinkAggregateMaximumBitrate           int64 = 218
	ProtocolIEIDPC5QoSParameters                              int64 = 219
	ProtocolIEIDAlternativeQoSParaSetList                     int64 = 220
	ProtocolIEIDCurrentQoSParaSetIndex                        int64 = 221
	ProtocolIEIDCEmodeBrestricted                             int64 = 222
	ProtocolIEIDEUTRAPagingeDRXInformation                    int64 = 223
	ProtocolIEIDCEmodeBSupportIndicator                       int64 = 224
	ProtocolIEIDLTEMIndication                                int64 = 225
	ProtocolIEIDEndIndication                                 int64 = 226
	ProtocolIEIDEDTSession                                    int64 = 227
	ProtocolIEIDUECapabilityInfoRequest                       int64 = 228
	ProtocolIEIDPDUSessionResourceFailedToResumeListRESReq    int64 = 229
	ProtocolIEIDPDUSessionResourceFailedToResumeListRESRes    int64 = 230
	ProtocolIEIDPDUSessionResourceSuspendListSUSReq           int64 = 231
	ProtocolIEIDPDUSessionResourceResumeListRESReq            int64 = 232
	ProtocolIEIDPDUSessionResourceResumeListRESRes            int64 = 233
	ProtocolIEIDUEUPCIoTSupport                               int64 = 234
	ProtocolIEIDSuspendRequestIndication                      int64 = 235
	ProtocolIEIDSuspendResponseIndication                     int64 = 236
	ProtocolIEIDRRCResumeCause                                int64 = 237
	ProtocolIEIDRGLevelWirelineAccessCharacteristics          int64 = 238
	ProtocolIEIDWAGFIdentityInformation                       int64 = 239
	ProtocolIEIDGlobalTNGFID                                  int64 = 240
	ProtocolIEIDGlobalTWIFID                                  int64 = 241
	ProtocolIEIDGlobalWAGFID                                  int64 = 242
	ProtocolIEIDUserLocationInformationWAGF                   int64 = 243
	ProtocolIEIDUserLocationInformationTNGF                   int64 = 244
	ProtocolIEIDAuthenticatedIndication                       int64 = 245
	ProtocolIEIDTNGFIdentityInformation                       int64 = 246
	ProtocolIEIDTWIFIdentityInformation                       int64 = 247
	ProtocolIEIDUserLocationInformationTWIF                   int64 = 248
	ProtocolIEIDDataForwardingResponseERABList                int64 = 249
	ProtocolIEIDIntersystemSONConfigurationTransferDL         int64 = 250
	ProtocolIEIDIntersystemSONConfigurationTransferUL         int64 = 251
	ProtocolIEIDSONInformationReport                          int64 = 252
	ProtocolIEIDUEHistoryInformationFromTheUE                 int64 = 253
	ProtocolIEIDManagementBasedMDTPLMNList                    int64 = 254
	ProtocolIEIDMDTConfiguration                              int64 = 255
	ProtocolIEIDPrivacyIndicator                              int64 = 256
	ProtocolIEIDTraceCollectionEntityURI                      int64 = 257
	ProtocolIEIDNPNSupport                                    int64 = 258
	ProtocolIEIDNPNAccessInformation                          int64 = 259
	ProtocolIEIDNPNPagingAssistanceInformation                int64 = 260
	ProtocolIEIDNPNMobilityInformation                        int64 = 261
	ProtocolIEIDTargettoSourceFailureTransparentContainer     int64 = 262
	ProtocolIEIDNID                                           int64 = 263
	ProtocolIEIDUERadioCapabilityID                           int64 = 264
	ProtocolIEIDUERadioCapabilityEUTRAFormat                  int64 = 265
	ProtocolIEIDDAPSRequestInfo                               int64 = 266
	ProtocolIEIDDAPSResponseInfoList                          int64 = 267
	ProtocolIEIDEarlyStatusTransferTransparentContainer       int64 = 268
	ProtocolIEIDNotifySourceNGRANNode                         int64 = 269
	ProtocolIEIDExtendedSliceSupportList                      int64 = 270
	ProtocolIEIDExtendedTAISliceSupportList                   int64 = 271
	ProtocolIEIDConfiguredTACIndication                       int64 = 272
	ProtocolIEIDExtendedRANNodeName                           int64 = 273
	ProtocolIEIDExtendedAMFName                               int64 = 274
	ProtocolIEIDGlobalCableID                                 int64 = 275
	ProtocolIEIDQosMonitoringReportingFrequency               int64 = 276
	ProtocolIEIDQosFlowParametersList                         int64 = 277
	ProtocolIEIDQosFlowFeedbackList                           int64 = 278
	ProtocolIEIDBurstArrivalTimeDownlink                      int64 = 279
	ProtocolIEIDExtendedUEIdentityIndexValue                  int64 = 280
	ProtocolIEIDPduSessionExpectedUEActivityBehaviour         int64 = 281
	ProtocolIEIDMicoAllPLMN                                   int64 = 282
	ProtocolIEIDQosFlowFailedToSetupList                      int64 = 283
	ProtocolIEIDSourceTNLAddrInfo                             int64 = 284
	ProtocolIEIDExtendedReportIntervalMDT                     int64 = 285
	ProtocolIEIDSourceNodeID                                  int64 = 286
	ProtocolIEIDNRNTNTAIInformation                           int64 = 287
	ProtocolIEIDUEContextReferenceAtSource                    int64 = 288
	ProtocolIEIDLastVisitedPSCellList                         int64 = 289
	ProtocolIEIDIntersystemSONInformationRequest              int64 = 290
	ProtocolIEIDIntersystemSONInformationReply                int64 = 291
	ProtocolIEIDEnergySavingIndication                        int64 = 292
	ProtocolIEIDIntersystemResourceStatusUpdate               int64 = 293
	ProtocolIEIDSuccessfulHandoverReportList                  int64 = 294
	ProtocolIEIDMBSAreaSessionID                              int64 = 295
	ProtocolIEIDMBSQoSFlowsToBeSetupList                      int64 = 296
	ProtocolIEIDMBSQoSFlowsToBeSetupModList                   int64 = 297
	ProtocolIEIDMBSServiceArea                                int64 = 298
	ProtocolIEIDMBSSessionID                                  int64 = 299
	ProtocolIEIDMBSDistributionReleaseRequestTransfer         int64 = 300
	ProtocolIEIDMBSDistributionSetupRequestTransfer           int64 = 301
	ProtocolIEIDMBSDistributionSetupResponseTransfer          int64 = 302
	ProtocolIEIDMBSDistributionSetupUnsuccessfulTransfer      int64 = 303
	ProtocolIEIDMulticastSessionActivationRequestTransfer     int64 = 304
	ProtocolIEIDMulticastSessionDeactivationRequestTransfer   int64 = 305
	ProtocolIEIDMulticastSessionUpdateRequestTransfer         int64 = 306
	ProtocolIEIDMulticastGroupPagingAreaList                  int64 = 307
	ProtocolIEIDMBSSupportIndicator                           int64 = 309
	ProtocolIEIDMBSSessionFailedtoSetupList                   int64 = 310
	ProtocolIEIDMBSSessionFailedtoSetuporModifyList           int64 = 311
	ProtocolIEIDMBSSessionSetupResponseList                   int64 = 312
	ProtocolIEIDMBSSessionSetuporModifyResponseList           int64 = 313
	ProtocolIEIDMBSSessionSetupFailureTransfer                int64 = 314
	ProtocolIEIDMBSSessionSetupRequestTransfer                int64 = 315
	ProtocolIEIDMBSSessionSetupResponseTransfer               int64 = 316
	ProtocolIEIDMBSSessionToReleaseList                       int64 = 317
	ProtocolIEIDMBSSessionSetupRequestList                    int64 = 318
	ProtocolIEIDMBSSessionSetuporModifyRequestList            int64 = 319
	ProtocolIEIDMBSActiveSessionInformationSourcetoTargetList int64 = 323
	ProtocolIEIDMBSActiveSessionInformationTargettoSourceList int64 = 324
	ProtocolIEIDOnboardingSupport                             int64 = 325
	ProtocolIEIDTimeSyncAssistanceInfo                        int64 = 326
	ProtocolIEIDSurvivalTime                                  int64 = 327
	ProtocolIEIDQMCConfigInfo                                 int64 = 328
	ProtocolIEIDQMCDeactivation                               int64 = 329
	ProtocolIEIDPDUSessionPairID                              int64 = 331
	ProtocolIEIDNRPagingeDRXInformation                       int64 = 332
	ProtocolIEIDRedCapIndication                              int64 = 333
	ProtocolIEIDTargetNSSAIInformation                        int64 = 334
	ProtocolIEIDUESliceMaximumBitRateList                     int64 = 335
	ProtocolIEIDM4ReportAmount                                int64 = 336
	ProtocolIEIDM5ReportAmount                                int64 = 337
	ProtocolIEIDM6ReportAmount                                int64 = 338
	ProtocolIEIDM7ReportAmount                                int64 = 339
	ProtocolIEIDIncludeBeamMeasurementsIndication             int64 = 340
	ProtocolIEIDExcessPacketDelayThresholdConfiguration       int64 = 341
	ProtocolIEIDPagingCause                                   int64 = 342
	ProtocolIEIDPagingCauseIndicationForVoiceService          int64 = 343
	ProtocolIEIDPEIPSassistanceInformation                    int64 = 344
	ProtocolIEIDFiveGProSeAuthorized                          int64 = 345
	ProtocolIEIDFiveGProSeUEPC5AggregateMaximumBitRate        int64 = 346
	ProtocolIEIDFiveGProSePC5QoSParameters                    int64 = 347
	ProtocolIEIDMBSSessionModificationFailureTransfer         int64 = 348
	ProtocolIEIDMBSSessionModificationRequestTransfer         int64 = 349
	ProtocolIEIDMBSSessionModificationResponseTransfer        int64 = 350
	ProtocolIEIDMBSQoSFlowToReleaseList                       int64 = 351
	ProtocolIEIDMBSSessionTNLInfo5GC                          int64 = 352
	ProtocolIEIDTAINSAGSupportList                            int64 = 353
	ProtocolIEIDSourceNodeTNLAddrInfo                         int64 = 354
	ProtocolIEIDNGAPIESupportInformationRequestList           int64 = 355
	ProtocolIEIDNGAPIESupportInformationResponseList          int64 = 356
	ProtocolIEIDMBSSessionFSAIDList                           int64 = 357
	ProtocolIEIDMBSSessionReleaseResponseTransfer             int64 = 358
	ProtocolIEIDManagementBasedMDTPLMNModificationList        int64 = 359
	ProtocolIEIDEarlyMeasurement                              int64 = 360
	ProtocolIEIDBeamMeasurementsReportConfiguration           int64 = 361
	ProtocolIEIDHFCNodeIDNew                                  int64 = 362
	ProtocolIEIDGlobalCableIDNew                              int64 = 363
	ProtocolIEIDTargetHomeENBID                               int64 = 364
	ProtocolIEIDHashedUEIdentityIndexValue                    int64 = 365
	ProtocolIEIDExtendedMobilityInformation                   int64 = 366
)
