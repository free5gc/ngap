package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSResourceItem struct {
	PRSResourceID        *PRSResourceID
	SequenceID           *int64                                           // valueLB:0,valueUB:4095
	REOffset             *int64                                           // valueExt,valueLB:0,valueUB:11
	ResourceSlotOffset   *int64                                           // valueLB:0,valueUB:511
	ResourceSymbolOffset *int64                                           // valueLB:0,valueUB:12
	QCLInfo              *PRSResourceQCLInfo                              // valueLB:0,valueUB:2,optional
	IEExtensions         *ProtocolExtensionContainerPRSResourceItemExtIEs // optional
}

func (x *PRSResourceItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSResourceItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSResourceID == nil {
		return errors.Errorf("PRSResourceID is missing")
	}
	// mandatory field
	if x.SequenceID == nil {
		return errors.Errorf("SequenceID is missing")
	}
	// mandatory field
	if x.REOffset == nil {
		return errors.Errorf("REOffset is missing")
	}
	// mandatory field
	if x.ResourceSlotOffset == nil {
		return errors.Errorf("ResourceSlotOffset is missing")
	}
	// mandatory field
	if x.ResourceSymbolOffset == nil {
		return errors.Errorf("ResourceSymbolOffset is missing")
	}
	// optional field
	if x.QCLInfo != nil {
		PRSResourceItemOptPresentFlag = append(PRSResourceItemOptPresentFlag, true)
	} else {
		PRSResourceItemOptPresentFlag = append(PRSResourceItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		PRSResourceItemOptPresentFlag = append(PRSResourceItemOptPresentFlag, true)
	} else {
		PRSResourceItemOptPresentFlag = append(PRSResourceItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSResourceItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceID marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 4095
	err = pd.WriteInteger(*(x.SequenceID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 11
	err = pd.WriteInteger(*(x.REOffset), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 511
	err = pd.WriteInteger(*(x.ResourceSlotOffset), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 12
	err = pd.WriteInteger(*(x.ResourceSymbolOffset), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// optional field
	if x.QCLInfo != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.QCLInfo.Write(pd)
		if err != nil {
			return errors.Wrap(err, "QCLInfo marshal failed")
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

func (x *PRSResourceItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSResourceItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&PRSResourceItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceID = new(PRSResourceID)
	err = x.PRSResourceID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceID error")
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 4095
	x.SequenceID = new(int64)
	*(x.SequenceID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 11
	x.REOffset = new(int64)
	*(x.REOffset), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 511
	x.ResourceSlotOffset = new(int64)
	*(x.ResourceSlotOffset), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 12
	x.ResourceSymbolOffset = new(int64)
	*(x.ResourceSymbolOffset), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 0)
	if PRSResourceItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.QCLInfo = new(PRSResourceQCLInfo)
		err = x.QCLInfo.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode QCLInfo error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if PRSResourceItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSResourceItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
