package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MulticastSessionDeactivationRequestTransfer struct {
	MBSSessionID *MBSSessionID                                                                // valueExt
	IEExtensions *ProtocolExtensionContainerMulticastSessionDeactivationRequestTransferExtIEs // optional
}

func (x *MulticastSessionDeactivationRequestTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MulticastSessionDeactivationRequestTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSSessionID == nil {
		return errors.Errorf("MBSSessionID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		MulticastSessionDeactivationRequestTransferOptPresentFlag = append(MulticastSessionDeactivationRequestTransferOptPresentFlag, true)
	} else {
		MulticastSessionDeactivationRequestTransferOptPresentFlag = append(MulticastSessionDeactivationRequestTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MulticastSessionDeactivationRequestTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSSessionID marshal failed")
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

func (x *MulticastSessionDeactivationRequestTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MulticastSessionDeactivationRequestTransferOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&MulticastSessionDeactivationRequestTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSSessionID = new(MBSSessionID)
	err = x.MBSSessionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSSessionID error")
	}

	// optional field (optPresentFlag index: 0)
	if MulticastSessionDeactivationRequestTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMulticastSessionDeactivationRequestTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
