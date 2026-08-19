package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceSetupRequestTransfer struct {
	ProtocolIEs *ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs
}

func (x *PDUSessionResourceSetupRequestTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceSetupRequestTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.ProtocolIEs == nil {
		return errors.Errorf("ProtocolIEs is missing")
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceSetupRequestTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ProtocolIEs.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ProtocolIEs marshal failed")
	}

	return nil
}

func (x *PDUSessionResourceSetupRequestTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceSetupRequestTransferOptPresentFlag := make([]bool, 0)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceSetupRequestTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ProtocolIEs = new(ProtocolIEContainerPDUSessionResourceSetupRequestTransferIEs)
	err = x.ProtocolIEs.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ProtocolIEs error")
	}

	return nil
}
