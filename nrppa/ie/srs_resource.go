package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	SRSResourceNrofSRSPortsPresentPort1  aper.Enumerated = 0
	SRSResourceNrofSRSPortsPresentPorts2 aper.Enumerated = 1
	SRSResourceNrofSRSPortsPresentPorts4 aper.Enumerated = 2
)

const ( /* Enum Type */
	SRSResourceNrofSymbolsPresentN1 aper.Enumerated = 0
	SRSResourceNrofSymbolsPresentN2 aper.Enumerated = 1
	SRSResourceNrofSymbolsPresentN4 aper.Enumerated = 2
)

const ( /* Enum Type */
	SRSResourceRepetitionFactorPresentN1 aper.Enumerated = 0
	SRSResourceRepetitionFactorPresentN2 aper.Enumerated = 1
	SRSResourceRepetitionFactorPresentN4 aper.Enumerated = 2
)

const ( /* Enum Type */
	SRSResourceGroupOrSequenceHoppingPresentNeither         aper.Enumerated = 0
	SRSResourceGroupOrSequenceHoppingPresentGroupHopping    aper.Enumerated = 1
	SRSResourceGroupOrSequenceHoppingPresentSequenceHopping aper.Enumerated = 2
)

type SRSResource struct {
	SRSResourceID          *SRSResourceID
	NrofSRSPorts           *aper.Enumerated                             // valueLB:0,valueUB:2
	TransmissionComb       *TransmissionComb                            // valueLB:0,valueUB:2
	StartPosition          *int64                                       // valueLB:0,valueUB:13
	NrofSymbols            *aper.Enumerated                             // valueLB:0,valueUB:2
	RepetitionFactor       *aper.Enumerated                             // valueLB:0,valueUB:2
	FreqDomainPosition     *int64                                       // valueLB:0,valueUB:67
	FreqDomainShift        *int64                                       // valueLB:0,valueUB:268
	CSRS                   *int64                                       // valueLB:0,valueUB:63
	BSRS                   *int64                                       // valueLB:0,valueUB:3
	BHop                   *int64                                       // valueLB:0,valueUB:3
	GroupOrSequenceHopping *aper.Enumerated                             // valueLB:0,valueUB:2
	ResourceType           *ResourceType                                // valueLB:0,valueUB:3
	SequenceId             *int64                                       // valueLB:0,valueUB:1023
	IEExtensions           *ProtocolExtensionContainerSRSResourceExtIEs // optional
}

func (x *SRSResource) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceID == nil {
		return errors.Errorf("SRSResourceID is missing")
	}
	// mandatory field
	if x.NrofSRSPorts == nil {
		return errors.Errorf("NrofSRSPorts is missing")
	}
	// mandatory field
	if x.TransmissionComb == nil {
		return errors.Errorf("TransmissionComb is missing")
	}
	// mandatory field
	if x.StartPosition == nil {
		return errors.Errorf("StartPosition is missing")
	}
	// mandatory field
	if x.NrofSymbols == nil {
		return errors.Errorf("NrofSymbols is missing")
	}
	// mandatory field
	if x.RepetitionFactor == nil {
		return errors.Errorf("RepetitionFactor is missing")
	}
	// mandatory field
	if x.FreqDomainPosition == nil {
		return errors.Errorf("FreqDomainPosition is missing")
	}
	// mandatory field
	if x.FreqDomainShift == nil {
		return errors.Errorf("FreqDomainShift is missing")
	}
	// mandatory field
	if x.CSRS == nil {
		return errors.Errorf("CSRS is missing")
	}
	// mandatory field
	if x.BSRS == nil {
		return errors.Errorf("BSRS is missing")
	}
	// mandatory field
	if x.BHop == nil {
		return errors.Errorf("BHop is missing")
	}
	// mandatory field
	if x.GroupOrSequenceHopping == nil {
		return errors.Errorf("GroupOrSequenceHopping is missing")
	}
	// mandatory field
	if x.ResourceType == nil {
		return errors.Errorf("ResourceType is missing")
	}
	// mandatory field
	if x.SequenceId == nil {
		return errors.Errorf("SequenceId is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SRSResourceOptPresentFlag = append(SRSResourceOptPresentFlag, true)
	} else {
		SRSResourceOptPresentFlag = append(SRSResourceOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSResourceID marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.NrofSRSPorts), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TransmissionComb.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TransmissionComb marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 13
	err = pd.WriteInteger(*(x.StartPosition), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.NrofSymbols), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.RepetitionFactor), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 67
	err = pd.WriteInteger(*(x.FreqDomainPosition), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 268
	err = pd.WriteInteger(*(x.FreqDomainShift), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 63
	err = pd.WriteInteger(*(x.CSRS), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteInteger(*(x.BSRS), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3
	err = pd.WriteInteger(*(x.BHop), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.GroupOrSequenceHopping), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ResourceType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ResourceType marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1023
	err = pd.WriteInteger(*(x.SequenceId), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *SRSResource) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSResourceID = new(SRSResourceID)
	err = x.SRSResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSResourceID error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.NrofSRSPorts = new(aper.Enumerated)
	*(x.NrofSRSPorts), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TransmissionComb = new(TransmissionComb)
	err = x.TransmissionComb.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TransmissionComb error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 13
	x.StartPosition = new(int64)
	*(x.StartPosition), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.NrofSymbols = new(aper.Enumerated)
	*(x.NrofSymbols), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.RepetitionFactor = new(aper.Enumerated)
	*(x.RepetitionFactor), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 67
	x.FreqDomainPosition = new(int64)
	*(x.FreqDomainPosition), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 268
	x.FreqDomainShift = new(int64)
	*(x.FreqDomainShift), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 63
	x.CSRS = new(int64)
	*(x.CSRS), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3
	x.BSRS = new(int64)
	*(x.BSRS), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3
	x.BHop = new(int64)
	*(x.BHop), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.GroupOrSequenceHopping = new(aper.Enumerated)
	*(x.GroupOrSequenceHopping), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ResourceType = new(ResourceType)
	err = x.ResourceType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ResourceType error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1023
	x.SequenceId = new(int64)
	*(x.SequenceId), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if SRSResourceOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSResourceExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
