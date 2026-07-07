package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SpatialDirectionInformation struct {
	NRPRSBeamInformation *NRPRSBeamInformation                                        // valueExt
	IEExtensions         *ProtocolExtensionContainerSpatialDirectionInformationExtIEs // optional
}

func (x *SpatialDirectionInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SpatialDirectionInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.NRPRSBeamInformation == nil {
		return errors.Errorf("NRPRSBeamInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SpatialDirectionInformationOptPresentFlag = append(SpatialDirectionInformationOptPresentFlag, true)
	} else {
		SpatialDirectionInformationOptPresentFlag = append(SpatialDirectionInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SpatialDirectionInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NRPRSBeamInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRPRSBeamInformation marshal failed")
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

func (x *SpatialDirectionInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SpatialDirectionInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SpatialDirectionInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRPRSBeamInformation = new(NRPRSBeamInformation)
	err = x.NRPRSBeamInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRPRSBeamInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if SpatialDirectionInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSpatialDirectionInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
