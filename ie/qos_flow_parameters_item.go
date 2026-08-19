package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowParametersItem struct {
	QosFlowIdentifier         *QosFlowIdentifier
	AlternativeQoSParaSetList *AlternativeQoSParaSetList                             // optional
	IEExtensions              *ProtocolExtensionContainerQosFlowParametersItemExtIEs // optional
}

func (x *QosFlowParametersItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowParametersItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// optional field
	if x.AlternativeQoSParaSetList != nil {
		QosFlowParametersItemOptPresentFlag = append(QosFlowParametersItemOptPresentFlag, true)
	} else {
		QosFlowParametersItemOptPresentFlag = append(QosFlowParametersItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowParametersItemOptPresentFlag = append(QosFlowParametersItemOptPresentFlag, true)
	} else {
		QosFlowParametersItemOptPresentFlag = append(QosFlowParametersItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowParametersItemOptPresentFlag, true)
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
	if x.AlternativeQoSParaSetList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AlternativeQoSParaSetList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AlternativeQoSParaSetList marshal failed")
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

func (x *QosFlowParametersItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowParametersItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&QosFlowParametersItemOptPresentFlag, true)
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
	if QosFlowParametersItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AlternativeQoSParaSetList = new(AlternativeQoSParaSetList)
		err = x.AlternativeQoSParaSetList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AlternativeQoSParaSetList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowParametersItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowParametersItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
