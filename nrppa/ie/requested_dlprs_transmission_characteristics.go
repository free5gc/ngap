package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type RequestedDLPRSTransmissionCharacteristics struct {
	RequestedDLPRSResourceSetList *RequestedDLPRSResourceSetList
	NumberofFrequencyLayers       *int64                                                                     // valueLB:1,valueUB:4,optional
	StartTimeAndDuration          *StartTimeAndDuration                                                      // valueExt,optional
	IEExtensions                  *ProtocolExtensionContainerRequestedDLPRSTransmissionCharacteristicsExtIEs // optional
}

func (x *RequestedDLPRSTransmissionCharacteristics) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedDLPRSTransmissionCharacteristicsOptPresentFlag := []bool{}
	// mandatory field
	if x.RequestedDLPRSResourceSetList == nil {
		return errors.Errorf("RequestedDLPRSResourceSetList is missing")
	}
	// optional field
	if x.NumberofFrequencyLayers != nil {
		RequestedDLPRSTransmissionCharacteristicsOptPresentFlag = append(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, true)
	} else {
		RequestedDLPRSTransmissionCharacteristicsOptPresentFlag = append(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, false)
	}
	// optional field
	if x.StartTimeAndDuration != nil {
		RequestedDLPRSTransmissionCharacteristicsOptPresentFlag = append(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, true)
	} else {
		RequestedDLPRSTransmissionCharacteristicsOptPresentFlag = append(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		RequestedDLPRSTransmissionCharacteristicsOptPresentFlag = append(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, true)
	} else {
		RequestedDLPRSTransmissionCharacteristicsOptPresentFlag = append(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RequestedDLPRSResourceSetList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RequestedDLPRSResourceSetList marshal failed")
	}

	// optional field
	if x.NumberofFrequencyLayers != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 1, 4
		err = pd.WriteInteger(*(x.NumberofFrequencyLayers), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
		}
	}

	// optional field
	if x.StartTimeAndDuration != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.StartTimeAndDuration.Write(pd)
		if err != nil {
			return errors.Wrap(err, "StartTimeAndDuration marshal failed")
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

func (x *RequestedDLPRSTransmissionCharacteristics) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedDLPRSTransmissionCharacteristicsOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&RequestedDLPRSTransmissionCharacteristicsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RequestedDLPRSResourceSetList = new(RequestedDLPRSResourceSetList)
	err = x.RequestedDLPRSResourceSetList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RequestedDLPRSResourceSetList error")
	}

	// optional field (optPresentFlag index: 0)
	if RequestedDLPRSTransmissionCharacteristicsOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 1, 4
		x.NumberofFrequencyLayers = new(int64)
		*(x.NumberofFrequencyLayers), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if RequestedDLPRSTransmissionCharacteristicsOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.StartTimeAndDuration = new(StartTimeAndDuration)
		err = x.StartTimeAndDuration.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode StartTimeAndDuration error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if RequestedDLPRSTransmissionCharacteristicsOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRequestedDLPRSTransmissionCharacteristicsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
