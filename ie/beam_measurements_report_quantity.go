package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	BeamMeasurementsReportQuantityRSRPPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	BeamMeasurementsReportQuantityRSRQPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	BeamMeasurementsReportQuantitySINRPresentTrue aper.Enumerated = 0
)

type BeamMeasurementsReportQuantity struct {
	RSRP         *aper.Enumerated                                                // valueExt,valueLB:0,valueUB:0
	RSRQ         *aper.Enumerated                                                // valueExt,valueLB:0,valueUB:0
	SINR         *aper.Enumerated                                                // valueExt,valueLB:0,valueUB:0
	IEExtensions *ProtocolExtensionContainerBeamMeasurementsReportQuantityExtIEs // optional
}

func (x *BeamMeasurementsReportQuantity) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	BeamMeasurementsReportQuantityOptPresentFlag := []bool{}
	// mandatory field
	if x.RSRP == nil {
		return errors.Errorf("RSRP is missing")
	}
	// mandatory field
	if x.RSRQ == nil {
		return errors.Errorf("RSRQ is missing")
	}
	// mandatory field
	if x.SINR == nil {
		return errors.Errorf("SINR is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		BeamMeasurementsReportQuantityOptPresentFlag = append(BeamMeasurementsReportQuantityOptPresentFlag, true)
	} else {
		BeamMeasurementsReportQuantityOptPresentFlag = append(BeamMeasurementsReportQuantityOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(BeamMeasurementsReportQuantityOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.RSRP), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.RSRQ), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.SINR), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// optional field
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *BeamMeasurementsReportQuantity) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	BeamMeasurementsReportQuantityOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&BeamMeasurementsReportQuantityOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.RSRP = new(aper.Enumerated)
	*(x.RSRP), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.RSRQ = new(aper.Enumerated)
	*(x.RSRQ), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.SINR = new(aper.Enumerated)
	*(x.SINR), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if BeamMeasurementsReportQuantityOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerBeamMeasurementsReportQuantityExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
