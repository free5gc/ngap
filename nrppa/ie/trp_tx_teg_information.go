package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPTxTEGInformation struct {
	TRPTxTEGID             *int64                                               // valueLB:0,valueUB:7
	TRPTxTimingErrorMargin *TimingErrorMargin                                   // valueExt,valueLB:0,valueUB:15
	IEExtensions           *ProtocolExtensionContainerTRPTxTEGInformationExtIEs // optional
}

func (x *TRPTxTEGInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPTxTEGInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPTxTEGID == nil {
		return errors.Errorf("TRPTxTEGID is missing")
	}
	// mandatory field
	if x.TRPTxTimingErrorMargin == nil {
		return errors.Errorf("TRPTxTimingErrorMargin is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPTxTEGInformationOptPresentFlag = append(TRPTxTEGInformationOptPresentFlag, true)
	} else {
		TRPTxTEGInformationOptPresentFlag = append(TRPTxTEGInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPTxTEGInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 7
	err = pd.WriteInteger(*(x.TRPTxTEGID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TRPTxTimingErrorMargin.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPTxTimingErrorMargin marshal failed")
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

func (x *TRPTxTEGInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPTxTEGInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPTxTEGInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 7
	x.TRPTxTEGID = new(int64)
	*(x.TRPTxTEGID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPTxTimingErrorMargin = new(TimingErrorMargin)
	err = x.TRPTxTimingErrorMargin.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPTxTimingErrorMargin error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPTxTEGInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPTxTEGInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
