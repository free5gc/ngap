package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type IntersystemCellActivationRequest struct {
	ActivationID        *int64 // valueExt,valueLB:0,valueUB:16384
	CellsToActivateList *CellsToActivateList
	IEExtensions        *ProtocolExtensionContainerIntersystemCellActivationRequestExtIEs // optional
}

func (x *IntersystemCellActivationRequest) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	IntersystemCellActivationRequestOptPresentFlag := []bool{}
	// mandatory field
	if x.ActivationID == nil {
		return errors.Errorf("ActivationID is missing")
	}
	// mandatory field
	if x.CellsToActivateList == nil {
		return errors.Errorf("CellsToActivateList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		IntersystemCellActivationRequestOptPresentFlag = append(IntersystemCellActivationRequestOptPresentFlag, true)
	} else {
		IntersystemCellActivationRequestOptPresentFlag = append(IntersystemCellActivationRequestOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(IntersystemCellActivationRequestOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 16384
	err = pd.WriteInteger(*(x.ActivationID), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.CellsToActivateList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "CellsToActivateList marshal failed")
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

func (x *IntersystemCellActivationRequest) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	IntersystemCellActivationRequestOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&IntersystemCellActivationRequestOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 16384
	x.ActivationID = new(int64)
	*(x.ActivationID), err = pd.ReadInteger(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.CellsToActivateList = new(CellsToActivateList)
	err = x.CellsToActivateList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode CellsToActivateList error")
	}

	// optional field (optPresentFlag index: 0)
	if IntersystemCellActivationRequestOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerIntersystemCellActivationRequestExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
