package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NonDynamic5QIDescriptor struct {
	FiveQI                 *FiveQI
	PriorityLevelQos       *PriorityLevelQos                                        // optional
	AveragingWindow        *AveragingWindow                                         // optional
	MaximumDataBurstVolume *MaximumDataBurstVolume                                  // optional
	IEExtensions           *ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs // optional
}

func (x *NonDynamic5QIDescriptor) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NonDynamic5QIDescriptorOptPresentFlag := []bool{}
	// mandatory field
	if x.FiveQI == nil {
		return errors.Errorf("FiveQI is missing")
	}
	// optional field
	if x.PriorityLevelQos != nil {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, true)
	} else {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.AveragingWindow != nil {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, true)
	} else {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.MaximumDataBurstVolume != nil {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, true)
	} else {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, true)
	} else {
		NonDynamic5QIDescriptorOptPresentFlag = append(NonDynamic5QIDescriptorOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NonDynamic5QIDescriptorOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.FiveQI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveQI marshal failed")
	}

	// optional field
	if x.PriorityLevelQos != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PriorityLevelQos.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PriorityLevelQos marshal failed")
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

func (x *NonDynamic5QIDescriptor) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NonDynamic5QIDescriptorOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&NonDynamic5QIDescriptorOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveQI = new(FiveQI)
	err = x.FiveQI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveQI error")
	}

	// optional field (optPresentFlag index: 0)
	if NonDynamic5QIDescriptorOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PriorityLevelQos = new(PriorityLevelQos)
		err = x.PriorityLevelQos.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PriorityLevelQos error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if NonDynamic5QIDescriptorOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.AveragingWindow = new(AveragingWindow)
		err = x.AveragingWindow.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AveragingWindow error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if NonDynamic5QIDescriptorOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.MaximumDataBurstVolume = new(MaximumDataBurstVolume)
		err = x.MaximumDataBurstVolume.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MaximumDataBurstVolume error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if NonDynamic5QIDescriptorOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNonDynamic5QIDescriptorExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
