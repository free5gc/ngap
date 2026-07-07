package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SpatialRelationInfo struct {
	SpatialRelationforResourceID *SpatialRelationforResourceID
	IEExtensions                 *ProtocolExtensionContainerSpatialRelationInfoExtIEs // optional
}

func (x *SpatialRelationInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.SpatialRelationforResourceID == nil {
		return errors.Errorf("SpatialRelationforResourceID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SpatialRelationInfoOptPresentFlag = append(SpatialRelationInfoOptPresentFlag, true)
	} else {
		SpatialRelationInfoOptPresentFlag = append(SpatialRelationInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SpatialRelationforResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SpatialRelationforResourceID marshal failed")
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

func (x *SpatialRelationInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationInfoOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SpatialRelationforResourceID = new(SpatialRelationforResourceID)
	err = x.SpatialRelationforResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SpatialRelationforResourceID error")
	}

	// optional field (optPresentFlag index: 0)
	if SpatialRelationInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSpatialRelationInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
