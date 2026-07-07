package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PathlossReferenceInformation struct {
	PathlossReferenceSignal *PathlossReferenceSignal                                      // valueLB:0,valueUB:2
	IEExtensions            *ProtocolExtensionContainerPathlossReferenceInformationExtIEs // optional
}

func (x *PathlossReferenceInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PathlossReferenceInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.PathlossReferenceSignal == nil {
		return errors.Errorf("PathlossReferenceSignal is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PathlossReferenceInformationOptPresentFlag = append(PathlossReferenceInformationOptPresentFlag, true)
	} else {
		PathlossReferenceInformationOptPresentFlag = append(PathlossReferenceInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PathlossReferenceInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PathlossReferenceSignal.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PathlossReferenceSignal marshal failed")
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

func (x *PathlossReferenceInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PathlossReferenceInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PathlossReferenceInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PathlossReferenceSignal = new(PathlossReferenceSignal)
	err = x.PathlossReferenceSignal.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PathlossReferenceSignal error")
	}

	// optional field (optPresentFlag index: 0)
	if PathlossReferenceInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPathlossReferenceInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
