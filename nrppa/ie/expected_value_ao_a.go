package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ExpectedValueAoA struct {
	Value int64 // valueLB:0,valueUB:3599
}

func (x *ExpectedValueAoA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 3599
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *ExpectedValueAoA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 3599
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}
