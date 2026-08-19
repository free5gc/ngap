package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	M6reportIntervalPresentMs120   aper.Enumerated = 0
	M6reportIntervalPresentMs240   aper.Enumerated = 1
	M6reportIntervalPresentMs480   aper.Enumerated = 2
	M6reportIntervalPresentMs640   aper.Enumerated = 3
	M6reportIntervalPresentMs1024  aper.Enumerated = 4
	M6reportIntervalPresentMs2048  aper.Enumerated = 5
	M6reportIntervalPresentMs5120  aper.Enumerated = 6
	M6reportIntervalPresentMs10240 aper.Enumerated = 7
	M6reportIntervalPresentMs20480 aper.Enumerated = 8
	M6reportIntervalPresentMs40960 aper.Enumerated = 9
	M6reportIntervalPresentMin1    aper.Enumerated = 10
	M6reportIntervalPresentMin6    aper.Enumerated = 11
	M6reportIntervalPresentMin12   aper.Enumerated = 12
	M6reportIntervalPresentMin30   aper.Enumerated = 13
)

type M6reportInterval struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:13
}

func (x *M6reportInterval) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 13
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *M6reportInterval) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 13
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
