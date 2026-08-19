package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type SNPNMobilityInformation struct {
	ServingNID   *NID
	IEExtensions *ProtocolExtensionContainerSNPNMobilityInformationExtIEs // optional
}

func (x *SNPNMobilityInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SNPNMobilityInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.ServingNID == nil {
		return errors.Errorf("ServingNID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SNPNMobilityInformationOptPresentFlag = append(SNPNMobilityInformationOptPresentFlag, true)
	} else {
		SNPNMobilityInformationOptPresentFlag = append(SNPNMobilityInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SNPNMobilityInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ServingNID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ServingNID marshal failed")
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

func (x *SNPNMobilityInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SNPNMobilityInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SNPNMobilityInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ServingNID = new(NID)
	err = x.ServingNID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ServingNID error")
	}

	// optional field (optPresentFlag index: 0)
	if SNPNMobilityInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSNPNMobilityInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
