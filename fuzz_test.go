//go:build go1.18
// +build go1.18

package ngap_test

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"

	"github.com/free5gc/ngap/message"
)

func FuzzNGAP(f *testing.F) {
	f.Fuzz(func(t *testing.T, d []byte) {
		// fuzzing code
		m, err := message.Parse(d)
		if err == nil {
			spew.Dump(m)
			// test re-encoding/re-decoding if the fuzz data is in valid format
			buf, err := m.MarshalBinary() //nolint
			// can't request no error here for MarshalBinary since Parse skipped decode error for
			// mandatory but ignored IE, while MarshalBinary requires all mandatory IEs to be provided
			if err == nil {
				_, err = message.Parse(buf)
				require.NoError(t, err)
			}
		} else {
			// check if the error is due to panic
			if strings.HasPrefix(err.Error(), "Parse(): panic") {
				require.NoError(t, err)
			}
		}
	})
}
