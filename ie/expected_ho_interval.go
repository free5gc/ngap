package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	ExpectedHOIntervalPresentSec15    aper.Enumerated = 0
	ExpectedHOIntervalPresentSec30    aper.Enumerated = 1
	ExpectedHOIntervalPresentSec60    aper.Enumerated = 2
	ExpectedHOIntervalPresentSec90    aper.Enumerated = 3
	ExpectedHOIntervalPresentSec120   aper.Enumerated = 4
	ExpectedHOIntervalPresentSec180   aper.Enumerated = 5
	ExpectedHOIntervalPresentLongTime aper.Enumerated = 6
)

type ExpectedHOInterval struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:6
}

func (x *ExpectedHOInterval) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 6
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *ExpectedHOInterval) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 6
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
