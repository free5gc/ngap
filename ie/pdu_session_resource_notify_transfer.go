package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList                                                // optional
	QosFlowReleasedList *QosFlowListWithCause                                             // optional
	IEExtensions        *ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs // optional
}

func (x *PDUSessionResourceNotifyTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceNotifyTransferOptPresentFlag := []bool{}
	// optional field
	if x.QosFlowNotifyList != nil {
		PDUSessionResourceNotifyTransferOptPresentFlag = append(PDUSessionResourceNotifyTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceNotifyTransferOptPresentFlag = append(PDUSessionResourceNotifyTransferOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowReleasedList != nil {
		PDUSessionResourceNotifyTransferOptPresentFlag = append(PDUSessionResourceNotifyTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceNotifyTransferOptPresentFlag = append(PDUSessionResourceNotifyTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceNotifyTransferOptPresentFlag = append(PDUSessionResourceNotifyTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceNotifyTransferOptPresentFlag = append(PDUSessionResourceNotifyTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceNotifyTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.QosFlowNotifyList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowNotifyList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowNotifyList marshal failed")
		}
	}

	// optional field
	if x.QosFlowReleasedList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowReleasedList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowReleasedList marshal failed")
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

func (x *PDUSessionResourceNotifyTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceNotifyTransferOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceNotifyTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceNotifyTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowNotifyList = new(QosFlowNotifyList)
		err = x.QosFlowNotifyList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowNotifyList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceNotifyTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowReleasedList = new(QosFlowListWithCause)
		err = x.QosFlowReleasedList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowReleasedList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PDUSessionResourceNotifyTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
