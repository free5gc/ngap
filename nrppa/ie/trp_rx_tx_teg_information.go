package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPRxTxTEGInformation struct {
	TRPRxTxTEGID             *int64                                                 // valueLB:0,valueUB:255
	TRPRxTxTimingErrorMargin *RxTxTimingErrorMargin                                 // valueExt,valueLB:0,valueUB:15
	IEExtensions             *ProtocolExtensionContainerTRPRxTxTEGInformationExtIEs // optional
}

func (x *TRPRxTxTEGInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPRxTxTEGInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPRxTxTEGID == nil {
		return errors.Errorf("TRPRxTxTEGID is missing")
	}
	// mandatory field
	if x.TRPRxTxTimingErrorMargin == nil {
		return errors.Errorf("TRPRxTxTimingErrorMargin is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPRxTxTEGInformationOptPresentFlag = append(TRPRxTxTEGInformationOptPresentFlag, true)
	} else {
		TRPRxTxTEGInformationOptPresentFlag = append(TRPRxTxTEGInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPRxTxTEGInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.TRPRxTxTEGID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TRPRxTxTimingErrorMargin.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPRxTxTimingErrorMargin marshal failed")
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

func (x *TRPRxTxTEGInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPRxTxTEGInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPRxTxTEGInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.TRPRxTxTEGID = new(int64)
	*(x.TRPRxTxTEGID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPRxTxTimingErrorMargin = new(RxTxTimingErrorMargin)
	err = x.TRPRxTxTimingErrorMargin.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPRxTxTimingErrorMargin error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPRxTxTEGInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPRxTxTEGInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
