package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AllocationAndRetentionPriority struct {
	PriorityLevelARP        *PriorityLevelARP
	PreEmptionCapability    *PreEmptionCapability                                           // valueExt,valueLB:0,valueUB:1
	PreEmptionVulnerability *PreEmptionVulnerability                                        // valueExt,valueLB:0,valueUB:1
	IEExtensions            *ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs // optional
}

func (x *AllocationAndRetentionPriority) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AllocationAndRetentionPriorityOptPresentFlag := []bool{}
	// mandatory field
	if x.PriorityLevelARP == nil {
		return errors.Errorf("PriorityLevelARP is missing")
	}
	// mandatory field
	if x.PreEmptionCapability == nil {
		return errors.Errorf("PreEmptionCapability is missing")
	}
	// mandatory field
	if x.PreEmptionVulnerability == nil {
		return errors.Errorf("PreEmptionVulnerability is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AllocationAndRetentionPriorityOptPresentFlag = append(AllocationAndRetentionPriorityOptPresentFlag, true)
	} else {
		AllocationAndRetentionPriorityOptPresentFlag = append(AllocationAndRetentionPriorityOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AllocationAndRetentionPriorityOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PriorityLevelARP.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PriorityLevelARP marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PreEmptionCapability.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PreEmptionCapability marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PreEmptionVulnerability.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PreEmptionVulnerability marshal failed")
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

func (x *AllocationAndRetentionPriority) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AllocationAndRetentionPriorityOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AllocationAndRetentionPriorityOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PriorityLevelARP = new(PriorityLevelARP)
	err = x.PriorityLevelARP.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PriorityLevelARP error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PreEmptionCapability = new(PreEmptionCapability)
	err = x.PreEmptionCapability.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PreEmptionCapability error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PreEmptionVulnerability = new(PreEmptionVulnerability)
	err = x.PreEmptionVulnerability.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PreEmptionVulnerability error")
	}

	// optional field (optPresentFlag index: 0)
	if AllocationAndRetentionPriorityOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAllocationAndRetentionPriorityExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
