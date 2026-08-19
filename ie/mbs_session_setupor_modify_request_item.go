package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSSessionSetuporModifyRequestItem struct {
	MBSSessionID                                 *MBSSessionID                                                       // valueExt
	MBSAreaSessionID                             *MBSAreaSessionID                                                   // optional
	AssociatedMBSQosFlowSetuporModifyRequestList *AssociatedMBSQosFlowSetuporModifyRequestList                       // optional
	MBSQosFlowToReleaseList                      *QosFlowListWithCause                                               // optional
	IEExtensions                                 *ProtocolExtensionContainerMBSSessionSetuporModifyRequestItemExtIEs // optional
}

func (x *MBSSessionSetuporModifyRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSSessionSetuporModifyRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSSessionID == nil {
		return errors.Errorf("MBSSessionID is missing")
	}
	// optional field
	if x.MBSAreaSessionID != nil {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, true)
	} else {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.AssociatedMBSQosFlowSetuporModifyRequestList != nil {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, true)
	} else {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.MBSQosFlowToReleaseList != nil {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, true)
	} else {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, true)
	} else {
		MBSSessionSetuporModifyRequestItemOptPresentFlag = append(MBSSessionSetuporModifyRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSSessionSetuporModifyRequestItemOptPresentFlag, true)
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
	if x.AssociatedMBSQosFlowSetuporModifyRequestList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AssociatedMBSQosFlowSetuporModifyRequestList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AssociatedMBSQosFlowSetuporModifyRequestList marshal failed")
		}
	}

	// optional field
	if x.MBSQosFlowToReleaseList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSQosFlowToReleaseList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSQosFlowToReleaseList marshal failed")
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

func (x *MBSSessionSetuporModifyRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSSessionSetuporModifyRequestItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&MBSSessionSetuporModifyRequestItemOptPresentFlag, true)
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
	if MBSSessionSetuporModifyRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MBSAreaSessionID = new(MBSAreaSessionID)
		err = x.MBSAreaSessionID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSAreaSessionID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSSessionSetuporModifyRequestItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.AssociatedMBSQosFlowSetuporModifyRequestList = new(AssociatedMBSQosFlowSetuporModifyRequestList)
		err = x.AssociatedMBSQosFlowSetuporModifyRequestList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AssociatedMBSQosFlowSetuporModifyRequestList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if MBSSessionSetuporModifyRequestItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.MBSQosFlowToReleaseList = new(QosFlowListWithCause)
		err = x.MBSQosFlowToReleaseList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSQosFlowToReleaseList error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if MBSSessionSetuporModifyRequestItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSSessionSetuporModifyRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
