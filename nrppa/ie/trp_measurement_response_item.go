package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPMeasurementResponseItem struct {
	TRPID             *TRPID
	MeasurementResult *TrpMeasurementResult
	IEExtensions      *ProtocolExtensionContainerTRPMeasurementResponseItemExtIEs // optional
}

func (x *TRPMeasurementResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementResponseItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// mandatory field
	if x.MeasurementResult == nil {
		return errors.Errorf("MeasurementResult is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPMeasurementResponseItemOptPresentFlag = append(TRPMeasurementResponseItemOptPresentFlag, true)
	} else {
		TRPMeasurementResponseItemOptPresentFlag = append(TRPMeasurementResponseItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementResponseItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MeasurementResult.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MeasurementResult marshal failed")
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

func (x *TRPMeasurementResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementResponseItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementResponseItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPID = new(TRPID)
	err = x.TRPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MeasurementResult = new(TrpMeasurementResult)
	err = x.MeasurementResult.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MeasurementResult error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPMeasurementResponseItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPMeasurementResponseItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
