package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestPagingMarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x00, 0xF1, 0x53})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	fivegTmsi_case1 := &ie.FiveGTMSI{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	// AMFID: "000000"
	amfSetId_case1 := &ie.AMFSetID{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 10,
		},
	}
	amfPointer_case1 := &ie.AMFPointer{
		Value: aper.BitString{
			Bytes:     []byte{0xC8},
			BitLength: 6,
		},
	}

	testCases := []struct {
		name     string
		input    *Paging
		expected []byte
	}{
		{
			name: "Case 1",
			input: &Paging{
				UEPagingIdentity: &ie.UEPagingIdentity{
					Choice: &ie.FiveGSTMSI{
						AMFSetID:   amfSetId_case1,
						AMFPointer: amfPointer_case1,
						FiveGTMSI:  fivegTmsi_case1,
					},
				},
				TAIListForPaging: &ie.TAIListForPaging{
					List: []ie.TAIListForPagingItem{
						{
							TAI: &ie.TAI{
								PLMNIdentity: plmnIdentity,
								TAC:          tac_case1,
							},
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x18, 0x40, 0x19, 0x00, 0x00, 0x02, 0x00,
				0x73, 0x40, 0x07, 0x00, 0x06, 0x40, 0x00, 0x00,
				0x00, 0x01, 0x00, 0x67, 0x40, 0x07, 0x00, 0x00,
				0xf1, 0x53, 0x00, 0x00, 0x01,
			},
		},
		{
			name: "Case 2: from ueranemu basic-k8s pipeline TestPaging",
			input: &Paging{
				UEPagingIdentity: &ie.UEPagingIdentity{
					Choice: &ie.FiveGSTMSI{
						AMFSetID: &ie.AMFSetID{
							Value: aper.BitString{
								Bytes:     []byte{0x00, 0x40},
								BitLength: 10,
							},
						},
						AMFPointer: &ie.AMFPointer{
							Value: aper.BitString{
								Bytes:     []byte{0x04},
								BitLength: 6,
							},
						},
						FiveGTMSI: &ie.FiveGTMSI{
							Value: aper.OctetString{0x06, 0x9a, 0x5f, 0x57},
						},
						IEExtensions: nil,
					},
				},
				PagingDRX: nil,
				TAIListForPaging: &ie.TAIListForPaging{
					List: []ie.TAIListForPagingItem{
						{
							TAI: &ie.TAI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: aper.OctetString{0x02, 0xf8, 0x39},
								},
								TAC: &ie.TAC{
									Value: aper.OctetString{0x00, 0x00, 0x01},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
					},
				},
				PagingPriority:              nil,
				UERadioCapabilityForPaging:  nil,
				PagingOrigin:                nil,
				AssistanceDataForPaging:     nil,
				NBIoTPagingEDRXInfo:         nil,
				NBIoTPagingDRX:              nil,
				EnhancedCoverageRestriction: nil,
				WUSAssistanceInformation:    nil,
				EUTRAPagingeDRXInformation:  nil,
				CEmodeBrestricted:           nil,
				NRPagingeDRXInformation:     nil,
				PagingCause:                 nil,
				PEIPSassistanceInformation:  nil,
			},
			expected: []byte{
				0x00, 0x18, 0x40, 0x19, 0x00, 0x00, 0x02, 0x00, 0x73, 0x40, 0x07, 0x00, 0x08, 0x20, 0x06, 0x9a,
				0x5f, 0x57, 0x00, 0x67, 0x40, 0x07, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01,
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

func TestPagingUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x00, 0xF1, 0x53})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	fivegTmsi_case1 := &ie.FiveGTMSI{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	// AMFID: "000000"
	amfSetId_case1 := &ie.AMFSetID{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 10,
		},
	}
	amfPointer_case1 := &ie.AMFPointer{
		Value: aper.BitString{
			Bytes:     []byte{0xC8},
			BitLength: 6,
		},
	}

	testCases := []struct {
		name     string
		input    []byte
		expected *Paging
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x18, 0x40, 0x19, 0x00, 0x00, 0x02, 0x00,
				0x73, 0x40, 0x07, 0x00, 0x06, 0x40, 0x00, 0x00,
				0x00, 0x01, 0x00, 0x67, 0x40, 0x07, 0x00, 0x00,
				0xf1, 0x53, 0x00, 0x00, 0x01,
			},
			expected: &Paging{
				UEPagingIdentity: &ie.UEPagingIdentity{
					Choice: &ie.FiveGSTMSI{
						AMFSetID:   amfSetId_case1,
						AMFPointer: amfPointer_case1,
						FiveGTMSI:  fivegTmsi_case1,
					},
				},
				TAIListForPaging: &ie.TAIListForPaging{
					List: []ie.TAIListForPagingItem{
						{
							TAI: &ie.TAI{
								PLMNIdentity: plmnIdentity,
								TAC:          tac_case1,
							},
						},
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu basic-k8s pipeline TestPaging",
			input: []byte{
				0x00, 0x18, 0x40, 0x19, 0x00, 0x00, 0x02, 0x00, 0x73, 0x40, 0x07, 0x00, 0x08, 0x20, 0x06, 0x9a,
				0x5f, 0x57, 0x00, 0x67, 0x40, 0x07, 0x00, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01,
			},
			expected: &Paging{
				UEPagingIdentity: &ie.UEPagingIdentity{
					Choice: &ie.FiveGSTMSI{
						AMFSetID: &ie.AMFSetID{
							Value: aper.BitString{
								Bytes:     []byte{0x00, 0x40},
								BitLength: 10,
							},
						},
						AMFPointer: &ie.AMFPointer{
							Value: aper.BitString{
								Bytes:     []byte{0x04},
								BitLength: 6,
							},
						},
						FiveGTMSI: &ie.FiveGTMSI{
							Value: aper.OctetString{0x06, 0x9a, 0x5f, 0x57},
						},
						IEExtensions: nil,
					},
				},
				PagingDRX: nil,
				TAIListForPaging: &ie.TAIListForPaging{
					List: []ie.TAIListForPagingItem{
						{
							TAI: &ie.TAI{
								PLMNIdentity: &ie.PLMNIdentity{
									Value: aper.OctetString{0x02, 0xf8, 0x39},
								},
								TAC: &ie.TAC{
									Value: aper.OctetString{0x00, 0x00, 0x01},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
					},
				},
				PagingPriority:              nil,
				UERadioCapabilityForPaging:  nil,
				PagingOrigin:                nil,
				AssistanceDataForPaging:     nil,
				NBIoTPagingEDRXInfo:         nil,
				NBIoTPagingDRX:              nil,
				EnhancedCoverageRestriction: nil,
				WUSAssistanceInformation:    nil,
				EUTRAPagingeDRXInformation:  nil,
				CEmodeBrestricted:           nil,
				NRPagingeDRXInformation:     nil,
				PagingCause:                 nil,
				PEIPSassistanceInformation:  nil,
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
