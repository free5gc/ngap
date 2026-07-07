package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NRPagingEDRXCyclePresentHfquarter aper.Enumerated = 0
	NRPagingEDRXCyclePresentHfhalf    aper.Enumerated = 1
	NRPagingEDRXCyclePresentHf1       aper.Enumerated = 2
	NRPagingEDRXCyclePresentHf2       aper.Enumerated = 3
	NRPagingEDRXCyclePresentHf4       aper.Enumerated = 4
	NRPagingEDRXCyclePresentHf8       aper.Enumerated = 5
	NRPagingEDRXCyclePresentHf16      aper.Enumerated = 6
	NRPagingEDRXCyclePresentHf32      aper.Enumerated = 7
	NRPagingEDRXCyclePresentHf64      aper.Enumerated = 8
	NRPagingEDRXCyclePresentHf128     aper.Enumerated = 9
	NRPagingEDRXCyclePresentHf256     aper.Enumerated = 10
	NRPagingEDRXCyclePresentHf512     aper.Enumerated = 11
	NRPagingEDRXCyclePresentHf1024    aper.Enumerated = 12
)

type NRPagingEDRXCycle struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:12
}

func (x *NRPagingEDRXCycle) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 12
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *NRPagingEDRXCycle) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 12
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
