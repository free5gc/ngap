package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TABasedMDT struct {
	TAListforMDT *TAListforMDT
	IEExtensions *ProtocolExtensionContainerTABasedMDTExtIEs // optional
}

func (x *TABasedMDT) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TABasedMDTOptPresentFlag := []bool{}
	// mandatory field
	if x.TAListforMDT == nil {
		return errors.Errorf("TAListforMDT is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TABasedMDTOptPresentFlag = append(TABasedMDTOptPresentFlag, true)
	} else {
		TABasedMDTOptPresentFlag = append(TABasedMDTOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TABasedMDTOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TAListforMDT.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAListforMDT marshal failed")
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

func (x *TABasedMDT) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TABasedMDTOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TABasedMDTOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAListforMDT = new(TAListforMDT)
	err = x.TAListforMDT.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAListforMDT error")
	}

	// optional field (optPresentFlag index: 0)
	if TABasedMDTOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTABasedMDTExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
