package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestUEContextReleaseRequestMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *UEContextReleaseRequest
		expected []byte
	}{
		{
			name: "Case 1",
			input: &UEContextReleaseRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 2348810240,
				},
				PDUSessionResourceListCxtRelReq: &ie.PDUSessionResourceListCxtRelReq{
					List: []ie.PDUSessionResourceItemCxtRelReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentUserInactivity,
					},
				},
			},
			expected: []byte{
				0x00, 0x2a, 0x40, 0x1f, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00,
				0x02, 0x00, 0x02, 0x00, 0x55, 0x00, 0x05, 0xc0, 0x8c, 0x00,
				0x00, 0x00, 0x00, 0x85, 0x00, 0x03, 0x00, 0x00, 0x05, 0x00,
				0x0f, 0x40, 0x02, 0x05, 0x00,
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestServiceRequest",
			input: &UEContextReleaseRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 885695317650,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceListCxtRelReq: &ie.PDUSessionResourceListCxtRelReq{
					List: []ie.PDUSessionResourceItemCxtRelReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: aper.Enumerated(1),
					},
				},
			},
			expected: []byte{
				0x00, 0x2a, 0x40, 0x20, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xce, 0x37, 0x8e, 0x06,
				0x92, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x85, 0x00, 0x03, 0x00, 0x00, 0x0a, 0x00, 0x0f,
				0x40, 0x02, 0x00, 0x40,
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

func TestUEContextReleaseRequestUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *UEContextReleaseRequest
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x2a, 0x40, 0x1f, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00,
				0x02, 0x00, 0x02, 0x00, 0x55, 0x00, 0x05, 0xc0, 0x8c, 0x00,
				0x00, 0x00, 0x00, 0x85, 0x00, 0x03, 0x00, 0x00, 0x05, 0x00,
				0x0f, 0x40, 0x02, 0x05, 0x00,
			},
			expected: &UEContextReleaseRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 2348810240,
				},
				PDUSessionResourceListCxtRelReq: &ie.PDUSessionResourceListCxtRelReq{
					List: []ie.PDUSessionResourceItemCxtRelReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: ie.CauseRadioNetworkPresentUserInactivity,
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestServiceRequest",
			input: []byte{
				0x00, 0x2a, 0x40, 0x20, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x00, 0x06, 0x80, 0xce, 0x37, 0x8e, 0x06,
				0x92, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x85, 0x00, 0x03, 0x00, 0x00, 0x0a, 0x00, 0x0f,
				0x40, 0x02, 0x00, 0x40,
			},
			expected: &UEContextReleaseRequest{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 885695317650,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				PDUSessionResourceListCxtRelReq: &ie.PDUSessionResourceListCxtRelReq{
					List: []ie.PDUSessionResourceItemCxtRelReq{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: aper.Enumerated(1),
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
