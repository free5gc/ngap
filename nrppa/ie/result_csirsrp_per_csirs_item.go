package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultCSIRSRPPerCSIRSItem struct {
	CSIRSIndex   *int64 // valueLB:0,valueUB:95
	ValueCSIRSRP *ValueRSRPNR
	IEExtensions *ProtocolExtensionContainerResultCSIRSRPPerCSIRSItemExtIEs // optional
}

func (x *ResultCSIRSRPPerCSIRSItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRPPerCSIRSItemOptPresentFlag := []bool{}
	// mandatory field
	if x.CSIRSIndex == nil {
		return errors.Errorf("CSIRSIndex is missing")
	}
	// mandatory field
	if x.ValueCSIRSRP == nil {
		return errors.Errorf("ValueCSIRSRP is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultCSIRSRPPerCSIRSItemOptPresentFlag = append(ResultCSIRSRPPerCSIRSItemOptPresentFlag, true)
	} else {
		ResultCSIRSRPPerCSIRSItemOptPresentFlag = append(ResultCSIRSRPPerCSIRSItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRPPerCSIRSItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 95
	err = pd.WriteInteger(*(x.CSIRSIndex), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ValueCSIRSRP.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ValueCSIRSRP marshal failed")
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

func (x *ResultCSIRSRPPerCSIRSItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRPPerCSIRSItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRPPerCSIRSItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 95
	x.CSIRSIndex = new(int64)
	*(x.CSIRSIndex), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ValueCSIRSRP = new(ValueRSRPNR)
	err = x.ValueCSIRSRP.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ValueCSIRSRP error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultCSIRSRPPerCSIRSItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultCSIRSRPPerCSIRSItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
