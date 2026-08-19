package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UENGAPIDPair struct {
	AMFUENGAPID  *AMFUENGAPID
	RANUENGAPID  *RANUENGAPID
	IEExtensions *ProtocolExtensionContainerUENGAPIDPairExtIEs // optional
}

func (x *UENGAPIDPair) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UENGAPIDPairOptPresentFlag := []bool{}
	// mandatory field
	if x.AMFUENGAPID == nil {
		return errors.Errorf("AMFUENGAPID is missing")
	}
	// mandatory field
	if x.RANUENGAPID == nil {
		return errors.Errorf("RANUENGAPID is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		UENGAPIDPairOptPresentFlag = append(UENGAPIDPairOptPresentFlag, true)
	} else {
		UENGAPIDPairOptPresentFlag = append(UENGAPIDPairOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UENGAPIDPairOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AMFUENGAPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AMFUENGAPID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.RANUENGAPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RANUENGAPID marshal failed")
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

func (x *UENGAPIDPair) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UENGAPIDPairOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&UENGAPIDPairOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AMFUENGAPID = new(AMFUENGAPID)
	err = x.AMFUENGAPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AMFUENGAPID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RANUENGAPID = new(RANUENGAPID)
	err = x.RANUENGAPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RANUENGAPID error")
	}

	// optional field (optPresentFlag index: 0)
	if UENGAPIDPairOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUENGAPIDPairExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
