package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowAddOrModifyRequestItem struct {
	QosFlowIdentifier         *QosFlowIdentifier
	QosFlowLevelQosParameters *QosFlowLevelQosParameters                                     // valueExt,optional
	ERABID                    *ERABID                                                        // optional
	IEExtensions              *ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs // optional
}

func (x *QosFlowAddOrModifyRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowAddOrModifyRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// optional field
	if x.QosFlowLevelQosParameters != nil {
		QosFlowAddOrModifyRequestItemOptPresentFlag = append(QosFlowAddOrModifyRequestItemOptPresentFlag, true)
	} else {
		QosFlowAddOrModifyRequestItemOptPresentFlag = append(QosFlowAddOrModifyRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.ERABID != nil {
		QosFlowAddOrModifyRequestItemOptPresentFlag = append(QosFlowAddOrModifyRequestItemOptPresentFlag, true)
	} else {
		QosFlowAddOrModifyRequestItemOptPresentFlag = append(QosFlowAddOrModifyRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowAddOrModifyRequestItemOptPresentFlag = append(QosFlowAddOrModifyRequestItemOptPresentFlag, true)
	} else {
		QosFlowAddOrModifyRequestItemOptPresentFlag = append(QosFlowAddOrModifyRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowAddOrModifyRequestItemOptPresentFlag, true)
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
	if x.QosFlowLevelQosParameters != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QosFlowLevelQosParameters.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QosFlowLevelQosParameters marshal failed")
		}
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

func (x *QosFlowAddOrModifyRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowAddOrModifyRequestItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&QosFlowAddOrModifyRequestItemOptPresentFlag, true)
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
	if QosFlowAddOrModifyRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QosFlowLevelQosParameters = new(QosFlowLevelQosParameters)
		err = x.QosFlowLevelQosParameters.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QosFlowLevelQosParameters error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowAddOrModifyRequestItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ERABID = new(ERABID)
		err = x.ERABID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ERABID error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if QosFlowAddOrModifyRequestItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowAddOrModifyRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
