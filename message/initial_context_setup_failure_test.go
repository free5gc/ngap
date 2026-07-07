package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestInitialContextSetupFailureMarshalBinary(t *testing.T) {
	t.Parallel()

	pDUSessionResourceSetupUnsuccessfulTransfer := ie.PDUSessionResourceSetupUnsuccessfulTransfer{
		Cause: &ie.Cause{
			Choice: &ie.CauseRadioNetwork{
				Value: ie.CauseRadioNetworkPresentFailureInRadioInterfaceProcedure,
			},
		},
	}
	pDUSessionResourceSetupUnsuccessfulTransferBytes, err := ie.MarshalBinary(&pDUSessionResourceSetupUnsuccessfulTransfer)
	require.NoError(t, err)
	pDUSessionResourceSetupUnsuccessfulTransferOS := aper.OctetString(pDUSessionResourceSetupUnsuccessfulTransferBytes)

	testCases := []struct {
		name     string
		input    *InitialContextSetupFailure
		expected []byte
	}{
		{
			name: "Case 1",
			input: &InitialContextSetupFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 4,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 142,
				},
				PDUSessionResourceFailedToSetupListCxtFail: &ie.PDUSessionResourceFailedToSetupListCxtFail{
					List: []ie.PDUSessionResourceFailedToSetupItemCxtFail{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 1,
							},
							PDUSessionResourceSetupUnsuccessfulTransfer: &pDUSessionResourceSetupUnsuccessfulTransferOS,
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentFailureInRadioInterfaceProcedure,
					},
				},
			},
			expected: []byte{
				0x40, 0x0e, 0x00, 0x1f, 0x00, 0x00, 0x04, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x04, 0x00, 0x55, 0x40, 0x02, 0x00,
				0x8e, 0x00, 0x84, 0x40, 0x06, 0x00, 0x00, 0x01, 0x02,
				0x00, 0xc0, 0x00, 0x0f, 0x40, 0x02, 0x06, 0x00,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := tc.input.MarshalBinary()
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, b)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestInitialContextSetupFailureUnmarshalBinary(t *testing.T) {
	t.Parallel()

	pDUSessionResourceSetupUnsuccessfulTransfer := ie.PDUSessionResourceSetupUnsuccessfulTransfer{
		Cause: &ie.Cause{
			Choice: &ie.CauseRadioNetwork{
				Value: ie.CauseRadioNetworkPresentFailureInRadioInterfaceProcedure,
			},
		},
	}
	pDUSessionResourceSetupUnsuccessfulTransferBytes, err := ie.MarshalBinary(&pDUSessionResourceSetupUnsuccessfulTransfer)
	require.NoError(t, err)
	pDUSessionResourceSetupUnsuccessfulTransferOS := aper.OctetString(pDUSessionResourceSetupUnsuccessfulTransferBytes)

	testCases := []struct {
		name     string
		input    []byte
		expected *InitialContextSetupFailure
	}{
		{
			name: "Case 1",
			input: []byte{
				0x40, 0x0e, 0x00, 0x1f, 0x00, 0x00, 0x04, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x04, 0x00, 0x55, 0x40, 0x02, 0x00,
				0x8e, 0x00, 0x84, 0x40, 0x06, 0x00, 0x00, 0x01, 0x02,
				0x00, 0xc0, 0x00, 0x0f, 0x40, 0x02, 0x06, 0x00,
			},
			expected: &InitialContextSetupFailure{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 4,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 142,
				},
				PDUSessionResourceFailedToSetupListCxtFail: &ie.PDUSessionResourceFailedToSetupListCxtFail{
					List: []ie.PDUSessionResourceFailedToSetupItemCxtFail{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 1,
							},
							PDUSessionResourceSetupUnsuccessfulTransfer: &pDUSessionResourceSetupUnsuccessfulTransferOS,
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentFailureInRadioInterfaceProcedure,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			msg, err := Parse(tc.input)
			if tc.expected != nil {
				require.NoError(t, err)
				require.Equal(t, tc.expected, msg)
			} else {
				require.Error(t, err)
			}
		})
	}
}
