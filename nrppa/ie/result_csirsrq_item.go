package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultCSIRSRQItem struct {
	NRPCI            *NRPCI
	NRARFCN          *NRARFCN
	CGINR            *CGINR                                             // valueExt,optional
	ValueCSIRSRQCell *ValueRSRQNR                                       // optional
	CSIRSRQPerCSIRS  *ResultCSIRSRQPerCSIRS                             // optional
	IEExtensions     *ProtocolExtensionContainerResultCSIRSRQItemExtIEs // optional
}

func (x *ResultCSIRSRQItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRQItemOptPresentFlag := []bool{}
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
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, true)
	} else {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, false)
	}
	// optional field
	if x.ValueCSIRSRQCell != nil {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, true)
	} else {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, false)
	}
	// optional field
	if x.CSIRSRQPerCSIRS != nil {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, true)
	} else {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, true)
	} else {
		ResultCSIRSRQItemOptPresentFlag = append(ResultCSIRSRQItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRQItemOptPresentFlag, true)
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
	if x.ValueCSIRSRQCell != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueCSIRSRQCell.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueCSIRSRQCell marshal failed")
		}
	}

	// optional field
	if x.CSIRSRQPerCSIRS != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CSIRSRQPerCSIRS.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CSIRSRQPerCSIRS marshal failed")
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

func (x *ResultCSIRSRQItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRQItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRQItemOptPresentFlag, true)
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
	if ResultCSIRSRQItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CGINR = new(CGINR)
		err = x.CGINR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGINR error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ResultCSIRSRQItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ValueCSIRSRQCell = new(ValueRSRQNR)
		err = x.ValueCSIRSRQCell.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueCSIRSRQCell error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ResultCSIRSRQItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.CSIRSRQPerCSIRS = new(ResultCSIRSRQPerCSIRS)
		err = x.CSIRSRQPerCSIRS.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CSIRSRQPerCSIRS error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ResultCSIRSRQItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultCSIRSRQItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
