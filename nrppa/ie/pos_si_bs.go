package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PosSIBsItem struct {
	PosSIBType                    *PosSIBType // valueExt,valueLB:0,valueUB:38
	PosSIBSegments                *PosSIBSegments
	AssistanceInformationMetaData *AssistanceInformationMetaData           // valueExt,optional
	BroadcastPriority             *int64                                   // valueExt,valueLB:1,valueUB:16,optional
	IEExtensions                  *ProtocolExtensionContainerPosSIBsExtIEs // optional
}

func (x *PosSIBsItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSIBsItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PosSIBType == nil {
		return errors.Errorf("PosSIBType is missing")
	}
	// mandatory field
	if x.PosSIBSegments == nil {
		return errors.Errorf("PosSIBSegments is missing")
	}
	// optional field
	if x.AssistanceInformationMetaData != nil {
		PosSIBsItemOptPresentFlag = append(PosSIBsItemOptPresentFlag, true)
	} else {
		PosSIBsItemOptPresentFlag = append(PosSIBsItemOptPresentFlag, false)
	}
	// optional field
	if x.BroadcastPriority != nil {
		PosSIBsItemOptPresentFlag = append(PosSIBsItemOptPresentFlag, true)
	} else {
		PosSIBsItemOptPresentFlag = append(PosSIBsItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PosSIBsItemOptPresentFlag = append(PosSIBsItemOptPresentFlag, true)
	} else {
		PosSIBsItemOptPresentFlag = append(PosSIBsItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PosSIBsItemOptPresentFlag, true)
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
	err = x.PosSIBSegments.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PosSIBSegments marshal failed")
	}

	// optional field
	if x.AssistanceInformationMetaData != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AssistanceInformationMetaData.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AssistanceInformationMetaData marshal failed")
		}
	}

	// optional field
	if x.BroadcastPriority != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 16
		err = pd.WriteInteger(*(x.BroadcastPriority), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *PosSIBsItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSIBsItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PosSIBsItemOptPresentFlag, true)
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
	x.PosSIBSegments = new(PosSIBSegments)
	err = x.PosSIBSegments.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PosSIBSegments error")
	}

	// optional field (optPresentFlag index: 0)
	if PosSIBsItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AssistanceInformationMetaData = new(AssistanceInformationMetaData)
		err = x.AssistanceInformationMetaData.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AssistanceInformationMetaData error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PosSIBsItemOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 16
		x.BroadcastPriority = new(int64)
		*(x.BroadcastPriority), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if PosSIBsItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPosSIBsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

type PosSIBs struct {
	List []PosSIBsItem // valueExt,sizeLB:1,sizeUB:32
}

func (x *PosSIBs) Write(pd *aper.PerBitData) error {
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

func (x *PosSIBs) Read(pd *aper.PerBitData) error {
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
	x.List = []PosSIBsItem{}
	for i := 0; i < int(numElementsList); i++ {
		var val PosSIBsItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}
