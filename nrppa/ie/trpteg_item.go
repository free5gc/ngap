package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPTEGItem struct {
	TRPTxTEGInformation *TRPTxTEGInformation // valueExt
	DlPRSResourceSetID  *PRSResourceSetID
	/* Sequence of = 35, FULL Name = struct TRPTEGItem__dl_PRSResourceID_List */
	/* Type Name = DLPRSResourceIDItem */
	/* Sequence Of Embed */
	DlPRSResourceIDList []DLPRSResourceIDItem                       // valueExt,sizeLB:1,sizeUB:64
	IEExtensions        *ProtocolExtensionContainerTRPTEGItemExtIEs // optional
}

func (x *TRPTEGItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPTEGItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPTxTEGInformation == nil {
		return errors.Errorf("TRPTxTEGInformation is missing")
	}
	// mandatory field
	if x.DlPRSResourceSetID == nil {
		return errors.Errorf("DlPRSResourceSetID is missing")
	}
	// mandatory field
	if x.DlPRSResourceIDList == nil {
		return errors.Errorf("DlPRSResourceIDList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPTEGItemOptPresentFlag = append(TRPTEGItemOptPresentFlag, true)
	} else {
		TRPTEGItemOptPresentFlag = append(TRPTEGItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPTEGItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPTxTEGInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPTxTEGInformation marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DlPRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DlPRSResourceSetID marshal failed")
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 64
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.DlPRSResourceIDList)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.DlPRSResourceIDList {
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

func (x *TRPTEGItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPTEGItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPTEGItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPTxTEGInformation = new(TRPTxTEGInformation)
	err = x.TRPTxTEGInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPTxTEGInformation error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DlPRSResourceSetID = new(PRSResourceSetID)
	err = x.DlPRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DlPRSResourceSetID error")
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 64
	var numElementsDlPRSResourceIDList uint64
	numElementsDlPRSResourceIDList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.DlPRSResourceIDList = []DLPRSResourceIDItem{}
	for i := 0; i < int(numElementsDlPRSResourceIDList); i++ {
		var val DLPRSResourceIDItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.DlPRSResourceIDList = append(x.DlPRSResourceIDList, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if TRPTEGItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPTEGItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
