package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSMuting struct {
	PRSMutingOption1 *PRSMutingOption1                          // valueExt,optional
	PRSMutingOption2 *PRSMutingOption2                          // valueExt,optional
	IEExtensions     *ProtocolExtensionContainerPRSMutingExtIEs // optional
}

func (x *PRSMuting) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingOptPresentFlag := []bool{}
	// optional field
	if x.PRSMutingOption1 != nil {
		PRSMutingOptPresentFlag = append(PRSMutingOptPresentFlag, true)
	} else {
		PRSMutingOptPresentFlag = append(PRSMutingOptPresentFlag, false)
	}
	// optional field
	if x.PRSMutingOption2 != nil {
		PRSMutingOptPresentFlag = append(PRSMutingOptPresentFlag, true)
	} else {
		PRSMutingOptPresentFlag = append(PRSMutingOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PRSMutingOptPresentFlag = append(PRSMutingOptPresentFlag, true)
	} else {
		PRSMutingOptPresentFlag = append(PRSMutingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PRSMutingOption1 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSMutingOption1.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSMutingOption1 marshal failed")
		}
	}

	// optional field
	if x.PRSMutingOption2 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSMutingOption2.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSMutingOption2 marshal failed")
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

func (x *PRSMuting) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if PRSMutingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PRSMutingOption1 = new(PRSMutingOption1)
		err = x.PRSMutingOption1.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSMutingOption1 error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PRSMutingOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PRSMutingOption2 = new(PRSMutingOption2)
		err = x.PRSMutingOption2.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSMutingOption2 error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PRSMutingOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSMutingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
