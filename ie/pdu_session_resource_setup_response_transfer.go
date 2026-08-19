package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceSetupResponseTransfer struct {
	DLQosFlowPerTNLInformation           *QosFlowPerTNLInformation                                                // valueExt
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList                                            // optional
	SecurityResult                       *SecurityResult                                                          // valueExt,optional
	QosFlowFailedToSetupList             *QosFlowListWithCause                                                    // optional
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs // optional
}

func (x *PDUSessionResourceSetupResponseTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceSetupResponseTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.DLQosFlowPerTNLInformation == nil {
		return errors.Errorf("DLQosFlowPerTNLInformation is missing")
	}
	// optional field
	if x.AdditionalDLQosFlowPerTNLInformation != nil {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.SecurityResult != nil {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowFailedToSetupList != nil {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupResponseTransferOptPresentFlag = append(PDUSessionResourceSetupResponseTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceSetupResponseTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DLQosFlowPerTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLQosFlowPerTNLInformation marshal failed")
	}

	// optional field
	if x.AdditionalDLQosFlowPerTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AdditionalDLQosFlowPerTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AdditionalDLQosFlowPerTNLInformation marshal failed")
		}
	}

	// optional field
	if x.SecurityResult != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SecurityResult.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SecurityResult marshal failed")
		}
	}

	// optional field
	if x.QosFlowFailedToSetupList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowFailedToSetupList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowFailedToSetupList marshal failed")
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

func (x *PDUSessionResourceSetupResponseTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceSetupResponseTransferOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceSetupResponseTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLQosFlowPerTNLInformation = new(QosFlowPerTNLInformation)
	err = x.DLQosFlowPerTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLQosFlowPerTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceSetupResponseTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalDLQosFlowPerTNLInformation = new(QosFlowPerTNLInformationList)
		err = x.AdditionalDLQosFlowPerTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalDLQosFlowPerTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceSetupResponseTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SecurityResult = new(SecurityResult)
		err = x.SecurityResult.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SecurityResult error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PDUSessionResourceSetupResponseTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowFailedToSetupList = new(QosFlowListWithCause)
		err = x.QosFlowFailedToSetupList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowFailedToSetupList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if PDUSessionResourceSetupResponseTransferOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceSetupResponseTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
