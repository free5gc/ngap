package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type M4Configuration struct {
	M4period     *M4period                                        // valueExt,valueLB:0,valueUB:4
	M4LinksToLog *LinksToLog                                      // valueExt,valueLB:0,valueUB:2
	IEExtensions *ProtocolExtensionContainerM4ConfigurationExtIEs // optional
}

func (x *M4Configuration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	M4ConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.M4period == nil {
		return errors.Errorf("M4period is missing")
	}
	// mandatory field
	if x.M4LinksToLog == nil {
		return errors.Errorf("M4LinksToLog is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		M4ConfigurationOptPresentFlag = append(M4ConfigurationOptPresentFlag, true)
	} else {
		M4ConfigurationOptPresentFlag = append(M4ConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(M4ConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.M4period.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M4period marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.M4LinksToLog.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M4LinksToLog marshal failed")
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

func (x *M4Configuration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	M4ConfigurationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&M4ConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M4period = new(M4period)
	err = x.M4period.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M4period error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M4LinksToLog = new(LinksToLog)
	err = x.M4LinksToLog.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M4LinksToLog error")
	}

	// optional field (optPresentFlag index: 0)
	if M4ConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerM4ConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
