package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSResourceSetItem struct {
	NumberOfSRSResourcePerSet    *int64                                              // valueExt,valueLB:1,valueUB:16,optional
	PeriodicityList              *PeriodicityList                                    // optional
	SpatialRelationInformation   *SpatialRelationInfo                                // valueExt,optional
	PathlossReferenceInformation *PathlossReferenceInformation                       // valueExt,optional
	IEExtensions                 *ProtocolExtensionContainerSRSResourceSetItemExtIEs // optional
}

func (x *SRSResourceSetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceSetItemOptPresentFlag := []bool{}
	// optional field
	if x.NumberOfSRSResourcePerSet != nil {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, true)
	} else {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.PeriodicityList != nil {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, true)
	} else {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.SpatialRelationInformation != nil {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, true)
	} else {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.PathlossReferenceInformation != nil {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, true)
	} else {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, true)
	} else {
		SRSResourceSetItemOptPresentFlag = append(SRSResourceSetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.NumberOfSRSResourcePerSet != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 16
		err = pd.WriteInteger(*(x.NumberOfSRSResourcePerSet), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.PeriodicityList != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PeriodicityList.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PeriodicityList marshal failed")
		}
	}

	// optional field
	if x.SpatialRelationInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SpatialRelationInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SpatialRelationInformation marshal failed")
		}
	}

	// optional field
	if x.PathlossReferenceInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PathlossReferenceInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PathlossReferenceInformation marshal failed")
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

func (x *SRSResourceSetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceSetItemOptPresentFlag := make([]bool, 5)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if SRSResourceSetItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 16
		x.NumberOfSRSResourcePerSet = new(int64)
		*(x.NumberOfSRSResourcePerSet), err = pd.ReadInteger(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if SRSResourceSetItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PeriodicityList = new(PeriodicityList)
		err = x.PeriodicityList.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PeriodicityList error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if SRSResourceSetItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SpatialRelationInformation = new(SpatialRelationInfo)
		err = x.SpatialRelationInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SpatialRelationInformation error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if SRSResourceSetItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.PathlossReferenceInformation = new(PathlossReferenceInformation)
		err = x.PathlossReferenceInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PathlossReferenceInformation error")
		}
	}

	// optional field (optPresentFlag index: 4)
	if SRSResourceSetItemOptPresentFlag[4] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSResourceSetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
