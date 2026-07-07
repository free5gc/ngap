package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceModifyConfirmTransfer struct {
	QosFlowModifyConfirmList      *QosFlowModifyConfirmList
	ULNGUUPTNLInformation         *UPTransportLayerInformation                                             // valueLB:0,valueUB:1
	AdditionalNGUUPTNLInformation *UPTransportLayerInformationPairList                                     // optional
	QosFlowFailedToModifyList     *QosFlowListWithCause                                                    // optional
	IEExtensions                  *ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs // optional
}

func (x *PDUSessionResourceModifyConfirmTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceModifyConfirmTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowModifyConfirmList == nil {
		return errors.Errorf("QosFlowModifyConfirmList is missing")
	}
	// mandatory field
	if x.ULNGUUPTNLInformation == nil {
		return errors.Errorf("ULNGUUPTNLInformation is missing")
	}
	// optional field
	if x.AdditionalNGUUPTNLInformation != nil {
		PDUSessionResourceModifyConfirmTransferOptPresentFlag = append(PDUSessionResourceModifyConfirmTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyConfirmTransferOptPresentFlag = append(PDUSessionResourceModifyConfirmTransferOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowFailedToModifyList != nil {
		PDUSessionResourceModifyConfirmTransferOptPresentFlag = append(PDUSessionResourceModifyConfirmTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyConfirmTransferOptPresentFlag = append(PDUSessionResourceModifyConfirmTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceModifyConfirmTransferOptPresentFlag = append(PDUSessionResourceModifyConfirmTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyConfirmTransferOptPresentFlag = append(PDUSessionResourceModifyConfirmTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceModifyConfirmTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowModifyConfirmList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowModifyConfirmList marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ULNGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ULNGUUPTNLInformation marshal failed")
	}

	// optional field
	if x.AdditionalNGUUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AdditionalNGUUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AdditionalNGUUPTNLInformation marshal failed")
		}
	}

	// optional field
	if x.QosFlowFailedToModifyList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowFailedToModifyList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowFailedToModifyList marshal failed")
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

func (x *PDUSessionResourceModifyConfirmTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceModifyConfirmTransferOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceModifyConfirmTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowModifyConfirmList = new(QosFlowModifyConfirmList)
	err = x.QosFlowModifyConfirmList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowModifyConfirmList error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.ULNGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ULNGUUPTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceModifyConfirmTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalNGUUPTNLInformation = new(UPTransportLayerInformationPairList)
		err = x.AdditionalNGUUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalNGUUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceModifyConfirmTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowFailedToModifyList = new(QosFlowListWithCause)
		err = x.QosFlowFailedToModifyList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowFailedToModifyList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PDUSessionResourceModifyConfirmTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceModifyConfirmTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
