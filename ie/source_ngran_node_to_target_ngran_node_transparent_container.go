package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SourceNGRANNodeToTargetNGRANNodeTransparentContainer struct {
	RRCContainer                      *RRCContainer
	PDUSessionResourceInformationList *PDUSessionResourceInformationList // optional
	ERABInformationList               *ERABInformationList               // optional
	TargetCellID                      *NGRANCGI                          // valueLB:0,valueUB:2
	IndexToRFSP                       *IndexToRFSP                       // optional
	UEHistoryInformation              *UEHistoryInformation
	IEExtensions                      *ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs // optional
}

func (x *SourceNGRANNodeToTargetNGRANNodeTransparentContainer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag := []bool{}
	// mandatory field
	if x.RRCContainer == nil {
		return errors.Errorf("RRCContainer is missing")
	}
	// optional field
	if x.PDUSessionResourceInformationList != nil {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, true)
	} else {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, false)
	}
	// optional field
	if x.ERABInformationList != nil {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, true)
	} else {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, false)
	}
	// mandatory field
	if x.TargetCellID == nil {
		return errors.Errorf("TargetCellID is missing")
	}
	// optional field
	if x.IndexToRFSP != nil {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, true)
	} else {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, false)
	}
	// mandatory field
	if x.UEHistoryInformation == nil {
		return errors.Errorf("UEHistoryInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, true)
	} else {
		SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag = append(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RRCContainer.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RRCContainer marshal failed")
	}

	// optional field
	if x.PDUSessionResourceInformationList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PDUSessionResourceInformationList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PDUSessionResourceInformationList marshal failed")
		}
	}

	// optional field
	if x.ERABInformationList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ERABInformationList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ERABInformationList marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TargetCellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargetCellID marshal failed")
	}

	// optional field
	if x.IndexToRFSP != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IndexToRFSP.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IndexToRFSP marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.UEHistoryInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UEHistoryInformation marshal failed")
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

func (x *SourceNGRANNodeToTargetNGRANNodeTransparentContainer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RRCContainer = new(RRCContainer)
	err = x.RRCContainer.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RRCContainer error")
	}

	// optional field (optPresentFlag index: 0)
	if SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PDUSessionResourceInformationList = new(PDUSessionResourceInformationList)
		err = x.PDUSessionResourceInformationList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PDUSessionResourceInformationList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ERABInformationList = new(ERABInformationList)
		err = x.ERABInformationList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ERABInformationList error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargetCellID = new(NGRANCGI)
	err = x.TargetCellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargetCellID error")
	}

	// optional field (optPresentFlag index: 2)
	if SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IndexToRFSP = new(IndexToRFSP)
		err = x.IndexToRFSP.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IndexToRFSP error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UEHistoryInformation = new(UEHistoryInformation)
	err = x.UEHistoryInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UEHistoryInformation error")
	}

	// optional field (optPresentFlag index: 3)
	if SourceNGRANNodeToTargetNGRANNodeTransparentContainerOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSourceNGRANNodeToTargetNGRANNodeTransparentContainerExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
