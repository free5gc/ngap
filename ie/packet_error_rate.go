package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PacketErrorRate struct {
	PERScalar    *int64                                           // valueExt,valueLB:0,valueUB:9
	PERExponent  *int64                                           // valueExt,valueLB:0,valueUB:9
	IEExtensions *ProtocolExtensionContainerPacketErrorRateExtIEs // optional
}

func (x *PacketErrorRate) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PacketErrorRateOptPresentFlag := []bool{}
	// mandatory field
	if x.PERScalar == nil {
		return errors.Errorf("PERScalar is missing")
	}
	// mandatory field
	if x.PERExponent == nil {
		return errors.Errorf("PERExponent is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PacketErrorRateOptPresentFlag = append(PacketErrorRateOptPresentFlag, true)
	} else {
		PacketErrorRateOptPresentFlag = append(PacketErrorRateOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PacketErrorRateOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 9
	err = pd.WriteInteger(*(x.PERScalar), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 9
	err = pd.WriteInteger(*(x.PERExponent), true, vLb, vUb)
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

func (x *PacketErrorRate) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PacketErrorRateOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PacketErrorRateOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 9
	x.PERScalar = new(int64)
	*(x.PERScalar), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 9
	x.PERExponent = new(int64)
	*(x.PERExponent), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if PacketErrorRateOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPacketErrorRateExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
