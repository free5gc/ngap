package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestUEContextReleaseCommandMarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *UEContextReleaseCommand
		expected []byte
	}{
		{
			name: "Case 1",
			input: &UEContextReleaseCommand{
				UENGAPIDs: &ie.UENGAPIDs{
					Choice: &ie.UENGAPIDPair{
						AMFUENGAPID: &ie.AMFUENGAPID{
							Value: 1,
						},
						RANUENGAPID: &ie.RANUENGAPID{
							Value: 3405774848,
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseNas{
						Value: ie.CauseNasPresentNormalRelease,
					},
				},
			},
			expected: []byte{
				0x00, 0x29, 0x00, 0x13, 0x00, 0x00, 0x02, 0x00, 0x72,
				0x00, 0x07, 0x00, 0x01, 0xc0, 0xcb, 0x00, 0x00, 0x00,
				0x00, 0x0f, 0x40, 0x01, 0x40,
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestServiceRequest",
			input: &UEContextReleaseCommand{
				UENGAPIDs: &ie.UENGAPIDs{
					Choice: &ie.UENGAPIDPair{
						AMFUENGAPID: &ie.AMFUENGAPID{
							Value: 885695317650,
						},
						RANUENGAPID: &ie.RANUENGAPID{
							Value: 1,
						},
						IEExtensions: nil, // Assuming nil since it's not provided
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseRadioNetwork{
						Value: aper.Enumerated(1),
					},
				},
			},
			expected: []byte{
				0x00, 0x29, 0x00, 0x15, 0x00, 0x00, 0x02, 0x00, 0x72, 0x00, 0x08, 0x08, 0xce, 0x37, 0x8e, 0x06,
				0x92, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x00, 0x40,
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

func TestUEContextReleaseCommandUnmarshalBinary(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []byte
		expected *UEContextReleaseCommand
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x29, 0x00, 0x13, 0x00, 0x00, 0x02, 0x00, 0x72,
				0x00, 0x07, 0x00, 0x01, 0xc0, 0xcb, 0x00, 0x00, 0x00,
				0x00, 0x0f, 0x40, 0x01, 0x40,
			},
			expected: &UEContextReleaseCommand{
				UENGAPIDs: &ie.UENGAPIDs{
					Choice: &ie.UENGAPIDPair{
						AMFUENGAPID: &ie.AMFUENGAPID{
							Value: 1,
						},
						RANUENGAPID: &ie.RANUENGAPID{
							Value: 3405774848,
						},
					},
				},
				Cause: &ie.Cause{
					Choice: &ie.CauseNas{
						Value: ie.CauseNasPresentNormalRelease,
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestServiceRequest",
			input: []byte{
				0x00, 0x29, 0x00, 0x15, 0x00, 0x00, 0x02, 0x00, 0x72, 0x00, 0x08, 0x08, 0xce, 0x37, 0x8e, 0x06,
				0x92, 0x00, 0x01, 0x00, 0x0f, 0x40, 0x02, 0x00, 0x40,
			},
			expected: &UEContextReleaseCommand{
				UENGAPIDs: &ie.UENGAPIDs{
					Choice: &ie.UENGAPIDPair{
						AMFUENGAPID: &ie.AMFUENGAPID{
							Value: 885695317650,
						},
						RANUENGAPID: &ie.RANUENGAPID{
							Value: 1,
						},
						IEExtensions: nil, // Assuming nil since it's not provided
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
