package message

import "github.com/free5gc/ngap/aper"

// dummy function to avoid unused error
func foo(args ...interface{}) {}

const (
	MessageTypeInitiatingMessage   int64 = 0
	MessageTypeSuccessfulOutcome   int64 = 1
	MessageTypeUnsuccessfulOutcome int64 = 2
)

const (
	CriticalityReject aper.Enumerated = 0
	CriticalityIgnore aper.Enumerated = 1
	CriticalityNotify aper.Enumerated = 2
)

const (
	ProcedureCodeErrorIndication                  int64 = 0
	ProcedureCodePrivateMessage                   int64 = 1
	ProcedureCodeECIDMeasurementInitiation        int64 = 2
	ProcedureCodeECIDMeasurementFailureIndication int64 = 3
	ProcedureCodeECIDMeasurementReport            int64 = 4
	ProcedureCodeECIDMeasurementTermination       int64 = 5
	ProcedureCodeOTDOAInformationExchange         int64 = 6
	ProcedureCodeAssistanceInformationControl     int64 = 7
	ProcedureCodeAssistanceInformationFeedback    int64 = 8
	ProcedureCodePositioningInformationExchange   int64 = 9
	ProcedureCodePositioningInformationUpdate     int64 = 10
	ProcedureCodeMeasurement                      int64 = 11
	ProcedureCodeMeasurementReport                int64 = 12
	ProcedureCodeMeasurementUpdate                int64 = 13
	ProcedureCodeMeasurementAbort                 int64 = 14
	ProcedureCodeMeasurementFailureIndication     int64 = 15
	ProcedureCodeTRPInformationExchange           int64 = 16
	ProcedureCodePositioningActivation            int64 = 17
	ProcedureCodePositioningDeactivation          int64 = 18
	ProcedureCodePRSConfigurationExchange         int64 = 19
	ProcedureCodeMeasurementPreconfiguration      int64 = 20
	ProcedureCodeMeasurementActivation            int64 = 21
)
