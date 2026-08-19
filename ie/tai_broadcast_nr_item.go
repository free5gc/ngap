package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TAIBroadcastNRItem struct {
	TAI                   *TAI // valueExt
	CompletedCellsInTAINR *CompletedCellsInTAINR
	IEExtensions          *ProtocolExtensionContainerTAIBroadcastNRItemExtIEs // optional
}

func (x *TAIBroadcastNRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TAIBroadcastNRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TAI == nil {
		return errors.Errorf("TAI is missing")
	}
	// mandatory field
	if x.CompletedCellsInTAINR == nil {
		return errors.Errorf("CompletedCellsInTAINR is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TAIBroadcastNRItemOptPresentFlag = append(TAIBroadcastNRItemOptPresentFlag, true)
	} else {
		TAIBroadcastNRItemOptPresentFlag = append(TAIBroadcastNRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TAIBroadcastNRItemOptPresentFlag, true)
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
	err = x.CompletedCellsInTAINR.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CompletedCellsInTAINR marshal failed")
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

func (x *TAIBroadcastNRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TAIBroadcastNRItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TAIBroadcastNRItemOptPresentFlag, true)
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
	x.CompletedCellsInTAINR = new(CompletedCellsInTAINR)
	err = x.CompletedCellsInTAINR.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CompletedCellsInTAINR error")
	}

	// optional field (optPresentFlag index: 0)
	if TAIBroadcastNRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTAIBroadcastNRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
