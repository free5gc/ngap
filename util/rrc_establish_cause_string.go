package util

import (
	"fmt"

	"github.com/free5gc/ngap/ie"
)

func RrcEstablishmentCauseToString(cause *ie.RRCEstablishmentCause) string {
	if cause == nil {
		return ""
	}
	switch cause.Value {
	case ie.RRCEstablishmentCausePresentEmergency:
		return "Emergency"
	case ie.RRCEstablishmentCausePresentHighPriorityAccess:
		return "HighPriorityAccess"
	case ie.RRCEstablishmentCausePresentMtAccess:
		return "MtAccess"
	case ie.RRCEstablishmentCausePresentMoSignalling:
		return "MoSignalling"
	case ie.RRCEstablishmentCausePresentMoData:
		return "MoData"
	case ie.RRCEstablishmentCausePresentMoVoiceCall:
		return "MoVoiceCall"
	case ie.RRCEstablishmentCausePresentMoVideoCall:
		return "MoVideoCall"
	case ie.RRCEstablishmentCausePresentMoSMS:
		return "MoSMS"
	case ie.RRCEstablishmentCausePresentMpsPriorityAccess:
		return "MpsPriorityAccess"
	case ie.RRCEstablishmentCausePresentMcsPriorityAccess:
		return "McsPriorityAccess"
	case ie.RRCEstablishmentCausePresentNotAvailable:
		return "NotAvailable"
	case ie.RRCEstablishmentCausePresentMoExceptionData:
		return "MoExceptionData"
	}
	return fmt.Sprintf("Unknown RRC Establishment Cause[%v]", cause.Value)
}
