package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowNotifyItem struct {
	QosFlowIdentifier *QosFlowIdentifier
	NotificationCause *NotificationCause                                 // valueExt,valueLB:0,valueUB:1
	IEExtensions      *ProtocolExtensionContainerQosFlowNotifyItemExtIEs // optional
}

func (x *QosFlowNotifyItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowNotifyItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// mandatory field
	if x.NotificationCause == nil {
		return errors.Errorf("NotificationCause is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowNotifyItemOptPresentFlag = append(QosFlowNotifyItemOptPresentFlag, true)
	} else {
		QosFlowNotifyItemOptPresentFlag = append(QosFlowNotifyItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowNotifyItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowIdentifier marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NotificationCause.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NotificationCause marshal failed")
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

func (x *QosFlowNotifyItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowNotifyItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&QosFlowNotifyItemOptPresentFlag, true)
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
	// Read struct defined elsewhere (Pointer)
	x.NotificationCause = new(NotificationCause)
	err = x.NotificationCause.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NotificationCause error")
	}

	// optional field (optPresentFlag index: 0)
	if QosFlowNotifyItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowNotifyItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
