package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SupportedTAItem struct {
	TAC               *TAC
	BroadcastPLMNList *BroadcastPLMNList
	IEExtensions      *ProtocolExtensionContainerSupportedTAItemExtIEs // optional
}

func (x *SupportedTAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SupportedTAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TAC == nil {
		return errors.Errorf("TAC is missing")
	}
	// mandatory field
	if x.BroadcastPLMNList == nil {
		return errors.Errorf("BroadcastPLMNList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SupportedTAItemOptPresentFlag = append(SupportedTAItemOptPresentFlag, true)
	} else {
		SupportedTAItemOptPresentFlag = append(SupportedTAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SupportedTAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TAC.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TAC marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.BroadcastPLMNList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "BroadcastPLMNList marshal failed")
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

func (x *SupportedTAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SupportedTAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SupportedTAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TAC = new(TAC)
	err = x.TAC.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TAC error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.BroadcastPLMNList = new(BroadcastPLMNList)
	err = x.BroadcastPLMNList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode BroadcastPLMNList error")
	}

	// optional field (optPresentFlag index: 0)
	if SupportedTAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSupportedTAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
