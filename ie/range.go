package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	RangePresentM50   aper.Enumerated = 0
	RangePresentM80   aper.Enumerated = 1
	RangePresentM180  aper.Enumerated = 2
	RangePresentM200  aper.Enumerated = 3
	RangePresentM350  aper.Enumerated = 4
	RangePresentM400  aper.Enumerated = 5
	RangePresentM500  aper.Enumerated = 6
	RangePresentM700  aper.Enumerated = 7
	RangePresentM1000 aper.Enumerated = 8
)

type Range struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:8
}

func (x *Range) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 8
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *Range) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 8
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
