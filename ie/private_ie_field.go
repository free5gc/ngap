package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PrivateMessageIEs struct {
	Id          *PrivateIEID
	Criticality *Criticality
	//	Value	PrivateMessageIEsValue `aper:"openType,referenceFieldName:Id"`
	// Open Type Values
}

func (x *PrivateMessageIEs) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PrivateMessageIEsOptPresentFlag := []bool{}
	// mandatory field
	if x.Id == nil {
		return errors.Errorf("Id is missing")
	}
	// mandatory field
	if x.Criticality == nil {
		return errors.Errorf("Criticality is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PrivateMessageIEsOptPresentFlag, false)
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

func (x *PrivateMessageIEs) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PrivateMessageIEsOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PrivateMessageIEsOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Id = new(PrivateIEID)
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
