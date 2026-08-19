package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSFrequencyHoppingConfigurationEUTRA struct {
	NoOfFreqHoppingBands *NumberOfFrequencyHoppingBands // valueExt,valueLB:0,valueUB:1
	/* Sequence of = 35, FULL Name = struct PRSFrequencyHoppingConfiguration_EUTRA__bandPositions */
	/* Type Name = NarrowBandIndex */
	/* Sequence Of Embed */
	BandPositions []NarrowBandIndex                                                       // sizeLB:1,sizeUB:7
	IEExtensions  *ProtocolExtensionContainerPRSFrequencyHoppingConfigurationEUTRAItemIEs // optional
}

func (x *PRSFrequencyHoppingConfigurationEUTRA) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag := []bool{}
	// mandatory field
	if x.NoOfFreqHoppingBands == nil {
		return errors.Errorf("NoOfFreqHoppingBands is missing")
	}
	// mandatory field
	if x.BandPositions == nil {
		return errors.Errorf("BandPositions is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag = append(PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag, true)
	} else {
		PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag = append(PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NoOfFreqHoppingBands.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NoOfFreqHoppingBands marshal failed")
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 7
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.BandPositions)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.BandPositions {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
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

func (x *PRSFrequencyHoppingConfigurationEUTRA) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NoOfFreqHoppingBands = new(NumberOfFrequencyHoppingBands)
	err = x.NoOfFreqHoppingBands.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NoOfFreqHoppingBands error")
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 7
	var numElementsBandPositions uint64
	numElementsBandPositions, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.BandPositions = []NarrowBandIndex{}
	for i := 0; i < int(numElementsBandPositions); i++ {
		var val NarrowBandIndex
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.BandPositions = append(x.BandPositions, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if PRSFrequencyHoppingConfigurationEUTRAOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSFrequencyHoppingConfigurationEUTRAItemIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
