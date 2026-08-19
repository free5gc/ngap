package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPElevationAngleListItem struct {
	TrpElevationAngle     *int64 // valueLB:0,valueUB:180
	TrpElevationAngleFine *int64 // valueLB:0,valueUB:9,optional
	/* Sequence of = 35, FULL Name = struct TRP_ElevationAngleList_Item__trp_beam_power_list */
	/* Type Name = TRPBeamPowerItem */
	/* Sequence Of Embed */
	TrpBeamPowerList []TRPBeamPowerItem                                         // valueExt,sizeLB:2,sizeUB:24
	IEExtensions     *ProtocolExtensionContainerTRPElevationAngleListItemExtIEs // optional
}

func (x *TRPElevationAngleListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPElevationAngleListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TrpElevationAngle == nil {
		return errors.Errorf("TrpElevationAngle is missing")
	}
	// optional field
	if x.TrpElevationAngleFine != nil {
		TRPElevationAngleListItemOptPresentFlag = append(TRPElevationAngleListItemOptPresentFlag, true)
	} else {
		TRPElevationAngleListItemOptPresentFlag = append(TRPElevationAngleListItemOptPresentFlag, false)
	}
	// mandatory field
	if x.TrpBeamPowerList == nil {
		return errors.Errorf("TrpBeamPowerList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPElevationAngleListItemOptPresentFlag = append(TRPElevationAngleListItemOptPresentFlag, true)
	} else {
		TRPElevationAngleListItemOptPresentFlag = append(TRPElevationAngleListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPElevationAngleListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 180
	err = pd.WriteInteger(*(x.TrpElevationAngle), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.TrpElevationAngleFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.TrpElevationAngleFine), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Sequence Of
	*sLb, *sUb = 2, 24
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.TrpBeamPowerList)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.TrpBeamPowerList {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
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

func (x *TRPElevationAngleListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPElevationAngleListItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPElevationAngleListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 180
	x.TrpElevationAngle = new(int64)
	*(x.TrpElevationAngle), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if TRPElevationAngleListItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.TrpElevationAngleFine = new(int64)
		*(x.TrpElevationAngleFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 2, 24
	var numElementsTrpBeamPowerList uint64
	numElementsTrpBeamPowerList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.TrpBeamPowerList = []TRPBeamPowerItem{}
	for i := 0; i < int(numElementsTrpBeamPowerList); i++ {
		var val TRPBeamPowerItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.TrpBeamPowerList = append(x.TrpBeamPowerList, val)
		}
	}

	// optional field (optPresentFlag index: 1)
	if TRPElevationAngleListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPElevationAngleListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
