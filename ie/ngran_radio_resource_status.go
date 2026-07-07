package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGRANRadioResourceStatus struct {
	DLGBRPRBUsageForMIMO    *int64                                                    // valueLB:0,valueUB:100
	ULGBRPRBUsageForMIMO    *int64                                                    // valueLB:0,valueUB:100
	DLNonGBRPRBUsageForMIMO *int64                                                    // valueLB:0,valueUB:100
	ULNonGBRPRBUsageForMIMO *int64                                                    // valueLB:0,valueUB:100
	DLTotalPRBUsageForMIMO  *int64                                                    // valueLB:0,valueUB:100
	ULTotalPRBUsageForMIMO  *int64                                                    // valueLB:0,valueUB:100
	IEExtensions            *ProtocolExtensionContainerNGRANRadioResourceStatusExtIEs // optional
}

func (x *NGRANRadioResourceStatus) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANRadioResourceStatusOptPresentFlag := []bool{}
	// mandatory field
	if x.DLGBRPRBUsageForMIMO == nil {
		return errors.Errorf("DLGBRPRBUsageForMIMO is missing")
	}
	// mandatory field
	if x.ULGBRPRBUsageForMIMO == nil {
		return errors.Errorf("ULGBRPRBUsageForMIMO is missing")
	}
	// mandatory field
	if x.DLNonGBRPRBUsageForMIMO == nil {
		return errors.Errorf("DLNonGBRPRBUsageForMIMO is missing")
	}
	// mandatory field
	if x.ULNonGBRPRBUsageForMIMO == nil {
		return errors.Errorf("ULNonGBRPRBUsageForMIMO is missing")
	}
	// mandatory field
	if x.DLTotalPRBUsageForMIMO == nil {
		return errors.Errorf("DLTotalPRBUsageForMIMO is missing")
	}
	// mandatory field
	if x.ULTotalPRBUsageForMIMO == nil {
		return errors.Errorf("ULTotalPRBUsageForMIMO is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANRadioResourceStatusOptPresentFlag = append(NGRANRadioResourceStatusOptPresentFlag, true)
	} else {
		NGRANRadioResourceStatusOptPresentFlag = append(NGRANRadioResourceStatusOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANRadioResourceStatusOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.DLGBRPRBUsageForMIMO), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.ULGBRPRBUsageForMIMO), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.DLNonGBRPRBUsageForMIMO), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.ULNonGBRPRBUsageForMIMO), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.DLTotalPRBUsageForMIMO), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 100
	err = pd.WriteInteger(*(x.ULTotalPRBUsageForMIMO), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *NGRANRadioResourceStatus) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANRadioResourceStatusOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGRANRadioResourceStatusOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.DLGBRPRBUsageForMIMO = new(int64)
	*(x.DLGBRPRBUsageForMIMO), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.ULGBRPRBUsageForMIMO = new(int64)
	*(x.ULGBRPRBUsageForMIMO), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.DLNonGBRPRBUsageForMIMO = new(int64)
	*(x.DLNonGBRPRBUsageForMIMO), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.ULNonGBRPRBUsageForMIMO = new(int64)
	*(x.ULNonGBRPRBUsageForMIMO), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.DLTotalPRBUsageForMIMO = new(int64)
	*(x.DLTotalPRBUsageForMIMO), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 100
	x.ULTotalPRBUsageForMIMO = new(int64)
	*(x.ULTotalPRBUsageForMIMO), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if NGRANRadioResourceStatusOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANRadioResourceStatusExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
