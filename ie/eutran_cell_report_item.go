package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EUTRANCellReportItem struct {
	ECGI                                  *EUTRACGI                                             // valueExt
	EUTRANCompositeAvailableCapacityGroup *EUTRANCompositeAvailableCapacityGroup                // valueExt
	EUTRANNumberOfActiveUEs               *EUTRANNumberOfActiveUEs                              // optional
	EUTRANNoofRRCConnections              *NGRANNoofRRCConnections                              // optional
	EUTRANRadioResourceStatus             *EUTRANRadioResourceStatus                            // valueExt,optional
	IEExtensions                          *ProtocolExtensionContainerEUTRANCellReportItemExtIEs // optional
}

func (x *EUTRANCellReportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EUTRANCellReportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.ECGI == nil {
		return errors.Errorf("ECGI is missing")
	}
	// mandatory field
	if x.EUTRANCompositeAvailableCapacityGroup == nil {
		return errors.Errorf("EUTRANCompositeAvailableCapacityGroup is missing")
	}
	// optional field
	if x.EUTRANNumberOfActiveUEs != nil {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, true)
	} else {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, false)
	}
	// optional field
	if x.EUTRANNoofRRCConnections != nil {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, true)
	} else {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, false)
	}
	// optional field
	if x.EUTRANRadioResourceStatus != nil {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, true)
	} else {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, true)
	} else {
		EUTRANCellReportItemOptPresentFlag = append(EUTRANCellReportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EUTRANCellReportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ECGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ECGI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRANCompositeAvailableCapacityGroup.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRANCompositeAvailableCapacityGroup marshal failed")
	}

	// optional field
	if x.EUTRANNumberOfActiveUEs != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.EUTRANNumberOfActiveUEs.Write(pd)
		if err != nil {
			return errors.Wrap(err, "EUTRANNumberOfActiveUEs marshal failed")
		}
	}

	// optional field
	if x.EUTRANNoofRRCConnections != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.EUTRANNoofRRCConnections.Write(pd)
		if err != nil {
			return errors.Wrap(err, "EUTRANNoofRRCConnections marshal failed")
		}
	}

	// optional field
	if x.EUTRANRadioResourceStatus != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.EUTRANRadioResourceStatus.Write(pd)
		if err != nil {
			return errors.Wrap(err, "EUTRANRadioResourceStatus marshal failed")
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

func (x *EUTRANCellReportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EUTRANCellReportItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&EUTRANCellReportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ECGI = new(EUTRACGI)
	err = x.ECGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ECGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRANCompositeAvailableCapacityGroup = new(EUTRANCompositeAvailableCapacityGroup)
	err = x.EUTRANCompositeAvailableCapacityGroup.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRANCompositeAvailableCapacityGroup error")
	}

	// optional field (optPresentFlag index: 0)
	if EUTRANCellReportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.EUTRANNumberOfActiveUEs = new(EUTRANNumberOfActiveUEs)
		err = x.EUTRANNumberOfActiveUEs.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode EUTRANNumberOfActiveUEs error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if EUTRANCellReportItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.EUTRANNoofRRCConnections = new(NGRANNoofRRCConnections)
		err = x.EUTRANNoofRRCConnections.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode EUTRANNoofRRCConnections error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if EUTRANCellReportItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.EUTRANRadioResourceStatus = new(EUTRANRadioResourceStatus)
		err = x.EUTRANRadioResourceStatus.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode EUTRANRadioResourceStatus error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if EUTRANCellReportItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEUTRANCellReportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
