package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AlternativeQoSParaSetItem struct {
	AlternativeQoSParaSetIndex *AlternativeQoSParaSetIndex
	GuaranteedFlowBitRateDL    *BitRate                                                   // optional
	GuaranteedFlowBitRateUL    *BitRate                                                   // optional
	PacketDelayBudget          *PacketDelayBudget                                         // optional
	PacketErrorRate            *PacketErrorRate                                           // valueExt,optional
	IEExtensions               *ProtocolExtensionContainerAlternativeQoSParaSetItemExtIEs // optional
}

func (x *AlternativeQoSParaSetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AlternativeQoSParaSetItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AlternativeQoSParaSetIndex == nil {
		return errors.Errorf("AlternativeQoSParaSetIndex is missing")
	}
	// optional field
	if x.GuaranteedFlowBitRateDL != nil {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, true)
	} else {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, false)
	}
	// optional field
	if x.GuaranteedFlowBitRateUL != nil {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, true)
	} else {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, false)
	}
	// optional field
	if x.PacketDelayBudget != nil {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, true)
	} else {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, false)
	}
	// optional field
	if x.PacketErrorRate != nil {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, true)
	} else {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, true)
	} else {
		AlternativeQoSParaSetItemOptPresentFlag = append(AlternativeQoSParaSetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AlternativeQoSParaSetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AlternativeQoSParaSetIndex.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AlternativeQoSParaSetIndex marshal failed")
	}

	// optional field
	if x.GuaranteedFlowBitRateDL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.GuaranteedFlowBitRateDL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "GuaranteedFlowBitRateDL marshal failed")
		}
	}

	// optional field
	if x.GuaranteedFlowBitRateUL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.GuaranteedFlowBitRateUL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "GuaranteedFlowBitRateUL marshal failed")
		}
	}

	// optional field
	if x.PacketDelayBudget != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PacketDelayBudget.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PacketDelayBudget marshal failed")
		}
	}

	// optional field
	if x.PacketErrorRate != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PacketErrorRate.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PacketErrorRate marshal failed")
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

func (x *AlternativeQoSParaSetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AlternativeQoSParaSetItemOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&AlternativeQoSParaSetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AlternativeQoSParaSetIndex = new(AlternativeQoSParaSetIndex)
	err = x.AlternativeQoSParaSetIndex.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AlternativeQoSParaSetIndex error")
	}

	// optional field (optPresentFlag index: 0)
	if AlternativeQoSParaSetItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.GuaranteedFlowBitRateDL = new(BitRate)
		err = x.GuaranteedFlowBitRateDL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode GuaranteedFlowBitRateDL error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AlternativeQoSParaSetItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.GuaranteedFlowBitRateUL = new(BitRate)
		err = x.GuaranteedFlowBitRateUL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode GuaranteedFlowBitRateUL error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if AlternativeQoSParaSetItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.PacketDelayBudget = new(PacketDelayBudget)
		err = x.PacketDelayBudget.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PacketDelayBudget error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if AlternativeQoSParaSetItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.PacketErrorRate = new(PacketErrorRate)
		err = x.PacketErrorRate.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PacketErrorRate error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if AlternativeQoSParaSetItemOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAlternativeQoSParaSetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
