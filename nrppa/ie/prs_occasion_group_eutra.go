package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PRSOccasionGroupEUTRAPresentOg2   aper.Enumerated = 0
	PRSOccasionGroupEUTRAPresentOg4   aper.Enumerated = 1
	PRSOccasionGroupEUTRAPresentOg8   aper.Enumerated = 2
	PRSOccasionGroupEUTRAPresentOg16  aper.Enumerated = 3
	PRSOccasionGroupEUTRAPresentOg32  aper.Enumerated = 4
	PRSOccasionGroupEUTRAPresentOg64  aper.Enumerated = 5
	PRSOccasionGroupEUTRAPresentOg128 aper.Enumerated = 6
)

type PRSOccasionGroupEUTRA struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:6
}

func (x *PRSOccasionGroupEUTRA) Write(pd *aper.PerBitData) error {
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

func (x *PRSOccasionGroupEUTRA) Read(pd *aper.PerBitData) error {
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
