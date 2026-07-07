package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSTransmissionOffPerResourceItem struct {
	PRSResourceSetID *PRSResourceSetID
	/* Sequence of = 35, FULL Name = struct PRSTransmissionOffPerResource_Item__pRSTransmissionOffIndicationPerResourceList */
	/* Type Name = PRSTransmissionOffIndicationPerResourceItem */
	/* Sequence Of Embed */
	PRSTransmissionOffIndicationPerResourceList []PRSTransmissionOffIndicationPerResourceItem                      // valueExt,sizeLB:1,sizeUB:64
	IEExtensions                                *ProtocolExtensionContainerPRSTransmissionOffPerResourceItemExtIEs // optional
}

func (x *PRSTransmissionOffPerResourceItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffPerResourceItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSResourceSetID == nil {
		return errors.Errorf("PRSResourceSetID is missing")
	}
	// mandatory field
	if x.PRSTransmissionOffIndicationPerResourceList == nil {
		return errors.Errorf("PRSTransmissionOffIndicationPerResourceList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSTransmissionOffPerResourceItemOptPresentFlag = append(PRSTransmissionOffPerResourceItemOptPresentFlag, true)
	} else {
		PRSTransmissionOffPerResourceItemOptPresentFlag = append(PRSTransmissionOffPerResourceItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffPerResourceItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceSetID marshal failed")
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 64
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.PRSTransmissionOffIndicationPerResourceList)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.PRSTransmissionOffIndicationPerResourceList {
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

func (x *PRSTransmissionOffPerResourceItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffPerResourceItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffPerResourceItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceSetID = new(PRSResourceSetID)
	err = x.PRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceSetID error")
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 64
	var numElementsPRSTransmissionOffIndicationPerResourceList uint64
	numElementsPRSTransmissionOffIndicationPerResourceList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.PRSTransmissionOffIndicationPerResourceList = []PRSTransmissionOffIndicationPerResourceItem{}
	for i := 0; i < int(numElementsPRSTransmissionOffIndicationPerResourceList); i++ {
		var val PRSTransmissionOffIndicationPerResourceItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.PRSTransmissionOffIndicationPerResourceList = append(x.PRSTransmissionOffIndicationPerResourceList, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if PRSTransmissionOffPerResourceItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSTransmissionOffPerResourceItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
