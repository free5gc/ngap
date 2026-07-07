package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowPerTNLInformationItem struct {
	QosFlowPerTNLInformation *QosFlowPerTNLInformation                                     // valueExt
	IEExtensions             *ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs // optional
}

func (x *QosFlowPerTNLInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowPerTNLInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowPerTNLInformation == nil {
		return errors.Errorf("QosFlowPerTNLInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowPerTNLInformationItemOptPresentFlag = append(QosFlowPerTNLInformationItemOptPresentFlag, true)
	} else {
		QosFlowPerTNLInformationItemOptPresentFlag = append(QosFlowPerTNLInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowPerTNLInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowPerTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowPerTNLInformation marshal failed")
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

func (x *QosFlowPerTNLInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowPerTNLInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&QosFlowPerTNLInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowPerTNLInformation = new(QosFlowPerTNLInformation)
	err = x.QosFlowPerTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowPerTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if QosFlowPerTNLInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowPerTNLInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
