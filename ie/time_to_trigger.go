package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	TimeToTriggerPresentMs0    aper.Enumerated = 0
	TimeToTriggerPresentMs40   aper.Enumerated = 1
	TimeToTriggerPresentMs64   aper.Enumerated = 2
	TimeToTriggerPresentMs80   aper.Enumerated = 3
	TimeToTriggerPresentMs100  aper.Enumerated = 4
	TimeToTriggerPresentMs128  aper.Enumerated = 5
	TimeToTriggerPresentMs160  aper.Enumerated = 6
	TimeToTriggerPresentMs256  aper.Enumerated = 7
	TimeToTriggerPresentMs320  aper.Enumerated = 8
	TimeToTriggerPresentMs480  aper.Enumerated = 9
	TimeToTriggerPresentMs512  aper.Enumerated = 10
	TimeToTriggerPresentMs640  aper.Enumerated = 11
	TimeToTriggerPresentMs1024 aper.Enumerated = 12
	TimeToTriggerPresentMs1280 aper.Enumerated = 13
	TimeToTriggerPresentMs2560 aper.Enumerated = 14
	TimeToTriggerPresentMs5120 aper.Enumerated = 15
)

type TimeToTrigger struct {
	Value aper.Enumerated // valueLB:0,valueUB:15
}

func (x *TimeToTrigger) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 15
	err = pd.WriteEnumerated(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *TimeToTrigger) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 15
	x.Value, err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
