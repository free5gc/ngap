package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &AssistanceInformationFailureList{}

type AssistanceInformationFailureItem struct {
	PosSIBType   *PosSIBType                                                       // valueExt,valueLB:0,valueUB:38
	Outcome      *Outcome                                                          // valueExt,valueLB:0,valueUB:0
	IEExtensions *ProtocolExtensionContainerAssistanceInformationFailureListExtIEs // optional
}

func (x *AssistanceInformationFailureItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationFailureItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PosSIBType == nil {
		return errors.Errorf("PosSIBType is missing")
	}
	// mandatory field
	if x.Outcome == nil {
		return errors.Errorf("Outcome is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AssistanceInformationFailureItemOptPresentFlag = append(AssistanceInformationFailureItemOptPresentFlag, true)
	} else {
		AssistanceInformationFailureItemOptPresentFlag = append(AssistanceInformationFailureItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationFailureItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PosSIBType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PosSIBType marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Outcome.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Outcome marshal failed")
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

func (x *AssistanceInformationFailureItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationFailureItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationFailureItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PosSIBType = new(PosSIBType)
	err = x.PosSIBType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PosSIBType error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Outcome = new(Outcome)
	err = x.Outcome.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Outcome error")
	}

	// optional field (optPresentFlag index: 0)
	if AssistanceInformationFailureItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAssistanceInformationFailureListExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

type AssistanceInformationFailureList struct {
	List []AssistanceInformationFailureItem // valueExt,sizeLB:1,sizeUB:32
}

func (x *AssistanceInformationFailureList) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Sequence Of
	*sLb, *sUb = 1, 32
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *AssistanceInformationFailureList) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 1, 32
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []AssistanceInformationFailureItem{}
	for i := 0; i < int(numElementsList); i++ {
		var val AssistanceInformationFailureItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}

func (x *AssistanceInformationFailureList) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *AssistanceInformationFailureList) ReadIE(pd *aper.PerBitData) error {
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
