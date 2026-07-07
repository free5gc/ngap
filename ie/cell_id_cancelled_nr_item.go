package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CellIDCancelledNRItem struct {
	NRCGI              *NRCGI // valueExt
	NumberOfBroadcasts *NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledNRItemExtIEs // optional
}

func (x *CellIDCancelledNRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CellIDCancelledNRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NRCGI == nil {
		return errors.Errorf("NRCGI is missing")
	}
	// mandatory field
	if x.NumberOfBroadcasts == nil {
		return errors.Errorf("NumberOfBroadcasts is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CellIDCancelledNRItemOptPresentFlag = append(CellIDCancelledNRItemOptPresentFlag, true)
	} else {
		CellIDCancelledNRItemOptPresentFlag = append(CellIDCancelledNRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CellIDCancelledNRItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRCGI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NumberOfBroadcasts.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NumberOfBroadcasts marshal failed")
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

func (x *CellIDCancelledNRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CellIDCancelledNRItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CellIDCancelledNRItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRCGI = new(NRCGI)
	err = x.NRCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRCGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NumberOfBroadcasts = new(NumberOfBroadcasts)
	err = x.NumberOfBroadcasts.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NumberOfBroadcasts error")
	}

	// optional field (optPresentFlag index: 0)
	if CellIDCancelledNRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCellIDCancelledNRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
