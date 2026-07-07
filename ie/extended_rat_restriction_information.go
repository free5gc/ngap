package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ExtendedRATRestrictionInformation struct {
	PrimaryRATRestriction   *aper.BitString                                                    // sizeExt,sizeLB:8,sizeUB:8
	SecondaryRATRestriction *aper.BitString                                                    // sizeExt,sizeLB:8,sizeUB:8
	IEExtensions            *ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs // optional
}

func (x *ExtendedRATRestrictionInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExtendedRATRestrictionInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.PrimaryRATRestriction == nil {
		return errors.Errorf("PrimaryRATRestriction is missing")
	}
	// mandatory field
	if x.SecondaryRATRestriction == nil {
		return errors.Errorf("SecondaryRATRestriction is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ExtendedRATRestrictionInformationOptPresentFlag = append(ExtendedRATRestrictionInformationOptPresentFlag, true)
	} else {
		ExtendedRATRestrictionInformationOptPresentFlag = append(ExtendedRATRestrictionInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExtendedRATRestrictionInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write BitString (Pointer)
	*sLb, *sUb = 8, 8
	err = pd.WriteBitString(*(x.PrimaryRATRestriction), true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}

	// Write BitString (Pointer)
	*sLb, *sUb = 8, 8
	err = pd.WriteBitString(*(x.SecondaryRATRestriction), true, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
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

func (x *ExtendedRATRestrictionInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExtendedRATRestrictionInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ExtendedRATRestrictionInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read BitString (Pointer)
	*sLb, *sUb = 8, 8
	x.PrimaryRATRestriction = new(aper.BitString)
	*(x.PrimaryRATRestriction), err = pd.ReadBitString(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}

	// mandatory field
	// Read BitString (Pointer)
	*sLb, *sUb = 8, 8
	x.SecondaryRATRestriction = new(aper.BitString)
	*(x.SecondaryRATRestriction), err = pd.ReadBitString(true, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}

	// optional field (optPresentFlag index: 0)
	if ExtendedRATRestrictionInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExtendedRATRestrictionInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
