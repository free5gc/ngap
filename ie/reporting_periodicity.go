package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	ReportingPeriodicityPresentStop    aper.Enumerated = 0
	ReportingPeriodicityPresentSingle  aper.Enumerated = 1
	ReportingPeriodicityPresentMs1000  aper.Enumerated = 2
	ReportingPeriodicityPresentMs2000  aper.Enumerated = 3
	ReportingPeriodicityPresentMs5000  aper.Enumerated = 4
	ReportingPeriodicityPresentMs10000 aper.Enumerated = 5
)

type ReportingPeriodicity struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:5
}

func (x *ReportingPeriodicity) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 5
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *ReportingPeriodicity) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 5
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
