package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPMeasurementQuantitiesListItem struct {
	TRPMeasurementQuantitiesItem     *TRPMeasurementQuantitiesItem                                     // valueExt,valueLB:0,valueUB:3
	TimingReportingGranularityFactor *int64                                                            // valueLB:0,valueUB:5,optional
	IEExtensions                     *ProtocolExtensionContainerTRPMeasurementQuantitiesListItemExtIEs // optional
}

func (x *TRPMeasurementQuantitiesListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementQuantitiesListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPMeasurementQuantitiesItem == nil {
		return errors.Errorf("TRPMeasurementQuantitiesItem is missing")
	}
	// optional field
	if x.TimingReportingGranularityFactor != nil {
		TRPMeasurementQuantitiesListItemOptPresentFlag = append(TRPMeasurementQuantitiesListItemOptPresentFlag, true)
	} else {
		TRPMeasurementQuantitiesListItemOptPresentFlag = append(TRPMeasurementQuantitiesListItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TRPMeasurementQuantitiesListItemOptPresentFlag = append(TRPMeasurementQuantitiesListItemOptPresentFlag, true)
	} else {
		TRPMeasurementQuantitiesListItemOptPresentFlag = append(TRPMeasurementQuantitiesListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementQuantitiesListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPMeasurementQuantitiesItem.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPMeasurementQuantitiesItem marshal failed")
	}

	// optional field
	if x.TimingReportingGranularityFactor != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 5
		err = pd.WriteInteger(*(x.TimingReportingGranularityFactor), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *TRPMeasurementQuantitiesListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementQuantitiesListItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementQuantitiesListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPMeasurementQuantitiesItem = new(TRPMeasurementQuantitiesItem)
	err = x.TRPMeasurementQuantitiesItem.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPMeasurementQuantitiesItem error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPMeasurementQuantitiesListItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 5
		x.TimingReportingGranularityFactor = new(int64)
		*(x.TimingReportingGranularityFactor), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if TRPMeasurementQuantitiesListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPMeasurementQuantitiesListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
