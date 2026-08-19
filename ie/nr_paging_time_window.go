package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NRPagingTimeWindowPresentS1  aper.Enumerated = 0
	NRPagingTimeWindowPresentS2  aper.Enumerated = 1
	NRPagingTimeWindowPresentS3  aper.Enumerated = 2
	NRPagingTimeWindowPresentS4  aper.Enumerated = 3
	NRPagingTimeWindowPresentS5  aper.Enumerated = 4
	NRPagingTimeWindowPresentS6  aper.Enumerated = 5
	NRPagingTimeWindowPresentS7  aper.Enumerated = 6
	NRPagingTimeWindowPresentS8  aper.Enumerated = 7
	NRPagingTimeWindowPresentS9  aper.Enumerated = 8
	NRPagingTimeWindowPresentS10 aper.Enumerated = 9
	NRPagingTimeWindowPresentS11 aper.Enumerated = 10
	NRPagingTimeWindowPresentS12 aper.Enumerated = 11
	NRPagingTimeWindowPresentS13 aper.Enumerated = 12
	NRPagingTimeWindowPresentS14 aper.Enumerated = 13
	NRPagingTimeWindowPresentS15 aper.Enumerated = 14
	NRPagingTimeWindowPresentS16 aper.Enumerated = 15
	NRPagingTimeWindowPresentS17 aper.Enumerated = 16
	NRPagingTimeWindowPresentS18 aper.Enumerated = 17
	NRPagingTimeWindowPresentS19 aper.Enumerated = 18
	NRPagingTimeWindowPresentS20 aper.Enumerated = 19
	NRPagingTimeWindowPresentS21 aper.Enumerated = 20
	NRPagingTimeWindowPresentS22 aper.Enumerated = 21
	NRPagingTimeWindowPresentS23 aper.Enumerated = 22
	NRPagingTimeWindowPresentS24 aper.Enumerated = 23
	NRPagingTimeWindowPresentS25 aper.Enumerated = 24
	NRPagingTimeWindowPresentS26 aper.Enumerated = 25
	NRPagingTimeWindowPresentS27 aper.Enumerated = 26
	NRPagingTimeWindowPresentS28 aper.Enumerated = 27
	NRPagingTimeWindowPresentS29 aper.Enumerated = 28
	NRPagingTimeWindowPresentS30 aper.Enumerated = 29
	NRPagingTimeWindowPresentS31 aper.Enumerated = 30
	NRPagingTimeWindowPresentS32 aper.Enumerated = 31
)

type NRPagingTimeWindow struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:15
}

func (x *NRPagingTimeWindow) Write(pd *aper.PerBitData) error {
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

func (x *NRPagingTimeWindow) Read(pd *aper.PerBitData) error {
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
