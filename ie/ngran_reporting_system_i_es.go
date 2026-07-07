package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGRANReportingSystemIEs struct {
	NGRANCellToReportList *NGRANCellToReportList
	IEExtensions          *ProtocolExtensionContainerNGRANReportingSystemIEsExtIEs // optional
}

func (x *NGRANReportingSystemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANReportingSystemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCellToReportList == nil {
		return errors.Errorf("NGRANCellToReportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANReportingSystemIEsOptPresentFlag = append(NGRANReportingSystemIEsOptPresentFlag, true)
	} else {
		NGRANReportingSystemIEsOptPresentFlag = append(NGRANReportingSystemIEsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANReportingSystemIEsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCellToReportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCellToReportList marshal failed")
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

func (x *NGRANReportingSystemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANReportingSystemIEsOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANReportingSystemIEsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCellToReportList = new(NGRANCellToReportList)
	err = x.NGRANCellToReportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCellToReportList error")
	}

	// optional field (optPresentFlag index: 0)
	if NGRANReportingSystemIEsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANReportingSystemIEsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
