package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TooearlyIntersystemHO struct {
	SourcecellID         *EUTRACGI                                              // valueExt
	FailurecellID        *NGRANCGI                                              // valueLB:0,valueUB:2
	UERLFReportContainer *UERLFReportContainer                                  // valueLB:0,valueUB:2,optional
	IEExtensions         *ProtocolExtensionContainerTooearlyIntersystemHOExtIEs // optional
}

func (x *TooearlyIntersystemHO) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TooearlyIntersystemHOOptPresentFlag := []bool{}
	// mandatory field
	if x.SourcecellID == nil {
		return errors.Errorf("SourcecellID is missing")
	}
	// mandatory field
	if x.FailurecellID == nil {
		return errors.Errorf("FailurecellID is missing")
	}
	// optional field
	if x.UERLFReportContainer != nil {
		TooearlyIntersystemHOOptPresentFlag = append(TooearlyIntersystemHOOptPresentFlag, true)
	} else {
		TooearlyIntersystemHOOptPresentFlag = append(TooearlyIntersystemHOOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TooearlyIntersystemHOOptPresentFlag = append(TooearlyIntersystemHOOptPresentFlag, true)
	} else {
		TooearlyIntersystemHOOptPresentFlag = append(TooearlyIntersystemHOOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TooearlyIntersystemHOOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SourcecellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SourcecellID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.FailurecellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FailurecellID marshal failed")
	}

	// optional field
	if x.UERLFReportContainer != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UERLFReportContainer.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UERLFReportContainer marshal failed")
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

func (x *TooearlyIntersystemHO) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TooearlyIntersystemHOOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TooearlyIntersystemHOOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SourcecellID = new(EUTRACGI)
	err = x.SourcecellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SourcecellID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FailurecellID = new(NGRANCGI)
	err = x.FailurecellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FailurecellID error")
	}

	// optional field (optPresentFlag index: 0)
	if TooearlyIntersystemHOOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UERLFReportContainer = new(UERLFReportContainer)
		err = x.UERLFReportContainer.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UERLFReportContainer error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TooearlyIntersystemHOOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTooearlyIntersystemHOExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
