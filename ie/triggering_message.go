package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	TriggeringMessagePresentInitiatingMessage   aper.Enumerated = 0
	TriggeringMessagePresentSuccessfulOutcome   aper.Enumerated = 1
	TriggeringMessagePresentUnsuccessfulOutcome aper.Enumerated = 2
)

type TriggeringMessage struct {
	Value aper.Enumerated // valueLB:0,valueUB:2
}

func (x *TriggeringMessage) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *TriggeringMessage) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 2
	x.Value, err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
