package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PosSIBSegmentsItem struct {
	AssistanceDataSIBelement *aper.OctetString
	IEExtensions             *ProtocolExtensionContainerPosSIBSegmentsExtIEs // optional
}

func (x *PosSIBSegmentsItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSIBSegmentsItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AssistanceDataSIBelement == nil {
		return errors.Errorf("AssistanceDataSIBelement is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PosSIBSegmentsItemOptPresentFlag = append(PosSIBSegmentsItemOptPresentFlag, true)
	} else {
		PosSIBSegmentsItemOptPresentFlag = append(PosSIBSegmentsItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PosSIBSegmentsItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write OctetString (Pointer)
	sLb, sUb = nil, nil
	err = pd.WriteOctetString(*(x.AssistanceDataSIBelement), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "octetString marshal failed")
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

func (x *PosSIBSegmentsItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSIBSegmentsItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PosSIBSegmentsItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read OctetString (Pointer)
	sLb, sUb = nil, nil
	x.AssistanceDataSIBelement = new(aper.OctetString)
	*(x.AssistanceDataSIBelement), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// optional field (optPresentFlag index: 0)
	if PosSIBSegmentsItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPosSIBSegmentsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

type PosSIBSegments struct {
	List []PosSIBSegmentsItem // valueExt,sizeLB:1,sizeUB:64
}

func (x *PosSIBSegments) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Sequence Of
	*sLb, *sUb = 1, 64
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

func (x *PosSIBSegments) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 1, 64
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []PosSIBSegmentsItem{}
	for i := 0; i < int(numElementsList); i++ {
		var val PosSIBSegmentsItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}
