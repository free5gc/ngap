package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AMFNameVisibleString struct {
	Value aper.VisibleString // sizeExt,sizeLB:1,sizeUB:150
}

func (x *AMFNameVisibleString) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Visible String
	*sLb, *sUb = 1, 150
	err = pd.WriteVisibleString(x.Value, true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "visibleString marshal failed")
	}
	return nil
}

func (x *AMFNameVisibleString) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Visible String
	*sLb, *sUb = 1, 150
	var Value_vs aper.VisibleString
	Value_vs, err = pd.ReadVisibleString(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode visibleString error"))
	}
	x.Value = Value_vs
	return nil
}
