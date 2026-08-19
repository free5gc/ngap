package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &UERadioCapabilityForPaging{}

type UERadioCapabilityForPaging struct {
	UERadioCapabilityForPagingOfNR    *UERadioCapabilityForPagingOfNR                             // optional
	UERadioCapabilityForPagingOfEUTRA *UERadioCapabilityForPagingOfEUTRA                          // optional
	IEExtensions                      *ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs // optional
}

func (x *UERadioCapabilityForPaging) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UERadioCapabilityForPagingOptPresentFlag := []bool{}
	// optional field
	if x.UERadioCapabilityForPagingOfNR != nil {
		UERadioCapabilityForPagingOptPresentFlag = append(UERadioCapabilityForPagingOptPresentFlag, true)
	} else {
		UERadioCapabilityForPagingOptPresentFlag = append(UERadioCapabilityForPagingOptPresentFlag, false)
	}
	// optional field
	if x.UERadioCapabilityForPagingOfEUTRA != nil {
		UERadioCapabilityForPagingOptPresentFlag = append(UERadioCapabilityForPagingOptPresentFlag, true)
	} else {
		UERadioCapabilityForPagingOptPresentFlag = append(UERadioCapabilityForPagingOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UERadioCapabilityForPagingOptPresentFlag = append(UERadioCapabilityForPagingOptPresentFlag, true)
	} else {
		UERadioCapabilityForPagingOptPresentFlag = append(UERadioCapabilityForPagingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UERadioCapabilityForPagingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.UERadioCapabilityForPagingOfNR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UERadioCapabilityForPagingOfNR.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UERadioCapabilityForPagingOfNR marshal failed")
		}
	}

	// optional field
	if x.UERadioCapabilityForPagingOfEUTRA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.UERadioCapabilityForPagingOfEUTRA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "UERadioCapabilityForPagingOfEUTRA marshal failed")
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

func (x *UERadioCapabilityForPaging) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UERadioCapabilityForPagingOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&UERadioCapabilityForPagingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if UERadioCapabilityForPagingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.UERadioCapabilityForPagingOfNR = new(UERadioCapabilityForPagingOfNR)
		err = x.UERadioCapabilityForPagingOfNR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UERadioCapabilityForPagingOfNR error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UERadioCapabilityForPagingOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.UERadioCapabilityForPagingOfEUTRA = new(UERadioCapabilityForPagingOfEUTRA)
		err = x.UERadioCapabilityForPagingOfEUTRA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode UERadioCapabilityForPagingOfEUTRA error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if UERadioCapabilityForPagingOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUERadioCapabilityForPagingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *UERadioCapabilityForPaging) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *UERadioCapabilityForPaging) ReadIE(pd *aper.PerBitData) error {
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
