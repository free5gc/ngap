package ngap

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeBackupAmfName(t *testing.T) {
	hs := "20150039000004000100060180616d663100600017004002f839cafe00060031302e3130302e3230302e313900" +
		"564001ff005000080002f83900000008"
	bs, err := hex.DecodeString(hs)
	require.NoError(t, err)
	_, err = Decoder(bs)
	require.NoError(t, err)
}
