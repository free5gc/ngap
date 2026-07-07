package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	M7ReportAmountMDTPresentR1       aper.Enumerated = 0
	M7ReportAmountMDTPresentR2       aper.Enumerated = 1
	M7ReportAmountMDTPresentR4       aper.Enumerated = 2
	M7ReportAmountMDTPresentR8       aper.Enumerated = 3
	M7ReportAmountMDTPresentR16      aper.Enumerated = 4
	M7ReportAmountMDTPresentR32      aper.Enumerated = 5
	M7ReportAmountMDTPresentR64      aper.Enumerated = 6
	M7ReportAmountMDTPresentInfinity aper.Enumerated = 7
)

type M7ReportAmountMDT struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:7
}

func (x *M7ReportAmountMDT) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 7
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *M7ReportAmountMDT) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 7
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
