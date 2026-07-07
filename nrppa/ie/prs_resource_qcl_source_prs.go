package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSResourceQCLSourcePRS struct {
	QCLSourcePRSResourceSetID *PRSResourceSetID
	QCLSourcePRSResourceID    *PRSResourceID                                           // optional
	IEExtensions              *ProtocolExtensionContainerPRSResourceQCLSourcePRSExtIEs // optional
}

func (x *PRSResourceQCLSourcePRS) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceQCLSourcePRSOptPresentFlag := []bool{}
	// mandatory field
	if x.QCLSourcePRSResourceSetID == nil {
		return errors.Errorf("QCLSourcePRSResourceSetID is missing")
	}
	// optional field
	if x.QCLSourcePRSResourceID != nil {
		PRSResourceQCLSourcePRSOptPresentFlag = append(PRSResourceQCLSourcePRSOptPresentFlag, true)
	} else {
		PRSResourceQCLSourcePRSOptPresentFlag = append(PRSResourceQCLSourcePRSOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PRSResourceQCLSourcePRSOptPresentFlag = append(PRSResourceQCLSourcePRSOptPresentFlag, true)
	} else {
		PRSResourceQCLSourcePRSOptPresentFlag = append(PRSResourceQCLSourcePRSOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceQCLSourcePRSOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.QCLSourcePRSResourceSetID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "QCLSourcePRSResourceSetID marshal failed")
	}

	// optional field
	if x.QCLSourcePRSResourceID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QCLSourcePRSResourceID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QCLSourcePRSResourceID marshal failed")
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

func (x *PRSResourceQCLSourcePRS) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceQCLSourcePRSOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceQCLSourcePRSOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.QCLSourcePRSResourceSetID = new(PRSResourceSetID)
	err = x.QCLSourcePRSResourceSetID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode QCLSourcePRSResourceSetID error")
	}

	// optional field (optPresentFlag index: 0)
	if PRSResourceQCLSourcePRSOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QCLSourcePRSResourceID = new(PRSResourceID)
		err = x.QCLSourcePRSResourceID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QCLSourcePRSResourceID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PRSResourceQCLSourcePRSOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSResourceQCLSourcePRSExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
