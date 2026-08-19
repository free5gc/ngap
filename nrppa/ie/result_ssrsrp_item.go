package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultSSRSRPItem struct {
	NRPCI           *NRPCI
	NRARFCN         *NRARFCN
	CGINR           *CGINR                                            // valueExt,optional
	ValueSSRSRPCell *ValueRSRPNR                                      // optional
	SSRSRPPerSSB    *ResultSSRSRPPerSSB                               // optional
	IEExtensions    *ProtocolExtensionContainerResultSSRSRPItemExtIEs // optional
}

func (x *ResultSSRSRPItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRPItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NRPCI == nil {
		return errors.Errorf("NRPCI is missing")
	}
	// mandatory field
	if x.NRARFCN == nil {
		return errors.Errorf("NRARFCN is missing")
	}
	// optional field
	if x.CGINR != nil {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, true)
	} else {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, false)
	}
	// optional field
	if x.ValueSSRSRPCell != nil {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, true)
	} else {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, false)
	}
	// optional field
	if x.SSRSRPPerSSB != nil {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, true)
	} else {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, true)
	} else {
		ResultSSRSRPItemOptPresentFlag = append(ResultSSRSRPItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRPItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRPCI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRPCI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NRARFCN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRARFCN marshal failed")
	}

	// optional field
	if x.CGINR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CGINR.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CGINR marshal failed")
		}
	}

	// optional field
	if x.ValueSSRSRPCell != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueSSRSRPCell.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueSSRSRPCell marshal failed")
		}
	}

	// optional field
	if x.SSRSRPPerSSB != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SSRSRPPerSSB.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SSRSRPPerSSB marshal failed")
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

func (x *ResultSSRSRPItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRPItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRPItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRPCI = new(NRPCI)
	err = x.NRPCI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRPCI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRARFCN = new(NRARFCN)
	err = x.NRARFCN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRARFCN error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultSSRSRPItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CGINR = new(CGINR)
		err = x.CGINR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGINR error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ResultSSRSRPItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ValueSSRSRPCell = new(ValueRSRPNR)
		err = x.ValueSSRSRPCell.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueSSRSRPCell error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ResultSSRSRPItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SSRSRPPerSSB = new(ResultSSRSRPPerSSB)
		err = x.SSRSRPPerSSB.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSRSRPPerSSB error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ResultSSRSRPItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultSSRSRPItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
