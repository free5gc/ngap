package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &TraceActivation{}

type TraceActivation struct {
	NGRANTraceID                   *NGRANTraceID
	InterfacesToTrace              *InterfacesToTrace
	TraceDepth                     *TraceDepth // valueExt,valueLB:0,valueUB:5
	TraceCollectionEntityIPAddress *TransportLayerAddress
	IEExtensions                   *ProtocolExtensionContainerTraceActivationExtIEs // optional
}

func (x *TraceActivation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TraceActivationOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANTraceID == nil {
		return errors.Errorf("NGRANTraceID is missing")
	}
	// mandatory field
	if x.InterfacesToTrace == nil {
		return errors.Errorf("InterfacesToTrace is missing")
	}
	// mandatory field
	if x.TraceDepth == nil {
		return errors.Errorf("TraceDepth is missing")
	}
	// mandatory field
	if x.TraceCollectionEntityIPAddress == nil {
		return errors.Errorf("TraceCollectionEntityIPAddress is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TraceActivationOptPresentFlag = append(TraceActivationOptPresentFlag, true)
	} else {
		TraceActivationOptPresentFlag = append(TraceActivationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TraceActivationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANTraceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANTraceID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.InterfacesToTrace.Write(pd)
	if err != nil {
		return errors.Wrap(err, "InterfacesToTrace marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TraceDepth.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TraceDepth marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TraceCollectionEntityIPAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TraceCollectionEntityIPAddress marshal failed")
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

func (x *TraceActivation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TraceActivationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TraceActivationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANTraceID = new(NGRANTraceID)
	err = x.NGRANTraceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANTraceID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.InterfacesToTrace = new(InterfacesToTrace)
	err = x.InterfacesToTrace.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode InterfacesToTrace error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TraceDepth = new(TraceDepth)
	err = x.TraceDepth.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TraceDepth error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TraceCollectionEntityIPAddress = new(TransportLayerAddress)
	err = x.TraceCollectionEntityIPAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TraceCollectionEntityIPAddress error")
	}

	// optional field (optPresentFlag index: 0)
	if TraceActivationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTraceActivationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *TraceActivation) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *TraceActivation) ReadIE(pd *aper.PerBitData) error {
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
