package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	LoggingIntervalPresentMs320    aper.Enumerated = 0
	LoggingIntervalPresentMs640    aper.Enumerated = 1
	LoggingIntervalPresentMs1280   aper.Enumerated = 2
	LoggingIntervalPresentMs2560   aper.Enumerated = 3
	LoggingIntervalPresentMs5120   aper.Enumerated = 4
	LoggingIntervalPresentMs10240  aper.Enumerated = 5
	LoggingIntervalPresentMs20480  aper.Enumerated = 6
	LoggingIntervalPresentMs30720  aper.Enumerated = 7
	LoggingIntervalPresentMs40960  aper.Enumerated = 8
	LoggingIntervalPresentMs61440  aper.Enumerated = 9
	LoggingIntervalPresentInfinity aper.Enumerated = 10
)

type LoggingInterval struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:10
}

func (x *LoggingInterval) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 10
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *LoggingInterval) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 10
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
