package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGRANReportingStatusIEs struct {
	NGRANCellReportList *NGRANCellReportList
	IEExtensions        *ProtocolExtensionContainerNGRANReportingStatusIEsExtIEs // optional
}

func (x *NGRANReportingStatusIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANReportingStatusIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCellReportList == nil {
		return errors.Errorf("NGRANCellReportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANReportingStatusIEsOptPresentFlag = append(NGRANReportingStatusIEsOptPresentFlag, true)
	} else {
		NGRANReportingStatusIEsOptPresentFlag = append(NGRANReportingStatusIEsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANReportingStatusIEsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCellReportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCellReportList marshal failed")
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

func (x *NGRANReportingStatusIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANReportingStatusIEsOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANReportingStatusIEsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCellReportList = new(NGRANCellReportList)
	err = x.NGRANCellReportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCellReportList error")
	}

	// optional field (optPresentFlag index: 0)
	if NGRANReportingStatusIEsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANReportingStatusIEsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
