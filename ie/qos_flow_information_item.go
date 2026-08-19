package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type QosFlowInformationItem struct {
	QosFlowIdentifier *QosFlowIdentifier
	DLForwarding      *DLForwarding                                           // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions      *ProtocolExtensionContainerQosFlowInformationItemExtIEs // optional
}

func (x *QosFlowInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	QosFlowInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.QosFlowIdentifier == nil {
		return errors.Errorf("QosFlowIdentifier is missing")
	}
	// optional field
	if x.DLForwarding != nil {
		QosFlowInformationItemOptPresentFlag = append(QosFlowInformationItemOptPresentFlag, true)
	} else {
		QosFlowInformationItemOptPresentFlag = append(QosFlowInformationItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		QosFlowInformationItemOptPresentFlag = append(QosFlowInformationItemOptPresentFlag, true)
	} else {
		QosFlowInformationItemOptPresentFlag = append(QosFlowInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(QosFlowInformationItemOptPresentFlag, true)
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
	if x.DLForwarding != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLForwarding.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLForwarding marshal failed")
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

func (x *QosFlowInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	QosFlowInformationItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&QosFlowInformationItemOptPresentFlag, true)
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
	if QosFlowInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLForwarding = new(DLForwarding)
		err = x.DLForwarding.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLForwarding error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if QosFlowInformationItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerQosFlowInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
