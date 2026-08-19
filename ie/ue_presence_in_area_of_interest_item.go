package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID *LocationReportingReferenceID
	UEPresence                   *UEPresence                                                     // valueExt,valueLB:0,valueUB:2
	IEExtensions                 *ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs // optional
}

func (x *UEPresenceInAreaOfInterestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEPresenceInAreaOfInterestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.LocationReportingReferenceID == nil {
		return errors.Errorf("LocationReportingReferenceID is missing")
	}
	// mandatory field
	if x.UEPresence == nil {
		return errors.Errorf("UEPresence is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UEPresenceInAreaOfInterestItemOptPresentFlag = append(UEPresenceInAreaOfInterestItemOptPresentFlag, true)
	} else {
		UEPresenceInAreaOfInterestItemOptPresentFlag = append(UEPresenceInAreaOfInterestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEPresenceInAreaOfInterestItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.LocationReportingReferenceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LocationReportingReferenceID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.UEPresence.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UEPresence marshal failed")
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

func (x *UEPresenceInAreaOfInterestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEPresenceInAreaOfInterestItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UEPresenceInAreaOfInterestItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LocationReportingReferenceID = new(LocationReportingReferenceID)
	err = x.LocationReportingReferenceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LocationReportingReferenceID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UEPresence = new(UEPresence)
	err = x.UEPresence.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UEPresence error")
	}

	// optional field (optPresentFlag index: 0)
	if UEPresenceInAreaOfInterestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
