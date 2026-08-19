package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPBeamPowerItem struct {
	PRSResourceSetID  *PRSResourceSetID // optional
	PRSResourceID     *PRSResourceID
	RelativePower     *int64                                            // valueLB:0,valueUB:30
	RelativePowerFine *int64                                            // valueLB:0,valueUB:9,optional
	IEExtensions      *ProtocolExtensionContainerTRPBeamPowerItemExtIEs // optional
}

func (x *TRPBeamPowerItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamPowerItemOptPresentFlag := []bool{}
	// optional field
	if x.PRSResourceSetID != nil {
		TRPBeamPowerItemOptPresentFlag = append(TRPBeamPowerItemOptPresentFlag, true)
	} else {
		TRPBeamPowerItemOptPresentFlag = append(TRPBeamPowerItemOptPresentFlag, false)
	}
	// mandatory field
	if x.PRSResourceID == nil {
		return errors.Errorf("PRSResourceID is missing")
	}
	// mandatory field
	if x.RelativePower == nil {
		return errors.Errorf("RelativePower is missing")
	}
	// optional field
	if x.RelativePowerFine != nil {
		TRPBeamPowerItemOptPresentFlag = append(TRPBeamPowerItemOptPresentFlag, true)
	} else {
		TRPBeamPowerItemOptPresentFlag = append(TRPBeamPowerItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TRPBeamPowerItemOptPresentFlag = append(TRPBeamPowerItemOptPresentFlag, true)
	} else {
		TRPBeamPowerItemOptPresentFlag = append(TRPBeamPowerItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamPowerItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PRSResourceSetID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSResourceSetID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSResourceSetID marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceID marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 30
	err = pd.WriteInteger(*(x.RelativePower), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.RelativePowerFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.RelativePowerFine), false, vLb, vUb)
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

func (x *TRPBeamPowerItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamPowerItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamPowerItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if TRPBeamPowerItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PRSResourceSetID = new(PRSResourceSetID)
		err = x.PRSResourceSetID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSResourceSetID error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceID = new(PRSResourceID)
	err = x.PRSResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceID error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 30
	x.RelativePower = new(int64)
	*(x.RelativePower), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 1)
	if TRPBeamPowerItemOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.RelativePowerFine = new(int64)
		*(x.RelativePowerFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if TRPBeamPowerItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPBeamPowerItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
