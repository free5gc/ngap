package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type OTDOAInformationTypeItemIEs struct {
	//	Value	OTDOAInformationTypeItemIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	OTDOAInformationTypeItem *OTDOAInformationTypeItem // valueExt,referenceFieldValue:10
}

func (x *OTDOAInformationTypeItemIEs) Id() *ProtocolIEID {
	if x.OTDOAInformationTypeItem != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOTDOAInformationTypeItem,
		}
	}
	return nil
}

func (x *OTDOAInformationTypeItemIEs) Criticality() *Criticality {
	if x.OTDOAInformationTypeItem != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *OTDOAInformationTypeItemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOAInformationTypeItemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOAInformationTypeItemIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.OTDOAInformationTypeItem != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OTDOAInformationTypeItem.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OTDOAInformationTypeItem marshal failed")
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

func (x *OTDOAInformationTypeItemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOAInformationTypeItemIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOAInformationTypeItemIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 10 {
		// Read struct defined elsewhere (Pointer)
		x.OTDOAInformationTypeItem = new(OTDOAInformationTypeItem)
		err = x.OTDOAInformationTypeItem.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OTDOAInformationTypeItem error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type SRSTypeExtIEs struct {
	//	Value	SRSTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SRSTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *SRSTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *SRSTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SRSTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *SRSTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SRSTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type AbortTransmissionExtIEs struct {
	//	Value	AbortTransmissionExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AbortTransmissionExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *AbortTransmissionExtIEs) Criticality() *Criticality {
	return nil
}

func (x *AbortTransmissionExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AbortTransmissionExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AbortTransmissionExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *AbortTransmissionExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AbortTransmissionExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AbortTransmissionExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type AngleMeasurementTypeExtIEs struct {
	//	Value	AngleMeasurementTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *AngleMeasurementTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *AngleMeasurementTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *AngleMeasurementTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AngleMeasurementTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AngleMeasurementTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *AngleMeasurementTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AngleMeasurementTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AngleMeasurementTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ARPLocationTypeExtIEs struct {
	//	Value	ARPLocationTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ARPLocationTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ARPLocationTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ARPLocationTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ARPLocationTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ARPLocationTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ARPLocationTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ARPLocationTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ARPLocationTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type BandwidthSRSExtIEs struct {
	//	Value	BandwidthSRSExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *BandwidthSRSExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *BandwidthSRSExtIEs) Criticality() *Criticality {
	return nil
}

func (x *BandwidthSRSExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	BandwidthSRSExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(BandwidthSRSExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *BandwidthSRSExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	BandwidthSRSExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&BandwidthSRSExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type CauseExtensionIE struct {
	//	Value	CauseExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *CauseExtensionIE) Id() *ProtocolIEID {
	return nil
}

func (x *CauseExtensionIE) Criticality() *Criticality {
	return nil
}

func (x *CauseExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CauseExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(CauseExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *CauseExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CauseExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&CauseExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type DLPRSMutingPatternExtIEs struct {
	//	Value	DLPRSMutingPatternExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSMutingPatternExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *DLPRSMutingPatternExtIEs) Criticality() *Criticality {
	return nil
}

func (x *DLPRSMutingPatternExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSMutingPatternExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSMutingPatternExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *DLPRSMutingPatternExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSMutingPatternExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSMutingPatternExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type DLPRSResourceSetARPLocationExtIEs struct {
	//	Value	DLPRSResourceSetARPLocationExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSResourceSetARPLocationExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *DLPRSResourceSetARPLocationExtIEs) Criticality() *Criticality {
	return nil
}

func (x *DLPRSResourceSetARPLocationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceSetARPLocationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceSetARPLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *DLPRSResourceSetARPLocationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceSetARPLocationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceSetARPLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type DLPRSResourceARPLocationExtIEs struct {
	//	Value	DLPRSResourceARPLocationExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *DLPRSResourceARPLocationExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *DLPRSResourceARPLocationExtIEs) Criticality() *Criticality {
	return nil
}

func (x *DLPRSResourceARPLocationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceARPLocationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceARPLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *DLPRSResourceARPLocationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceARPLocationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceARPLocationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type GNBRxTxTimeDiffMeasExtIEs struct {
	//	Value	GNBRxTxTimeDiffMeasExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *GNBRxTxTimeDiffMeasExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *GNBRxTxTimeDiffMeasExtIEs) Criticality() *Criticality {
	return nil
}

func (x *GNBRxTxTimeDiffMeasExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GNBRxTxTimeDiffMeasExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(GNBRxTxTimeDiffMeasExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *GNBRxTxTimeDiffMeasExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GNBRxTxTimeDiffMeasExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&GNBRxTxTimeDiffMeasExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type LoSNLoSInformationExtIEs struct {
	//	Value	LoSNLoSInformationExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *LoSNLoSInformationExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *LoSNLoSInformationExtIEs) Criticality() *Criticality {
	return nil
}

func (x *LoSNLoSInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LoSNLoSInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(LoSNLoSInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *LoSNLoSInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LoSNLoSInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&LoSNLoSInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type MeasurementQuantitiesItemIEs struct {
	//	Value	MeasurementQuantitiesItemIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	MeasurementQuantitiesItem *MeasurementQuantitiesItem // valueExt,referenceFieldValue:11
}

func (x *MeasurementQuantitiesItemIEs) Id() *ProtocolIEID {
	if x.MeasurementQuantitiesItem != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementQuantitiesItem,
		}
	}
	return nil
}

func (x *MeasurementQuantitiesItemIEs) Criticality() *Criticality {
	if x.MeasurementQuantitiesItem != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *MeasurementQuantitiesItemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementQuantitiesItemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementQuantitiesItemIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.MeasurementQuantitiesItem != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementQuantitiesItem.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementQuantitiesItem marshal failed")
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

func (x *MeasurementQuantitiesItemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementQuantitiesItemIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementQuantitiesItemIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 11 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementQuantitiesItem = new(MeasurementQuantitiesItem)
		err = x.MeasurementQuantitiesItem.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementQuantitiesItem error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasuredResultsValueExtensionIE struct {
	//	Value	MeasuredResultsValueExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	ResultSSRSRP     *ResultSSRSRP  // refFieldVal:32
	ResultSSRSRQ     *ResultSSRSRQ  // refFieldVal:33
	ResultCSIRSRP    *ResultCSIRSRP // refFieldVal:34
	ResultCSIRSRQ    *ResultCSIRSRQ // refFieldVal:35
	AngleOfArrivalNR *ULAoA         // valueExt,referenceFieldValue:36
	NRTADV           *NRTADV        // refFieldVal:94
}

func (x *MeasuredResultsValueExtensionIE) Id() *ProtocolIEID {
	if x.ResultSSRSRP != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResultSSRSRP,
		}
	}
	if x.ResultSSRSRQ != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResultSSRSRQ,
		}
	}
	if x.ResultCSIRSRP != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResultCSIRSRP,
		}
	}
	if x.ResultCSIRSRQ != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResultCSIRSRQ,
		}
	}
	if x.AngleOfArrivalNR != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDAngleOfArrivalNR,
		}
	}
	if x.NRTADV != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDNRTADV,
		}
	}
	return nil
}

func (x *MeasuredResultsValueExtensionIE) Criticality() *Criticality {
	if x.ResultSSRSRP != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.ResultSSRSRQ != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.ResultCSIRSRP != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.ResultCSIRSRQ != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.AngleOfArrivalNR != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.NRTADV != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasuredResultsValueExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasuredResultsValueExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasuredResultsValueExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.ResultSSRSRP != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResultSSRSRP.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResultSSRSRP marshal failed")
		}
	} else if x.ResultSSRSRQ != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResultSSRSRQ.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResultSSRSRQ marshal failed")
		}
	} else if x.ResultCSIRSRP != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResultCSIRSRP.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResultCSIRSRP marshal failed")
		}
	} else if x.ResultCSIRSRQ != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResultCSIRSRQ.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResultCSIRSRQ marshal failed")
		}
	} else if x.AngleOfArrivalNR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AngleOfArrivalNR.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "AngleOfArrivalNR marshal failed")
		}
	} else if x.NRTADV != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NRTADV.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "NRTADV marshal failed")
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

func (x *MeasuredResultsValueExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasuredResultsValueExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasuredResultsValueExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 32 {
		// Read struct defined elsewhere (Pointer)
		x.ResultSSRSRP = new(ResultSSRSRP)
		err = x.ResultSSRSRP.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResultSSRSRP error")
		}
	} else if id.Value == 33 {
		// Read struct defined elsewhere (Pointer)
		x.ResultSSRSRQ = new(ResultSSRSRQ)
		err = x.ResultSSRSRQ.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResultSSRSRQ error")
		}
	} else if id.Value == 34 {
		// Read struct defined elsewhere (Pointer)
		x.ResultCSIRSRP = new(ResultCSIRSRP)
		err = x.ResultCSIRSRP.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResultCSIRSRP error")
		}
	} else if id.Value == 35 {
		// Read struct defined elsewhere (Pointer)
		x.ResultCSIRSRQ = new(ResultCSIRSRQ)
		err = x.ResultCSIRSRQ.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResultCSIRSRQ error")
		}
	} else if id.Value == 36 {
		// Read struct defined elsewhere (Pointer)
		x.AngleOfArrivalNR = new(ULAoA)
		err = x.AngleOfArrivalNR.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode AngleOfArrivalNR error")
		}
	} else if id.Value == 94 {
		// Read struct defined elsewhere (Pointer)
		x.NRTADV = new(NRTADV)
		err = x.NRTADV.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode NRTADV error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MultipleULAoAItemExtIEs struct {
	//	Value	MultipleULAoAItemExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *MultipleULAoAItemExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *MultipleULAoAItemExtIEs) Criticality() *Criticality {
	return nil
}

func (x *MultipleULAoAItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MultipleULAoAItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MultipleULAoAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *MultipleULAoAItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MultipleULAoAItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MultipleULAoAItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type NGRANCellExtensionIE struct {
	//	Value	NGRANCellExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *NGRANCellExtensionIE) Id() *ProtocolIEID {
	return nil
}

func (x *NGRANCellExtensionIE) Criticality() *Criticality {
	return nil
}

func (x *NGRANCellExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANCellExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(NGRANCellExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *NGRANCellExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANCellExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&NGRANCellExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type OTDOACellInformationItemExtensionIE struct {
	//	Value	OTDOACellInformationItemExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TDDConfigEUTRAItem      *TDDConfigEUTRAItem         // valueExt,referenceFieldValue:22
	CGINR                   *CGINR                      // valueExt,referenceFieldValue:58
	SFNInitialisationTimeNR *SFNInitialisationTimeEUTRA // refFieldVal:59
}

func (x *OTDOACellInformationItemExtensionIE) Id() *ProtocolIEID {
	if x.TDDConfigEUTRAItem != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTDDConfigEUTRAItem,
		}
	}
	if x.CGINR != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCGINR,
		}
	}
	if x.SFNInitialisationTimeNR != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSFNInitialisationTimeNR,
		}
	}
	return nil
}

func (x *OTDOACellInformationItemExtensionIE) Criticality() *Criticality {
	if x.TDDConfigEUTRAItem != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CGINR != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SFNInitialisationTimeNR != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *OTDOACellInformationItemExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOACellInformationItemExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOACellInformationItemExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TDDConfigEUTRAItem != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TDDConfigEUTRAItem.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TDDConfigEUTRAItem marshal failed")
		}
	} else if x.CGINR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CGINR.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CGINR marshal failed")
		}
	} else if x.SFNInitialisationTimeNR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SFNInitialisationTimeNR.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SFNInitialisationTimeNR marshal failed")
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

func (x *OTDOACellInformationItemExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOACellInformationItemExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOACellInformationItemExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 22 {
		// Read struct defined elsewhere (Pointer)
		x.TDDConfigEUTRAItem = new(TDDConfigEUTRAItem)
		err = x.TDDConfigEUTRAItem.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TDDConfigEUTRAItem error")
		}
	} else if id.Value == 58 {
		// Read struct defined elsewhere (Pointer)
		x.CGINR = new(CGINR)
		err = x.CGINR.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CGINR error")
		}
	} else if id.Value == 59 {
		// Read struct defined elsewhere (Pointer)
		x.SFNInitialisationTimeNR = new(SFNInitialisationTimeEUTRA)
		err = x.SFNInitialisationTimeNR.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SFNInitialisationTimeNR error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type OtherRATMeasurementQuantitiesItemIEs struct {
	//	Value	OtherRATMeasurementQuantitiesItemIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	OtherRATMeasurementQuantitiesItem *OtherRATMeasurementQuantitiesItem // valueExt,referenceFieldValue:16
}

func (x *OtherRATMeasurementQuantitiesItemIEs) Id() *ProtocolIEID {
	if x.OtherRATMeasurementQuantitiesItem != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOtherRATMeasurementQuantitiesItem,
		}
	}
	return nil
}

func (x *OtherRATMeasurementQuantitiesItemIEs) Criticality() *Criticality {
	if x.OtherRATMeasurementQuantitiesItem != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *OtherRATMeasurementQuantitiesItemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OtherRATMeasurementQuantitiesItemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OtherRATMeasurementQuantitiesItemIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.OtherRATMeasurementQuantitiesItem != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OtherRATMeasurementQuantitiesItem.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OtherRATMeasurementQuantitiesItem marshal failed")
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

func (x *OtherRATMeasurementQuantitiesItemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OtherRATMeasurementQuantitiesItemIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OtherRATMeasurementQuantitiesItemIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 16 {
		// Read struct defined elsewhere (Pointer)
		x.OtherRATMeasurementQuantitiesItem = new(OtherRATMeasurementQuantitiesItem)
		err = x.OtherRATMeasurementQuantitiesItem.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OtherRATMeasurementQuantitiesItem error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type OtherRATMeasuredResultsValueExtensionIE struct {
	//	Value	OtherRATMeasuredResultsValueExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	ResultNR    *ResultNR    // refFieldVal:55
	ResultEUTRA *ResultEUTRA // refFieldVal:56
}

func (x *OtherRATMeasuredResultsValueExtensionIE) Id() *ProtocolIEID {
	if x.ResultNR != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResultNR,
		}
	}
	if x.ResultEUTRA != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResultEUTRA,
		}
	}
	return nil
}

func (x *OtherRATMeasuredResultsValueExtensionIE) Criticality() *Criticality {
	if x.ResultNR != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.ResultEUTRA != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *OtherRATMeasuredResultsValueExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OtherRATMeasuredResultsValueExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OtherRATMeasuredResultsValueExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.ResultNR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResultNR.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResultNR marshal failed")
		}
	} else if x.ResultEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResultEUTRA.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResultEUTRA marshal failed")
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

func (x *OtherRATMeasuredResultsValueExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OtherRATMeasuredResultsValueExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OtherRATMeasuredResultsValueExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 55 {
		// Read struct defined elsewhere (Pointer)
		x.ResultNR = new(ResultNR)
		err = x.ResultNR.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResultNR error")
		}
	} else if id.Value == 56 {
		// Read struct defined elsewhere (Pointer)
		x.ResultEUTRA = new(ResultEUTRA)
		err = x.ResultEUTRA.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResultEUTRA error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PathlossReferenceSignalExtensionIE struct {
	//	Value	PathlossReferenceSignalExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PathlossReferenceSignalExtensionIE) Id() *ProtocolIEID {
	return nil
}

func (x *PathlossReferenceSignalExtensionIE) Criticality() *Criticality {
	return nil
}

func (x *PathlossReferenceSignalExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PathlossReferenceSignalExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PathlossReferenceSignalExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *PathlossReferenceSignalExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PathlossReferenceSignalExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PathlossReferenceSignalExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type PosResourceSetTypeExtIEs struct {
	//	Value	PosResourceSetTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PosResourceSetTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *PosResourceSetTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *PosResourceSetTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosResourceSetTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PosResourceSetTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *PosResourceSetTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosResourceSetTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PosResourceSetTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type PRSMutingConfigurationEUTRAExtensionIE struct {
	//	Value	PRSMutingConfigurationEUTRAExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSMutingConfigurationEUTRAExtensionIE) Id() *ProtocolIEID {
	return nil
}

func (x *PRSMutingConfigurationEUTRAExtensionIE) Criticality() *Criticality {
	return nil
}

func (x *PRSMutingConfigurationEUTRAExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSMutingConfigurationEUTRAExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSMutingConfigurationEUTRAExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *PRSMutingConfigurationEUTRAExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSMutingConfigurationEUTRAExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSMutingConfigurationEUTRAExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type PRSResourceQCLInfoExtIEs struct {
	//	Value	PRSResourceQCLInfoExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSResourceQCLInfoExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *PRSResourceQCLInfoExtIEs) Criticality() *Criticality {
	return nil
}

func (x *PRSResourceQCLInfoExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceQCLInfoExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceQCLInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *PRSResourceQCLInfoExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceQCLInfoExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceQCLInfoExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type PRSTransmissionOffIndicationExtIEs struct {
	//	Value	PRSTransmissionOffIndicationExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PRSTransmissionOffIndicationExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *PRSTransmissionOffIndicationExtIEs) Criticality() *Criticality {
	return nil
}

func (x *PRSTransmissionOffIndicationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffIndicationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffIndicationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *PRSTransmissionOffIndicationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffIndicationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffIndicationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ReferenceSignalExtensionIE struct {
	//	Value	ReferenceSignalExtensionIEValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ReferenceSignalExtensionIE) Id() *ProtocolIEID {
	return nil
}

func (x *ReferenceSignalExtensionIE) Criticality() *Criticality {
	return nil
}

func (x *ReferenceSignalExtensionIE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ReferenceSignalExtensionIEOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ReferenceSignalExtensionIEOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ReferenceSignalExtensionIE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ReferenceSignalExtensionIEOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ReferenceSignalExtensionIEOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ReferencePointExtIEs struct {
	//	Value	ReferencePointExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ReferencePointExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ReferencePointExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ReferencePointExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ReferencePointExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ReferencePointExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ReferencePointExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ReferencePointExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ReferencePointExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type RelativePathDelayExtIEs struct {
	//	Value	RelativePathDelayExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *RelativePathDelayExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *RelativePathDelayExtIEs) Criticality() *Criticality {
	return nil
}

func (x *RelativePathDelayExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RelativePathDelayExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(RelativePathDelayExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *RelativePathDelayExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RelativePathDelayExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&RelativePathDelayExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ResourceSetTypeExtIEs struct {
	//	Value	ResourceSetTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceSetTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ResourceSetTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ResourceSetTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceSetTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceSetTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ResourceSetTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceSetTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceSetTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ResourceTypeExtIEs struct {
	//	Value	ResourceTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ResourceTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ResourceTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ResourceTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ResourceTypePosExtIEs struct {
	//	Value	ResourceTypePosExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ResourceTypePosExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ResourceTypePosExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ResourceTypePosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResourceTypePosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ResourceTypePosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ResourceTypePosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResourceTypePosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ResourceTypePosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type SpatialInformationPosExtIEs struct {
	//	Value	SpatialInformationPosExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SpatialInformationPosExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *SpatialInformationPosExtIEs) Criticality() *Criticality {
	return nil
}

func (x *SpatialInformationPosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialInformationPosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SpatialInformationPosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *SpatialInformationPosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialInformationPosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SpatialInformationPosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type SSBBurstPositionExtIEs struct {
	//	Value	SSBBurstPositionExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *SSBBurstPositionExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *SSBBurstPositionExtIEs) Criticality() *Criticality {
	return nil
}

func (x *SSBBurstPositionExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBBurstPositionExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(SSBBurstPositionExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *SSBBurstPositionExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBBurstPositionExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&SSBBurstPositionExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type StartRBIndexExtIEs struct {
	//	Value	StartRBIndexExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *StartRBIndexExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *StartRBIndexExtIEs) Criticality() *Criticality {
	return nil
}

func (x *StartRBIndexExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	StartRBIndexExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(StartRBIndexExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *StartRBIndexExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	StartRBIndexExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&StartRBIndexExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type TRPTEGInformationExtIEs struct {
	//	Value	TRPTEGInformationExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPTEGInformationExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *TRPTEGInformationExtIEs) Criticality() *Criticality {
	return nil
}

func (x *TRPTEGInformationExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPTEGInformationExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *TRPTEGInformationExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPTEGInformationExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPTEGInformationExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type TimeStampSlotIndexExtIEs struct {
	//	Value	TimeStampSlotIndexExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SCS480 *SCS480 // refFieldVal:119
	SCS960 *SCS960 // refFieldVal:120
}

func (x *TimeStampSlotIndexExtIEs) Id() *ProtocolIEID {
	if x.SCS480 != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSCS480,
		}
	}
	if x.SCS960 != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSCS960,
		}
	}
	return nil
}

func (x *TimeStampSlotIndexExtIEs) Criticality() *Criticality {
	if x.SCS480 != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.SCS960 != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *TimeStampSlotIndexExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TimeStampSlotIndexExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TimeStampSlotIndexExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SCS480 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SCS480.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SCS480 marshal failed")
		}
	} else if x.SCS960 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SCS960.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SCS960 marshal failed")
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

func (x *TimeStampSlotIndexExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TimeStampSlotIndexExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TimeStampSlotIndexExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 119 {
		// Read struct defined elsewhere (Pointer)
		x.SCS480 = new(SCS480)
		err = x.SCS480.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SCS480 error")
		}
	} else if id.Value == 120 {
		// Read struct defined elsewhere (Pointer)
		x.SCS960 = new(SCS960)
		err = x.SCS960.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SCS960 error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TransmissionCombExtIEs struct {
	//	Value	TransmissionCombExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TransmissionCombn8 *TransmissionCombn8 // refFieldVal:111
}

func (x *TransmissionCombExtIEs) Id() *ProtocolIEID {
	if x.TransmissionCombn8 != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTransmissionCombn8,
		}
	}
	return nil
}

func (x *TransmissionCombExtIEs) Criticality() *Criticality {
	if x.TransmissionCombn8 != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *TransmissionCombExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TransmissionCombExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TransmissionCombExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TransmissionCombn8 != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TransmissionCombn8.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TransmissionCombn8 marshal failed")
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

func (x *TransmissionCombExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TransmissionCombExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TransmissionCombExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 111 {
		// Read struct defined elsewhere (Pointer)
		x.TransmissionCombn8 = new(TransmissionCombn8)
		err = x.TransmissionCombn8.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TransmissionCombn8 error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TransmissionCombPosExtIEs struct {
	//	Value	TransmissionCombPosExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TransmissionCombPosExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *TransmissionCombPosExtIEs) Criticality() *Criticality {
	return nil
}

func (x *TransmissionCombPosExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TransmissionCombPosExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TransmissionCombPosExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *TransmissionCombPosExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TransmissionCombPosExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TransmissionCombPosExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ChoiceTRPBeamInfoItemExtIEs struct {
	//	Value	ChoiceTRPBeamInfoItemExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ChoiceTRPBeamInfoItemExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ChoiceTRPBeamInfoItemExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ChoiceTRPBeamInfoItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ChoiceTRPBeamInfoItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ChoiceTRPBeamInfoItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ChoiceTRPBeamInfoItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ChoiceTRPBeamInfoItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ChoiceTRPBeamInfoItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type TrpMeasuredResultsValueExtIEs struct {
	//	Value	TrpMeasuredResultsValueExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	ZoA           *ZoA           // valueExt,referenceFieldValue:71
	MultipleULAoA *MultipleULAoA // valueExt,referenceFieldValue:74
	ULSRSRSRPP    *ULSRSRSRPP    // valueExt,referenceFieldValue:75
}

func (x *TrpMeasuredResultsValueExtIEs) Id() *ProtocolIEID {
	if x.ZoA != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDZoA,
		}
	}
	if x.MultipleULAoA != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMultipleULAoA,
		}
	}
	if x.ULSRSRSRPP != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDULSRSRSRPP,
		}
	}
	return nil
}

func (x *TrpMeasuredResultsValueExtIEs) Criticality() *Criticality {
	if x.ZoA != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.MultipleULAoA != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ULSRSRSRPP != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *TrpMeasuredResultsValueExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasuredResultsValueExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasuredResultsValueExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.ZoA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ZoA.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ZoA marshal failed")
		}
	} else if x.MultipleULAoA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MultipleULAoA.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MultipleULAoA marshal failed")
		}
	} else if x.ULSRSRSRPP != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ULSRSRSRPP.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ULSRSRSRPP marshal failed")
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

func (x *TrpMeasuredResultsValueExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasuredResultsValueExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasuredResultsValueExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 71 {
		// Read struct defined elsewhere (Pointer)
		x.ZoA = new(ZoA)
		err = x.ZoA.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ZoA error")
		}
	} else if id.Value == 74 {
		// Read struct defined elsewhere (Pointer)
		x.MultipleULAoA = new(MultipleULAoA)
		err = x.MultipleULAoA.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MultipleULAoA error")
		}
	} else if id.Value == 75 {
		// Read struct defined elsewhere (Pointer)
		x.ULSRSRSRPP = new(ULSRSRSRPP)
		err = x.ULSRSRSRPP.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ULSRSRSRPP error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TrpMeasurementQualityExtIEs struct {
	//	Value	TrpMeasurementQualityExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TrpMeasurementQualityExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *TrpMeasurementQualityExtIEs) Criticality() *Criticality {
	return nil
}

func (x *TrpMeasurementQualityExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TrpMeasurementQualityExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TrpMeasurementQualityExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *TrpMeasurementQualityExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TrpMeasurementQualityExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TrpMeasurementQualityExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type TRPInformationTypeResponseItemExtIEs struct {
	//	Value	TRPInformationTypeResponseItemExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TRPType                   *TRPType                   // valueExt,referenceFieldValue:62
	OnDemandPRS               *OnDemandPRSInfo           // valueExt,referenceFieldValue:68
	TRPTxTEGAssociation       *TRPTxTEGAssociation       // refFieldVal:84
	TRPBeamAntennaInformation *TRPBeamAntennaInformation // valueExt,referenceFieldValue:93
}

func (x *TRPInformationTypeResponseItemExtIEs) Id() *ProtocolIEID {
	if x.TRPType != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPType,
		}
	}
	if x.OnDemandPRS != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOnDemandPRS,
		}
	}
	if x.TRPTxTEGAssociation != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPTxTEGAssociation,
		}
	}
	if x.TRPBeamAntennaInformation != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPBeamAntennaInformation,
		}
	}
	return nil
}

func (x *TRPInformationTypeResponseItemExtIEs) Criticality() *Criticality {
	if x.TRPType != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.OnDemandPRS != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.TRPTxTEGAssociation != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.TRPBeamAntennaInformation != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *TRPInformationTypeResponseItemExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationTypeResponseItemExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationTypeResponseItemExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TRPType != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPType.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPType marshal failed")
		}
	} else if x.OnDemandPRS != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OnDemandPRS.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OnDemandPRS marshal failed")
		}
	} else if x.TRPTxTEGAssociation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPTxTEGAssociation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPTxTEGAssociation marshal failed")
		}
	} else if x.TRPBeamAntennaInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPBeamAntennaInformation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPBeamAntennaInformation marshal failed")
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

func (x *TRPInformationTypeResponseItemExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationTypeResponseItemExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationTypeResponseItemExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 62 {
		// Read struct defined elsewhere (Pointer)
		x.TRPType = new(TRPType)
		err = x.TRPType.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPType error")
		}
	} else if id.Value == 68 {
		// Read struct defined elsewhere (Pointer)
		x.OnDemandPRS = new(OnDemandPRSInfo)
		err = x.OnDemandPRS.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OnDemandPRS error")
		}
	} else if id.Value == 84 {
		// Read struct defined elsewhere (Pointer)
		x.TRPTxTEGAssociation = new(TRPTxTEGAssociation)
		err = x.TRPTxTEGAssociation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPTxTEGAssociation error")
		}
	} else if id.Value == 93 {
		// Read struct defined elsewhere (Pointer)
		x.TRPBeamAntennaInformation = new(TRPBeamAntennaInformation)
		err = x.TRPBeamAntennaInformation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPBeamAntennaInformation error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TRPInformationTypeItemTRPReq struct {
	//	Value	TRPInformationTypeItemTRPReqValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TRPInformationTypeItem *TRPInformationTypeItem // valueExt,referenceFieldValue:57
}

func (x *TRPInformationTypeItemTRPReq) Id() *ProtocolIEID {
	if x.TRPInformationTypeItem != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPInformationTypeItem,
		}
	}
	return nil
}

func (x *TRPInformationTypeItemTRPReq) Criticality() *Criticality {
	if x.TRPInformationTypeItem != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *TRPInformationTypeItemTRPReq) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationTypeItemTRPReqOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationTypeItemTRPReqOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TRPInformationTypeItem != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPInformationTypeItem.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPInformationTypeItem marshal failed")
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

func (x *TRPInformationTypeItemTRPReq) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationTypeItemTRPReqOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationTypeItemTRPReqOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 57 {
		// Read struct defined elsewhere (Pointer)
		x.TRPInformationTypeItem = new(TRPInformationTypeItem)
		err = x.TRPInformationTypeItem.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPInformationTypeItem error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TRPPositionDefinitionTypeExtIEs struct {
	//	Value	TRPPositionDefinitionTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPPositionDefinitionTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *TRPPositionDefinitionTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *TRPPositionDefinitionTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPositionDefinitionTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPPositionDefinitionTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *TRPPositionDefinitionTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPositionDefinitionTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPPositionDefinitionTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type TRPPositionDirectAccuracyExtIEs struct {
	//	Value	TRPPositionDirectAccuracyExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPPositionDirectAccuracyExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *TRPPositionDirectAccuracyExtIEs) Criticality() *Criticality {
	return nil
}

func (x *TRPPositionDirectAccuracyExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPositionDirectAccuracyExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPPositionDirectAccuracyExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *TRPPositionDirectAccuracyExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPositionDirectAccuracyExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPPositionDirectAccuracyExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type TRPReferencePointTypeExtIEs struct {
	//	Value	TRPReferencePointTypeExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *TRPReferencePointTypeExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *TRPReferencePointTypeExtIEs) Criticality() *Criticality {
	return nil
}

func (x *TRPReferencePointTypeExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPReferencePointTypeExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPReferencePointTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *TRPReferencePointTypeExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPReferencePointTypeExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPReferencePointTypeExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type ULRTOAMeasExtIEs struct {
	//	Value	ULRTOAMeasExtIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *ULRTOAMeasExtIEs) Id() *ProtocolIEID {
	return nil
}

func (x *ULRTOAMeasExtIEs) Criticality() *Criticality {
	return nil
}

func (x *ULRTOAMeasExtIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULRTOAMeasExtIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ULRTOAMeasExtIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
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

func (x *ULRTOAMeasExtIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULRTOAMeasExtIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ULRTOAMeasExtIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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

type WLANMeasurementQuantitiesItemIEs struct {
	//	Value	WLANMeasurementQuantitiesItemIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	WLANMeasurementQuantitiesItem *WLANMeasurementQuantitiesItem // valueExt,referenceFieldValue:20
}

func (x *WLANMeasurementQuantitiesItemIEs) Id() *ProtocolIEID {
	if x.WLANMeasurementQuantitiesItem != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDWLANMeasurementQuantitiesItem,
		}
	}
	return nil
}

func (x *WLANMeasurementQuantitiesItemIEs) Criticality() *Criticality {
	if x.WLANMeasurementQuantitiesItem != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *WLANMeasurementQuantitiesItemIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	WLANMeasurementQuantitiesItemIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(WLANMeasurementQuantitiesItemIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.WLANMeasurementQuantitiesItem != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WLANMeasurementQuantitiesItem.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "WLANMeasurementQuantitiesItem marshal failed")
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

func (x *WLANMeasurementQuantitiesItemIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	WLANMeasurementQuantitiesItemIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&WLANMeasurementQuantitiesItemIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 20 {
		// Read struct defined elsewhere (Pointer)
		x.WLANMeasurementQuantitiesItem = new(WLANMeasurementQuantitiesItem)
		err = x.WLANMeasurementQuantitiesItem.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode WLANMeasurementQuantitiesItem error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ECIDMeasurementInitiationRequestIEs struct {
	//	Value	ECIDMeasurementInitiationRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFUEMeasurementID            *UEMeasurementID               // refFieldVal:2
	ReportCharacteristics         *ReportCharacteristics         // valueExt,referenceFieldValue:3
	MeasurementPeriodicity        *MeasurementPeriodicity        // valueExt,referenceFieldValue:4
	MeasurementQuantities         *MeasurementQuantities         // refFieldVal:5
	OtherRATMeasurementQuantities *OtherRATMeasurementQuantities // refFieldVal:15
	WLANMeasurementQuantities     *WLANMeasurementQuantities     // refFieldVal:19
	MeasurementPeriodicityNRAoA   *MeasurementPeriodicityNRAoA   // valueExt,referenceFieldValue:105
}

func (x *ECIDMeasurementInitiationRequestIEs) Id() *ProtocolIEID {
	if x.LMFUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFUEMeasurementID,
		}
	}
	if x.ReportCharacteristics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDReportCharacteristics,
		}
	}
	if x.MeasurementPeriodicity != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementPeriodicity,
		}
	}
	if x.MeasurementQuantities != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementQuantities,
		}
	}
	if x.OtherRATMeasurementQuantities != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOtherRATMeasurementQuantities,
		}
	}
	if x.WLANMeasurementQuantities != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDWLANMeasurementQuantities,
		}
	}
	if x.MeasurementPeriodicityNRAoA != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementPeriodicityNRAoA,
		}
	}
	return nil
}

func (x *ECIDMeasurementInitiationRequestIEs) Criticality() *Criticality {
	if x.LMFUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ReportCharacteristics != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.MeasurementPeriodicity != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.MeasurementQuantities != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.OtherRATMeasurementQuantities != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.WLANMeasurementQuantities != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementPeriodicityNRAoA != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *ECIDMeasurementInitiationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementInitiationRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementInitiationRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFUEMeasurementID marshal failed")
		}
	} else if x.ReportCharacteristics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ReportCharacteristics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ReportCharacteristics marshal failed")
		}
	} else if x.MeasurementPeriodicity != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementPeriodicity.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementPeriodicity marshal failed")
		}
	} else if x.MeasurementQuantities != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementQuantities.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementQuantities marshal failed")
		}
	} else if x.OtherRATMeasurementQuantities != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OtherRATMeasurementQuantities.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OtherRATMeasurementQuantities marshal failed")
		}
	} else if x.WLANMeasurementQuantities != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WLANMeasurementQuantities.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "WLANMeasurementQuantities marshal failed")
		}
	} else if x.MeasurementPeriodicityNRAoA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementPeriodicityNRAoA.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementPeriodicityNRAoA marshal failed")
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

func (x *ECIDMeasurementInitiationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementInitiationRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementInitiationRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 2 {
		// Read struct defined elsewhere (Pointer)
		x.LMFUEMeasurementID = new(UEMeasurementID)
		err = x.LMFUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFUEMeasurementID error")
		}
	} else if id.Value == 3 {
		// Read struct defined elsewhere (Pointer)
		x.ReportCharacteristics = new(ReportCharacteristics)
		err = x.ReportCharacteristics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ReportCharacteristics error")
		}
	} else if id.Value == 4 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementPeriodicity = new(MeasurementPeriodicity)
		err = x.MeasurementPeriodicity.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementPeriodicity error")
		}
	} else if id.Value == 5 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementQuantities = new(MeasurementQuantities)
		err = x.MeasurementQuantities.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementQuantities error")
		}
	} else if id.Value == 15 {
		// Read struct defined elsewhere (Pointer)
		x.OtherRATMeasurementQuantities = new(OtherRATMeasurementQuantities)
		err = x.OtherRATMeasurementQuantities.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OtherRATMeasurementQuantities error")
		}
	} else if id.Value == 19 {
		// Read struct defined elsewhere (Pointer)
		x.WLANMeasurementQuantities = new(WLANMeasurementQuantities)
		err = x.WLANMeasurementQuantities.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode WLANMeasurementQuantities error")
		}
	} else if id.Value == 105 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementPeriodicityNRAoA = new(MeasurementPeriodicityNRAoA)
		err = x.MeasurementPeriodicityNRAoA.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementPeriodicityNRAoA error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ECIDMeasurementInitiationResponseIEs struct {
	//	Value	ECIDMeasurementInitiationResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFUEMeasurementID        *UEMeasurementID           // refFieldVal:2
	RANUEMeasurementID        *UEMeasurementID           // refFieldVal:6
	ECIDMeasurementResult     *ECIDMeasurementResult     // valueExt,referenceFieldValue:7
	CriticalityDiagnostics    *CriticalityDiagnostics    // valueExt,referenceFieldValue:1
	CellPortionID             *CellPortionID             // refFieldVal:14
	OtherRATMeasurementResult *OtherRATMeasurementResult // refFieldVal:17
	WLANMeasurementResult     *WLANMeasurementResult     // refFieldVal:21
}

func (x *ECIDMeasurementInitiationResponseIEs) Id() *ProtocolIEID {
	if x.LMFUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFUEMeasurementID,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANUEMeasurementID,
		}
	}
	if x.ECIDMeasurementResult != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDECIDMeasurementResult,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	if x.CellPortionID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCellPortionID,
		}
	}
	if x.OtherRATMeasurementResult != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOtherRATMeasurementResult,
		}
	}
	if x.WLANMeasurementResult != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDWLANMeasurementResult,
		}
	}
	return nil
}

func (x *ECIDMeasurementInitiationResponseIEs) Criticality() *Criticality {
	if x.LMFUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ECIDMeasurementResult != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CellPortionID != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.OtherRATMeasurementResult != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.WLANMeasurementResult != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *ECIDMeasurementInitiationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementInitiationResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementInitiationResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFUEMeasurementID marshal failed")
		}
	} else if x.RANUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANUEMeasurementID marshal failed")
		}
	} else if x.ECIDMeasurementResult != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ECIDMeasurementResult.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ECIDMeasurementResult marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
		}
	} else if x.CellPortionID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CellPortionID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CellPortionID marshal failed")
		}
	} else if x.OtherRATMeasurementResult != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OtherRATMeasurementResult.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OtherRATMeasurementResult marshal failed")
		}
	} else if x.WLANMeasurementResult != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.WLANMeasurementResult.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "WLANMeasurementResult marshal failed")
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

func (x *ECIDMeasurementInitiationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementInitiationResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementInitiationResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 2 {
		// Read struct defined elsewhere (Pointer)
		x.LMFUEMeasurementID = new(UEMeasurementID)
		err = x.LMFUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFUEMeasurementID error")
		}
	} else if id.Value == 6 {
		// Read struct defined elsewhere (Pointer)
		x.RANUEMeasurementID = new(UEMeasurementID)
		err = x.RANUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANUEMeasurementID error")
		}
	} else if id.Value == 7 {
		// Read struct defined elsewhere (Pointer)
		x.ECIDMeasurementResult = new(ECIDMeasurementResult)
		err = x.ECIDMeasurementResult.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ECIDMeasurementResult error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else if id.Value == 14 {
		// Read struct defined elsewhere (Pointer)
		x.CellPortionID = new(CellPortionID)
		err = x.CellPortionID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CellPortionID error")
		}
	} else if id.Value == 17 {
		// Read struct defined elsewhere (Pointer)
		x.OtherRATMeasurementResult = new(OtherRATMeasurementResult)
		err = x.OtherRATMeasurementResult.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OtherRATMeasurementResult error")
		}
	} else if id.Value == 21 {
		// Read struct defined elsewhere (Pointer)
		x.WLANMeasurementResult = new(WLANMeasurementResult)
		err = x.WLANMeasurementResult.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode WLANMeasurementResult error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ECIDMeasurementInitiationFailureIEs struct {
	//	Value	ECIDMeasurementInitiationFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFUEMeasurementID     *UEMeasurementID        // refFieldVal:2
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *ECIDMeasurementInitiationFailureIEs) Id() *ProtocolIEID {
	if x.LMFUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFUEMeasurementID,
		}
	}
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *ECIDMeasurementInitiationFailureIEs) Criticality() *Criticality {
	if x.LMFUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *ECIDMeasurementInitiationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementInitiationFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementInitiationFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFUEMeasurementID marshal failed")
		}
	} else if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *ECIDMeasurementInitiationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementInitiationFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementInitiationFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 2 {
		// Read struct defined elsewhere (Pointer)
		x.LMFUEMeasurementID = new(UEMeasurementID)
		err = x.LMFUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFUEMeasurementID error")
		}
	} else if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ECIDMeasurementFailureIndicationIEs struct {
	//	Value	ECIDMeasurementFailureIndicationIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFUEMeasurementID *UEMeasurementID // refFieldVal:2
	RANUEMeasurementID *UEMeasurementID // refFieldVal:6
	Cause              *Cause           // refFieldVal:0,valueLB:0,valueUB:3
}

func (x *ECIDMeasurementFailureIndicationIEs) Id() *ProtocolIEID {
	if x.LMFUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFUEMeasurementID,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANUEMeasurementID,
		}
	}
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	return nil
}

func (x *ECIDMeasurementFailureIndicationIEs) Criticality() *Criticality {
	if x.LMFUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *ECIDMeasurementFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementFailureIndicationIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementFailureIndicationIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFUEMeasurementID marshal failed")
		}
	} else if x.RANUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANUEMeasurementID marshal failed")
		}
	} else if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
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

func (x *ECIDMeasurementFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementFailureIndicationIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementFailureIndicationIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 2 {
		// Read struct defined elsewhere (Pointer)
		x.LMFUEMeasurementID = new(UEMeasurementID)
		err = x.LMFUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFUEMeasurementID error")
		}
	} else if id.Value == 6 {
		// Read struct defined elsewhere (Pointer)
		x.RANUEMeasurementID = new(UEMeasurementID)
		err = x.RANUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANUEMeasurementID error")
		}
	} else if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ECIDMeasurementReportIEs struct {
	//	Value	ECIDMeasurementReportIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFUEMeasurementID    *UEMeasurementID       // refFieldVal:2
	RANUEMeasurementID    *UEMeasurementID       // refFieldVal:6
	ECIDMeasurementResult *ECIDMeasurementResult // valueExt,referenceFieldValue:7
	CellPortionID         *CellPortionID         // refFieldVal:14
}

func (x *ECIDMeasurementReportIEs) Id() *ProtocolIEID {
	if x.LMFUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFUEMeasurementID,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANUEMeasurementID,
		}
	}
	if x.ECIDMeasurementResult != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDECIDMeasurementResult,
		}
	}
	if x.CellPortionID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCellPortionID,
		}
	}
	return nil
}

func (x *ECIDMeasurementReportIEs) Criticality() *Criticality {
	if x.LMFUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ECIDMeasurementResult != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CellPortionID != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *ECIDMeasurementReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementReportIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementReportIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFUEMeasurementID marshal failed")
		}
	} else if x.RANUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANUEMeasurementID marshal failed")
		}
	} else if x.ECIDMeasurementResult != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ECIDMeasurementResult.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ECIDMeasurementResult marshal failed")
		}
	} else if x.CellPortionID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CellPortionID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CellPortionID marshal failed")
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

func (x *ECIDMeasurementReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementReportIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementReportIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 2 {
		// Read struct defined elsewhere (Pointer)
		x.LMFUEMeasurementID = new(UEMeasurementID)
		err = x.LMFUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFUEMeasurementID error")
		}
	} else if id.Value == 6 {
		// Read struct defined elsewhere (Pointer)
		x.RANUEMeasurementID = new(UEMeasurementID)
		err = x.RANUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANUEMeasurementID error")
		}
	} else if id.Value == 7 {
		// Read struct defined elsewhere (Pointer)
		x.ECIDMeasurementResult = new(ECIDMeasurementResult)
		err = x.ECIDMeasurementResult.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ECIDMeasurementResult error")
		}
	} else if id.Value == 14 {
		// Read struct defined elsewhere (Pointer)
		x.CellPortionID = new(CellPortionID)
		err = x.CellPortionID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CellPortionID error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ECIDMeasurementTerminationCommandIEs struct {
	//	Value	ECIDMeasurementTerminationCommandIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFUEMeasurementID *UEMeasurementID // refFieldVal:2
	RANUEMeasurementID *UEMeasurementID // refFieldVal:6
}

func (x *ECIDMeasurementTerminationCommandIEs) Id() *ProtocolIEID {
	if x.LMFUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFUEMeasurementID,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANUEMeasurementID,
		}
	}
	return nil
}

func (x *ECIDMeasurementTerminationCommandIEs) Criticality() *Criticality {
	if x.LMFUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANUEMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *ECIDMeasurementTerminationCommandIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementTerminationCommandIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementTerminationCommandIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFUEMeasurementID marshal failed")
		}
	} else if x.RANUEMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANUEMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANUEMeasurementID marshal failed")
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

func (x *ECIDMeasurementTerminationCommandIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementTerminationCommandIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementTerminationCommandIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 2 {
		// Read struct defined elsewhere (Pointer)
		x.LMFUEMeasurementID = new(UEMeasurementID)
		err = x.LMFUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFUEMeasurementID error")
		}
	} else if id.Value == 6 {
		// Read struct defined elsewhere (Pointer)
		x.RANUEMeasurementID = new(UEMeasurementID)
		err = x.RANUEMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANUEMeasurementID error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type OTDOAInformationRequestIEs struct {
	//	Value	OTDOAInformationRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	OTDOAInformationTypeGroup *OTDOAInformationType // refFieldVal:9
}

func (x *OTDOAInformationRequestIEs) Id() *ProtocolIEID {
	if x.OTDOAInformationTypeGroup != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOTDOAInformationTypeGroup,
		}
	}
	return nil
}

func (x *OTDOAInformationRequestIEs) Criticality() *Criticality {
	if x.OTDOAInformationTypeGroup != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *OTDOAInformationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOAInformationRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOAInformationRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.OTDOAInformationTypeGroup != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OTDOAInformationTypeGroup.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OTDOAInformationTypeGroup marshal failed")
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

func (x *OTDOAInformationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOAInformationRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOAInformationRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 9 {
		// Read struct defined elsewhere (Pointer)
		x.OTDOAInformationTypeGroup = new(OTDOAInformationType)
		err = x.OTDOAInformationTypeGroup.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OTDOAInformationTypeGroup error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type OTDOAInformationResponseIEs struct {
	//	Value	OTDOAInformationResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	OTDOACells             *OTDOACells             // refFieldVal:8
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *OTDOAInformationResponseIEs) Id() *ProtocolIEID {
	if x.OTDOACells != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDOTDOACells,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *OTDOAInformationResponseIEs) Criticality() *Criticality {
	if x.OTDOACells != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *OTDOAInformationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOAInformationResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOAInformationResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.OTDOACells != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.OTDOACells.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "OTDOACells marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *OTDOAInformationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOAInformationResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOAInformationResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 8 {
		// Read struct defined elsewhere (Pointer)
		x.OTDOACells = new(OTDOACells)
		err = x.OTDOACells.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode OTDOACells error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type OTDOAInformationFailureIEs struct {
	//	Value	OTDOAInformationFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *OTDOAInformationFailureIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *OTDOAInformationFailureIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *OTDOAInformationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OTDOAInformationFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(OTDOAInformationFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *OTDOAInformationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OTDOAInformationFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&OTDOAInformationFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type AssistanceInformationControlIEs struct {
	//	Value	AssistanceInformationControlIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	AssistanceInformation     *AssistanceInformation     // valueExt,referenceFieldValue:23
	Broadcast                 *Broadcast                 // valueExt,referenceFieldValue:24
	PositioningBroadcastCells *PositioningBroadcastCells // refFieldVal:38
}

func (x *AssistanceInformationControlIEs) Id() *ProtocolIEID {
	if x.AssistanceInformation != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDAssistanceInformation,
		}
	}
	if x.Broadcast != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDBroadcast,
		}
	}
	if x.PositioningBroadcastCells != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPositioningBroadcastCells,
		}
	}
	return nil
}

func (x *AssistanceInformationControlIEs) Criticality() *Criticality {
	if x.AssistanceInformation != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.Broadcast != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.PositioningBroadcastCells != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *AssistanceInformationControlIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationControlIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationControlIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.AssistanceInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AssistanceInformation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "AssistanceInformation marshal failed")
		}
	} else if x.Broadcast != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Broadcast.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Broadcast marshal failed")
		}
	} else if x.PositioningBroadcastCells != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PositioningBroadcastCells.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PositioningBroadcastCells marshal failed")
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

func (x *AssistanceInformationControlIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationControlIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationControlIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 23 {
		// Read struct defined elsewhere (Pointer)
		x.AssistanceInformation = new(AssistanceInformation)
		err = x.AssistanceInformation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode AssistanceInformation error")
		}
	} else if id.Value == 24 {
		// Read struct defined elsewhere (Pointer)
		x.Broadcast = new(Broadcast)
		err = x.Broadcast.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Broadcast error")
		}
	} else if id.Value == 38 {
		// Read struct defined elsewhere (Pointer)
		x.PositioningBroadcastCells = new(PositioningBroadcastCells)
		err = x.PositioningBroadcastCells.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PositioningBroadcastCells error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type AssistanceInformationFeedbackIEs struct {
	//	Value	AssistanceInformationFeedbackIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	AssistanceInformationFailureList *AssistanceInformationFailureList // refFieldVal:25
	PositioningBroadcastCells        *PositioningBroadcastCells        // refFieldVal:38
	CriticalityDiagnostics           *CriticalityDiagnostics           // valueExt,referenceFieldValue:1
}

func (x *AssistanceInformationFeedbackIEs) Id() *ProtocolIEID {
	if x.AssistanceInformationFailureList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDAssistanceInformationFailureList,
		}
	}
	if x.PositioningBroadcastCells != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPositioningBroadcastCells,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *AssistanceInformationFeedbackIEs) Criticality() *Criticality {
	if x.AssistanceInformationFailureList != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.PositioningBroadcastCells != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *AssistanceInformationFeedbackIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationFeedbackIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationFeedbackIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.AssistanceInformationFailureList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AssistanceInformationFailureList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "AssistanceInformationFailureList marshal failed")
		}
	} else if x.PositioningBroadcastCells != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PositioningBroadcastCells.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PositioningBroadcastCells marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *AssistanceInformationFeedbackIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationFeedbackIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationFeedbackIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 25 {
		// Read struct defined elsewhere (Pointer)
		x.AssistanceInformationFailureList = new(AssistanceInformationFailureList)
		err = x.AssistanceInformationFailureList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode AssistanceInformationFailureList error")
		}
	} else if id.Value == 38 {
		// Read struct defined elsewhere (Pointer)
		x.PositioningBroadcastCells = new(PositioningBroadcastCells)
		err = x.PositioningBroadcastCells.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PositioningBroadcastCells error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type ErrorIndicationIEs struct {
	//	Value	ErrorIndicationIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *ErrorIndicationIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *ErrorIndicationIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *ErrorIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ErrorIndicationIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(ErrorIndicationIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *ErrorIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ErrorIndicationIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&ErrorIndicationIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningInformationRequestIEs struct {
	//	Value	PositioningInformationRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	RequestedSRSTransmissionCharacteristics *RequestedSRSTransmissionCharacteristics // valueExt,referenceFieldValue:12
	UEReportingInformation                  *UEReportingInformation                  // valueExt,referenceFieldValue:73
	UETEGInfoRequest                        *UETEGInfoRequest                        // valueExt,referenceFieldValue:90
	UETEGReportingPeriodicity               *UETEGReportingPeriodicity               // valueExt,referenceFieldValue:99
}

func (x *PositioningInformationRequestIEs) Id() *ProtocolIEID {
	if x.RequestedSRSTransmissionCharacteristics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRequestedSRSTransmissionCharacteristics,
		}
	}
	if x.UEReportingInformation != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDUEReportingInformation,
		}
	}
	if x.UETEGInfoRequest != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDUETEGInfoRequest,
		}
	}
	if x.UETEGReportingPeriodicity != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDUETEGReportingPeriodicity,
		}
	}
	return nil
}

func (x *PositioningInformationRequestIEs) Criticality() *Criticality {
	if x.RequestedSRSTransmissionCharacteristics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.UEReportingInformation != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.UETEGInfoRequest != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.UETEGReportingPeriodicity != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *PositioningInformationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningInformationRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningInformationRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.RequestedSRSTransmissionCharacteristics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RequestedSRSTransmissionCharacteristics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RequestedSRSTransmissionCharacteristics marshal failed")
		}
	} else if x.UEReportingInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UEReportingInformation.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "UEReportingInformation marshal failed")
		}
	} else if x.UETEGInfoRequest != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UETEGInfoRequest.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "UETEGInfoRequest marshal failed")
		}
	} else if x.UETEGReportingPeriodicity != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UETEGReportingPeriodicity.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "UETEGReportingPeriodicity marshal failed")
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

func (x *PositioningInformationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningInformationRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningInformationRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 12 {
		// Read struct defined elsewhere (Pointer)
		x.RequestedSRSTransmissionCharacteristics = new(RequestedSRSTransmissionCharacteristics)
		err = x.RequestedSRSTransmissionCharacteristics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RequestedSRSTransmissionCharacteristics error")
		}
	} else if id.Value == 73 {
		// Read struct defined elsewhere (Pointer)
		x.UEReportingInformation = new(UEReportingInformation)
		err = x.UEReportingInformation.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode UEReportingInformation error")
		}
	} else if id.Value == 90 {
		// Read struct defined elsewhere (Pointer)
		x.UETEGInfoRequest = new(UETEGInfoRequest)
		err = x.UETEGInfoRequest.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode UETEGInfoRequest error")
		}
	} else if id.Value == 99 {
		// Read struct defined elsewhere (Pointer)
		x.UETEGReportingPeriodicity = new(UETEGReportingPeriodicity)
		err = x.UETEGReportingPeriodicity.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode UETEGReportingPeriodicity error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningInformationResponseIEs struct {
	//	Value	PositioningInformationResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSConfiguration       *SRSConfiguration       // valueExt,referenceFieldValue:26
	SFNInitialisationTime  *RelativeTime1900       // refFieldVal:54
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
	UETxTEGAssociationList *UETxTEGAssociationList // refFieldVal:81
}

func (x *PositioningInformationResponseIEs) Id() *ProtocolIEID {
	if x.SRSConfiguration != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSRSConfiguration,
		}
	}
	if x.SFNInitialisationTime != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSFNInitialisationTime,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	if x.UETxTEGAssociationList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDUETxTEGAssociationList,
		}
	}
	return nil
}

func (x *PositioningInformationResponseIEs) Criticality() *Criticality {
	if x.SRSConfiguration != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SFNInitialisationTime != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.UETxTEGAssociationList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningInformationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningInformationResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningInformationResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSConfiguration.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSConfiguration marshal failed")
		}
	} else if x.SFNInitialisationTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SFNInitialisationTime.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SFNInitialisationTime marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
		}
	} else if x.UETxTEGAssociationList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UETxTEGAssociationList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "UETxTEGAssociationList marshal failed")
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

func (x *PositioningInformationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningInformationResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningInformationResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 26 {
		// Read struct defined elsewhere (Pointer)
		x.SRSConfiguration = new(SRSConfiguration)
		err = x.SRSConfiguration.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSConfiguration error")
		}
	} else if id.Value == 54 {
		// Read struct defined elsewhere (Pointer)
		x.SFNInitialisationTime = new(RelativeTime1900)
		err = x.SFNInitialisationTime.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SFNInitialisationTime error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else if id.Value == 81 {
		// Read struct defined elsewhere (Pointer)
		x.UETxTEGAssociationList = new(UETxTEGAssociationList)
		err = x.UETxTEGAssociationList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode UETxTEGAssociationList error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningInformationFailureIEs struct {
	//	Value	PositioningInformationFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *PositioningInformationFailureIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *PositioningInformationFailureIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningInformationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningInformationFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningInformationFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *PositioningInformationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningInformationFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningInformationFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningInformationUpdateIEs struct {
	//	Value	PositioningInformationUpdateIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSConfiguration       *SRSConfiguration       // valueExt,referenceFieldValue:26
	SFNInitialisationTime  *RelativeTime1900       // refFieldVal:54
	UETxTEGAssociationList *UETxTEGAssociationList // refFieldVal:81
	SRSTransmissionStatus  *SRSTransmissionStatus  // valueExt,referenceFieldValue:106
}

func (x *PositioningInformationUpdateIEs) Id() *ProtocolIEID {
	if x.SRSConfiguration != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSRSConfiguration,
		}
	}
	if x.SFNInitialisationTime != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSFNInitialisationTime,
		}
	}
	if x.UETxTEGAssociationList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDUETxTEGAssociationList,
		}
	}
	if x.SRSTransmissionStatus != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSRSTransmissionStatus,
		}
	}
	return nil
}

func (x *PositioningInformationUpdateIEs) Criticality() *Criticality {
	if x.SRSConfiguration != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SFNInitialisationTime != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.UETxTEGAssociationList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SRSTransmissionStatus != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningInformationUpdateIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningInformationUpdateIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningInformationUpdateIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSConfiguration.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSConfiguration marshal failed")
		}
	} else if x.SFNInitialisationTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SFNInitialisationTime.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SFNInitialisationTime marshal failed")
		}
	} else if x.UETxTEGAssociationList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UETxTEGAssociationList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "UETxTEGAssociationList marshal failed")
		}
	} else if x.SRSTransmissionStatus != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSTransmissionStatus.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSTransmissionStatus marshal failed")
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

func (x *PositioningInformationUpdateIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningInformationUpdateIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningInformationUpdateIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 26 {
		// Read struct defined elsewhere (Pointer)
		x.SRSConfiguration = new(SRSConfiguration)
		err = x.SRSConfiguration.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSConfiguration error")
		}
	} else if id.Value == 54 {
		// Read struct defined elsewhere (Pointer)
		x.SFNInitialisationTime = new(RelativeTime1900)
		err = x.SFNInitialisationTime.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SFNInitialisationTime error")
		}
	} else if id.Value == 81 {
		// Read struct defined elsewhere (Pointer)
		x.UETxTEGAssociationList = new(UETxTEGAssociationList)
		err = x.UETxTEGAssociationList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode UETxTEGAssociationList error")
		}
	} else if id.Value == 106 {
		// Read struct defined elsewhere (Pointer)
		x.SRSTransmissionStatus = new(SRSTransmissionStatus)
		err = x.SRSTransmissionStatus.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSTransmissionStatus error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementRequestIEs struct {
	//	Value	MeasurementRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID                           *MeasurementID                              // refFieldVal:39
	TRPMeasurementRequestList                  *TRPMeasurementRequestList                  // refFieldVal:41
	ReportCharacteristics                      *ReportCharacteristics                      // valueExt,referenceFieldValue:3
	MeasurementPeriodicity                     *MeasurementPeriodicity                     // valueExt,referenceFieldValue:4
	TRPMeasurementQuantities                   *TRPMeasurementQuantities                   // refFieldVal:52
	SFNInitialisationTime                      *RelativeTime1900                           // refFieldVal:54
	SRSConfiguration                           *SRSConfiguration                           // valueExt,referenceFieldValue:26
	MeasurementBeamInfoRequest                 *MeasurementBeamInfoRequest                 // valueExt,referenceFieldValue:31
	SystemFrameNumber                          *SystemFrameNumber                          // refFieldVal:49
	SlotNumber                                 *SlotNumber                                 // refFieldVal:50
	MeasurementPeriodicityExtended             *MeasurementPeriodicityExtended             // valueExt,referenceFieldValue:64
	ResponseTime                               *ResponseTime                               // valueExt,referenceFieldValue:72
	MeasurementCharacteristicsRequestIndicator *MeasurementCharacteristicsRequestIndicator // refFieldVal:92
	MeasurementTimeOccasion                    *MeasurementTimeOccasion                    // valueExt,referenceFieldValue:91
	MeasurementAmount                          *MeasurementAmount                          // refFieldVal:95
}

func (x *MeasurementRequestIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.TRPMeasurementRequestList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPMeasurementRequestList,
		}
	}
	if x.ReportCharacteristics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDReportCharacteristics,
		}
	}
	if x.MeasurementPeriodicity != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementPeriodicity,
		}
	}
	if x.TRPMeasurementQuantities != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPMeasurementQuantities,
		}
	}
	if x.SFNInitialisationTime != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSFNInitialisationTime,
		}
	}
	if x.SRSConfiguration != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSRSConfiguration,
		}
	}
	if x.MeasurementBeamInfoRequest != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementBeamInfoRequest,
		}
	}
	if x.SystemFrameNumber != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSystemFrameNumber,
		}
	}
	if x.SlotNumber != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSlotNumber,
		}
	}
	if x.MeasurementPeriodicityExtended != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementPeriodicityExtended,
		}
	}
	if x.ResponseTime != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDResponseTime,
		}
	}
	if x.MeasurementCharacteristicsRequestIndicator != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementCharacteristicsRequestIndicator,
		}
	}
	if x.MeasurementTimeOccasion != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementTimeOccasion,
		}
	}
	if x.MeasurementAmount != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementAmount,
		}
	}
	return nil
}

func (x *MeasurementRequestIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.TRPMeasurementRequestList != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ReportCharacteristics != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.MeasurementPeriodicity != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.TRPMeasurementQuantities != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.SFNInitialisationTime != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SRSConfiguration != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementBeamInfoRequest != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SystemFrameNumber != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SlotNumber != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementPeriodicityExtended != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ResponseTime != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementCharacteristicsRequestIndicator != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementTimeOccasion != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementAmount != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.TRPMeasurementRequestList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPMeasurementRequestList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPMeasurementRequestList marshal failed")
		}
	} else if x.ReportCharacteristics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ReportCharacteristics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ReportCharacteristics marshal failed")
		}
	} else if x.MeasurementPeriodicity != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementPeriodicity.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementPeriodicity marshal failed")
		}
	} else if x.TRPMeasurementQuantities != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPMeasurementQuantities.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPMeasurementQuantities marshal failed")
		}
	} else if x.SFNInitialisationTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SFNInitialisationTime.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SFNInitialisationTime marshal failed")
		}
	} else if x.SRSConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSConfiguration.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSConfiguration marshal failed")
		}
	} else if x.MeasurementBeamInfoRequest != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementBeamInfoRequest.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementBeamInfoRequest marshal failed")
		}
	} else if x.SystemFrameNumber != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SystemFrameNumber.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SystemFrameNumber marshal failed")
		}
	} else if x.SlotNumber != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SlotNumber.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SlotNumber marshal failed")
		}
	} else if x.MeasurementPeriodicityExtended != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementPeriodicityExtended.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementPeriodicityExtended marshal failed")
		}
	} else if x.ResponseTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ResponseTime.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ResponseTime marshal failed")
		}
	} else if x.MeasurementCharacteristicsRequestIndicator != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementCharacteristicsRequestIndicator.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementCharacteristicsRequestIndicator marshal failed")
		}
	} else if x.MeasurementTimeOccasion != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementTimeOccasion.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementTimeOccasion marshal failed")
		}
	} else if x.MeasurementAmount != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementAmount.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementAmount marshal failed")
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

func (x *MeasurementRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 41 {
		// Read struct defined elsewhere (Pointer)
		x.TRPMeasurementRequestList = new(TRPMeasurementRequestList)
		err = x.TRPMeasurementRequestList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPMeasurementRequestList error")
		}
	} else if id.Value == 3 {
		// Read struct defined elsewhere (Pointer)
		x.ReportCharacteristics = new(ReportCharacteristics)
		err = x.ReportCharacteristics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ReportCharacteristics error")
		}
	} else if id.Value == 4 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementPeriodicity = new(MeasurementPeriodicity)
		err = x.MeasurementPeriodicity.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementPeriodicity error")
		}
	} else if id.Value == 52 {
		// Read struct defined elsewhere (Pointer)
		x.TRPMeasurementQuantities = new(TRPMeasurementQuantities)
		err = x.TRPMeasurementQuantities.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPMeasurementQuantities error")
		}
	} else if id.Value == 54 {
		// Read struct defined elsewhere (Pointer)
		x.SFNInitialisationTime = new(RelativeTime1900)
		err = x.SFNInitialisationTime.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SFNInitialisationTime error")
		}
	} else if id.Value == 26 {
		// Read struct defined elsewhere (Pointer)
		x.SRSConfiguration = new(SRSConfiguration)
		err = x.SRSConfiguration.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSConfiguration error")
		}
	} else if id.Value == 31 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementBeamInfoRequest = new(MeasurementBeamInfoRequest)
		err = x.MeasurementBeamInfoRequest.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementBeamInfoRequest error")
		}
	} else if id.Value == 49 {
		// Read struct defined elsewhere (Pointer)
		x.SystemFrameNumber = new(SystemFrameNumber)
		err = x.SystemFrameNumber.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SystemFrameNumber error")
		}
	} else if id.Value == 50 {
		// Read struct defined elsewhere (Pointer)
		x.SlotNumber = new(SlotNumber)
		err = x.SlotNumber.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SlotNumber error")
		}
	} else if id.Value == 64 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementPeriodicityExtended = new(MeasurementPeriodicityExtended)
		err = x.MeasurementPeriodicityExtended.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementPeriodicityExtended error")
		}
	} else if id.Value == 72 {
		// Read struct defined elsewhere (Pointer)
		x.ResponseTime = new(ResponseTime)
		err = x.ResponseTime.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ResponseTime error")
		}
	} else if id.Value == 92 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementCharacteristicsRequestIndicator = new(MeasurementCharacteristicsRequestIndicator)
		err = x.MeasurementCharacteristicsRequestIndicator.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementCharacteristicsRequestIndicator error")
		}
	} else if id.Value == 91 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementTimeOccasion = new(MeasurementTimeOccasion)
		err = x.MeasurementTimeOccasion.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementTimeOccasion error")
		}
	} else if id.Value == 95 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementAmount = new(MeasurementAmount)
		err = x.MeasurementAmount.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementAmount error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementResponseIEs struct {
	//	Value	MeasurementResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID           *MeasurementID              // refFieldVal:39
	RANMeasurementID           *MeasurementID              // refFieldVal:40
	TRPMeasurementResponseList *TRPMeasurementResponseList // refFieldVal:42
	CriticalityDiagnostics     *CriticalityDiagnostics     // valueExt,referenceFieldValue:1
}

func (x *MeasurementResponseIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.RANMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANMeasurementID,
		}
	}
	if x.TRPMeasurementResponseList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPMeasurementResponseList,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *MeasurementResponseIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.TRPMeasurementResponseList != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.RANMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANMeasurementID marshal failed")
		}
	} else if x.TRPMeasurementResponseList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPMeasurementResponseList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPMeasurementResponseList marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *MeasurementResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 40 {
		// Read struct defined elsewhere (Pointer)
		x.RANMeasurementID = new(MeasurementID)
		err = x.RANMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANMeasurementID error")
		}
	} else if id.Value == 42 {
		// Read struct defined elsewhere (Pointer)
		x.TRPMeasurementResponseList = new(TRPMeasurementResponseList)
		err = x.TRPMeasurementResponseList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPMeasurementResponseList error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementFailureIEs struct {
	//	Value	MeasurementFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID       *MeasurementID          // refFieldVal:39
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *MeasurementFailureIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *MeasurementFailureIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *MeasurementFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementReportIEs struct {
	//	Value	MeasurementReportIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID         *MeasurementID              // refFieldVal:39
	RANMeasurementID         *MeasurementID              // refFieldVal:40
	TRPMeasurementReportList *TRPMeasurementResponseList // refFieldVal:43
}

func (x *MeasurementReportIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.RANMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANMeasurementID,
		}
	}
	if x.TRPMeasurementReportList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPMeasurementReportList,
		}
	}
	return nil
}

func (x *MeasurementReportIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.TRPMeasurementReportList != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *MeasurementReportIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementReportIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementReportIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.RANMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANMeasurementID marshal failed")
		}
	} else if x.TRPMeasurementReportList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPMeasurementReportList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPMeasurementReportList marshal failed")
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

func (x *MeasurementReportIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementReportIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementReportIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 40 {
		// Read struct defined elsewhere (Pointer)
		x.RANMeasurementID = new(MeasurementID)
		err = x.RANMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANMeasurementID error")
		}
	} else if id.Value == 43 {
		// Read struct defined elsewhere (Pointer)
		x.TRPMeasurementReportList = new(TRPMeasurementResponseList)
		err = x.TRPMeasurementReportList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPMeasurementReportList error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementUpdateIEs struct {
	//	Value	MeasurementUpdateIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID                           *MeasurementID                              // refFieldVal:39
	RANMeasurementID                           *MeasurementID                              // refFieldVal:40
	SRSConfiguration                           *SRSConfiguration                           // valueExt,referenceFieldValue:26
	TRPMeasurementUpdateList                   *TRPMeasurementUpdateList                   // refFieldVal:70
	MeasurementCharacteristicsRequestIndicator *MeasurementCharacteristicsRequestIndicator // refFieldVal:92
	MeasurementTimeOccasion                    *MeasurementTimeOccasion                    // valueExt,referenceFieldValue:91
}

func (x *MeasurementUpdateIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.RANMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANMeasurementID,
		}
	}
	if x.SRSConfiguration != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSRSConfiguration,
		}
	}
	if x.TRPMeasurementUpdateList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPMeasurementUpdateList,
		}
	}
	if x.MeasurementCharacteristicsRequestIndicator != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementCharacteristicsRequestIndicator,
		}
	}
	if x.MeasurementTimeOccasion != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDMeasurementTimeOccasion,
		}
	}
	return nil
}

func (x *MeasurementUpdateIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.SRSConfiguration != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.TRPMeasurementUpdateList != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.MeasurementCharacteristicsRequestIndicator != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.MeasurementTimeOccasion != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementUpdateIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementUpdateIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementUpdateIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.RANMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANMeasurementID marshal failed")
		}
	} else if x.SRSConfiguration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSConfiguration.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSConfiguration marshal failed")
		}
	} else if x.TRPMeasurementUpdateList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPMeasurementUpdateList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPMeasurementUpdateList marshal failed")
		}
	} else if x.MeasurementCharacteristicsRequestIndicator != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementCharacteristicsRequestIndicator.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementCharacteristicsRequestIndicator marshal failed")
		}
	} else if x.MeasurementTimeOccasion != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasurementTimeOccasion.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "MeasurementTimeOccasion marshal failed")
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

func (x *MeasurementUpdateIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementUpdateIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementUpdateIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 40 {
		// Read struct defined elsewhere (Pointer)
		x.RANMeasurementID = new(MeasurementID)
		err = x.RANMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANMeasurementID error")
		}
	} else if id.Value == 26 {
		// Read struct defined elsewhere (Pointer)
		x.SRSConfiguration = new(SRSConfiguration)
		err = x.SRSConfiguration.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSConfiguration error")
		}
	} else if id.Value == 70 {
		// Read struct defined elsewhere (Pointer)
		x.TRPMeasurementUpdateList = new(TRPMeasurementUpdateList)
		err = x.TRPMeasurementUpdateList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPMeasurementUpdateList error")
		}
	} else if id.Value == 92 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementCharacteristicsRequestIndicator = new(MeasurementCharacteristicsRequestIndicator)
		err = x.MeasurementCharacteristicsRequestIndicator.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementCharacteristicsRequestIndicator error")
		}
	} else if id.Value == 91 {
		// Read struct defined elsewhere (Pointer)
		x.MeasurementTimeOccasion = new(MeasurementTimeOccasion)
		err = x.MeasurementTimeOccasion.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode MeasurementTimeOccasion error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementAbortIEs struct {
	//	Value	MeasurementAbortIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID *MeasurementID // refFieldVal:39
	RANMeasurementID *MeasurementID // refFieldVal:40
}

func (x *MeasurementAbortIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.RANMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANMeasurementID,
		}
	}
	return nil
}

func (x *MeasurementAbortIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *MeasurementAbortIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementAbortIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementAbortIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.RANMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANMeasurementID marshal failed")
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

func (x *MeasurementAbortIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementAbortIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementAbortIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 40 {
		// Read struct defined elsewhere (Pointer)
		x.RANMeasurementID = new(MeasurementID)
		err = x.RANMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANMeasurementID error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementFailureIndicationIEs struct {
	//	Value	MeasurementFailureIndicationIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	LMFMeasurementID *MeasurementID // refFieldVal:39
	RANMeasurementID *MeasurementID // refFieldVal:40
	Cause            *Cause         // refFieldVal:0,valueLB:0,valueUB:3
}

func (x *MeasurementFailureIndicationIEs) Id() *ProtocolIEID {
	if x.LMFMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDLMFMeasurementID,
		}
	}
	if x.RANMeasurementID != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRANMeasurementID,
		}
	}
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	return nil
}

func (x *MeasurementFailureIndicationIEs) Criticality() *Criticality {
	if x.LMFMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.RANMeasurementID != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementFailureIndicationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementFailureIndicationIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementFailureIndicationIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.LMFMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LMFMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "LMFMeasurementID marshal failed")
		}
	} else if x.RANMeasurementID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANMeasurementID.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RANMeasurementID marshal failed")
		}
	} else if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
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

func (x *MeasurementFailureIndicationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementFailureIndicationIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementFailureIndicationIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 39 {
		// Read struct defined elsewhere (Pointer)
		x.LMFMeasurementID = new(MeasurementID)
		err = x.LMFMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode LMFMeasurementID error")
		}
	} else if id.Value == 40 {
		// Read struct defined elsewhere (Pointer)
		x.RANMeasurementID = new(MeasurementID)
		err = x.RANMeasurementID.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RANMeasurementID error")
		}
	} else if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TRPInformationRequestIEs struct {
	//	Value	TRPInformationRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TRPList                      *TRPList                      // refFieldVal:47
	TRPInformationTypeListTRPReq *TRPInformationTypeListTRPReq // refFieldVal:29
}

func (x *TRPInformationRequestIEs) Id() *ProtocolIEID {
	if x.TRPList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPList,
		}
	}
	if x.TRPInformationTypeListTRPReq != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPInformationTypeListTRPReq,
		}
	}
	return nil
}

func (x *TRPInformationRequestIEs) Criticality() *Criticality {
	if x.TRPList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.TRPInformationTypeListTRPReq != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	return nil
}

func (x *TRPInformationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TRPList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPList marshal failed")
		}
	} else if x.TRPInformationTypeListTRPReq != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPInformationTypeListTRPReq.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPInformationTypeListTRPReq marshal failed")
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

func (x *TRPInformationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 47 {
		// Read struct defined elsewhere (Pointer)
		x.TRPList = new(TRPList)
		err = x.TRPList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPList error")
		}
	} else if id.Value == 29 {
		// Read struct defined elsewhere (Pointer)
		x.TRPInformationTypeListTRPReq = new(TRPInformationTypeListTRPReq)
		err = x.TRPInformationTypeListTRPReq.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPInformationTypeListTRPReq error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TRPInformationResponseIEs struct {
	//	Value	TRPInformationResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TRPInformationListTRPResp *TRPInformationListTRPResp // refFieldVal:30
	CriticalityDiagnostics    *CriticalityDiagnostics    // valueExt,referenceFieldValue:1
}

func (x *TRPInformationResponseIEs) Id() *ProtocolIEID {
	if x.TRPInformationListTRPResp != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPInformationListTRPResp,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *TRPInformationResponseIEs) Criticality() *Criticality {
	if x.TRPInformationListTRPResp != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *TRPInformationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TRPInformationListTRPResp != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPInformationListTRPResp.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPInformationListTRPResp marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *TRPInformationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 30 {
		// Read struct defined elsewhere (Pointer)
		x.TRPInformationListTRPResp = new(TRPInformationListTRPResp)
		err = x.TRPInformationListTRPResp.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPInformationListTRPResp error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type TRPInformationFailureIEs struct {
	//	Value	TRPInformationFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *TRPInformationFailureIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *TRPInformationFailureIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *TRPInformationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPInformationFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(TRPInformationFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *TRPInformationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPInformationFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&TRPInformationFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningActivationRequestIEs struct {
	//	Value	PositioningActivationRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	SRSType        *SRSType          // refFieldVal:44,valueLB:0,valueUB:2
	ActivationTime *RelativeTime1900 // refFieldVal:45
}

func (x *PositioningActivationRequestIEs) Id() *ProtocolIEID {
	if x.SRSType != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSRSType,
		}
	}
	if x.ActivationTime != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDActivationTime,
		}
	}
	return nil
}

func (x *PositioningActivationRequestIEs) Criticality() *Criticality {
	if x.SRSType != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.ActivationTime != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningActivationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningActivationRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningActivationRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.SRSType != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SRSType.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SRSType marshal failed")
		}
	} else if x.ActivationTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ActivationTime.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "ActivationTime marshal failed")
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

func (x *PositioningActivationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningActivationRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningActivationRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 44 {
		// Read struct defined elsewhere (Pointer)
		x.SRSType = new(SRSType)
		err = x.SRSType.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SRSType error")
		}
	} else if id.Value == 45 {
		// Read struct defined elsewhere (Pointer)
		x.ActivationTime = new(RelativeTime1900)
		err = x.ActivationTime.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode ActivationTime error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningActivationResponseIEs struct {
	//	Value	PositioningActivationResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
	SystemFrameNumber      *SystemFrameNumber      // refFieldVal:49
	SlotNumber             *SlotNumber             // refFieldVal:50
}

func (x *PositioningActivationResponseIEs) Id() *ProtocolIEID {
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	if x.SystemFrameNumber != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSystemFrameNumber,
		}
	}
	if x.SlotNumber != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDSlotNumber,
		}
	}
	return nil
}

func (x *PositioningActivationResponseIEs) Criticality() *Criticality {
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SystemFrameNumber != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.SlotNumber != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningActivationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningActivationResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningActivationResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
		}
	} else if x.SystemFrameNumber != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SystemFrameNumber.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SystemFrameNumber marshal failed")
		}
	} else if x.SlotNumber != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SlotNumber.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "SlotNumber marshal failed")
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

func (x *PositioningActivationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningActivationResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningActivationResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else if id.Value == 49 {
		// Read struct defined elsewhere (Pointer)
		x.SystemFrameNumber = new(SystemFrameNumber)
		err = x.SystemFrameNumber.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SystemFrameNumber error")
		}
	} else if id.Value == 50 {
		// Read struct defined elsewhere (Pointer)
		x.SlotNumber = new(SlotNumber)
		err = x.SlotNumber.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode SlotNumber error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningActivationFailureIEs struct {
	//	Value	PositioningActivationFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *PositioningActivationFailureIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *PositioningActivationFailureIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningActivationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningActivationFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningActivationFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *PositioningActivationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningActivationFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningActivationFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PositioningDeactivationIEs struct {
	//	Value	PositioningDeactivationIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	AbortTransmission *AbortTransmission // refFieldVal:53,valueLB:0,valueUB:2
}

func (x *PositioningDeactivationIEs) Id() *ProtocolIEID {
	if x.AbortTransmission != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDAbortTransmission,
		}
	}
	return nil
}

func (x *PositioningDeactivationIEs) Criticality() *Criticality {
	if x.AbortTransmission != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PositioningDeactivationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PositioningDeactivationIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PositioningDeactivationIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.AbortTransmission != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AbortTransmission.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "AbortTransmission marshal failed")
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

func (x *PositioningDeactivationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PositioningDeactivationIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PositioningDeactivationIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 53 {
		// Read struct defined elsewhere (Pointer)
		x.AbortTransmission = new(AbortTransmission)
		err = x.AbortTransmission.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode AbortTransmission error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PRSConfigurationRequestIEs struct {
	//	Value	PRSConfigurationRequestIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	PRSConfigRequestType *PRSConfigRequestType // valueExt,referenceFieldValue:89
	PRSTRPList           *PRSTRPList           // refFieldVal:66
}

func (x *PRSConfigurationRequestIEs) Id() *ProtocolIEID {
	if x.PRSConfigRequestType != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPRSConfigRequestType,
		}
	}
	if x.PRSTRPList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPRSTRPList,
		}
	}
	return nil
}

func (x *PRSConfigurationRequestIEs) Criticality() *Criticality {
	if x.PRSConfigRequestType != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.PRSTRPList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PRSConfigurationRequestIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSConfigurationRequestIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSConfigurationRequestIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.PRSConfigRequestType != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSConfigRequestType.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PRSConfigRequestType marshal failed")
		}
	} else if x.PRSTRPList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSTRPList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PRSTRPList marshal failed")
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

func (x *PRSConfigurationRequestIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSConfigurationRequestIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSConfigurationRequestIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 89 {
		// Read struct defined elsewhere (Pointer)
		x.PRSConfigRequestType = new(PRSConfigRequestType)
		err = x.PRSConfigRequestType.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PRSConfigRequestType error")
		}
	} else if id.Value == 66 {
		// Read struct defined elsewhere (Pointer)
		x.PRSTRPList = new(PRSTRPList)
		err = x.PRSTRPList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PRSTRPList error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PRSConfigurationResponseIEs struct {
	//	Value	PRSConfigurationResponseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	PRSTransmissionTRPList *PRSTransmissionTRPList // refFieldVal:67
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *PRSConfigurationResponseIEs) Id() *ProtocolIEID {
	if x.PRSTransmissionTRPList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPRSTransmissionTRPList,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *PRSConfigurationResponseIEs) Criticality() *Criticality {
	if x.PRSTransmissionTRPList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PRSConfigurationResponseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSConfigurationResponseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSConfigurationResponseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.PRSTransmissionTRPList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSTransmissionTRPList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PRSTransmissionTRPList marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *PRSConfigurationResponseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSConfigurationResponseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSConfigurationResponseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 67 {
		// Read struct defined elsewhere (Pointer)
		x.PRSTransmissionTRPList = new(PRSTransmissionTRPList)
		err = x.PRSTransmissionTRPList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PRSTransmissionTRPList error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type PRSConfigurationFailureIEs struct {
	//	Value	PRSConfigurationFailureIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *PRSConfigurationFailureIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *PRSConfigurationFailureIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *PRSConfigurationFailureIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSConfigurationFailureIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PRSConfigurationFailureIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *PRSConfigurationFailureIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSConfigurationFailureIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PRSConfigurationFailureIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementPreconfigurationRequiredIEs struct {
	//	Value	MeasurementPreconfigurationRequiredIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	TRPPRSInformationList *TRPPRSInformationList // refFieldVal:87
}

func (x *MeasurementPreconfigurationRequiredIEs) Id() *ProtocolIEID {
	if x.TRPPRSInformationList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDTRPPRSInformationList,
		}
	}
	return nil
}

func (x *MeasurementPreconfigurationRequiredIEs) Criticality() *Criticality {
	if x.TRPPRSInformationList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementPreconfigurationRequiredIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementPreconfigurationRequiredIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementPreconfigurationRequiredIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.TRPPRSInformationList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPPRSInformationList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "TRPPRSInformationList marshal failed")
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

func (x *MeasurementPreconfigurationRequiredIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementPreconfigurationRequiredIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementPreconfigurationRequiredIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 87 {
		// Read struct defined elsewhere (Pointer)
		x.TRPPRSInformationList = new(TRPPRSInformationList)
		err = x.TRPPRSInformationList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode TRPPRSInformationList error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementPreconfigurationConfirmIEs struct {
	//	Value	MeasurementPreconfigurationConfirmIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	PreconfigurationResult *PreconfigurationResult // refFieldVal:97
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *MeasurementPreconfigurationConfirmIEs) Id() *ProtocolIEID {
	if x.PreconfigurationResult != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPreconfigurationResult,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *MeasurementPreconfigurationConfirmIEs) Criticality() *Criticality {
	if x.PreconfigurationResult != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementPreconfigurationConfirmIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementPreconfigurationConfirmIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementPreconfigurationConfirmIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.PreconfigurationResult != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PreconfigurationResult.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PreconfigurationResult marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *MeasurementPreconfigurationConfirmIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementPreconfigurationConfirmIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementPreconfigurationConfirmIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 97 {
		// Read struct defined elsewhere (Pointer)
		x.PreconfigurationResult = new(PreconfigurationResult)
		err = x.PreconfigurationResult.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PreconfigurationResult error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementPreconfigurationRefuseIEs struct {
	//	Value	MeasurementPreconfigurationRefuseIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	Cause                  *Cause                  // refFieldVal:0,valueLB:0,valueUB:3
	CriticalityDiagnostics *CriticalityDiagnostics // valueExt,referenceFieldValue:1
}

func (x *MeasurementPreconfigurationRefuseIEs) Id() *ProtocolIEID {
	if x.Cause != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCause,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDCriticalityDiagnostics,
		}
	}
	return nil
}

func (x *MeasurementPreconfigurationRefuseIEs) Criticality() *Criticality {
	if x.Cause != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	if x.CriticalityDiagnostics != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementPreconfigurationRefuseIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementPreconfigurationRefuseIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementPreconfigurationRefuseIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.Cause != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Cause.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "Cause marshal failed")
		}
	} else if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
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

func (x *MeasurementPreconfigurationRefuseIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementPreconfigurationRefuseIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementPreconfigurationRefuseIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 0 {
		// Read struct defined elsewhere (Pointer)
		x.Cause = new(Cause)
		err = x.Cause.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode Cause error")
		}
	} else if id.Value == 1 {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}

type MeasurementActivationIEs struct {
	//	Value	MeasurementActivationIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
	RequestType             *RequestType             // valueExt,referenceFieldValue:98
	PRSMeasurementsInfoList *PRSMeasurementsInfoList // refFieldVal:88
}

func (x *MeasurementActivationIEs) Id() *ProtocolIEID {
	if x.RequestType != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDRequestType,
		}
	}
	if x.PRSMeasurementsInfoList != nil {
		return &ProtocolIEID{
			Value: ProtocolIEIDPRSMeasurementsInfoList,
		}
	}
	return nil
}

func (x *MeasurementActivationIEs) Criticality() *Criticality {
	if x.RequestType != nil {
		return &Criticality{
			Value: CriticalityPresentReject,
		}
	}
	if x.PRSMeasurementsInfoList != nil {
		return &Criticality{
			Value: CriticalityPresentIgnore,
		}
	}
	return nil
}

func (x *MeasurementActivationIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementActivationIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id() == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality() == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementActivationIEsOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Id().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Id marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Criticality().Write(pd)
	if err != nil {
		return errors.Wrap(err, "Criticality marshal failed")
	}

	// Write Open Type
	pdOpenType := aper.NewPerBitData(nil)

	if x.RequestType != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RequestType.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "RequestType marshal failed")
		}
	} else if x.PRSMeasurementsInfoList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSMeasurementsInfoList.Write(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "PRSMeasurementsInfoList marshal failed")
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

func (x *MeasurementActivationIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementActivationIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&MeasurementActivationIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	id := new(ProtocolIEID)
	err = id.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Id error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	crit := new(Criticality)
	err = crit.Read(pd)
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
	if id.Value == 98 {
		// Read struct defined elsewhere (Pointer)
		x.RequestType = new(RequestType)
		err = x.RequestType.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode RequestType error")
		}
	} else if id.Value == 88 {
		// Read struct defined elsewhere (Pointer)
		x.PRSMeasurementsInfoList = new(PRSMeasurementsInfoList)
		err = x.PRSMeasurementsInfoList.Read(pdOpenType)
		if err != nil {
			return errors.Wrap(err, "decode PRSMeasurementsInfoList error")
		}
	} else {
		return errors.Errorf("unknown reference field value %d", id.Value)
	}

	return nil
}
