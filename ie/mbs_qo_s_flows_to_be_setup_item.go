package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSQoSFlowsToBeSetupItem struct {
	MBSqosFlowIdentifier         *QosFlowIdentifier
	MBSqosFlowLevelQosParameters *QosFlowLevelQosParameters                                // valueExt
	IEExtensions                 *ProtocolExtensionContainerMBSQoSFlowsToBeSetupItemExtIEs // optional
}

func (x *MBSQoSFlowsToBeSetupItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSQoSFlowsToBeSetupItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSqosFlowIdentifier == nil {
		return errors.Errorf("MBSqosFlowIdentifier is missing")
	}
	// mandatory field
	if x.MBSqosFlowLevelQosParameters == nil {
		return errors.Errorf("MBSqosFlowLevelQosParameters is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		MBSQoSFlowsToBeSetupItemOptPresentFlag = append(MBSQoSFlowsToBeSetupItemOptPresentFlag, true)
	} else {
		MBSQoSFlowsToBeSetupItemOptPresentFlag = append(MBSQoSFlowsToBeSetupItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSQoSFlowsToBeSetupItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSqosFlowIdentifier.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSqosFlowIdentifier marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MBSqosFlowLevelQosParameters.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSqosFlowLevelQosParameters marshal failed")
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

func (x *MBSQoSFlowsToBeSetupItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSQoSFlowsToBeSetupItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&MBSQoSFlowsToBeSetupItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSqosFlowIdentifier = new(QosFlowIdentifier)
	err = x.MBSqosFlowIdentifier.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSqosFlowIdentifier error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSqosFlowLevelQosParameters = new(QosFlowLevelQosParameters)
	err = x.MBSqosFlowLevelQosParameters.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSqosFlowLevelQosParameters error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSQoSFlowsToBeSetupItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSQoSFlowsToBeSetupItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
