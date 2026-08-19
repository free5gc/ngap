package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	EUTRAPagingTimeWindowPresentS1  aper.Enumerated = 0
	EUTRAPagingTimeWindowPresentS2  aper.Enumerated = 1
	EUTRAPagingTimeWindowPresentS3  aper.Enumerated = 2
	EUTRAPagingTimeWindowPresentS4  aper.Enumerated = 3
	EUTRAPagingTimeWindowPresentS5  aper.Enumerated = 4
	EUTRAPagingTimeWindowPresentS6  aper.Enumerated = 5
	EUTRAPagingTimeWindowPresentS7  aper.Enumerated = 6
	EUTRAPagingTimeWindowPresentS8  aper.Enumerated = 7
	EUTRAPagingTimeWindowPresentS9  aper.Enumerated = 8
	EUTRAPagingTimeWindowPresentS10 aper.Enumerated = 9
	EUTRAPagingTimeWindowPresentS11 aper.Enumerated = 10
	EUTRAPagingTimeWindowPresentS12 aper.Enumerated = 11
	EUTRAPagingTimeWindowPresentS13 aper.Enumerated = 12
	EUTRAPagingTimeWindowPresentS14 aper.Enumerated = 13
	EUTRAPagingTimeWindowPresentS15 aper.Enumerated = 14
	EUTRAPagingTimeWindowPresentS16 aper.Enumerated = 15
)

type EUTRAPagingTimeWindow struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:15
}

func (x *EUTRAPagingTimeWindow) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 15
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *EUTRAPagingTimeWindow) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 15
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
