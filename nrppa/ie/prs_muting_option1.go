package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PRSMutingOption1MutingBitRepetitionFactorPresentN1 aper.Enumerated = 0
	PRSMutingOption1MutingBitRepetitionFactorPresentN2 aper.Enumerated = 1
	PRSMutingOption1MutingBitRepetitionFactorPresentN4 aper.Enumerated = 2
	PRSMutingOption1MutingBitRepetitionFactorPresentN8 aper.Enumerated = 3
)

type PRSMutingOption1 struct {
	MutingPattern             *DLPRSMutingPattern                               // valueLB:0,valueUB:6
	MutingBitRepetitionFactor *aper.Enumerated                                  // valueExt,valueLB:0,valueUB:3
	IEExtensions              *ProtocolExtensionContainerPRSMutingOption1ExtIEs // optional
}

func (x *PRSMutingOption1) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingOption1OptPresentFlag := []bool{}
	// mandatory field
	if x.MutingPattern == nil {
		return errors.Errorf("MutingPattern is missing")
	}
	// mandatory field
	if x.MutingBitRepetitionFactor == nil {
		return errors.Errorf("MutingBitRepetitionFactor is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSMutingOption1OptPresentFlag = append(PRSMutingOption1OptPresentFlag, true)
	} else {
		PRSMutingOption1OptPresentFlag = append(PRSMutingOption1OptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingOption1OptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MutingPattern.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MutingPattern marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteEnumerated(*(x.MutingBitRepetitionFactor), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *PRSMutingOption1) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingOption1OptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingOption1OptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MutingPattern = new(DLPRSMutingPattern)
	err = x.MutingPattern.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MutingPattern error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 3
	x.MutingBitRepetitionFactor = new(aper.Enumerated)
	*(x.MutingBitRepetitionFactor), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if PRSMutingOption1OptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSMutingOption1ExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
