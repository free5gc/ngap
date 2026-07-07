package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AMFNameUTF8String struct {
	Value aper.UTF8String // sizeExt,sizeLB:1,sizeUB:150
}

func (x *AMFNameUTF8String) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write UTF8 String
	*sLb, *sUb = 1, 150
	err = pd.WriteUTF8String(x.Value, true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "utf8String marshal failed")
	}
	return nil
}

func (x *AMFNameUTF8String) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read UTF8 String
	*sLb, *sUb = 1, 150
	var Value_utf8s aper.UTF8String
	Value_utf8s, err = pd.ReadUTF8String(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode utf8String error"))
	}
	x.Value = Value_utf8s
	return nil
}
