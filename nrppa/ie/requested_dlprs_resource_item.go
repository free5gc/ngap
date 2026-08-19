package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type RequestedDLPRSResourceItem struct {
	QCLInfo      *PRSResourceQCLInfo                                         // valueLB:0,valueUB:2,optional
	IEExtensions *ProtocolExtensionContainerRequestedDLPRSResourceItemExtIEs // optional
}

func (x *RequestedDLPRSResourceItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	RequestedDLPRSResourceItemOptPresentFlag := []bool{}
	// optional field
	if x.QCLInfo != nil {
		RequestedDLPRSResourceItemOptPresentFlag = append(RequestedDLPRSResourceItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceItemOptPresentFlag = append(RequestedDLPRSResourceItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		RequestedDLPRSResourceItemOptPresentFlag = append(RequestedDLPRSResourceItemOptPresentFlag, true)
	} else {
		RequestedDLPRSResourceItemOptPresentFlag = append(RequestedDLPRSResourceItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(RequestedDLPRSResourceItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.QCLInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QCLInfo.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QCLInfo marshal failed")
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

func (x *RequestedDLPRSResourceItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	RequestedDLPRSResourceItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&RequestedDLPRSResourceItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if RequestedDLPRSResourceItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QCLInfo = new(PRSResourceQCLInfo)
		err = x.QCLInfo.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QCLInfo error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if RequestedDLPRSResourceItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerRequestedDLPRSResourceItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
