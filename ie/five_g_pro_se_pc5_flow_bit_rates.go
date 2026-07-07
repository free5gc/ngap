package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type FiveGProSePC5FlowBitRates struct {
	FiveGproSeguaranteedFlowBitRate *BitRate
	FiveGproSemaximumFlowBitRate    *BitRate
	IEExtensions                    *ProtocolExtensionContainerFiveGProSePC5FlowBitRatesExtIEs // optional
}

func (x *FiveGProSePC5FlowBitRates) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FiveGProSePC5FlowBitRatesOptPresentFlag := []bool{}
	// mandatory field
	if x.FiveGproSeguaranteedFlowBitRate == nil {
		return errors.Errorf("FiveGproSeguaranteedFlowBitRate is missing")
	}
	// mandatory field
	if x.FiveGproSemaximumFlowBitRate == nil {
		return errors.Errorf("FiveGproSemaximumFlowBitRate is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		FiveGProSePC5FlowBitRatesOptPresentFlag = append(FiveGProSePC5FlowBitRatesOptPresentFlag, true)
	} else {
		FiveGProSePC5FlowBitRatesOptPresentFlag = append(FiveGProSePC5FlowBitRatesOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FiveGProSePC5FlowBitRatesOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.FiveGproSeguaranteedFlowBitRate.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveGproSeguaranteedFlowBitRate marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.FiveGproSemaximumFlowBitRate.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveGproSemaximumFlowBitRate marshal failed")
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

func (x *FiveGProSePC5FlowBitRates) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FiveGProSePC5FlowBitRatesOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&FiveGProSePC5FlowBitRatesOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveGproSeguaranteedFlowBitRate = new(BitRate)
	err = x.FiveGproSeguaranteedFlowBitRate.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveGproSeguaranteedFlowBitRate error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveGproSemaximumFlowBitRate = new(BitRate)
	err = x.FiveGproSemaximumFlowBitRate.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveGproSemaximumFlowBitRate error")
	}

	// optional field (optPresentFlag index: 0)
	if FiveGProSePC5FlowBitRatesOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFiveGProSePC5FlowBitRatesExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
