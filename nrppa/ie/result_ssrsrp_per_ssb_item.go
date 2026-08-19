package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultSSRSRPPerSSBItem struct {
	SSBIndex     *SSBIndex
	ValueSSRSRP  *ValueRSRPNR
	IEExtensions *ProtocolExtensionContainerResultSSRSRPPerSSBItemExtIEs // optional
}

func (x *ResultSSRSRPPerSSBItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRPPerSSBItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SSBIndex == nil {
		return errors.Errorf("SSBIndex is missing")
	}
	// mandatory field
	if x.ValueSSRSRP == nil {
		return errors.Errorf("ValueSSRSRP is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultSSRSRPPerSSBItemOptPresentFlag = append(ResultSSRSRPPerSSBItemOptPresentFlag, true)
	} else {
		ResultSSRSRPPerSSBItemOptPresentFlag = append(ResultSSRSRPPerSSBItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRPPerSSBItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SSBIndex.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SSBIndex marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ValueSSRSRP.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ValueSSRSRP marshal failed")
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

func (x *ResultSSRSRPPerSSBItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRPPerSSBItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRPPerSSBItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SSBIndex = new(SSBIndex)
	err = x.SSBIndex.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SSBIndex error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ValueSSRSRP = new(ValueRSRPNR)
	err = x.ValueSSRSRP.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ValueSSRSRP error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultSSRSRPPerSSBItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultSSRSRPPerSSBItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
