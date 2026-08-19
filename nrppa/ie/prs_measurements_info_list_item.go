package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PRSMeasurementsInfoListItemMeasPRSPeriodicityPresentMs20  aper.Enumerated = 0
	PRSMeasurementsInfoListItemMeasPRSPeriodicityPresentMs40  aper.Enumerated = 1
	PRSMeasurementsInfoListItemMeasPRSPeriodicityPresentMs80  aper.Enumerated = 2
	PRSMeasurementsInfoListItemMeasPRSPeriodicityPresentMs160 aper.Enumerated = 3
)

const ( /* Enum Type */
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs1dot5 aper.Enumerated = 0
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs3     aper.Enumerated = 1
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs3dot5 aper.Enumerated = 2
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs4     aper.Enumerated = 3
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs5dot5 aper.Enumerated = 4
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs6     aper.Enumerated = 5
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs10    aper.Enumerated = 6
	PRSMeasurementsInfoListItemMeasurementPRSLengthPresentMs20    aper.Enumerated = 7
)

type PRSMeasurementsInfoListItem struct {
	PointA               *int64                                                       // valueLB:0,valueUB:3279165
	MeasPRSPeriodicity   *aper.Enumerated                                             // valueExt,valueLB:0,valueUB:3
	MeasPRSOffset        *int64                                                       // valueExt,valueLB:0,valueUB:159
	MeasurementPRSLength *aper.Enumerated                                             // valueLB:0,valueUB:7
	IEExtensions         *ProtocolExtensionContainerPRSMeasurementsInfoListItemExtIEs // optional
}

func (x *PRSMeasurementsInfoListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMeasurementsInfoListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PointA == nil {
		return errors.Errorf("PointA is missing")
	}
	// mandatory field
	if x.MeasPRSPeriodicity == nil {
		return errors.Errorf("MeasPRSPeriodicity is missing")
	}
	// mandatory field
	if x.MeasPRSOffset == nil {
		return errors.Errorf("MeasPRSOffset is missing")
	}
	// mandatory field
	if x.MeasurementPRSLength == nil {
		return errors.Errorf("MeasurementPRSLength is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSMeasurementsInfoListItemOptPresentFlag = append(PRSMeasurementsInfoListItemOptPresentFlag, true)
	} else {
		PRSMeasurementsInfoListItemOptPresentFlag = append(PRSMeasurementsInfoListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSMeasurementsInfoListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(*(x.PointA), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.MeasPRSPeriodicity), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 159
	err = pd.WriteInteger(*(x.MeasPRSOffset), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 7
	err = pd.WriteEnumerated(*(x.MeasurementPRSLength), false, vLb, vUb)
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

func (x *PRSMeasurementsInfoListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMeasurementsInfoListItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSMeasurementsInfoListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	x.PointA = new(int64)
	*(x.PointA), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.MeasPRSPeriodicity = new(aper.Enumerated)
	*(x.MeasPRSPeriodicity), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 159
	x.MeasPRSOffset = new(int64)
	*(x.MeasPRSOffset), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 7
	x.MeasurementPRSLength = new(aper.Enumerated)
	*(x.MeasurementPRSLength), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if PRSMeasurementsInfoListItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSMeasurementsInfoListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
