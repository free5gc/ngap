package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type RxTEG struct {
	TRPRxTEGInformation *TRPRxTEGInformation                   // valueExt
	TRPTxTEGInformation *TRPTxTEGInformation                   // valueExt
	IEExtensions        *ProtocolExtensionContainerRxTEGExtIEs // optional
}

func (x *RxTEG) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RxTEGOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPRxTEGInformation == nil {
		return errors.Errorf("TRPRxTEGInformation is missing")
	}
	// mandatory field
	if x.TRPTxTEGInformation == nil {
		return errors.Errorf("TRPTxTEGInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RxTEGOptPresentFlag = append(RxTEGOptPresentFlag, true)
	} else {
		RxTEGOptPresentFlag = append(RxTEGOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RxTEGOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPRxTEGInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPRxTEGInformation marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TRPTxTEGInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPTxTEGInformation marshal failed")
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

func (x *RxTEG) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RxTEGOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RxTEGOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPRxTEGInformation = new(TRPRxTEGInformation)
	err = x.TRPRxTEGInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPRxTEGInformation error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPTxTEGInformation = new(TRPTxTEGInformation)
	err = x.TRPTxTEGInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPTxTEGInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if RxTEGOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRxTEGExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
