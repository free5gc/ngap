package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	EUTRAPagingEDRXCyclePresentHfhalf aper.Enumerated = 0
	EUTRAPagingEDRXCyclePresentHf1    aper.Enumerated = 1
	EUTRAPagingEDRXCyclePresentHf2    aper.Enumerated = 2
	EUTRAPagingEDRXCyclePresentHf4    aper.Enumerated = 3
	EUTRAPagingEDRXCyclePresentHf6    aper.Enumerated = 4
	EUTRAPagingEDRXCyclePresentHf8    aper.Enumerated = 5
	EUTRAPagingEDRXCyclePresentHf10   aper.Enumerated = 6
	EUTRAPagingEDRXCyclePresentHf12   aper.Enumerated = 7
	EUTRAPagingEDRXCyclePresentHf14   aper.Enumerated = 8
	EUTRAPagingEDRXCyclePresentHf16   aper.Enumerated = 9
	EUTRAPagingEDRXCyclePresentHf32   aper.Enumerated = 10
	EUTRAPagingEDRXCyclePresentHf64   aper.Enumerated = 11
	EUTRAPagingEDRXCyclePresentHf128  aper.Enumerated = 12
	EUTRAPagingEDRXCyclePresentHf256  aper.Enumerated = 13
)

type EUTRAPagingEDRXCycle struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:13
}

func (x *EUTRAPagingEDRXCycle) Write(pd *aper.PerBitData) error {
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

func (x *EUTRAPagingEDRXCycle) Read(pd *aper.PerBitData) error {
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
