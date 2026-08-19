package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type LCSToGCSTranslation struct {
	Alpha        *int64                                               // valueLB:0,valueUB:3599
	Beta         *int64                                               // valueLB:0,valueUB:3599
	Gamma        *int64                                               // valueLB:0,valueUB:3599
	IEExtensions *ProtocolExtensionContainerLCSToGCSTranslationExtIEs // optional
}

func (x *LCSToGCSTranslation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LCSToGCSTranslationOptPresentFlag := []bool{}
	// mandatory field
	if x.Alpha == nil {
		return errors.Errorf("Alpha is missing")
	}
	// mandatory field
	if x.Beta == nil {
		return errors.Errorf("Beta is missing")
	}
	// mandatory field
	if x.Gamma == nil {
		return errors.Errorf("Gamma is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		LCSToGCSTranslationOptPresentFlag = append(LCSToGCSTranslationOptPresentFlag, true)
	} else {
		LCSToGCSTranslationOptPresentFlag = append(LCSToGCSTranslationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LCSToGCSTranslationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3599
	err = pd.WriteInteger(*(x.Alpha), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3599
	err = pd.WriteInteger(*(x.Beta), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3599
	err = pd.WriteInteger(*(x.Gamma), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *LCSToGCSTranslation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LCSToGCSTranslationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&LCSToGCSTranslationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3599
	x.Alpha = new(int64)
	*(x.Alpha), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3599
	x.Beta = new(int64)
	*(x.Beta), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3599
	x.Gamma = new(int64)
	*(x.Gamma), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if LCSToGCSTranslationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLCSToGCSTranslationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
