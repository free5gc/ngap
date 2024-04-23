package ngap

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/free5gc/aper"
	"github.com/free5gc/ngap/ngapType"
)

func TestDecodeBackupAmfName(t *testing.T) {
	hs := "20150039000004000100060180616d663100600017004002f839cafe00060031302e3130302e3230302e313900" +
		"564001ff005000080002f83900000008"
	bs, err := hex.DecodeString(hs)
	require.NoError(t, err)
	_, err = Decoder(bs)
	require.NoError(t, err)
}

func TestDecoder(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantPdu *ngapType.NGAPPDU
		wantErr bool
	}{
		{
			name: "TestUncomprehendedProcedure",
			args: args{
				b: []byte{
					0x00, 0x55, 0x00, 0x17, 0x00, 0x00, 0x02, 0x00, 0x52, 0x40, 0x06, 0x01, 0x80, 0x54, 0x45, 0x53,
					0x54, 0x00, 0x52, 0x40, 0x06, 0x01, 0x80, 0x54, 0x45, 0x53, 0x54,
				},
			},
			wantPdu: &ngapType.NGAPPDU{
				Present: ngapType.NGAPPDUPresentInitiatingMessage,
				InitiatingMessage: &ngapType.InitiatingMessage{
					ProcedureCode: ngapType.ProcedureCode{
						Value: 0x55,
					},
					Criticality: ngapType.Criticality{
						Value: ngapType.CriticalityPresentReject,
					},
					Value: ngapType.InitiatingMessageValue{
						Present: 0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "TestUncomprehendedIE",
			args: args{
				b: []byte{
					0x00, 0x15, 0x00, 0x17, 0x00, 0x00, 0x02, 0x55, 0x55, 0x40, 0x06, 0x01, 0x80, 0x54, 0x45, 0x53,
					0x54, 0x00, 0x52, 0x40, 0x06, 0x01, 0x80, 0x54, 0x45, 0x53, 0x54,
				},
			},
			wantPdu: &ngapType.NGAPPDU{
				Present: ngapType.NGAPPDUPresentInitiatingMessage,
				InitiatingMessage: &ngapType.InitiatingMessage{
					ProcedureCode: ngapType.ProcedureCode{
						Value: ngapType.ProcedureCodeNGSetup,
					},
					Criticality: ngapType.Criticality{
						Value: ngapType.CriticalityPresentReject,
					},
					Value: ngapType.InitiatingMessageValue{
						Present: ngapType.InitiatingMessagePresentNGSetupRequest,
						NGSetupRequest: &ngapType.NGSetupRequest{
							ProtocolIEs: ngapType.ProtocolIEContainerNGSetupRequestIEs{
								List: []ngapType.NGSetupRequestIEs{
									{
										Id: ngapType.ProtocolIEID{
											Value: 0x5555,
										},
										Criticality: ngapType.Criticality{
											Value: ngapType.CriticalityPresentIgnore,
										},
										Value: ngapType.NGSetupRequestIEsValue{
											Present: 0,
										},
									},
									{
										Id: ngapType.ProtocolIEID{
											Value: ngapType.ProtocolIEIDRANNodeName,
										},
										Criticality: ngapType.Criticality{
											Value: ngapType.CriticalityPresentIgnore,
										},
										Value: ngapType.NGSetupRequestIEsValue{
											Present: ngapType.NGSetupRequestIEsPresentRANNodeName,
											RANNodeName: &ngapType.RANNodeName{
												Value: "TEST",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test AMFStatusIndication with BackupAMFName IE",
			args: args{
				b: []byte{
					0x00, 0x01, 0x40, 0x15, 0x00, 0x00, 0x01, 0x00, 0x78, 0x00, 0x0e, 0x00, 0x20, 0x02,
					0xf8, 0x39, 0xca, 0xfe, 0x00, 0x01, 0x80, 0x41, 0x4d, 0x46, 0x31, 0x00, 0x00, 0x00,
				},
			},
			wantPdu: &ngapType.NGAPPDU{
				Present: ngapType.NGAPPDUPresentInitiatingMessage,
				InitiatingMessage: &ngapType.InitiatingMessage{
					ProcedureCode: ngapType.ProcedureCode{
						Value: ngapType.ProcedureCodeAMFStatusIndication,
					},
					Criticality: ngapType.Criticality{
						Value: ngapType.CriticalityPresentIgnore,
					},
					Value: ngapType.InitiatingMessageValue{
						Present: ngapType.InitiatingMessagePresentAMFStatusIndication,
						AMFStatusIndication: &ngapType.AMFStatusIndication{
							ProtocolIEs: ngapType.ProtocolIEContainerAMFStatusIndicationIEs{
								List: []ngapType.AMFStatusIndicationIEs{
									{
										Id: ngapType.ProtocolIEID{
											Value: ngapType.ProtocolIEIDUnavailableGUAMIList,
										},
										Criticality: ngapType.Criticality{
											Value: ngapType.CriticalityPresentReject,
										},
										Value: ngapType.AMFStatusIndicationIEsValue{
											Present: ngapType.AMFStatusIndicationIEsPresentUnavailableGUAMIList,
											UnavailableGUAMIList: &ngapType.UnavailableGUAMIList{
												List: []ngapType.UnavailableGUAMIItem{
													{
														GUAMI: ngapType.GUAMI{
															PLMNIdentity: ngapType.PLMNIdentity{
																Value: aper.OctetString{
																	0x02, 0xf8, 0x39,
																},
															},
															AMFRegionID: ngapType.AMFRegionID{
																Value: aper.BitString{
																	Bytes: []byte{
																		0xca,
																	},
																	BitLength: 8,
																},
															},
															AMFSetID: ngapType.AMFSetID{
																Value: aper.BitString{
																	Bytes: []byte{
																		0xfe, 0x00,
																	},
																	BitLength: 10,
																},
															},
															AMFPointer: ngapType.AMFPointer{
																Value: aper.BitString{
																	Bytes: []byte{
																		0x00,
																	},
																	BitLength: 6,
																},
															},
														},
														BackupAMFName: &ngapType.AMFName{
															Value: "AMF1",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotPdu, err := Decoder(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPdu, tt.wantPdu) {
				t.Errorf("Decoder() = %v, want %v", gotPdu, tt.wantPdu)
			}
		})
	}
}
