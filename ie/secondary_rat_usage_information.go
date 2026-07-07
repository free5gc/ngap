package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SecondaryRATUsageInformation struct {
	PDUSessionUsageReport   *PDUSessionUsageReport                                        // valueExt,optional
	QosFlowsUsageReportList *QoSFlowsUsageReportList                                      // optional
	IEExtension             *ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs // optional
}

func (x *SecondaryRATUsageInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SecondaryRATUsageInformationOptPresentFlag := []bool{}
	// optional field
	if x.PDUSessionUsageReport != nil {
		SecondaryRATUsageInformationOptPresentFlag = append(SecondaryRATUsageInformationOptPresentFlag, true)
	} else {
		SecondaryRATUsageInformationOptPresentFlag = append(SecondaryRATUsageInformationOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowsUsageReportList != nil {
		SecondaryRATUsageInformationOptPresentFlag = append(SecondaryRATUsageInformationOptPresentFlag, true)
	} else {
		SecondaryRATUsageInformationOptPresentFlag = append(SecondaryRATUsageInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtension != nil {
		SecondaryRATUsageInformationOptPresentFlag = append(SecondaryRATUsageInformationOptPresentFlag, true)
	} else {
		SecondaryRATUsageInformationOptPresentFlag = append(SecondaryRATUsageInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SecondaryRATUsageInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PDUSessionUsageReport != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PDUSessionUsageReport.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PDUSessionUsageReport marshal failed")
		}
	}

	// optional field
	if x.QosFlowsUsageReportList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowsUsageReportList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowsUsageReportList marshal failed")
		}
	}

	// optional field
	if x.IEExtension != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtension.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtension marshal failed")
		}
	}

	return nil
}

func (x *SecondaryRATUsageInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SecondaryRATUsageInformationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&SecondaryRATUsageInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if SecondaryRATUsageInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PDUSessionUsageReport = new(PDUSessionUsageReport)
		err = x.PDUSessionUsageReport.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PDUSessionUsageReport error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SecondaryRATUsageInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowsUsageReportList = new(QoSFlowsUsageReportList)
		err = x.QosFlowsUsageReportList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowsUsageReportList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if SecondaryRATUsageInformationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerSecondaryRATUsageInformationExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
