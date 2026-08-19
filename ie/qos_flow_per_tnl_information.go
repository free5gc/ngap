package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowPerTNLInformation struct {
	UPTransportLayerInformation *UPTransportLayerInformation // valueLB:0,valueUB:1
	AssociatedQosFlowList       *AssociatedQosFlowList
	IEExtensions                *ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs // optional
}

func (x *QosFlowPerTNLInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowPerTNLInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.UPTransportLayerInformation == nil {
		return errors.Errorf("UPTransportLayerInformation is missing")
	}
	// mandatory field
	if x.AssociatedQosFlowList == nil {
		return errors.Errorf("AssociatedQosFlowList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowPerTNLInformationOptPresentFlag = append(QosFlowPerTNLInformationOptPresentFlag, true)
	} else {
		QosFlowPerTNLInformationOptPresentFlag = append(QosFlowPerTNLInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowPerTNLInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.UPTransportLayerInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UPTransportLayerInformation marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AssociatedQosFlowList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AssociatedQosFlowList marshal failed")
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

func (x *QosFlowPerTNLInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowPerTNLInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&QosFlowPerTNLInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UPTransportLayerInformation = new(UPTransportLayerInformation)
	err = x.UPTransportLayerInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UPTransportLayerInformation error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AssociatedQosFlowList = new(AssociatedQosFlowList)
	err = x.AssociatedQosFlowList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AssociatedQosFlowList error")
	}

	// optional field (optPresentFlag index: 0)
	if QosFlowPerTNLInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowPerTNLInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
