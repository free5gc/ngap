package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	RxTxTimingErrorMarginPresentTc0dot5 aper.Enumerated = 0
	RxTxTimingErrorMarginPresentTc1     aper.Enumerated = 1
	RxTxTimingErrorMarginPresentTc2     aper.Enumerated = 2
	RxTxTimingErrorMarginPresentTc4     aper.Enumerated = 3
	RxTxTimingErrorMarginPresentTc8     aper.Enumerated = 4
	RxTxTimingErrorMarginPresentTc12    aper.Enumerated = 5
	RxTxTimingErrorMarginPresentTc16    aper.Enumerated = 6
	RxTxTimingErrorMarginPresentTc20    aper.Enumerated = 7
	RxTxTimingErrorMarginPresentTc24    aper.Enumerated = 8
	RxTxTimingErrorMarginPresentTc32    aper.Enumerated = 9
	RxTxTimingErrorMarginPresentTc40    aper.Enumerated = 10
	RxTxTimingErrorMarginPresentTc48    aper.Enumerated = 11
	RxTxTimingErrorMarginPresentTc64    aper.Enumerated = 12
	RxTxTimingErrorMarginPresentTc80    aper.Enumerated = 13
	RxTxTimingErrorMarginPresentTc96    aper.Enumerated = 14
	RxTxTimingErrorMarginPresentTc128   aper.Enumerated = 15
)

type RxTxTimingErrorMargin struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:15
}

func (x *RxTxTimingErrorMargin) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 15
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *RxTxTimingErrorMargin) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 15
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
