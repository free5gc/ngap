package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	CauseProtocolPresentTransferSyntaxError                          aper.Enumerated = 0
	CauseProtocolPresentAbstractSyntaxErrorReject                    aper.Enumerated = 1
	CauseProtocolPresentAbstractSyntaxErrorIgnoreAndNotify           aper.Enumerated = 2
	CauseProtocolPresentMessageNotCompatibleWithReceiverState        aper.Enumerated = 3
	CauseProtocolPresentSemanticError                                aper.Enumerated = 4
	CauseProtocolPresentUnspecified                                  aper.Enumerated = 5
	CauseProtocolPresentAbstractSyntaxErrorFalselyConstructedMessage aper.Enumerated = 6
)

type CauseProtocol struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:6
}

func (x *CauseProtocol) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 6
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *CauseProtocol) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 6
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
