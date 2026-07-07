package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type GNBRxTxTimeDiff struct {
	RxTxTimeDiff       *GNBRxTxTimeDiffMeas                             // valueLB:0,valueUB:6
	AdditionalPathList *AdditionalPathList                              // optional
	IEExtensions       *ProtocolExtensionContainerGNBRxTxTimeDiffExtIEs // optional
}

func (x *GNBRxTxTimeDiff) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GNBRxTxTimeDiffOptPresentFlag := []bool{}
	// mandatory field
	if x.RxTxTimeDiff == nil {
		return errors.Errorf("RxTxTimeDiff is missing")
	}
	// optional field
	if x.AdditionalPathList != nil {
		GNBRxTxTimeDiffOptPresentFlag = append(GNBRxTxTimeDiffOptPresentFlag, true)
	} else {
		GNBRxTxTimeDiffOptPresentFlag = append(GNBRxTxTimeDiffOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		GNBRxTxTimeDiffOptPresentFlag = append(GNBRxTxTimeDiffOptPresentFlag, true)
	} else {
		GNBRxTxTimeDiffOptPresentFlag = append(GNBRxTxTimeDiffOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GNBRxTxTimeDiffOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RxTxTimeDiff.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RxTxTimeDiff marshal failed")
	}

	// optional field
	if x.AdditionalPathList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AdditionalPathList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AdditionalPathList marshal failed")
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

func (x *GNBRxTxTimeDiff) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GNBRxTxTimeDiffOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&GNBRxTxTimeDiffOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RxTxTimeDiff = new(GNBRxTxTimeDiffMeas)
	err = x.RxTxTimeDiff.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RxTxTimeDiff error")
	}

	// optional field (optPresentFlag index: 0)
	if GNBRxTxTimeDiffOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AdditionalPathList = new(AdditionalPathList)
		err = x.AdditionalPathList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AdditionalPathList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if GNBRxTxTimeDiffOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGNBRxTxTimeDiffExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
