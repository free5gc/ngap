package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PosSIBTypePresentPosSibType11   aper.Enumerated = 0
	PosSIBTypePresentPosSibType12   aper.Enumerated = 1
	PosSIBTypePresentPosSibType13   aper.Enumerated = 2
	PosSIBTypePresentPosSibType14   aper.Enumerated = 3
	PosSIBTypePresentPosSibType15   aper.Enumerated = 4
	PosSIBTypePresentPosSibType16   aper.Enumerated = 5
	PosSIBTypePresentPosSibType17   aper.Enumerated = 6
	PosSIBTypePresentPosSibType18   aper.Enumerated = 7
	PosSIBTypePresentPosSibType21   aper.Enumerated = 8
	PosSIBTypePresentPosSibType22   aper.Enumerated = 9
	PosSIBTypePresentPosSibType23   aper.Enumerated = 10
	PosSIBTypePresentPosSibType24   aper.Enumerated = 11
	PosSIBTypePresentPosSibType25   aper.Enumerated = 12
	PosSIBTypePresentPosSibType26   aper.Enumerated = 13
	PosSIBTypePresentPosSibType27   aper.Enumerated = 14
	PosSIBTypePresentPosSibType28   aper.Enumerated = 15
	PosSIBTypePresentPosSibType29   aper.Enumerated = 16
	PosSIBTypePresentPosSibType210  aper.Enumerated = 17
	PosSIBTypePresentPosSibType211  aper.Enumerated = 18
	PosSIBTypePresentPosSibType212  aper.Enumerated = 19
	PosSIBTypePresentPosSibType213  aper.Enumerated = 20
	PosSIBTypePresentPosSibType214  aper.Enumerated = 21
	PosSIBTypePresentPosSibType215  aper.Enumerated = 22
	PosSIBTypePresentPosSibType216  aper.Enumerated = 23
	PosSIBTypePresentPosSibType217  aper.Enumerated = 24
	PosSIBTypePresentPosSibType218  aper.Enumerated = 25
	PosSIBTypePresentPosSibType219  aper.Enumerated = 26
	PosSIBTypePresentPosSibType220  aper.Enumerated = 27
	PosSIBTypePresentPosSibType221  aper.Enumerated = 28
	PosSIBTypePresentPosSibType222  aper.Enumerated = 29
	PosSIBTypePresentPosSibType223  aper.Enumerated = 30
	PosSIBTypePresentPosSibType224  aper.Enumerated = 31
	PosSIBTypePresentPosSibType225  aper.Enumerated = 32
	PosSIBTypePresentPosSibType31   aper.Enumerated = 33
	PosSIBTypePresentPosSibType41   aper.Enumerated = 34
	PosSIBTypePresentPosSibType51   aper.Enumerated = 35
	PosSIBTypePresentPosSibType61   aper.Enumerated = 36
	PosSIBTypePresentPosSibType62   aper.Enumerated = 37
	PosSIBTypePresentPosSibType63   aper.Enumerated = 38
	PosSIBTypePresentPosSibType19   aper.Enumerated = 39
	PosSIBTypePresentPosSibType110  aper.Enumerated = 40
	PosSIBTypePresentPosSibType64   aper.Enumerated = 41
	PosSIBTypePresentPosSibType65   aper.Enumerated = 42
	PosSIBTypePresentPosSibType66   aper.Enumerated = 43
	PosSIBTypePresentPosSibType217a aper.Enumerated = 44
	PosSIBTypePresentPosSibType218a aper.Enumerated = 45
	PosSIBTypePresentPosSibType220a aper.Enumerated = 46
)

type PosSIBType struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:38
}

func (x *PosSIBType) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 38
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *PosSIBType) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 38
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
