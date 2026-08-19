package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ARPLocationInformationItem struct {
	ARPID           *ARPID
	ARPLocationType *ARPLocationType                                        // valueLB:0,valueUB:2
	IEExtensions    *ProtocolExtensionContainerARPLocationInformationExtIEs // optional
}

func (x *ARPLocationInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ARPLocationInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.ARPID == nil {
		return errors.Errorf("ARPID is missing")
	}
	// mandatory field
	if x.ARPLocationType == nil {
		return errors.Errorf("ARPLocationType is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		ARPLocationInformationItemOptPresentFlag = append(ARPLocationInformationItemOptPresentFlag, true)
	} else {
		ARPLocationInformationItemOptPresentFlag = append(ARPLocationInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ARPLocationInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.ARPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ARPID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ARPLocationType.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ARPLocationType marshal failed")
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

func (x *ARPLocationInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ARPLocationInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&ARPLocationInformationItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ARPID = new(ARPID)
	err = x.ARPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ARPID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ARPLocationType = new(ARPLocationType)
	err = x.ARPLocationType.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ARPLocationType error")
	}

	// optional field (optPresentFlag index: 0)
	if ARPLocationInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerARPLocationInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
