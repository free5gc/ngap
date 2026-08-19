package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type COUNTValueForPDCPSN18 struct {
	PDCPSN18     *int64                                                 // valueLB:0,valueUB:262143
	HFNPDCPSN18  *int64                                                 // valueLB:0,valueUB:16383
	IEExtensions *ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs // optional
}

func (x *COUNTValueForPDCPSN18) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	COUNTValueForPDCPSN18OptPresentFlag := []bool{}
	// mandatory field
	if x.PDCPSN18 == nil {
		return errors.Errorf("PDCPSN18 is missing")
	}
	// mandatory field
	if x.HFNPDCPSN18 == nil {
		return errors.Errorf("HFNPDCPSN18 is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		COUNTValueForPDCPSN18OptPresentFlag = append(COUNTValueForPDCPSN18OptPresentFlag, true)
	} else {
		COUNTValueForPDCPSN18OptPresentFlag = append(COUNTValueForPDCPSN18OptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(COUNTValueForPDCPSN18OptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 262143
	err = pd.WriteInteger(*(x.PDCPSN18), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 16383
	err = pd.WriteInteger(*(x.HFNPDCPSN18), false, vLb, vUb)
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

func (x *COUNTValueForPDCPSN18) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	COUNTValueForPDCPSN18OptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&COUNTValueForPDCPSN18OptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 262143
	x.PDCPSN18 = new(int64)
	*(x.PDCPSN18), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 16383
	x.HFNPDCPSN18 = new(int64)
	*(x.HFNPDCPSN18), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if COUNTValueForPDCPSN18OptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCOUNTValueForPDCPSN18ExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
