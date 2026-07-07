package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	CNTypeRestrictionsForEquivalentItemCnTypePresentEpcForbidden    aper.Enumerated = 0
	CNTypeRestrictionsForEquivalentItemCnTypePresentFiveGCForbidden aper.Enumerated = 1
)

type CNTypeRestrictionsForEquivalentItem struct {
	PlmnIdentity *PLMNIdentity
	CnType       *aper.Enumerated                                                     // valueExt,valueLB:0,valueUB:1
	IEExtensions *ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs // optional
}

func (x *CNTypeRestrictionsForEquivalentItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	CNTypeRestrictionsForEquivalentItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PlmnIdentity == nil {
		return errors.Errorf("PlmnIdentity is missing")
	}
	// mandatory field
	if x.CnType == nil {
		return errors.Errorf("CnType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		CNTypeRestrictionsForEquivalentItemOptPresentFlag = append(CNTypeRestrictionsForEquivalentItemOptPresentFlag, true)
	} else {
		CNTypeRestrictionsForEquivalentItemOptPresentFlag = append(CNTypeRestrictionsForEquivalentItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(CNTypeRestrictionsForEquivalentItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PlmnIdentity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PlmnIdentity marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.CnType), true, vLb, vUb)
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

func (x *CNTypeRestrictionsForEquivalentItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	CNTypeRestrictionsForEquivalentItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&CNTypeRestrictionsForEquivalentItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PlmnIdentity = new(PLMNIdentity)
	err = x.PlmnIdentity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PlmnIdentity error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.CnType = new(aper.Enumerated)
	*(x.CnType), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if CNTypeRestrictionsForEquivalentItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerCNTypeRestrictionsForEquivalentItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
