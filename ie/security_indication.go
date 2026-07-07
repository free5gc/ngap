package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &SecurityIndication{}

type SecurityIndication struct {
	IntegrityProtectionIndication       *IntegrityProtectionIndication                      // valueExt,valueLB:0,valueUB:2
	ConfidentialityProtectionIndication *ConfidentialityProtectionIndication                // valueExt,valueLB:0,valueUB:2
	MaximumIntegrityProtectedDataRateUL *MaximumIntegrityProtectedDataRate                  // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions                        *ProtocolExtensionContainerSecurityIndicationExtIEs // optional
}

func (x *SecurityIndication) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SecurityIndicationOptPresentFlag := []bool{}
	// mandatory field
	if x.IntegrityProtectionIndication == nil {
		return errors.Errorf("IntegrityProtectionIndication is missing")
	}
	// mandatory field
	if x.ConfidentialityProtectionIndication == nil {
		return errors.Errorf("ConfidentialityProtectionIndication is missing")
	}
	// optional field
	if x.MaximumIntegrityProtectedDataRateUL != nil {
		SecurityIndicationOptPresentFlag = append(SecurityIndicationOptPresentFlag, true)
	} else {
		SecurityIndicationOptPresentFlag = append(SecurityIndicationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SecurityIndicationOptPresentFlag = append(SecurityIndicationOptPresentFlag, true)
	} else {
		SecurityIndicationOptPresentFlag = append(SecurityIndicationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SecurityIndicationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.IntegrityProtectionIndication.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IntegrityProtectionIndication marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ConfidentialityProtectionIndication.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ConfidentialityProtectionIndication marshal failed")
	}

	// optional field
	if x.MaximumIntegrityProtectedDataRateUL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MaximumIntegrityProtectedDataRateUL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MaximumIntegrityProtectedDataRateUL marshal failed")
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

func (x *SecurityIndication) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SecurityIndicationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SecurityIndicationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IntegrityProtectionIndication = new(IntegrityProtectionIndication)
	err = x.IntegrityProtectionIndication.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IntegrityProtectionIndication error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ConfidentialityProtectionIndication = new(ConfidentialityProtectionIndication)
	err = x.ConfidentialityProtectionIndication.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ConfidentialityProtectionIndication error")
	}

	// optional field (optPresentFlag index: 0)
	if SecurityIndicationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MaximumIntegrityProtectedDataRateUL = new(MaximumIntegrityProtectedDataRate)
		err = x.MaximumIntegrityProtectedDataRateUL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MaximumIntegrityProtectedDataRateUL error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SecurityIndicationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSecurityIndicationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *SecurityIndication) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *SecurityIndication) ReadIE(pd *aper.PerBitData) error {
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
