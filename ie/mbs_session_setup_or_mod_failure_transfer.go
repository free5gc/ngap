package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSSessionSetupOrModFailureTransfer struct {
	Cause                  *Cause                                                               // valueLB:0,valueUB:5
	CriticalityDiagnostics *CriticalityDiagnostics                                              // valueExt,optional
	IEExtensions           *ProtocolExtensionContainerMBSSessionSetupOrModFailureTransferExtIEs // optional
}

func (x *MBSSessionSetupOrModFailureTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSSessionSetupOrModFailureTransferOptPresentFlag := []bool{}
	// mandatory field
	if x.Cause == nil {
		return errors.Errorf("Cause is missing")
	}
	// optional field
	if x.CriticalityDiagnostics != nil {
		MBSSessionSetupOrModFailureTransferOptPresentFlag = append(MBSSessionSetupOrModFailureTransferOptPresentFlag, true)
	} else {
		MBSSessionSetupOrModFailureTransferOptPresentFlag = append(MBSSessionSetupOrModFailureTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSSessionSetupOrModFailureTransferOptPresentFlag = append(MBSSessionSetupOrModFailureTransferOptPresentFlag, true)
	} else {
		MBSSessionSetupOrModFailureTransferOptPresentFlag = append(MBSSessionSetupOrModFailureTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSSessionSetupOrModFailureTransferOptPresentFlag, true)
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

func (x *MBSSessionSetupOrModFailureTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSSessionSetupOrModFailureTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MBSSessionSetupOrModFailureTransferOptPresentFlag, true)
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
	if MBSSessionSetupOrModFailureTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CriticalityDiagnostics = new(CriticalityDiagnostics)
		err = x.CriticalityDiagnostics.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CriticalityDiagnostics error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSSessionSetupOrModFailureTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSSessionSetupOrModFailureTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
