package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type Dynamic5QIDescriptor struct {
	PriorityLevelQos       *PriorityLevelQos
	PacketDelayBudget      *PacketDelayBudget
	PacketErrorRate        *PacketErrorRate                                      // valueExt
	FiveQI                 *FiveQI                                               // optional
	DelayCritical          *DelayCritical                                        // valueExt,valueLB:0,valueUB:1,optional
	AveragingWindow        *AveragingWindow                                      // optional
	MaximumDataBurstVolume *MaximumDataBurstVolume                               // optional
	IEExtensions           *ProtocolExtensionContainerDynamic5QIDescriptorExtIEs // optional
}

func (x *Dynamic5QIDescriptor) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	Dynamic5QIDescriptorOptPresentFlag := []bool{}
	// mandatory field
	if x.PriorityLevelQos == nil {
		return errors.Errorf("PriorityLevelQos is missing")
	}
	// mandatory field
	if x.PacketDelayBudget == nil {
		return errors.Errorf("PacketDelayBudget is missing")
	}
	// mandatory field
	if x.PacketErrorRate == nil {
		return errors.Errorf("PacketErrorRate is missing")
	}
	// optional field
	if x.FiveQI != nil {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, true)
	} else {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.DelayCritical != nil {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, true)
	} else {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.AveragingWindow != nil {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, true)
	} else {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.MaximumDataBurstVolume != nil {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, true)
	} else {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, true)
	} else {
		Dynamic5QIDescriptorOptPresentFlag = append(Dynamic5QIDescriptorOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(Dynamic5QIDescriptorOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PriorityLevelQos.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PriorityLevelQos marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PacketDelayBudget.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PacketDelayBudget marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PacketErrorRate.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PacketErrorRate marshal failed")
	}

	// optional field
	if x.FiveQI != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveQI.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveQI marshal failed")
		}
	}

	// optional field
	if x.DelayCritical != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DelayCritical.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DelayCritical marshal failed")
		}
	}

	// optional field
	if x.AveragingWindow != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AveragingWindow.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AveragingWindow marshal failed")
		}
	}

	// optional field
	if x.MaximumDataBurstVolume != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MaximumDataBurstVolume.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MaximumDataBurstVolume marshal failed")
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

func (x *Dynamic5QIDescriptor) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	Dynamic5QIDescriptorOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&Dynamic5QIDescriptorOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PriorityLevelQos = new(PriorityLevelQos)
	err = x.PriorityLevelQos.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PriorityLevelQos error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PacketDelayBudget = new(PacketDelayBudget)
	err = x.PacketDelayBudget.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PacketDelayBudget error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PacketErrorRate = new(PacketErrorRate)
	err = x.PacketErrorRate.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PacketErrorRate error")
	}

	// optional field (optPresentFlag index: 0)
	if Dynamic5QIDescriptorOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.FiveQI = new(FiveQI)
		err = x.FiveQI.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveQI error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if Dynamic5QIDescriptorOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.DelayCritical = new(DelayCritical)
		err = x.DelayCritical.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DelayCritical error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if Dynamic5QIDescriptorOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.AveragingWindow = new(AveragingWindow)
		err = x.AveragingWindow.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AveragingWindow error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if Dynamic5QIDescriptorOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.MaximumDataBurstVolume = new(MaximumDataBurstVolume)
		err = x.MaximumDataBurstVolume.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MaximumDataBurstVolume error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if Dynamic5QIDescriptorOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDynamic5QIDescriptorExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
