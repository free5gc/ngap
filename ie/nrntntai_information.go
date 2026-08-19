package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NRNTNTAIInformation struct {
	ServingPLMN                 *PLMNIdentity
	TACListInNRNTN              *TACListInNRNTN
	UELocationDerivedTACInNRNTN *TAC                                                 // optional
	IEExtensions                *ProtocolExtensionContainerNRNTNTAIInformationExtIEs // optional
}

func (x *NRNTNTAIInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRNTNTAIInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.ServingPLMN == nil {
		return errors.Errorf("ServingPLMN is missing")
	}
	// mandatory field
	if x.TACListInNRNTN == nil {
		return errors.Errorf("TACListInNRNTN is missing")
	}
	// optional field
	if x.UELocationDerivedTACInNRNTN != nil {
		NRNTNTAIInformationOptPresentFlag = append(NRNTNTAIInformationOptPresentFlag, true)
	} else {
		NRNTNTAIInformationOptPresentFlag = append(NRNTNTAIInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		NRNTNTAIInformationOptPresentFlag = append(NRNTNTAIInformationOptPresentFlag, true)
	} else {
		NRNTNTAIInformationOptPresentFlag = append(NRNTNTAIInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRNTNTAIInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ServingPLMN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ServingPLMN marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TACListInNRNTN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TACListInNRNTN marshal failed")
	}

	// optional field
	if x.UELocationDerivedTACInNRNTN != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UELocationDerivedTACInNRNTN.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UELocationDerivedTACInNRNTN marshal failed")
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

func (x *NRNTNTAIInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRNTNTAIInformationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&NRNTNTAIInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ServingPLMN = new(PLMNIdentity)
	err = x.ServingPLMN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ServingPLMN error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TACListInNRNTN = new(TACListInNRNTN)
	err = x.TACListInNRNTN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TACListInNRNTN error")
	}

	// optional field (optPresentFlag index: 0)
	if NRNTNTAIInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UELocationDerivedTACInNRNTN = new(TAC)
		err = x.UELocationDerivedTACInNRNTN.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UELocationDerivedTACInNRNTN error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if NRNTNTAIInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNRNTNTAIInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
