package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	NumberOfTRPRxTEGPresentTwo   aper.Enumerated = 0
	NumberOfTRPRxTEGPresentThree aper.Enumerated = 1
	NumberOfTRPRxTEGPresentFour  aper.Enumerated = 2
	NumberOfTRPRxTEGPresentSix   aper.Enumerated = 3
	NumberOfTRPRxTEGPresentEight aper.Enumerated = 4
)

type NumberOfTRPRxTEG struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:4
}

func (x *NumberOfTRPRxTEG) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 4
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *NumberOfTRPRxTEG) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 4
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
