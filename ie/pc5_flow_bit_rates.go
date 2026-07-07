package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PC5FlowBitRates struct {
	GuaranteedFlowBitRate *BitRate
	MaximumFlowBitRate    *BitRate
	IEExtensions          *ProtocolExtensionContainerPC5FlowBitRatesExtIEs // optional
}

func (x *PC5FlowBitRates) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PC5FlowBitRatesOptPresentFlag := []bool{}
	// mandatory field
	if x.GuaranteedFlowBitRate == nil {
		return errors.Errorf("GuaranteedFlowBitRate is missing")
	}
	// mandatory field
	if x.MaximumFlowBitRate == nil {
		return errors.Errorf("MaximumFlowBitRate is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PC5FlowBitRatesOptPresentFlag = append(PC5FlowBitRatesOptPresentFlag, true)
	} else {
		PC5FlowBitRatesOptPresentFlag = append(PC5FlowBitRatesOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PC5FlowBitRatesOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.GuaranteedFlowBitRate.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GuaranteedFlowBitRate marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MaximumFlowBitRate.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MaximumFlowBitRate marshal failed")
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

func (x *PC5FlowBitRates) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PC5FlowBitRatesOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PC5FlowBitRatesOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GuaranteedFlowBitRate = new(BitRate)
	err = x.GuaranteedFlowBitRate.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GuaranteedFlowBitRate error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MaximumFlowBitRate = new(BitRate)
	err = x.MaximumFlowBitRate.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MaximumFlowBitRate error")
	}

	// optional field (optPresentFlag index: 0)
	if PC5FlowBitRatesOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPC5FlowBitRatesExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
