package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &SNSSAI{}

type SNSSAI struct {
	SST          *SST
	SD           *SD                                     // optional
	IEExtensions *ProtocolExtensionContainerSNSSAIExtIEs // optional
}

func (x *SNSSAI) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SNSSAIOptPresentFlag := []bool{}
	// mandatory field
	if x.SST == nil {
		return errors.Errorf("SST is missing")
	}
	// optional field
	if x.SD != nil {
		SNSSAIOptPresentFlag = append(SNSSAIOptPresentFlag, true)
	} else {
		SNSSAIOptPresentFlag = append(SNSSAIOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SNSSAIOptPresentFlag = append(SNSSAIOptPresentFlag, true)
	} else {
		SNSSAIOptPresentFlag = append(SNSSAIOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SNSSAIOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SST.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SST marshal failed")
	}

	// optional field
	if x.SD != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SD.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SD marshal failed")
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

func (x *SNSSAI) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SNSSAIOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SNSSAIOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SST = new(SST)
	err = x.SST.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SST error")
	}

	// optional field (optPresentFlag index: 0)
	if SNSSAIOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SD = new(SD)
		err = x.SD.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SD error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if SNSSAIOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSNSSAIExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *SNSSAI) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *SNSSAI) ReadIE(pd *aper.PerBitData) error {
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
