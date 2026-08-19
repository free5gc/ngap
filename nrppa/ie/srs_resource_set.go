package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSResourceSet struct {
	SRSResourceSetID  *int64 // valueLB:0,valueUB:15
	SRSResourceIDList *SRSResourceIDList
	ResourceSetType   *ResourceSetType                                // valueLB:0,valueUB:3
	IEExtensions      *ProtocolExtensionContainerSRSResourceSetExtIEs // optional
}

func (x *SRSResourceSet) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSResourceSetOptPresentFlag := []bool{}
	// mandatory field
	if x.SRSResourceSetID == nil {
		return errors.Errorf("SRSResourceSetID is missing")
	}
	// mandatory field
	if x.SRSResourceIDList == nil {
		return errors.Errorf("SRSResourceIDList is missing")
	}
	// mandatory field
	if x.ResourceSetType == nil {
		return errors.Errorf("ResourceSetType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		SRSResourceSetOptPresentFlag = append(SRSResourceSetOptPresentFlag, true)
	} else {
		SRSResourceSetOptPresentFlag = append(SRSResourceSetOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSResourceSetOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 15
	err = pd.WriteInteger(*(x.SRSResourceSetID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.SRSResourceIDList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SRSResourceIDList marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ResourceSetType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ResourceSetType marshal failed")
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

func (x *SRSResourceSet) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSResourceSetOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&SRSResourceSetOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 15
	x.SRSResourceSetID = new(int64)
	*(x.SRSResourceSetID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SRSResourceIDList = new(SRSResourceIDList)
	err = x.SRSResourceIDList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SRSResourceIDList error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ResourceSetType = new(ResourceSetType)
	err = x.ResourceSetType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ResourceSetType error")
	}

	// optional field (optPresentFlag index: 0)
	if SRSResourceSetOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSResourceSetExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
