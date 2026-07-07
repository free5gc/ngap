package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	PDUSessionUsageReportRATTypePresentNr              aper.Enumerated = 0
	PDUSessionUsageReportRATTypePresentEutra           aper.Enumerated = 1
	PDUSessionUsageReportRATTypePresentNrUnlicensed    aper.Enumerated = 2
	PDUSessionUsageReportRATTypePresentEUtraUnlicensed aper.Enumerated = 3
)

type PDUSessionUsageReport struct {
	RATType                   *aper.Enumerated // valueExt,valueLB:0,valueUB:1
	PDUSessionTimedReportList *VolumeTimedReportList
	IEExtensions              *ProtocolExtensionContainerPDUSessionUsageReportExtIEs // optional
}

func (x *PDUSessionUsageReport) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionUsageReportOptPresentFlag := []bool{}
	// mandatory field
	if x.RATType == nil {
		return errors.Errorf("RATType is missing")
	}
	// mandatory field
	if x.PDUSessionTimedReportList == nil {
		return errors.Errorf("PDUSessionTimedReportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionUsageReportOptPresentFlag = append(PDUSessionUsageReportOptPresentFlag, true)
	} else {
		PDUSessionUsageReportOptPresentFlag = append(PDUSessionUsageReportOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionUsageReportOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.RATType), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PDUSessionTimedReportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PDUSessionTimedReportList marshal failed")
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

func (x *PDUSessionUsageReport) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionUsageReportOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionUsageReportOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.RATType = new(aper.Enumerated)
	*(x.RATType), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PDUSessionTimedReportList = new(VolumeTimedReportList)
	err = x.PDUSessionTimedReportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PDUSessionTimedReportList error")
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionUsageReportOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionUsageReportExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
