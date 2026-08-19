package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TAICancelledEUTRAItem struct {
	TAI                      *TAI // valueExt
	CancelledCellsInTAIEUTRA *CancelledCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs // optional
}

func (x *TAICancelledEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TAICancelledEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TAI == nil {
		return errors.Errorf("TAI is missing")
	}
	// mandatory field
	if x.CancelledCellsInTAIEUTRA == nil {
		return errors.Errorf("CancelledCellsInTAIEUTRA is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TAICancelledEUTRAItemOptPresentFlag = append(TAICancelledEUTRAItemOptPresentFlag, true)
	} else {
		TAICancelledEUTRAItemOptPresentFlag = append(TAICancelledEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TAICancelledEUTRAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.CancelledCellsInTAIEUTRA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CancelledCellsInTAIEUTRA marshal failed")
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

func (x *TAICancelledEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TAICancelledEUTRAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TAICancelledEUTRAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAI = new(TAI)
	err = x.TAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CancelledCellsInTAIEUTRA = new(CancelledCellsInTAIEUTRA)
	err = x.CancelledCellsInTAIEUTRA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CancelledCellsInTAIEUTRA error")
	}

	// optional field (optPresentFlag index: 0)
	if TAICancelledEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
