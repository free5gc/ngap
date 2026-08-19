package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &TimeSyncAssistanceInfo{}

const ( /* Enum Type */
	TimeSyncAssistanceInfoTimeDistributionIndicationPresentEnabled  aper.Enumerated = 0
	TimeSyncAssistanceInfoTimeDistributionIndicationPresentDisabled aper.Enumerated = 1
)

type TimeSyncAssistanceInfo struct {
	TimeDistributionIndication *aper.Enumerated                                        // valueExt,valueLB:0,valueUB:1
	UUTimeSyncErrorBudget      *int64                                                  // valueExt,valueLB:1,valueUB:1000000,optional
	IEExtensions               *ProtocolExtensionContainerTimeSyncAssistanceInfoExtIEs // optional
}

func (x *TimeSyncAssistanceInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TimeSyncAssistanceInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.TimeDistributionIndication == nil {
		return errors.Errorf("TimeDistributionIndication is missing")
	}
	// optional field
	if x.UUTimeSyncErrorBudget != nil {
		TimeSyncAssistanceInfoOptPresentFlag = append(TimeSyncAssistanceInfoOptPresentFlag, true)
	} else {
		TimeSyncAssistanceInfoOptPresentFlag = append(TimeSyncAssistanceInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TimeSyncAssistanceInfoOptPresentFlag = append(TimeSyncAssistanceInfoOptPresentFlag, true)
	} else {
		TimeSyncAssistanceInfoOptPresentFlag = append(TimeSyncAssistanceInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TimeSyncAssistanceInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.TimeDistributionIndication), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// optional field
	if x.UUTimeSyncErrorBudget != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 1000000
		err = pd.WriteInteger(*(x.UUTimeSyncErrorBudget), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *TimeSyncAssistanceInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TimeSyncAssistanceInfoOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TimeSyncAssistanceInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.TimeDistributionIndication = new(aper.Enumerated)
	*(x.TimeDistributionIndication), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if TimeSyncAssistanceInfoOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 1000000
		x.UUTimeSyncErrorBudget = new(int64)
		*(x.UUTimeSyncErrorBudget), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if TimeSyncAssistanceInfoOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTimeSyncAssistanceInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *TimeSyncAssistanceInfo) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *TimeSyncAssistanceInfo) ReadIE(pd *aper.PerBitData) error {
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
