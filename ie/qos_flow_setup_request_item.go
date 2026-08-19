package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowSetupRequestItem struct {
	QosFlowIdentifier         *QosFlowIdentifier
	QosFlowLevelQosParameters *QosFlowLevelQosParameters                               // valueExt
	ERABID                    *ERABID                                                  // optional
	IEExtensions              *ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs // optional
}

func (x *QosFlowSetupRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowSetupRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// mandatory field
	if x.QosFlowLevelQosParameters == nil {
		return errors.Errorf("QosFlowLevelQosParameters is missing")
	}
	// optional field
	if x.ERABID != nil {
		QosFlowSetupRequestItemOptPresentFlag = append(QosFlowSetupRequestItemOptPresentFlag, true)
	} else {
		QosFlowSetupRequestItemOptPresentFlag = append(QosFlowSetupRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowSetupRequestItemOptPresentFlag = append(QosFlowSetupRequestItemOptPresentFlag, true)
	} else {
		QosFlowSetupRequestItemOptPresentFlag = append(QosFlowSetupRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowSetupRequestItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowIdentifier marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowLevelQosParameters.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowLevelQosParameters marshal failed")
	}

	// optional field
	if x.ERABID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ERABID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ERABID marshal failed")
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

func (x *QosFlowSetupRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowSetupRequestItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&QosFlowSetupRequestItemOptPresentFlag, true)
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

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowLevelQosParameters = new(QosFlowLevelQosParameters)
	err = x.QosFlowLevelQosParameters.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowLevelQosParameters error")
	}

	// optional field (optPresentFlag index: 0)
	if QosFlowSetupRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ERABID = new(ERABID)
		err = x.ERABID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ERABID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowSetupRequestItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowSetupRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
