package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EmergencyAreaID struct {
	Value aper.OctetString // sizeLB:3,sizeUB:3
}

func (x *EmergencyAreaID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write OctetString
	*sLb, *sUb = 3, 3
	err = pd.WriteOctetString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "octetString marshal failed")
	}
	return nil
}

func (x *EmergencyAreaID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read OctetString
	*sLb, *sUb = 3, 3
	x.Value, err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}
	return nil
}
