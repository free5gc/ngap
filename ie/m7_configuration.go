package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type M7Configuration struct {
	M7period     *M7period
	M7LinksToLog *LinksToLog                                      // valueExt,valueLB:0,valueUB:2
	IEExtensions *ProtocolExtensionContainerM7ConfigurationExtIEs // optional
}

func (x *M7Configuration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	M7ConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.M7period == nil {
		return errors.Errorf("M7period is missing")
	}
	// mandatory field
	if x.M7LinksToLog == nil {
		return errors.Errorf("M7LinksToLog is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		M7ConfigurationOptPresentFlag = append(M7ConfigurationOptPresentFlag, true)
	} else {
		M7ConfigurationOptPresentFlag = append(M7ConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(M7ConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.M7period.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M7period marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.M7LinksToLog.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M7LinksToLog marshal failed")
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

func (x *M7Configuration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	M7ConfigurationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&M7ConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M7period = new(M7period)
	err = x.M7period.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M7period error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M7LinksToLog = new(LinksToLog)
	err = x.M7LinksToLog.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M7LinksToLog error")
	}

	// optional field (optPresentFlag index: 0)
	if M7ConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerM7ConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
