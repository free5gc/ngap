package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SpatialRelationforResourceIDItem struct {
	ReferenceSignal *ReferenceSignal                                                  // valueLB:0,valueUB:5
	IEExtensions    *ProtocolExtensionContainerSpatialRelationforResourceIDItemExtIEs // optional
}

func (x *SpatialRelationforResourceIDItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationforResourceIDItemOptPresentFlag := []bool{}
	// mandatory field
	if x.ReferenceSignal == nil {
		return errors.Errorf("ReferenceSignal is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SpatialRelationforResourceIDItemOptPresentFlag = append(SpatialRelationforResourceIDItemOptPresentFlag, true)
	} else {
		SpatialRelationforResourceIDItemOptPresentFlag = append(SpatialRelationforResourceIDItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationforResourceIDItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ReferenceSignal.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReferenceSignal marshal failed")
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

func (x *SpatialRelationforResourceIDItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationforResourceIDItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationforResourceIDItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReferenceSignal = new(ReferenceSignal)
	err = x.ReferenceSignal.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReferenceSignal error")
	}

	// optional field (optPresentFlag index: 0)
	if SpatialRelationforResourceIDItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSpatialRelationforResourceIDItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
