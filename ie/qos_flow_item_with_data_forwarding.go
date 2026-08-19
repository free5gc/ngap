package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowItemWithDataForwarding struct {
	QosFlowIdentifier      *QosFlowIdentifier
	DataForwardingAccepted *DataForwardingAccepted                                        // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions           *ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs // optional
}

func (x *QosFlowItemWithDataForwarding) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowItemWithDataForwardingOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// optional field
	if x.DataForwardingAccepted != nil {
		QosFlowItemWithDataForwardingOptPresentFlag = append(QosFlowItemWithDataForwardingOptPresentFlag, true)
	} else {
		QosFlowItemWithDataForwardingOptPresentFlag = append(QosFlowItemWithDataForwardingOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowItemWithDataForwardingOptPresentFlag = append(QosFlowItemWithDataForwardingOptPresentFlag, true)
	} else {
		QosFlowItemWithDataForwardingOptPresentFlag = append(QosFlowItemWithDataForwardingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowItemWithDataForwardingOptPresentFlag, true)
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
	if x.DataForwardingAccepted != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DataForwardingAccepted.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DataForwardingAccepted marshal failed")
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

func (x *QosFlowItemWithDataForwarding) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowItemWithDataForwardingOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&QosFlowItemWithDataForwardingOptPresentFlag, true)
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
	if QosFlowItemWithDataForwardingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DataForwardingAccepted = new(DataForwardingAccepted)
		err = x.DataForwardingAccepted.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DataForwardingAccepted error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowItemWithDataForwardingOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowItemWithDataForwardingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
