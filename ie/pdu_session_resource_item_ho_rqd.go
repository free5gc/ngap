package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceItemHORqd struct {
	PDUSessionID             *PDUSessionID
	HandoverRequiredTransfer *aper.OctetString
	IEExtensions             *ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs // optional
}

func (x *PDUSessionResourceItemHORqd) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceItemHORqdOptPresentFlag := []bool{}
	// mandatory field
	if x.PDUSessionID == nil {
		return errors.Errorf("PDUSessionID is missing")
	}
	// mandatory field
	if x.HandoverRequiredTransfer == nil {
		return errors.Errorf("HandoverRequiredTransfer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceItemHORqdOptPresentFlag = append(PDUSessionResourceItemHORqdOptPresentFlag, true)
	} else {
		PDUSessionResourceItemHORqdOptPresentFlag = append(PDUSessionResourceItemHORqdOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceItemHORqdOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PDUSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PDUSessionID marshal failed")
	}

	// Write OctetString (Pointer)
	sLb, sUb = nil, nil
	err = pd.WriteOctetString(*(x.HandoverRequiredTransfer), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "octetString marshal failed")
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

func (x *PDUSessionResourceItemHORqd) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceItemHORqdOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceItemHORqdOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PDUSessionID = new(PDUSessionID)
	err = x.PDUSessionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PDUSessionID error")
	}

	// mandatory field
	// Read OctetString (Pointer)
	sLb, sUb = nil, nil
	x.HandoverRequiredTransfer = new(aper.OctetString)
	*(x.HandoverRequiredTransfer), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceItemHORqdOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceItemHORqdExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
