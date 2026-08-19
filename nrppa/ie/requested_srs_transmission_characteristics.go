package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &RequestedSRSTransmissionCharacteristics{}

const ( /* Enum Type */
	RequestedSRSTransmissionCharacteristicsResourceTypePresentPeriodic       aper.Enumerated = 0
	RequestedSRSTransmissionCharacteristicsResourceTypePresentSemiPersistent aper.Enumerated = 1
	RequestedSRSTransmissionCharacteristicsResourceTypePresentAperiodic      aper.Enumerated = 2
)

type RequestedSRSTransmissionCharacteristics struct {
	NumberOfTransmissions *int64           // valueExt,valueLB:0,valueUB:500,optional
	ResourceType          *aper.Enumerated // valueExt,valueLB:0,valueUB:2
	Bandwidth             *BandwidthSRS    // valueLB:0,valueUB:2
	/* Sequence of = 35, FULL Name = struct RequestedSRSTransmissionCharacteristics__listOfSRSResourceSet */
	/* Type Name = SRSResourceSetItem */
	/* Sequence Of Embed */
	ListOfSRSResourceSet []SRSResourceSetItem                                                     // valueExt,sizeLB:1,sizeUB:16
	SSBInformation       *SSBInfo                                                                 // valueExt,optional
	IEExtensions         *ProtocolExtensionContainerRequestedSRSTransmissionCharacteristicsExtIEs // optional
}

func (x *RequestedSRSTransmissionCharacteristics) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedSRSTransmissionCharacteristicsOptPresentFlag := []bool{}
	// optional field
	if x.NumberOfTransmissions != nil {
		RequestedSRSTransmissionCharacteristicsOptPresentFlag = append(RequestedSRSTransmissionCharacteristicsOptPresentFlag, true)
	} else {
		RequestedSRSTransmissionCharacteristicsOptPresentFlag = append(RequestedSRSTransmissionCharacteristicsOptPresentFlag, false)
	}
	// mandatory field
	if x.ResourceType == nil {
		return errors.Errorf("ResourceType is missing")
	}
	// mandatory field
	if x.Bandwidth == nil {
		return errors.Errorf("Bandwidth is missing")
	}
	// mandatory field
	if x.ListOfSRSResourceSet == nil {
		return errors.Errorf("ListOfSRSResourceSet is missing")
	}
	// optional field
	if x.SSBInformation != nil {
		RequestedSRSTransmissionCharacteristicsOptPresentFlag = append(RequestedSRSTransmissionCharacteristicsOptPresentFlag, true)
	} else {
		RequestedSRSTransmissionCharacteristicsOptPresentFlag = append(RequestedSRSTransmissionCharacteristicsOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		RequestedSRSTransmissionCharacteristicsOptPresentFlag = append(RequestedSRSTransmissionCharacteristicsOptPresentFlag, true)
	} else {
		RequestedSRSTransmissionCharacteristicsOptPresentFlag = append(RequestedSRSTransmissionCharacteristicsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RequestedSRSTransmissionCharacteristicsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.NumberOfTransmissions != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 500
		err = pd.WriteInteger(*(x.NumberOfTransmissions), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.ResourceType), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.Bandwidth.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Bandwidth marshal failed")
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 16
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.ListOfSRSResourceSet)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.ListOfSRSResourceSet {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	// optional field
	if x.SSBInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SSBInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SSBInformation marshal failed")
		}
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

func (x *RequestedSRSTransmissionCharacteristics) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedSRSTransmissionCharacteristicsOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&RequestedSRSTransmissionCharacteristicsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if RequestedSRSTransmissionCharacteristicsOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 500
		x.NumberOfTransmissions = new(int64)
		*(x.NumberOfTransmissions), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.ResourceType = new(aper.Enumerated)
	*(x.ResourceType), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Bandwidth = new(BandwidthSRS)
	err = x.Bandwidth.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Bandwidth error")
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 16
	var numElementsListOfSRSResourceSet uint64
	numElementsListOfSRSResourceSet, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.ListOfSRSResourceSet = []SRSResourceSetItem{}
	for i := 0; i < int(numElementsListOfSRSResourceSet); i++ {
		var val SRSResourceSetItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.ListOfSRSResourceSet = append(x.ListOfSRSResourceSet, val)
		}
	}

	// optional field (optPresentFlag index: 1)
	if RequestedSRSTransmissionCharacteristicsOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SSBInformation = new(SSBInfo)
		err = x.SSBInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSBInformation error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if RequestedSRSTransmissionCharacteristicsOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRequestedSRSTransmissionCharacteristicsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *RequestedSRSTransmissionCharacteristics) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *RequestedSRSTransmissionCharacteristics) ReadIE(pd *aper.PerBitData) error {
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
