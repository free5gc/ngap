package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type M5Configuration struct {
	M5period     *M5period                                        // valueExt,valueLB:0,valueUB:4
	M5LinksToLog *LinksToLog                                      // valueExt,valueLB:0,valueUB:2
	IEExtensions *ProtocolExtensionContainerM5ConfigurationExtIEs // optional
}

func (x *M5Configuration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	M5ConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.M5period == nil {
		return errors.Errorf("M5period is missing")
	}
	// mandatory field
	if x.M5LinksToLog == nil {
		return errors.Errorf("M5LinksToLog is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		M5ConfigurationOptPresentFlag = append(M5ConfigurationOptPresentFlag, true)
	} else {
		M5ConfigurationOptPresentFlag = append(M5ConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(M5ConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.M5period.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M5period marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.M5LinksToLog.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M5LinksToLog marshal failed")
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

func (x *M5Configuration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	M5ConfigurationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&M5ConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M5period = new(M5period)
	err = x.M5period.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M5period error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M5LinksToLog = new(LinksToLog)
	err = x.M5LinksToLog.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M5LinksToLog error")
	}

	// optional field (optPresentFlag index: 0)
	if M5ConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerM5ConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
