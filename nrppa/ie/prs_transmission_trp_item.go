package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSTransmissionTRPItem struct {
	TRPID            *TRPID
	PRSConfiguration *PRSConfiguration                                       // valueExt
	IEExtensions     *ProtocolExtensionContainerPRSTransmissionTRPItemExtIEs // optional
}

func (x *PRSTransmissionTRPItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSTransmissionTRPItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// mandatory field
	if x.PRSConfiguration == nil {
		return errors.Errorf("PRSConfiguration is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSTransmissionTRPItemOptPresentFlag = append(PRSTransmissionTRPItemOptPresentFlag, true)
	} else {
		PRSTransmissionTRPItemOptPresentFlag = append(PRSTransmissionTRPItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSTransmissionTRPItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PRSConfiguration.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSConfiguration marshal failed")
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

func (x *PRSTransmissionTRPItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSTransmissionTRPItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSTransmissionTRPItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPID = new(TRPID)
	err = x.TRPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSConfiguration = new(PRSConfiguration)
	err = x.PRSConfiguration.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSConfiguration error")
	}

	// optional field (optPresentFlag index: 0)
	if PRSTransmissionTRPItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSTransmissionTRPItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
