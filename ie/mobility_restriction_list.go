package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &MobilityRestrictionList{}

type MobilityRestrictionList struct {
	ServingPLMN              *PLMNIdentity
	EquivalentPLMNs          *EquivalentPLMNs                                         // optional
	RATRestrictions          *RATRestrictions                                         // optional
	ForbiddenAreaInformation *ForbiddenAreaInformation                                // optional
	ServiceAreaInformation   *ServiceAreaInformation                                  // optional
	IEExtensions             *ProtocolExtensionContainerMobilityRestrictionListExtIEs // optional
}

func (x *MobilityRestrictionList) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MobilityRestrictionListOptPresentFlag := []bool{}
	// mandatory field
	if x.ServingPLMN == nil {
		return errors.Errorf("ServingPLMN is missing")
	}
	// optional field
	if x.EquivalentPLMNs != nil {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, true)
	} else {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, false)
	}
	// optional field
	if x.RATRestrictions != nil {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, true)
	} else {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, false)
	}
	// optional field
	if x.ForbiddenAreaInformation != nil {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, true)
	} else {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, false)
	}
	// optional field
	if x.ServiceAreaInformation != nil {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, true)
	} else {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, true)
	} else {
		MobilityRestrictionListOptPresentFlag = append(MobilityRestrictionListOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MobilityRestrictionListOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ServingPLMN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ServingPLMN marshal failed")
	}

	// optional field
	if x.EquivalentPLMNs != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.EquivalentPLMNs.Write(pd)
		if err != nil {
			return errors.Wrap(err, "EquivalentPLMNs marshal failed")
		}
	}

	// optional field
	if x.RATRestrictions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RATRestrictions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RATRestrictions marshal failed")
		}
	}

	// optional field
	if x.ForbiddenAreaInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ForbiddenAreaInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ForbiddenAreaInformation marshal failed")
		}
	}

	// optional field
	if x.ServiceAreaInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ServiceAreaInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ServiceAreaInformation marshal failed")
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

func (x *MobilityRestrictionList) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MobilityRestrictionListOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&MobilityRestrictionListOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ServingPLMN = new(PLMNIdentity)
	err = x.ServingPLMN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ServingPLMN error")
	}

	// optional field (optPresentFlag index: 0)
	if MobilityRestrictionListOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.EquivalentPLMNs = new(EquivalentPLMNs)
		err = x.EquivalentPLMNs.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode EquivalentPLMNs error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MobilityRestrictionListOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.RATRestrictions = new(RATRestrictions)
		err = x.RATRestrictions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RATRestrictions error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if MobilityRestrictionListOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.ForbiddenAreaInformation = new(ForbiddenAreaInformation)
		err = x.ForbiddenAreaInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ForbiddenAreaInformation error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if MobilityRestrictionListOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.ServiceAreaInformation = new(ServiceAreaInformation)
		err = x.ServiceAreaInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ServiceAreaInformation error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if MobilityRestrictionListOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMobilityRestrictionListExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *MobilityRestrictionList) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *MobilityRestrictionList) ReadIE(pd *aper.PerBitData) error {
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
