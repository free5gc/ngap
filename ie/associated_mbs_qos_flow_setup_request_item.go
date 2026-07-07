package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AssociatedMBSQosFlowSetupRequestItem struct {
	MBSQosFlowIdentifier               *QosFlowIdentifier
	AssociatedUnicastQosFlowIdentifier *QosFlowIdentifier
	IEExtensions                       *ProtocolExtensionContainerAssociatedMBSQosFlowSetupRequestItemExtIEs // optional
}

func (x *AssociatedMBSQosFlowSetupRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssociatedMBSQosFlowSetupRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSQosFlowIdentifier == nil {
		return errors.Errorf("MBSQosFlowIdentifier is missing")
	}
	// mandatory field
	if x.AssociatedUnicastQosFlowIdentifier == nil {
		return errors.Errorf("AssociatedUnicastQosFlowIdentifier is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AssociatedMBSQosFlowSetupRequestItemOptPresentFlag = append(AssociatedMBSQosFlowSetupRequestItemOptPresentFlag, true)
	} else {
		AssociatedMBSQosFlowSetupRequestItemOptPresentFlag = append(AssociatedMBSQosFlowSetupRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AssociatedMBSQosFlowSetupRequestItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSQosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSQosFlowIdentifier marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AssociatedUnicastQosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AssociatedUnicastQosFlowIdentifier marshal failed")
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

func (x *AssociatedMBSQosFlowSetupRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssociatedMBSQosFlowSetupRequestItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AssociatedMBSQosFlowSetupRequestItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSQosFlowIdentifier = new(QosFlowIdentifier)
	err = x.MBSQosFlowIdentifier.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSQosFlowIdentifier error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AssociatedUnicastQosFlowIdentifier = new(QosFlowIdentifier)
	err = x.AssociatedUnicastQosFlowIdentifier.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AssociatedUnicastQosFlowIdentifier error")
	}

	// optional field (optPresentFlag index: 0)
	if AssociatedMBSQosFlowSetupRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAssociatedMBSQosFlowSetupRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
