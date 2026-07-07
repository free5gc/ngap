package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/ie"
)

func TestPathSwitchRequestAcknowledgeMarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	sst_case1 := &ie.SST{Value: aper.OctetString([]byte{0x01})}
	sd1_case1 := &ie.SD{Value: aper.OctetString([]byte{0x01, 0x02, 0x03})}
	sd2_case1 := &ie.SD{Value: aper.OctetString([]byte{0x11, 0x22, 0x33})}
	nrEncryptAlgo := &ie.NRencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	nrIntegrityAlgo := &ie.NRintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
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
	securityKey := &ie.SecurityKey{
		Value: aper.BitString{
			Bytes: []byte{
				0x94, 0xdd, 0x95, 0x9f, 0x2a, 0x33, 0xbf, 0x3b,
				0x66, 0x29, 0x5e, 0x64, 0x10, 0xd7, 0x55, 0x30,
				0x80, 0x99, 0x7e, 0x12, 0xbe, 0x61, 0x90, 0x9d,
				0xe4, 0x0e, 0x0a, 0xe5, 0xb4, 0xf4, 0x25, 0x25,
			},
			BitLength: 256,
		},
	}

	// pathSwitchRequestAcknowledgeTransfer
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x64},
			BitLength: 32,
		},
	}
	gtpteid := &ie.GTPTEID{
		Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x03}),
	}
	pathSwitchRequestAcknowledgeTransfer := ie.PathSwitchRequestAcknowledgeTransfer{
		ULNGUUPTNLInformation: &ie.UPTransportLayerInformation{
			Choice: &ie.GTPTunnel{
				TransportLayerAddress: ipv4Addr,
				GTPTEID:               gtpteid,
			},
		},
		SecurityIndication: &ie.SecurityIndication{
			IntegrityProtectionIndication: &ie.IntegrityProtectionIndication{
				Value: ie.IntegrityProtectionIndicationPresentNotNeeded,
			},
			ConfidentialityProtectionIndication: &ie.ConfidentialityProtectionIndication{
				Value: ie.ConfidentialityProtectionIndicationPresentNotNeeded,
			},
		},
	}
	pathSwitchRequestAcknowledgeTransferBytes, err := ie.MarshalBinary(
		&pathSwitchRequestAcknowledgeTransfer)
	require.NoError(t, err)
	pathSwitchRequestAcknowledgeTransferOS := aper.OctetString(
		pathSwitchRequestAcknowledgeTransferBytes)

	testCases := []struct {
		name     string
		input    *PathSwitchRequestAcknowledge
		expected []byte
	}{
		{
			name: "Case 1",
			input: &PathSwitchRequestAcknowledge{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 2818572289,
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms:             nrEncryptAlgo,
					NRintegrityProtectionAlgorithms:    nrIntegrityAlgo,
					EUTRAencryptionAlgorithms:          eutraEncryptAlgo,
					EUTRAintegrityProtectionAlgorithms: eutraIntegrityAlgo,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: securityKey,
				},
				PDUSessionResourceSwitchedList: &ie.PDUSessionResourceSwitchedList{
					List: []ie.PDUSessionResourceSwitchedItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							PathSwitchRequestAcknowledgeTransfer: &pathSwitchRequestAcknowledgeTransferOS,
						},
					},
				},
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd1_case1,
							},
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd2_case1,
							},
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x19, 0x00, 0x66, 0x00, 0x00, 0x06, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x02, 0x00, 0x55, 0x40, 0x05, 0xc0,
				0xa8, 0x00, 0x00, 0x01, 0x00, 0x77, 0x00, 0x09, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x5d, 0x00, 0x21, 0x10, 0x94, 0xdd, 0x95, 0x9f, 0x2a,
				0x33, 0xbf, 0x3b, 0x66, 0x29, 0x5e, 0x64, 0x10, 0xd7,
				0x55, 0x30, 0x80, 0x99, 0x7e, 0x12, 0xbe, 0x61, 0x90,
				0x9d, 0xe4, 0x0e, 0x0a, 0xe5, 0xb4, 0xf4, 0x25, 0x25,
				0x00, 0x4d, 0x40, 0x10, 0x00, 0x00, 0x05, 0x0c, 0x60,
				0x1f, 0xac, 0x10, 0x03, 0x64, 0x00, 0x00, 0x00, 0x03,
				0x09, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x22, 0x01, 0x01,
				0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestXnHandover",
			input: &PathSwitchRequestAcknowledge{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1014734679650,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms: &ie.NRencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xc0, 0x00},
							BitLength: 16,
						},
					},
					NRintegrityProtectionAlgorithms: &ie.NRintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xc0, 0x00},
							BitLength: 16,
						},
					},
					EUTRAencryptionAlgorithms: &ie.EUTRAencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0x00, 0x00},
							BitLength: 16,
						},
					},
					EUTRAintegrityProtectionAlgorithms: &ie.EUTRAintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0x00, 0x00},
							BitLength: 16,
						},
					},
					IEExtensions: nil,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: &ie.SecurityKey{
						Value: aper.BitString{
							Bytes: []byte{
								0x95, 0x74, 0x7e, 0x31, 0x52, 0xdf, 0x0d, 0xe2,
								0xe7, 0xee, 0x36, 0x6d, 0xf4, 0x6e, 0xb4, 0x83,
								0x39, 0xd1, 0xe8, 0x52, 0x46, 0xca, 0x58, 0x5b,
								0x43, 0xbb, 0x67, 0xcb, 0x41, 0xd6, 0x80, 0x0a,
							},
							BitLength: 256,
						},
					},
					IEExtensions: nil,
				},
				NewSecurityContextInd: nil,
				PDUSessionResourceSwitchedList: &ie.PDUSessionResourceSwitchedList{
					List: []ie.PDUSessionResourceSwitchedItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPathSwitchRequestAcknowledgeTransferMarshalBinary
							PathSwitchRequestAcknowledgeTransfer: &aper.OctetString{
								0x40, 0x1f, 0xac, 0x10, 0x1f, 0x3d, 0x92, 0xf7, 0x52, 0x1f,
							},
							IEExtensions: nil,
						},
					},
				},
				PDUSessionResourceReleasedListPSAck: nil,
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x01, 0x02, 0x03},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x11, 0x22, 0x33},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
					},
				},
			},
			expected: []byte{
				0x20, 0x19, 0x00, 0x65, 0x00, 0x00, 0x06, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xec, 0x42, 0xe6, 0x6e,
				0x62, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x77, 0x00, 0x09, 0x18, 0x00, 0x0c, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x21, 0x10, 0x95, 0x74, 0x7e, 0x31, 0x52, 0xdf, 0x0d,
				0xe2, 0xe7, 0xee, 0x36, 0x6d, 0xf4, 0x6e, 0xb4, 0x83, 0x39, 0xd1, 0xe8, 0x52, 0x46, 0xca, 0x58,
				0x5b, 0x43, 0xbb, 0x67, 0xcb, 0x41, 0xd6, 0x80, 0x0a, 0x00, 0x4d, 0x40, 0x0e, 0x00, 0x00, 0x0a,
				0x0a, 0x40, 0x1f, 0xac, 0x10, 0x1f, 0x3d, 0x92, 0xf7, 0x52, 0x1f, 0x00, 0x00, 0x00, 0x0a, 0x22,
				0x01, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
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

func TestPathSwitchRequestAcknowledgeUnmarshalBinary(t *testing.T) {
	t.Parallel()

	// Case 1
	sst_case1 := &ie.SST{Value: aper.OctetString([]byte{0x01})}
	sd1_case1 := &ie.SD{Value: aper.OctetString([]byte{0x01, 0x02, 0x03})}
	sd2_case1 := &ie.SD{Value: aper.OctetString([]byte{0x11, 0x22, 0x33})}
	nrEncryptAlgo := &ie.NRencryptionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
			BitLength: 16,
		},
	}
	nrIntegrityAlgo := &ie.NRintegrityProtectionAlgorithms{
		Value: aper.BitString{
			Bytes:     []byte{0x00, 0x00},
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
	securityKey := &ie.SecurityKey{
		Value: aper.BitString{
			Bytes: []byte{
				0x94, 0xdd, 0x95, 0x9f, 0x2a, 0x33, 0xbf, 0x3b,
				0x66, 0x29, 0x5e, 0x64, 0x10, 0xd7, 0x55, 0x30,
				0x80, 0x99, 0x7e, 0x12, 0xbe, 0x61, 0x90, 0x9d,
				0xe4, 0x0e, 0x0a, 0xe5, 0xb4, 0xf4, 0x25, 0x25,
			},
			BitLength: 256,
		},
	}

	// pathSwitchRequestAcknowledgeTransfer
	ipv4Addr := &ie.TransportLayerAddress{
		Value: aper.BitString{
			Bytes:     []byte{0xac, 0x10, 0x03, 0x64},
			BitLength: 32,
		},
	}
	gtpteid := &ie.GTPTEID{
		Value: aper.OctetString([]byte{0x00, 0x00, 0x00, 0x03}),
	}
	pathSwitchRequestAcknowledgeTransfer := ie.PathSwitchRequestAcknowledgeTransfer{
		ULNGUUPTNLInformation: &ie.UPTransportLayerInformation{
			Choice: &ie.GTPTunnel{
				TransportLayerAddress: ipv4Addr,
				GTPTEID:               gtpteid,
			},
		},
		SecurityIndication: &ie.SecurityIndication{
			IntegrityProtectionIndication: &ie.IntegrityProtectionIndication{
				Value: ie.IntegrityProtectionIndicationPresentNotNeeded,
			},
			ConfidentialityProtectionIndication: &ie.ConfidentialityProtectionIndication{
				Value: ie.ConfidentialityProtectionIndicationPresentNotNeeded,
			},
		},
	}
	pathSwitchRequestAcknowledgeTransferBytes, err := ie.MarshalBinary(
		&pathSwitchRequestAcknowledgeTransfer)
	require.NoError(t, err)
	pathSwitchRequestAcknowledgeTransferOS := aper.OctetString(
		pathSwitchRequestAcknowledgeTransferBytes)

	testCases := []struct {
		name     string
		input    []byte
		expected *PathSwitchRequestAcknowledge
	}{
		{
			name: "Case 1",
			input: []byte{
				0x20, 0x19, 0x00, 0x66, 0x00, 0x00, 0x06, 0x00, 0x0a,
				0x40, 0x02, 0x00, 0x02, 0x00, 0x55, 0x40, 0x05, 0xc0,
				0xa8, 0x00, 0x00, 0x01, 0x00, 0x77, 0x00, 0x09, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x5d, 0x00, 0x21, 0x10, 0x94, 0xdd, 0x95, 0x9f, 0x2a,
				0x33, 0xbf, 0x3b, 0x66, 0x29, 0x5e, 0x64, 0x10, 0xd7,
				0x55, 0x30, 0x80, 0x99, 0x7e, 0x12, 0xbe, 0x61, 0x90,
				0x9d, 0xe4, 0x0e, 0x0a, 0xe5, 0xb4, 0xf4, 0x25, 0x25,
				0x00, 0x4d, 0x40, 0x10, 0x00, 0x00, 0x05, 0x0c, 0x60,
				0x1f, 0xac, 0x10, 0x03, 0x64, 0x00, 0x00, 0x00, 0x03,
				0x09, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x22, 0x01, 0x01,
				0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
			},
			expected: &PathSwitchRequestAcknowledge{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 2,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 2818572289,
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms:             nrEncryptAlgo,
					NRintegrityProtectionAlgorithms:    nrIntegrityAlgo,
					EUTRAencryptionAlgorithms:          eutraEncryptAlgo,
					EUTRAintegrityProtectionAlgorithms: eutraIntegrityAlgo,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: securityKey,
				},
				PDUSessionResourceSwitchedList: &ie.PDUSessionResourceSwitchedList{
					List: []ie.PDUSessionResourceSwitchedItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 5,
							},
							PathSwitchRequestAcknowledgeTransfer: &pathSwitchRequestAcknowledgeTransferOS,
						},
					},
				},
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd1_case1,
							},
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: sst_case1,
								SD:  sd2_case1,
							},
						},
					},
				},
			},
		},
		{
			name: "Case 2: from ueranemu k8s-basic pipeline TestXnHandover",
			input: []byte{
				0x20, 0x19, 0x00, 0x65, 0x00, 0x00, 0x06, 0x00, 0x0a, 0x40, 0x06, 0x80, 0xec, 0x42, 0xe6, 0x6e,
				0x62, 0x00, 0x55, 0x40, 0x02, 0x00, 0x01, 0x00, 0x77, 0x00, 0x09, 0x18, 0x00, 0x0c, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x5d, 0x00, 0x21, 0x10, 0x95, 0x74, 0x7e, 0x31, 0x52, 0xdf, 0x0d,
				0xe2, 0xe7, 0xee, 0x36, 0x6d, 0xf4, 0x6e, 0xb4, 0x83, 0x39, 0xd1, 0xe8, 0x52, 0x46, 0xca, 0x58,
				0x5b, 0x43, 0xbb, 0x67, 0xcb, 0x41, 0xd6, 0x80, 0x0a, 0x00, 0x4d, 0x40, 0x0e, 0x00, 0x00, 0x0a,
				0x0a, 0x40, 0x1f, 0xac, 0x10, 0x1f, 0x3d, 0x92, 0xf7, 0x52, 0x1f, 0x00, 0x00, 0x00, 0x0a, 0x22,
				0x01, 0x01, 0x02, 0x03, 0x10, 0x08, 0x11, 0x22, 0x33,
			},
			expected: &PathSwitchRequestAcknowledge{
				AMFUENGAPID: &ie.AMFUENGAPID{
					Value: 1014734679650,
				},
				RANUENGAPID: &ie.RANUENGAPID{
					Value: 1,
				},
				UESecurityCapabilities: &ie.UESecurityCapabilities{
					NRencryptionAlgorithms: &ie.NRencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xc0, 0x00},
							BitLength: 16,
						},
					},
					NRintegrityProtectionAlgorithms: &ie.NRintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0xc0, 0x00},
							BitLength: 16,
						},
					},
					EUTRAencryptionAlgorithms: &ie.EUTRAencryptionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0x00, 0x00},
							BitLength: 16,
						},
					},
					EUTRAintegrityProtectionAlgorithms: &ie.EUTRAintegrityProtectionAlgorithms{
						Value: aper.BitString{
							Bytes:     []byte{0x00, 0x00},
							BitLength: 16,
						},
					},
					IEExtensions: nil,
				},
				SecurityContext: &ie.SecurityContext{
					NextHopChainingCount: &ie.NextHopChainingCount{
						Value: 2,
					},
					NextHopNH: &ie.SecurityKey{
						Value: aper.BitString{
							Bytes: []byte{
								0x95, 0x74, 0x7e, 0x31, 0x52, 0xdf, 0x0d, 0xe2,
								0xe7, 0xee, 0x36, 0x6d, 0xf4, 0x6e, 0xb4, 0x83,
								0x39, 0xd1, 0xe8, 0x52, 0x46, 0xca, 0x58, 0x5b,
								0x43, 0xbb, 0x67, 0xcb, 0x41, 0xd6, 0x80, 0x0a,
							},
							BitLength: 256,
						},
					},
					IEExtensions: nil,
				},
				NewSecurityContextInd: nil,
				PDUSessionResourceSwitchedList: &ie.PDUSessionResourceSwitchedList{
					List: []ie.PDUSessionResourceSwitchedItem{
						{
							PDUSessionID: &ie.PDUSessionID{
								Value: 10,
							},
							// will be tested at TestPathSwitchRequestAcknowledgeTransferUnmarshalBinary
							PathSwitchRequestAcknowledgeTransfer: &aper.OctetString{
								0x40, 0x1f, 0xac, 0x10, 0x1f, 0x3d, 0x92, 0xf7, 0x52, 0x1f,
							},
							IEExtensions: nil,
						},
					},
				},
				PDUSessionResourceReleasedListPSAck: nil,
				AllowedNSSAI: &ie.AllowedNSSAI{
					List: []ie.AllowedNSSAIItem{
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x01, 0x02, 0x03},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
						{
							SNSSAI: &ie.SNSSAI{
								SST: &ie.SST{
									Value: aper.OctetString{0x01},
								},
								SD: &ie.SD{
									Value: aper.OctetString{0x11, 0x22, 0x33},
								},
								IEExtensions: nil,
							},
							IEExtensions: nil,
						},
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
