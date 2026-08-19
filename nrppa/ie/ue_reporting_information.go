package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &UEReportingInformation{}

const ( /* Enum Type */
	UEReportingInformationReportingAmountPresentMa0  aper.Enumerated = 0
	UEReportingInformationReportingAmountPresentMa1  aper.Enumerated = 1
	UEReportingInformationReportingAmountPresentMa2  aper.Enumerated = 2
	UEReportingInformationReportingAmountPresentMa4  aper.Enumerated = 3
	UEReportingInformationReportingAmountPresentMa8  aper.Enumerated = 4
	UEReportingInformationReportingAmountPresentMa16 aper.Enumerated = 5
	UEReportingInformationReportingAmountPresentMa32 aper.Enumerated = 6
	UEReportingInformationReportingAmountPresentMa64 aper.Enumerated = 7
)

const ( /* Enum Type */
	UEReportingInformationReportingIntervalPresentNone      aper.Enumerated = 0
	UEReportingInformationReportingIntervalPresentOne       aper.Enumerated = 1
	UEReportingInformationReportingIntervalPresentTwo       aper.Enumerated = 2
	UEReportingInformationReportingIntervalPresentFour      aper.Enumerated = 3
	UEReportingInformationReportingIntervalPresentEight     aper.Enumerated = 4
	UEReportingInformationReportingIntervalPresentTen       aper.Enumerated = 5
	UEReportingInformationReportingIntervalPresentSixteen   aper.Enumerated = 6
	UEReportingInformationReportingIntervalPresentTwenty    aper.Enumerated = 7
	UEReportingInformationReportingIntervalPresentThirtyTwo aper.Enumerated = 8
	UEReportingInformationReportingIntervalPresentSixtyFour aper.Enumerated = 9
)

type UEReportingInformation struct {
	ReportingAmount   *aper.Enumerated                                        // valueLB:0,valueUB:7
	ReportingInterval *aper.Enumerated                                        // valueExt,valueLB:0,valueUB:9
	IEExtensions      *ProtocolExtensionContainerUEReportingInformationExtIEs // optional
}

func (x *UEReportingInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEReportingInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.ReportingAmount == nil {
		return errors.Errorf("ReportingAmount is missing")
	}
	// mandatory field
	if x.ReportingInterval == nil {
		return errors.Errorf("ReportingInterval is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UEReportingInformationOptPresentFlag = append(UEReportingInformationOptPresentFlag, true)
	} else {
		UEReportingInformationOptPresentFlag = append(UEReportingInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEReportingInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 7
	err = pd.WriteEnumerated(*(x.ReportingAmount), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 9
	err = pd.WriteEnumerated(*(x.ReportingInterval), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *UEReportingInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEReportingInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UEReportingInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 7
	x.ReportingAmount = new(aper.Enumerated)
	*(x.ReportingAmount), err = pd.ReadEnumerated(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 9
	x.ReportingInterval = new(aper.Enumerated)
	*(x.ReportingInterval), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if UEReportingInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEReportingInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *UEReportingInformation) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *UEReportingInformation) ReadIE(pd *aper.PerBitData) error {
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
