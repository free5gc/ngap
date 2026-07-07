package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PeriodicityItemPresentMs0dot125 aper.Enumerated = 0
	PeriodicityItemPresentMs0dot25  aper.Enumerated = 1
	PeriodicityItemPresentMs0dot5   aper.Enumerated = 2
	PeriodicityItemPresentMs0dot625 aper.Enumerated = 3
	PeriodicityItemPresentMs1       aper.Enumerated = 4
	PeriodicityItemPresentMs1dot25  aper.Enumerated = 5
	PeriodicityItemPresentMs2       aper.Enumerated = 6
	PeriodicityItemPresentMs2dot5   aper.Enumerated = 7
	PeriodicityItemPresentMs4dot    aper.Enumerated = 8
	PeriodicityItemPresentMs5       aper.Enumerated = 9
	PeriodicityItemPresentMs8       aper.Enumerated = 10
	PeriodicityItemPresentMs10      aper.Enumerated = 11
	PeriodicityItemPresentMs16      aper.Enumerated = 12
	PeriodicityItemPresentMs20      aper.Enumerated = 13
	PeriodicityItemPresentMs32      aper.Enumerated = 14
	PeriodicityItemPresentMs40      aper.Enumerated = 15
	PeriodicityItemPresentMs64      aper.Enumerated = 16
	PeriodicityItemPresentMs80m     aper.Enumerated = 17
	PeriodicityItemPresentMs160     aper.Enumerated = 18
	PeriodicityItemPresentMs320     aper.Enumerated = 19
	PeriodicityItemPresentMs640m    aper.Enumerated = 20
	PeriodicityItemPresentMs1280    aper.Enumerated = 21
	PeriodicityItemPresentMs2560    aper.Enumerated = 22
	PeriodicityItemPresentMs5120    aper.Enumerated = 23
	PeriodicityItemPresentMs10240   aper.Enumerated = 24
)

type PeriodicityItem struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:24
}

func (x *PeriodicityItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 24
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *PeriodicityItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 24
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
