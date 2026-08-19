package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceSetupItemSURes struct {
	PDUSessionID                            *PDUSessionID
	PDUSessionResourceSetupResponseTransfer *aper.OctetString
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs // optional
}

func (x *PDUSessionResourceSetupItemSURes) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceSetupItemSUResOptPresentFlag := []bool{}
	// mandatory field
	if x.PDUSessionID == nil {
		return errors.Errorf("PDUSessionID is missing")
	}
	// mandatory field
	if x.PDUSessionResourceSetupResponseTransfer == nil {
		return errors.Errorf("PDUSessionResourceSetupResponseTransfer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceSetupItemSUResOptPresentFlag = append(PDUSessionResourceSetupItemSUResOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupItemSUResOptPresentFlag = append(PDUSessionResourceSetupItemSUResOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceSetupItemSUResOptPresentFlag, true)
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
	err = pd.WriteOctetString(*(x.PDUSessionResourceSetupResponseTransfer), false, sLb, sUb)
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

func (x *PDUSessionResourceSetupItemSURes) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceSetupItemSUResOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceSetupItemSUResOptPresentFlag, true)
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
	x.PDUSessionResourceSetupResponseTransfer = new(aper.OctetString)
	*(x.PDUSessionResourceSetupResponseTransfer), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceSetupItemSUResOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceSetupItemSUResExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
