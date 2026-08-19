package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceModifyItemModReq struct {
	PDUSessionID                            *PDUSessionID
	NASPDU                                  *NASPDU // optional
	PDUSessionResourceModifyRequestTransfer *aper.OctetString
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs // optional
}

func (x *PDUSessionResourceModifyItemModReq) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceModifyItemModReqOptPresentFlag := []bool{}
	// mandatory field
	if x.PDUSessionID == nil {
		return errors.Errorf("PDUSessionID is missing")
	}
	// optional field
	if x.NASPDU != nil {
		PDUSessionResourceModifyItemModReqOptPresentFlag = append(PDUSessionResourceModifyItemModReqOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyItemModReqOptPresentFlag = append(PDUSessionResourceModifyItemModReqOptPresentFlag, false)
	}
	// mandatory field
	if x.PDUSessionResourceModifyRequestTransfer == nil {
		return errors.Errorf("PDUSessionResourceModifyRequestTransfer is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceModifyItemModReqOptPresentFlag = append(PDUSessionResourceModifyItemModReqOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyItemModReqOptPresentFlag = append(PDUSessionResourceModifyItemModReqOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceModifyItemModReqOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PDUSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PDUSessionID marshal failed")
	}

	// optional field
	if x.NASPDU != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NASPDU.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NASPDU marshal failed")
		}
	}

	// Write OctetString (Pointer)
	sLb, sUb = nil, nil
	err = pd.WriteOctetString(*(x.PDUSessionResourceModifyRequestTransfer), false, sLb, sUb)
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

func (x *PDUSessionResourceModifyItemModReq) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceModifyItemModReqOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceModifyItemModReqOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceModifyItemModReqOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NASPDU = new(NASPDU)
		err = x.NASPDU.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NASPDU error")
		}
	}

	// mandatory field
	// Read OctetString (Pointer)
	sLb, sUb = nil, nil
	x.PDUSessionResourceModifyRequestTransfer = new(aper.OctetString)
	*(x.PDUSessionResourceModifyRequestTransfer), err = pd.ReadOctetString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode octetstring error"))
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceModifyItemModReqOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceModifyItemModReqExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
