package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PRSResourceSetItemSubcarrierSpacingPresentKHz15  aper.Enumerated = 0
	PRSResourceSetItemSubcarrierSpacingPresentKHz30  aper.Enumerated = 1
	PRSResourceSetItemSubcarrierSpacingPresentKHz60  aper.Enumerated = 2
	PRSResourceSetItemSubcarrierSpacingPresentKHz120 aper.Enumerated = 3
)

const ( /* Enum Type */
	PRSResourceSetItemCombSizePresentN2  aper.Enumerated = 0
	PRSResourceSetItemCombSizePresentN4  aper.Enumerated = 1
	PRSResourceSetItemCombSizePresentN6  aper.Enumerated = 2
	PRSResourceSetItemCombSizePresentN12 aper.Enumerated = 3
)

const ( /* Enum Type */
	PRSResourceSetItemCPTypePresentNormal   aper.Enumerated = 0
	PRSResourceSetItemCPTypePresentExtended aper.Enumerated = 1
)

const ( /* Enum Type */
	PRSResourceSetItemResourceSetPeriodicityPresentN4     aper.Enumerated = 0
	PRSResourceSetItemResourceSetPeriodicityPresentN5     aper.Enumerated = 1
	PRSResourceSetItemResourceSetPeriodicityPresentN8     aper.Enumerated = 2
	PRSResourceSetItemResourceSetPeriodicityPresentN10    aper.Enumerated = 3
	PRSResourceSetItemResourceSetPeriodicityPresentN16    aper.Enumerated = 4
	PRSResourceSetItemResourceSetPeriodicityPresentN20    aper.Enumerated = 5
	PRSResourceSetItemResourceSetPeriodicityPresentN32    aper.Enumerated = 6
	PRSResourceSetItemResourceSetPeriodicityPresentN40    aper.Enumerated = 7
	PRSResourceSetItemResourceSetPeriodicityPresentN64    aper.Enumerated = 8
	PRSResourceSetItemResourceSetPeriodicityPresentN80    aper.Enumerated = 9
	PRSResourceSetItemResourceSetPeriodicityPresentN160   aper.Enumerated = 10
	PRSResourceSetItemResourceSetPeriodicityPresentN320   aper.Enumerated = 11
	PRSResourceSetItemResourceSetPeriodicityPresentN640   aper.Enumerated = 12
	PRSResourceSetItemResourceSetPeriodicityPresentN1280  aper.Enumerated = 13
	PRSResourceSetItemResourceSetPeriodicityPresentN2560  aper.Enumerated = 14
	PRSResourceSetItemResourceSetPeriodicityPresentN5120  aper.Enumerated = 15
	PRSResourceSetItemResourceSetPeriodicityPresentN10240 aper.Enumerated = 16
	PRSResourceSetItemResourceSetPeriodicityPresentN20480 aper.Enumerated = 17
	PRSResourceSetItemResourceSetPeriodicityPresentN40960 aper.Enumerated = 18
	PRSResourceSetItemResourceSetPeriodicityPresentN81920 aper.Enumerated = 19
	PRSResourceSetItemResourceSetPeriodicityPresentN128   aper.Enumerated = 20
	PRSResourceSetItemResourceSetPeriodicityPresentN256   aper.Enumerated = 21
	PRSResourceSetItemResourceSetPeriodicityPresentN512   aper.Enumerated = 22
)

const ( /* Enum Type */
	PRSResourceSetItemResourceRepetitionFactorPresentRf1  aper.Enumerated = 0
	PRSResourceSetItemResourceRepetitionFactorPresentRf2  aper.Enumerated = 1
	PRSResourceSetItemResourceRepetitionFactorPresentRf4  aper.Enumerated = 2
	PRSResourceSetItemResourceRepetitionFactorPresentRf6  aper.Enumerated = 3
	PRSResourceSetItemResourceRepetitionFactorPresentRf8  aper.Enumerated = 4
	PRSResourceSetItemResourceRepetitionFactorPresentRf16 aper.Enumerated = 5
	PRSResourceSetItemResourceRepetitionFactorPresentRf32 aper.Enumerated = 6
)

const ( /* Enum Type */
	PRSResourceSetItemResourceTimeGapPresentTg1  aper.Enumerated = 0
	PRSResourceSetItemResourceTimeGapPresentTg2  aper.Enumerated = 1
	PRSResourceSetItemResourceTimeGapPresentTg4  aper.Enumerated = 2
	PRSResourceSetItemResourceTimeGapPresentTg8  aper.Enumerated = 3
	PRSResourceSetItemResourceTimeGapPresentTg16 aper.Enumerated = 4
	PRSResourceSetItemResourceTimeGapPresentTg32 aper.Enumerated = 5
)

const ( /* Enum Type */
	PRSResourceSetItemResourceNumberofSymbolsPresentN2  aper.Enumerated = 0
	PRSResourceSetItemResourceNumberofSymbolsPresentN4  aper.Enumerated = 1
	PRSResourceSetItemResourceNumberofSymbolsPresentN6  aper.Enumerated = 2
	PRSResourceSetItemResourceNumberofSymbolsPresentN12 aper.Enumerated = 3
)

type PRSResourceSetItem struct {
	PRSResourceSetID         *PRSResourceSetID
	SubcarrierSpacing        *aper.Enumerated // valueExt,valueLB:0,valueUB:3
	PRSbandwidth             *int64           // valueLB:1,valueUB:63
	StartPRB                 *int64           // valueLB:0,valueUB:2176
	PointA                   *int64           // valueLB:0,valueUB:3279165
	CombSize                 *aper.Enumerated // valueExt,valueLB:0,valueUB:3
	CPType                   *aper.Enumerated // valueExt,valueLB:0,valueUB:1
	ResourceSetPeriodicity   *aper.Enumerated // valueExt,valueLB:0,valueUB:19
	ResourceSetSlotOffset    *int64           // valueExt,valueLB:0,valueUB:81919
	ResourceRepetitionFactor *aper.Enumerated // valueExt,valueLB:0,valueUB:6
	ResourceTimeGap          *aper.Enumerated // valueExt,valueLB:0,valueUB:5
	ResourceNumberofSymbols  *aper.Enumerated // valueExt,valueLB:0,valueUB:3
	PRSMuting                *PRSMuting       // valueExt,optional
	PRSResourceTransmitPower *int64           // valueLB:-60,valueUB:50
	PRSResourceList          *PRSResourceList
	IEExtensions             *ProtocolExtensionContainerPRSResourceSetItemExtIEs // optional
}

func (x *PRSResourceSetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceSetItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSResourceSetID == nil {
		return errors.Errorf("PRSResourceSetID is missing")
	}
	// mandatory field
	if x.SubcarrierSpacing == nil {
		return errors.Errorf("SubcarrierSpacing is missing")
	}
	// mandatory field
	if x.PRSbandwidth == nil {
		return errors.Errorf("PRSbandwidth is missing")
	}
	// mandatory field
	if x.StartPRB == nil {
		return errors.Errorf("StartPRB is missing")
	}
	// mandatory field
	if x.PointA == nil {
		return errors.Errorf("PointA is missing")
	}
	// mandatory field
	if x.CombSize == nil {
		return errors.Errorf("CombSize is missing")
	}
	// mandatory field
	if x.CPType == nil {
		return errors.Errorf("CPType is missing")
	}
	// mandatory field
	if x.ResourceSetPeriodicity == nil {
		return errors.Errorf("ResourceSetPeriodicity is missing")
	}
	// mandatory field
	if x.ResourceSetSlotOffset == nil {
		return errors.Errorf("ResourceSetSlotOffset is missing")
	}
	// mandatory field
	if x.ResourceRepetitionFactor == nil {
		return errors.Errorf("ResourceRepetitionFactor is missing")
	}
	// mandatory field
	if x.ResourceTimeGap == nil {
		return errors.Errorf("ResourceTimeGap is missing")
	}
	// mandatory field
	if x.ResourceNumberofSymbols == nil {
		return errors.Errorf("ResourceNumberofSymbols is missing")
	}
	// optional field
	if x.PRSMuting != nil {
		PRSResourceSetItemOptPresentFlag = append(PRSResourceSetItemOptPresentFlag, true)
	} else {
		PRSResourceSetItemOptPresentFlag = append(PRSResourceSetItemOptPresentFlag, false)
	}
	// mandatory field
	if x.PRSResourceTransmitPower == nil {
		return errors.Errorf("PRSResourceTransmitPower is missing")
	}
	// mandatory field
	if x.PRSResourceList == nil {
		return errors.Errorf("PRSResourceList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSResourceSetItemOptPresentFlag = append(PRSResourceSetItemOptPresentFlag, true)
	} else {
		PRSResourceSetItemOptPresentFlag = append(PRSResourceSetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceSetID marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.SubcarrierSpacing), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 1, 63
	err = pd.WriteInteger(*(x.PRSbandwidth), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 2176
	err = pd.WriteInteger(*(x.StartPRB), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(*(x.PointA), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.CombSize), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.CPType), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 19
	err = pd.WriteEnumerated(*(x.ResourceSetPeriodicity), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 81919
	err = pd.WriteInteger(*(x.ResourceSetSlotOffset), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 6
	err = pd.WriteEnumerated(*(x.ResourceRepetitionFactor), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 5
	err = pd.WriteEnumerated(*(x.ResourceTimeGap), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.ResourceNumberofSymbols), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// optional field
	if x.PRSMuting != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSMuting.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSMuting marshal failed")
		}
	}

	// Write Integer (Pointer)
	*vLb, *vUb = -60, 50
	err = pd.WriteInteger(*(x.PRSResourceTransmitPower), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceList marshal failed")
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

func (x *PRSResourceSetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceSetItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceSetID = new(PRSResourceSetID)
	err = x.PRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceSetID error")
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
	*vLb, *vUb = 1, 63
	x.PRSbandwidth = new(int64)
	*(x.PRSbandwidth), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 2176
	x.StartPRB = new(int64)
	*(x.StartPRB), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	x.PointA = new(int64)
	*(x.PointA), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.CombSize = new(aper.Enumerated)
	*(x.CombSize), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.CPType = new(aper.Enumerated)
	*(x.CPType), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 19
	x.ResourceSetPeriodicity = new(aper.Enumerated)
	*(x.ResourceSetPeriodicity), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 81919
	x.ResourceSetSlotOffset = new(int64)
	*(x.ResourceSetSlotOffset), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 6
	x.ResourceRepetitionFactor = new(aper.Enumerated)
	*(x.ResourceRepetitionFactor), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 5
	x.ResourceTimeGap = new(aper.Enumerated)
	*(x.ResourceTimeGap), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.ResourceNumberofSymbols = new(aper.Enumerated)
	*(x.ResourceNumberofSymbols), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if PRSResourceSetItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PRSMuting = new(PRSMuting)
		err = x.PRSMuting.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSMuting error")
		}
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = -60, 50
	x.PRSResourceTransmitPower = new(int64)
	*(x.PRSResourceTransmitPower), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceList = new(PRSResourceList)
	err = x.PRSResourceList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceList error")
	}

	// optional field (optPresentFlag index: 1)
	if PRSResourceSetItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSResourceSetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
