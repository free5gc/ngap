package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSServiceAreaInformation struct {
	MBSServiceAreaCellList *MBSServiceAreaCellList                                    // optional
	MBSServiceAreaTAIList  *MBSServiceAreaTAIList                                     // optional
	IEExtensions           *ProtocolExtensionContainerMBSServiceAreaInformationExtIEs // optional
}

func (x *MBSServiceAreaInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSServiceAreaInformationOptPresentFlag := []bool{}
	// optional field
	if x.MBSServiceAreaCellList != nil {
		MBSServiceAreaInformationOptPresentFlag = append(MBSServiceAreaInformationOptPresentFlag, true)
	} else {
		MBSServiceAreaInformationOptPresentFlag = append(MBSServiceAreaInformationOptPresentFlag, false)
	}
	// optional field
	if x.MBSServiceAreaTAIList != nil {
		MBSServiceAreaInformationOptPresentFlag = append(MBSServiceAreaInformationOptPresentFlag, true)
	} else {
		MBSServiceAreaInformationOptPresentFlag = append(MBSServiceAreaInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSServiceAreaInformationOptPresentFlag = append(MBSServiceAreaInformationOptPresentFlag, true)
	} else {
		MBSServiceAreaInformationOptPresentFlag = append(MBSServiceAreaInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSServiceAreaInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.MBSServiceAreaCellList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSServiceAreaCellList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSServiceAreaCellList marshal failed")
		}
	}

	// optional field
	if x.MBSServiceAreaTAIList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSServiceAreaTAIList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSServiceAreaTAIList marshal failed")
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

func (x *MBSServiceAreaInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSServiceAreaInformationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&MBSServiceAreaInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if MBSServiceAreaInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MBSServiceAreaCellList = new(MBSServiceAreaCellList)
		err = x.MBSServiceAreaCellList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSServiceAreaCellList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSServiceAreaInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MBSServiceAreaTAIList = new(MBSServiceAreaTAIList)
		err = x.MBSServiceAreaTAIList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSServiceAreaTAIList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if MBSServiceAreaInformationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSServiceAreaInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
