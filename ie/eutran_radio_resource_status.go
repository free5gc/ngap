package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EUTRANRadioResourceStatus struct {
	DLGBRPRBUsage             *int64                                                     // valueLB:0,valueUB:100
	ULGBRPRBUsage             *int64                                                     // valueLB:0,valueUB:100
	DLNonGBRPRBUsage          *int64                                                     // valueLB:0,valueUB:100
	ULNonGBRPRBUsage          *int64                                                     // valueLB:0,valueUB:100
	DLTotalPRBUsage           *int64                                                     // valueLB:0,valueUB:100
	ULTotalPRBUsage           *int64                                                     // valueLB:0,valueUB:100
	DLSchedulingPDCCHCCEUsage *int64                                                     // valueLB:0,valueUB:100,optional
	ULSchedulingPDCCHCCEUsage *int64                                                     // valueLB:0,valueUB:100,optional
	IEExtensions              *ProtocolExtensionContainerEUTRANRadioResourceStatusExtIEs // optional
}

func (x *EUTRANRadioResourceStatus) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EUTRANRadioResourceStatusOptPresentFlag := []bool{}
	// mandatory field
	if x.DLGBRPRBUsage == nil {
		return errors.Errorf("DLGBRPRBUsage is missing")
	}
	// mandatory field
	if x.ULGBRPRBUsage == nil {
		return errors.Errorf("ULGBRPRBUsage is missing")
	}
	// mandatory field
	if x.DLNonGBRPRBUsage == nil {
		return errors.Errorf("DLNonGBRPRBUsage is missing")
	}
	// mandatory field
	if x.ULNonGBRPRBUsage == nil {
		return errors.Errorf("ULNonGBRPRBUsage is missing")
	}
	// mandatory field
	if x.DLTotalPRBUsage == nil {
		return errors.Errorf("DLTotalPRBUsage is missing")
	}
	// mandatory field
	if x.ULTotalPRBUsage == nil {
		return errors.Errorf("ULTotalPRBUsage is missing")
	}
	// optional field
	if x.DLSchedulingPDCCHCCEUsage != nil {
		EUTRANRadioResourceStatusOptPresentFlag = append(EUTRANRadioResourceStatusOptPresentFlag, true)
	} else {
		EUTRANRadioResourceStatusOptPresentFlag = append(EUTRANRadioResourceStatusOptPresentFlag, false)
	}
	// optional field
	if x.ULSchedulingPDCCHCCEUsage != nil {
		EUTRANRadioResourceStatusOptPresentFlag = append(EUTRANRadioResourceStatusOptPresentFlag, true)
	} else {
		EUTRANRadioResourceStatusOptPresentFlag = append(EUTRANRadioResourceStatusOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		EUTRANRadioResourceStatusOptPresentFlag = append(EUTRANRadioResourceStatusOptPresentFlag, true)
	} else {
		EUTRANRadioResourceStatusOptPresentFlag = append(EUTRANRadioResourceStatusOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EUTRANRadioResourceStatusOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.DLGBRPRBUsage), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.ULGBRPRBUsage), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.DLNonGBRPRBUsage), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.ULNonGBRPRBUsage), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.DLTotalPRBUsage), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.ULTotalPRBUsage), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.DLSchedulingPDCCHCCEUsage != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 100
		err = pd.WriteInteger(*(x.DLSchedulingPDCCHCCEUsage), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.ULSchedulingPDCCHCCEUsage != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 100
		err = pd.WriteInteger(*(x.ULSchedulingPDCCHCCEUsage), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *EUTRANRadioResourceStatus) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EUTRANRadioResourceStatusOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&EUTRANRadioResourceStatusOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.DLGBRPRBUsage = new(int64)
	*(x.DLGBRPRBUsage), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.ULGBRPRBUsage = new(int64)
	*(x.ULGBRPRBUsage), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.DLNonGBRPRBUsage = new(int64)
	*(x.DLNonGBRPRBUsage), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.ULNonGBRPRBUsage = new(int64)
	*(x.ULNonGBRPRBUsage), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.DLTotalPRBUsage = new(int64)
	*(x.DLTotalPRBUsage), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.ULTotalPRBUsage = new(int64)
	*(x.ULTotalPRBUsage), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if EUTRANRadioResourceStatusOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 100
		x.DLSchedulingPDCCHCCEUsage = new(int64)
		*(x.DLSchedulingPDCCHCCEUsage), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if EUTRANRadioResourceStatusOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 100
		x.ULSchedulingPDCCHCCEUsage = new(int64)
		*(x.ULSchedulingPDCCHCCEUsage), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if EUTRANRadioResourceStatusOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEUTRANRadioResourceStatusExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
