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
	ProtocolIEIDCause                                      int64 = 0
	ProtocolIEIDCriticalityDiagnostics                     int64 = 1
	ProtocolIEIDLMFUEMeasurementID                         int64 = 2
	ProtocolIEIDReportCharacteristics                      int64 = 3
	ProtocolIEIDMeasurementPeriodicity                     int64 = 4
	ProtocolIEIDMeasurementQuantities                      int64 = 5
	ProtocolIEIDRANUEMeasurementID                         int64 = 6
	ProtocolIEIDECIDMeasurementResult                      int64 = 7
	ProtocolIEIDOTDOACells                                 int64 = 8
	ProtocolIEIDOTDOAInformationTypeGroup                  int64 = 9
	ProtocolIEIDOTDOAInformationTypeItem                   int64 = 10
	ProtocolIEIDMeasurementQuantitiesItem                  int64 = 11
	ProtocolIEIDRequestedSRSTransmissionCharacteristics    int64 = 12
	ProtocolIEIDCellPortionID                              int64 = 14
	ProtocolIEIDOtherRATMeasurementQuantities              int64 = 15
	ProtocolIEIDOtherRATMeasurementQuantitiesItem          int64 = 16
	ProtocolIEIDOtherRATMeasurementResult                  int64 = 17
	ProtocolIEIDWLANMeasurementQuantities                  int64 = 19
	ProtocolIEIDWLANMeasurementQuantitiesItem              int64 = 20
	ProtocolIEIDWLANMeasurementResult                      int64 = 21
	ProtocolIEIDTDDConfigEUTRAItem                         int64 = 22
	ProtocolIEIDAssistanceInformation                      int64 = 23
	ProtocolIEIDBroadcast                                  int64 = 24
	ProtocolIEIDAssistanceInformationFailureList           int64 = 25
	ProtocolIEIDSRSConfiguration                           int64 = 26
	ProtocolIEIDMeasurementResult                          int64 = 27
	ProtocolIEIDTRPID                                      int64 = 28
	ProtocolIEIDTRPInformationTypeListTRPReq               int64 = 29
	ProtocolIEIDTRPInformationListTRPResp                  int64 = 30
	ProtocolIEIDMeasurementBeamInfoRequest                 int64 = 31
	ProtocolIEIDResultSSRSRP                               int64 = 32
	ProtocolIEIDResultSSRSRQ                               int64 = 33
	ProtocolIEIDResultCSIRSRP                              int64 = 34
	ProtocolIEIDResultCSIRSRQ                              int64 = 35
	ProtocolIEIDAngleOfArrivalNR                           int64 = 36
	ProtocolIEIDGeographicalCoordinates                    int64 = 37
	ProtocolIEIDPositioningBroadcastCells                  int64 = 38
	ProtocolIEIDLMFMeasurementID                           int64 = 39
	ProtocolIEIDRANMeasurementID                           int64 = 40
	ProtocolIEIDTRPMeasurementRequestList                  int64 = 41
	ProtocolIEIDTRPMeasurementResponseList                 int64 = 42
	ProtocolIEIDTRPMeasurementReportList                   int64 = 43
	ProtocolIEIDSRSType                                    int64 = 44
	ProtocolIEIDActivationTime                             int64 = 45
	ProtocolIEIDSRSResourceSetID                           int64 = 46
	ProtocolIEIDTRPList                                    int64 = 47
	ProtocolIEIDSRSSpatialRelation                         int64 = 48
	ProtocolIEIDSystemFrameNumber                          int64 = 49
	ProtocolIEIDSlotNumber                                 int64 = 50
	ProtocolIEIDSRSResourceTrigger                         int64 = 51
	ProtocolIEIDTRPMeasurementQuantities                   int64 = 52
	ProtocolIEIDAbortTransmission                          int64 = 53
	ProtocolIEIDSFNInitialisationTime                      int64 = 54
	ProtocolIEIDResultNR                                   int64 = 55
	ProtocolIEIDResultEUTRA                                int64 = 56
	ProtocolIEIDTRPInformationTypeItem                     int64 = 57
	ProtocolIEIDCGINR                                      int64 = 58
	ProtocolIEIDSFNInitialisationTimeNR                    int64 = 59
	ProtocolIEIDCellID                                     int64 = 60
	ProtocolIEIDSrsFrequency                               int64 = 61
	ProtocolIEIDTRPType                                    int64 = 62
	ProtocolIEIDSRSSpatialRelationPerSRSResource           int64 = 63
	ProtocolIEIDMeasurementPeriodicityExtended             int64 = 64
	ProtocolIEIDPRSResourceID                              int64 = 65
	ProtocolIEIDPRSTRPList                                 int64 = 66
	ProtocolIEIDPRSTransmissionTRPList                     int64 = 67
	ProtocolIEIDOnDemandPRS                                int64 = 68
	ProtocolIEIDAoASearchWindow                            int64 = 69
	ProtocolIEIDTRPMeasurementUpdateList                   int64 = 70
	ProtocolIEIDZoA                                        int64 = 71
	ProtocolIEIDResponseTime                               int64 = 72
	ProtocolIEIDUEReportingInformation                     int64 = 73
	ProtocolIEIDMultipleULAoA                              int64 = 74
	ProtocolIEIDULSRSRSRPP                                 int64 = 75
	ProtocolIEIDSRSResourcetype                            int64 = 76
	ProtocolIEIDExtendedAdditionalPathList                 int64 = 77
	ProtocolIEIDARPLocationInfo                            int64 = 78
	ProtocolIEIDARPID                                      int64 = 79
	ProtocolIEIDLoSNLoSInformation                         int64 = 80
	ProtocolIEIDUETxTEGAssociationList                     int64 = 81
	ProtocolIEIDNumberOfTRPRxTEG                           int64 = 82
	ProtocolIEIDNumberOfTRPRxTxTEG                         int64 = 83
	ProtocolIEIDTRPTxTEGAssociation                        int64 = 84
	ProtocolIEIDTRPTEGInformation                          int64 = 85
	ProtocolIEIDTRPRxTEGInformation                        int64 = 86
	ProtocolIEIDTRPPRSInformationList                      int64 = 87
	ProtocolIEIDPRSMeasurementsInfoList                    int64 = 88
	ProtocolIEIDPRSConfigRequestType                       int64 = 89
	ProtocolIEIDUETEGInfoRequest                           int64 = 90
	ProtocolIEIDMeasurementTimeOccasion                    int64 = 91
	ProtocolIEIDMeasurementCharacteristicsRequestIndicator int64 = 92
	ProtocolIEIDTRPBeamAntennaInformation                  int64 = 93
	ProtocolIEIDNRTADV                                     int64 = 94
	ProtocolIEIDMeasurementAmount                          int64 = 95
	ProtocolIEIDPathPower                                  int64 = 96
	ProtocolIEIDPreconfigurationResult                     int64 = 97
	ProtocolIEIDRequestType                                int64 = 98
	ProtocolIEIDUETEGReportingPeriodicity                  int64 = 99
	ProtocolIEIDSRSPortIndex                               int64 = 100
	ProtocolIEIDProcedureCode101NotToBeUsed                int64 = 101
	ProtocolIEIDProcedureCode102NotToBeUsed                int64 = 102
	ProtocolIEIDProcedureCode103NotToBeUsed                int64 = 103
	ProtocolIEIDUETxTimingErrorMargin                      int64 = 104
	ProtocolIEIDMeasurementPeriodicityNRAoA                int64 = 105
	ProtocolIEIDSRSTransmissionStatus                      int64 = 106
	ProtocolIEIDNrofSymbolsExtended                        int64 = 107
	ProtocolIEIDRepetitionFactorExtended                   int64 = 108
	ProtocolIEIDStartRBHopping                             int64 = 109
	ProtocolIEIDStartRBIndex                               int64 = 110
	ProtocolIEIDTransmissionCombn8                         int64 = 111
	ProtocolIEIDSCS480                                     int64 = 119
	ProtocolIEIDSCS960                                     int64 = 120
)
