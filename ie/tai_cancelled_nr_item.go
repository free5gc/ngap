package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TAICancelledNRItem struct {
	TAI                   *TAI // valueExt
	CancelledCellsInTAINR *CancelledCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAICancelledNRItemExtIEs // optional
}

func (x *TAICancelledNRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TAICancelledNRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TAI == nil {
		return errors.Errorf("TAI is missing")
	}
	// mandatory field
	if x.CancelledCellsInTAINR == nil {
		return errors.Errorf("CancelledCellsInTAINR is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TAICancelledNRItemOptPresentFlag = append(TAICancelledNRItemOptPresentFlag, true)
	} else {
		TAICancelledNRItemOptPresentFlag = append(TAICancelledNRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TAICancelledNRItemOptPresentFlag, true)
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
	err = x.CancelledCellsInTAINR.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CancelledCellsInTAINR marshal failed")
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

func (x *TAICancelledNRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TAICancelledNRItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TAICancelledNRItemOptPresentFlag, true)
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
	x.CancelledCellsInTAINR = new(CancelledCellsInTAINR)
	err = x.CancelledCellsInTAINR.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CancelledCellsInTAINR error")
	}

	// optional field (optPresentFlag index: 0)
	if TAICancelledNRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTAICancelledNRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
