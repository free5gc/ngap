package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PosResourceSetTypeSemipersistentPossemiPersistentSetPresentTrue aper.Enumerated = 0
)

type PosResourceSetTypeSemiPersistent struct {
	PossemiPersistentSet *aper.Enumerated                                                  // valueExt,valueLB:0,valueUB:0
	IEExtensions         *ProtocolExtensionContainerPosResourceSetTypeSemiPersistentExtIEs // optional
}

func (x *PosResourceSetTypeSemiPersistent) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosResourceSetTypeSemiPersistentOptPresentFlag := []bool{}
	// mandatory field
	if x.PossemiPersistentSet == nil {
		return errors.Errorf("PossemiPersistentSet is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PosResourceSetTypeSemiPersistentOptPresentFlag = append(PosResourceSetTypeSemiPersistentOptPresentFlag, true)
	} else {
		PosResourceSetTypeSemiPersistentOptPresentFlag = append(PosResourceSetTypeSemiPersistentOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PosResourceSetTypeSemiPersistentOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	err = pd.WriteEnumerated(*(x.PossemiPersistentSet), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *PosResourceSetTypeSemiPersistent) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosResourceSetTypeSemiPersistentOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PosResourceSetTypeSemiPersistentOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 0
	x.PossemiPersistentSet = new(aper.Enumerated)
	*(x.PossemiPersistentSet), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if PosResourceSetTypeSemiPersistentOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPosResourceSetTypeSemiPersistentExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
