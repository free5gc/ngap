package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type FiveGProSePC5QoSFlowItem struct {
	FiveGproSepQI             *FiveQI
	FiveGproSepc5FlowBitRates *FiveGProSePC5FlowBitRates                                // valueExt,optional
	FiveGproSerange           *Range                                                    // valueExt,valueLB:0,valueUB:8,optional
	IEExtensions              *ProtocolExtensionContainerFiveGProSePC5QoSFlowItemExtIEs // optional
}

func (x *FiveGProSePC5QoSFlowItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FiveGProSePC5QoSFlowItemOptPresentFlag := []bool{}
	// mandatory field
	if x.FiveGproSepQI == nil {
		return errors.Errorf("FiveGproSepQI is missing")
	}
	// optional field
	if x.FiveGproSepc5FlowBitRates != nil {
		FiveGProSePC5QoSFlowItemOptPresentFlag = append(FiveGProSePC5QoSFlowItemOptPresentFlag, true)
	} else {
		FiveGProSePC5QoSFlowItemOptPresentFlag = append(FiveGProSePC5QoSFlowItemOptPresentFlag, false)
	}
	// optional field
	if x.FiveGproSerange != nil {
		FiveGProSePC5QoSFlowItemOptPresentFlag = append(FiveGProSePC5QoSFlowItemOptPresentFlag, true)
	} else {
		FiveGProSePC5QoSFlowItemOptPresentFlag = append(FiveGProSePC5QoSFlowItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		FiveGProSePC5QoSFlowItemOptPresentFlag = append(FiveGProSePC5QoSFlowItemOptPresentFlag, true)
	} else {
		FiveGProSePC5QoSFlowItemOptPresentFlag = append(FiveGProSePC5QoSFlowItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FiveGProSePC5QoSFlowItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.FiveGproSepQI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveGproSepQI marshal failed")
	}

	// optional field
	if x.FiveGproSepc5FlowBitRates != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGproSepc5FlowBitRates.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGproSepc5FlowBitRates marshal failed")
		}
	}

	// optional field
	if x.FiveGproSerange != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGproSerange.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGproSerange marshal failed")
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

func (x *FiveGProSePC5QoSFlowItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FiveGProSePC5QoSFlowItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&FiveGProSePC5QoSFlowItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveGproSepQI = new(FiveQI)
	err = x.FiveGproSepQI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveGproSepQI error")
	}

	// optional field (optPresentFlag index: 0)
	if FiveGProSePC5QoSFlowItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGproSepc5FlowBitRates = new(FiveGProSePC5FlowBitRates)
		err = x.FiveGproSepc5FlowBitRates.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGproSepc5FlowBitRates error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if FiveGProSePC5QoSFlowItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGproSerange = new(Range)
		err = x.FiveGproSerange.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGproSerange error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if FiveGProSePC5QoSFlowItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFiveGProSePC5QoSFlowItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
