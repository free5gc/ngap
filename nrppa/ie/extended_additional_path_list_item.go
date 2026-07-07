package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type ExtendedAdditionalPathListItem struct {
	RelativeTimeOfPath *RelativePathDelay                                              // valueLB:0,valueUB:6
	PathQuality        *TrpMeasurementQuality                                          // valueLB:0,valueUB:2,optional
	MultipleULAoA      *MultipleULAoA                                                  // valueExt,optional
	PathPower          *ULSRSRSRPP                                                     // valueExt,optional
	IEExtensions       *ProtocolExtensionContainerExtendedAdditionalPathListItemExtIEs // optional
}

func (x *ExtendedAdditionalPathListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	ExtendedAdditionalPathListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.RelativeTimeOfPath == nil {
		return errors.Errorf("RelativeTimeOfPath is missing")
	}
	// optional field
	if x.PathQuality != nil {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, true)
	} else {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, false)
	}
	// optional field
	if x.MultipleULAoA != nil {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, true)
	} else {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, false)
	}
	// optional field
	if x.PathPower != nil {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, true)
	} else {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, true)
	} else {
		ExtendedAdditionalPathListItemOptPresentFlag = append(ExtendedAdditionalPathListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(ExtendedAdditionalPathListItemOptPresentFlag, true)
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
	if x.MultipleULAoA != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MultipleULAoA.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MultipleULAoA marshal failed")
		}
	}

	// optional field
	if x.PathPower != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PathPower.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PathPower marshal failed")
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

func (x *ExtendedAdditionalPathListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	ExtendedAdditionalPathListItemOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&ExtendedAdditionalPathListItemOptPresentFlag, true)
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
	if ExtendedAdditionalPathListItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PathQuality = new(TrpMeasurementQuality)
		err = x.PathQuality.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PathQuality error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if ExtendedAdditionalPathListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.MultipleULAoA = new(MultipleULAoA)
		err = x.MultipleULAoA.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MultipleULAoA error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if ExtendedAdditionalPathListItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.PathPower = new(ULSRSRSRPP)
		err = x.PathPower.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PathPower error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if ExtendedAdditionalPathListItemOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerExtendedAdditionalPathListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
