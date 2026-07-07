package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	TRPMeasurementQuantitiesItemPresentGNBRxTxTimeDiff aper.Enumerated = 0
	TRPMeasurementQuantitiesItemPresentULSRSRSRP       aper.Enumerated = 1
	TRPMeasurementQuantitiesItemPresentULAoA           aper.Enumerated = 2
	TRPMeasurementQuantitiesItemPresentULRTOA          aper.Enumerated = 3
	TRPMeasurementQuantitiesItemPresentMultipleULAoA   aper.Enumerated = 4
	TRPMeasurementQuantitiesItemPresentULSRSRSRPP      aper.Enumerated = 5
)

type TRPMeasurementQuantitiesItem struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:3
}

func (x *TRPMeasurementQuantitiesItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *TRPMeasurementQuantitiesItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 3
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
