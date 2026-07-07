package message

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/free5gc/ngap/aper"
	"github.com/free5gc/ngap/nrppa/ie"
	"github.com/stretchr/testify/require"
)

func TestECIDMeasurementReport(t *testing.T) {
	// Create comprehensive E-CID Measurement Report
	report := &ECIDMeasurementReport{
		NRPPATransactionID: 12345,
		LMFUEMeasurementID: &ie.UEMeasurementID{
			Value: 1,
		},
		RANUEMeasurementID: &ie.UEMeasurementID{
			Value: 1,
		},
		ECIDMeasurementResult: &ie.ECIDMeasurementResult{
			// Serving cell information
			ServingCellID: &ie.NGRANCGI{
				PLMNIdentity: &ie.PLMNIdentity{
					Value: aper.OctetString("\x02\xf8\x39"), // MCC=208, MNC=93
				},
				NGRANcell: &ie.NGRANCell{
					Choice: &ie.NRCellIdentifier{
						Value: aper.BitString{
							Bytes:     []byte{0x00, 0x00, 0x00, 0x00, 0x10},
							BitLength: 36,
						},
					},
				},
			},
			ServingCellTAC: &ie.TAC{
				Value: aper.OctetString("\x00\x01\x02"), // TAC=000102
			},
		},
	}

	encodedPdu, err := report.MarshalBinary()
	require.NoError(t, err)

	// Encoded result can be verified with https://www.marben-products.com/decoder-asn1-nr/
	t.Logf("encoded result: %s", hex.EncodeToString(encodedPdu))

	decodedMsg, err := Parse(encodedPdu)
	require.NoError(t, err)

	require.True(t, reflect.DeepEqual(report, decodedMsg.(*ECIDMeasurementReport)))

}
