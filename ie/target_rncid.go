package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &TargetRNCID{}

type TargetRNCID struct {
	LAI           *LAI // valueExt
	RNCID         *RNCID
	ExtendedRNCID *ExtendedRNCID                               // optional
	IEExtensions  *ProtocolExtensionContainerTargetRNCIDExtIEs // optional
}

func (x *TargetRNCID) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TargetRNCIDOptPresentFlag := []bool{}
	// mandatory field
	if x.LAI == nil {
		return errors.Errorf("LAI is missing")
	}
	// mandatory field
	if x.RNCID == nil {
		return errors.Errorf("RNCID is missing")
	}
	// optional field
	if x.ExtendedRNCID != nil {
		TargetRNCIDOptPresentFlag = append(TargetRNCIDOptPresentFlag, true)
	} else {
		TargetRNCIDOptPresentFlag = append(TargetRNCIDOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TargetRNCIDOptPresentFlag = append(TargetRNCIDOptPresentFlag, true)
	} else {
		TargetRNCIDOptPresentFlag = append(TargetRNCIDOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TargetRNCIDOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.LAI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LAI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.RNCID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RNCID marshal failed")
	}

	// optional field
	if x.ExtendedRNCID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ExtendedRNCID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ExtendedRNCID marshal failed")
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

func (x *TargetRNCID) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TargetRNCIDOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TargetRNCIDOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LAI = new(LAI)
	err = x.LAI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LAI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RNCID = new(RNCID)
	err = x.RNCID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RNCID error")
	}

	// optional field (optPresentFlag index: 0)
	if TargetRNCIDOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.ExtendedRNCID = new(ExtendedRNCID)
		err = x.ExtendedRNCID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ExtendedRNCID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TargetRNCIDOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTargetRNCIDExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *TargetRNCID) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *TargetRNCID) ReadIE(pd *aper.PerBitData) error {
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
