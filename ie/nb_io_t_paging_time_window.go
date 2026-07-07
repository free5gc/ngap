package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NBIoTPagingTimeWindowPresentS1  aper.Enumerated = 0
	NBIoTPagingTimeWindowPresentS2  aper.Enumerated = 1
	NBIoTPagingTimeWindowPresentS3  aper.Enumerated = 2
	NBIoTPagingTimeWindowPresentS4  aper.Enumerated = 3
	NBIoTPagingTimeWindowPresentS5  aper.Enumerated = 4
	NBIoTPagingTimeWindowPresentS6  aper.Enumerated = 5
	NBIoTPagingTimeWindowPresentS7  aper.Enumerated = 6
	NBIoTPagingTimeWindowPresentS8  aper.Enumerated = 7
	NBIoTPagingTimeWindowPresentS9  aper.Enumerated = 8
	NBIoTPagingTimeWindowPresentS10 aper.Enumerated = 9
	NBIoTPagingTimeWindowPresentS11 aper.Enumerated = 10
	NBIoTPagingTimeWindowPresentS12 aper.Enumerated = 11
	NBIoTPagingTimeWindowPresentS13 aper.Enumerated = 12
	NBIoTPagingTimeWindowPresentS14 aper.Enumerated = 13
	NBIoTPagingTimeWindowPresentS15 aper.Enumerated = 14
	NBIoTPagingTimeWindowPresentS16 aper.Enumerated = 15
)

type NBIoTPagingTimeWindow struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:15
}

func (x *NBIoTPagingTimeWindow) Write(pd *aper.PerBitData) error {
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

func (x *NBIoTPagingTimeWindow) Read(pd *aper.PerBitData) error {
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
