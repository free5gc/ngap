package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ServedGUAMIItem struct {
	GUAMI         *GUAMI                                           // valueExt
	BackupAMFName *AMFName                                         // sizeExt,sizeLB:1,sizeUB:150,optional
	IEExtensions  *ProtocolExtensionContainerServedGUAMIItemExtIEs // optional
}

func (x *ServedGUAMIItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ServedGUAMIItemOptPresentFlag := []bool{}
	// mandatory field
	if x.GUAMI == nil {
		return errors.Errorf("GUAMI is missing")
	}
	// optional field
	if x.BackupAMFName != nil {
		ServedGUAMIItemOptPresentFlag = append(ServedGUAMIItemOptPresentFlag, true)
	} else {
		ServedGUAMIItemOptPresentFlag = append(ServedGUAMIItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ServedGUAMIItemOptPresentFlag = append(ServedGUAMIItemOptPresentFlag, true)
	} else {
		ServedGUAMIItemOptPresentFlag = append(ServedGUAMIItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ServedGUAMIItemOptPresentFlag, true)
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

func (x *ServedGUAMIItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ServedGUAMIItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ServedGUAMIItemOptPresentFlag, true)
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
	if ServedGUAMIItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.BackupAMFName = new(AMFName)
		err = x.BackupAMFName.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BackupAMFName error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ServedGUAMIItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerServedGUAMIItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
