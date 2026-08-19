package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NRencryptionAlgorithms struct {
	Value aper.BitString // sizeExt,sizeLB:16,sizeUB:16
}

func (x *NRencryptionAlgorithms) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 16, 16
	err = pd.WriteBitString(x.Value, true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *NRencryptionAlgorithms) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 16, 16
	x.Value, err = pd.ReadBitString(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}
