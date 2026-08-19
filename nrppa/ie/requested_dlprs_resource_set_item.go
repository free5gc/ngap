package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	RequestedDLPRSResourceSetItemCombSizePresentN2  aper.Enumerated = 0
	RequestedDLPRSResourceSetItemCombSizePresentN4  aper.Enumerated = 1
	RequestedDLPRSResourceSetItemCombSizePresentN6  aper.Enumerated = 2
	RequestedDLPRSResourceSetItemCombSizePresentN12 aper.Enumerated = 3
)

const ( /* Enum Type */
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN4     aper.Enumerated = 0
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN5     aper.Enumerated = 1
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN8     aper.Enumerated = 2
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN10    aper.Enumerated = 3
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN16    aper.Enumerated = 4
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN20    aper.Enumerated = 5
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN32    aper.Enumerated = 6
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN40    aper.Enumerated = 7
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN64    aper.Enumerated = 8
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN80    aper.Enumerated = 9
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN160   aper.Enumerated = 10
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN320   aper.Enumerated = 11
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN640   aper.Enumerated = 12
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN1280  aper.Enumerated = 13
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN2560  aper.Enumerated = 14
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN5120  aper.Enumerated = 15
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN10240 aper.Enumerated = 16
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN20480 aper.Enumerated = 17
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN40960 aper.Enumerated = 18
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN81920 aper.Enumerated = 19
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN128   aper.Enumerated = 20
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN256   aper.Enumerated = 21
	RequestedDLPRSResourceSetItemResourceSetPeriodicityPresentN512   aper.Enumerated = 22
)

const ( /* Enum Type */
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf1  aper.Enumerated = 0
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf2  aper.Enumerated = 1
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf4  aper.Enumerated = 2
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf6  aper.Enumerated = 3
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf8  aper.Enumerated = 4
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf16 aper.Enumerated = 5
	RequestedDLPRSResourceSetItemResourceRepetitionFactorPresentRf32 aper.Enumerated = 6
)

const ( /* Enum Type */
	RequestedDLPRSResourceSetItemResourceNumberofSymbolsPresentN2  aper.Enumerated = 0
	RequestedDLPRSResourceSetItemResourceNumberofSymbolsPresentN4  aper.Enumerated = 1
	RequestedDLPRSResourceSetItemResourceNumberofSymbolsPresentN6  aper.Enumerated = 2
	RequestedDLPRSResourceSetItemResourceNumberofSymbolsPresentN12 aper.Enumerated = 3
)

type RequestedDLPRSResourceSetItem struct {
	PRSbandwidth                    *int64                                                         // valueLB:1,valueUB:63,optional
	CombSize                        *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:3,optional
	ResourceSetPeriodicity          *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:19,optional
	ResourceRepetitionFactor        *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:6,optional
	ResourceNumberofSymbols         *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:3,optional
	RequestedDLPRSResourceList      *RequestedDLPRSResourceList                                    // optional
	ResourceSetStartTimeAndDuration *StartTimeAndDuration                                          // valueExt,optional
	IEExtensions                    *ProtocolExtensionContainerRequestedDLPRSResourceSetItemExtIEs // optional
}

func (x *RequestedDLPRSResourceSetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedDLPRSResourceSetItemOptPresentFlag := []bool{}
	// optional field
	if x.PRSbandwidth != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.CombSize != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.ResourceSetPeriodicity != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.ResourceRepetitionFactor != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.ResourceNumberofSymbols != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.RequestedDLPRSResourceList != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.ResourceSetStartTimeAndDuration != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceSetItemOptPresentFlag = append(RequestedDLPRSResourceSetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RequestedDLPRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PRSbandwidth != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 63
		err = pd.WriteInteger(*(x.PRSbandwidth), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.CombSize != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 3
		err = pd.WriteEnumerated(*(x.CombSize), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.ResourceSetPeriodicity != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 19
		err = pd.WriteEnumerated(*(x.ResourceSetPeriodicity), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.ResourceRepetitionFactor != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 6
		err = pd.WriteEnumerated(*(x.ResourceRepetitionFactor), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.ResourceNumberofSymbols != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 3
		err = pd.WriteEnumerated(*(x.ResourceNumberofSymbols), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.RequestedDLPRSResourceList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RequestedDLPRSResourceList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RequestedDLPRSResourceList marshal failed")
		}
	}

	// optional field
	if x.ResourceSetStartTimeAndDuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResourceSetStartTimeAndDuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ResourceSetStartTimeAndDuration marshal failed")
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

func (x *RequestedDLPRSResourceSetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedDLPRSResourceSetItemOptPresentFlag := make([]bool, 8)
	err = pd.ReadSequencePreambleBitMap(&RequestedDLPRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if RequestedDLPRSResourceSetItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 63
		x.PRSbandwidth = new(int64)
		*(x.PRSbandwidth), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if RequestedDLPRSResourceSetItemOptPresentFlag[1] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 3
		x.CombSize = new(aper.Enumerated)
		*(x.CombSize), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if RequestedDLPRSResourceSetItemOptPresentFlag[2] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 19
		x.ResourceSetPeriodicity = new(aper.Enumerated)
		*(x.ResourceSetPeriodicity), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if RequestedDLPRSResourceSetItemOptPresentFlag[3] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 6
		x.ResourceRepetitionFactor = new(aper.Enumerated)
		*(x.ResourceRepetitionFactor), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 4)
	if RequestedDLPRSResourceSetItemOptPresentFlag[4] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 3
		x.ResourceNumberofSymbols = new(aper.Enumerated)
		*(x.ResourceNumberofSymbols), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 5)
	if RequestedDLPRSResourceSetItemOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.RequestedDLPRSResourceList = new(RequestedDLPRSResourceList)
		err = x.RequestedDLPRSResourceList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RequestedDLPRSResourceList error")
		}
	}

	// optional field (optPresentFlag index: 6)
	if RequestedDLPRSResourceSetItemOptPresentFlag[6] {
		// Read struct defined elsewhere (Pointer)
		x.ResourceSetStartTimeAndDuration = new(StartTimeAndDuration)
		err = x.ResourceSetStartTimeAndDuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ResourceSetStartTimeAndDuration error")
		}
	}

	// optional field (optPresentFlag index: 7)
	if RequestedDLPRSResourceSetItemOptPresentFlag[7] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRequestedDLPRSResourceSetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
