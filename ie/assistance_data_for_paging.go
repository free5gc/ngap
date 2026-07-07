package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &AssistanceDataForPaging{}

type AssistanceDataForPaging struct {
	AssistanceDataForRecommendedCells *AssistanceDataForRecommendedCells                       // valueExt,optional
	PagingAttemptInformation          *PagingAttemptInformation                                // valueExt,optional
	IEExtensions                      *ProtocolExtensionContainerAssistanceDataForPagingExtIEs // optional
}

func (x *AssistanceDataForPaging) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceDataForPagingOptPresentFlag := []bool{}
	// optional field
	if x.AssistanceDataForRecommendedCells != nil {
		AssistanceDataForPagingOptPresentFlag = append(AssistanceDataForPagingOptPresentFlag, true)
	} else {
		AssistanceDataForPagingOptPresentFlag = append(AssistanceDataForPagingOptPresentFlag, false)
	}
	// optional field
	if x.PagingAttemptInformation != nil {
		AssistanceDataForPagingOptPresentFlag = append(AssistanceDataForPagingOptPresentFlag, true)
	} else {
		AssistanceDataForPagingOptPresentFlag = append(AssistanceDataForPagingOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AssistanceDataForPagingOptPresentFlag = append(AssistanceDataForPagingOptPresentFlag, true)
	} else {
		AssistanceDataForPagingOptPresentFlag = append(AssistanceDataForPagingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceDataForPagingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.AssistanceDataForRecommendedCells != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AssistanceDataForRecommendedCells.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AssistanceDataForRecommendedCells marshal failed")
		}
	}

	// optional field
	if x.PagingAttemptInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PagingAttemptInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PagingAttemptInformation marshal failed")
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

func (x *AssistanceDataForPaging) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceDataForPagingOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&AssistanceDataForPagingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if AssistanceDataForPagingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AssistanceDataForRecommendedCells = new(AssistanceDataForRecommendedCells)
		err = x.AssistanceDataForRecommendedCells.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AssistanceDataForRecommendedCells error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AssistanceDataForPagingOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PagingAttemptInformation = new(PagingAttemptInformation)
		err = x.PagingAttemptInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PagingAttemptInformation error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if AssistanceDataForPagingOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAssistanceDataForPagingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *AssistanceDataForPaging) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *AssistanceDataForPaging) ReadIE(pd *aper.PerBitData) error {
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
