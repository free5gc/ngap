package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceItemCxtRelCpl struct {
	PDUSessionID *PDUSessionID
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs // optional
}

func (x *PDUSessionResourceItemCxtRelCpl) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceItemCxtRelCplOptPresentFlag := []bool{}
	// mandatory field
	if x.PDUSessionID == nil {
		return errors.Errorf("PDUSessionID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceItemCxtRelCplOptPresentFlag = append(PDUSessionResourceItemCxtRelCplOptPresentFlag, true)
	} else {
		PDUSessionResourceItemCxtRelCplOptPresentFlag = append(PDUSessionResourceItemCxtRelCplOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceItemCxtRelCplOptPresentFlag, true)
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
	if x.IEExtensions != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtensions.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtensions marshal failed")
		}
	}

	return nil
}

func (x *PDUSessionResourceItemCxtRelCpl) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceItemCxtRelCplOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceItemCxtRelCplOptPresentFlag, true)
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
	if PDUSessionResourceItemCxtRelCplOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceItemCxtRelCplExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
