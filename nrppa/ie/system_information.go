package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SystemInformationItem struct {
	BroadcastPeriodicity *BroadcastPeriodicity // valueExt,valueLB:0,valueUB:6
	PosSIBs              *PosSIBs
	IEExtensions         *ProtocolExtensionContainerSystemInformationExtIEs // optional
}

func (x *SystemInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SystemInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.BroadcastPeriodicity == nil {
		return errors.Errorf("BroadcastPeriodicity is missing")
	}
	// mandatory field
	if x.PosSIBs == nil {
		return errors.Errorf("PosSIBs is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SystemInformationItemOptPresentFlag = append(SystemInformationItemOptPresentFlag, true)
	} else {
		SystemInformationItemOptPresentFlag = append(SystemInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SystemInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.BroadcastPeriodicity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "BroadcastPeriodicity marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PosSIBs.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PosSIBs marshal failed")
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

func (x *SystemInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SystemInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SystemInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.BroadcastPeriodicity = new(BroadcastPeriodicity)
	err = x.BroadcastPeriodicity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode BroadcastPeriodicity error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PosSIBs = new(PosSIBs)
	err = x.PosSIBs.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PosSIBs error")
	}

	// optional field (optPresentFlag index: 0)
	if SystemInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSystemInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

type SystemInformation struct {
	List []SystemInformationItem // valueExt,sizeLB:1,sizeUB:32
}

func (x *SystemInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Sequence Of
	*sLb, *sUb = 1, 32
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.List)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.List {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	return nil
}

func (x *SystemInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Sequence Of
	*sLb, *sUb = 1, 32
	var numElementsList uint64
	numElementsList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.List = []SystemInformationItem{}
	for i := 0; i < int(numElementsList); i++ {
		var val SystemInformationItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.List = append(x.List, val)
		}
	}

	return nil
}
