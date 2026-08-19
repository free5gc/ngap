package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AreaOfInterestItem struct {
	AreaOfInterest               *AreaOfInterest // valueExt
	LocationReportingReferenceID *LocationReportingReferenceID
	IEExtensions                 *ProtocolExtensionContainerAreaOfInterestItemExtIEs // optional
}

func (x *AreaOfInterestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AreaOfInterestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AreaOfInterest == nil {
		return errors.Errorf("AreaOfInterest is missing")
	}
	// mandatory field
	if x.LocationReportingReferenceID == nil {
		return errors.Errorf("LocationReportingReferenceID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AreaOfInterestItemOptPresentFlag = append(AreaOfInterestItemOptPresentFlag, true)
	} else {
		AreaOfInterestItemOptPresentFlag = append(AreaOfInterestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AreaOfInterestItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AreaOfInterest.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AreaOfInterest marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.LocationReportingReferenceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LocationReportingReferenceID marshal failed")
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

func (x *AreaOfInterestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AreaOfInterestItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AreaOfInterestItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AreaOfInterest = new(AreaOfInterest)
	err = x.AreaOfInterest.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AreaOfInterest error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LocationReportingReferenceID = new(LocationReportingReferenceID)
	err = x.LocationReportingReferenceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LocationReportingReferenceID error")
	}

	// optional field (optPresentFlag index: 0)
	if AreaOfInterestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAreaOfInterestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
