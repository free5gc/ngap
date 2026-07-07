package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &SONConfigurationTransfer{}

type SONConfigurationTransfer struct {
	TargetRANNodeIDSON     *TargetRANNodeIDSON                                       // valueExt
	SourceRANNodeID        *SourceRANNodeID                                          // valueExt
	SONInformation         *SONInformation                                           // valueLB:0,valueUB:2
	XnTNLConfigurationInfo *XnTNLConfigurationInfo                                   // valueExt,optional
	IEExtensions           *ProtocolExtensionContainerSONConfigurationTransferExtIEs // optional
}

func (x *SONConfigurationTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SONConfigurationTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.TargetRANNodeIDSON == nil {
		return errors.Errorf("TargetRANNodeIDSON is missing")
	}
	// mandatory field
	if x.SourceRANNodeID == nil {
		return errors.Errorf("SourceRANNodeID is missing")
	}
	// mandatory field
	if x.SONInformation == nil {
		return errors.Errorf("SONInformation is missing")
	}
	// optional field
	if x.XnTNLConfigurationInfo != nil {
		SONConfigurationTransferOptPresentFlag = append(SONConfigurationTransferOptPresentFlag, true)
	} else {
		SONConfigurationTransferOptPresentFlag = append(SONConfigurationTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SONConfigurationTransferOptPresentFlag = append(SONConfigurationTransferOptPresentFlag, true)
	} else {
		SONConfigurationTransferOptPresentFlag = append(SONConfigurationTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SONConfigurationTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TargetRANNodeIDSON.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargetRANNodeIDSON marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SourceRANNodeID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SourceRANNodeID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SONInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SONInformation marshal failed")
	}

	// optional field
	if x.XnTNLConfigurationInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.XnTNLConfigurationInfo.Write(pd)
		if err != nil {
			return errors.Wrap(err, "XnTNLConfigurationInfo marshal failed")
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

func (x *SONConfigurationTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SONConfigurationTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SONConfigurationTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargetRANNodeIDSON = new(TargetRANNodeIDSON)
	err = x.TargetRANNodeIDSON.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargetRANNodeIDSON error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SourceRANNodeID = new(SourceRANNodeID)
	err = x.SourceRANNodeID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SourceRANNodeID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SONInformation = new(SONInformation)
	err = x.SONInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SONInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if SONConfigurationTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.XnTNLConfigurationInfo = new(XnTNLConfigurationInfo)
		err = x.XnTNLConfigurationInfo.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode XnTNLConfigurationInfo error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SONConfigurationTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSONConfigurationTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *SONConfigurationTransfer) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.WriteSequencePreambleBitMap(optPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: id
	err = id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: criticality
	err = criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: value (open type)
	pdOpenType := aper.NewPerBitData(nil)
	err = x.Write(pdOpenType)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}

	return nil
}

func (x *SONConfigurationTransfer) ReadIE(pd *aper.PerBitData) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.ReadSequencePreambleBitMap(&optPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err = protocolIECriticality.Read(pd)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	// sequence element: value (open type)
	var pdOpenTypeBytes []byte
	pdOpenTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	pdOpenType := aper.NewPerBitData(pdOpenTypeBytes)
	err = x.Read(pdOpenType)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	return nil
}
