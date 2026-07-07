package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type PNINPNMobilityInformation struct {
	AllowedPNINPIList *AllowedPNINPNList
	IEExtensions      *ProtocolExtensionContainerPNINPNMobilityInformationExtIEs // optional
}

func (x *PNINPNMobilityInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PNINPNMobilityInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.AllowedPNINPIList == nil {
		return errors.Errorf("AllowedPNINPIList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PNINPNMobilityInformationOptPresentFlag = append(PNINPNMobilityInformationOptPresentFlag, true)
	} else {
		PNINPNMobilityInformationOptPresentFlag = append(PNINPNMobilityInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PNINPNMobilityInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AllowedPNINPIList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AllowedPNINPIList marshal failed")
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

func (x *PNINPNMobilityInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PNINPNMobilityInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PNINPNMobilityInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AllowedPNINPIList = new(AllowedPNINPNList)
	err = x.AllowedPNINPIList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AllowedPNINPIList error")
	}

	// optional field (optPresentFlag index: 0)
	if PNINPNMobilityInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPNINPNMobilityInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
