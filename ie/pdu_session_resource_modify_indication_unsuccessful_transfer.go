package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceModifyIndicationUnsuccessfulTransfer struct {
	Cause        *Cause                                                                                  // valueLB:0,valueUB:5
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs // optional
}

func (x *PDUSessionResourceModifyIndicationUnsuccessfulTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.Cause == nil {
		return errors.Errorf("Cause is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag = append(PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag = append(PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Cause.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Cause marshal failed")
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

func (x *PDUSessionResourceModifyIndicationUnsuccessfulTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Cause = new(Cause)
	err = x.Cause.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Cause error")
	}

	// optional field (optPresentFlag index: 0)
	if PDUSessionResourceModifyIndicationUnsuccessfulTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceModifyIndicationUnsuccessfulTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
