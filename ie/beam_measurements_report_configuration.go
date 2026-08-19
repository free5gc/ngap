package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type BeamMeasurementsReportConfiguration struct {
	BeamMeasurementsReportQuantity *BeamMeasurementsReportQuantity                                      // valueExt,optional
	MaxNrofRSIndexesToReport       *MaxNrofRSIndexesToReport                                            // optional
	IEExtensions                   *ProtocolExtensionContainerBeamMeasurementsReportConfigurationExtIEs // optional
}

func (x *BeamMeasurementsReportConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	BeamMeasurementsReportConfigurationOptPresentFlag := []bool{}
	// optional field
	if x.BeamMeasurementsReportQuantity != nil {
		BeamMeasurementsReportConfigurationOptPresentFlag = append(BeamMeasurementsReportConfigurationOptPresentFlag, true)
	} else {
		BeamMeasurementsReportConfigurationOptPresentFlag = append(BeamMeasurementsReportConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.MaxNrofRSIndexesToReport != nil {
		BeamMeasurementsReportConfigurationOptPresentFlag = append(BeamMeasurementsReportConfigurationOptPresentFlag, true)
	} else {
		BeamMeasurementsReportConfigurationOptPresentFlag = append(BeamMeasurementsReportConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		BeamMeasurementsReportConfigurationOptPresentFlag = append(BeamMeasurementsReportConfigurationOptPresentFlag, true)
	} else {
		BeamMeasurementsReportConfigurationOptPresentFlag = append(BeamMeasurementsReportConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(BeamMeasurementsReportConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.BeamMeasurementsReportQuantity != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.BeamMeasurementsReportQuantity.Write(pd)
		if err != nil {
			return errors.Wrap(err, "BeamMeasurementsReportQuantity marshal failed")
		}
	}

	// optional field
	if x.MaxNrofRSIndexesToReport != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MaxNrofRSIndexesToReport.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MaxNrofRSIndexesToReport marshal failed")
		}
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

func (x *BeamMeasurementsReportConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	BeamMeasurementsReportConfigurationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&BeamMeasurementsReportConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if BeamMeasurementsReportConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.BeamMeasurementsReportQuantity = new(BeamMeasurementsReportQuantity)
		err = x.BeamMeasurementsReportQuantity.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BeamMeasurementsReportQuantity error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if BeamMeasurementsReportConfigurationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MaxNrofRSIndexesToReport = new(MaxNrofRSIndexesToReport)
		err = x.MaxNrofRSIndexesToReport.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MaxNrofRSIndexesToReport error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if BeamMeasurementsReportConfigurationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerBeamMeasurementsReportConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
