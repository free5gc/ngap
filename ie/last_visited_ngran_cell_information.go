package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LastVisitedNGRANCellInformation struct {
	GlobalCellID                          *NGRANCGI // valueLB:0,valueUB:2
	CellType                              *CellType // valueExt
	TimeUEStayedInCell                    *TimeUEStayedInCell
	TimeUEStayedInCellEnhancedGranularity *TimeUEStayedInCellEnhancedGranularity                           // optional
	HOCauseValue                          *Cause                                                           // valueLB:0,valueUB:5,optional
	IEExtensions                          *ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs // optional
}

func (x *LastVisitedNGRANCellInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LastVisitedNGRANCellInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.GlobalCellID == nil {
		return errors.Errorf("GlobalCellID is missing")
	}
	// mandatory field
	if x.CellType == nil {
		return errors.Errorf("CellType is missing")
	}
	// mandatory field
	if x.TimeUEStayedInCell == nil {
		return errors.Errorf("TimeUEStayedInCell is missing")
	}
	// optional field
	if x.TimeUEStayedInCellEnhancedGranularity != nil {
		LastVisitedNGRANCellInformationOptPresentFlag = append(LastVisitedNGRANCellInformationOptPresentFlag, true)
	} else {
		LastVisitedNGRANCellInformationOptPresentFlag = append(LastVisitedNGRANCellInformationOptPresentFlag, false)
	}
	// optional field
	if x.HOCauseValue != nil {
		LastVisitedNGRANCellInformationOptPresentFlag = append(LastVisitedNGRANCellInformationOptPresentFlag, true)
	} else {
		LastVisitedNGRANCellInformationOptPresentFlag = append(LastVisitedNGRANCellInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		LastVisitedNGRANCellInformationOptPresentFlag = append(LastVisitedNGRANCellInformationOptPresentFlag, true)
	} else {
		LastVisitedNGRANCellInformationOptPresentFlag = append(LastVisitedNGRANCellInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LastVisitedNGRANCellInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.GlobalCellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GlobalCellID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.CellType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CellType marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TimeUEStayedInCell.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TimeUEStayedInCell marshal failed")
	}

	// optional field
	if x.TimeUEStayedInCellEnhancedGranularity != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TimeUEStayedInCellEnhancedGranularity.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TimeUEStayedInCellEnhancedGranularity marshal failed")
		}
	}

	// optional field
	if x.HOCauseValue != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.HOCauseValue.Write(pd)
		if err != nil {
			return errors.Wrap(err, "HOCauseValue marshal failed")
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

func (x *LastVisitedNGRANCellInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LastVisitedNGRANCellInformationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&LastVisitedNGRANCellInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GlobalCellID = new(NGRANCGI)
	err = x.GlobalCellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GlobalCellID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CellType = new(CellType)
	err = x.CellType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CellType error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TimeUEStayedInCell = new(TimeUEStayedInCell)
	err = x.TimeUEStayedInCell.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TimeUEStayedInCell error")
	}

	// optional field (optPresentFlag index: 0)
	if LastVisitedNGRANCellInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TimeUEStayedInCellEnhancedGranularity = new(TimeUEStayedInCellEnhancedGranularity)
		err = x.TimeUEStayedInCellEnhancedGranularity.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TimeUEStayedInCellEnhancedGranularity error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if LastVisitedNGRANCellInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.HOCauseValue = new(Cause)
		err = x.HOCauseValue.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode HOCauseValue error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if LastVisitedNGRANCellInformationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLastVisitedNGRANCellInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
