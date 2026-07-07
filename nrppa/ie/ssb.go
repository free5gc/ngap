package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SSB struct {
	PCINR        *int64                               // valueLB:0,valueUB:1007
	SsbIndex     *SSBIndex                            // optional
	IEExtensions *ProtocolExtensionContainerSSBExtIEs // optional
}

func (x *SSB) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBOptPresentFlag := []bool{}
	// mandatory field
	if x.PCINR == nil {
		return errors.Errorf("PCINR is missing")
	}
	// optional field
	if x.SsbIndex != nil {
		SSBOptPresentFlag = append(SSBOptPresentFlag, true)
	} else {
		SSBOptPresentFlag = append(SSBOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SSBOptPresentFlag = append(SSBOptPresentFlag, true)
	} else {
		SSBOptPresentFlag = append(SSBOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SSBOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1007
	err = pd.WriteInteger(*(x.PCINR), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.SsbIndex != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SsbIndex.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SsbIndex marshal failed")
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

func (x *SSB) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SSBOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1007
	x.PCINR = new(int64)
	*(x.PCINR), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if SSBOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SsbIndex = new(SSBIndex)
		err = x.SsbIndex.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SsbIndex error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SSBOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSSBExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
