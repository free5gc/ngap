package ngap

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

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
