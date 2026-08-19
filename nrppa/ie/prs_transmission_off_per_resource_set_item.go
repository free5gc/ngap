package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSTransmissionOffPerResourceSetItem struct {
	PRSResourceSetID *PRSResourceSetID
	IEExtensions     *ProtocolExtensionContainerPRSTransmissionOffPerResourceSetItemExtIEs // optional
}

func (x *PRSTransmissionOffPerResourceSetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionOffPerResourceSetItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSResourceSetID == nil {
		return errors.Errorf("PRSResourceSetID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSTransmissionOffPerResourceSetItemOptPresentFlag = append(PRSTransmissionOffPerResourceSetItemOptPresentFlag, true)
	} else {
		PRSTransmissionOffPerResourceSetItemOptPresentFlag = append(PRSTransmissionOffPerResourceSetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionOffPerResourceSetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceSetID marshal failed")
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

func (x *PRSTransmissionOffPerResourceSetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionOffPerResourceSetItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionOffPerResourceSetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceSetID = new(PRSResourceSetID)
	err = x.PRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceSetID error")
	}

	// optional field (optPresentFlag index: 0)
	if PRSTransmissionOffPerResourceSetItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSTransmissionOffPerResourceSetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
