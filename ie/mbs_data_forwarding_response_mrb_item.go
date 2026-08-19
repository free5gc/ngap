package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSDataForwardingResponseMRBItem struct {
	MRBID                        *MRBID
	DLForwardingUPTNLInformation *UPTransportLayerInformation                                      // valueLB:0,valueUB:1
	MRBProgressInformation       *MRBProgressInformation                                           // valueLB:0,valueUB:2,optional
	IEExtensions                 *ProtocolExtensionContainerMBSDataForwardingResponseMRBItemExtIEs // optional
}

func (x *MBSDataForwardingResponseMRBItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSDataForwardingResponseMRBItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MRBID == nil {
		return errors.Errorf("MRBID is missing")
	}
	// mandatory field
	if x.DLForwardingUPTNLInformation == nil {
		return errors.Errorf("DLForwardingUPTNLInformation is missing")
	}
	// optional field
	if x.MRBProgressInformation != nil {
		MBSDataForwardingResponseMRBItemOptPresentFlag = append(MBSDataForwardingResponseMRBItemOptPresentFlag, true)
	} else {
		MBSDataForwardingResponseMRBItemOptPresentFlag = append(MBSDataForwardingResponseMRBItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSDataForwardingResponseMRBItemOptPresentFlag = append(MBSDataForwardingResponseMRBItemOptPresentFlag, true)
	} else {
		MBSDataForwardingResponseMRBItemOptPresentFlag = append(MBSDataForwardingResponseMRBItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSDataForwardingResponseMRBItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MRBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MRBID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DLForwardingUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLForwardingUPTNLInformation marshal failed")
	}

	// optional field
	if x.MRBProgressInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MRBProgressInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MRBProgressInformation marshal failed")
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

func (x *MBSDataForwardingResponseMRBItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSDataForwardingResponseMRBItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MBSDataForwardingResponseMRBItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MRBID = new(MRBID)
	err = x.MRBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MRBID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLForwardingUPTNLInformation = new(UPTransportLayerInformation)
	err = x.DLForwardingUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLForwardingUPTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSDataForwardingResponseMRBItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MRBProgressInformation = new(MRBProgressInformation)
		err = x.MRBProgressInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MRBProgressInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSDataForwardingResponseMRBItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSDataForwardingResponseMRBItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
