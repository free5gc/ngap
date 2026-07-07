package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	IntersystemUnnecessaryHOEarlyIRATHOPresentTrue  aper.Enumerated = 0
	IntersystemUnnecessaryHOEarlyIRATHOPresentFalse aper.Enumerated = 1
)

type IntersystemUnnecessaryHO struct {
	SourcecellID      *NGRANCGI        // valueLB:0,valueUB:2
	TargetcellID      *EUTRACGI        // valueExt
	EarlyIRATHO       *aper.Enumerated // valueExt,valueLB:0,valueUB:1
	CandidateCellList *CandidateCellList
	IEExtensions      *ProtocolExtensionContainerIntersystemUnnecessaryHOExtIEs // optional
}

func (x *IntersystemUnnecessaryHO) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemUnnecessaryHOOptPresentFlag := []bool{}
	// mandatory field
	if x.SourcecellID == nil {
		return errors.Errorf("SourcecellID is missing")
	}
	// mandatory field
	if x.TargetcellID == nil {
		return errors.Errorf("TargetcellID is missing")
	}
	// mandatory field
	if x.EarlyIRATHO == nil {
		return errors.Errorf("EarlyIRATHO is missing")
	}
	// mandatory field
	if x.CandidateCellList == nil {
		return errors.Errorf("CandidateCellList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemUnnecessaryHOOptPresentFlag = append(IntersystemUnnecessaryHOOptPresentFlag, true)
	} else {
		IntersystemUnnecessaryHOOptPresentFlag = append(IntersystemUnnecessaryHOOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemUnnecessaryHOOptPresentFlag, true)
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
	err = x.TargetcellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargetcellID marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.EarlyIRATHO), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.CandidateCellList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CandidateCellList marshal failed")
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

func (x *IntersystemUnnecessaryHO) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemUnnecessaryHOOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemUnnecessaryHOOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SourcecellID = new(NGRANCGI)
	err = x.SourcecellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SourcecellID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargetcellID = new(EUTRACGI)
	err = x.TargetcellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargetcellID error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.EarlyIRATHO = new(aper.Enumerated)
	*(x.EarlyIRATHO), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CandidateCellList = new(CandidateCellList)
	err = x.CandidateCellList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CandidateCellList error")
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemUnnecessaryHOOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemUnnecessaryHOExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
