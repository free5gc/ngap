package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultNRItem struct {
	NRPCI           *NRPCI
	NRARFCN         *NRARFCN
	ValueSSRSRPCell *ValueRSRPNR                                  // optional
	ValueSSRSRQCell *ValueRSRQNR                                  // optional
	SSRSRPPerSSB    *ResultSSRSRPPerSSB                           // optional
	SSRSRQPerSSB    *ResultSSRSRQPerSSB                           // optional
	CGINR           *CGINR                                        // valueExt,optional
	IEExtensions    *ProtocolExtensionContainerResultNRItemExtIEs // optional
}

func (x *ResultNRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultNRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NRPCI == nil {
		return errors.Errorf("NRPCI is missing")
	}
	// mandatory field
	if x.NRARFCN == nil {
		return errors.Errorf("NRARFCN is missing")
	}
	// optional field
	if x.ValueSSRSRPCell != nil {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, true)
	} else {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, false)
	}
	// optional field
	if x.ValueSSRSRQCell != nil {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, true)
	} else {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, false)
	}
	// optional field
	if x.SSRSRPPerSSB != nil {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, true)
	} else {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, false)
	}
	// optional field
	if x.SSRSRQPerSSB != nil {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, true)
	} else {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, false)
	}
	// optional field
	if x.CGINR != nil {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, true)
	} else {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, true)
	} else {
		ResultNRItemOptPresentFlag = append(ResultNRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultNRItemOptPresentFlag, true)
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
	if x.ValueSSRSRPCell != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueSSRSRPCell.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueSSRSRPCell marshal failed")
		}
	}

	// optional field
	if x.ValueSSRSRQCell != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueSSRSRQCell.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueSSRSRQCell marshal failed")
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
	if x.SSRSRQPerSSB != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SSRSRQPerSSB.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SSRSRQPerSSB marshal failed")
		}
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
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *ResultNRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultNRItemOptPresentFlag := make([]bool, 6)
	err = pd.ReadSequencePreambleBitMap(&ResultNRItemOptPresentFlag, true)
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
	if ResultNRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ValueSSRSRPCell = new(ValueRSRPNR)
		err = x.ValueSSRSRPCell.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueSSRSRPCell error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ResultNRItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ValueSSRSRQCell = new(ValueRSRQNR)
		err = x.ValueSSRSRQCell.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueSSRSRQCell error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ResultNRItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SSRSRPPerSSB = new(ResultSSRSRPPerSSB)
		err = x.SSRSRPPerSSB.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSRSRPPerSSB error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ResultNRItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.SSRSRQPerSSB = new(ResultSSRSRQPerSSB)
		err = x.SSRSRQPerSSB.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSRSRQPerSSB error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if ResultNRItemOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.CGINR = new(CGINR)
		err = x.CGINR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGINR error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if ResultNRItemOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultNRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
