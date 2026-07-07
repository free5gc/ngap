package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type DLPRSResourceARP struct {
	DlPRSResourceID          *PRSResourceID
	DLPRSResourceARPLocation *DLPRSResourceARPLocation                         // valueLB:0,valueUB:2
	IEExtensions             *ProtocolExtensionContainerDLPRSResourceARPExtIEs // optional
}

func (x *DLPRSResourceARP) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSResourceARPOptPresentFlag := []bool{}
	// mandatory field
	if x.DlPRSResourceID == nil {
		return errors.Errorf("DlPRSResourceID is missing")
	}
	// mandatory field
	if x.DLPRSResourceARPLocation == nil {
		return errors.Errorf("DLPRSResourceARPLocation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		DLPRSResourceARPOptPresentFlag = append(DLPRSResourceARPOptPresentFlag, true)
	} else {
		DLPRSResourceARPOptPresentFlag = append(DLPRSResourceARPOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSResourceARPOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DlPRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DlPRSResourceID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DLPRSResourceARPLocation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DLPRSResourceARPLocation marshal failed")
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

func (x *DLPRSResourceARP) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSResourceARPOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&DLPRSResourceARPOptPresentFlag, true)
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

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DLPRSResourceARPLocation = new(DLPRSResourceARPLocation)
	err = x.DLPRSResourceARPLocation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DLPRSResourceARPLocation error")
	}

	// optional field (optPresentFlag index: 0)
	if DLPRSResourceARPOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDLPRSResourceARPExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
