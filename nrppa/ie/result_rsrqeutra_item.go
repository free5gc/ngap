package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultRSRQEUTRAItem struct {
	PCIEUTRA       *PCIEUTRA
	EARFCN         *EARFCN
	CGIUTRA        *CGIEUTRA // valueExt,optional
	ValueRSRQEUTRA *ValueRSRQEUTRA
	IEExtensions   *ProtocolExtensionContainerResultRSRQEUTRAItemExtIEs // optional
}

func (x *ResultRSRQEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultRSRQEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PCIEUTRA == nil {
		return errors.Errorf("PCIEUTRA is missing")
	}
	// mandatory field
	if x.EARFCN == nil {
		return errors.Errorf("EARFCN is missing")
	}
	// optional field
	if x.CGIUTRA != nil {
		ResultRSRQEUTRAItemOptPresentFlag = append(ResultRSRQEUTRAItemOptPresentFlag, true)
	} else {
		ResultRSRQEUTRAItemOptPresentFlag = append(ResultRSRQEUTRAItemOptPresentFlag, false)
	}
	// mandatory field
	if x.ValueRSRQEUTRA == nil {
		return errors.Errorf("ValueRSRQEUTRA is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultRSRQEUTRAItemOptPresentFlag = append(ResultRSRQEUTRAItemOptPresentFlag, true)
	} else {
		ResultRSRQEUTRAItemOptPresentFlag = append(ResultRSRQEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultRSRQEUTRAItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PCIEUTRA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PCIEUTRA marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.EARFCN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EARFCN marshal failed")
	}

	// optional field
	if x.CGIUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CGIUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CGIUTRA marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ValueRSRQEUTRA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ValueRSRQEUTRA marshal failed")
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

func (x *ResultRSRQEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultRSRQEUTRAItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ResultRSRQEUTRAItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PCIEUTRA = new(PCIEUTRA)
	err = x.PCIEUTRA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PCIEUTRA error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EARFCN = new(EARFCN)
	err = x.EARFCN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EARFCN error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultRSRQEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CGIUTRA = new(CGIEUTRA)
		err = x.CGIUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGIUTRA error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ValueRSRQEUTRA = new(ValueRSRQEUTRA)
	err = x.ValueRSRQEUTRA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ValueRSRQEUTRA error")
	}

	// optional field (optPresentFlag index: 1)
	if ResultRSRQEUTRAItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultRSRQEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
