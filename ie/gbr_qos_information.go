package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type GBRQosInformation struct {
	MaximumFlowBitRateDL    *BitRate
	MaximumFlowBitRateUL    *BitRate
	GuaranteedFlowBitRateDL *BitRate
	GuaranteedFlowBitRateUL *BitRate
	NotificationControl     *NotificationControl                               // valueExt,valueLB:0,valueUB:0,optional
	MaximumPacketLossRateDL *PacketLossRate                                    // optional
	MaximumPacketLossRateUL *PacketLossRate                                    // optional
	IEExtensions            *ProtocolExtensionContainerGBRQosInformationExtIEs // optional
}

func (x *GBRQosInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GBRQosInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.MaximumFlowBitRateDL == nil {
		return errors.Errorf("MaximumFlowBitRateDL is missing")
	}
	// mandatory field
	if x.MaximumFlowBitRateUL == nil {
		return errors.Errorf("MaximumFlowBitRateUL is missing")
	}
	// mandatory field
	if x.GuaranteedFlowBitRateDL == nil {
		return errors.Errorf("GuaranteedFlowBitRateDL is missing")
	}
	// mandatory field
	if x.GuaranteedFlowBitRateUL == nil {
		return errors.Errorf("GuaranteedFlowBitRateUL is missing")
	}
	// optional field
	if x.NotificationControl != nil {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, true)
	} else {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, false)
	}
	// optional field
	if x.MaximumPacketLossRateDL != nil {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, true)
	} else {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, false)
	}
	// optional field
	if x.MaximumPacketLossRateUL != nil {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, true)
	} else {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, true)
	} else {
		GBRQosInformationOptPresentFlag = append(GBRQosInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GBRQosInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MaximumFlowBitRateDL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MaximumFlowBitRateDL marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MaximumFlowBitRateUL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MaximumFlowBitRateUL marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.GuaranteedFlowBitRateDL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GuaranteedFlowBitRateDL marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.GuaranteedFlowBitRateUL.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GuaranteedFlowBitRateUL marshal failed")
	}

	// optional field
	if x.NotificationControl != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NotificationControl.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NotificationControl marshal failed")
		}
	}

	// optional field
	if x.MaximumPacketLossRateDL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MaximumPacketLossRateDL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MaximumPacketLossRateDL marshal failed")
		}
	}

	// optional field
	if x.MaximumPacketLossRateUL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MaximumPacketLossRateUL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MaximumPacketLossRateUL marshal failed")
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

func (x *GBRQosInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GBRQosInformationOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&GBRQosInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MaximumFlowBitRateDL = new(BitRate)
	err = x.MaximumFlowBitRateDL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MaximumFlowBitRateDL error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MaximumFlowBitRateUL = new(BitRate)
	err = x.MaximumFlowBitRateUL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MaximumFlowBitRateUL error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GuaranteedFlowBitRateDL = new(BitRate)
	err = x.GuaranteedFlowBitRateDL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GuaranteedFlowBitRateDL error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GuaranteedFlowBitRateUL = new(BitRate)
	err = x.GuaranteedFlowBitRateUL.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GuaranteedFlowBitRateUL error")
	}

	// optional field (optPresentFlag index: 0)
	if GBRQosInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NotificationControl = new(NotificationControl)
		err = x.NotificationControl.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NotificationControl error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if GBRQosInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MaximumPacketLossRateDL = new(PacketLossRate)
		err = x.MaximumPacketLossRateDL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MaximumPacketLossRateDL error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if GBRQosInformationOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.MaximumPacketLossRateUL = new(PacketLossRate)
		err = x.MaximumPacketLossRateUL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MaximumPacketLossRateUL error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if GBRQosInformationOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGBRQosInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
