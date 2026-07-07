package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &ULAoA{}

type ULAoA struct {
	AzimuthAoA          *int64                                 // valueLB:0,valueUB:3599
	ZenithAoA           *int64                                 // valueLB:0,valueUB:1799,optional
	LCSToGCSTranslation *LCSToGCSTranslation                   // valueExt,optional
	IEExtensions        *ProtocolExtensionContainerULAoAExtIEs // optional
}

func (x *ULAoA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ULAoAOptPresentFlag := []bool{}
	// mandatory field
	if x.AzimuthAoA == nil {
		return errors.Errorf("AzimuthAoA is missing")
	}
	// optional field
	if x.ZenithAoA != nil {
		ULAoAOptPresentFlag = append(ULAoAOptPresentFlag, true)
	} else {
		ULAoAOptPresentFlag = append(ULAoAOptPresentFlag, false)
	}
	// optional field
	if x.LCSToGCSTranslation != nil {
		ULAoAOptPresentFlag = append(ULAoAOptPresentFlag, true)
	} else {
		ULAoAOptPresentFlag = append(ULAoAOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ULAoAOptPresentFlag = append(ULAoAOptPresentFlag, true)
	} else {
		ULAoAOptPresentFlag = append(ULAoAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ULAoAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3599
	err = pd.WriteInteger(*(x.AzimuthAoA), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.ZenithAoA != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 1799
		err = pd.WriteInteger(*(x.ZenithAoA), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.LCSToGCSTranslation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LCSToGCSTranslation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "LCSToGCSTranslation marshal failed")
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

func (x *ULAoA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ULAoAOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&ULAoAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3599
	x.AzimuthAoA = new(int64)
	*(x.AzimuthAoA), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if ULAoAOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 1799
		x.ZenithAoA = new(int64)
		*(x.ZenithAoA), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if ULAoAOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.LCSToGCSTranslation = new(LCSToGCSTranslation)
		err = x.LCSToGCSTranslation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode LCSToGCSTranslation error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ULAoAOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerULAoAExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *ULAoA) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.WriteSequencePreambleBitMap(optPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: id
	err = id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: criticality
	err = criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: value (open type)
	pdOpenType := aper.NewPerBitData(nil)
	err = x.Write(pdOpenType)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}

	return nil
}

func (x *ULAoA) ReadIE(pd *aper.PerBitData) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.ReadSequencePreambleBitMap(&optPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err = protocolIECriticality.Read(pd)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	// sequence element: value (open type)
	var pdOpenTypeBytes []byte
	pdOpenTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	pdOpenType := aper.NewPerBitData(pdOpenTypeBytes)
	err = x.Read(pdOpenType)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	return nil
}
