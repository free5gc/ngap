package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &UserLocationInformationTWIF{}

type UserLocationInformationTWIF struct {
	TWAPID       *TWAPID
	IPAddress    *TransportLayerAddress
	PortNumber   *PortNumber                                                  // optional
	IEExtensions *ProtocolExtensionContainerUserLocationInformationTWIFExtIEs // optional
}

func (x *UserLocationInformationTWIF) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UserLocationInformationTWIFOptPresentFlag := []bool{}
	// mandatory field
	if x.TWAPID == nil {
		return errors.Errorf("TWAPID is missing")
	}
	// mandatory field
	if x.IPAddress == nil {
		return errors.Errorf("IPAddress is missing")
	}
	// optional field
	if x.PortNumber != nil {
		UserLocationInformationTWIFOptPresentFlag = append(UserLocationInformationTWIFOptPresentFlag, true)
	} else {
		UserLocationInformationTWIFOptPresentFlag = append(UserLocationInformationTWIFOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UserLocationInformationTWIFOptPresentFlag = append(UserLocationInformationTWIFOptPresentFlag, true)
	} else {
		UserLocationInformationTWIFOptPresentFlag = append(UserLocationInformationTWIFOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UserLocationInformationTWIFOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TWAPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TWAPID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.IPAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IPAddress marshal failed")
	}

	// optional field
	if x.PortNumber != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PortNumber.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PortNumber marshal failed")
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

func (x *UserLocationInformationTWIF) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UserLocationInformationTWIFOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UserLocationInformationTWIFOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TWAPID = new(TWAPID)
	err = x.TWAPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TWAPID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.IPAddress = new(TransportLayerAddress)
	err = x.IPAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode IPAddress error")
	}

	// optional field (optPresentFlag index: 0)
	if UserLocationInformationTWIFOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PortNumber = new(PortNumber)
		err = x.PortNumber.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PortNumber error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UserLocationInformationTWIFOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUserLocationInformationTWIFExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *UserLocationInformationTWIF) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *UserLocationInformationTWIF) ReadIE(pd *aper.PerBitData) error {
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
