package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type CancelledCellsInTAIEUTRAItem struct {
	EUTRACGI           *EUTRACGI // valueExt
	NumberOfBroadcasts *NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs // optional
}

func (x *CancelledCellsInTAIEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CancelledCellsInTAIEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.EUTRACGI == nil {
		return errors.Errorf("EUTRACGI is missing")
	}
	// mandatory field
	if x.NumberOfBroadcasts == nil {
		return errors.Errorf("NumberOfBroadcasts is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CancelledCellsInTAIEUTRAItemOptPresentFlag = append(CancelledCellsInTAIEUTRAItemOptPresentFlag, true)
	} else {
		CancelledCellsInTAIEUTRAItemOptPresentFlag = append(CancelledCellsInTAIEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CancelledCellsInTAIEUTRAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRACGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRACGI marshal failed")
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

func (x *CancelledCellsInTAIEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CancelledCellsInTAIEUTRAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CancelledCellsInTAIEUTRAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRACGI = new(EUTRACGI)
	err = x.EUTRACGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRACGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NumberOfBroadcasts = new(NumberOfBroadcasts)
	err = x.NumberOfBroadcasts.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NumberOfBroadcasts error")
	}

	// optional field (optPresentFlag index: 0)
	if CancelledCellsInTAIEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCancelledCellsInTAIEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
