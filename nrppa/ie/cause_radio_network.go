package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	CauseRadioNetworkPresentUnspecified                          aper.Enumerated = 0
	CauseRadioNetworkPresentRequestedItemNotSupported            aper.Enumerated = 1
	CauseRadioNetworkPresentRequestedItemTemporarilyNotAvailable aper.Enumerated = 2
	CauseRadioNetworkPresentServingNGRANNodeChanged              aper.Enumerated = 3
	CauseRadioNetworkPresentRequestedItemNotSupportedOnTime      aper.Enumerated = 4
)

type CauseRadioNetwork struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:2
}

func (x *CauseRadioNetwork) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *CauseRadioNetwork) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 2
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
