package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PRSBandwidthEUTRAPresentBw6   aper.Enumerated = 0
	PRSBandwidthEUTRAPresentBw15  aper.Enumerated = 1
	PRSBandwidthEUTRAPresentBw25  aper.Enumerated = 2
	PRSBandwidthEUTRAPresentBw50  aper.Enumerated = 3
	PRSBandwidthEUTRAPresentBw75  aper.Enumerated = 4
	PRSBandwidthEUTRAPresentBw100 aper.Enumerated = 5
)

type PRSBandwidthEUTRA struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:5
}

func (x *PRSBandwidthEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 5
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *PRSBandwidthEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 5
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
