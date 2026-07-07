package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	PosSRSResourceItemNrofSymbolsPresentN1  aper.Enumerated = 0
	PosSRSResourceItemNrofSymbolsPresentN2  aper.Enumerated = 1
	PosSRSResourceItemNrofSymbolsPresentN4  aper.Enumerated = 2
	PosSRSResourceItemNrofSymbolsPresentN8  aper.Enumerated = 3
	PosSRSResourceItemNrofSymbolsPresentN12 aper.Enumerated = 4
)

const ( /* Enum Type */
	PosSRSResourceItemGroupOrSequenceHoppingPresentNeither         aper.Enumerated = 0
	PosSRSResourceItemGroupOrSequenceHoppingPresentGroupHopping    aper.Enumerated = 1
	PosSRSResourceItemGroupOrSequenceHoppingPresentSequenceHopping aper.Enumerated = 2
)

type PosSRSResourceItem struct {
	SrsPosResourceId       *SRSPosResourceID
	TransmissionCombPos    *TransmissionCombPos                                // valueLB:0,valueUB:3
	StartPosition          *int64                                              // valueLB:0,valueUB:13
	NrofSymbols            *aper.Enumerated                                    // valueLB:0,valueUB:4
	FreqDomainShift        *int64                                              // valueLB:0,valueUB:268
	CSRS                   *int64                                              // valueLB:0,valueUB:63
	GroupOrSequenceHopping *aper.Enumerated                                    // valueLB:0,valueUB:2
	ResourceTypePos        *ResourceTypePos                                    // valueLB:0,valueUB:3
	SequenceId             *int64                                              // valueLB:0,valueUB:65535
	SpatialRelationPos     *SpatialRelationPos                                 // valueLB:0,valueUB:2,optional
	IEExtensions           *ProtocolExtensionContainerPosSRSResourceItemExtIEs // optional
}

func (x *PosSRSResourceItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSRSResourceItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SrsPosResourceId == nil {
		return errors.Errorf("SrsPosResourceId is missing")
	}
	// mandatory field
	if x.TransmissionCombPos == nil {
		return errors.Errorf("TransmissionCombPos is missing")
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
	if x.FreqDomainShift == nil {
		return errors.Errorf("FreqDomainShift is missing")
	}
	// mandatory field
	if x.CSRS == nil {
		return errors.Errorf("CSRS is missing")
	}
	// mandatory field
	if x.GroupOrSequenceHopping == nil {
		return errors.Errorf("GroupOrSequenceHopping is missing")
	}
	// mandatory field
	if x.ResourceTypePos == nil {
		return errors.Errorf("ResourceTypePos is missing")
	}
	// mandatory field
	if x.SequenceId == nil {
		return errors.Errorf("SequenceId is missing")
	}
	// optional field
	if x.SpatialRelationPos != nil {
		PosSRSResourceItemOptPresentFlag = append(PosSRSResourceItemOptPresentFlag, true)
	} else {
		PosSRSResourceItemOptPresentFlag = append(PosSRSResourceItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PosSRSResourceItemOptPresentFlag = append(PosSRSResourceItemOptPresentFlag, true)
	} else {
		PosSRSResourceItemOptPresentFlag = append(PosSRSResourceItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PosSRSResourceItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SrsPosResourceId.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SrsPosResourceId marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TransmissionCombPos.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TransmissionCombPos marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 13
	err = pd.WriteInteger(*(x.StartPosition), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 4
	err = pd.WriteEnumerated(*(x.NrofSymbols), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	err = pd.WriteEnumerated(*(x.GroupOrSequenceHopping), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ResourceTypePos.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ResourceTypePos marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 65535
	err = pd.WriteInteger(*(x.SequenceId), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.SpatialRelationPos != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SpatialRelationPos.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SpatialRelationPos marshal failed")
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

func (x *PosSRSResourceItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSRSResourceItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PosSRSResourceItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SrsPosResourceId = new(SRSPosResourceID)
	err = x.SrsPosResourceId.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SrsPosResourceId error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TransmissionCombPos = new(TransmissionCombPos)
	err = x.TransmissionCombPos.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TransmissionCombPos error")
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
	*vLb, *vUb = 0, 4
	x.NrofSymbols = new(aper.Enumerated)
	*(x.NrofSymbols), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
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
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 2
	x.GroupOrSequenceHopping = new(aper.Enumerated)
	*(x.GroupOrSequenceHopping), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ResourceTypePos = new(ResourceTypePos)
	err = x.ResourceTypePos.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ResourceTypePos error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 65535
	x.SequenceId = new(int64)
	*(x.SequenceId), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if PosSRSResourceItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SpatialRelationPos = new(SpatialRelationPos)
		err = x.SpatialRelationPos.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SpatialRelationPos error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PosSRSResourceItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPosSRSResourceItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
