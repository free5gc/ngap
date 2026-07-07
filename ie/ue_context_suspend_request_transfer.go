package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEContextSuspendRequestTransfer struct {
	SuspendIndicator *SuspendIndicator                                                // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions     *ProtocolExtensionContainerUEContextSuspendRequestTransferExtIEs // optional
}

func (x *UEContextSuspendRequestTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEContextSuspendRequestTransferOptPresentFlag := []bool{}
	// optional field
	if x.SuspendIndicator != nil {
		UEContextSuspendRequestTransferOptPresentFlag = append(UEContextSuspendRequestTransferOptPresentFlag, true)
	} else {
		UEContextSuspendRequestTransferOptPresentFlag = append(UEContextSuspendRequestTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEContextSuspendRequestTransferOptPresentFlag = append(UEContextSuspendRequestTransferOptPresentFlag, true)
	} else {
		UEContextSuspendRequestTransferOptPresentFlag = append(UEContextSuspendRequestTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEContextSuspendRequestTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.SuspendIndicator != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SuspendIndicator.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SuspendIndicator marshal failed")
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

func (x *UEContextSuspendRequestTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEContextSuspendRequestTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UEContextSuspendRequestTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if UEContextSuspendRequestTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SuspendIndicator = new(SuspendIndicator)
		err = x.SuspendIndicator.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SuspendIndicator error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UEContextSuspendRequestTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEContextSuspendRequestTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
