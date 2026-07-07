package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPPositionDirect struct {
	Accuracy     *TRPPositionDirectAccuracy                         // valueLB:0,valueUB:2
	IEExtensions *ProtocolExtensionContainerTRPPositionDirectExtIEs // optional
}

func (x *TRPPositionDirect) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPositionDirectOptPresentFlag := []bool{}
	// mandatory field
	if x.Accuracy == nil {
		return errors.Errorf("Accuracy is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPPositionDirectOptPresentFlag = append(TRPPositionDirectOptPresentFlag, true)
	} else {
		TRPPositionDirectOptPresentFlag = append(TRPPositionDirectOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPPositionDirectOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Accuracy.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Accuracy marshal failed")
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

func (x *TRPPositionDirect) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPositionDirectOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&TRPPositionDirectOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Accuracy = new(TRPPositionDirectAccuracy)
	err = x.Accuracy.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Accuracy error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPPositionDirectOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPPositionDirectExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
