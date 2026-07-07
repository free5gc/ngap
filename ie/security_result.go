package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SecurityResult struct {
	IntegrityProtectionResult       *IntegrityProtectionResult                      // valueExt,valueLB:0,valueUB:1
	ConfidentialityProtectionResult *ConfidentialityProtectionResult                // valueExt,valueLB:0,valueUB:1
	IEExtensions                    *ProtocolExtensionContainerSecurityResultExtIEs // optional
}

func (x *SecurityResult) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SecurityResultOptPresentFlag := []bool{}
	// mandatory field
	if x.IntegrityProtectionResult == nil {
		return errors.Errorf("IntegrityProtectionResult is missing")
	}
	// mandatory field
	if x.ConfidentialityProtectionResult == nil {
		return errors.Errorf("ConfidentialityProtectionResult is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SecurityResultOptPresentFlag = append(SecurityResultOptPresentFlag, true)
	} else {
		SecurityResultOptPresentFlag = append(SecurityResultOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SecurityResultOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.IntegrityProtectionResult.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IntegrityProtectionResult marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ConfidentialityProtectionResult.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ConfidentialityProtectionResult marshal failed")
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

func (x *SecurityResult) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SecurityResultOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SecurityResultOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IntegrityProtectionResult = new(IntegrityProtectionResult)
	err = x.IntegrityProtectionResult.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IntegrityProtectionResult error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ConfidentialityProtectionResult = new(ConfidentialityProtectionResult)
	err = x.ConfidentialityProtectionResult.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ConfidentialityProtectionResult error")
	}

	// optional field (optPresentFlag index: 0)
	if SecurityResultOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSecurityResultExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
