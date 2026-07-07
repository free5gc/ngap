package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &NRPagingeDRXInformation{}

type NRPagingeDRXInformation struct {
	NRPagingEDRXCycle  *NRPagingEDRXCycle                                       // valueExt,valueLB:0,valueUB:12
	NRPagingTimeWindow *NRPagingTimeWindow                                      // valueExt,valueLB:0,valueUB:15,optional
	IEExtensions       *ProtocolExtensionContainerNRPagingeDRXInformationExtIEs // optional
}

func (x *NRPagingeDRXInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRPagingeDRXInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.NRPagingEDRXCycle == nil {
		return errors.Errorf("NRPagingEDRXCycle is missing")
	}
	// optional field
	if x.NRPagingTimeWindow != nil {
		NRPagingeDRXInformationOptPresentFlag = append(NRPagingeDRXInformationOptPresentFlag, true)
	} else {
		NRPagingeDRXInformationOptPresentFlag = append(NRPagingeDRXInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		NRPagingeDRXInformationOptPresentFlag = append(NRPagingeDRXInformationOptPresentFlag, true)
	} else {
		NRPagingeDRXInformationOptPresentFlag = append(NRPagingeDRXInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRPagingeDRXInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRPagingEDRXCycle.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRPagingEDRXCycle marshal failed")
	}

	// optional field
	if x.NRPagingTimeWindow != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NRPagingTimeWindow.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NRPagingTimeWindow marshal failed")
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

func (x *NRPagingeDRXInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRPagingeDRXInformationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&NRPagingeDRXInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRPagingEDRXCycle = new(NRPagingEDRXCycle)
	err = x.NRPagingEDRXCycle.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRPagingEDRXCycle error")
	}

	// optional field (optPresentFlag index: 0)
	if NRPagingeDRXInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NRPagingTimeWindow = new(NRPagingTimeWindow)
		err = x.NRPagingTimeWindow.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NRPagingTimeWindow error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if NRPagingeDRXInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNRPagingeDRXInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *NRPagingeDRXInformation) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *NRPagingeDRXInformation) ReadIE(pd *aper.PerBitData) error {
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
