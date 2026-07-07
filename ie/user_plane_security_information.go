package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UserPlaneSecurityInformation struct {
	SecurityResult     *SecurityResult                                               // valueExt
	SecurityIndication *SecurityIndication                                           // valueExt
	IEExtensions       *ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs // optional
}

func (x *UserPlaneSecurityInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UserPlaneSecurityInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.SecurityResult == nil {
		return errors.Errorf("SecurityResult is missing")
	}
	// mandatory field
	if x.SecurityIndication == nil {
		return errors.Errorf("SecurityIndication is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UserPlaneSecurityInformationOptPresentFlag = append(UserPlaneSecurityInformationOptPresentFlag, true)
	} else {
		UserPlaneSecurityInformationOptPresentFlag = append(UserPlaneSecurityInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UserPlaneSecurityInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SecurityResult.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SecurityResult marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SecurityIndication.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SecurityIndication marshal failed")
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

func (x *UserPlaneSecurityInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UserPlaneSecurityInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UserPlaneSecurityInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SecurityResult = new(SecurityResult)
	err = x.SecurityResult.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SecurityResult error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SecurityIndication = new(SecurityIndication)
	err = x.SecurityIndication.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SecurityIndication error")
	}

	// optional field (optPresentFlag index: 0)
	if UserPlaneSecurityInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUserPlaneSecurityInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
