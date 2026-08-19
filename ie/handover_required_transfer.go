package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability                         // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions                     *ProtocolExtensionContainerHandoverRequiredTransferExtIEs // optional
}

func (x *HandoverRequiredTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	HandoverRequiredTransferOptPresentFlag := []bool{}
	// optional field
	if x.DirectForwardingPathAvailability != nil {
		HandoverRequiredTransferOptPresentFlag = append(HandoverRequiredTransferOptPresentFlag, true)
	} else {
		HandoverRequiredTransferOptPresentFlag = append(HandoverRequiredTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		HandoverRequiredTransferOptPresentFlag = append(HandoverRequiredTransferOptPresentFlag, true)
	} else {
		HandoverRequiredTransferOptPresentFlag = append(HandoverRequiredTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(HandoverRequiredTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.DirectForwardingPathAvailability != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DirectForwardingPathAvailability.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DirectForwardingPathAvailability marshal failed")
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

func (x *HandoverRequiredTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	HandoverRequiredTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&HandoverRequiredTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if HandoverRequiredTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DirectForwardingPathAvailability = new(DirectForwardingPathAvailability)
		err = x.DirectForwardingPathAvailability.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DirectForwardingPathAvailability error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if HandoverRequiredTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerHandoverRequiredTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
