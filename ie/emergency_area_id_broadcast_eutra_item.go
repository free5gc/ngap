package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EmergencyAreaIDBroadcastEUTRAItem struct {
	EmergencyAreaID          *EmergencyAreaID
	CompletedCellsInEAIEUTRA *CompletedCellsInEAIEUTRA
	IEExtensions             *ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs // optional
}

func (x *EmergencyAreaIDBroadcastEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.EmergencyAreaID == nil {
		return errors.Errorf("EmergencyAreaID is missing")
	}
	// mandatory field
	if x.CompletedCellsInEAIEUTRA == nil {
		return errors.Errorf("CompletedCellsInEAIEUTRA is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag = append(EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag, true)
	} else {
		EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag = append(EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag, true)
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
	err = x.CompletedCellsInEAIEUTRA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CompletedCellsInEAIEUTRA marshal failed")
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

func (x *EmergencyAreaIDBroadcastEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag, true)
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
	x.CompletedCellsInEAIEUTRA = new(CompletedCellsInEAIEUTRA)
	err = x.CompletedCellsInEAIEUTRA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CompletedCellsInEAIEUTRA error")
	}

	// optional field (optPresentFlag index: 0)
	if EmergencyAreaIDBroadcastEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEmergencyAreaIDBroadcastEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
