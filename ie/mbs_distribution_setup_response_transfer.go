package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSDistributionSetupResponseTransfer struct {
	MBSSessionID                     *MBSSessionID                     // valueExt
	MBSAreaSessionID                 *MBSAreaSessionID                 // optional
	SharedNGUMulticastTNLInformation *SharedNGUMulticastTNLInformation // valueExt,optional
	MBSQoSFlowsToBeSetupList         *MBSQoSFlowsToBeSetupList
	MBSSessionStatus                 *MBSSessionStatus                                                     // valueExt,valueLB:0,valueUB:1
	MBSServiceArea                   *MBSServiceArea                                                       // valueLB:0,valueUB:2,optional
	IEExtensions                     *ProtocolExtensionContainerMBSDistributionSetupResponseTransferExtIEs // optional
}

func (x *MBSDistributionSetupResponseTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSDistributionSetupResponseTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSSessionID == nil {
		return errors.Errorf("MBSSessionID is missing")
	}
	// optional field
	if x.MBSAreaSessionID != nil {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, true)
	} else {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.SharedNGUMulticastTNLInformation != nil {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, true)
	} else {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, false)
	}
	// mandatory field
	if x.MBSQoSFlowsToBeSetupList == nil {
		return errors.Errorf("MBSQoSFlowsToBeSetupList is missing")
	}
	// mandatory field
	if x.MBSSessionStatus == nil {
		return errors.Errorf("MBSSessionStatus is missing")
	}
	// optional field
	if x.MBSServiceArea != nil {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, true)
	} else {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, true)
	} else {
		MBSDistributionSetupResponseTransferOptPresentFlag = append(MBSDistributionSetupResponseTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSDistributionSetupResponseTransferOptPresentFlag, true)
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
	if x.SharedNGUMulticastTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SharedNGUMulticastTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SharedNGUMulticastTNLInformation marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MBSQoSFlowsToBeSetupList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSQoSFlowsToBeSetupList marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MBSSessionStatus.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSSessionStatus marshal failed")
	}

	// optional field
	if x.MBSServiceArea != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSServiceArea.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSServiceArea marshal failed")
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

func (x *MBSDistributionSetupResponseTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSDistributionSetupResponseTransferOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&MBSDistributionSetupResponseTransferOptPresentFlag, true)
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
	if MBSDistributionSetupResponseTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MBSAreaSessionID = new(MBSAreaSessionID)
		err = x.MBSAreaSessionID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSAreaSessionID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSDistributionSetupResponseTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SharedNGUMulticastTNLInformation = new(SharedNGUMulticastTNLInformation)
		err = x.SharedNGUMulticastTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SharedNGUMulticastTNLInformation error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSQoSFlowsToBeSetupList = new(MBSQoSFlowsToBeSetupList)
	err = x.MBSQoSFlowsToBeSetupList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSQoSFlowsToBeSetupList error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSSessionStatus = new(MBSSessionStatus)
	err = x.MBSSessionStatus.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSSessionStatus error")
	}

	// optional field (optPresentFlag index: 2)
	if MBSDistributionSetupResponseTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.MBSServiceArea = new(MBSServiceArea)
		err = x.MBSServiceArea.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSServiceArea error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if MBSDistributionSetupResponseTransferOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSDistributionSetupResponseTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
