package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	ExcessPacketDelayThresholdValuePresentMs0dot25 aper.Enumerated = 0
	ExcessPacketDelayThresholdValuePresentMs0dot5  aper.Enumerated = 1
	ExcessPacketDelayThresholdValuePresentMs1      aper.Enumerated = 2
	ExcessPacketDelayThresholdValuePresentMs2      aper.Enumerated = 3
	ExcessPacketDelayThresholdValuePresentMs4      aper.Enumerated = 4
	ExcessPacketDelayThresholdValuePresentMs5      aper.Enumerated = 5
	ExcessPacketDelayThresholdValuePresentMs10     aper.Enumerated = 6
	ExcessPacketDelayThresholdValuePresentMs20     aper.Enumerated = 7
	ExcessPacketDelayThresholdValuePresentMs30     aper.Enumerated = 8
	ExcessPacketDelayThresholdValuePresentMs40     aper.Enumerated = 9
	ExcessPacketDelayThresholdValuePresentMs50     aper.Enumerated = 10
	ExcessPacketDelayThresholdValuePresentMs60     aper.Enumerated = 11
	ExcessPacketDelayThresholdValuePresentMs70     aper.Enumerated = 12
	ExcessPacketDelayThresholdValuePresentMs80     aper.Enumerated = 13
	ExcessPacketDelayThresholdValuePresentMs90     aper.Enumerated = 14
	ExcessPacketDelayThresholdValuePresentMs100    aper.Enumerated = 15
	ExcessPacketDelayThresholdValuePresentMs150    aper.Enumerated = 16
	ExcessPacketDelayThresholdValuePresentMs300    aper.Enumerated = 17
	ExcessPacketDelayThresholdValuePresentMs500    aper.Enumerated = 18
)

type ExcessPacketDelayThresholdValue struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:18
}

func (x *ExcessPacketDelayThresholdValue) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 18
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *ExcessPacketDelayThresholdValue) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 18
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
