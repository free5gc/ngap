package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &FiveGSTMSI{}

type FiveGSTMSI struct {
	AMFSetID     *AMFSetID
	AMFPointer   *AMFPointer
	FiveGTMSI    *FiveGTMSI
	IEExtensions *ProtocolExtensionContainerFiveGSTMSIExtIEs // optional
}

func (x *FiveGSTMSI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FiveGSTMSIOptPresentFlag := []bool{}
	// mandatory field
	if x.AMFSetID == nil {
		return errors.Errorf("AMFSetID is missing")
	}
	// mandatory field
	if x.AMFPointer == nil {
		return errors.Errorf("AMFPointer is missing")
	}
	// mandatory field
	if x.FiveGTMSI == nil {
		return errors.Errorf("FiveGTMSI is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		FiveGSTMSIOptPresentFlag = append(FiveGSTMSIOptPresentFlag, true)
	} else {
		FiveGSTMSIOptPresentFlag = append(FiveGSTMSIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FiveGSTMSIOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

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

	// Write struct defined elsewhere (Pointer)
	err = x.FiveGTMSI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveGTMSI marshal failed")
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

func (x *FiveGSTMSI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FiveGSTMSIOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&FiveGSTMSIOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

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

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveGTMSI = new(FiveGTMSI)
	err = x.FiveGTMSI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveGTMSI error")
	}

	// optional field (optPresentFlag index: 0)
	if FiveGSTMSIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFiveGSTMSIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *FiveGSTMSI) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *FiveGSTMSI) ReadIE(pd *aper.PerBitData) error {
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
