package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	RepetitionFactorExtendedPresentN3  aper.Enumerated = 0
	RepetitionFactorExtendedPresentN5  aper.Enumerated = 1
	RepetitionFactorExtendedPresentN6  aper.Enumerated = 2
	RepetitionFactorExtendedPresentN7  aper.Enumerated = 3
	RepetitionFactorExtendedPresentN8  aper.Enumerated = 4
	RepetitionFactorExtendedPresentN10 aper.Enumerated = 5
	RepetitionFactorExtendedPresentN12 aper.Enumerated = 6
	RepetitionFactorExtendedPresentN14 aper.Enumerated = 7
)

type RepetitionFactorExtended struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:7
}

func (x *RepetitionFactorExtended) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 7
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *RepetitionFactorExtended) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 7
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
