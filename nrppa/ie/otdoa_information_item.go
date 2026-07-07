package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	OTDOAInformationItemPresentPci                              aper.Enumerated = 0
	OTDOAInformationItemPresentCGI                              aper.Enumerated = 1
	OTDOAInformationItemPresentTac                              aper.Enumerated = 2
	OTDOAInformationItemPresentEarfcn                           aper.Enumerated = 3
	OTDOAInformationItemPresentPrsBandwidth                     aper.Enumerated = 4
	OTDOAInformationItemPresentPrsConfigIndex                   aper.Enumerated = 5
	OTDOAInformationItemPresentCpLength                         aper.Enumerated = 6
	OTDOAInformationItemPresentNoDlFrames                       aper.Enumerated = 7
	OTDOAInformationItemPresentNoAntennaPorts                   aper.Enumerated = 8
	OTDOAInformationItemPresentSFNInitTime                      aper.Enumerated = 9
	OTDOAInformationItemPresentNGRANAccessPointPosition         aper.Enumerated = 10
	OTDOAInformationItemPresentPrsmutingconfiguration           aper.Enumerated = 11
	OTDOAInformationItemPresentPrsid                            aper.Enumerated = 12
	OTDOAInformationItemPresentTpid                             aper.Enumerated = 13
	OTDOAInformationItemPresentTpType                           aper.Enumerated = 14
	OTDOAInformationItemPresentCrsCPlength                      aper.Enumerated = 15
	OTDOAInformationItemPresentDlBandwidth                      aper.Enumerated = 16
	OTDOAInformationItemPresentMultipleprsConfigurationsperCell aper.Enumerated = 17
	OTDOAInformationItemPresentPrsOccasionGroup                 aper.Enumerated = 18
	OTDOAInformationItemPresentPrsFrequencyHoppingConfiguration aper.Enumerated = 19
	OTDOAInformationItemPresentTddConfig                        aper.Enumerated = 20
)

type OTDOAInformationItem struct {
	Value aper.Enumerated // valueExt,valueLB:0,valueUB:19
}

func (x *OTDOAInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Enumerated
	*vLb, *vUb = 0, 19
	err = pd.WriteEnumerated(x.Value, true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}
	return nil
}

func (x *OTDOAInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Enumerated
	*vLb, *vUb = 0, 19
	x.Value, err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}
	return nil
}
