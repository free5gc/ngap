package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UnavailableGUAMIItem struct {
	GUAMI                        *GUAMI                                                // valueExt
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval                         // valueExt,valueLB:0,valueUB:0,optional
	BackupAMFName                *AMFName                                              // sizeExt,sizeLB:1,sizeUB:150,optional
	IEExtensions                 *ProtocolExtensionContainerUnavailableGUAMIItemExtIEs // optional
}

func (x *UnavailableGUAMIItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UnavailableGUAMIItemOptPresentFlag := []bool{}
	// mandatory field
	if x.GUAMI == nil {
		return errors.Errorf("GUAMI is missing")
	}
	// optional field
	if x.TimerApproachForGUAMIRemoval != nil {
		UnavailableGUAMIItemOptPresentFlag = append(UnavailableGUAMIItemOptPresentFlag, true)
	} else {
		UnavailableGUAMIItemOptPresentFlag = append(UnavailableGUAMIItemOptPresentFlag, false)
	}
	// optional field
	if x.BackupAMFName != nil {
		UnavailableGUAMIItemOptPresentFlag = append(UnavailableGUAMIItemOptPresentFlag, true)
	} else {
		UnavailableGUAMIItemOptPresentFlag = append(UnavailableGUAMIItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UnavailableGUAMIItemOptPresentFlag = append(UnavailableGUAMIItemOptPresentFlag, true)
	} else {
		UnavailableGUAMIItemOptPresentFlag = append(UnavailableGUAMIItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UnavailableGUAMIItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.GUAMI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "GUAMI marshal failed")
	}

	// optional field
	if x.TimerApproachForGUAMIRemoval != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TimerApproachForGUAMIRemoval.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TimerApproachForGUAMIRemoval marshal failed")
		}
	}

	// optional field
	if x.BackupAMFName != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.BackupAMFName.Write(pd)
		if err != nil {
			return errors.Wrap(err, "BackupAMFName marshal failed")
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

func (x *UnavailableGUAMIItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UnavailableGUAMIItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&UnavailableGUAMIItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.GUAMI = new(GUAMI)
	err = x.GUAMI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode GUAMI error")
	}

	// optional field (optPresentFlag index: 0)
	if UnavailableGUAMIItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TimerApproachForGUAMIRemoval = new(TimerApproachForGUAMIRemoval)
		err = x.TimerApproachForGUAMIRemoval.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TimerApproachForGUAMIRemoval error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UnavailableGUAMIItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.BackupAMFName = new(AMFName)
		err = x.BackupAMFName.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BackupAMFName error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if UnavailableGUAMIItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUnavailableGUAMIItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
