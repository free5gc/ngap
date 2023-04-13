package ngap_test

import (
	"testing"

	"github.com/free5gc/ngap"
)

func FuzzNGAP(f *testing.F) {
	f.Fuzz(func(t *testing.T, d []byte) {
		//nolint:errcheck // fuzzing code
		ngap.Decoder(d)
	})
}
