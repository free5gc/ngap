package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LastVisitedCellItem struct {
	LastVisitedCellInformation *LastVisitedCellInformation                          // valueLB:0,valueUB:4
	IEExtensions               *ProtocolExtensionContainerLastVisitedCellItemExtIEs // optional
}

func (x *LastVisitedCellItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LastVisitedCellItemOptPresentFlag := []bool{}
	// mandatory field
	if x.LastVisitedCellInformation == nil {
		return errors.Errorf("LastVisitedCellInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		LastVisitedCellItemOptPresentFlag = append(LastVisitedCellItemOptPresentFlag, true)
	} else {
		LastVisitedCellItemOptPresentFlag = append(LastVisitedCellItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LastVisitedCellItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.LastVisitedCellInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "LastVisitedCellInformation marshal failed")
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

func (x *LastVisitedCellItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LastVisitedCellItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&LastVisitedCellItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.LastVisitedCellInformation = new(LastVisitedCellInformation)
	err = x.LastVisitedCellInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode LastVisitedCellInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if LastVisitedCellItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLastVisitedCellItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
