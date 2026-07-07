package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	ActiveULBWPSubcarrierSpacingPresentKHz15  aper.Enumerated = 0
	ActiveULBWPSubcarrierSpacingPresentKHz30  aper.Enumerated = 1
	ActiveULBWPSubcarrierSpacingPresentKHz60  aper.Enumerated = 2
	ActiveULBWPSubcarrierSpacingPresentKHz120 aper.Enumerated = 3
	ActiveULBWPSubcarrierSpacingPresentKHz480 aper.Enumerated = 4
	ActiveULBWPSubcarrierSpacingPresentKHz960 aper.Enumerated = 5
)

const ( /* Enum Type */
	ActiveULBWPCyclicPrefixPresentNormal   aper.Enumerated = 0
	ActiveULBWPCyclicPrefixPresentExtended aper.Enumerated = 1
)

const ( /* Enum Type */
	ActiveULBWPShift7dot5kHzPresentTrue aper.Enumerated = 0
)

type ActiveULBWP struct {
	LocationAndBandwidth    *int64                                       // valueExt,valueLB:0,valueUB:37949
	SubcarrierSpacing       *aper.Enumerated                             // valueExt,valueLB:0,valueUB:3
	CyclicPrefix            *aper.Enumerated                             // valueLB:0,valueUB:1
	TxDirectCurrentLocation *int64                                       // valueExt,valueLB:0,valueUB:3301
	Shift7dot5kHz           *aper.Enumerated                             // valueExt,valueLB:0,valueUB:0,optional
	SRSConfig               *SRSConfig                                   // valueExt
	IEExtensions            *ProtocolExtensionContainerActiveULBWPExtIEs // optional
}

func (x *ActiveULBWP) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ActiveULBWPOptPresentFlag := []bool{}
	// mandatory field
	if x.LocationAndBandwidth == nil {
		return errors.Errorf("LocationAndBandwidth is missing")
	}
	// mandatory field
	if x.SubcarrierSpacing == nil {
		return errors.Errorf("SubcarrierSpacing is missing")
	}
	// mandatory field
	if x.CyclicPrefix == nil {
		return errors.Errorf("CyclicPrefix is missing")
	}
	// mandatory field
	if x.TxDirectCurrentLocation == nil {
		return errors.Errorf("TxDirectCurrentLocation is missing")
	}
	// optional field
	if x.Shift7dot5kHz != nil {
		ActiveULBWPOptPresentFlag = append(ActiveULBWPOptPresentFlag, true)
	} else {
		ActiveULBWPOptPresentFlag = append(ActiveULBWPOptPresentFlag, false)
	}
	// mandatory field
	if x.SRSConfig == nil {
		return errors.Errorf("SRSConfig is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ActiveULBWPOptPresentFlag = append(ActiveULBWPOptPresentFlag, true)
	} else {
		ActiveULBWPOptPresentFlag = append(ActiveULBWPOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ActiveULBWPOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 37949
	err = pd.WriteInteger(*(x.LocationAndBandwidth), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.SubcarrierSpacing), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.CyclicPrefix), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3301
	err = pd.WriteInteger(*(x.TxDirectCurrentLocation), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.Shift7dot5kHz != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.Shift7dot5kHz), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SRSConfig.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSConfig marshal failed")
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

func (x *ActiveULBWP) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ActiveULBWPOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ActiveULBWPOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 37949
	x.LocationAndBandwidth = new(int64)
	*(x.LocationAndBandwidth), err = pd.ReadInteger(true, vLb, vUb)
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
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.CyclicPrefix = new(aper.Enumerated)
	*(x.CyclicPrefix), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3301
	x.TxDirectCurrentLocation = new(int64)
	*(x.TxDirectCurrentLocation), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if ActiveULBWPOptPresentFlag[0] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.Shift7dot5kHz = new(aper.Enumerated)
		*(x.Shift7dot5kHz), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSConfig = new(SRSConfig)
	err = x.SRSConfig.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSConfig error")
	}

	// optional field (optPresentFlag index: 1)
	if ActiveULBWPOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerActiveULBWPExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
