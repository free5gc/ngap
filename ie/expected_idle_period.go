package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ExpectedIdlePeriod struct {
	Value int64
}

func (x *ExpectedIdlePeriod) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	if (x.Value <= 30) || (x.Value == 40) || (x.Value == 50) || (x.Value == 60) || (x.Value == 80) || (x.Value == 100) ||
		(x.Value == 120) || (x.Value == 150) || (x.Value == 180) || (x.Value == 181) {
		*vLb, *vUb = 1, 181
		err = pd.WriteInteger(x.Value, true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	} else {
		err = pd.WriteInteger(x.Value, true, nil, nil)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}
	return nil
}

func (x *ExpectedIdlePeriod) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	if (x.Value <= 30) || (x.Value == 40) || (x.Value == 50) || (x.Value == 60) || (x.Value == 80) || (x.Value == 100) ||
		(x.Value == 120) || (x.Value == 150) || (x.Value == 180) || (x.Value == 181) {
		*vLb, *vUb = 1, 181
		x.Value, err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer unmarshal failed")
		}
	} else {
		x.Value, err = pd.ReadInteger(true, nil, nil)
		if err != nil {
			return errors.Wrap(err, "integer unmarshal failed")
		}
	}
	return nil
}
