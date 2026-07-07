package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type DLPRSResourceSetARP struct {
	DlPRSResourceSetID          *PRSResourceSetID
	DLPRSResourceSetARPLocation *DLPRSResourceSetARPLocation // valueLB:0,valueUB:2
	/* Sequence of = 35, FULL Name = struct DLPRSResourceSetARP__listofDL_PRSResourceARP */
	/* Type Name = DLPRSResourceARP */
	/* Sequence Of Embed */
	ListofDLPRSResourceARP []DLPRSResourceARP                                   // valueExt,sizeLB:1,sizeUB:64
	IEExtensions           *ProtocolExtensionContainerDLPRSResourceSetARPExtIEs // optional
}

func (x *DLPRSResourceSetARP) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceSetARPOptPresentFlag := []bool{}
	// mandatory field
	if x.DlPRSResourceSetID == nil {
		return errors.Errorf("DlPRSResourceSetID is missing")
	}
	// mandatory field
	if x.DLPRSResourceSetARPLocation == nil {
		return errors.Errorf("DLPRSResourceSetARPLocation is missing")
	}
	// mandatory field
	if x.ListofDLPRSResourceARP == nil {
		return errors.Errorf("ListofDLPRSResourceARP is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		DLPRSResourceSetARPOptPresentFlag = append(DLPRSResourceSetARPOptPresentFlag, true)
	} else {
		DLPRSResourceSetARPOptPresentFlag = append(DLPRSResourceSetARPOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceSetARPOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DlPRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DlPRSResourceSetID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DLPRSResourceSetARPLocation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLPRSResourceSetARPLocation marshal failed")
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 64
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.ListofDLPRSResourceARP)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.ListofDLPRSResourceARP {
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

func (x *DLPRSResourceSetARP) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceSetARPOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceSetARPOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DlPRSResourceSetID = new(PRSResourceSetID)
	err = x.DlPRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DlPRSResourceSetID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLPRSResourceSetARPLocation = new(DLPRSResourceSetARPLocation)
	err = x.DLPRSResourceSetARPLocation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLPRSResourceSetARPLocation error")
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 64
	var numElementsListofDLPRSResourceARP uint64
	numElementsListofDLPRSResourceARP, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.ListofDLPRSResourceARP = []DLPRSResourceARP{}
	for i := 0; i < int(numElementsListofDLPRSResourceARP); i++ {
		var val DLPRSResourceARP
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.ListofDLPRSResourceARP = append(x.ListofDLPRSResourceARP, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if DLPRSResourceSetARPOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDLPRSResourceSetARPExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
