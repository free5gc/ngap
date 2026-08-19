package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PC5QoSFlowItem struct {
	PQI             *FiveQI
	Pc5FlowBitRates *PC5FlowBitRates                                // valueExt,optional
	Range           *Range                                          // valueExt,valueLB:0,valueUB:8,optional
	IEExtensions    *ProtocolExtensionContainerPC5QoSFlowItemExtIEs // optional
}

func (x *PC5QoSFlowItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PC5QoSFlowItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PQI == nil {
		return errors.Errorf("PQI is missing")
	}
	// optional field
	if x.Pc5FlowBitRates != nil {
		PC5QoSFlowItemOptPresentFlag = append(PC5QoSFlowItemOptPresentFlag, true)
	} else {
		PC5QoSFlowItemOptPresentFlag = append(PC5QoSFlowItemOptPresentFlag, false)
	}
	// optional field
	if x.Range != nil {
		PC5QoSFlowItemOptPresentFlag = append(PC5QoSFlowItemOptPresentFlag, true)
	} else {
		PC5QoSFlowItemOptPresentFlag = append(PC5QoSFlowItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PC5QoSFlowItemOptPresentFlag = append(PC5QoSFlowItemOptPresentFlag, true)
	} else {
		PC5QoSFlowItemOptPresentFlag = append(PC5QoSFlowItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PC5QoSFlowItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PQI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PQI marshal failed")
	}

	// optional field
	if x.Pc5FlowBitRates != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Pc5FlowBitRates.Write(pd)
		if err != nil {
			return errors.Wrap(err, "Pc5FlowBitRates marshal failed")
		}
	}

	// optional field
	if x.Range != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Range.Write(pd)
		if err != nil {
			return errors.Wrap(err, "Range marshal failed")
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

func (x *PC5QoSFlowItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PC5QoSFlowItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PC5QoSFlowItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PQI = new(FiveQI)
	err = x.PQI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PQI error")
	}

	// optional field (optPresentFlag index: 0)
	if PC5QoSFlowItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.Pc5FlowBitRates = new(PC5FlowBitRates)
		err = x.Pc5FlowBitRates.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode Pc5FlowBitRates error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PC5QoSFlowItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.Range = new(Range)
		err = x.Range.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode Range error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PC5QoSFlowItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPC5QoSFlowItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
