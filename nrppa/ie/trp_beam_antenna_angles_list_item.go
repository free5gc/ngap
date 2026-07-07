package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPBeamAntennaAnglesListItem struct {
	TrpAzimuthAngle     *int64 // valueLB:0,valueUB:359
	TrpAzimuthAngleFine *int64 // valueLB:0,valueUB:9,optional
	/* Sequence of = 35, FULL Name = struct TRP_BeamAntennaAnglesList_Item__trp_elevation_angle_list */
	/* Type Name = TRPElevationAngleListItem */
	/* Sequence Of Embed */
	TrpElevationAngleList []TRPElevationAngleListItem                                   // valueExt,sizeLB:1,sizeUB:1801
	IEExtensions          *ProtocolExtensionContainerTRPBeamAntennaAnglesListItemExtIEs // optional
}

func (x *TRPBeamAntennaAnglesListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamAntennaAnglesListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TrpAzimuthAngle == nil {
		return errors.Errorf("TrpAzimuthAngle is missing")
	}
	// optional field
	if x.TrpAzimuthAngleFine != nil {
		TRPBeamAntennaAnglesListItemOptPresentFlag = append(TRPBeamAntennaAnglesListItemOptPresentFlag, true)
	} else {
		TRPBeamAntennaAnglesListItemOptPresentFlag = append(TRPBeamAntennaAnglesListItemOptPresentFlag, false)
	}
	// mandatory field
	if x.TrpElevationAngleList == nil {
		return errors.Errorf("TrpElevationAngleList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPBeamAntennaAnglesListItemOptPresentFlag = append(TRPBeamAntennaAnglesListItemOptPresentFlag, true)
	} else {
		TRPBeamAntennaAnglesListItemOptPresentFlag = append(TRPBeamAntennaAnglesListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamAntennaAnglesListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 359
	err = pd.WriteInteger(*(x.TrpAzimuthAngle), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.TrpAzimuthAngleFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.TrpAzimuthAngleFine), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 1801
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.TrpElevationAngleList)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.TrpElevationAngleList {
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

func (x *TRPBeamAntennaAnglesListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamAntennaAnglesListItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamAntennaAnglesListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 359
	x.TrpAzimuthAngle = new(int64)
	*(x.TrpAzimuthAngle), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if TRPBeamAntennaAnglesListItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.TrpAzimuthAngleFine = new(int64)
		*(x.TrpAzimuthAngleFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 1801
	var numElementsTrpElevationAngleList uint64
	numElementsTrpElevationAngleList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.TrpElevationAngleList = []TRPElevationAngleListItem{}
	for i := 0; i < int(numElementsTrpElevationAngleList); i++ {
		var val TRPElevationAngleListItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.TrpElevationAngleList = append(x.TrpElevationAngleList, val)
		}
	}

	// optional field (optPresentFlag index: 1)
	if TRPBeamAntennaAnglesListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPBeamAntennaAnglesListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
