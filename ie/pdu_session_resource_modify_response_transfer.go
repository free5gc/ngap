package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation                *UPTransportLayerInformation                                              // valueLB:0,valueUB:1,optional
	ULNGUUPTNLInformation                *UPTransportLayerInformation                                              // valueLB:0,valueUB:1,optional
	QosFlowAddOrModifyResponseList       *QosFlowAddOrModifyResponseList                                           // optional
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList                                             // optional
	QosFlowFailedToAddOrModifyList       *QosFlowListWithCause                                                     // optional
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs // optional
}

func (x *PDUSessionResourceModifyResponseTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceModifyResponseTransferOptPresentFlag := []bool{}
	// optional field
	if x.DLNGUUPTNLInformation != nil {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.ULNGUUPTNLInformation != nil {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowAddOrModifyResponseList != nil {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.AdditionalDLQosFlowPerTNLInformation != nil {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowFailedToAddOrModifyList != nil {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyResponseTransferOptPresentFlag = append(PDUSessionResourceModifyResponseTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.DLNGUUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLNGUUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLNGUUPTNLInformation marshal failed")
		}
	}

	// optional field
	if x.ULNGUUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ULNGUUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ULNGUUPTNLInformation marshal failed")
		}
	}

	// optional field
	if x.QosFlowAddOrModifyResponseList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowAddOrModifyResponseList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowAddOrModifyResponseList marshal failed")
		}
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
	if x.QosFlowFailedToAddOrModifyList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowFailedToAddOrModifyList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowFailedToAddOrModifyList marshal failed")
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

func (x *PDUSessionResourceModifyResponseTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceModifyResponseTransferOptPresentFlag := make([]bool, 6)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceModifyResponseTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceModifyResponseTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
		err = x.DLNGUUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLNGUUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceModifyResponseTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
		err = x.ULNGUUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ULNGUUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PDUSessionResourceModifyResponseTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowAddOrModifyResponseList = new(QosFlowAddOrModifyResponseList)
		err = x.QosFlowAddOrModifyResponseList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowAddOrModifyResponseList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if PDUSessionResourceModifyResponseTransferOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalDLQosFlowPerTNLInformation = new(QosFlowPerTNLInformationList)
		err = x.AdditionalDLQosFlowPerTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalDLQosFlowPerTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if PDUSessionResourceModifyResponseTransferOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowFailedToAddOrModifyList = new(QosFlowListWithCause)
		err = x.QosFlowFailedToAddOrModifyList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowFailedToAddOrModifyList error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if PDUSessionResourceModifyResponseTransferOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
