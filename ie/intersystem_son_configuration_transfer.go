package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &IntersystemSONConfigurationTransfer{}

type IntersystemSONConfigurationTransfer struct {
	TransferType              *IntersystemSONTransferType                                          // valueLB:0,valueUB:2
	IntersystemSONInformation *IntersystemSONInformation                                           // valueLB:0,valueUB:1
	IEExtensions              *ProtocolExtensionContainerIntersystemSONConfigurationTransferExtIEs // optional
}

func (x *IntersystemSONConfigurationTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemSONConfigurationTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.TransferType == nil {
		return errors.Errorf("TransferType is missing")
	}
	// mandatory field
	if x.IntersystemSONInformation == nil {
		return errors.Errorf("IntersystemSONInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemSONConfigurationTransferOptPresentFlag = append(IntersystemSONConfigurationTransferOptPresentFlag, true)
	} else {
		IntersystemSONConfigurationTransferOptPresentFlag = append(IntersystemSONConfigurationTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemSONConfigurationTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TransferType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TransferType marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.IntersystemSONInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IntersystemSONInformation marshal failed")
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

func (x *IntersystemSONConfigurationTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemSONConfigurationTransferOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemSONConfigurationTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TransferType = new(IntersystemSONTransferType)
	err = x.TransferType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TransferType error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IntersystemSONInformation = new(IntersystemSONInformation)
	err = x.IntersystemSONInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IntersystemSONInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemSONConfigurationTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemSONConfigurationTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *IntersystemSONConfigurationTransfer) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *IntersystemSONConfigurationTransfer) ReadIE(pd *aper.PerBitData) error {
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
