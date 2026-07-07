package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGRANCellReportItem struct {
	NGRANCGI                             *NGRANCGI                                            // valueLB:0,valueUB:2
	NGRANCompositeAvailableCapacityGroup *EUTRANCompositeAvailableCapacityGroup               // valueExt
	NGRANNumberOfActiveUEs               *NGRANNumberOfActiveUEs                              // optional
	NGRANNoofRRCConnections              *NGRANNoofRRCConnections                             // optional
	NGRANRadioResourceStatus             *NGRANRadioResourceStatus                            // valueExt,optional
	IEExtensions                         *ProtocolExtensionContainerNGRANCellReportItemExtIEs // optional
}

func (x *NGRANCellReportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANCellReportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCGI == nil {
		return errors.Errorf("NGRANCGI is missing")
	}
	// mandatory field
	if x.NGRANCompositeAvailableCapacityGroup == nil {
		return errors.Errorf("NGRANCompositeAvailableCapacityGroup is missing")
	}
	// optional field
	if x.NGRANNumberOfActiveUEs != nil {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, true)
	} else {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, false)
	}
	// optional field
	if x.NGRANNoofRRCConnections != nil {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, true)
	} else {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, false)
	}
	// optional field
	if x.NGRANRadioResourceStatus != nil {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, true)
	} else {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, true)
	} else {
		NGRANCellReportItemOptPresentFlag = append(NGRANCellReportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANCellReportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCGI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCompositeAvailableCapacityGroup.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCompositeAvailableCapacityGroup marshal failed")
	}

	// optional field
	if x.NGRANNumberOfActiveUEs != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NGRANNumberOfActiveUEs.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NGRANNumberOfActiveUEs marshal failed")
		}
	}

	// optional field
	if x.NGRANNoofRRCConnections != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NGRANNoofRRCConnections.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NGRANNoofRRCConnections marshal failed")
		}
	}

	// optional field
	if x.NGRANRadioResourceStatus != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NGRANRadioResourceStatus.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NGRANRadioResourceStatus marshal failed")
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

func (x *NGRANCellReportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANCellReportItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&NGRANCellReportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCGI = new(NGRANCGI)
	err = x.NGRANCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCompositeAvailableCapacityGroup = new(EUTRANCompositeAvailableCapacityGroup)
	err = x.NGRANCompositeAvailableCapacityGroup.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCompositeAvailableCapacityGroup error")
	}

	// optional field (optPresentFlag index: 0)
	if NGRANCellReportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NGRANNumberOfActiveUEs = new(NGRANNumberOfActiveUEs)
		err = x.NGRANNumberOfActiveUEs.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NGRANNumberOfActiveUEs error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if NGRANCellReportItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.NGRANNoofRRCConnections = new(NGRANNoofRRCConnections)
		err = x.NGRANNoofRRCConnections.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NGRANNoofRRCConnections error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if NGRANCellReportItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.NGRANRadioResourceStatus = new(NGRANRadioResourceStatus)
		err = x.NGRANRadioResourceStatus.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NGRANRadioResourceStatus error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if NGRANCellReportItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANCellReportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
