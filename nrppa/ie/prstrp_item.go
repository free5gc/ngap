package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSTRPItem struct {
	TRPID                                     *TRPID
	RequestedDLPRSTransmissionCharacteristics *RequestedDLPRSTransmissionCharacteristics  // valueExt,optional
	PRSTransmissionOffInformation             *PRSTransmissionOffInformation              // valueExt,optional
	IEExtensions                              *ProtocolExtensionContainerPRSTRPItemExtIEs // optional
}

func (x *PRSTRPItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTRPItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// optional field
	if x.RequestedDLPRSTransmissionCharacteristics != nil {
		PRSTRPItemOptPresentFlag = append(PRSTRPItemOptPresentFlag, true)
	} else {
		PRSTRPItemOptPresentFlag = append(PRSTRPItemOptPresentFlag, false)
	}
	// optional field
	if x.PRSTransmissionOffInformation != nil {
		PRSTRPItemOptPresentFlag = append(PRSTRPItemOptPresentFlag, true)
	} else {
		PRSTRPItemOptPresentFlag = append(PRSTRPItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PRSTRPItemOptPresentFlag = append(PRSTRPItemOptPresentFlag, true)
	} else {
		PRSTRPItemOptPresentFlag = append(PRSTRPItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSTRPItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPID marshal failed")
	}

	// optional field
	if x.RequestedDLPRSTransmissionCharacteristics != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RequestedDLPRSTransmissionCharacteristics.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RequestedDLPRSTransmissionCharacteristics marshal failed")
		}
	}

	// optional field
	if x.PRSTransmissionOffInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSTransmissionOffInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSTransmissionOffInformation marshal failed")
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

func (x *PRSTRPItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTRPItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&PRSTRPItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPID = new(TRPID)
	err = x.TRPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPID error")
	}

	// optional field (optPresentFlag index: 0)
	if PRSTRPItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.RequestedDLPRSTransmissionCharacteristics = new(RequestedDLPRSTransmissionCharacteristics)
		err = x.RequestedDLPRSTransmissionCharacteristics.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RequestedDLPRSTransmissionCharacteristics error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PRSTRPItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PRSTransmissionOffInformation = new(PRSTransmissionOffInformation)
		err = x.PRSTransmissionOffInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSTransmissionOffInformation error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if PRSTRPItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSTRPItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
