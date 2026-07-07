package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type LastVisitedPSCellInformation struct {
	PSCellID     *NGRANCGI                                                     // valueLB:0,valueUB:2,optional
	TimeStay     *int64                                                        // valueLB:0,valueUB:40950
	IEExtensions *ProtocolExtensionContainerLastVisitedPSCellInformationExtIEs // optional
}

func (x *LastVisitedPSCellInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	LastVisitedPSCellInformationOptPresentFlag := []bool{}
	// optional field
	if x.PSCellID != nil {
		LastVisitedPSCellInformationOptPresentFlag = append(LastVisitedPSCellInformationOptPresentFlag, true)
	} else {
		LastVisitedPSCellInformationOptPresentFlag = append(LastVisitedPSCellInformationOptPresentFlag, false)
	}
	// mandatory field
	if x.TimeStay == nil {
		return errors.Errorf("TimeStay is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		LastVisitedPSCellInformationOptPresentFlag = append(LastVisitedPSCellInformationOptPresentFlag, true)
	} else {
		LastVisitedPSCellInformationOptPresentFlag = append(LastVisitedPSCellInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(LastVisitedPSCellInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PSCellID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PSCellID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PSCellID marshal failed")
		}
	}

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 40950
	err = pd.WriteInteger(*(x.TimeStay), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
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

func (x *LastVisitedPSCellInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	LastVisitedPSCellInformationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&LastVisitedPSCellInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if LastVisitedPSCellInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PSCellID = new(NGRANCGI)
		err = x.PSCellID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PSCellID error")
		}
	}

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 40950
	x.TimeStay = new(int64)
	*(x.TimeStay), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// optional field (optPresentFlag index: 1)
	if LastVisitedPSCellInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerLastVisitedPSCellInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
