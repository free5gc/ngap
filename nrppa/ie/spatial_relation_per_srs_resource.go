package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SpatialRelationPerSRSResource struct {
	SpatialRelationPerSRSResourceList *SpatialRelationPerSRSResourceList
	IEExtensions                      *ProtocolExtensionContainerSpatialRelationPerSRSResourceExtIEs // optional
}

func (x *SpatialRelationPerSRSResource) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationPerSRSResourceOptPresentFlag := []bool{}
	// mandatory field
	if x.SpatialRelationPerSRSResourceList == nil {
		return errors.Errorf("SpatialRelationPerSRSResourceList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SpatialRelationPerSRSResourceOptPresentFlag = append(SpatialRelationPerSRSResourceOptPresentFlag, true)
	} else {
		SpatialRelationPerSRSResourceOptPresentFlag = append(SpatialRelationPerSRSResourceOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationPerSRSResourceOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SpatialRelationPerSRSResourceList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SpatialRelationPerSRSResourceList marshal failed")
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

func (x *SpatialRelationPerSRSResource) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationPerSRSResourceOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationPerSRSResourceOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SpatialRelationPerSRSResourceList = new(SpatialRelationPerSRSResourceList)
	err = x.SpatialRelationPerSRSResourceList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SpatialRelationPerSRSResourceList error")
	}

	// optional field (optPresentFlag index: 0)
	if SpatialRelationPerSRSResourceOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSpatialRelationPerSRSResourceExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
