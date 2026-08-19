package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type LCSToGCSTranslationItem struct {
	Alpha        *int64                                                   // valueLB:0,valueUB:359
	AlphaFine    *int64                                                   // valueLB:0,valueUB:9,optional
	Beta         *int64                                                   // valueLB:0,valueUB:359
	BetaFine     *int64                                                   // valueLB:0,valueUB:9,optional
	Gamma        *int64                                                   // valueLB:0,valueUB:359
	GammaFine    *int64                                                   // valueLB:0,valueUB:9,optional
	IEExtensions *ProtocolExtensionContainerLCSToGCSTranslationItemExtIEs // optional
}

func (x *LCSToGCSTranslationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LCSToGCSTranslationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.Alpha == nil {
		return errors.Errorf("Alpha is missing")
	}
	// optional field
	if x.AlphaFine != nil {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, true)
	} else {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, false)
	}
	// mandatory field
	if x.Beta == nil {
		return errors.Errorf("Beta is missing")
	}
	// optional field
	if x.BetaFine != nil {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, true)
	} else {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, false)
	}
	// mandatory field
	if x.Gamma == nil {
		return errors.Errorf("Gamma is missing")
	}
	// optional field
	if x.GammaFine != nil {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, true)
	} else {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, true)
	} else {
		LCSToGCSTranslationItemOptPresentFlag = append(LCSToGCSTranslationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LCSToGCSTranslationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 359
	err = pd.WriteInteger(*(x.Alpha), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.AlphaFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.AlphaFine), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 359
	err = pd.WriteInteger(*(x.Beta), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.BetaFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.BetaFine), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 359
	err = pd.WriteInteger(*(x.Gamma), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.GammaFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.GammaFine), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *LCSToGCSTranslationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LCSToGCSTranslationItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&LCSToGCSTranslationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 359
	x.Alpha = new(int64)
	*(x.Alpha), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if LCSToGCSTranslationItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.AlphaFine = new(int64)
		*(x.AlphaFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 359
	x.Beta = new(int64)
	*(x.Beta), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 1)
	if LCSToGCSTranslationItemOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.BetaFine = new(int64)
		*(x.BetaFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 359
	x.Gamma = new(int64)
	*(x.Gamma), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 2)
	if LCSToGCSTranslationItemOptPresentFlag[2] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.GammaFine = new(int64)
		*(x.GammaFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if LCSToGCSTranslationItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLCSToGCSTranslationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
