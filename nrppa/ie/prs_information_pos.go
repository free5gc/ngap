package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSInformationPos struct {
	PRSIDPos            *int64                                             // valueLB:0,valueUB:255
	PRSResourceSetIDPos *int64                                             // valueLB:0,valueUB:7
	PRSResourceIDPos    *int64                                             // valueLB:0,valueUB:63,optional
	IEExtensions        *ProtocolExtensionContainerPRSInformationPosExtIEs // optional
}

func (x *PRSInformationPos) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSInformationPosOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSIDPos == nil {
		return errors.Errorf("PRSIDPos is missing")
	}
	// mandatory field
	if x.PRSResourceSetIDPos == nil {
		return errors.Errorf("PRSResourceSetIDPos is missing")
	}
	// optional field
	if x.PRSResourceIDPos != nil {
		PRSInformationPosOptPresentFlag = append(PRSInformationPosOptPresentFlag, true)
	} else {
		PRSInformationPosOptPresentFlag = append(PRSInformationPosOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PRSInformationPosOptPresentFlag = append(PRSInformationPosOptPresentFlag, true)
	} else {
		PRSInformationPosOptPresentFlag = append(PRSInformationPosOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSInformationPosOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.PRSIDPos), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 7
	err = pd.WriteInteger(*(x.PRSResourceSetIDPos), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.PRSResourceIDPos != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 63
		err = pd.WriteInteger(*(x.PRSResourceIDPos), false, vLb, vUb)
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

func (x *PRSInformationPos) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSInformationPosOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PRSInformationPosOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.PRSIDPos = new(int64)
	*(x.PRSIDPos), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 7
	x.PRSResourceSetIDPos = new(int64)
	*(x.PRSResourceSetIDPos), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if PRSInformationPosOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 63
		x.PRSResourceIDPos = new(int64)
		*(x.PRSResourceIDPos), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if PRSInformationPosOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSInformationPosExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
