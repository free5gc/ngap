package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &UESecurityCapabilities{}

type UESecurityCapabilities struct {
	NRencryptionAlgorithms             *NRencryptionAlgorithms
	NRintegrityProtectionAlgorithms    *NRintegrityProtectionAlgorithms
	EUTRAencryptionAlgorithms          *EUTRAencryptionAlgorithms
	EUTRAintegrityProtectionAlgorithms *EUTRAintegrityProtectionAlgorithms
	IEExtensions                       *ProtocolExtensionContainerUESecurityCapabilitiesExtIEs // optional
}

func (x *UESecurityCapabilities) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UESecurityCapabilitiesOptPresentFlag := []bool{}
	// mandatory field
	if x.NRencryptionAlgorithms == nil {
		return errors.Errorf("NRencryptionAlgorithms is missing")
	}
	// mandatory field
	if x.NRintegrityProtectionAlgorithms == nil {
		return errors.Errorf("NRintegrityProtectionAlgorithms is missing")
	}
	// mandatory field
	if x.EUTRAencryptionAlgorithms == nil {
		return errors.Errorf("EUTRAencryptionAlgorithms is missing")
	}
	// mandatory field
	if x.EUTRAintegrityProtectionAlgorithms == nil {
		return errors.Errorf("EUTRAintegrityProtectionAlgorithms is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UESecurityCapabilitiesOptPresentFlag = append(UESecurityCapabilitiesOptPresentFlag, true)
	} else {
		UESecurityCapabilitiesOptPresentFlag = append(UESecurityCapabilitiesOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UESecurityCapabilitiesOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRencryptionAlgorithms.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRencryptionAlgorithms marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NRintegrityProtectionAlgorithms.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRintegrityProtectionAlgorithms marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRAencryptionAlgorithms.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRAencryptionAlgorithms marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRAintegrityProtectionAlgorithms.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRAintegrityProtectionAlgorithms marshal failed")
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

func (x *UESecurityCapabilities) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UESecurityCapabilitiesOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UESecurityCapabilitiesOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRencryptionAlgorithms = new(NRencryptionAlgorithms)
	err = x.NRencryptionAlgorithms.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRencryptionAlgorithms error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRintegrityProtectionAlgorithms = new(NRintegrityProtectionAlgorithms)
	err = x.NRintegrityProtectionAlgorithms.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRintegrityProtectionAlgorithms error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRAencryptionAlgorithms = new(EUTRAencryptionAlgorithms)
	err = x.EUTRAencryptionAlgorithms.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRAencryptionAlgorithms error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRAintegrityProtectionAlgorithms = new(EUTRAintegrityProtectionAlgorithms)
	err = x.EUTRAintegrityProtectionAlgorithms.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRAintegrityProtectionAlgorithms error")
	}

	// optional field (optPresentFlag index: 0)
	if UESecurityCapabilitiesOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUESecurityCapabilitiesExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *UESecurityCapabilities) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *UESecurityCapabilities) ReadIE(pd *aper.PerBitData) error {
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
