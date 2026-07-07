package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPRxTEGInformation struct {
	TRPRxTEGID             *int64                                               // valueLB:0,valueUB:31
	TRPRxTimingErrorMargin *TimingErrorMargin                                   // valueExt,valueLB:0,valueUB:15
	IEExtensions           *ProtocolExtensionContainerTRPRxTEGInformationExtIEs // optional
}

func (x *TRPRxTEGInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPRxTEGInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPRxTEGID == nil {
		return errors.Errorf("TRPRxTEGID is missing")
	}
	// mandatory field
	if x.TRPRxTimingErrorMargin == nil {
		return errors.Errorf("TRPRxTimingErrorMargin is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPRxTEGInformationOptPresentFlag = append(TRPRxTEGInformationOptPresentFlag, true)
	} else {
		TRPRxTEGInformationOptPresentFlag = append(TRPRxTEGInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPRxTEGInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 31
	err = pd.WriteInteger(*(x.TRPRxTEGID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TRPRxTimingErrorMargin.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPRxTimingErrorMargin marshal failed")
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

func (x *TRPRxTEGInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPRxTEGInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPRxTEGInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 31
	x.TRPRxTEGID = new(int64)
	*(x.TRPRxTEGID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPRxTimingErrorMargin = new(TimingErrorMargin)
	err = x.TRPRxTimingErrorMargin.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPRxTimingErrorMargin error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPRxTEGInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPRxTEGInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
