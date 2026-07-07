package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &ECIDMeasurementResult{}

type ECIDMeasurementResult struct {
	ServingCellID            *NGRANCGI // valueExt
	ServingCellTAC           *TAC
	NGRANAccessPointPosition *NGRANAccessPointPosition                              // valueExt,optional
	MeasuredResults          *MeasuredResults                                       // optional
	IEExtensions             *ProtocolExtensionContainerECIDMeasurementResultExtIEs // optional
}

func (x *ECIDMeasurementResult) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ECIDMeasurementResultOptPresentFlag := []bool{}
	// mandatory field
	if x.ServingCellID == nil {
		return errors.Errorf("ServingCellID is missing")
	}
	// mandatory field
	if x.ServingCellTAC == nil {
		return errors.Errorf("ServingCellTAC is missing")
	}
	// optional field
	if x.NGRANAccessPointPosition != nil {
		ECIDMeasurementResultOptPresentFlag = append(ECIDMeasurementResultOptPresentFlag, true)
	} else {
		ECIDMeasurementResultOptPresentFlag = append(ECIDMeasurementResultOptPresentFlag, false)
	}
	// optional field
	if x.MeasuredResults != nil {
		ECIDMeasurementResultOptPresentFlag = append(ECIDMeasurementResultOptPresentFlag, true)
	} else {
		ECIDMeasurementResultOptPresentFlag = append(ECIDMeasurementResultOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ECIDMeasurementResultOptPresentFlag = append(ECIDMeasurementResultOptPresentFlag, true)
	} else {
		ECIDMeasurementResultOptPresentFlag = append(ECIDMeasurementResultOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ECIDMeasurementResultOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ServingCellID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ServingCellID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ServingCellTAC.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ServingCellTAC marshal failed")
	}

	// optional field
	if x.NGRANAccessPointPosition != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.NGRANAccessPointPosition.Write(pd)
		if err != nil {
			return errors.Wrap(err, "NGRANAccessPointPosition marshal failed")
		}
	}

	// optional field
	if x.MeasuredResults != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MeasuredResults.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MeasuredResults marshal failed")
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

func (x *ECIDMeasurementResult) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ECIDMeasurementResultOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&ECIDMeasurementResultOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ServingCellID = new(NGRANCGI)
	err = x.ServingCellID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ServingCellID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ServingCellTAC = new(TAC)
	err = x.ServingCellTAC.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ServingCellTAC error")
	}

	// optional field (optPresentFlag index: 0)
	if ECIDMeasurementResultOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.NGRANAccessPointPosition = new(NGRANAccessPointPosition)
		err = x.NGRANAccessPointPosition.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode NGRANAccessPointPosition error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ECIDMeasurementResultOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MeasuredResults = new(MeasuredResults)
		err = x.MeasuredResults.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MeasuredResults error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ECIDMeasurementResultOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerECIDMeasurementResultExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *ECIDMeasurementResult) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *ECIDMeasurementResult) ReadIE(pd *aper.PerBitData) error {
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
