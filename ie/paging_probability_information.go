package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	PagingProbabilityInformationPresentP00  aper.Enumerated = 0
	PagingProbabilityInformationPresentP05  aper.Enumerated = 1
	PagingProbabilityInformationPresentP10  aper.Enumerated = 2
	PagingProbabilityInformationPresentP15  aper.Enumerated = 3
	PagingProbabilityInformationPresentP20  aper.Enumerated = 4
	PagingProbabilityInformationPresentP25  aper.Enumerated = 5
	PagingProbabilityInformationPresentP30  aper.Enumerated = 6
	PagingProbabilityInformationPresentP35  aper.Enumerated = 7
	PagingProbabilityInformationPresentP40  aper.Enumerated = 8
	PagingProbabilityInformationPresentP45  aper.Enumerated = 9
	PagingProbabilityInformationPresentP50  aper.Enumerated = 10
	PagingProbabilityInformationPresentP55  aper.Enumerated = 11
	PagingProbabilityInformationPresentP60  aper.Enumerated = 12
	PagingProbabilityInformationPresentP65  aper.Enumerated = 13
	PagingProbabilityInformationPresentP70  aper.Enumerated = 14
	PagingProbabilityInformationPresentP75  aper.Enumerated = 15
	PagingProbabilityInformationPresentP80  aper.Enumerated = 16
	PagingProbabilityInformationPresentP85  aper.Enumerated = 17
	PagingProbabilityInformationPresentP90  aper.Enumerated = 18
	PagingProbabilityInformationPresentP95  aper.Enumerated = 19
	PagingProbabilityInformationPresentP100 aper.Enumerated = 20
)

type PagingProbabilityInformation struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:20
}

func (x *PagingProbabilityInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 20
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *PagingProbabilityInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 20
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
