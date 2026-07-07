package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEContextResumeRequestTransfer struct {
	QosFlowFailedToResumeList *QosFlowListWithCause                                           // optional
	IEExtensions              *ProtocolExtensionContainerUEContextResumeRequestTransferExtIEs // optional
}

func (x *UEContextResumeRequestTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEContextResumeRequestTransferOptPresentFlag := []bool{}
	// optional field
	if x.QosFlowFailedToResumeList != nil {
		UEContextResumeRequestTransferOptPresentFlag = append(UEContextResumeRequestTransferOptPresentFlag, true)
	} else {
		UEContextResumeRequestTransferOptPresentFlag = append(UEContextResumeRequestTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEContextResumeRequestTransferOptPresentFlag = append(UEContextResumeRequestTransferOptPresentFlag, true)
	} else {
		UEContextResumeRequestTransferOptPresentFlag = append(UEContextResumeRequestTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEContextResumeRequestTransferOptPresentFlag, true)
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

func (x *UEContextResumeRequestTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEContextResumeRequestTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UEContextResumeRequestTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if UEContextResumeRequestTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowFailedToResumeList = new(QosFlowListWithCause)
		err = x.QosFlowFailedToResumeList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowFailedToResumeList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UEContextResumeRequestTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEContextResumeRequestTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
