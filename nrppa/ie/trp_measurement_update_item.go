package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPMeasurementUpdateItem struct {
	TRPID                *TRPID
	AoAWindowInformation *AoAAssistanceInfo                                        // valueExt,optional
	IEExtensions         *ProtocolExtensionContainerTRPMeasurementUpdateItemExtIEs // optional
}

func (x *TRPMeasurementUpdateItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementUpdateItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// optional field
	if x.AoAWindowInformation != nil {
		TRPMeasurementUpdateItemOptPresentFlag = append(TRPMeasurementUpdateItemOptPresentFlag, true)
	} else {
		TRPMeasurementUpdateItemOptPresentFlag = append(TRPMeasurementUpdateItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TRPMeasurementUpdateItemOptPresentFlag = append(TRPMeasurementUpdateItemOptPresentFlag, true)
	} else {
		TRPMeasurementUpdateItemOptPresentFlag = append(TRPMeasurementUpdateItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementUpdateItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPID marshal failed")
	}

	// optional field
	if x.AoAWindowInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AoAWindowInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AoAWindowInformation marshal failed")
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

func (x *TRPMeasurementUpdateItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementUpdateItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementUpdateItemOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if TRPMeasurementUpdateItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AoAWindowInformation = new(AoAAssistanceInfo)
		err = x.AoAWindowInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AoAWindowInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TRPMeasurementUpdateItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPMeasurementUpdateItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
