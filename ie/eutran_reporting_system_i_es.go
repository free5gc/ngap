package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EUTRANReportingSystemIEs struct {
	EUTRANCellToReportList *EUTRANCellToReportList
	IEExtensions           *ProtocolExtensionContainerEUTRANReportingSystemIEsExtIEs // optional
}

func (x *EUTRANReportingSystemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EUTRANReportingSystemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.EUTRANCellToReportList == nil {
		return errors.Errorf("EUTRANCellToReportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EUTRANReportingSystemIEsOptPresentFlag = append(EUTRANReportingSystemIEsOptPresentFlag, true)
	} else {
		EUTRANReportingSystemIEsOptPresentFlag = append(EUTRANReportingSystemIEsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EUTRANReportingSystemIEsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRANCellToReportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRANCellToReportList marshal failed")
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

func (x *EUTRANReportingSystemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EUTRANReportingSystemIEsOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EUTRANReportingSystemIEsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRANCellToReportList = new(EUTRANCellToReportList)
	err = x.EUTRANCellToReportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRANCellToReportList error")
	}

	// optional field (optPresentFlag index: 0)
	if EUTRANReportingSystemIEsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEUTRANReportingSystemIEsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
