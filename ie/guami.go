package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &GUAMI{}

type GUAMI struct {
	PLMNIdentity *PLMNIdentity
	AMFRegionID  *AMFRegionID
	AMFSetID     *AMFSetID
	AMFPointer   *AMFPointer
	IEExtensions *ProtocolExtensionContainerGUAMIExtIEs // optional
}

func (x *GUAMI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GUAMIOptPresentFlag := []bool{}
	// mandatory field
	if x.PLMNIdentity == nil {
		return errors.Errorf("PLMNIdentity is missing")
	}
	// mandatory field
	if x.AMFRegionID == nil {
		return errors.Errorf("AMFRegionID is missing")
	}
	// mandatory field
	if x.AMFSetID == nil {
		return errors.Errorf("AMFSetID is missing")
	}
	// mandatory field
	if x.AMFPointer == nil {
		return errors.Errorf("AMFPointer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		GUAMIOptPresentFlag = append(GUAMIOptPresentFlag, true)
	} else {
		GUAMIOptPresentFlag = append(GUAMIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GUAMIOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PLMNIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PLMNIdentity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AMFRegionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AMFRegionID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AMFSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AMFSetID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.AMFPointer.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AMFPointer marshal failed")
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

func (x *GUAMI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GUAMIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&GUAMIOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PLMNIdentity = new(PLMNIdentity)
	err = x.PLMNIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PLMNIdentity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AMFRegionID = new(AMFRegionID)
	err = x.AMFRegionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AMFRegionID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AMFSetID = new(AMFSetID)
	err = x.AMFSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AMFSetID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AMFPointer = new(AMFPointer)
	err = x.AMFPointer.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AMFPointer error")
	}

	// optional field (optPresentFlag index: 0)
	if GUAMIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGUAMIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *GUAMI) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *GUAMI) ReadIE(pd *aper.PerBitData) error {
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
