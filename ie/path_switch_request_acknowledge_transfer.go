package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation                                          // valueLB:0,valueUB:1,optional
	SecurityIndication    *SecurityIndication                                                   // valueExt,optional
	IEExtensions          *ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs // optional
}

func (x *PathSwitchRequestAcknowledgeTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PathSwitchRequestAcknowledgeTransferOptPresentFlag := []bool{}
	// optional field
	if x.ULNGUUPTNLInformation != nil {
		PathSwitchRequestAcknowledgeTransferOptPresentFlag = append(PathSwitchRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		PathSwitchRequestAcknowledgeTransferOptPresentFlag = append(PathSwitchRequestAcknowledgeTransferOptPresentFlag, false)
	}
	// optional field
	if x.SecurityIndication != nil {
		PathSwitchRequestAcknowledgeTransferOptPresentFlag = append(PathSwitchRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		PathSwitchRequestAcknowledgeTransferOptPresentFlag = append(PathSwitchRequestAcknowledgeTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PathSwitchRequestAcknowledgeTransferOptPresentFlag = append(PathSwitchRequestAcknowledgeTransferOptPresentFlag, true)
	} else {
		PathSwitchRequestAcknowledgeTransferOptPresentFlag = append(PathSwitchRequestAcknowledgeTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PathSwitchRequestAcknowledgeTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.ULNGUUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ULNGUUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ULNGUUPTNLInformation marshal failed")
		}
	}

	// optional field
	if x.SecurityIndication != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SecurityIndication.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SecurityIndication marshal failed")
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

func (x *PathSwitchRequestAcknowledgeTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PathSwitchRequestAcknowledgeTransferOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PathSwitchRequestAcknowledgeTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if PathSwitchRequestAcknowledgeTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
		err = x.ULNGUUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ULNGUUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PathSwitchRequestAcknowledgeTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SecurityIndication = new(SecurityIndication)
		err = x.SecurityIndication.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SecurityIndication error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PathSwitchRequestAcknowledgeTransferOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
