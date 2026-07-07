package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &PagingAssisDataforCEcapabUE{}

type PagingAssisDataforCEcapabUE struct {
	EUTRACGI                 *EUTRACGI // valueExt
	CoverageEnhancementLevel *CoverageEnhancementLevel
	IEExtensions             *ProtocolExtensionContainerPagingAssisDataforCEcapabUEExtIEs // optional
}

func (x *PagingAssisDataforCEcapabUE) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PagingAssisDataforCEcapabUEOptPresentFlag := []bool{}
	// mandatory field
	if x.EUTRACGI == nil {
		return errors.Errorf("EUTRACGI is missing")
	}
	// mandatory field
	if x.CoverageEnhancementLevel == nil {
		return errors.Errorf("CoverageEnhancementLevel is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PagingAssisDataforCEcapabUEOptPresentFlag = append(PagingAssisDataforCEcapabUEOptPresentFlag, true)
	} else {
		PagingAssisDataforCEcapabUEOptPresentFlag = append(PagingAssisDataforCEcapabUEOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PagingAssisDataforCEcapabUEOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.EUTRACGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EUTRACGI marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.CoverageEnhancementLevel.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CoverageEnhancementLevel marshal failed")
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

func (x *PagingAssisDataforCEcapabUE) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PagingAssisDataforCEcapabUEOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PagingAssisDataforCEcapabUEOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EUTRACGI = new(EUTRACGI)
	err = x.EUTRACGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EUTRACGI error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CoverageEnhancementLevel = new(CoverageEnhancementLevel)
	err = x.CoverageEnhancementLevel.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CoverageEnhancementLevel error")
	}

	// optional field (optPresentFlag index: 0)
	if PagingAssisDataforCEcapabUEOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPagingAssisDataforCEcapabUEExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *PagingAssisDataforCEcapabUE) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *PagingAssisDataforCEcapabUE) ReadIE(pd *aper.PerBitData) error {
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
