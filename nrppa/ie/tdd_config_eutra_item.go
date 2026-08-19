package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &TDDConfigEUTRAItem{}

const ( /* Enum Type */
	TDDConfigEUTRAItemSubframeAssignmentPresentSa0 aper.Enumerated = 0
	TDDConfigEUTRAItemSubframeAssignmentPresentSa1 aper.Enumerated = 1
	TDDConfigEUTRAItemSubframeAssignmentPresentSa2 aper.Enumerated = 2
	TDDConfigEUTRAItemSubframeAssignmentPresentSa3 aper.Enumerated = 3
	TDDConfigEUTRAItemSubframeAssignmentPresentSa4 aper.Enumerated = 4
	TDDConfigEUTRAItemSubframeAssignmentPresentSa5 aper.Enumerated = 5
	TDDConfigEUTRAItemSubframeAssignmentPresentSa6 aper.Enumerated = 6
)

type TDDConfigEUTRAItem struct {
	SubframeAssignment *aper.Enumerated                                        // valueExt,valueLB:0,valueUB:6
	IEExtensions       *ProtocolExtensionContainerTDDConfigEUTRAItemItemExtIEs // optional
}

func (x *TDDConfigEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TDDConfigEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SubframeAssignment == nil {
		return errors.Errorf("SubframeAssignment is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TDDConfigEUTRAItemOptPresentFlag = append(TDDConfigEUTRAItemOptPresentFlag, true)
	} else {
		TDDConfigEUTRAItemOptPresentFlag = append(TDDConfigEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TDDConfigEUTRAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 6
	err = pd.WriteEnumerated(*(x.SubframeAssignment), true, vLb, vUb)
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

func (x *TDDConfigEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TDDConfigEUTRAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TDDConfigEUTRAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 6
	x.SubframeAssignment = new(aper.Enumerated)
	*(x.SubframeAssignment), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if TDDConfigEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTDDConfigEUTRAItemItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *TDDConfigEUTRAItem) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *TDDConfigEUTRAItem) ReadIE(pd *aper.PerBitData) error {
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
