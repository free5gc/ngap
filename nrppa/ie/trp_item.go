package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPItem struct {
	TRPID        *TRPID
	IEExtensions *ProtocolExtensionContainerTRPItemExtIEs // optional
}

func (x *TRPItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPItemOptPresentFlag = append(TRPItemOptPresentFlag, true)
	} else {
		TRPItemOptPresentFlag = append(TRPItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPID marshal failed")
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

func (x *TRPItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPID = new(TRPID)
	err = x.TRPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPID error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
