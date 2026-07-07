package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGRANCellToReportItem struct {
	NGRANCGI     *NGRANCGI                                              // valueLB:0,valueUB:2
	IEExtensions *ProtocolExtensionContainerNGRANCellToReportItemExtIEs // optional
}

func (x *NGRANCellToReportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANCellToReportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCGI == nil {
		return errors.Errorf("NGRANCGI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANCellToReportItemOptPresentFlag = append(NGRANCellToReportItemOptPresentFlag, true)
	} else {
		NGRANCellToReportItemOptPresentFlag = append(NGRANCellToReportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANCellToReportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCGI marshal failed")
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

func (x *NGRANCellToReportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANCellToReportItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANCellToReportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCGI = new(NGRANCGI)
	err = x.NGRANCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCGI error")
	}

	// optional field (optPresentFlag index: 0)
	if NGRANCellToReportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANCellToReportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
