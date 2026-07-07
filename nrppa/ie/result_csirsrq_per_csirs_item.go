package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultCSIRSRQPerCSIRSItem struct {
	CSIRSIndex   *int64 // valueLB:0,valueUB:95
	ValueCSIRSRQ *ValueRSRQNR
	IEExtensions *ProtocolExtensionContainerResultCSIRSRQPerCSIRSItemExtIEs // optional
}

func (x *ResultCSIRSRQPerCSIRSItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultCSIRSRQPerCSIRSItemOptPresentFlag := []bool{}
	// mandatory field
	if x.CSIRSIndex == nil {
		return errors.Errorf("CSIRSIndex is missing")
	}
	// mandatory field
	if x.ValueCSIRSRQ == nil {
		return errors.Errorf("ValueCSIRSRQ is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultCSIRSRQPerCSIRSItemOptPresentFlag = append(ResultCSIRSRQPerCSIRSItemOptPresentFlag, true)
	} else {
		ResultCSIRSRQPerCSIRSItemOptPresentFlag = append(ResultCSIRSRQPerCSIRSItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultCSIRSRQPerCSIRSItemOptPresentFlag, true)
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
	err = x.ValueCSIRSRQ.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ValueCSIRSRQ marshal failed")
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

func (x *ResultCSIRSRQPerCSIRSItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultCSIRSRQPerCSIRSItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResultCSIRSRQPerCSIRSItemOptPresentFlag, true)
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
	x.ValueCSIRSRQ = new(ValueRSRQNR)
	err = x.ValueCSIRSRQ.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ValueCSIRSRQ error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultCSIRSRQPerCSIRSItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultCSIRSRQPerCSIRSItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
