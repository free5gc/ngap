package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AdditionalDLUPTNLInformationForHOItem struct {
	AdditionalDLNGUUPTNLInformation        *UPTransportLayerInformation // valueLB:0,valueUB:1
	AdditionalQosFlowSetupResponseList     *QosFlowListWithDataForwarding
	AdditionalDLForwardingUPTNLInformation *UPTransportLayerInformation                                           // valueLB:0,valueUB:1,optional
	IEExtensions                           *ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs // optional
}

func (x *AdditionalDLUPTNLInformationForHOItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AdditionalDLUPTNLInformationForHOItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AdditionalDLNGUUPTNLInformation == nil {
		return errors.Errorf("AdditionalDLNGUUPTNLInformation is missing")
	}
	// mandatory field
	if x.AdditionalQosFlowSetupResponseList == nil {
		return errors.Errorf("AdditionalQosFlowSetupResponseList is missing")
	}
	// optional field
	if x.AdditionalDLForwardingUPTNLInformation != nil {
		AdditionalDLUPTNLInformationForHOItemOptPresentFlag = append(AdditionalDLUPTNLInformationForHOItemOptPresentFlag, true)
	} else {
		AdditionalDLUPTNLInformationForHOItemOptPresentFlag = append(AdditionalDLUPTNLInformationForHOItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AdditionalDLUPTNLInformationForHOItemOptPresentFlag = append(AdditionalDLUPTNLInformationForHOItemOptPresentFlag, true)
	} else {
		AdditionalDLUPTNLInformationForHOItemOptPresentFlag = append(AdditionalDLUPTNLInformationForHOItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AdditionalDLUPTNLInformationForHOItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AdditionalDLNGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AdditionalDLNGUUPTNLInformation marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AdditionalQosFlowSetupResponseList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AdditionalQosFlowSetupResponseList marshal failed")
	}

	// optional field
	if x.AdditionalDLForwardingUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AdditionalDLForwardingUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AdditionalDLForwardingUPTNLInformation marshal failed")
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

func (x *AdditionalDLUPTNLInformationForHOItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AdditionalDLUPTNLInformationForHOItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&AdditionalDLUPTNLInformationForHOItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AdditionalDLNGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.AdditionalDLNGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AdditionalDLNGUUPTNLInformation error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AdditionalQosFlowSetupResponseList = new(QosFlowListWithDataForwarding)
	err = x.AdditionalQosFlowSetupResponseList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AdditionalQosFlowSetupResponseList error")
	}

	// optional field (optPresentFlag index: 0)
	if AdditionalDLUPTNLInformationForHOItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalDLForwardingUPTNLInformation = new(UPTransportLayerInformation)
		err = x.AdditionalDLForwardingUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalDLForwardingUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AdditionalDLUPTNLInformationForHOItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAdditionalDLUPTNLInformationForHOItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
