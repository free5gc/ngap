package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type DLPRS struct {
	Prsid              *int64 // valueLB:0,valueUB:255
	DlPRSResourceSetID *PRSResourceSetID
	DlPRSResourceID    *PRSResourceID                         // optional
	IEExtensions       *ProtocolExtensionContainerDLPRSExtIEs // optional
}

func (x *DLPRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DLPRSOptPresentFlag := []bool{}
	// mandatory field
	if x.Prsid == nil {
		return errors.Errorf("Prsid is missing")
	}
	// mandatory field
	if x.DlPRSResourceSetID == nil {
		return errors.Errorf("DlPRSResourceSetID is missing")
	}
	// optional field
	if x.DlPRSResourceID != nil {
		DLPRSOptPresentFlag = append(DLPRSOptPresentFlag, true)
	} else {
		DLPRSOptPresentFlag = append(DLPRSOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		DLPRSOptPresentFlag = append(DLPRSOptPresentFlag, true)
	} else {
		DLPRSOptPresentFlag = append(DLPRSOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DLPRSOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 255
	err = pd.WriteInteger(*(x.Prsid), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.DlPRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DlPRSResourceSetID marshal failed")
	}

	// optional field
	if x.DlPRSResourceID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DlPRSResourceID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DlPRSResourceID marshal failed")
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

func (x *DLPRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DLPRSOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&DLPRSOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 255
	x.Prsid = new(int64)
	*(x.Prsid), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DlPRSResourceSetID = new(PRSResourceSetID)
	err = x.DlPRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DlPRSResourceSetID error")
	}

	// optional field (optPresentFlag index: 0)
	if DLPRSOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DlPRSResourceID = new(PRSResourceID)
		err = x.DlPRSResourceID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DlPRSResourceID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if DLPRSOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDLPRSExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
