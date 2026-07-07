package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &SourceToTargetAMFInformationReroute{}

type SourceToTargetAMFInformationReroute struct {
	ConfiguredNSSAI     *ConfiguredNSSAI                                                     // optional
	RejectedNSSAIinPLMN *RejectedNSSAIinPLMN                                                 // optional
	RejectedNSSAIinTA   *RejectedNSSAIinTA                                                   // optional
	IEExtensions        *ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs // optional
}

func (x *SourceToTargetAMFInformationReroute) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SourceToTargetAMFInformationRerouteOptPresentFlag := []bool{}
	// optional field
	if x.ConfiguredNSSAI != nil {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, true)
	} else {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, false)
	}
	// optional field
	if x.RejectedNSSAIinPLMN != nil {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, true)
	} else {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, false)
	}
	// optional field
	if x.RejectedNSSAIinTA != nil {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, true)
	} else {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, true)
	} else {
		SourceToTargetAMFInformationRerouteOptPresentFlag = append(SourceToTargetAMFInformationRerouteOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SourceToTargetAMFInformationRerouteOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.ConfiguredNSSAI != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ConfiguredNSSAI.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ConfiguredNSSAI marshal failed")
		}
	}

	// optional field
	if x.RejectedNSSAIinPLMN != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RejectedNSSAIinPLMN.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RejectedNSSAIinPLMN marshal failed")
		}
	}

	// optional field
	if x.RejectedNSSAIinTA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RejectedNSSAIinTA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RejectedNSSAIinTA marshal failed")
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

func (x *SourceToTargetAMFInformationReroute) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SourceToTargetAMFInformationRerouteOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&SourceToTargetAMFInformationRerouteOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if SourceToTargetAMFInformationRerouteOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ConfiguredNSSAI = new(ConfiguredNSSAI)
		err = x.ConfiguredNSSAI.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ConfiguredNSSAI error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SourceToTargetAMFInformationRerouteOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.RejectedNSSAIinPLMN = new(RejectedNSSAIinPLMN)
		err = x.RejectedNSSAIinPLMN.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RejectedNSSAIinPLMN error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if SourceToTargetAMFInformationRerouteOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.RejectedNSSAIinTA = new(RejectedNSSAIinTA)
		err = x.RejectedNSSAIinTA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RejectedNSSAIinTA error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if SourceToTargetAMFInformationRerouteOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSourceToTargetAMFInformationRerouteExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *SourceToTargetAMFInformationReroute) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *SourceToTargetAMFInformationReroute) ReadIE(pd *aper.PerBitData) error {
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
