package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &OnDemandPRSInfo{}

type OnDemandPRSInfo struct {
	OnDemandPRSRequestAllowed             *aper.BitString                                  // sizeLB:16,sizeUB:16
	AllowedResourceSetPeriodicityValues   *aper.BitString                                  // sizeLB:24,sizeUB:24,optional
	AllowedPRSBandwidthValues             *aper.BitString                                  // sizeLB:64,sizeUB:64,optional
	AllowedResourceRepetitionFactorValues *aper.BitString                                  // sizeLB:8,sizeUB:8,optional
	AllowedResourceNumberOfSymbolsValues  *aper.BitString                                  // sizeLB:8,sizeUB:8,optional
	AllowedCombSizeValues                 *aper.BitString                                  // sizeLB:8,sizeUB:8,optional
	IEExtensions                          *ProtocolExtensionContainerOnDemandPRSInfoExtIEs // optional
}

func (x *OnDemandPRSInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OnDemandPRSInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.OnDemandPRSRequestAllowed == nil {
		return errors.Errorf("OnDemandPRSRequestAllowed is missing")
	}
	// optional field
	if x.AllowedResourceSetPeriodicityValues != nil {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, true)
	} else {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, false)
	}
	// optional field
	if x.AllowedPRSBandwidthValues != nil {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, true)
	} else {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, false)
	}
	// optional field
	if x.AllowedResourceRepetitionFactorValues != nil {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, true)
	} else {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, false)
	}
	// optional field
	if x.AllowedResourceNumberOfSymbolsValues != nil {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, true)
	} else {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, false)
	}
	// optional field
	if x.AllowedCombSizeValues != nil {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, true)
	} else {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, true)
	} else {
		OnDemandPRSInfoOptPresentFlag = append(OnDemandPRSInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(OnDemandPRSInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write BitString (Pointer)
	*sLb, *sUb = 16, 16
	err = pd.WriteBitString(*(x.OnDemandPRSRequestAllowed), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}

	// optional field
	if x.AllowedResourceSetPeriodicityValues != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 24, 24
		err = pd.WriteBitString(*(x.AllowedResourceSetPeriodicityValues), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
	}

	// optional field
	if x.AllowedPRSBandwidthValues != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 64, 64
		err = pd.WriteBitString(*(x.AllowedPRSBandwidthValues), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
	}

	// optional field
	if x.AllowedResourceRepetitionFactorValues != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 8, 8
		err = pd.WriteBitString(*(x.AllowedResourceRepetitionFactorValues), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
	}

	// optional field
	if x.AllowedResourceNumberOfSymbolsValues != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 8, 8
		err = pd.WriteBitString(*(x.AllowedResourceNumberOfSymbolsValues), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
		}
	}

	// optional field
	if x.AllowedCombSizeValues != nil {
		// Write BitString (Pointer)
		*sLb, *sUb = 8, 8
		err = pd.WriteBitString(*(x.AllowedCombSizeValues), false, sLb, sUb)
		if err != nil {
			return errors.Wrap(err, "bitString marshal failed")
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

func (x *OnDemandPRSInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OnDemandPRSInfoOptPresentFlag := make([]bool, 6)
	err = pd.ReadSequencePreambleBitMap(&OnDemandPRSInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read BitString (Pointer)
	*sLb, *sUb = 16, 16
	x.OnDemandPRSRequestAllowed = new(aper.BitString)
	*(x.OnDemandPRSRequestAllowed), err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}

	// optional field (optPresentFlag index: 0)
	if OnDemandPRSInfoOptPresentFlag[0] {
		// Read BitString (Pointer)
		*sLb, *sUb = 24, 24
		x.AllowedResourceSetPeriodicityValues = new(aper.BitString)
		*(x.AllowedResourceSetPeriodicityValues), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if OnDemandPRSInfoOptPresentFlag[1] {
		// Read BitString (Pointer)
		*sLb, *sUb = 64, 64
		x.AllowedPRSBandwidthValues = new(aper.BitString)
		*(x.AllowedPRSBandwidthValues), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if OnDemandPRSInfoOptPresentFlag[2] {
		// Read BitString (Pointer)
		*sLb, *sUb = 8, 8
		x.AllowedResourceRepetitionFactorValues = new(aper.BitString)
		*(x.AllowedResourceRepetitionFactorValues), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if OnDemandPRSInfoOptPresentFlag[3] {
		// Read BitString (Pointer)
		*sLb, *sUb = 8, 8
		x.AllowedResourceNumberOfSymbolsValues = new(aper.BitString)
		*(x.AllowedResourceNumberOfSymbolsValues), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 4)
	if OnDemandPRSInfoOptPresentFlag[4] {
		// Read BitString (Pointer)
		*sLb, *sUb = 8, 8
		x.AllowedCombSizeValues = new(aper.BitString)
		*(x.AllowedCombSizeValues), err = pd.ReadBitString(false, sLb, sUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
		}
	}

	// optional field (optPresentFlag index: 5)
	if OnDemandPRSInfoOptPresentFlag[5] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerOnDemandPRSInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *OnDemandPRSInfo) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *OnDemandPRSInfo) ReadIE(pd *aper.PerBitData) error {
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
