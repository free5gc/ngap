package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultEUTRAItem struct {
	PCIEUTRA       *PCIEUTRA
	EARFCN         *EARFCN
	ValueRSRPEUTRA *ValueRSRPEUTRA                                  // optional
	ValueRSRQEUTRA *ValueRSRQEUTRA                                  // optional
	CGIEUTRA       *CGIEUTRA                                        // valueExt,optional
	IEExtensions   *ProtocolExtensionContainerResultEUTRAItemExtIEs // optional
}

func (x *ResultEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PCIEUTRA == nil {
		return errors.Errorf("PCIEUTRA is missing")
	}
	// mandatory field
	if x.EARFCN == nil {
		return errors.Errorf("EARFCN is missing")
	}
	// optional field
	if x.ValueRSRPEUTRA != nil {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, true)
	} else {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, false)
	}
	// optional field
	if x.ValueRSRQEUTRA != nil {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, true)
	} else {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, false)
	}
	// optional field
	if x.CGIEUTRA != nil {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, true)
	} else {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, true)
	} else {
		ResultEUTRAItemOptPresentFlag = append(ResultEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultEUTRAItemOptPresentFlag, true)
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
	if x.ValueRSRPEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueRSRPEUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueRSRPEUTRA marshal failed")
		}
	}

	// optional field
	if x.ValueRSRQEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueRSRQEUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueRSRQEUTRA marshal failed")
		}
	}

	// optional field
	if x.CGIEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CGIEUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CGIEUTRA marshal failed")
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

func (x *ResultEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultEUTRAItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ResultEUTRAItemOptPresentFlag, true)
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
	if ResultEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ValueRSRPEUTRA = new(ValueRSRPEUTRA)
		err = x.ValueRSRPEUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueRSRPEUTRA error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ResultEUTRAItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ValueRSRQEUTRA = new(ValueRSRQEUTRA)
		err = x.ValueRSRQEUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueRSRQEUTRA error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ResultEUTRAItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.CGIEUTRA = new(CGIEUTRA)
		err = x.CGIEUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGIEUTRA error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ResultEUTRAItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
