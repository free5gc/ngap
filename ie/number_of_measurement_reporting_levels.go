package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NumberOfMeasurementReportingLevelsPresentN2  aper.Enumerated = 0
	NumberOfMeasurementReportingLevelsPresentN3  aper.Enumerated = 1
	NumberOfMeasurementReportingLevelsPresentN4  aper.Enumerated = 2
	NumberOfMeasurementReportingLevelsPresentN5  aper.Enumerated = 3
	NumberOfMeasurementReportingLevelsPresentN10 aper.Enumerated = 4
	NumberOfMeasurementReportingLevelsPresentN0  aper.Enumerated = 5
)

type NumberOfMeasurementReportingLevels struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:4
}

func (x *NumberOfMeasurementReportingLevels) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 4
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *NumberOfMeasurementReportingLevels) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 4
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
