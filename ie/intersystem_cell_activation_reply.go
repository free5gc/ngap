package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type IntersystemCellActivationReply struct {
	ActivatedCellList *ActivatedCellList
	ActivationID      *int64                                                          // valueExt,valueLB:0,valueUB:16384
	IEExtensions      *ProtocolExtensionContainerIntersystemCellActivationReplyExtIEs // optional
}

func (x *IntersystemCellActivationReply) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemCellActivationReplyOptPresentFlag := []bool{}
	// mandatory field
	if x.ActivatedCellList == nil {
		return errors.Errorf("ActivatedCellList is missing")
	}
	// mandatory field
	if x.ActivationID == nil {
		return errors.Errorf("ActivationID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemCellActivationReplyOptPresentFlag = append(IntersystemCellActivationReplyOptPresentFlag, true)
	} else {
		IntersystemCellActivationReplyOptPresentFlag = append(IntersystemCellActivationReplyOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemCellActivationReplyOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ActivatedCellList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ActivatedCellList marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 16384
	err = pd.WriteInteger(*(x.ActivationID), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *IntersystemCellActivationReply) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemCellActivationReplyOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemCellActivationReplyOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ActivatedCellList = new(ActivatedCellList)
	err = x.ActivatedCellList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ActivatedCellList error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 16384
	x.ActivationID = new(int64)
	*(x.ActivationID), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemCellActivationReplyOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemCellActivationReplyExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
