package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceModifyIndicationTransfer struct {
	DLQosFlowPerTNLInformation           *QosFlowPerTNLInformation                                                   // valueExt
	AdditionalDLQosFlowPerTNLInformation *QosFlowPerTNLInformationList                                               // optional
	IEExtensions                         *ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs // optional
}

func (x *PDUSessionResourceModifyIndicationTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceModifyIndicationTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.DLQosFlowPerTNLInformation == nil {
		return errors.Errorf("DLQosFlowPerTNLInformation is missing")
	}
	// optional field
	if x.AdditionalDLQosFlowPerTNLInformation != nil {
		PDUSessionResourceModifyIndicationTransferOptPresentFlag = append(PDUSessionResourceModifyIndicationTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyIndicationTransferOptPresentFlag = append(PDUSessionResourceModifyIndicationTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceModifyIndicationTransferOptPresentFlag = append(PDUSessionResourceModifyIndicationTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyIndicationTransferOptPresentFlag = append(PDUSessionResourceModifyIndicationTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceModifyIndicationTransferOptPresentFlag, true)
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
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *PDUSessionResourceModifyIndicationTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceModifyIndicationTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceModifyIndicationTransferOptPresentFlag, true)
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
	if PDUSessionResourceModifyIndicationTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalDLQosFlowPerTNLInformation = new(QosFlowPerTNLInformationList)
		err = x.AdditionalDLQosFlowPerTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalDLQosFlowPerTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceModifyIndicationTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
