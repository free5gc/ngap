package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	QoSFlowsUsageReportItemRATTypePresentNr              aper.Enumerated = 0
	QoSFlowsUsageReportItemRATTypePresentEutra           aper.Enumerated = 1
	QoSFlowsUsageReportItemRATTypePresentNrUnlicensed    aper.Enumerated = 2
	QoSFlowsUsageReportItemRATTypePresentEUtraUnlicensed aper.Enumerated = 3
)

type QoSFlowsUsageReportItem struct {
	QosFlowIdentifier       *QosFlowIdentifier
	RATType                 *aper.Enumerated // valueExt,valueLB:0,valueUB:1
	QoSFlowsTimedReportList *VolumeTimedReportList
	IEExtensions            *ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs // optional
}

func (x *QoSFlowsUsageReportItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QoSFlowsUsageReportItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// mandatory field
	if x.RATType == nil {
		return errors.Errorf("RATType is missing")
	}
	// mandatory field
	if x.QoSFlowsTimedReportList == nil {
		return errors.Errorf("QoSFlowsTimedReportList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		QoSFlowsUsageReportItemOptPresentFlag = append(QoSFlowsUsageReportItemOptPresentFlag, true)
	} else {
		QoSFlowsUsageReportItemOptPresentFlag = append(QoSFlowsUsageReportItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QoSFlowsUsageReportItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowIdentifier marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.RATType), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.QoSFlowsTimedReportList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QoSFlowsTimedReportList marshal failed")
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

func (x *QoSFlowsUsageReportItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QoSFlowsUsageReportItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&QoSFlowsUsageReportItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowIdentifier = new(QosFlowIdentifier)
	err = x.QosFlowIdentifier.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowIdentifier error")
	}

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
	x.QoSFlowsTimedReportList = new(VolumeTimedReportList)
	err = x.QoSFlowsTimedReportList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QoSFlowsTimedReportList error")
	}

	// optional field (optPresentFlag index: 0)
	if QoSFlowsUsageReportItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQoSFlowsUsageReportItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
