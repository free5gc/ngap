package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	MeasurementQuantitiesValuePresentCellID             aper.Enumerated = 0
	MeasurementQuantitiesValuePresentAngleOfArrival     aper.Enumerated = 1
	MeasurementQuantitiesValuePresentTimingAdvanceType1 aper.Enumerated = 2
	MeasurementQuantitiesValuePresentTimingAdvanceType2 aper.Enumerated = 3
	MeasurementQuantitiesValuePresentRSRP               aper.Enumerated = 4
	MeasurementQuantitiesValuePresentRSRQ               aper.Enumerated = 5
	MeasurementQuantitiesValuePresentSSRSRP             aper.Enumerated = 6
	MeasurementQuantitiesValuePresentSSRSRQ             aper.Enumerated = 7
	MeasurementQuantitiesValuePresentCSIRSRP            aper.Enumerated = 8
	MeasurementQuantitiesValuePresentCSIRSRQ            aper.Enumerated = 9
	MeasurementQuantitiesValuePresentAngleOfArrivalNR   aper.Enumerated = 10
	MeasurementQuantitiesValuePresentTimingAdvanceNR    aper.Enumerated = 11
)

type MeasurementQuantitiesValue struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:5
}

func (x *MeasurementQuantitiesValue) Write(pd *aper.PerBitData) error {
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

func (x *MeasurementQuantitiesValue) Read(pd *aper.PerBitData) error {
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
