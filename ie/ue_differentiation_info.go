package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &UEDifferentiationInfo{}

const ( /* Enum Type */
	UEDifferentiationInfoPeriodicCommunicationIndicatorPresentPeriodically aper.Enumerated = 0
	UEDifferentiationInfoPeriodicCommunicationIndicatorPresentOndemand     aper.Enumerated = 1
)

const ( /* Enum Type */
	UEDifferentiationInfoStationaryIndicationPresentStationary aper.Enumerated = 0
	UEDifferentiationInfoStationaryIndicationPresentMobile     aper.Enumerated = 1
)

const ( /* Enum Type */
	UEDifferentiationInfoTrafficProfilePresentSinglePacket    aper.Enumerated = 0
	UEDifferentiationInfoTrafficProfilePresentDualPackets     aper.Enumerated = 1
	UEDifferentiationInfoTrafficProfilePresentMultiplePackets aper.Enumerated = 2
)

const ( /* Enum Type */
	UEDifferentiationInfoBatteryIndicationPresentBatteryPowered                             aper.Enumerated = 0
	UEDifferentiationInfoBatteryIndicationPresentBatteryPoweredNotRechargeableOrReplaceable aper.Enumerated = 1
	UEDifferentiationInfoBatteryIndicationPresentNotBatteryPowered                          aper.Enumerated = 2
)

type UEDifferentiationInfo struct {
	PeriodicCommunicationIndicator *aper.Enumerated                                       // valueExt,valueLB:0,valueUB:1,optional
	PeriodicTime                   *int64                                                 // valueExt,valueLB:1,valueUB:3600,optional
	ScheduledCommunicationTime     *ScheduledCommunicationTime                            // valueExt,optional
	StationaryIndication           *aper.Enumerated                                       // valueExt,valueLB:0,valueUB:1,optional
	TrafficProfile                 *aper.Enumerated                                       // valueExt,valueLB:0,valueUB:2,optional
	BatteryIndication              *aper.Enumerated                                       // valueExt,valueLB:0,valueUB:2,optional
	IEExtensions                   *ProtocolExtensionContainerUEDifferentiationInfoExtIEs // optional
}

func (x *UEDifferentiationInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEDifferentiationInfoOptPresentFlag := []bool{}
	// optional field
	if x.PeriodicCommunicationIndicator != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}
	// optional field
	if x.PeriodicTime != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}
	// optional field
	if x.ScheduledCommunicationTime != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}
	// optional field
	if x.StationaryIndication != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}
	// optional field
	if x.TrafficProfile != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}
	// optional field
	if x.BatteryIndication != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, true)
	} else {
		UEDifferentiationInfoOptPresentFlag = append(UEDifferentiationInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEDifferentiationInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PeriodicCommunicationIndicator != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 1
		err = pd.WriteEnumerated(*(x.PeriodicCommunicationIndicator), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.PeriodicTime != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 3600
		err = pd.WriteInteger(*(x.PeriodicTime), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.ScheduledCommunicationTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ScheduledCommunicationTime.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ScheduledCommunicationTime marshal failed")
		}
	}

	// optional field
	if x.StationaryIndication != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 1
		err = pd.WriteEnumerated(*(x.StationaryIndication), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.TrafficProfile != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 2
		err = pd.WriteEnumerated(*(x.TrafficProfile), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.BatteryIndication != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 2
		err = pd.WriteEnumerated(*(x.BatteryIndication), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
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

func (x *UEDifferentiationInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEDifferentiationInfoOptPresentFlag := make([]bool, 7)
	err = pd.ReadSequencePreambleBitMap(&UEDifferentiationInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if UEDifferentiationInfoOptPresentFlag[0] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 1
		x.PeriodicCommunicationIndicator = new(aper.Enumerated)
		*(x.PeriodicCommunicationIndicator), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if UEDifferentiationInfoOptPresentFlag[1] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 3600
		x.PeriodicTime = new(int64)
		*(x.PeriodicTime), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if UEDifferentiationInfoOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.ScheduledCommunicationTime = new(ScheduledCommunicationTime)
		err = x.ScheduledCommunicationTime.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ScheduledCommunicationTime error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if UEDifferentiationInfoOptPresentFlag[3] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 1
		x.StationaryIndication = new(aper.Enumerated)
		*(x.StationaryIndication), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 4)
	if UEDifferentiationInfoOptPresentFlag[4] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 2
		x.TrafficProfile = new(aper.Enumerated)
		*(x.TrafficProfile), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 5)
	if UEDifferentiationInfoOptPresentFlag[5] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 2
		x.BatteryIndication = new(aper.Enumerated)
		*(x.BatteryIndication), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 6)
	if UEDifferentiationInfoOptPresentFlag[6] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEDifferentiationInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *UEDifferentiationInfo) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *UEDifferentiationInfo) ReadIE(pd *aper.PerBitData) error {
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
