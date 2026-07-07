package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSTransmissionOffInformation struct {
	PRSTransmissionOffIndication *PRSTransmissionOffIndication                                  // valueLB:0,valueUB:3
	IEExtensions                 *ProtocolExtensionContainerPRSTransmissionOffInformationExtIEs // optional
}

func (x *PRSTransmissionOffInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSTransmissionOffIndication == nil {
		return errors.Errorf("PRSTransmissionOffIndication is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSTransmissionOffInformationOptPresentFlag = append(PRSTransmissionOffInformationOptPresentFlag, true)
	} else {
		PRSTransmissionOffInformationOptPresentFlag = append(PRSTransmissionOffInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSTransmissionOffIndication.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSTransmissionOffIndication marshal failed")
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

func (x *PRSTransmissionOffInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSTransmissionOffIndication = new(PRSTransmissionOffIndication)
	err = x.PRSTransmissionOffIndication.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSTransmissionOffIndication error")
	}

	// optional field (optPresentFlag index: 0)
	if PRSTransmissionOffInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSTransmissionOffInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
