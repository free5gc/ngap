package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type RxTxTEG struct {
	TRPRxTxTEGInformation *TRPRxTxTEGInformation                   // valueExt
	TRPTxTEGInformation   *TRPTxTEGInformation                     // valueExt,optional
	IEExtensions          *ProtocolExtensionContainerRxTxTEGExtIEs // optional
}

func (x *RxTxTEG) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RxTxTEGOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPRxTxTEGInformation == nil {
		return errors.Errorf("TRPRxTxTEGInformation is missing")
	}
	// optional field
	if x.TRPTxTEGInformation != nil {
		RxTxTEGOptPresentFlag = append(RxTxTEGOptPresentFlag, true)
	} else {
		RxTxTEGOptPresentFlag = append(RxTxTEGOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		RxTxTEGOptPresentFlag = append(RxTxTEGOptPresentFlag, true)
	} else {
		RxTxTEGOptPresentFlag = append(RxTxTEGOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RxTxTEGOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPRxTxTEGInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPRxTxTEGInformation marshal failed")
	}

	// optional field
	if x.TRPTxTEGInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TRPTxTEGInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TRPTxTEGInformation marshal failed")
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

func (x *RxTxTEG) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RxTxTEGOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&RxTxTEGOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPRxTxTEGInformation = new(TRPRxTxTEGInformation)
	err = x.TRPRxTxTEGInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPRxTxTEGInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if RxTxTEGOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TRPTxTEGInformation = new(TRPTxTEGInformation)
		err = x.TRPTxTEGInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TRPTxTEGInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if RxTxTEGOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRxTxTEGExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
