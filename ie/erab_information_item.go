package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ERABInformationItem struct {
	ERABID       *ERABID
	DLForwarding *DLForwarding                                        // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions *ProtocolExtensionContainerERABInformationItemExtIEs // optional
}

func (x *ERABInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ERABInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.ERABID == nil {
		return errors.Errorf("ERABID is missing")
	}
	// optional field
	if x.DLForwarding != nil {
		ERABInformationItemOptPresentFlag = append(ERABInformationItemOptPresentFlag, true)
	} else {
		ERABInformationItemOptPresentFlag = append(ERABInformationItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ERABInformationItemOptPresentFlag = append(ERABInformationItemOptPresentFlag, true)
	} else {
		ERABInformationItemOptPresentFlag = append(ERABInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ERABInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ERABID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ERABID marshal failed")
	}

	// optional field
	if x.DLForwarding != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLForwarding.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLForwarding marshal failed")
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

func (x *ERABInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ERABInformationItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&ERABInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ERABID = new(ERABID)
	err = x.ERABID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ERABID error")
	}

	// optional field (optPresentFlag index: 0)
	if ERABInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLForwarding = new(DLForwarding)
		err = x.DLForwarding.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLForwarding error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ERABInformationItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerERABInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
