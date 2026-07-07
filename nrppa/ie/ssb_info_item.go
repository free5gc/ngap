package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SSBInfoItem struct {
	SSBConfiguration *TFConfiguration                             // valueExt
	PCINR            *int64                                       // valueLB:0,valueUB:1007
	IEExtensions     *ProtocolExtensionContainerSSBInfoItemExtIEs // optional
}

func (x *SSBInfoItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBInfoItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SSBConfiguration == nil {
		return errors.Errorf("SSBConfiguration is missing")
	}
	// mandatory field
	if x.PCINR == nil {
		return errors.Errorf("PCINR is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SSBInfoItemOptPresentFlag = append(SSBInfoItemOptPresentFlag, true)
	} else {
		SSBInfoItemOptPresentFlag = append(SSBInfoItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SSBInfoItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SSBConfiguration.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SSBConfiguration marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1007
	err = pd.WriteInteger(*(x.PCINR), false, vLb, vUb)
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

func (x *SSBInfoItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBInfoItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SSBInfoItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SSBConfiguration = new(TFConfiguration)
	err = x.SSBConfiguration.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SSBConfiguration error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1007
	x.PCINR = new(int64)
	*(x.PCINR), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if SSBInfoItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSSBInfoItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
