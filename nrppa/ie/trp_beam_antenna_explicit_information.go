package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPBeamAntennaExplicitInformation struct {
	TrpBeamAntennaAngles *TRPBeamAntennaAngles
	LcsToGcsTranslation  *LCSToGCSTranslation                                               // valueExt,optional
	IEExtensions         *ProtocolExtensionContainerTRPBeamAntennaExplicitInformationExtIEs // optional
}

func (x *TRPBeamAntennaExplicitInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPBeamAntennaExplicitInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.TrpBeamAntennaAngles == nil {
		return errors.Errorf("TrpBeamAntennaAngles is missing")
	}
	// optional field
	if x.LcsToGcsTranslation != nil {
		TRPBeamAntennaExplicitInformationOptPresentFlag = append(TRPBeamAntennaExplicitInformationOptPresentFlag, true)
	} else {
		TRPBeamAntennaExplicitInformationOptPresentFlag = append(TRPBeamAntennaExplicitInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TRPBeamAntennaExplicitInformationOptPresentFlag = append(TRPBeamAntennaExplicitInformationOptPresentFlag, true)
	} else {
		TRPBeamAntennaExplicitInformationOptPresentFlag = append(TRPBeamAntennaExplicitInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPBeamAntennaExplicitInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TrpBeamAntennaAngles.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TrpBeamAntennaAngles marshal failed")
	}

	// optional field
	if x.LcsToGcsTranslation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.LcsToGcsTranslation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "LcsToGcsTranslation marshal failed")
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

func (x *TRPBeamAntennaExplicitInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPBeamAntennaExplicitInformationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPBeamAntennaExplicitInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TrpBeamAntennaAngles = new(TRPBeamAntennaAngles)
	err = x.TrpBeamAntennaAngles.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TrpBeamAntennaAngles error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPBeamAntennaExplicitInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.LcsToGcsTranslation = new(LCSToGCSTranslation)
		err = x.LcsToGcsTranslation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode LcsToGcsTranslation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TRPBeamAntennaExplicitInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPBeamAntennaExplicitInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
