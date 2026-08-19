package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type HandoverCommandTransfer struct {
	DLForwardingUPTNLInformation  *UPTransportLayerInformation                             // valueLB:0,valueUB:1,optional
	QosFlowToBeForwardedList      *QosFlowToBeForwardedList                                // optional
	DataForwardingResponseDRBList *DataForwardingResponseDRBList                           // optional
	IEExtensions                  *ProtocolExtensionContainerHandoverCommandTransferExtIEs // optional
}

func (x *HandoverCommandTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	HandoverCommandTransferOptPresentFlag := []bool{}
	// optional field
	if x.DLForwardingUPTNLInformation != nil {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, true)
	} else {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, false)
	}
	// optional field
	if x.QosFlowToBeForwardedList != nil {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, true)
	} else {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, false)
	}
	// optional field
	if x.DataForwardingResponseDRBList != nil {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, true)
	} else {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, true)
	} else {
		HandoverCommandTransferOptPresentFlag = append(HandoverCommandTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(HandoverCommandTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.DLForwardingUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLForwardingUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLForwardingUPTNLInformation marshal failed")
		}
	}

	// optional field
	if x.QosFlowToBeForwardedList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowToBeForwardedList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowToBeForwardedList marshal failed")
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

func (x *HandoverCommandTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	HandoverCommandTransferOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&HandoverCommandTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if HandoverCommandTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLForwardingUPTNLInformation = new(UPTransportLayerInformation)
		err = x.DLForwardingUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLForwardingUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if HandoverCommandTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowToBeForwardedList = new(QosFlowToBeForwardedList)
		err = x.QosFlowToBeForwardedList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowToBeForwardedList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if HandoverCommandTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.DataForwardingResponseDRBList = new(DataForwardingResponseDRBList)
		err = x.DataForwardingResponseDRBList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DataForwardingResponseDRBList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if HandoverCommandTransferOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerHandoverCommandTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
