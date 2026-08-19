package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SSBInfo struct {
	/* Sequence of = 35, FULL Name = struct SSBInfo__listOfSSBInfo */
	/* Type Name = SSBInfoItem */
	/* Sequence Of Embed */
	ListOfSSBInfo []SSBInfoItem                            // valueExt,sizeLB:1,sizeUB:255
	IEExtensions  *ProtocolExtensionContainerSSBInfoExtIEs // optional
}

func (x *SSBInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SSBInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.ListOfSSBInfo == nil {
		return errors.Errorf("ListOfSSBInfo is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SSBInfoOptPresentFlag = append(SSBInfoOptPresentFlag, true)
	} else {
		SSBInfoOptPresentFlag = append(SSBInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SSBInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Sequence Of
	*sLb, *sUb = 1, 255
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.ListOfSSBInfo)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.ListOfSSBInfo {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
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

func (x *SSBInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SSBInfoOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SSBInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 255
	var numElementsListOfSSBInfo uint64
	numElementsListOfSSBInfo, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.ListOfSSBInfo = []SSBInfoItem{}
	for i := 0; i < int(numElementsListOfSSBInfo); i++ {
		var val SSBInfoItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.ListOfSSBInfo = append(x.ListOfSSBInfo, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if SSBInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSSBInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
