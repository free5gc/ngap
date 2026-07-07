package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	NumberOfTRPRxTxTEGPresentTwo   aper.Enumerated = 0
	NumberOfTRPRxTxTEGPresentThree aper.Enumerated = 1
	NumberOfTRPRxTxTEGPresentFour  aper.Enumerated = 2
	NumberOfTRPRxTxTEGPresentSix   aper.Enumerated = 3
	NumberOfTRPRxTxTEGPresentEight aper.Enumerated = 4
)

type NumberOfTRPRxTxTEG struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:4
}

func (x *NumberOfTRPRxTxTEG) Write(pd *aper.PerBitData) error {
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

func (x *NumberOfTRPRxTxTEG) Read(pd *aper.PerBitData) error {
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
