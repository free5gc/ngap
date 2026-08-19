package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &TRPBeamAntennaInformation{}

type TRPBeamAntennaInformation struct {
	ChoiceTRPBeamAntennaInfoItem *ChoiceTRPBeamAntennaInfoItem                              // valueLB:0,valueUB:3
	IEExtensions                 *ProtocolExtensionContainerTRPBeamAntennaInformationExtIEs // optional
}

func (x *TRPBeamAntennaInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamAntennaInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.ChoiceTRPBeamAntennaInfoItem == nil {
		return errors.Errorf("ChoiceTRPBeamAntennaInfoItem is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPBeamAntennaInformationOptPresentFlag = append(TRPBeamAntennaInformationOptPresentFlag, true)
	} else {
		TRPBeamAntennaInformationOptPresentFlag = append(TRPBeamAntennaInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamAntennaInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ChoiceTRPBeamAntennaInfoItem.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ChoiceTRPBeamAntennaInfoItem marshal failed")
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

func (x *TRPBeamAntennaInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamAntennaInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamAntennaInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ChoiceTRPBeamAntennaInfoItem = new(ChoiceTRPBeamAntennaInfoItem)
	err = x.ChoiceTRPBeamAntennaInfoItem.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ChoiceTRPBeamAntennaInfoItem error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPBeamAntennaInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPBeamAntennaInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *TRPBeamAntennaInformation) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *TRPBeamAntennaInformation) ReadIE(pd *aper.PerBitData) error {
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
