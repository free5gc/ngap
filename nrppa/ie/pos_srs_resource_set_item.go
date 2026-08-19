package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PosSRSResourceSetItem struct {
	PossrsResourceSetID        *int64 // valueLB:0,valueUB:15
	PossRSResourceIDPerSetList *PosSRSResourceIDPerSetList
	PosresourceSetType         *PosResourceSetType                                    // valueLB:0,valueUB:3
	IEExtensions               *ProtocolExtensionContainerPosSRSResourceSetItemExtIEs // optional
}

func (x *PosSRSResourceSetItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PosSRSResourceSetItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PossrsResourceSetID == nil {
		return errors.Errorf("PossrsResourceSetID is missing")
	}
	// mandatory field
	if x.PossRSResourceIDPerSetList == nil {
		return errors.Errorf("PossRSResourceIDPerSetList is missing")
	}
	// mandatory field
	if x.PosresourceSetType == nil {
		return errors.Errorf("PosresourceSetType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PosSRSResourceSetItemOptPresentFlag = append(PosSRSResourceSetItemOptPresentFlag, true)
	} else {
		PosSRSResourceSetItemOptPresentFlag = append(PosSRSResourceSetItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PosSRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 15
	err = pd.WriteInteger(*(x.PossrsResourceSetID), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PossRSResourceIDPerSetList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PossRSResourceIDPerSetList marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PosresourceSetType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PosresourceSetType marshal failed")
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

func (x *PosSRSResourceSetItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PosSRSResourceSetItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PosSRSResourceSetItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 15
	x.PossrsResourceSetID = new(int64)
	*(x.PossrsResourceSetID), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PossRSResourceIDPerSetList = new(PosSRSResourceIDPerSetList)
	err = x.PossRSResourceIDPerSetList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PossRSResourceIDPerSetList error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PosresourceSetType = new(PosResourceSetType)
	err = x.PosresourceSetType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PosresourceSetType error")
	}

	// optional field (optPresentFlag index: 0)
	if PosSRSResourceSetItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPosSRSResourceSetItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
