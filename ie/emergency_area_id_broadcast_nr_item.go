package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       *EmergencyAreaID
	CompletedCellsInEAINR *CompletedCellsInEAINR
	IEExtensions          *ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs // optional
}

func (x *EmergencyAreaIDBroadcastNRItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EmergencyAreaIDBroadcastNRItemOptPresentFlag := []bool{}
	// mandatory field
	if x.EmergencyAreaID == nil {
		return errors.Errorf("EmergencyAreaID is missing")
	}
	// mandatory field
	if x.CompletedCellsInEAINR == nil {
		return errors.Errorf("CompletedCellsInEAINR is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EmergencyAreaIDBroadcastNRItemOptPresentFlag = append(EmergencyAreaIDBroadcastNRItemOptPresentFlag, true)
	} else {
		EmergencyAreaIDBroadcastNRItemOptPresentFlag = append(EmergencyAreaIDBroadcastNRItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EmergencyAreaIDBroadcastNRItemOptPresentFlag, true)
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
	err = x.CompletedCellsInEAINR.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CompletedCellsInEAINR marshal failed")
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

func (x *EmergencyAreaIDBroadcastNRItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EmergencyAreaIDBroadcastNRItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EmergencyAreaIDBroadcastNRItemOptPresentFlag, true)
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
	x.CompletedCellsInEAINR = new(CompletedCellsInEAINR)
	err = x.CompletedCellsInEAINR.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CompletedCellsInEAINR error")
	}

	// optional field (optPresentFlag index: 0)
	if EmergencyAreaIDBroadcastNRItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
