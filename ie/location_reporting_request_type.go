package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &LocationReportingRequestType{}

type LocationReportingRequestType struct {
	EventType                                 *EventType                                                    // valueExt,valueLB:0,valueUB:5
	ReportArea                                *ReportArea                                                   // valueExt,valueLB:0,valueUB:0
	AreaOfInterestList                        *AreaOfInterestList                                           // optional
	LocationReportingReferenceIDToBeCancelled *LocationReportingReferenceID                                 // optional
	IEExtensions                              *ProtocolExtensionContainerLocationReportingRequestTypeExtIEs // optional
}

func (x *LocationReportingRequestType) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LocationReportingRequestTypeOptPresentFlag := []bool{}
	// mandatory field
	if x.EventType == nil {
		return errors.Errorf("EventType is missing")
	}
	// mandatory field
	if x.ReportArea == nil {
		return errors.Errorf("ReportArea is missing")
	}
	// optional field
	if x.AreaOfInterestList != nil {
		LocationReportingRequestTypeOptPresentFlag = append(LocationReportingRequestTypeOptPresentFlag, true)
	} else {
		LocationReportingRequestTypeOptPresentFlag = append(LocationReportingRequestTypeOptPresentFlag, false)
	}
	// optional field
	if x.LocationReportingReferenceIDToBeCancelled != nil {
		LocationReportingRequestTypeOptPresentFlag = append(LocationReportingRequestTypeOptPresentFlag, true)
	} else {
		LocationReportingRequestTypeOptPresentFlag = append(LocationReportingRequestTypeOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		LocationReportingRequestTypeOptPresentFlag = append(LocationReportingRequestTypeOptPresentFlag, true)
	} else {
		LocationReportingRequestTypeOptPresentFlag = append(LocationReportingRequestTypeOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LocationReportingRequestTypeOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.EventType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EventType marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ReportArea.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ReportArea marshal failed")
	}

	// optional field
	if x.AreaOfInterestList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AreaOfInterestList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AreaOfInterestList marshal failed")
		}
	}

	// optional field
	if x.LocationReportingReferenceIDToBeCancelled != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LocationReportingReferenceIDToBeCancelled.Write(pd)
		if err != nil {
			return errors.Wrap(err, "LocationReportingReferenceIDToBeCancelled marshal failed")
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

func (x *LocationReportingRequestType) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LocationReportingRequestTypeOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&LocationReportingRequestTypeOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EventType = new(EventType)
	err = x.EventType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EventType error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ReportArea = new(ReportArea)
	err = x.ReportArea.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ReportArea error")
	}

	// optional field (optPresentFlag index: 0)
	if LocationReportingRequestTypeOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AreaOfInterestList = new(AreaOfInterestList)
		err = x.AreaOfInterestList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AreaOfInterestList error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if LocationReportingRequestTypeOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.LocationReportingReferenceIDToBeCancelled = new(LocationReportingReferenceID)
		err = x.LocationReportingReferenceIDToBeCancelled.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode LocationReportingReferenceIDToBeCancelled error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if LocationReportingRequestTypeOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLocationReportingRequestTypeExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *LocationReportingRequestType) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *LocationReportingRequestType) ReadIE(pd *aper.PerBitData) error {
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
