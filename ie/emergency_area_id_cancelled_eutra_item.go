package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type EmergencyAreaIDCancelledEUTRAItem struct {
	EmergencyAreaID          *EmergencyAreaID
	CancelledCellsInEAIEUTRA *CancelledCellsInEAIEUTRA
	IEExtensions             *ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs // optional
}

func (x *EmergencyAreaIDCancelledEUTRAItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	EmergencyAreaIDCancelledEUTRAItemOptPresentFlag := []bool{}
	// mandatory field
	if x.EmergencyAreaID == nil {
		return errors.Errorf("EmergencyAreaID is missing")
	}
	// mandatory field
	if x.CancelledCellsInEAIEUTRA == nil {
		return errors.Errorf("CancelledCellsInEAIEUTRA is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		EmergencyAreaIDCancelledEUTRAItemOptPresentFlag = append(EmergencyAreaIDCancelledEUTRAItemOptPresentFlag, true)
	} else {
		EmergencyAreaIDCancelledEUTRAItemOptPresentFlag = append(EmergencyAreaIDCancelledEUTRAItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(EmergencyAreaIDCancelledEUTRAItemOptPresentFlag, true)
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
	err = x.CancelledCellsInEAIEUTRA.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CancelledCellsInEAIEUTRA marshal failed")
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

func (x *EmergencyAreaIDCancelledEUTRAItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	EmergencyAreaIDCancelledEUTRAItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&EmergencyAreaIDCancelledEUTRAItemOptPresentFlag, true)
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
	x.CancelledCellsInEAIEUTRA = new(CancelledCellsInEAIEUTRA)
	err = x.CancelledCellsInEAIEUTRA.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CancelledCellsInEAIEUTRA error")
	}

	// optional field (optPresentFlag index: 0)
	if EmergencyAreaIDCancelledEUTRAItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerEmergencyAreaIDCancelledEUTRAItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
