package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type M6Configuration struct {
	M6reportInterval *M6reportInterval                                // valueExt,valueLB:0,valueUB:13
	M6LinksToLog     *LinksToLog                                      // valueExt,valueLB:0,valueUB:2
	IEExtensions     *ProtocolExtensionContainerM6ConfigurationExtIEs // optional
}

func (x *M6Configuration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	M6ConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.M6reportInterval == nil {
		return errors.Errorf("M6reportInterval is missing")
	}
	// mandatory field
	if x.M6LinksToLog == nil {
		return errors.Errorf("M6LinksToLog is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		M6ConfigurationOptPresentFlag = append(M6ConfigurationOptPresentFlag, true)
	} else {
		M6ConfigurationOptPresentFlag = append(M6ConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(M6ConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.M6reportInterval.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M6reportInterval marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.M6LinksToLog.Write(pd)
	if err != nil {
		return errors.Wrap(err, "M6LinksToLog marshal failed")
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

func (x *M6Configuration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	M6ConfigurationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&M6ConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M6reportInterval = new(M6reportInterval)
	err = x.M6reportInterval.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M6reportInterval error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.M6LinksToLog = new(LinksToLog)
	err = x.M6LinksToLog.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode M6LinksToLog error")
	}

	// optional field (optPresentFlag index: 0)
	if M6ConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerM6ConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
