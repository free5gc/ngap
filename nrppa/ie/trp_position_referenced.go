package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPPositionReferenced struct {
	ReferencePoint     *ReferencePoint                                        // valueLB:0,valueUB:3
	ReferencePointType *TRPReferencePointType                                 // valueLB:0,valueUB:2
	IEExtensions       *ProtocolExtensionContainerTRPPositionReferencedExtIEs // optional
}

func (x *TRPPositionReferenced) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPositionReferencedOptPresentFlag := []bool{}
	// mandatory field
	if x.ReferencePoint == nil {
		return errors.Errorf("ReferencePoint is missing")
	}
	// mandatory field
	if x.ReferencePointType == nil {
		return errors.Errorf("ReferencePointType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPPositionReferencedOptPresentFlag = append(TRPPositionReferencedOptPresentFlag, true)
	} else {
		TRPPositionReferencedOptPresentFlag = append(TRPPositionReferencedOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPPositionReferencedOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ReferencePoint.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReferencePoint marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ReferencePointType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReferencePointType marshal failed")
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

func (x *TRPPositionReferenced) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPositionReferencedOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPPositionReferencedOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReferencePoint = new(ReferencePoint)
	err = x.ReferencePoint.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReferencePoint error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReferencePointType = new(TRPReferencePointType)
	err = x.ReferencePointType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReferencePointType error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPPositionReferencedOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPPositionReferencedExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
