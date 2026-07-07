package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSMappingandDataForwardingRequestItem struct {
	MRBID                  *MRBID
	MBSQoSFlowList         *MBSQoSFlowList
	MRBProgressInformation *MRBProgressInformation                                                 // valueLB:0,valueUB:2,optional
	IEExtensions           *ProtocolExtensionContainerMBSMappingandDataForwardingRequestItemExtIEs // optional
}

func (x *MBSMappingandDataForwardingRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSMappingandDataForwardingRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MRBID == nil {
		return errors.Errorf("MRBID is missing")
	}
	// mandatory field
	if x.MBSQoSFlowList == nil {
		return errors.Errorf("MBSQoSFlowList is missing")
	}
	// optional field
	if x.MRBProgressInformation != nil {
		MBSMappingandDataForwardingRequestItemOptPresentFlag = append(MBSMappingandDataForwardingRequestItemOptPresentFlag, true)
	} else {
		MBSMappingandDataForwardingRequestItemOptPresentFlag = append(MBSMappingandDataForwardingRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSMappingandDataForwardingRequestItemOptPresentFlag = append(MBSMappingandDataForwardingRequestItemOptPresentFlag, true)
	} else {
		MBSMappingandDataForwardingRequestItemOptPresentFlag = append(MBSMappingandDataForwardingRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSMappingandDataForwardingRequestItemOptPresentFlag, true)
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
	err = x.MBSQoSFlowList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSQoSFlowList marshal failed")
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

func (x *MBSMappingandDataForwardingRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSMappingandDataForwardingRequestItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MBSMappingandDataForwardingRequestItemOptPresentFlag, true)
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
	x.MBSQoSFlowList = new(MBSQoSFlowList)
	err = x.MBSQoSFlowList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSQoSFlowList error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSMappingandDataForwardingRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MRBProgressInformation = new(MRBProgressInformation)
		err = x.MRBProgressInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MRBProgressInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSMappingandDataForwardingRequestItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSMappingandDataForwardingRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
