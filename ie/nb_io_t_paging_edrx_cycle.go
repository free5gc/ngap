package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NBIoTPagingEDRXCyclePresentHf2    aper.Enumerated = 0
	NBIoTPagingEDRXCyclePresentHf4    aper.Enumerated = 1
	NBIoTPagingEDRXCyclePresentHf6    aper.Enumerated = 2
	NBIoTPagingEDRXCyclePresentHf8    aper.Enumerated = 3
	NBIoTPagingEDRXCyclePresentHf10   aper.Enumerated = 4
	NBIoTPagingEDRXCyclePresentHf12   aper.Enumerated = 5
	NBIoTPagingEDRXCyclePresentHf14   aper.Enumerated = 6
	NBIoTPagingEDRXCyclePresentHf16   aper.Enumerated = 7
	NBIoTPagingEDRXCyclePresentHf32   aper.Enumerated = 8
	NBIoTPagingEDRXCyclePresentHf64   aper.Enumerated = 9
	NBIoTPagingEDRXCyclePresentHf128  aper.Enumerated = 10
	NBIoTPagingEDRXCyclePresentHf256  aper.Enumerated = 11
	NBIoTPagingEDRXCyclePresentHf512  aper.Enumerated = 12
	NBIoTPagingEDRXCyclePresentHf1024 aper.Enumerated = 13
)

type NBIoTPagingEDRXCycle struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:13
}

func (x *NBIoTPagingEDRXCycle) Write(pd *aper.PerBitData) error {
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

func (x *NBIoTPagingEDRXCycle) Read(pd *aper.PerBitData) error {
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
