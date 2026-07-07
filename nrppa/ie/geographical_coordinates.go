package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type GeographicalCoordinates struct {
	TRPPositionDefinitionType *TRPPositionDefinitionType                               // valueLB:0,valueUB:2
	DLPRSResourceCoordinates  *DLPRSResourceCoordinates                                // valueExt,optional
	IEExtensions              *ProtocolExtensionContainerGeographicalCoordinatesExtIEs // optional
}

func (x *GeographicalCoordinates) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	GeographicalCoordinatesOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPPositionDefinitionType == nil {
		return errors.Errorf("TRPPositionDefinitionType is missing")
	}
	// optional field
	if x.DLPRSResourceCoordinates != nil {
		GeographicalCoordinatesOptPresentFlag = append(GeographicalCoordinatesOptPresentFlag, true)
	} else {
		GeographicalCoordinatesOptPresentFlag = append(GeographicalCoordinatesOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		GeographicalCoordinatesOptPresentFlag = append(GeographicalCoordinatesOptPresentFlag, true)
	} else {
		GeographicalCoordinatesOptPresentFlag = append(GeographicalCoordinatesOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(GeographicalCoordinatesOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPPositionDefinitionType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPPositionDefinitionType marshal failed")
	}

	// optional field
	if x.DLPRSResourceCoordinates != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLPRSResourceCoordinates.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLPRSResourceCoordinates marshal failed")
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

func (x *GeographicalCoordinates) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	GeographicalCoordinatesOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&GeographicalCoordinatesOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPPositionDefinitionType = new(TRPPositionDefinitionType)
	err = x.TRPPositionDefinitionType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPPositionDefinitionType error")
	}

	// optional field (optPresentFlag index: 0)
	if GeographicalCoordinatesOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLPRSResourceCoordinates = new(DLPRSResourceCoordinates)
		err = x.DLPRSResourceCoordinates.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLPRSResourceCoordinates error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if GeographicalCoordinatesOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerGeographicalCoordinatesExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
