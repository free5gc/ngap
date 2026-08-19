package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation                          // valueLB:0,valueUB:1
	DLNGUUPTNLInformation *UPTransportLayerInformation                          // valueLB:0,valueUB:1
	IEExtensions          *ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs // optional
}

func (x *ULNGUUPTNLModifyItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULNGUUPTNLModifyItemOptPresentFlag := []bool{}
	// mandatory field
	if x.ULNGUUPTNLInformation == nil {
		return errors.Errorf("ULNGUUPTNLInformation is missing")
	}
	// mandatory field
	if x.DLNGUUPTNLInformation == nil {
		return errors.Errorf("DLNGUUPTNLInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ULNGUUPTNLModifyItemOptPresentFlag = append(ULNGUUPTNLModifyItemOptPresentFlag, true)
	} else {
		ULNGUUPTNLModifyItemOptPresentFlag = append(ULNGUUPTNLModifyItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ULNGUUPTNLModifyItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ULNGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ULNGUUPTNLInformation marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DLNGUUPTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLNGUUPTNLInformation marshal failed")
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

func (x *ULNGUUPTNLModifyItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULNGUUPTNLModifyItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ULNGUUPTNLModifyItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ULNGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.ULNGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ULNGUUPTNLInformation error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLNGUUPTNLInformation = new(UPTransportLayerInformation)
	err = x.DLNGUUPTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLNGUUPTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if ULNGUUPTNLModifyItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
