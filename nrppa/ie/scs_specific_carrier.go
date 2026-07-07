package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	SCSSpecificCarrierSubcarrierSpacingPresentKHz15  aper.Enumerated = 0
	SCSSpecificCarrierSubcarrierSpacingPresentKHz30  aper.Enumerated = 1
	SCSSpecificCarrierSubcarrierSpacingPresentKHz60  aper.Enumerated = 2
	SCSSpecificCarrierSubcarrierSpacingPresentKHz120 aper.Enumerated = 3
	SCSSpecificCarrierSubcarrierSpacingPresentKHz480 aper.Enumerated = 4
	SCSSpecificCarrierSubcarrierSpacingPresentKHz960 aper.Enumerated = 5
)

type SCSSpecificCarrier struct {
	OffsetToCarrier   *int64                                              // valueExt,valueLB:0,valueUB:2199
	SubcarrierSpacing *aper.Enumerated                                    // valueExt,valueLB:0,valueUB:3
	CarrierBandwidth  *int64                                              // valueExt,valueLB:1,valueUB:275
	IEExtensions      *ProtocolExtensionContainerSCSSpecificCarrierExtIEs // optional
}

func (x *SCSSpecificCarrier) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SCSSpecificCarrierOptPresentFlag := []bool{}
	// mandatory field
	if x.OffsetToCarrier == nil {
		return errors.Errorf("OffsetToCarrier is missing")
	}
	// mandatory field
	if x.SubcarrierSpacing == nil {
		return errors.Errorf("SubcarrierSpacing is missing")
	}
	// mandatory field
	if x.CarrierBandwidth == nil {
		return errors.Errorf("CarrierBandwidth is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SCSSpecificCarrierOptPresentFlag = append(SCSSpecificCarrierOptPresentFlag, true)
	} else {
		SCSSpecificCarrierOptPresentFlag = append(SCSSpecificCarrierOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SCSSpecificCarrierOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 2199
	err = pd.WriteInteger(*(x.OffsetToCarrier), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.SubcarrierSpacing), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 1, 275
	err = pd.WriteInteger(*(x.CarrierBandwidth), true, vLb, vUb)
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

func (x *SCSSpecificCarrier) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SCSSpecificCarrierOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SCSSpecificCarrierOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 2199
	x.OffsetToCarrier = new(int64)
	*(x.OffsetToCarrier), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.SubcarrierSpacing = new(aper.Enumerated)
	*(x.SubcarrierSpacing), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 1, 275
	x.CarrierBandwidth = new(int64)
	*(x.CarrierBandwidth), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if SCSSpecificCarrierOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSCSSpecificCarrierExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
