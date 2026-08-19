package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AreaOfInterest struct {
	AreaOfInterestTAIList     *AreaOfInterestTAIList                          // optional
	AreaOfInterestCellList    *AreaOfInterestCellList                         // optional
	AreaOfInterestRANNodeList *AreaOfInterestRANNodeList                      // optional
	IEExtensions              *ProtocolExtensionContainerAreaOfInterestExtIEs // optional
}

func (x *AreaOfInterest) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AreaOfInterestOptPresentFlag := []bool{}
	// optional field
	if x.AreaOfInterestTAIList != nil {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, true)
	} else {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, false)
	}
	// optional field
	if x.AreaOfInterestCellList != nil {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, true)
	} else {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, false)
	}
	// optional field
	if x.AreaOfInterestRANNodeList != nil {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, true)
	} else {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, true)
	} else {
		AreaOfInterestOptPresentFlag = append(AreaOfInterestOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AreaOfInterestOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.AreaOfInterestTAIList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AreaOfInterestTAIList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AreaOfInterestTAIList marshal failed")
		}
	}

	// optional field
	if x.AreaOfInterestCellList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AreaOfInterestCellList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AreaOfInterestCellList marshal failed")
		}
	}

	// optional field
	if x.AreaOfInterestRANNodeList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AreaOfInterestRANNodeList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AreaOfInterestRANNodeList marshal failed")
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

func (x *AreaOfInterest) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AreaOfInterestOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&AreaOfInterestOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if AreaOfInterestOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AreaOfInterestTAIList = new(AreaOfInterestTAIList)
		err = x.AreaOfInterestTAIList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AreaOfInterestTAIList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AreaOfInterestOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.AreaOfInterestCellList = new(AreaOfInterestCellList)
		err = x.AreaOfInterestCellList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AreaOfInterestCellList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if AreaOfInterestOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.AreaOfInterestRANNodeList = new(AreaOfInterestRANNodeList)
		err = x.AreaOfInterestRANNodeList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AreaOfInterestRANNodeList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if AreaOfInterestOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAreaOfInterestExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
