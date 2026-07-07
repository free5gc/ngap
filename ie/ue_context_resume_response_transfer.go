package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEContextResumeResponseTransfer struct {
	QosFlowFailedToResumeList *QosFlowListWithCause                                            // optional
	IEExtensions              *ProtocolExtensionContainerUEContextResumeResponseTransferExtIEs // optional
}

func (x *UEContextResumeResponseTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEContextResumeResponseTransferOptPresentFlag := []bool{}
	// optional field
	if x.QosFlowFailedToResumeList != nil {
		UEContextResumeResponseTransferOptPresentFlag = append(UEContextResumeResponseTransferOptPresentFlag, true)
	} else {
		UEContextResumeResponseTransferOptPresentFlag = append(UEContextResumeResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEContextResumeResponseTransferOptPresentFlag = append(UEContextResumeResponseTransferOptPresentFlag, true)
	} else {
		UEContextResumeResponseTransferOptPresentFlag = append(UEContextResumeResponseTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEContextResumeResponseTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.QosFlowFailedToResumeList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowFailedToResumeList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowFailedToResumeList marshal failed")
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

func (x *UEContextResumeResponseTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEContextResumeResponseTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UEContextResumeResponseTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if UEContextResumeResponseTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowFailedToResumeList = new(QosFlowListWithCause)
		err = x.QosFlowFailedToResumeList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowFailedToResumeList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UEContextResumeResponseTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEContextResumeResponseTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
