package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	ReportIntervalMDTPresentMs120   aper.Enumerated = 0
	ReportIntervalMDTPresentMs240   aper.Enumerated = 1
	ReportIntervalMDTPresentMs480   aper.Enumerated = 2
	ReportIntervalMDTPresentMs640   aper.Enumerated = 3
	ReportIntervalMDTPresentMs1024  aper.Enumerated = 4
	ReportIntervalMDTPresentMs2048  aper.Enumerated = 5
	ReportIntervalMDTPresentMs5120  aper.Enumerated = 6
	ReportIntervalMDTPresentMs10240 aper.Enumerated = 7
	ReportIntervalMDTPresentMin1    aper.Enumerated = 8
	ReportIntervalMDTPresentMin6    aper.Enumerated = 9
	ReportIntervalMDTPresentMin12   aper.Enumerated = 10
	ReportIntervalMDTPresentMin30   aper.Enumerated = 11
	ReportIntervalMDTPresentMin60   aper.Enumerated = 12
)

type ReportIntervalMDT struct {
	Value aper.Enumerated // valueLB:0,valueUB:12
}

func (x *ReportIntervalMDT) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 12
	err = pd.WriteEnumerated(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *ReportIntervalMDT) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 12
	x.Value, err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
