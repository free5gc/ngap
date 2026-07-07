package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestPathSwitchRequestMarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	gtpTeid := &ie.GTPTEID{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	nrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x80, 0x00, 0x00, 0x00, 0x10},
			BitLength: 36,
		},
	}
	nrEncryptAlgo := &ie.NRencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	nrIntegrityAlgo := &ie.NRintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x40, 0x00},
			BitLength: 16,
		},
	}
	eutraEncryptAlgo := &ie.EUTRAencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	eutraIntegrityAlgo := &ie.EUTRAintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}

	// pathSwitchRequestTransfer
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x02},
			BitLength: 32,
		},
	}
	pathSwitchRequestTransfer := ie.PathSwitchRequestTransfer{
		DLNGUUPTNLInformation: &ie.UPTransportLayerInformation{
			Choice: &ie.GTPTunnel{
				TransportLayerAddress: ipv4Addr,
				GTPTEID:               gtpTeid,
			},
		},
		QosFlowAcceptedList: &ie.QosFlowAcceptedList{
			List: []ie.QosFlowAcceptedItem{
				{
					QosFlowIdentifier: &ie.QosFlowIdentifier{
						Value: 9,
					},
				},
			},
		},
	}
	pathSwitchRequestTransferBytes, err := ie.MarshalBinary(
		&pathSwitchRequestTransfer)
	require.NoError(t, err)
	pathSwitchRequestTransferOS := aper.OctetString(
		pathSwitchRequestTransferBytes)

	testCases := []struct {
		name     string
		input    *PathSwitchRequest
		expected []byte
	}{
		{
			name: "Case 1",
			input: &PathSwitchRequest{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 2818572289,
				},
				SourceAMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case1,
							NRCellIdentity: nrCellIdentity,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms:             nrEncryptAlgo,
					NRintegrityProtectionAlgorithms:    nrIntegrityAlgo,
					EUTRAencryptionAlgorithms:          eutraEncryptAlgo,
					EUTRAintegrityProtectionAlgorithms: eutraIntegrityAlgo,
				},
				PDUSessionResourceToBeSwitchedDLList: &ie.PDUSessionResourceToBeSwitchedDLList{
					List: []ie.PDUSessionResourceToBeSwitchedDLItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							PathSwitchRequestTransfer: &pathSwitchRequestTransferOS,
						},
					},
				},
			},
			expected: []byte{
				0x00, 0x19, 0x00, 0x46, 0x00, 0x00, 0x05, 0x00,
				0x55, 0x00, 0x05, 0xc0, 0xa8, 0x00, 0x00, 0x01,
				0x00, 0x64, 0x00, 0x02, 0x00, 0x02, 0x00, 0x79,
				0x40, 0x0f, 0x40, 0x13, 0x30, 0x01, 0x80, 0x00,
				0x00, 0x00, 0x10, 0x13, 0x30, 0x01, 0x00, 0x00,
				0x01, 0x00, 0x77, 0x40, 0x09, 0x00, 0x00, 0x04,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4c,
				0x00, 0x10, 0x00, 0x00, 0x05, 0x0c, 0x00, 0x1f,
				0xac, 0x10, 0x03, 0x02, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x12,
			},
		},
		{
			name: "Case 2 - from ueranemu k8s-basic pipeline TestXnHandover",
			input: &PathSwitchRequest{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				SourceAMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1014734679650,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x01, 0x03, 0x00, 0x10},
									BitLength: 36,
								},
							},
							IEExtensions: nil,
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x01},
							},
							IEExtensions: nil,
						},
						TimeStamp:    nil,
						IEExtensions: nil,
					},
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms: &ie.NRencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					NRintegrityProtectionAlgorithms: &ie.NRintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					EUTRAencryptionAlgorithms: &ie.EUTRAencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					EUTRAintegrityProtectionAlgorithms: &ie.EUTRAintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					IEExtensions: nil,
				},
				PDUSessionResourceToBeSwitchedDLList: &ie.PDUSessionResourceToBeSwitchedDLList{
					List: []ie.PDUSessionResourceToBeSwitchedDLItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested in TestPathSwitchRequestTransfer
							PathSwitchRequestTransfer: &aper.OctetString{0x00, 0x1f, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02},
							IEExtensions:              nil,
						},
					},
				},
				PDUSessionResourceFailedToSetupListPSReq: nil,
				RRCResumeCause:                           nil,
				RedCapIndication:                         nil,
			},
			expected: []byte{
				0x00, 0x19, 0x00, 0x47, 0x00, 0x00, 0x05, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x64, 0x00,
				0x06, 0x80, 0xec, 0x42, 0xe6, 0x6e, 0x62, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00,
				0x01, 0x03, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01, 0x00, 0x77, 0x40, 0x09, 0x1f, 0xff,
				0xef, 0xff, 0xf7, 0xff, 0xfb, 0xff, 0xfc, 0x00, 0x4c, 0x00, 0x10, 0x00, 0x00, 0x0a, 0x0c, 0x00,
				0x1f, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02,
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

func TestPathSwitchRequestUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	plmnIdentity_case1 := &ie.PLMNIdentity{Value: aper.OctetString([]byte{0x13, 0x30, 0x01})}
	tac_case1 := &ie.TAC{Value: aper.OctetString([]byte{0x00, 0x00, 0x01})}
	gtpTeid := &ie.GTPTEID{Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x01})}
	nrCellIdentity := &ie.NRCellIdentity{
		Value: aper.BitString{
			Bytes:     []byte{0x80, 0x00, 0x00, 0x00, 0x10},
			BitLength: 36,
		},
	}
	nrEncryptAlgo := &ie.NRencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	nrIntegrityAlgo := &ie.NRintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x40, 0x00},
			BitLength: 16,
		},
	}
	eutraEncryptAlgo := &ie.EUTRAencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	eutraIntegrityAlgo := &ie.EUTRAintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}

	// pathSwitchRequestTransfer
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x02},
			BitLength: 32,
		},
	}
	pathSwitchRequestTransfer := &ie.PathSwitchRequestTransfer{
		DLNGUUPTNLInformation: &ie.UPTransportLayerInformation{
			Choice: &ie.GTPTunnel{
				TransportLayerAddress: ipv4Addr,
				GTPTEID:               gtpTeid,
			},
		},
		QosFlowAcceptedList: &ie.QosFlowAcceptedList{
			List: []ie.QosFlowAcceptedItem{
				{
					QosFlowIdentifier: &ie.QosFlowIdentifier{
						Value: 9,
					},
				},
			},
		},
	}
	pathSwitchRequestTransferBytes, err := ie.MarshalBinary(pathSwitchRequestTransfer)
	require.NoError(t, err)
	pathSwitchRequestTransferOS := aper.OctetString(pathSwitchRequestTransferBytes)

	testCases := []struct {
		name     string
		input    []byte
		expected *PathSwitchRequest
	}{
		{
			name: "Case 1",
			input: []byte{
				0x00, 0x19, 0x00, 0x46, 0x00, 0x00, 0x05, 0x00,
				0x55, 0x00, 0x05, 0xc0, 0xa8, 0x00, 0x00, 0x01,
				0x00, 0x64, 0x00, 0x02, 0x00, 0x02, 0x00, 0x79,
				0x40, 0x0f, 0x40, 0x13, 0x30, 0x01, 0x80, 0x00,
				0x00, 0x00, 0x10, 0x13, 0x30, 0x01, 0x00, 0x00,
				0x01, 0x00, 0x77, 0x40, 0x09, 0x00, 0x00, 0x04,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4c,
				0x00, 0x10, 0x00, 0x00, 0x05, 0x0c, 0x00, 0x1f,
				0xac, 0x10, 0x03, 0x02, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x12,
			},
			expected: &PathSwitchRequest{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 2818572289,
				},
				SourceAMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity:   plmnIdentity_case1,
							NRCellIdentity: nrCellIdentity,
						},
						TAI: &ie.TAI{
							PLMNIdentity: plmnIdentity_case1,
							TAC:          tac_case1,
						},
					},
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms:             nrEncryptAlgo,
					NRintegrityProtectionAlgorithms:    nrIntegrityAlgo,
					EUTRAencryptionAlgorithms:          eutraEncryptAlgo,
					EUTRAintegrityProtectionAlgorithms: eutraIntegrityAlgo,
				},
				PDUSessionResourceToBeSwitchedDLList: &ie.PDUSessionResourceToBeSwitchedDLList{
					List: []ie.PDUSessionResourceToBeSwitchedDLItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							PathSwitchRequestTransfer: &pathSwitchRequestTransferOS,
						},
					},
				},
			},
		},
		{
			name: "Case 2 - from ueranemu k8s-basic pipeline TestXnHandover",
			input: []byte{
				0x00, 0x19, 0x00, 0x47, 0x00, 0x00, 0x05, 0x00, 0x55, 0x00, 0x02, 0x00, 0x01, 0x00, 0x64, 0x00,
				0x06, 0x80, 0xec, 0x42, 0xe6, 0x6e, 0x62, 0x00, 0x79, 0x40, 0x0f, 0x40, 0x02, 0xf8, 0x39, 0x00,
				0x01, 0x03, 0x00, 0x10, 0x02, 0xf8, 0x39, 0x00, 0x00, 0x01, 0x00, 0x77, 0x40, 0x09, 0x1f, 0xff,
				0xef, 0xff, 0xf7, 0xff, 0xfb, 0xff, 0xfc, 0x00, 0x4c, 0x00, 0x10, 0x00, 0x00, 0x0a, 0x0c, 0x00,
				0x1f, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02,
			},
			expected: &PathSwitchRequest{
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				SourceAMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1014734679650,
				},
				UserLocationInformation: &ie.UserLocationInformation{
					Choice: &ie.UserLocationInformationNR{
						NRCGI: &ie.NRCGI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							NRCellIdentity: &ie.NRCellIdentity{
								Value: aper.BitString{
									Bytes:     []byte{0x00, 0x01, 0x03, 0x00, 0x10},
									BitLength: 36,
								},
							},
							IEExtensions: nil,
						},
						TAI: &ie.TAI{
							PLMNIdentity: &ie.PLMNIdentity{
								Value: aper.OctetString{0x02, 0xf8, 0x39},
							},
							TAC: &ie.TAC{
								Value: aper.OctetString{0x00, 0x00, 0x01},
							},
							IEExtensions: nil,
						},
						TimeStamp:    nil,
						IEExtensions: nil,
					},
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms: &ie.NRencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					NRintegrityProtectionAlgorithms: &ie.NRintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					EUTRAencryptionAlgorithms: &ie.EUTRAencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					EUTRAintegrityProtectionAlgorithms: &ie.EUTRAintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xff, 0xff},
							BitLength: 16,
						},
					},
					IEExtensions: nil,
				},
				PDUSessionResourceToBeSwitchedDLList: &ie.PDUSessionResourceToBeSwitchedDLList{
					List: []ie.PDUSessionResourceToBeSwitchedDLItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							PathSwitchRequestTransfer: &aper.OctetString{0x00, 0x1f, 0xac, 0x10, 0x1f, 0x33, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02},
							IEExtensions:              nil,
						},
					},
				},
				PDUSessionResourceFailedToSetupListPSReq: nil,
				RRCResumeCause:                           nil,
				RedCapIndication:                         nil,
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
