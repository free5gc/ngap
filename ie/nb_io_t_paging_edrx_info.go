package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &NBIoTPagingEDRXInfo{}

type NBIoTPagingEDRXInfo struct {
	NBIoTPagingEDRXCycle  *NBIoTPagingEDRXCycle                                // valueExt,valueLB:0,valueUB:13
	NBIoTPagingTimeWindow *NBIoTPagingTimeWindow                               // valueExt,valueLB:0,valueUB:15,optional
	IEExtensions          *ProtocolExtensionContainerNBIoTPagingEDRXInfoExtIEs // optional
}

func (x *NBIoTPagingEDRXInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NBIoTPagingEDRXInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.NBIoTPagingEDRXCycle == nil {
		return errors.Errorf("NBIoTPagingEDRXCycle is missing")
	}
	// optional field
	if x.NBIoTPagingTimeWindow != nil {
		NBIoTPagingEDRXInfoOptPresentFlag = append(NBIoTPagingEDRXInfoOptPresentFlag, true)
	} else {
		NBIoTPagingEDRXInfoOptPresentFlag = append(NBIoTPagingEDRXInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		NBIoTPagingEDRXInfoOptPresentFlag = append(NBIoTPagingEDRXInfoOptPresentFlag, true)
	} else {
		NBIoTPagingEDRXInfoOptPresentFlag = append(NBIoTPagingEDRXInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NBIoTPagingEDRXInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NBIoTPagingEDRXCycle.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NBIoTPagingEDRXCycle marshal failed")
	}

	// optional field
	if x.NBIoTPagingTimeWindow != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NBIoTPagingTimeWindow.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NBIoTPagingTimeWindow marshal failed")
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

func (x *NBIoTPagingEDRXInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NBIoTPagingEDRXInfoOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&NBIoTPagingEDRXInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NBIoTPagingEDRXCycle = new(NBIoTPagingEDRXCycle)
	err = x.NBIoTPagingEDRXCycle.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NBIoTPagingEDRXCycle error")
	}

	// optional field (optPresentFlag index: 0)
	if NBIoTPagingEDRXInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NBIoTPagingTimeWindow = new(NBIoTPagingTimeWindow)
		err = x.NBIoTPagingTimeWindow.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NBIoTPagingTimeWindow error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if NBIoTPagingEDRXInfoOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNBIoTPagingEDRXInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *NBIoTPagingEDRXInfo) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *NBIoTPagingEDRXInfo) ReadIE(pd *aper.PerBitData) error {
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
