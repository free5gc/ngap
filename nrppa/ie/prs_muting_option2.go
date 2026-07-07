package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSMutingOption2 struct {
	MutingPattern *DLPRSMutingPattern                               // valueLB:0,valueUB:6
	IEExtensions  *ProtocolExtensionContainerPRSMutingOption2ExtIEs // optional
}

func (x *PRSMutingOption2) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingOption2OptPresentFlag := []bool{}
	// mandatory field
	if x.MutingPattern == nil {
		return errors.Errorf("MutingPattern is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSMutingOption2OptPresentFlag = append(PRSMutingOption2OptPresentFlag, true)
	} else {
		PRSMutingOption2OptPresentFlag = append(PRSMutingOption2OptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingOption2OptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MutingPattern.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MutingPattern marshal failed")
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

func (x *PRSMutingOption2) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingOption2OptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingOption2OptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if PRSMutingOption2OptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSMutingOption2ExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
