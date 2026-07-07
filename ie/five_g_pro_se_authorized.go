package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &FiveGProSeAuthorized{}

type FiveGProSeAuthorized struct {
	FiveGProSeDirectDiscovery        *FiveGProSeDirectDiscovery                            // valueExt,valueLB:0,valueUB:1,optional
	FiveGProSeDirectCommunication    *FiveGProSeDirectCommunication                        // valueExt,valueLB:0,valueUB:1,optional
	FiveGProSeLayer2UEtoNetworkRelay *FiveGProSeLayer2UEtoNetworkRelay                     // valueExt,valueLB:0,valueUB:1,optional
	FiveGProSeLayer3UEtoNetworkRelay *FiveGProSeLayer3UEtoNetworkRelay                     // valueExt,valueLB:0,valueUB:1,optional
	FiveGProSeLayer2RemoteUE         *FiveGProSeLayer2RemoteUE                             // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions                     *ProtocolExtensionContainerFiveGProSeAuthorizedExtIEs // optional
}

func (x *FiveGProSeAuthorized) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FiveGProSeAuthorizedOptPresentFlag := []bool{}
	// optional field
	if x.FiveGProSeDirectDiscovery != nil {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, true)
	} else {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.FiveGProSeDirectCommunication != nil {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, true)
	} else {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.FiveGProSeLayer2UEtoNetworkRelay != nil {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, true)
	} else {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.FiveGProSeLayer3UEtoNetworkRelay != nil {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, true)
	} else {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.FiveGProSeLayer2RemoteUE != nil {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, true)
	} else {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, true)
	} else {
		FiveGProSeAuthorizedOptPresentFlag = append(FiveGProSeAuthorizedOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FiveGProSeAuthorizedOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.FiveGProSeDirectDiscovery != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGProSeDirectDiscovery.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGProSeDirectDiscovery marshal failed")
		}
	}

	// optional field
	if x.FiveGProSeDirectCommunication != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGProSeDirectCommunication.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGProSeDirectCommunication marshal failed")
		}
	}

	// optional field
	if x.FiveGProSeLayer2UEtoNetworkRelay != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGProSeLayer2UEtoNetworkRelay.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGProSeLayer2UEtoNetworkRelay marshal failed")
		}
	}

	// optional field
	if x.FiveGProSeLayer3UEtoNetworkRelay != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGProSeLayer3UEtoNetworkRelay.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGProSeLayer3UEtoNetworkRelay marshal failed")
		}
	}

	// optional field
	if x.FiveGProSeLayer2RemoteUE != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGProSeLayer2RemoteUE.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGProSeLayer2RemoteUE marshal failed")
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

func (x *FiveGProSeAuthorized) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FiveGProSeAuthorizedOptPresentFlag := make([]bool, 6)
	err = pd.ReadSequencePreambleBitMap(&FiveGProSeAuthorizedOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if FiveGProSeAuthorizedOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGProSeDirectDiscovery = new(FiveGProSeDirectDiscovery)
		err = x.FiveGProSeDirectDiscovery.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGProSeDirectDiscovery error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if FiveGProSeAuthorizedOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGProSeDirectCommunication = new(FiveGProSeDirectCommunication)
		err = x.FiveGProSeDirectCommunication.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGProSeDirectCommunication error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if FiveGProSeAuthorizedOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGProSeLayer2UEtoNetworkRelay = new(FiveGProSeLayer2UEtoNetworkRelay)
		err = x.FiveGProSeLayer2UEtoNetworkRelay.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGProSeLayer2UEtoNetworkRelay error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if FiveGProSeAuthorizedOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGProSeLayer3UEtoNetworkRelay = new(FiveGProSeLayer3UEtoNetworkRelay)
		err = x.FiveGProSeLayer3UEtoNetworkRelay.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGProSeLayer3UEtoNetworkRelay error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if FiveGProSeAuthorizedOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGProSeLayer2RemoteUE = new(FiveGProSeLayer2RemoteUE)
		err = x.FiveGProSeLayer2RemoteUE.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGProSeLayer2RemoteUE error")
		}
	}

	// optional field (optPresentFlag index: 5)
	if FiveGProSeAuthorizedOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFiveGProSeAuthorizedExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *FiveGProSeAuthorized) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *FiveGProSeAuthorized) ReadIE(pd *aper.PerBitData) error {
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
