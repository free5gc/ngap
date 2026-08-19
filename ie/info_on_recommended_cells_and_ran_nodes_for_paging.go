package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &InfoOnRecommendedCellsAndRANNodesForPaging{}

type InfoOnRecommendedCellsAndRANNodesForPaging struct {
	RecommendedCellsForPaging  *RecommendedCellsForPaging                                                  // valueExt
	RecommendRANNodesForPaging *RecommendedRANNodesForPaging                                               // valueExt
	IEExtensions               *ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs // optional
}

func (x *InfoOnRecommendedCellsAndRANNodesForPaging) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag := []bool{}
	// mandatory field
	if x.RecommendedCellsForPaging == nil {
		return errors.Errorf("RecommendedCellsForPaging is missing")
	}
	// mandatory field
	if x.RecommendRANNodesForPaging == nil {
		return errors.Errorf("RecommendRANNodesForPaging is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag = append(InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag, true)
	} else {
		InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag = append(InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RecommendedCellsForPaging.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RecommendedCellsForPaging marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.RecommendRANNodesForPaging.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RecommendRANNodesForPaging marshal failed")
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

func (x *InfoOnRecommendedCellsAndRANNodesForPaging) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RecommendedCellsForPaging = new(RecommendedCellsForPaging)
	err = x.RecommendedCellsForPaging.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RecommendedCellsForPaging error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RecommendRANNodesForPaging = new(RecommendedRANNodesForPaging)
	err = x.RecommendRANNodesForPaging.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RecommendRANNodesForPaging error")
	}

	// optional field (optPresentFlag index: 0)
	if InfoOnRecommendedCellsAndRANNodesForPagingOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerInfoOnRecommendedCellsAndRANNodesForPagingExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}

func (x *InfoOnRecommendedCellsAndRANNodesForPaging) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
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

func (x *InfoOnRecommendedCellsAndRANNodesForPaging) ReadIE(pd *aper.PerBitData) error {
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
