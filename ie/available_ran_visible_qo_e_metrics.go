package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	AvailableRANVisibleQoEMetricsApplicationLayerBufferLevelListPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	AvailableRANVisibleQoEMetricsPlayoutDelayForMediaStartupPresentTrue aper.Enumerated = 0
)

type AvailableRANVisibleQoEMetrics struct {
	ApplicationLayerBufferLevelList *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:0,optional
	PlayoutDelayForMediaStartup     *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:0,optional
	IEExtensions                    *ProtocolExtensionContainerAvailableRANVisibleQoEMetricsExtIEs // optional
}

func (x *AvailableRANVisibleQoEMetrics) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AvailableRANVisibleQoEMetricsOptPresentFlag := []bool{}
	// optional field
	if x.ApplicationLayerBufferLevelList != nil {
		AvailableRANVisibleQoEMetricsOptPresentFlag = append(AvailableRANVisibleQoEMetricsOptPresentFlag, true)
	} else {
		AvailableRANVisibleQoEMetricsOptPresentFlag = append(AvailableRANVisibleQoEMetricsOptPresentFlag, false)
	}
	// optional field
	if x.PlayoutDelayForMediaStartup != nil {
		AvailableRANVisibleQoEMetricsOptPresentFlag = append(AvailableRANVisibleQoEMetricsOptPresentFlag, true)
	} else {
		AvailableRANVisibleQoEMetricsOptPresentFlag = append(AvailableRANVisibleQoEMetricsOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AvailableRANVisibleQoEMetricsOptPresentFlag = append(AvailableRANVisibleQoEMetricsOptPresentFlag, true)
	} else {
		AvailableRANVisibleQoEMetricsOptPresentFlag = append(AvailableRANVisibleQoEMetricsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AvailableRANVisibleQoEMetricsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.ApplicationLayerBufferLevelList != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.ApplicationLayerBufferLevelList), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.PlayoutDelayForMediaStartup != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.PlayoutDelayForMediaStartup), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
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

func (x *AvailableRANVisibleQoEMetrics) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AvailableRANVisibleQoEMetricsOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&AvailableRANVisibleQoEMetricsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if AvailableRANVisibleQoEMetricsOptPresentFlag[0] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.ApplicationLayerBufferLevelList = new(aper.Enumerated)
		*(x.ApplicationLayerBufferLevelList), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if AvailableRANVisibleQoEMetricsOptPresentFlag[1] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.PlayoutDelayForMediaStartup = new(aper.Enumerated)
		*(x.PlayoutDelayForMediaStartup), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if AvailableRANVisibleQoEMetricsOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAvailableRANVisibleQoEMetricsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
