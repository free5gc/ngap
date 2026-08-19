package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ResultSSRSRQPerSSBItem struct {
	SSBIndex     *SSBIndex
	ValueSSRSRQ  *ValueRSRQNR
	IEExtensions *ProtocolExtensionContainerResultSSRSRQPerSSBItemExtIEs // optional
}

func (x *ResultSSRSRQPerSSBItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ResultSSRSRQPerSSBItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SSBIndex == nil {
		return errors.Errorf("SSBIndex is missing")
	}
	// mandatory field
	if x.ValueSSRSRQ == nil {
		return errors.Errorf("ValueSSRSRQ is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ResultSSRSRQPerSSBItemOptPresentFlag = append(ResultSSRSRQPerSSBItemOptPresentFlag, true)
	} else {
		ResultSSRSRQPerSSBItemOptPresentFlag = append(ResultSSRSRQPerSSBItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ResultSSRSRQPerSSBItemOptPresentFlag, true)
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
	err = x.ValueSSRSRQ.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ValueSSRSRQ marshal failed")
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

func (x *ResultSSRSRQPerSSBItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ResultSSRSRQPerSSBItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ResultSSRSRQPerSSBItemOptPresentFlag, true)
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
	x.ValueSSRSRQ = new(ValueRSRQNR)
	err = x.ValueSSRSRQ.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ValueSSRSRQ error")
	}

	// optional field (optPresentFlag index: 0)
	if ResultSSRSRQPerSSBItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerResultSSRSRQPerSSBItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
