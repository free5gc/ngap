package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSActiveSessionInformationSourcetoTargetItem struct {
	MBSSessionID                           *MBSSessionID     // valueExt
	MBSAreaSessionID                       *MBSAreaSessionID // optional
	MBSServiceArea                         *MBSServiceArea   // valueLB:0,valueUB:2,optional
	MBSQoSFlowsToBeSetupList               *MBSQoSFlowsToBeSetupList
	MBSMappingandDataForwardingRequestList *MBSMappingandDataForwardingRequestList                                        // optional
	IEExtensions                           *ProtocolExtensionContainerMBSActiveSessionInformationSourcetoTargetItemExtIEs // optional
}

func (x *MBSActiveSessionInformationSourcetoTargetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSSessionID == nil {
		return errors.Errorf("MBSSessionID is missing")
	}
	// optional field
	if x.MBSAreaSessionID != nil {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, true)
	} else {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, false)
	}
	// optional field
	if x.MBSServiceArea != nil {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, true)
	} else {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, false)
	}
	// mandatory field
	if x.MBSQoSFlowsToBeSetupList == nil {
		return errors.Errorf("MBSQoSFlowsToBeSetupList is missing")
	}
	// optional field
	if x.MBSMappingandDataForwardingRequestList != nil {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, true)
	} else {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, true)
	} else {
		MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag = append(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSSessionID marshal failed")
	}

	// optional field
	if x.MBSAreaSessionID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSAreaSessionID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSAreaSessionID marshal failed")
		}
	}

	// optional field
	if x.MBSServiceArea != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSServiceArea.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSServiceArea marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MBSQoSFlowsToBeSetupList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSQoSFlowsToBeSetupList marshal failed")
	}

	// optional field
	if x.MBSMappingandDataForwardingRequestList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSMappingandDataForwardingRequestList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSMappingandDataForwardingRequestList marshal failed")
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

func (x *MBSActiveSessionInformationSourcetoTargetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSSessionID = new(MBSSessionID)
	err = x.MBSSessionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSSessionID error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MBSAreaSessionID = new(MBSAreaSessionID)
		err = x.MBSAreaSessionID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSAreaSessionID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MBSServiceArea = new(MBSServiceArea)
		err = x.MBSServiceArea.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSServiceArea error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSQoSFlowsToBeSetupList = new(MBSQoSFlowsToBeSetupList)
	err = x.MBSQoSFlowsToBeSetupList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSQoSFlowsToBeSetupList error")
	}

	// optional field (optPresentFlag index: 2)
	if MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.MBSMappingandDataForwardingRequestList = new(MBSMappingandDataForwardingRequestList)
		err = x.MBSMappingandDataForwardingRequestList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSMappingandDataForwardingRequestList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if MBSActiveSessionInformationSourcetoTargetItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSActiveSessionInformationSourcetoTargetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
