package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EmergencyAreaIDCancelledNRItem struct {
	EmergencyAreaID       *EmergencyAreaID
	CancelledCellsInEAINR *CancelledCellsInEAINR
	IEExtensions          *ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs // optional
}

func (x *EmergencyAreaIDCancelledNRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EmergencyAreaIDCancelledNRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.EmergencyAreaID == nil {
		return errors.Errorf("EmergencyAreaID is missing")
	}
	// mandatory field
	if x.CancelledCellsInEAINR == nil {
		return errors.Errorf("CancelledCellsInEAINR is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EmergencyAreaIDCancelledNRItemOptPresentFlag = append(EmergencyAreaIDCancelledNRItemOptPresentFlag, true)
	} else {
		EmergencyAreaIDCancelledNRItemOptPresentFlag = append(EmergencyAreaIDCancelledNRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EmergencyAreaIDCancelledNRItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.EmergencyAreaID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "EmergencyAreaID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.CancelledCellsInEAINR.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CancelledCellsInEAINR marshal failed")
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

func (x *EmergencyAreaIDCancelledNRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EmergencyAreaIDCancelledNRItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EmergencyAreaIDCancelledNRItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.EmergencyAreaID = new(EmergencyAreaID)
	err = x.EmergencyAreaID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode EmergencyAreaID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CancelledCellsInEAINR = new(CancelledCellsInEAINR)
	err = x.CancelledCellsInEAINR.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CancelledCellsInEAINR error")
	}

	// optional field (optPresentFlag index: 0)
	if EmergencyAreaIDCancelledNRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEmergencyAreaIDCancelledNRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
