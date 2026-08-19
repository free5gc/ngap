package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowAddOrModifyResponseItem struct {
	QosFlowIdentifier *QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs // optional
}

func (x *QosFlowAddOrModifyResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowAddOrModifyResponseItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowAddOrModifyResponseItemOptPresentFlag = append(QosFlowAddOrModifyResponseItemOptPresentFlag, true)
	} else {
		QosFlowAddOrModifyResponseItemOptPresentFlag = append(QosFlowAddOrModifyResponseItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowAddOrModifyResponseItemOptPresentFlag, true)
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
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *QosFlowAddOrModifyResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowAddOrModifyResponseItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&QosFlowAddOrModifyResponseItemOptPresentFlag, true)
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
	if QosFlowAddOrModifyResponseItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowAddOrModifyResponseItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
