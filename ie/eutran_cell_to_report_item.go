package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EUTRANCellToReportItem struct {
	ECGI         *EUTRACGI                                               // valueExt
	IEExtensions *ProtocolExtensionContainerEUTRANCellToReportItemExtIEs // optional
}

func (x *EUTRANCellToReportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EUTRANCellToReportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.ECGI == nil {
		return errors.Errorf("ECGI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EUTRANCellToReportItemOptPresentFlag = append(EUTRANCellToReportItemOptPresentFlag, true)
	} else {
		EUTRANCellToReportItemOptPresentFlag = append(EUTRANCellToReportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EUTRANCellToReportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ECGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ECGI marshal failed")
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

func (x *EUTRANCellToReportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EUTRANCellToReportItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EUTRANCellToReportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ECGI = new(EUTRACGI)
	err = x.ECGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ECGI error")
	}

	// optional field (optPresentFlag index: 0)
	if EUTRANCellToReportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEUTRANCellToReportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
