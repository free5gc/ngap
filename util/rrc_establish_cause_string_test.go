package util

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/ie"
)

func TestRrcEstablishmentCauseToString(t *testing.T) {
	tcs := []struct {
		name   string
		cause  *ie.RRCEstablishmentCause
		expect string
	}{
		{
			name:   "cause is nil",
			cause:  nil,
			expect: "",
		},
		{
			name: "cause: Emergency",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentEmergency,
			},
			expect: "Emergency",
		},
		{
			name: "cause: HighPriorityAccess",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentHighPriorityAccess,
			},
			expect: "HighPriorityAccess",
		},
		{
			name: "cause: MtAccess",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMtAccess,
			},
			expect: "MtAccess",
		},
		{
			name: "cause: MoSignalling",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMoSignalling,
			},
			expect: "MoSignalling",
		},
		{
			name: "cause: MoData",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMoData,
			},
			expect: "MoData",
		},
		{
			name: "cause: MoVoiceCall",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMoVoiceCall,
			},
			expect: "MoVoiceCall",
		},
		{
			name: "cause: MoVideoCall",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMoVideoCall,
			},
			expect: "MoVideoCall",
		},
		{
			name: "cause: MoSMS",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMoSMS,
			},
			expect: "MoSMS",
		},
		{
			name: "cause: MpsPriorityAccess",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMpsPriorityAccess,
			},
			expect: "MpsPriorityAccess",
		},
		{
			name: "cause: McsPriorityAccess",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMcsPriorityAccess,
			},
			expect: "McsPriorityAccess",
		},
		{
			name: "cause: NotAvailable",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentNotAvailable,
			},
			expect: "NotAvailable",
		},
		{
			name: "cause: MoExceptionData",
			cause: &ie.RRCEstablishmentCause{
				Value: ie.RRCEstablishmentCausePresentMoExceptionData,
			},
			expect: "MoExceptionData",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.name, func(t *testing.T) {
			s := RrcEstablishmentCauseToString(tc.cause)
			require.Equal(t, tc.expect, s)
		})
	}
}
