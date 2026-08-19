package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ExcessPacketDelayThresholdItem struct {
	FiveQi                          *FiveQI
	ExcessPacketDelayThresholdValue *ExcessPacketDelayThresholdValue                                // valueExt,valueLB:0,valueUB:18
	IEExtensions                    *ProtocolExtensionContainerExcessPacketDelayThresholdItemExtIEs // optional
}

func (x *ExcessPacketDelayThresholdItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExcessPacketDelayThresholdItemOptPresentFlag := []bool{}
	// mandatory field
	if x.FiveQi == nil {
		return errors.Errorf("FiveQi is missing")
	}
	// mandatory field
	if x.ExcessPacketDelayThresholdValue == nil {
		return errors.Errorf("ExcessPacketDelayThresholdValue is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ExcessPacketDelayThresholdItemOptPresentFlag = append(ExcessPacketDelayThresholdItemOptPresentFlag, true)
	} else {
		ExcessPacketDelayThresholdItemOptPresentFlag = append(ExcessPacketDelayThresholdItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExcessPacketDelayThresholdItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.FiveQi.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveQi marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ExcessPacketDelayThresholdValue.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ExcessPacketDelayThresholdValue marshal failed")
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

func (x *ExcessPacketDelayThresholdItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExcessPacketDelayThresholdItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ExcessPacketDelayThresholdItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveQi = new(FiveQI)
	err = x.FiveQi.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveQi error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ExcessPacketDelayThresholdValue = new(ExcessPacketDelayThresholdValue)
	err = x.ExcessPacketDelayThresholdValue.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ExcessPacketDelayThresholdValue error")
	}

	// optional field (optPresentFlag index: 0)
	if ExcessPacketDelayThresholdItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExcessPacketDelayThresholdItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
