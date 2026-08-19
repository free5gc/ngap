package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type UETxTEGAssociationItem struct {
	UETxTEGID            *int64 // valueLB:0,valueUB:7
	PosSRSResourceIDList *PosSRSResourceIDList
	TimeStamp            *TimeStamp                                              // valueExt
	CarrierFreq          *CarrierFreq                                            // valueExt,optional
	IEExtensions         *ProtocolExtensionContainerUETxTEGAssociationItemExtIEs // optional
}

func (x *UETxTEGAssociationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UETxTEGAssociationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.UETxTEGID == nil {
		return errors.Errorf("UETxTEGID is missing")
	}
	// mandatory field
	if x.PosSRSResourceIDList == nil {
		return errors.Errorf("PosSRSResourceIDList is missing")
	}
	// mandatory field
	if x.TimeStamp == nil {
		return errors.Errorf("TimeStamp is missing")
	}
	// optional field
	if x.CarrierFreq != nil {
		UETxTEGAssociationItemOptPresentFlag = append(UETxTEGAssociationItemOptPresentFlag, true)
	} else {
		UETxTEGAssociationItemOptPresentFlag = append(UETxTEGAssociationItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UETxTEGAssociationItemOptPresentFlag = append(UETxTEGAssociationItemOptPresentFlag, true)
	} else {
		UETxTEGAssociationItemOptPresentFlag = append(UETxTEGAssociationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UETxTEGAssociationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 7
	err = pd.WriteInteger(*(x.UETxTEGID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PosSRSResourceIDList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PosSRSResourceIDList marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TimeStamp.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TimeStamp marshal failed")
	}

	// optional field
	if x.CarrierFreq != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CarrierFreq.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CarrierFreq marshal failed")
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

func (x *UETxTEGAssociationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UETxTEGAssociationItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&UETxTEGAssociationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 7
	x.UETxTEGID = new(int64)
	*(x.UETxTEGID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PosSRSResourceIDList = new(PosSRSResourceIDList)
	err = x.PosSRSResourceIDList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PosSRSResourceIDList error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TimeStamp = new(TimeStamp)
	err = x.TimeStamp.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TimeStamp error")
	}

	// optional field (optPresentFlag index: 0)
	if UETxTEGAssociationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CarrierFreq = new(CarrierFreq)
		err = x.CarrierFreq.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CarrierFreq error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UETxTEGAssociationItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUETxTEGAssociationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
