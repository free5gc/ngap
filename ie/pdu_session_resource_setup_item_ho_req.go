package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceSetupItemHOReq struct {
	PDUSessionID            *PDUSessionID
	SNSSAI                  *SNSSAI // valueExt
	HandoverRequestTransfer *aper.OctetString
	IEExtensions            *ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs // optional
}

func (x *PDUSessionResourceSetupItemHOReq) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceSetupItemHOReqOptPresentFlag := []bool{}
	// mandatory field
	if x.PDUSessionID == nil {
		return errors.Errorf("PDUSessionID is missing")
	}
	// mandatory field
	if x.SNSSAI == nil {
		return errors.Errorf("SNSSAI is missing")
	}
	// mandatory field
	if x.HandoverRequestTransfer == nil {
		return errors.Errorf("HandoverRequestTransfer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceSetupItemHOReqOptPresentFlag = append(PDUSessionResourceSetupItemHOReqOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupItemHOReqOptPresentFlag = append(PDUSessionResourceSetupItemHOReqOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceSetupItemHOReqOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PDUSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PDUSessionID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SNSSAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SNSSAI marshal failed")
	}

	// Write OctetString (Pointer)
	sLb, sUb = nil, nil
	err = pd.WriteOctetString(*(x.HandoverRequestTransfer), false, sLb, sUb)
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

func (x *PDUSessionResourceSetupItemHOReq) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceSetupItemHOReqOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceSetupItemHOReqOptPresentFlag, true)
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
	// Read struct defined elsewhere (Pointer)
	x.SNSSAI = new(SNSSAI)
	err = x.SNSSAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SNSSAI error")
	}

	// mandatory field
	// Read OctetString (Pointer)
	sLb, sUb = nil, nil
	x.HandoverRequestTransfer = new(aper.OctetString)
	*(x.HandoverRequestTransfer), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceSetupItemHOReqOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceSetupItemHOReqExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
