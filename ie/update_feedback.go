package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UpdateFeedback struct {
	Value aper.BitString // sizeExt,sizeLB:8,sizeUB:8
}

func (x *UpdateFeedback) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 8, 8
	err = pd.WriteBitString(x.Value, true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *UpdateFeedback) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 8, 8
	x.Value, err = pd.ReadBitString(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}
