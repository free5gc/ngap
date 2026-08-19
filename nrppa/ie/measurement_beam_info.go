package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type MeasurementBeamInfo struct {
	PRSResourceID    *PRSResourceID                                       // optional
	PRSResourceSetID *PRSResourceSetID                                    // optional
	SSBIndex         *SSBIndex                                            // optional
	IEExtensions     *ProtocolExtensionContainerMeasurementBeamInfoExtIEs // optional
}

func (x *MeasurementBeamInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MeasurementBeamInfoOptPresentFlag := []bool{}
	// optional field
	if x.PRSResourceID != nil {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, true)
	} else {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, false)
	}
	// optional field
	if x.PRSResourceSetID != nil {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, true)
	} else {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, false)
	}
	// optional field
	if x.SSBIndex != nil {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, true)
	} else {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, true)
	} else {
		MeasurementBeamInfoOptPresentFlag = append(MeasurementBeamInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MeasurementBeamInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.PRSResourceID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSResourceID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSResourceID marshal failed")
		}
	}

	// optional field
	if x.PRSResourceSetID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.PRSResourceSetID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "PRSResourceSetID marshal failed")
		}
	}

	// optional field
	if x.SSBIndex != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SSBIndex.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SSBIndex marshal failed")
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

func (x *MeasurementBeamInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MeasurementBeamInfoOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&MeasurementBeamInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if MeasurementBeamInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.PRSResourceID = new(PRSResourceID)
		err = x.PRSResourceID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSResourceID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MeasurementBeamInfoOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.PRSResourceSetID = new(PRSResourceSetID)
		err = x.PRSResourceSetID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode PRSResourceSetID error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if MeasurementBeamInfoOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.SSBIndex = new(SSBIndex)
		err = x.SSBIndex.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SSBIndex error")
		}
	}

	// optional field (optPresentFlag index: 3)
	if MeasurementBeamInfoOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMeasurementBeamInfoExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
