package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type COUNTValueForPDCPSN12 struct {
	PDCPSN12     *int64                                                 // valueLB:0,valueUB:4095
	HFNPDCPSN12  *int64                                                 // valueLB:0,valueUB:1048575
	IEExtensions *ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs // optional
}

func (x *COUNTValueForPDCPSN12) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	COUNTValueForPDCPSN12OptPresentFlag := []bool{}
	// mandatory field
	if x.PDCPSN12 == nil {
		return errors.Errorf("PDCPSN12 is missing")
	}
	// mandatory field
	if x.HFNPDCPSN12 == nil {
		return errors.Errorf("HFNPDCPSN12 is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		COUNTValueForPDCPSN12OptPresentFlag = append(COUNTValueForPDCPSN12OptPresentFlag, true)
	} else {
		COUNTValueForPDCPSN12OptPresentFlag = append(COUNTValueForPDCPSN12OptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(COUNTValueForPDCPSN12OptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 4095
	err = pd.WriteInteger(*(x.PDCPSN12), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 1048575
	err = pd.WriteInteger(*(x.HFNPDCPSN12), false, vLb, vUb)
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

func (x *COUNTValueForPDCPSN12) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	COUNTValueForPDCPSN12OptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&COUNTValueForPDCPSN12OptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 4095
	x.PDCPSN12 = new(int64)
	*(x.PDCPSN12), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 1048575
	x.HFNPDCPSN12 = new(int64)
	*(x.HFNPDCPSN12), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if COUNTValueForPDCPSN12OptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCOUNTValueForPDCPSN12ExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
