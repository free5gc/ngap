package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	TFConfigurationSSBSubcarrierSpacingPresentKHz15  aper.Enumerated = 0
	TFConfigurationSSBSubcarrierSpacingPresentKHz30  aper.Enumerated = 1
	TFConfigurationSSBSubcarrierSpacingPresentKHz120 aper.Enumerated = 2
	TFConfigurationSSBSubcarrierSpacingPresentKHz240 aper.Enumerated = 3
	TFConfigurationSSBSubcarrierSpacingPresentKHz60  aper.Enumerated = 4
	TFConfigurationSSBSubcarrierSpacingPresentKHz480 aper.Enumerated = 5
	TFConfigurationSSBSubcarrierSpacingPresentKHz960 aper.Enumerated = 6
)

const ( /* Enum Type */
	TFConfigurationSSBPeriodicityPresentMs5   aper.Enumerated = 0
	TFConfigurationSSBPeriodicityPresentMs10  aper.Enumerated = 1
	TFConfigurationSSBPeriodicityPresentMs20  aper.Enumerated = 2
	TFConfigurationSSBPeriodicityPresentMs40  aper.Enumerated = 3
	TFConfigurationSSBPeriodicityPresentMs80  aper.Enumerated = 4
	TFConfigurationSSBPeriodicityPresentMs160 aper.Enumerated = 5
)

type TFConfiguration struct {
	SSBFrequency          *int64                                           // valueLB:0,valueUB:3279165
	SSBSubcarrierSpacing  *aper.Enumerated                                 // valueExt,valueLB:0,valueUB:3
	SSBTransmitPower      *int64                                           // valueLB:-60,valueUB:50
	SSBPeriodicity        *aper.Enumerated                                 // valueExt,valueLB:0,valueUB:5
	SSBHalfFrameOffset    *int64                                           // valueLB:0,valueUB:1
	SSBSFNOffset          *int64                                           // valueLB:0,valueUB:15
	SSBBurstPosition      *SSBBurstPosition                                // valueLB:0,valueUB:3,optional
	SFNInitialisationTime *RelativeTime1900                                // optional
	IEExtensions          *ProtocolExtensionContainerTFConfigurationExtIEs // optional
}

func (x *TFConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TFConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.SSBFrequency == nil {
		return errors.Errorf("SSBFrequency is missing")
	}
	// mandatory field
	if x.SSBSubcarrierSpacing == nil {
		return errors.Errorf("SSBSubcarrierSpacing is missing")
	}
	// mandatory field
	if x.SSBTransmitPower == nil {
		return errors.Errorf("SSBTransmitPower is missing")
	}
	// mandatory field
	if x.SSBPeriodicity == nil {
		return errors.Errorf("SSBPeriodicity is missing")
	}
	// mandatory field
	if x.SSBHalfFrameOffset == nil {
		return errors.Errorf("SSBHalfFrameOffset is missing")
	}
	// mandatory field
	if x.SSBSFNOffset == nil {
		return errors.Errorf("SSBSFNOffset is missing")
	}
	// optional field
	if x.SSBBurstPosition != nil {
		TFConfigurationOptPresentFlag = append(TFConfigurationOptPresentFlag, true)
	} else {
		TFConfigurationOptPresentFlag = append(TFConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.SFNInitialisationTime != nil {
		TFConfigurationOptPresentFlag = append(TFConfigurationOptPresentFlag, true)
	} else {
		TFConfigurationOptPresentFlag = append(TFConfigurationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TFConfigurationOptPresentFlag = append(TFConfigurationOptPresentFlag, true)
	} else {
		TFConfigurationOptPresentFlag = append(TFConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TFConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(*(x.SSBFrequency), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.SSBSubcarrierSpacing), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -60, 50
	err = pd.WriteInteger(*(x.SSBTransmitPower), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 5
	err = pd.WriteEnumerated(*(x.SSBPeriodicity), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteInteger(*(x.SSBHalfFrameOffset), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 15
	err = pd.WriteInteger(*(x.SSBSFNOffset), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.SSBBurstPosition != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SSBBurstPosition.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SSBBurstPosition marshal failed")
		}
	}

	// optional field
	if x.SFNInitialisationTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SFNInitialisationTime.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SFNInitialisationTime marshal failed")
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

func (x *TFConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TFConfigurationOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&TFConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	x.SSBFrequency = new(int64)
	*(x.SSBFrequency), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.SSBSubcarrierSpacing = new(aper.Enumerated)
	*(x.SSBSubcarrierSpacing), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -60, 50
	x.SSBTransmitPower = new(int64)
	*(x.SSBTransmitPower), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 5
	x.SSBPeriodicity = new(aper.Enumerated)
	*(x.SSBPeriodicity), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1
	x.SSBHalfFrameOffset = new(int64)
	*(x.SSBHalfFrameOffset), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 15
	x.SSBSFNOffset = new(int64)
	*(x.SSBSFNOffset), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if TFConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SSBBurstPosition = new(SSBBurstPosition)
		err = x.SSBBurstPosition.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSBBurstPosition error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TFConfigurationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SFNInitialisationTime = new(RelativeTime1900)
		err = x.SFNInitialisationTime.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SFNInitialisationTime error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if TFConfigurationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTFConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
