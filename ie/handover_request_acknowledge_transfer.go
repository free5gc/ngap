package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type HandoverRequestAcknowledgeTransfer struct {
	DLNGUUPTNLInformation         *UPTransportLayerInformation // valueLB:0,valueUB:1
	DLForwardingUPTNLInformation  *UPTransportLayerInformation // valueLB:0,valueUB:1,optional
	SecurityResult                *SecurityResult              // valueExt,optional
	QosFlowSetupResponseList      *QosFlowListWithDataForwarding
	QosFlowFailedToSetupList      *QosFlowListWithCause                                               // optional
	DataForwardingResponseDRBList *DataForwardingResponseDRBList                                      // optional
	IEExtensions                  *ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs // optional
}

func (x *HandoverRequestAcknowledgeTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	HandoverRequestAcknowledgeTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.DLNGUUPTNLInformation == nil {
		return errors.Errorf("DLNGUUPTNLInformation is missing")
	}
	// optional field
	if x.DLForwardingUPTNLInformation != nil {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, false)
	}
	// optional field
	if x.SecurityResult != nil {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, false)
	}
	// mandatory field
	if x.QosFlowSetupResponseList == nil {
		return errors.Errorf("QosFlowSetupResponseList is missing")
	}
	// optional field
	if x.QosFlowFailedToSetupList != nil {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, false)
	}
	// optional field
	if x.DataForwardingResponseDRBList != nil {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		HandoverRequestAcknowledgeTransferOptPresentFlag = append(HandoverRequestAcknowledgeTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DLNGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLNGUUPTNLInformation marshal failed")
	}

	// optional field
	if x.DLForwardingUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLForwardingUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLForwardingUPTNLInformation marshal failed")
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

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowSetupResponseList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowSetupResponseList marshal failed")
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
	if x.DataForwardingResponseDRBList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DataForwardingResponseDRBList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DataForwardingResponseDRBList marshal failed")
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

func (x *HandoverRequestAcknowledgeTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	HandoverRequestAcknowledgeTransferOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&HandoverRequestAcknowledgeTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.DLNGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLNGUUPTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if HandoverRequestAcknowledgeTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLForwardingUPTNLInformation = new(UPTransportLayerInformation)
		err = x.DLForwardingUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLForwardingUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if HandoverRequestAcknowledgeTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SecurityResult = new(SecurityResult)
		err = x.SecurityResult.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SecurityResult error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowSetupResponseList = new(QosFlowListWithDataForwarding)
	err = x.QosFlowSetupResponseList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowSetupResponseList error")
	}

	// optional field (optPresentFlag index: 2)
	if HandoverRequestAcknowledgeTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowFailedToSetupList = new(QosFlowListWithCause)
		err = x.QosFlowFailedToSetupList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowFailedToSetupList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if HandoverRequestAcknowledgeTransferOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.DataForwardingResponseDRBList = new(DataForwardingResponseDRBList)
		err = x.DataForwardingResponseDRBList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DataForwardingResponseDRBList error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if HandoverRequestAcknowledgeTransferOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerHandoverRequestAcknowledgeTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
