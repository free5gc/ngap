package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPMeasurementRequestItem struct {
	TRPID                   *TRPID
	SearchWindowInformation *SearchWindowInformation                                   // valueExt,optional
	IEExtensions            *ProtocolExtensionContainerTRPMeasurementRequestItemExtIEs // optional
}

func (x *TRPMeasurementRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPMeasurementRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// optional field
	if x.SearchWindowInformation != nil {
		TRPMeasurementRequestItemOptPresentFlag = append(TRPMeasurementRequestItemOptPresentFlag, true)
	} else {
		TRPMeasurementRequestItemOptPresentFlag = append(TRPMeasurementRequestItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TRPMeasurementRequestItemOptPresentFlag = append(TRPMeasurementRequestItemOptPresentFlag, true)
	} else {
		TRPMeasurementRequestItemOptPresentFlag = append(TRPMeasurementRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPMeasurementRequestItemOptPresentFlag, true)
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
	if x.SearchWindowInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SearchWindowInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SearchWindowInformation marshal failed")
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

func (x *TRPMeasurementRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPMeasurementRequestItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPMeasurementRequestItemOptPresentFlag, true)
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
	if TRPMeasurementRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SearchWindowInformation = new(SearchWindowInformation)
		err = x.SearchWindowInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SearchWindowInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TRPMeasurementRequestItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPMeasurementRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
