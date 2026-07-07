package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type DLPRSResourceIDItem struct {
	DlPRSResourceID *PRSResourceID
	IEExtensions    *ProtocolExtensionContainerDLPRSResourceItemExtIEs // optional
}

func (x *DLPRSResourceIDItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceIDItemOptPresentFlag := []bool{}
	// mandatory field
	if x.DlPRSResourceID == nil {
		return errors.Errorf("DlPRSResourceID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		DLPRSResourceIDItemOptPresentFlag = append(DLPRSResourceIDItemOptPresentFlag, true)
	} else {
		DLPRSResourceIDItemOptPresentFlag = append(DLPRSResourceIDItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceIDItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DlPRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DlPRSResourceID marshal failed")
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

func (x *DLPRSResourceIDItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceIDItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceIDItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DlPRSResourceID = new(PRSResourceID)
	err = x.DlPRSResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DlPRSResourceID error")
	}

	// optional field (optPresentFlag index: 0)
	if DLPRSResourceIDItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDLPRSResourceItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
