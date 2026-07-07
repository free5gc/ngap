package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &FiveGProSePC5QoSParameters{}

type FiveGProSePC5QoSParameters struct {
	FiveGProSepc5QoSFlowList           *FiveGProSePC5QoSFlowList
	FiveGProSepc5LinkAggregateBitRates *BitRate                                                    // optional
	IEExtensions                       *ProtocolExtensionContainerFiveGProSePC5QoSParametersExtIEs // optional
}

func (x *FiveGProSePC5QoSParameters) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	FiveGProSePC5QoSParametersOptPresentFlag := []bool{}
	// mandatory field
	if x.FiveGProSepc5QoSFlowList == nil {
		return errors.Errorf("FiveGProSepc5QoSFlowList is missing")
	}
	// optional field
	if x.FiveGProSepc5LinkAggregateBitRates != nil {
		FiveGProSePC5QoSParametersOptPresentFlag = append(FiveGProSePC5QoSParametersOptPresentFlag, true)
	} else {
		FiveGProSePC5QoSParametersOptPresentFlag = append(FiveGProSePC5QoSParametersOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		FiveGProSePC5QoSParametersOptPresentFlag = append(FiveGProSePC5QoSParametersOptPresentFlag, true)
	} else {
		FiveGProSePC5QoSParametersOptPresentFlag = append(FiveGProSePC5QoSParametersOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(FiveGProSePC5QoSParametersOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.FiveGProSepc5QoSFlowList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FiveGProSepc5QoSFlowList marshal failed")
	}

	// optional field
	if x.FiveGProSepc5LinkAggregateBitRates != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.FiveGProSepc5LinkAggregateBitRates.Write(pd)
		if err != nil {
			return errors.Wrap(err, "FiveGProSepc5LinkAggregateBitRates marshal failed")
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

func (x *FiveGProSePC5QoSParameters) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	FiveGProSePC5QoSParametersOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&FiveGProSePC5QoSParametersOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FiveGProSepc5QoSFlowList = new(FiveGProSePC5QoSFlowList)
	err = x.FiveGProSepc5QoSFlowList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FiveGProSepc5QoSFlowList error")
	}

	// optional field (optPresentFlag index: 0)
	if FiveGProSePC5QoSParametersOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.FiveGProSepc5LinkAggregateBitRates = new(BitRate)
		err = x.FiveGProSepc5LinkAggregateBitRates.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode FiveGProSepc5LinkAggregateBitRates error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if FiveGProSePC5QoSParametersOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerFiveGProSePC5QoSParametersExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *FiveGProSePC5QoSParameters) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *FiveGProSePC5QoSParameters) ReadIE(pd *aper.PerBitData) error {
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
