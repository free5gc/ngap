package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &ExtendedRANNodeName{}

type ExtendedRANNodeName struct {
	RANNodeNameVisibleString *RANNodeNameVisibleString                            // sizeExt,sizeLB:1,sizeUB:150,optional
	RANNodeNameUTF8String    *RANNodeNameUTF8String                               // sizeExt,sizeLB:1,sizeUB:150,optional
	IEExtensions             *ProtocolExtensionContainerExtendedRANNodeNameExtIEs // optional
}

func (x *ExtendedRANNodeName) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExtendedRANNodeNameOptPresentFlag := []bool{}
	// optional field
	if x.RANNodeNameVisibleString != nil {
		ExtendedRANNodeNameOptPresentFlag = append(ExtendedRANNodeNameOptPresentFlag, true)
	} else {
		ExtendedRANNodeNameOptPresentFlag = append(ExtendedRANNodeNameOptPresentFlag, false)
	}
	// optional field
	if x.RANNodeNameUTF8String != nil {
		ExtendedRANNodeNameOptPresentFlag = append(ExtendedRANNodeNameOptPresentFlag, true)
	} else {
		ExtendedRANNodeNameOptPresentFlag = append(ExtendedRANNodeNameOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExtendedRANNodeNameOptPresentFlag = append(ExtendedRANNodeNameOptPresentFlag, true)
	} else {
		ExtendedRANNodeNameOptPresentFlag = append(ExtendedRANNodeNameOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExtendedRANNodeNameOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.RANNodeNameVisibleString != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANNodeNameVisibleString.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RANNodeNameVisibleString marshal failed")
		}
	}

	// optional field
	if x.RANNodeNameUTF8String != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANNodeNameUTF8String.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RANNodeNameUTF8String marshal failed")
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

func (x *ExtendedRANNodeName) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExtendedRANNodeNameOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&ExtendedRANNodeNameOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if ExtendedRANNodeNameOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.RANNodeNameVisibleString = new(RANNodeNameVisibleString)
		err = x.RANNodeNameVisibleString.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RANNodeNameVisibleString error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExtendedRANNodeNameOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.RANNodeNameUTF8String = new(RANNodeNameUTF8String)
		err = x.RANNodeNameUTF8String.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RANNodeNameUTF8String error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ExtendedRANNodeNameOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExtendedRANNodeNameExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *ExtendedRANNodeName) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *ExtendedRANNodeName) ReadIE(pd *aper.PerBitData) error {
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
