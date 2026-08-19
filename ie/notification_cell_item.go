package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NotificationCellItemNotifyFlagPresentActivated   aper.Enumerated = 0
	NotificationCellItemNotifyFlagPresentDeactivated aper.Enumerated = 1
)

type NotificationCellItem struct {
	NGRANCGI     *NGRANCGI                                             // valueLB:0,valueUB:2
	NotifyFlag   *aper.Enumerated                                      // valueExt,valueLB:0,valueUB:1
	IEExtensions *ProtocolExtensionContainerNotificationCellItemExtIEs // optional
}

func (x *NotificationCellItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NotificationCellItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NGRANCGI == nil {
		return errors.Errorf("NGRANCGI is missing")
	}
	// mandatory field
	if x.NotifyFlag == nil {
		return errors.Errorf("NotifyFlag is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NotificationCellItemOptPresentFlag = append(NotificationCellItemOptPresentFlag, true)
	} else {
		NotificationCellItemOptPresentFlag = append(NotificationCellItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NotificationCellItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NGRANCGI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NGRANCGI marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.NotifyFlag), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *NotificationCellItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NotificationCellItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NotificationCellItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NGRANCGI = new(NGRANCGI)
	err = x.NGRANCGI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NGRANCGI error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.NotifyFlag = new(aper.Enumerated)
	*(x.NotifyFlag), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if NotificationCellItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNotificationCellItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
