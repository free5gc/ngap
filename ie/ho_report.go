package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	HOReportHandoverReportTypePresentHoTooEarly          aper.Enumerated = 0
	HOReportHandoverReportTypePresentHoToWrongCell       aper.Enumerated = 1
	HOReportHandoverReportTypePresentIntersystemPingPong aper.Enumerated = 2
)

type HOReport struct {
	HandoverReportType     *aper.Enumerated                          // valueExt,valueLB:0,valueUB:2
	HandoverCause          *Cause                                    // valueLB:0,valueUB:5
	SourcecellCGI          *NGRANCGI                                 // valueLB:0,valueUB:2
	TargetcellCGI          *NGRANCGI                                 // valueLB:0,valueUB:2
	ReestablishmentcellCGI *NGRANCGI                                 // valueLB:0,valueUB:2,optional
	SourcecellCRNTI        *aper.BitString                           // sizeLB:16,sizeUB:16,optional
	TargetcellinEUTRAN     *EUTRACGI                                 // valueExt,optional
	MobilityInformation    *MobilityInformation                      // optional
	UERLFReportContainer   *UERLFReportContainer                     // valueLB:0,valueUB:2,optional
	IEExtensions           *ProtocolExtensionContainerHOReportExtIEs // optional
}

func (x *HOReport) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	HOReportOptPresentFlag := []bool{}
	// mandatory field
	if x.HandoverReportType == nil {
		return errors.Errorf("HandoverReportType is missing")
	}
	// mandatory field
	if x.HandoverCause == nil {
		return errors.Errorf("HandoverCause is missing")
	}
	// mandatory field
	if x.SourcecellCGI == nil {
		return errors.Errorf("SourcecellCGI is missing")
	}
	// mandatory field
	if x.TargetcellCGI == nil {
		return errors.Errorf("TargetcellCGI is missing")
	}
	// optional field
	if x.ReestablishmentcellCGI != nil {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, true)
	} else {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, false)
	}
	// optional field
	if x.SourcecellCRNTI != nil {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, true)
	} else {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, false)
	}
	// optional field
	if x.TargetcellinEUTRAN != nil {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, true)
	} else {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, false)
	}
	// optional field
	if x.MobilityInformation != nil {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, true)
	} else {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, false)
	}
	// optional field
	if x.UERLFReportContainer != nil {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, true)
	} else {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, true)
	} else {
		HOReportOptPresentFlag = append(HOReportOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(HOReportOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.HandoverReportType), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.HandoverCause.Write(pd)
	if err != nil {
		return errors.Wrap(err, "HandoverCause marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SourcecellCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SourcecellCGI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TargetcellCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargetcellCGI marshal failed")
	}

	// optional field
	if x.ReestablishmentcellCGI != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ReestablishmentcellCGI.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ReestablishmentcellCGI marshal failed")
		}
	}

	// optional field
	if x.SourcecellCRNTI != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 16, 16
		err = pd.WriteBitString(*(x.SourcecellCRNTI), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
	}

	// optional field
	if x.TargetcellinEUTRAN != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TargetcellinEUTRAN.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TargetcellinEUTRAN marshal failed")
		}
	}

	// optional field
	if x.MobilityInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MobilityInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MobilityInformation marshal failed")
		}
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

func (x *HOReport) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	HOReportOptPresentFlag := make([]bool, 6)
	err = pd.ReadSequencePreambleBitMap(&HOReportOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.HandoverReportType = new(aper.Enumerated)
	*(x.HandoverReportType), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.HandoverCause = new(Cause)
	err = x.HandoverCause.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode HandoverCause error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SourcecellCGI = new(NGRANCGI)
	err = x.SourcecellCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SourcecellCGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargetcellCGI = new(NGRANCGI)
	err = x.TargetcellCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargetcellCGI error")
	}

	// optional field (optPresentFlag index: 0)
	if HOReportOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ReestablishmentcellCGI = new(NGRANCGI)
		err = x.ReestablishmentcellCGI.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ReestablishmentcellCGI error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if HOReportOptPresentFlag[1] {
		// Read BitString (Pointer)
		*sLb, *sUb = 16, 16
		x.SourcecellCRNTI = new(aper.BitString)
		*(x.SourcecellCRNTI), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if HOReportOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.TargetcellinEUTRAN = new(EUTRACGI)
		err = x.TargetcellinEUTRAN.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TargetcellinEUTRAN error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if HOReportOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.MobilityInformation = new(MobilityInformation)
		err = x.MobilityInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MobilityInformation error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if HOReportOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.UERLFReportContainer = new(UERLFReportContainer)
		err = x.UERLFReportContainer.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UERLFReportContainer error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if HOReportOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerHOReportExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
