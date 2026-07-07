package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSAngleItem struct {
	NRPRSAzimuth       *int64                                        // valueLB:0,valueUB:359
	NRPRSAzimuthFine   *int64                                        // valueLB:0,valueUB:9,optional
	NRPRSElevation     *int64                                        // valueLB:0,valueUB:180,optional
	NRPRSElevationFine *int64                                        // valueLB:0,valueUB:9,optional
	IEExtensions       *ProtocolExtensionContainerPRSAngleItemExtIEs // optional
}

func (x *PRSAngleItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSAngleItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NRPRSAzimuth == nil {
		return errors.Errorf("NRPRSAzimuth is missing")
	}
	// optional field
	if x.NRPRSAzimuthFine != nil {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, true)
	} else {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, false)
	}
	// optional field
	if x.NRPRSElevation != nil {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, true)
	} else {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, false)
	}
	// optional field
	if x.NRPRSElevationFine != nil {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, true)
	} else {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, true)
	} else {
		PRSAngleItemOptPresentFlag = append(PRSAngleItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSAngleItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 359
	err = pd.WriteInteger(*(x.NRPRSAzimuth), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.NRPRSAzimuthFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.NRPRSAzimuthFine), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.NRPRSElevation != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 180
		err = pd.WriteInteger(*(x.NRPRSElevation), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.NRPRSElevationFine != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 9
		err = pd.WriteInteger(*(x.NRPRSElevationFine), false, vLb, vUb)
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

func (x *PRSAngleItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSAngleItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&PRSAngleItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 359
	x.NRPRSAzimuth = new(int64)
	*(x.NRPRSAzimuth), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if PRSAngleItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.NRPRSAzimuthFine = new(int64)
		*(x.NRPRSAzimuthFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if PRSAngleItemOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 180
		x.NRPRSElevation = new(int64)
		*(x.NRPRSElevation), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if PRSAngleItemOptPresentFlag[2] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 9
		x.NRPRSElevationFine = new(int64)
		*(x.NRPRSElevationFine), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if PRSAngleItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSAngleItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
