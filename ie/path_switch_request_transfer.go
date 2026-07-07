package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PathSwitchRequestTransfer struct {
	DLNGUUPTNLInformation        *UPTransportLayerInformation  // valueLB:0,valueUB:1
	DLNGUTNLInformationReused    *DLNGUTNLInformationReused    // valueExt,valueLB:0,valueUB:0,optional
	UserPlaneSecurityInformation *UserPlaneSecurityInformation // valueExt,optional
	QosFlowAcceptedList          *QosFlowAcceptedList
	IEExtensions                 *ProtocolExtensionContainerPathSwitchRequestTransferExtIEs // optional
}

func (x *PathSwitchRequestTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PathSwitchRequestTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.DLNGUUPTNLInformation == nil {
		return errors.Errorf("DLNGUUPTNLInformation is missing")
	}
	// optional field
	if x.DLNGUTNLInformationReused != nil {
		PathSwitchRequestTransferOptPresentFlag = append(PathSwitchRequestTransferOptPresentFlag, true)
	} else {
		PathSwitchRequestTransferOptPresentFlag = append(PathSwitchRequestTransferOptPresentFlag, false)
	}
	// optional field
	if x.UserPlaneSecurityInformation != nil {
		PathSwitchRequestTransferOptPresentFlag = append(PathSwitchRequestTransferOptPresentFlag, true)
	} else {
		PathSwitchRequestTransferOptPresentFlag = append(PathSwitchRequestTransferOptPresentFlag, false)
	}
	// mandatory field
	if x.QosFlowAcceptedList == nil {
		return errors.Errorf("QosFlowAcceptedList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PathSwitchRequestTransferOptPresentFlag = append(PathSwitchRequestTransferOptPresentFlag, true)
	} else {
		PathSwitchRequestTransferOptPresentFlag = append(PathSwitchRequestTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PathSwitchRequestTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DLNGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLNGUUPTNLInformation marshal failed")
	}

	// optional field
	if x.DLNGUTNLInformationReused != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLNGUTNLInformationReused.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLNGUTNLInformationReused marshal failed")
		}
	}

	// optional field
	if x.UserPlaneSecurityInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UserPlaneSecurityInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UserPlaneSecurityInformation marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.QosFlowAcceptedList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QosFlowAcceptedList marshal failed")
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

func (x *PathSwitchRequestTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PathSwitchRequestTransferOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PathSwitchRequestTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.DLNGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLNGUUPTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if PathSwitchRequestTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLNGUTNLInformationReused = new(DLNGUTNLInformationReused)
		err = x.DLNGUTNLInformationReused.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLNGUTNLInformationReused error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PathSwitchRequestTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.UserPlaneSecurityInformation = new(UserPlaneSecurityInformation)
		err = x.UserPlaneSecurityInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UserPlaneSecurityInformation error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QosFlowAcceptedList = new(QosFlowAcceptedList)
	err = x.QosFlowAcceptedList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QosFlowAcceptedList error")
	}

	// optional field (optPresentFlag index: 2)
	if PathSwitchRequestTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPathSwitchRequestTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
