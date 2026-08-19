package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type AdditionalPathListItem struct {
	RelativeTimeOfPath *RelativePathDelay                                      // valueLB:0,valueUB:6
	PathQuality        *TrpMeasurementQuality                                  // valueLB:0,valueUB:2,optional
	IEExtensions       *ProtocolExtensionContainerAdditionalPathListItemExtIEs // optional
}

func (x *AdditionalPathListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AdditionalPathListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.RelativeTimeOfPath == nil {
		return errors.Errorf("RelativeTimeOfPath is missing")
	}
	// optional field
	if x.PathQuality != nil {
		AdditionalPathListItemOptPresentFlag = append(AdditionalPathListItemOptPresentFlag, true)
	} else {
		AdditionalPathListItemOptPresentFlag = append(AdditionalPathListItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AdditionalPathListItemOptPresentFlag = append(AdditionalPathListItemOptPresentFlag, true)
	} else {
		AdditionalPathListItemOptPresentFlag = append(AdditionalPathListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AdditionalPathListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.RelativeTimeOfPath.Write(pd)
	if err != nil {
		return errors.Wrap(err, "RelativeTimeOfPath marshal failed")
	}

	// optional field
	if x.PathQuality != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PathQuality.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PathQuality marshal failed")
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

func (x *AdditionalPathListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AdditionalPathListItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&AdditionalPathListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.RelativeTimeOfPath = new(RelativePathDelay)
	err = x.RelativeTimeOfPath.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode RelativeTimeOfPath error")
	}

	// optional field (optPresentFlag index: 0)
	if AdditionalPathListItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PathQuality = new(TrpMeasurementQuality)
		err = x.PathQuality.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PathQuality error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if AdditionalPathListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAdditionalPathListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
