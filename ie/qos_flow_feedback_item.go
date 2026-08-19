package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowFeedbackItem struct {
	QosFlowIdentifier     *QosFlowIdentifier
	UpdateFeedback        *UpdateFeedback                                      // optional
	CNpacketDelayBudgetDL *ExtendedPacketDelayBudget                           // optional
	CNpacketDelayBudgetUL *ExtendedPacketDelayBudget                           // optional
	IEExtensions          *ProtocolExtensionContainerQosFlowFeedbackItemExtIEs // optional
}

func (x *QosFlowFeedbackItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowFeedbackItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// optional field
	if x.UpdateFeedback != nil {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, true)
	} else {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, false)
	}
	// optional field
	if x.CNpacketDelayBudgetDL != nil {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, true)
	} else {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, false)
	}
	// optional field
	if x.CNpacketDelayBudgetUL != nil {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, true)
	} else {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, true)
	} else {
		QosFlowFeedbackItemOptPresentFlag = append(QosFlowFeedbackItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowFeedbackItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowIdentifier marshal failed")
	}

	// optional field
	if x.UpdateFeedback != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UpdateFeedback.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UpdateFeedback marshal failed")
		}
	}

	// optional field
	if x.CNpacketDelayBudgetDL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CNpacketDelayBudgetDL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CNpacketDelayBudgetDL marshal failed")
		}
	}

	// optional field
	if x.CNpacketDelayBudgetUL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CNpacketDelayBudgetUL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CNpacketDelayBudgetUL marshal failed")
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

func (x *QosFlowFeedbackItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowFeedbackItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&QosFlowFeedbackItemOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if QosFlowFeedbackItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UpdateFeedback = new(UpdateFeedback)
		err = x.UpdateFeedback.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UpdateFeedback error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowFeedbackItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.CNpacketDelayBudgetDL = new(ExtendedPacketDelayBudget)
		err = x.CNpacketDelayBudgetDL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CNpacketDelayBudgetDL error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if QosFlowFeedbackItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.CNpacketDelayBudgetUL = new(ExtendedPacketDelayBudget)
		err = x.CNpacketDelayBudgetUL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CNpacketDelayBudgetUL error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if QosFlowFeedbackItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowFeedbackItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
