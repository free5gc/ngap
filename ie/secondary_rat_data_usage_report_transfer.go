package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SecondaryRATDataUsageReportTransfer struct {
	SecondaryRATUsageInformation *SecondaryRATUsageInformation                                        // valueExt,optional
	IEExtensions                 *ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs // optional
}

func (x *SecondaryRATDataUsageReportTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SecondaryRATDataUsageReportTransferOptPresentFlag := []bool{}
	// optional field
	if x.SecondaryRATUsageInformation != nil {
		SecondaryRATDataUsageReportTransferOptPresentFlag = append(SecondaryRATDataUsageReportTransferOptPresentFlag, true)
	} else {
		SecondaryRATDataUsageReportTransferOptPresentFlag = append(SecondaryRATDataUsageReportTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SecondaryRATDataUsageReportTransferOptPresentFlag = append(SecondaryRATDataUsageReportTransferOptPresentFlag, true)
	} else {
		SecondaryRATDataUsageReportTransferOptPresentFlag = append(SecondaryRATDataUsageReportTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SecondaryRATDataUsageReportTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.SecondaryRATUsageInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SecondaryRATUsageInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SecondaryRATUsageInformation marshal failed")
		}
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

func (x *SecondaryRATDataUsageReportTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SecondaryRATDataUsageReportTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SecondaryRATDataUsageReportTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if SecondaryRATDataUsageReportTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SecondaryRATUsageInformation = new(SecondaryRATUsageInformation)
		err = x.SecondaryRATUsageInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SecondaryRATUsageInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SecondaryRATDataUsageReportTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSecondaryRATDataUsageReportTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
