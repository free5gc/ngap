package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultRSRPEUTRAItem struct {
	PCIEUTRA       *PCIEUTRA
	EARFCN         *EARFCN
	CGIEUTRA       *CGIEUTRA // valueExt,optional
	ValueRSRPEUTRA *ValueRSRPEUTRA
	IEExtensions   *ProtocolExtensionContainerResultRSRPEUTRAItemExtIEs // optional
}

func (x *ResultRSRPEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultRSRPEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PCIEUTRA == nil {
		return errors.Errorf("PCIEUTRA is missing")
	}
	// mandatory field
	if x.EARFCN == nil {
		return errors.Errorf("EARFCN is missing")
	}
	// optional field
	if x.CGIEUTRA != nil {
		ResultRSRPEUTRAItemOptPresentFlag = append(ResultRSRPEUTRAItemOptPresentFlag, true)
	} else {
		ResultRSRPEUTRAItemOptPresentFlag = append(ResultRSRPEUTRAItemOptPresentFlag, false)
	}
	// mandatory field
	if x.ValueRSRPEUTRA == nil {
		return errors.Errorf("ValueRSRPEUTRA is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultRSRPEUTRAItemOptPresentFlag = append(ResultRSRPEUTRAItemOptPresentFlag, true)
	} else {
		ResultRSRPEUTRAItemOptPresentFlag = append(ResultRSRPEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultRSRPEUTRAItemOptPresentFlag, true)
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
	if x.CGIEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CGIEUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CGIEUTRA marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ValueRSRPEUTRA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ValueRSRPEUTRA marshal failed")
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

func (x *ResultRSRPEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultRSRPEUTRAItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ResultRSRPEUTRAItemOptPresentFlag, true)
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
	if ResultRSRPEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CGIEUTRA = new(CGIEUTRA)
		err = x.CGIEUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGIEUTRA error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ValueRSRPEUTRA = new(ValueRSRPEUTRA)
	err = x.ValueRSRPEUTRA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ValueRSRPEUTRA error")
	}

	// optional field (optPresentFlag index: 1)
	if ResultRSRPEUTRAItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultRSRPEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
