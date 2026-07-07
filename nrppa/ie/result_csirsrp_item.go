package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultCSIRSRPItem struct {
	NRPCI            *NRPCI
	NRARFCN          *NRARFCN
	CGINR            *CGINR                                             // valueExt,optional
	ValueCSIRSRPCell *ValueRSRPNR                                       // optional
	CSIRSRPPerCSIRS  *ResultCSIRSRPPerCSIRS                             // optional
	IEExtensions     *ProtocolExtensionContainerResultCSIRSRPItemExtIEs // optional
}

func (x *ResultCSIRSRPItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRPItemOptPresentFlag := []bool{}
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
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, true)
	} else {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, false)
	}
	// optional field
	if x.ValueCSIRSRPCell != nil {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, true)
	} else {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, false)
	}
	// optional field
	if x.CSIRSRPPerCSIRS != nil {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, true)
	} else {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, true)
	} else {
		ResultCSIRSRPItemOptPresentFlag = append(ResultCSIRSRPItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRPItemOptPresentFlag, true)
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
	if x.ValueCSIRSRPCell != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ValueCSIRSRPCell.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ValueCSIRSRPCell marshal failed")
		}
	}

	// optional field
	if x.CSIRSRPPerCSIRS != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CSIRSRPPerCSIRS.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CSIRSRPPerCSIRS marshal failed")
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

func (x *ResultCSIRSRPItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRPItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRPItemOptPresentFlag, true)
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
	if ResultCSIRSRPItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CGINR = new(CGINR)
		err = x.CGINR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGINR error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ResultCSIRSRPItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ValueCSIRSRPCell = new(ValueRSRPNR)
		err = x.ValueCSIRSRPCell.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ValueCSIRSRPCell error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ResultCSIRSRPItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.CSIRSRPPerCSIRS = new(ResultCSIRSRPPerCSIRS)
		err = x.CSIRSRPPerCSIRS.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CSIRSRPPerCSIRS error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ResultCSIRSRPItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultCSIRSRPItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
