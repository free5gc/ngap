package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &TargetHomeENBID{}

type TargetHomeENBID struct {
	PLMNidentity   *PLMNIdentity
	HomeENBID      *aper.BitString                                  // sizeLB:28,sizeUB:28
	SelectedEPSTAI *EPSTAI                                          // valueExt
	IEExtensions   *ProtocolExtensionContainerTargetHomeENBIDExtIEs // optional
}

func (x *TargetHomeENBID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TargetHomeENBIDOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNidentity == nil {
		return errors.Errorf("PLMNidentity is missing")
	}
	// mandatory field
	if x.HomeENBID == nil {
		return errors.Errorf("HomeENBID is missing")
	}
	// mandatory field
	if x.SelectedEPSTAI == nil {
		return errors.Errorf("SelectedEPSTAI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TargetHomeENBIDOptPresentFlag = append(TargetHomeENBIDOptPresentFlag, true)
	} else {
		TargetHomeENBIDOptPresentFlag = append(TargetHomeENBIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TargetHomeENBIDOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNidentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNidentity marshal failed")
	}

	// Write BitString (Pointer)
	*sLb, *sUb = 28, 28
	err = pd.WriteBitString(*(x.HomeENBID), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SelectedEPSTAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SelectedEPSTAI marshal failed")
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

func (x *TargetHomeENBID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TargetHomeENBIDOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TargetHomeENBIDOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNidentity = new(PLMNIdentity)
	err = x.PLMNidentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNidentity error")
	}

	// mandatory field
	// Read BitString (Pointer)
	*sLb, *sUb = 28, 28
	x.HomeENBID = new(aper.BitString)
	*(x.HomeENBID), err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SelectedEPSTAI = new(EPSTAI)
	err = x.SelectedEPSTAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SelectedEPSTAI error")
	}

	// optional field (optPresentFlag index: 0)
	if TargetHomeENBIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTargetHomeENBIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *TargetHomeENBID) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.WriteSequencePreambleBitMap(optPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: id
	err = id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: criticality
	err = criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: value (open type)
	pdOpenType := aper.NewPerBitData(nil)
	err = x.Write(pdOpenType)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}

	return nil
}

func (x *TargetHomeENBID) ReadIE(pd *aper.PerBitData) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.ReadSequencePreambleBitMap(&optPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err = protocolIECriticality.Read(pd)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	// sequence element: value (open type)
	var pdOpenTypeBytes []byte
	pdOpenTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	pdOpenType := aper.NewPerBitData(pdOpenTypeBytes)
	err = x.Read(pdOpenType)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	return nil
}
