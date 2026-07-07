package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type DLPRSResourceCoordinates struct {
	/* Sequence of = 35, FULL Name = struct DLPRSResourceCoordinates__listofDL_PRSResourceSetARP */
	/* Type Name = DLPRSResourceSetARP */
	/* Sequence Of Embed */
	ListofDLPRSResourceSetARP []DLPRSResourceSetARP                                     // valueExt,sizeLB:1,sizeUB:2
	IEExtensions              *ProtocolExtensionContainerDLPRSResourceCoordinatesExtIEs // optional
}

func (x *DLPRSResourceCoordinates) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceCoordinatesOptPresentFlag := []bool{}
	// mandatory field
	if x.ListofDLPRSResourceSetARP == nil {
		return errors.Errorf("ListofDLPRSResourceSetARP is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		DLPRSResourceCoordinatesOptPresentFlag = append(DLPRSResourceCoordinatesOptPresentFlag, true)
	} else {
		DLPRSResourceCoordinatesOptPresentFlag = append(DLPRSResourceCoordinatesOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceCoordinatesOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Sequence Of
	*sLb, *sUb = 1, 2
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.ListofDLPRSResourceSetARP)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.ListofDLPRSResourceSetARP {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
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

func (x *DLPRSResourceCoordinates) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceCoordinatesOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceCoordinatesOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 2
	var numElementsListofDLPRSResourceSetARP uint64
	numElementsListofDLPRSResourceSetARP, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.ListofDLPRSResourceSetARP = []DLPRSResourceSetARP{}
	for i := 0; i < int(numElementsListofDLPRSResourceSetARP); i++ {
		var val DLPRSResourceSetARP
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.ListofDLPRSResourceSetARP = append(x.ListofDLPRSResourceSetARP, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if DLPRSResourceCoordinatesOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDLPRSResourceCoordinatesExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
