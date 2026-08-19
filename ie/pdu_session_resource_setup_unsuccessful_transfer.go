package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  *Cause                                                                       // valueLB:0,valueUB:5
	CriticalityDiagnostics *CriticalityDiagnostics                                                      // valueExt,optional
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs // optional
}

func (x *PDUSessionResourceSetupUnsuccessfulTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.Cause == nil {
		return errors.Errorf("Cause is missing")
	}
	// optional field
	if x.CriticalityDiagnostics != nil {
		PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag = append(PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag = append(PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag = append(PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag, true)
	} else {
		PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag = append(PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag, true)
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
	if x.CriticalityDiagnostics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CriticalityDiagnostics.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CriticalityDiagnostics marshal failed")
		}
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

func (x *PDUSessionResourceSetupUnsuccessfulTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag, true)
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
	if PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PDUSessionResourceSetupUnsuccessfulTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
