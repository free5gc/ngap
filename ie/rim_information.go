package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	RIMInformationRIMRSDetectionPresentRsDetected    aper.Enumerated = 0
	RIMInformationRIMRSDetectionPresentRsDisappeared aper.Enumerated = 1
)

type RIMInformation struct {
	TargetgNBSetID *GNBSetID
	RIMRSDetection *aper.Enumerated                                // valueExt,valueLB:0,valueUB:1
	IEExtensions   *ProtocolExtensionContainerRIMInformationExtIEs // optional
}

func (x *RIMInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RIMInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.TargetgNBSetID == nil {
		return errors.Errorf("TargetgNBSetID is missing")
	}
	// mandatory field
	if x.RIMRSDetection == nil {
		return errors.Errorf("RIMRSDetection is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		RIMInformationOptPresentFlag = append(RIMInformationOptPresentFlag, true)
	} else {
		RIMInformationOptPresentFlag = append(RIMInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RIMInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TargetgNBSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TargetgNBSetID marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.RIMRSDetection), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *RIMInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RIMInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&RIMInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TargetgNBSetID = new(GNBSetID)
	err = x.TargetgNBSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TargetgNBSetID error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.RIMRSDetection = new(aper.Enumerated)
	*(x.RIMRSDetection), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if RIMInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRIMInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
