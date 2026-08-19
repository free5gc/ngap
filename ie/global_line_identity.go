package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GlobalLineIdentity struct {
	Value aper.OctetString
}

func (x *GlobalLineIdentity) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write OctetString
	sLb, sUb = nil, nil
	err = pd.WriteOctetString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "octetString marshal failed")
	}
	return nil
}

func (x *GlobalLineIdentity) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read OctetString
	sLb, sUb = nil, nil
	x.Value, err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}
	return nil
}
