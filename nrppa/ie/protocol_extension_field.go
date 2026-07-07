package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type OTDOAInformationTypeItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	OTDOAInformationTypeItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *OTDOAInformationTypeItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOAInformationTypeItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOAInformationTypeItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *OTDOAInformationTypeItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOAInformationTypeItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOAInformationTypeItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SemipersistentSRSExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SemipersistentSRSExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSSpatialRelation               *SpatialRelationInfo           // valueExt,referenceFieldValue:48
	SRSSpatialRelationPerSRSResource *SpatialRelationPerSRSResource // valueExt,referenceFieldValue:63
}

func (x *SemipersistentSRSExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SemipersistentSRSExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SemipersistentSRSExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSSpatialRelation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSSpatialRelation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSSpatialRelation marshal failed")
		}
	} else if x.SRSSpatialRelationPerSRSResource != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSSpatialRelationPerSRSResource.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSSpatialRelationPerSRSResource marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SemipersistentSRSExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SemipersistentSRSExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SemipersistentSRSExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 48 {
		// Read struct defined elsewhere (Pointer)
		x.SRSSpatialRelation = new(SpatialRelationInfo)
		err = x.SRSSpatialRelation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSSpatialRelation error")
		}
	} else if x.Id.Value == 63 {
		// Read struct defined elsewhere (Pointer)
		x.SRSSpatialRelationPerSRSResource = new(SpatialRelationPerSRSResource)
		err = x.SRSSpatialRelationPerSRSResource.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSSpatialRelationPerSRSResource error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type AperiodicSRSExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	AperiodicSRSExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AperiodicSRSExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AperiodicSRSExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AperiodicSRSExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *AperiodicSRSExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AperiodicSRSExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AperiodicSRSExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ActiveULBWPExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ActiveULBWPExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ActiveULBWPExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ActiveULBWPExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ActiveULBWPExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ActiveULBWPExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ActiveULBWPExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ActiveULBWPExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type AdditionalPathListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	AdditionalPathListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	MultipleULAoA *MultipleULAoA // valueExt,referenceFieldValue:74
	PathPower     *ULSRSRSRPP    // valueExt,referenceFieldValue:96
}

func (x *AdditionalPathListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AdditionalPathListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AdditionalPathListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.MultipleULAoA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MultipleULAoA.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MultipleULAoA marshal failed")
		}
	} else if x.PathPower != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PathPower.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PathPower marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *AdditionalPathListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AdditionalPathListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AdditionalPathListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 74 {
		// Read struct defined elsewhere (Pointer)
		x.MultipleULAoA = new(MultipleULAoA)
		err = x.MultipleULAoA.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MultipleULAoA error")
		}
	} else if x.Id.Value == 96 {
		// Read struct defined elsewhere (Pointer)
		x.PathPower = new(ULSRSRSRPP)
		err = x.PathPower.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PathPower error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type ExtendedAdditionalPathListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ExtendedAdditionalPathListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ExtendedAdditionalPathListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExtendedAdditionalPathListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ExtendedAdditionalPathListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ExtendedAdditionalPathListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExtendedAdditionalPathListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ExtendedAdditionalPathListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type AoAAssistanceInfoExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	AoAAssistanceInfoExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AoAAssistanceInfoExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AoAAssistanceInfoExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AoAAssistanceInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *AoAAssistanceInfoExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AoAAssistanceInfoExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AoAAssistanceInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ExpectedULAoAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ExpectedULAoAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ExpectedULAoAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedULAoAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedULAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ExpectedULAoAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedULAoAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ExpectedULAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ExpectedZoAOnlyExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ExpectedZoAOnlyExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ExpectedZoAOnlyExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedZoAOnlyExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedZoAOnlyExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ExpectedZoAOnlyExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedZoAOnlyExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ExpectedZoAOnlyExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ExpectedAzimuthAoAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ExpectedAzimuthAoAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ExpectedAzimuthAoAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedAzimuthAoAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedAzimuthAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ExpectedAzimuthAoAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedAzimuthAoAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ExpectedAzimuthAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ExpectedZenithAoAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ExpectedZenithAoAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ExpectedZenithAoAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExpectedZenithAoAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ExpectedZenithAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ExpectedZenithAoAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExpectedZenithAoAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ExpectedZenithAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ARPLocationInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ARPLocationInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ARPLocationInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ARPLocationInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ARPLocationInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ARPLocationInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ARPLocationInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ARPLocationInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type AssistanceInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	AssistanceInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AssistanceInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *AssistanceInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type AssistanceInformationFailureListExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	AssistanceInformationFailureListExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AssistanceInformationFailureListExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationFailureListExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationFailureListExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *AssistanceInformationFailureListExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationFailureListExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationFailureListExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type AssistanceInformationMetaDataExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	AssistanceInformationMetaDataExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AssistanceInformationMetaDataExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationMetaDataExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationMetaDataExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *AssistanceInformationMetaDataExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationMetaDataExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationMetaDataExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type CarrierFreqExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	CarrierFreqExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *CarrierFreqExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CarrierFreqExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(CarrierFreqExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *CarrierFreqExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CarrierFreqExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&CarrierFreqExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type CGIEUTRAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	CGIEUTRAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *CGIEUTRAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CGIEUTRAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(CGIEUTRAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *CGIEUTRAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CGIEUTRAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&CGIEUTRAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type CGINRExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	CGINRExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *CGINRExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CGINRExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(CGINRExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *CGINRExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CGINRExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&CGINRExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type CriticalityDiagnosticsExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	CriticalityDiagnosticsExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *CriticalityDiagnosticsExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CriticalityDiagnosticsExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(CriticalityDiagnosticsExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *CriticalityDiagnosticsExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CriticalityDiagnosticsExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&CriticalityDiagnosticsExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type CriticalityDiagnosticsIEListExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	CriticalityDiagnosticsIEListExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *CriticalityDiagnosticsIEListExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CriticalityDiagnosticsIEListExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(CriticalityDiagnosticsIEListExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *CriticalityDiagnosticsIEListExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CriticalityDiagnosticsIEListExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&CriticalityDiagnosticsIEListExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type DLPRSExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	DLPRSExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *DLPRSExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type DLPRSResourceCoordinatesExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	DLPRSResourceCoordinatesExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSResourceCoordinatesExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceCoordinatesExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceCoordinatesExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *DLPRSResourceCoordinatesExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceCoordinatesExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceCoordinatesExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type DLPRSResourceSetARPExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	DLPRSResourceSetARPExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSResourceSetARPExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceSetARPExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceSetARPExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *DLPRSResourceSetARPExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceSetARPExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceSetARPExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type DLPRSResourceARPExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	DLPRSResourceARPExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSResourceARPExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceARPExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceARPExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *DLPRSResourceARPExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceARPExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceARPExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ECIDMeasurementResultExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ECIDMeasurementResultExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	GeographicalCoordinates *GeographicalCoordinates // valueExt,referenceFieldValue:37
}

func (x *ECIDMeasurementResultExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementResultExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementResultExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.GeographicalCoordinates != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.GeographicalCoordinates.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "GeographicalCoordinates marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ECIDMeasurementResultExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementResultExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementResultExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 37 {
		// Read struct defined elsewhere (Pointer)
		x.GeographicalCoordinates = new(GeographicalCoordinates)
		err = x.GeographicalCoordinates.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode GeographicalCoordinates error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type GeographicalCoordinatesExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	GeographicalCoordinatesExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	ARPLocationInfo *ARPLocationInformation // refFieldVal:78
}

func (x *GeographicalCoordinatesExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GeographicalCoordinatesExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(GeographicalCoordinatesExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.ARPLocationInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ARPLocationInfo.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ARPLocationInfo marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *GeographicalCoordinatesExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GeographicalCoordinatesExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&GeographicalCoordinatesExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 78 {
		// Read struct defined elsewhere (Pointer)
		x.ARPLocationInfo = new(ARPLocationInformation)
		err = x.ARPLocationInfo.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ARPLocationInfo error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type GNBRxTxTimeDiffExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	GNBRxTxTimeDiffExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	ExtendedAdditionalPathList *ExtendedAdditionalPathList // refFieldVal:77
	TRPTEGInformation          *TRPTEGInformation          // refFieldVal:85,valueLB:0,valueUB:2
}

func (x *GNBRxTxTimeDiffExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GNBRxTxTimeDiffExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(GNBRxTxTimeDiffExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.ExtendedAdditionalPathList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExtendedAdditionalPathList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ExtendedAdditionalPathList marshal failed")
		}
	} else if x.TRPTEGInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPTEGInformation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPTEGInformation marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *GNBRxTxTimeDiffExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GNBRxTxTimeDiffExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&GNBRxTxTimeDiffExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 77 {
		// Read struct defined elsewhere (Pointer)
		x.ExtendedAdditionalPathList = new(ExtendedAdditionalPathList)
		err = x.ExtendedAdditionalPathList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ExtendedAdditionalPathList error")
		}
	} else if x.Id.Value == 85 {
		// Read struct defined elsewhere (Pointer)
		x.TRPTEGInformation = new(TRPTEGInformation)
		err = x.TRPTEGInformation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPTEGInformation error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type LCSToGCSTranslationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	LCSToGCSTranslationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *LCSToGCSTranslationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LCSToGCSTranslationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(LCSToGCSTranslationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *LCSToGCSTranslationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LCSToGCSTranslationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&LCSToGCSTranslationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type LCSToGCSTranslationItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	LCSToGCSTranslationItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *LCSToGCSTranslationItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LCSToGCSTranslationItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(LCSToGCSTranslationItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *LCSToGCSTranslationItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LCSToGCSTranslationItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&LCSToGCSTranslationItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type LocationUncertaintyExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	LocationUncertaintyExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *LocationUncertaintyExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LocationUncertaintyExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(LocationUncertaintyExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *LocationUncertaintyExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LocationUncertaintyExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&LocationUncertaintyExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type MeasurementBeamInfoExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	MeasurementBeamInfoExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *MeasurementBeamInfoExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementBeamInfoExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementBeamInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *MeasurementBeamInfoExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementBeamInfoExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementBeamInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type MeasurementQuantitiesValueExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	MeasurementQuantitiesValueExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *MeasurementQuantitiesValueExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementQuantitiesValueExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementQuantitiesValueExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *MeasurementQuantitiesValueExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementQuantitiesValueExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementQuantitiesValueExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type MultipleULAoAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	MultipleULAoAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *MultipleULAoAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MultipleULAoAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MultipleULAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *MultipleULAoAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MultipleULAoAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MultipleULAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type NGRANAccessPointPositionExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	NGRANAccessPointPositionExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *NGRANAccessPointPositionExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANAccessPointPositionExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(NGRANAccessPointPositionExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *NGRANAccessPointPositionExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANAccessPointPositionExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&NGRANAccessPointPositionExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type NGRANHighAccuracyAccessPointPositionExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	NGRANHighAccuracyAccessPointPositionExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *NGRANHighAccuracyAccessPointPositionExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANHighAccuracyAccessPointPositionExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(NGRANHighAccuracyAccessPointPositionExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *NGRANHighAccuracyAccessPointPositionExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANHighAccuracyAccessPointPositionExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&NGRANHighAccuracyAccessPointPositionExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type NGRANCGIExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	NGRANCGIExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *NGRANCGIExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANCGIExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(NGRANCGIExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *NGRANCGIExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANCGIExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&NGRANCGIExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type NRPRSBeamInformationIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	NRPRSBeamInformationIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *NRPRSBeamInformationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRPRSBeamInformationIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(NRPRSBeamInformationIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *NRPRSBeamInformationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRPRSBeamInformationIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&NRPRSBeamInformationIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type NRPRSBeamInformationItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	NRPRSBeamInformationItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *NRPRSBeamInformationItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRPRSBeamInformationItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(NRPRSBeamInformationItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *NRPRSBeamInformationItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRPRSBeamInformationItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&NRPRSBeamInformationItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type OnDemandPRSInfoExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	OnDemandPRSInfoExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *OnDemandPRSInfoExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OnDemandPRSInfoExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OnDemandPRSInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *OnDemandPRSInfoExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OnDemandPRSInfoExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OnDemandPRSInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type OTDOACellsExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	OTDOACellsExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *OTDOACellsExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOACellsExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOACellsExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *OTDOACellsExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOACellsExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOACellsExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type OtherRATMeasurementQuantitiesValueExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	OtherRATMeasurementQuantitiesValueExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *OtherRATMeasurementQuantitiesValueExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OtherRATMeasurementQuantitiesValueExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OtherRATMeasurementQuantitiesValueExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *OtherRATMeasurementQuantitiesValueExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OtherRATMeasurementQuantitiesValueExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OtherRATMeasurementQuantitiesValueExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PathlossReferenceInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PathlossReferenceInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PathlossReferenceInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PathlossReferenceInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PathlossReferenceInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PathlossReferenceInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PathlossReferenceInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PathlossReferenceInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosSIBsExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosSIBsExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosSIBsExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSIBsExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosSIBsExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosSIBsExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSIBsExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosSIBsExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosSIBSegmentsExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosSIBSegmentsExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosSIBSegmentsExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSIBSegmentsExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosSIBSegmentsExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosSIBSegmentsExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSIBSegmentsExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosSIBSegmentsExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosSRSResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosSRSResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosSRSResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSRSResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosSRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosSRSResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSRSResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosSRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosSRSResourceSetItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosSRSResourceSetItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosSRSResourceSetItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSRSResourceSetItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosSRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosSRSResourceSetItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSRSResourceSetItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosSRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosResourceSetTypePeriodicExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosResourceSetTypePeriodicExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosResourceSetTypePeriodicExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosResourceSetTypePeriodicExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosResourceSetTypePeriodicExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosResourceSetTypePeriodicExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosResourceSetTypePeriodicExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosResourceSetTypePeriodicExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosResourceSetTypeSemiPersistentExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosResourceSetTypeSemiPersistentExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosResourceSetTypeSemiPersistentExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosResourceSetTypeSemiPersistentExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosResourceSetTypeSemiPersistentExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosResourceSetTypeSemiPersistentExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosResourceSetTypeSemiPersistentExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosResourceSetTypeSemiPersistentExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PosResourceSetTypeAperiodicExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PosResourceSetTypeAperiodicExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosResourceSetTypeAperiodicExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosResourceSetTypeAperiodicExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosResourceSetTypeAperiodicExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PosResourceSetTypeAperiodicExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosResourceSetTypeAperiodicExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosResourceSetTypeAperiodicExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSAngleItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSAngleItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	PRSResourceID *PRSResourceID // refFieldVal:65
}

func (x *PRSAngleItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSAngleItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSAngleItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.PRSResourceID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSResourceID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PRSResourceID marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSAngleItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSAngleItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSAngleItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 65 {
		// Read struct defined elsewhere (Pointer)
		x.PRSResourceID = new(PRSResourceID)
		err = x.PRSResourceID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PRSResourceID error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type PRSInformationPosExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSInformationPosExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSInformationPosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSInformationPosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSInformationPosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSInformationPosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSInformationPosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSInformationPosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSConfigurationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSConfigurationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSConfigurationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSConfigurationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSConfigurationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSConfigurationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSConfigurationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSConfigurationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSFrequencyHoppingConfigurationEUTRAItemIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSFrequencyHoppingConfigurationEUTRAItemIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSFrequencyHoppingConfigurationEUTRAItemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSFrequencyHoppingConfigurationEUTRAItemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSFrequencyHoppingConfigurationEUTRAItemIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSFrequencyHoppingConfigurationEUTRAItemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSFrequencyHoppingConfigurationEUTRAItemIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSFrequencyHoppingConfigurationEUTRAItemIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSMeasurementsInfoListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSMeasurementsInfoListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSMeasurementsInfoListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMeasurementsInfoListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSMeasurementsInfoListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSMeasurementsInfoListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMeasurementsInfoListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSMeasurementsInfoListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSMutingExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSMutingExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSMutingExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSMutingExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSMutingOption1ExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSMutingOption1ExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSMutingOption1ExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingOption1ExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingOption1ExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSMutingOption1ExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingOption1ExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingOption1ExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSMutingOption2ExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSMutingOption2ExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSMutingOption2ExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingOption2ExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingOption2ExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSMutingOption2ExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingOption2ExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingOption2ExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSResourceQCLSourceSSBExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSResourceQCLSourceSSBExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSResourceQCLSourceSSBExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceQCLSourceSSBExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceQCLSourceSSBExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSResourceQCLSourceSSBExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceQCLSourceSSBExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceQCLSourceSSBExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSResourceQCLSourcePRSExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSResourceQCLSourcePRSExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSResourceQCLSourcePRSExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceQCLSourcePRSExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceQCLSourcePRSExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSResourceQCLSourcePRSExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceQCLSourcePRSExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceQCLSourcePRSExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSResourceSetItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSResourceSetItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSResourceSetItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceSetItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSResourceSetItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceSetItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSTransmissionOffPerResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSTransmissionOffPerResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTransmissionOffPerResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffPerResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffPerResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSTransmissionOffPerResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffPerResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffPerResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSTransmissionOffIndicationPerResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSTransmissionOffIndicationPerResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTransmissionOffIndicationPerResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffIndicationPerResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffIndicationPerResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSTransmissionOffIndicationPerResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffIndicationPerResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffIndicationPerResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSTransmissionOffInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSTransmissionOffInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTransmissionOffInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSTransmissionOffInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSTransmissionOffPerResourceSetItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSTransmissionOffPerResourceSetItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTransmissionOffPerResourceSetItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffPerResourceSetItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffPerResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSTransmissionOffPerResourceSetItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffPerResourceSetItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffPerResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSTRPItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSTRPItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTRPItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTRPItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSTRPItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTRPItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type PRSTransmissionTRPItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	PRSTransmissionTRPItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTransmissionTRPItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionTRPItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionTRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *PRSTransmissionTRPItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionTRPItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionTRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RelativeGeodeticLocationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RelativeGeodeticLocationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RelativeGeodeticLocationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RelativeGeodeticLocationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RelativeGeodeticLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RelativeGeodeticLocationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RelativeGeodeticLocationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RelativeGeodeticLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RelativeCartesianLocationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RelativeCartesianLocationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RelativeCartesianLocationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RelativeCartesianLocationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RelativeCartesianLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RelativeCartesianLocationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RelativeCartesianLocationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RelativeCartesianLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RequestedDLPRSTransmissionCharacteristicsExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RequestedDLPRSTransmissionCharacteristicsExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RequestedDLPRSTransmissionCharacteristicsExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedDLPRSTransmissionCharacteristicsExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RequestedDLPRSTransmissionCharacteristicsExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RequestedDLPRSTransmissionCharacteristicsExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedDLPRSTransmissionCharacteristicsExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RequestedDLPRSTransmissionCharacteristicsExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RequestedDLPRSResourceSetItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RequestedDLPRSResourceSetItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RequestedDLPRSResourceSetItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedDLPRSResourceSetItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RequestedDLPRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RequestedDLPRSResourceSetItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedDLPRSResourceSetItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RequestedDLPRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RequestedDLPRSResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RequestedDLPRSResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RequestedDLPRSResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedDLPRSResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RequestedDLPRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RequestedDLPRSResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedDLPRSResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RequestedDLPRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RequestedSRSTransmissionCharacteristicsExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RequestedSRSTransmissionCharacteristicsExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SrsFrequency *SrsFrequency // refFieldVal:61
}

func (x *RequestedSRSTransmissionCharacteristicsExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedSRSTransmissionCharacteristicsExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RequestedSRSTransmissionCharacteristicsExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SrsFrequency != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SrsFrequency.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SrsFrequency marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RequestedSRSTransmissionCharacteristicsExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedSRSTransmissionCharacteristicsExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RequestedSRSTransmissionCharacteristicsExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 61 {
		// Read struct defined elsewhere (Pointer)
		x.SrsFrequency = new(SrsFrequency)
		err = x.SrsFrequency.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SrsFrequency error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type SRSResourceSetItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSResourceSetItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSSpatialRelationPerSRSResource *SpatialRelationPerSRSResource // valueExt,referenceFieldValue:63
}

func (x *SRSResourceSetItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceSetItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSSpatialRelationPerSRSResource != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSSpatialRelationPerSRSResource.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSSpatialRelationPerSRSResource marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSResourceSetItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceSetItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceSetItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 63 {
		// Read struct defined elsewhere (Pointer)
		x.SRSSpatialRelationPerSRSResource = new(SpatialRelationPerSRSResource)
		err = x.SRSSpatialRelationPerSRSResource.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSSpatialRelationPerSRSResource error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type ResourceSetTypePeriodicExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceSetTypePeriodicExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceSetTypePeriodicExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypePeriodicExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypePeriodicExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceSetTypePeriodicExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypePeriodicExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypePeriodicExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceSetTypeSemiPersistentExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceSetTypeSemiPersistentExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceSetTypeSemiPersistentExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypeSemiPersistentExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypeSemiPersistentExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceSetTypeSemiPersistentExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypeSemiPersistentExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypeSemiPersistentExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceSetTypeAperiodicExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceSetTypeAperiodicExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceSetTypeAperiodicExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypeAperiodicExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypeAperiodicExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceSetTypeAperiodicExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypeAperiodicExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypeAperiodicExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceTypePeriodicExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceTypePeriodicExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypePeriodicExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypePeriodicExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypePeriodicExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceTypePeriodicExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypePeriodicExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypePeriodicExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceTypeSemiPersistentExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceTypeSemiPersistentExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypeSemiPersistentExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeSemiPersistentExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeSemiPersistentExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceTypeSemiPersistentExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeSemiPersistentExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeSemiPersistentExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceTypeAperiodicExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceTypeAperiodicExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypeAperiodicExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeAperiodicExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeAperiodicExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceTypeAperiodicExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeAperiodicExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeAperiodicExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceTypePeriodicPosExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceTypePeriodicPosExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypePeriodicPosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypePeriodicPosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypePeriodicPosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceTypePeriodicPosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypePeriodicPosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypePeriodicPosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceTypeSemiPersistentPosExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceTypeSemiPersistentPosExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypeSemiPersistentPosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeSemiPersistentPosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeSemiPersistentPosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceTypeSemiPersistentPosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeSemiPersistentPosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeSemiPersistentPosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResourceTypeAperiodicPosExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResourceTypeAperiodicPosExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypeAperiodicPosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeAperiodicPosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeAperiodicPosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResourceTypeAperiodicPosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeAperiodicPosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeAperiodicPosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResponseTimeExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResponseTimeExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResponseTimeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResponseTimeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResponseTimeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResponseTimeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResponseTimeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResponseTimeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultCSIRSRPItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultCSIRSRPItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultCSIRSRPItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRPItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultCSIRSRPItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRPItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultCSIRSRPPerCSIRSItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultCSIRSRPPerCSIRSItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultCSIRSRPPerCSIRSItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRPPerCSIRSItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRPPerCSIRSItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultCSIRSRPPerCSIRSItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRPPerCSIRSItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRPPerCSIRSItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultCSIRSRQItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultCSIRSRQItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultCSIRSRQItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRQItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRQItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultCSIRSRQItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRQItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRQItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultCSIRSRQPerCSIRSItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultCSIRSRQPerCSIRSItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultCSIRSRQPerCSIRSItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRQPerCSIRSItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRQPerCSIRSItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultCSIRSRQPerCSIRSItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRQPerCSIRSItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRQPerCSIRSItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultEUTRAItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultEUTRAItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultEUTRAItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultEUTRAItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultEUTRAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultEUTRAItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultEUTRAItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultEUTRAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultRSRPEUTRAItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultRSRPEUTRAItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultRSRPEUTRAItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultRSRPEUTRAItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultRSRPEUTRAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultRSRPEUTRAItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultRSRPEUTRAItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultRSRPEUTRAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultRSRQEUTRAItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultRSRQEUTRAItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultRSRQEUTRAItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultRSRQEUTRAItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultRSRQEUTRAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultRSRQEUTRAItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultRSRQEUTRAItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultRSRQEUTRAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultSSRSRPItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultSSRSRPItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultSSRSRPItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRPItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultSSRSRPItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRPItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultSSRSRPPerSSBItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultSSRSRPPerSSBItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultSSRSRPPerSSBItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRPPerSSBItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRPPerSSBItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultSSRSRPPerSSBItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRPPerSSBItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRPPerSSBItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultSSRSRQItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultSSRSRQItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultSSRSRQItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRQItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRQItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultSSRSRQItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRQItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRQItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultSSRSRQPerSSBItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultSSRSRQPerSSBItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultSSRSRQPerSSBItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRQPerSSBItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRQPerSSBItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultSSRSRQPerSSBItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRQPerSSBItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRQPerSSBItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultGERANItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultGERANItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultGERANItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultGERANItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultGERANItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultGERANItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultGERANItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultGERANItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultNRItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultNRItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultNRItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultNRItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultNRItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultNRItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultNRItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultNRItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ResultUTRANItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ResultUTRANItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResultUTRANItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultUTRANItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResultUTRANItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ResultUTRANItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultUTRANItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResultUTRANItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SCSSpecificCarrierExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SCSSpecificCarrierExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SCSSpecificCarrierExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SCSSpecificCarrierExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SCSSpecificCarrierExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SCSSpecificCarrierExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SCSSpecificCarrierExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SCSSpecificCarrierExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SearchWindowInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SearchWindowInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SearchWindowInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SearchWindowInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SearchWindowInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SearchWindowInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SearchWindowInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SearchWindowInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SpatialDirectionInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SpatialDirectionInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SpatialDirectionInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialDirectionInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SpatialDirectionInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SpatialDirectionInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialDirectionInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SpatialDirectionInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SpatialRelationInfoExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SpatialRelationInfoExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SpatialRelationInfoExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationInfoExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SpatialRelationInfoExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationInfoExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SpatialRelationforResourceIDItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SpatialRelationforResourceIDItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SpatialRelationforResourceIDItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationforResourceIDItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationforResourceIDItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SpatialRelationforResourceIDItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationforResourceIDItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationforResourceIDItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SpatialRelationPerSRSResourceExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SpatialRelationPerSRSResourceExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SpatialRelationPerSRSResourceExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationPerSRSResourceExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationPerSRSResourceExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SpatialRelationPerSRSResourceExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationPerSRSResourceExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationPerSRSResourceExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SpatialRelationPerSRSResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SpatialRelationPerSRSResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SpatialRelationPerSRSResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialRelationPerSRSResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SpatialRelationPerSRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SpatialRelationPerSRSResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialRelationPerSRSResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SpatialRelationPerSRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SRSConfigExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSConfigExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSConfigExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSConfigExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSConfigExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSConfigExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSConfigExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSConfigExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SRSCarrierListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSCarrierListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSCarrierListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSCarrierListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSCarrierListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSCarrierListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSCarrierListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSCarrierListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SRSConfigurationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSConfigurationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSConfigurationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSConfigurationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSConfigurationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSConfigurationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSConfigurationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSConfigurationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SRSResourceExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSResourceExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	NrofSymbolsExtended      *NrofSymbolsExtended      // valueExt,referenceFieldValue:107
	RepetitionFactorExtended *RepetitionFactorExtended // valueExt,referenceFieldValue:108
	StartRBHopping           *StartRBHopping           // refFieldVal:109
	StartRBIndex             *StartRBIndex             // refFieldVal:110,valueLB:0,valueUB:2
}

func (x *SRSResourceExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.NrofSymbolsExtended != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NrofSymbolsExtended.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "NrofSymbolsExtended marshal failed")
		}
	} else if x.RepetitionFactorExtended != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RepetitionFactorExtended.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RepetitionFactorExtended marshal failed")
		}
	} else if x.StartRBHopping != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.StartRBHopping.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "StartRBHopping marshal failed")
		}
	} else if x.StartRBIndex != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.StartRBIndex.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "StartRBIndex marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSResourceExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 107 {
		// Read struct defined elsewhere (Pointer)
		x.NrofSymbolsExtended = new(NrofSymbolsExtended)
		err = x.NrofSymbolsExtended.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode NrofSymbolsExtended error")
		}
	} else if x.Id.Value == 108 {
		// Read struct defined elsewhere (Pointer)
		x.RepetitionFactorExtended = new(RepetitionFactorExtended)
		err = x.RepetitionFactorExtended.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RepetitionFactorExtended error")
		}
	} else if x.Id.Value == 109 {
		// Read struct defined elsewhere (Pointer)
		x.StartRBHopping = new(StartRBHopping)
		err = x.StartRBHopping.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode StartRBHopping error")
		}
	} else if x.Id.Value == 110 {
		// Read struct defined elsewhere (Pointer)
		x.StartRBIndex = new(StartRBIndex)
		err = x.StartRBIndex.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode StartRBIndex error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type SRSResourceSetExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSResourceSetExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSResourceSetExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceSetExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceSetExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSResourceSetExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceSetExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceSetExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SRSResourceTriggerExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSResourceTriggerExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSResourceTriggerExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceTriggerExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceTriggerExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSResourceTriggerExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceTriggerExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceTriggerExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SRSResourcetypeExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSResourcetypeExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSPortIndex *SRSPortIndex // valueExt,referenceFieldValue:100
}

func (x *SRSResourcetypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourcetypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourcetypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSPortIndex != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSPortIndex.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSPortIndex marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSResourcetypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourcetypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSResourcetypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 100 {
		// Read struct defined elsewhere (Pointer)
		x.SRSPortIndex = new(SRSPortIndex)
		err = x.SRSPortIndex.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSPortIndex error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type SSBInfoExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SSBInfoExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SSBInfoExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBInfoExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SSBInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SSBInfoExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBInfoExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SSBInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SSBInfoItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SSBInfoItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SSBInfoItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBInfoItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SSBInfoItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SSBInfoItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBInfoItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SSBInfoItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SSBExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SSBExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SSBExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SSBExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SSBExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SSBExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type StartTimeAndDurationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	StartTimeAndDurationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *StartTimeAndDurationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	StartTimeAndDurationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(StartTimeAndDurationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *StartTimeAndDurationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	StartTimeAndDurationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&StartTimeAndDurationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type SystemInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SystemInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SystemInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SystemInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SystemInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SystemInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SystemInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SystemInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TDDConfigEUTRAItemItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TDDConfigEUTRAItemItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TDDConfigEUTRAItemItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TDDConfigEUTRAItemItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TDDConfigEUTRAItemItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TDDConfigEUTRAItemItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TDDConfigEUTRAItemItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TDDConfigEUTRAItemItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RxTxTEGExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RxTxTEGExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RxTxTEGExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RxTxTEGExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RxTxTEGExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RxTxTEGExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RxTxTEGExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RxTxTEGExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type RxTEGExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	RxTEGExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RxTEGExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RxTEGExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RxTEGExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *RxTEGExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RxTEGExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RxTEGExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TFConfigurationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TFConfigurationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TFConfigurationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TFConfigurationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TFConfigurationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TFConfigurationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TFConfigurationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TFConfigurationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TimeStampExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TimeStampExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TimeStampExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TimeStampExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TimeStampExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TimeStampExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TimeStampExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TimeStampExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TransmissionCombn8ExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TransmissionCombn8ExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TransmissionCombn8ExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TransmissionCombn8ExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TransmissionCombn8ExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TransmissionCombn8ExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TransmissionCombn8ExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TransmissionCombn8ExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPBeamAntennaInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPBeamAntennaInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPBeamAntennaInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamAntennaInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamAntennaInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPBeamAntennaInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamAntennaInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamAntennaInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPBeamAntennaExplicitInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPBeamAntennaExplicitInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPBeamAntennaExplicitInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamAntennaExplicitInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamAntennaExplicitInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPBeamAntennaExplicitInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamAntennaExplicitInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamAntennaExplicitInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPBeamAntennaAnglesListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPBeamAntennaAnglesListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPBeamAntennaAnglesListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamAntennaAnglesListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamAntennaAnglesListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPBeamAntennaAnglesListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamAntennaAnglesListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamAntennaAnglesListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPElevationAngleListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPElevationAngleListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPElevationAngleListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPElevationAngleListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPElevationAngleListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPElevationAngleListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPElevationAngleListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPElevationAngleListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPBeamPowerItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPBeamPowerItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPBeamPowerItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamPowerItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamPowerItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPBeamPowerItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamPowerItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamPowerItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPMeasurementQuantitiesListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPMeasurementQuantitiesListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPMeasurementQuantitiesListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementQuantitiesListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementQuantitiesListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPMeasurementQuantitiesListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementQuantitiesListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementQuantitiesListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TrpMeasurementResultItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TrpMeasurementResultItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSResourcetype    *SRSResourcetype    // valueExt,referenceFieldValue:76
	ARPID              *ARPID              // refFieldVal:79
	LoSNLoSInformation *LoSNLoSInformation // refFieldVal:80,valueLB:0,valueUB:2
}

func (x *TrpMeasurementResultItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementResultItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementResultItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSResourcetype != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSResourcetype.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSResourcetype marshal failed")
		}
	} else if x.ARPID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ARPID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ARPID marshal failed")
		}
	} else if x.LoSNLoSInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LoSNLoSInformation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LoSNLoSInformation marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TrpMeasurementResultItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementResultItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementResultItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 76 {
		// Read struct defined elsewhere (Pointer)
		x.SRSResourcetype = new(SRSResourcetype)
		err = x.SRSResourcetype.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSResourcetype error")
		}
	} else if x.Id.Value == 79 {
		// Read struct defined elsewhere (Pointer)
		x.ARPID = new(ARPID)
		err = x.ARPID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ARPID error")
		}
	} else if x.Id.Value == 80 {
		// Read struct defined elsewhere (Pointer)
		x.LoSNLoSInformation = new(LoSNLoSInformation)
		err = x.LoSNLoSInformation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LoSNLoSInformation error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type TrpMeasurementTimingQualityExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TrpMeasurementTimingQualityExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TrpMeasurementTimingQualityExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementTimingQualityExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementTimingQualityExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TrpMeasurementTimingQualityExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementTimingQualityExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementTimingQualityExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TrpMeasurementAngleQualityExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TrpMeasurementAngleQualityExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TrpMeasurementAngleQualityExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementAngleQualityExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementAngleQualityExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TrpMeasurementAngleQualityExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementAngleQualityExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementAngleQualityExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPMeasurementRequestItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPMeasurementRequestItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	CellID             *CGINR              // valueExt,referenceFieldValue:60
	AoASearchWindow    *AoAAssistanceInfo  // valueExt,referenceFieldValue:69
	NumberOfTRPRxTEG   *NumberOfTRPRxTEG   // valueExt,referenceFieldValue:82
	NumberOfTRPRxTxTEG *NumberOfTRPRxTxTEG // valueExt,referenceFieldValue:83
}

func (x *TRPMeasurementRequestItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementRequestItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementRequestItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.CellID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CellID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CellID marshal failed")
		}
	} else if x.AoASearchWindow != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AoASearchWindow.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "AoASearchWindow marshal failed")
		}
	} else if x.NumberOfTRPRxTEG != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NumberOfTRPRxTEG.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "NumberOfTRPRxTEG marshal failed")
		}
	} else if x.NumberOfTRPRxTxTEG != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NumberOfTRPRxTxTEG.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "NumberOfTRPRxTxTEG marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPMeasurementRequestItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementRequestItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementRequestItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 60 {
		// Read struct defined elsewhere (Pointer)
		x.CellID = new(CGINR)
		err = x.CellID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CellID error")
		}
	} else if x.Id.Value == 69 {
		// Read struct defined elsewhere (Pointer)
		x.AoASearchWindow = new(AoAAssistanceInfo)
		err = x.AoASearchWindow.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode AoASearchWindow error")
		}
	} else if x.Id.Value == 82 {
		// Read struct defined elsewhere (Pointer)
		x.NumberOfTRPRxTEG = new(NumberOfTRPRxTEG)
		err = x.NumberOfTRPRxTEG.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode NumberOfTRPRxTEG error")
		}
	} else if x.Id.Value == 83 {
		// Read struct defined elsewhere (Pointer)
		x.NumberOfTRPRxTxTEG = new(NumberOfTRPRxTxTEG)
		err = x.NumberOfTRPRxTxTEG.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode NumberOfTRPRxTxTEG error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type TRPMeasurementResponseItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPMeasurementResponseItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	CellID *CGINR // valueExt,referenceFieldValue:60
}

func (x *TRPMeasurementResponseItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementResponseItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementResponseItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.CellID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CellID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CellID marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPMeasurementResponseItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementResponseItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementResponseItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 60 {
		// Read struct defined elsewhere (Pointer)
		x.CellID = new(CGINR)
		err = x.CellID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CellID error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type TRPMeasurementUpdateItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPMeasurementUpdateItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	NumberOfTRPRxTEG   *NumberOfTRPRxTEG   // valueExt,referenceFieldValue:82
	NumberOfTRPRxTxTEG *NumberOfTRPRxTxTEG // valueExt,referenceFieldValue:83
}

func (x *TRPMeasurementUpdateItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementUpdateItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementUpdateItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.NumberOfTRPRxTEG != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NumberOfTRPRxTEG.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "NumberOfTRPRxTEG marshal failed")
		}
	} else if x.NumberOfTRPRxTxTEG != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NumberOfTRPRxTxTEG.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "NumberOfTRPRxTxTEG marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPMeasurementUpdateItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementUpdateItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementUpdateItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 82 {
		// Read struct defined elsewhere (Pointer)
		x.NumberOfTRPRxTEG = new(NumberOfTRPRxTEG)
		err = x.NumberOfTRPRxTEG.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode NumberOfTRPRxTEG error")
		}
	} else if x.Id.Value == 83 {
		// Read struct defined elsewhere (Pointer)
		x.NumberOfTRPRxTxTEG = new(NumberOfTRPRxTxTEG)
		err = x.NumberOfTRPRxTxTEG.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode NumberOfTRPRxTxTEG error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type TRPInformationTRPRespExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPInformationTRPRespExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPInformationTRPRespExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationTRPRespExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationTRPRespExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPInformationTRPRespExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationTRPRespExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationTRPRespExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPPositionDirectExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPPositionDirectExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPPositionDirectExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPositionDirectExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPPositionDirectExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPPositionDirectExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPositionDirectExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPPositionDirectExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPPositionReferencedExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPPositionReferencedExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPPositionReferencedExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPositionReferencedExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPPositionReferencedExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPPositionReferencedExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPositionReferencedExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPPositionReferencedExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPPRSInformationListItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPPRSInformationListItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPPRSInformationListItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPRSInformationListItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPPRSInformationListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPPRSInformationListItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPRSInformationListItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPPRSInformationListItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPRxTEGInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPRxTEGInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPRxTEGInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPRxTEGInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPRxTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPRxTEGInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPRxTEGInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPRxTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPRxTxTEGInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPRxTxTEGInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPRxTxTEGInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPRxTxTEGInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPRxTxTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPRxTxTEGInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPRxTxTEGInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPRxTxTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPTxTEGInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPTxTEGInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPTxTEGInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPTxTEGInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPTxTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPTxTEGInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPTxTEGInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPTxTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type TRPTEGItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	TRPTEGItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPTEGItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPTEGItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPTEGItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *TRPTEGItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPTEGItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPTEGItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type DLPRSResourceItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	DLPRSResourceItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSResourceItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *DLPRSResourceItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type UEReportingInformationExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	UEReportingInformationExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *UEReportingInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEReportingInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(UEReportingInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *UEReportingInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEReportingInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&UEReportingInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type UETxTEGAssociationItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	UETxTEGAssociationItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	UETxTimingErrorMargin *TimingErrorMargin // valueExt,referenceFieldValue:104
}

func (x *UETxTEGAssociationItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UETxTEGAssociationItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(UETxTEGAssociationItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.UETxTimingErrorMargin != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UETxTimingErrorMargin.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "UETxTimingErrorMargin marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *UETxTEGAssociationItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UETxTEGAssociationItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&UETxTEGAssociationItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 104 {
		// Read struct defined elsewhere (Pointer)
		x.UETxTimingErrorMargin = new(TimingErrorMargin)
		err = x.UETxTimingErrorMargin.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode UETxTimingErrorMargin error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type SRSResourceIDItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	SRSResourceIDItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSResourceIDItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceIDItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceIDItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *SRSResourceIDItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceIDItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceIDItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ULAoAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ULAoAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ULAoAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULAoAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ULAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ULAoAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULAoAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ULAoAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ULRTOAMeasurementExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ULRTOAMeasurementExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	ExtendedAdditionalPathList *ExtendedAdditionalPathList // refFieldVal:77
	TRPRxTEGInformation        *TRPRxTEGInformation        // valueExt,referenceFieldValue:86
}

func (x *ULRTOAMeasurementExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULRTOAMeasurementExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ULRTOAMeasurementExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.ExtendedAdditionalPathList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExtendedAdditionalPathList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ExtendedAdditionalPathList marshal failed")
		}
	} else if x.TRPRxTEGInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPRxTEGInformation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPRxTEGInformation marshal failed")
		}
	} else {
		return errors.Errorf("no open type value is present")
	}

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ULRTOAMeasurementExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULRTOAMeasurementExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ULRTOAMeasurementExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	if x.Id.Value == 77 {
		// Read struct defined elsewhere (Pointer)
		x.ExtendedAdditionalPathList = new(ExtendedAdditionalPathList)
		err = x.ExtendedAdditionalPathList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ExtendedAdditionalPathList error")
		}
	} else if x.Id.Value == 86 {
		// Read struct defined elsewhere (Pointer)
		x.TRPRxTEGInformation = new(TRPRxTEGInformation)
		err = x.TRPRxTEGInformation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPRxTEGInformation error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", x.Id.Value)
	}

	return nil
}

type ULSRSRSRPPExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ULSRSRSRPPExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ULSRSRSRPPExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULSRSRSRPPExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ULSRSRSRPPExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ULSRSRSRPPExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULSRSRSRPPExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ULSRSRSRPPExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type WLANMeasurementQuantitiesValueExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	WLANMeasurementQuantitiesValueExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *WLANMeasurementQuantitiesValueExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	WLANMeasurementQuantitiesValueExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(WLANMeasurementQuantitiesValueExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *WLANMeasurementQuantitiesValueExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	WLANMeasurementQuantitiesValueExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&WLANMeasurementQuantitiesValueExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type WLANMeasurementResultItemExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	WLANMeasurementResultItemExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *WLANMeasurementResultItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	WLANMeasurementResultItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(WLANMeasurementResultItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *WLANMeasurementResultItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	WLANMeasurementResultItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&WLANMeasurementResultItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}

type ZoAExtIEs struct {
	Id          *ProtocolIEID
	Criticality *Criticality
	//	ExtensionValue	ZoAExtIEsExtensionValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ZoAExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ZoAExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ZoAExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	// No defined open type values

	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "opentype marshal failed")
	}

	return nil
}

func (x *ZoAExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ZoAExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ZoAExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(ProtocolIEID)
	err = x.Id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Criticality = new(Criticality)
	err = x.Criticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Criticality error")
	}

	// mandatory field
	// Read Open Type
	var pdOpenType *aper.PerBitData
	var openTypeBytes []byte

	// Read Open Type Bytes
	openTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode opentype error"))
	}
	pdOpenType = aper.NewPerBitData(openTypeBytes)

	// dummy function to avoid unsed error for pdOpenType
	foo(pdOpenType)

	// Read Open Type - referenceField: Id
	// No defined open type values

	return nil
}
