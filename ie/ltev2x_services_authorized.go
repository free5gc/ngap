package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &LTEV2XServicesAuthorized{}

type LTEV2XServicesAuthorized struct {
	VehicleUE    *VehicleUE                                                // valueExt,valueLB:0,valueUB:1,optional
	PedestrianUE *PedestrianUE                                             // valueExt,valueLB:0,valueUB:1,optional
	IEExtensions *ProtocolExtensionContainerLTEV2XServicesAuthorizedExtIEs // optional
}

func (x *LTEV2XServicesAuthorized) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LTEV2XServicesAuthorizedOptPresentFlag := []bool{}
	// optional field
	if x.VehicleUE != nil {
		LTEV2XServicesAuthorizedOptPresentFlag = append(LTEV2XServicesAuthorizedOptPresentFlag, true)
	} else {
		LTEV2XServicesAuthorizedOptPresentFlag = append(LTEV2XServicesAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.PedestrianUE != nil {
		LTEV2XServicesAuthorizedOptPresentFlag = append(LTEV2XServicesAuthorizedOptPresentFlag, true)
	} else {
		LTEV2XServicesAuthorizedOptPresentFlag = append(LTEV2XServicesAuthorizedOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		LTEV2XServicesAuthorizedOptPresentFlag = append(LTEV2XServicesAuthorizedOptPresentFlag, true)
	} else {
		LTEV2XServicesAuthorizedOptPresentFlag = append(LTEV2XServicesAuthorizedOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LTEV2XServicesAuthorizedOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.VehicleUE != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.VehicleUE.Write(pd)
		if err != nil {
			return errors.Wrap(err, "VehicleUE marshal failed")
		}
	}

	// optional field
	if x.PedestrianUE != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PedestrianUE.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PedestrianUE marshal failed")
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

func (x *LTEV2XServicesAuthorized) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LTEV2XServicesAuthorizedOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&LTEV2XServicesAuthorizedOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if LTEV2XServicesAuthorizedOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.VehicleUE = new(VehicleUE)
		err = x.VehicleUE.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode VehicleUE error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if LTEV2XServicesAuthorizedOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PedestrianUE = new(PedestrianUE)
		err = x.PedestrianUE.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PedestrianUE error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if LTEV2XServicesAuthorizedOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLTEV2XServicesAuthorizedExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *LTEV2XServicesAuthorized) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *LTEV2XServicesAuthorized) ReadIE(pd *aper.PerBitData) error {
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
