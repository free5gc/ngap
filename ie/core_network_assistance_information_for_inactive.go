package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &CoreNetworkAssistanceInformationForInactive{}

type CoreNetworkAssistanceInformationForInactive struct {
	UEIdentityIndexValue            *UEIdentityIndexValue // valueLB:0,valueUB:1
	UESpecificDRX                   *PagingDRX            // valueExt,valueLB:0,valueUB:3,optional
	PeriodicRegistrationUpdateTimer *PeriodicRegistrationUpdateTimer
	MICOModeIndication              *MICOModeIndication // valueExt,valueLB:0,valueUB:0,optional
	TAIListForInactive              *TAIListForInactive
	ExpectedUEBehaviour             *ExpectedUEBehaviour                                                         // valueExt,optional
	IEExtensions                    *ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs // optional
}

func (x *CoreNetworkAssistanceInformationForInactive) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CoreNetworkAssistanceInformationForInactiveOptPresentFlag := []bool{}
	// mandatory field
	if x.UEIdentityIndexValue == nil {
		return errors.Errorf("UEIdentityIndexValue is missing")
	}
	// optional field
	if x.UESpecificDRX != nil {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, true)
	} else {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, false)
	}
	// mandatory field
	if x.PeriodicRegistrationUpdateTimer == nil {
		return errors.Errorf("PeriodicRegistrationUpdateTimer is missing")
	}
	// optional field
	if x.MICOModeIndication != nil {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, true)
	} else {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, false)
	}
	// mandatory field
	if x.TAIListForInactive == nil {
		return errors.Errorf("TAIListForInactive is missing")
	}
	// optional field
	if x.ExpectedUEBehaviour != nil {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, true)
	} else {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, true)
	} else {
		CoreNetworkAssistanceInformationForInactiveOptPresentFlag = append(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CoreNetworkAssistanceInformationForInactiveOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.UEIdentityIndexValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UEIdentityIndexValue marshal failed")
	}

	// optional field
	if x.UESpecificDRX != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UESpecificDRX.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UESpecificDRX marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PeriodicRegistrationUpdateTimer.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PeriodicRegistrationUpdateTimer marshal failed")
	}

	// optional field
	if x.MICOModeIndication != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MICOModeIndication.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MICOModeIndication marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TAIListForInactive.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAIListForInactive marshal failed")
	}

	// optional field
	if x.ExpectedUEBehaviour != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExpectedUEBehaviour.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExpectedUEBehaviour marshal failed")
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

func (x *CoreNetworkAssistanceInformationForInactive) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CoreNetworkAssistanceInformationForInactiveOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&CoreNetworkAssistanceInformationForInactiveOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UEIdentityIndexValue = new(UEIdentityIndexValue)
	err = x.UEIdentityIndexValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UEIdentityIndexValue error")
	}

	// optional field (optPresentFlag index: 0)
	if CoreNetworkAssistanceInformationForInactiveOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UESpecificDRX = new(PagingDRX)
		err = x.UESpecificDRX.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UESpecificDRX error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PeriodicRegistrationUpdateTimer = new(PeriodicRegistrationUpdateTimer)
	err = x.PeriodicRegistrationUpdateTimer.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PeriodicRegistrationUpdateTimer error")
	}

	// optional field (optPresentFlag index: 1)
	if CoreNetworkAssistanceInformationForInactiveOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MICOModeIndication = new(MICOModeIndication)
		err = x.MICOModeIndication.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MICOModeIndication error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAIListForInactive = new(TAIListForInactive)
	err = x.TAIListForInactive.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAIListForInactive error")
	}

	// optional field (optPresentFlag index: 2)
	if CoreNetworkAssistanceInformationForInactiveOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.ExpectedUEBehaviour = new(ExpectedUEBehaviour)
		err = x.ExpectedUEBehaviour.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExpectedUEBehaviour error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if CoreNetworkAssistanceInformationForInactiveOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCoreNetworkAssistanceInformationForInactiveExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *CoreNetworkAssistanceInformationForInactive) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *CoreNetworkAssistanceInformationForInactive) ReadIE(pd *aper.PerBitData) error {
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
