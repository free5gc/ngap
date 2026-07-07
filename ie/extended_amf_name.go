package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &ExtendedAMFName{}

type ExtendedAMFName struct {
	AMFNameVisibleString *AMFNameVisibleString                            // sizeExt,sizeLB:1,sizeUB:150,optional
	AMFNameUTF8String    *AMFNameUTF8String                               // sizeExt,sizeLB:1,sizeUB:150,optional
	IEExtensions         *ProtocolExtensionContainerExtendedAMFNameExtIEs // optional
}

func (x *ExtendedAMFName) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExtendedAMFNameOptPresentFlag := []bool{}
	// optional field
	if x.AMFNameVisibleString != nil {
		ExtendedAMFNameOptPresentFlag = append(ExtendedAMFNameOptPresentFlag, true)
	} else {
		ExtendedAMFNameOptPresentFlag = append(ExtendedAMFNameOptPresentFlag, false)
	}
	// optional field
	if x.AMFNameUTF8String != nil {
		ExtendedAMFNameOptPresentFlag = append(ExtendedAMFNameOptPresentFlag, true)
	} else {
		ExtendedAMFNameOptPresentFlag = append(ExtendedAMFNameOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExtendedAMFNameOptPresentFlag = append(ExtendedAMFNameOptPresentFlag, true)
	} else {
		ExtendedAMFNameOptPresentFlag = append(ExtendedAMFNameOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExtendedAMFNameOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.AMFNameVisibleString != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AMFNameVisibleString.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AMFNameVisibleString marshal failed")
		}
	}

	// optional field
	if x.AMFNameUTF8String != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AMFNameUTF8String.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AMFNameUTF8String marshal failed")
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

func (x *ExtendedAMFName) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExtendedAMFNameOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&ExtendedAMFNameOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if ExtendedAMFNameOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AMFNameVisibleString = new(AMFNameVisibleString)
		err = x.AMFNameVisibleString.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AMFNameVisibleString error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExtendedAMFNameOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.AMFNameUTF8String = new(AMFNameUTF8String)
		err = x.AMFNameUTF8String.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AMFNameUTF8String error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ExtendedAMFNameOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExtendedAMFNameExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *ExtendedAMFName) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *ExtendedAMFName) ReadIE(pd *aper.PerBitData) error {
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
