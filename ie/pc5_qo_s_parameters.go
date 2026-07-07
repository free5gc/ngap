package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &PC5QoSParameters{}

type PC5QoSParameters struct {
	Pc5QoSFlowList           *PC5QoSFlowList
	Pc5LinkAggregateBitRates *BitRate                                          // optional
	IEExtensions             *ProtocolExtensionContainerPC5QoSParametersExtIEs // optional
}

func (x *PC5QoSParameters) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PC5QoSParametersOptPresentFlag := []bool{}
	// mandatory field
	if x.Pc5QoSFlowList == nil {
		return errors.Errorf("Pc5QoSFlowList is missing")
	}
	// optional field
	if x.Pc5LinkAggregateBitRates != nil {
		PC5QoSParametersOptPresentFlag = append(PC5QoSParametersOptPresentFlag, true)
	} else {
		PC5QoSParametersOptPresentFlag = append(PC5QoSParametersOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PC5QoSParametersOptPresentFlag = append(PC5QoSParametersOptPresentFlag, true)
	} else {
		PC5QoSParametersOptPresentFlag = append(PC5QoSParametersOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PC5QoSParametersOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Pc5QoSFlowList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Pc5QoSFlowList marshal failed")
	}

	// optional field
	if x.Pc5LinkAggregateBitRates != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.Pc5LinkAggregateBitRates.Write(pd)
		if err != nil {
			return errors.Wrap(err, "Pc5LinkAggregateBitRates marshal failed")
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

func (x *PC5QoSParameters) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PC5QoSParametersOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PC5QoSParametersOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Pc5QoSFlowList = new(PC5QoSFlowList)
	err = x.Pc5QoSFlowList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Pc5QoSFlowList error")
	}

	// optional field (optPresentFlag index: 0)
	if PC5QoSParametersOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.Pc5LinkAggregateBitRates = new(BitRate)
		err = x.Pc5LinkAggregateBitRates.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode Pc5LinkAggregateBitRates error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PC5QoSParametersOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPC5QoSParametersExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *PC5QoSParameters) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *PC5QoSParameters) ReadIE(pd *aper.PerBitData) error {
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
