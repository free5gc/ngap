package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NRFrequencyBandItem struct {
	NrFrequencyBand *NRFrequencyBand
	IEExtension     *ProtocolExtensionContainerNRFrequencyBandItemExtIEs // optional
}

func (x *NRFrequencyBandItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRFrequencyBandItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NrFrequencyBand == nil {
		return errors.Errorf("NrFrequencyBand is missing")
	}
	// optional field
	if x.IEExtension != nil {
		NRFrequencyBandItemOptPresentFlag = append(NRFrequencyBandItemOptPresentFlag, true)
	} else {
		NRFrequencyBandItemOptPresentFlag = append(NRFrequencyBandItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRFrequencyBandItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NrFrequencyBand.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NrFrequencyBand marshal failed")
	}

	// optional field
	if x.IEExtension != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.IEExtension.Write(pd)
		if err != nil {
			return errors.Wrap(err, "IEExtension marshal failed")
		}
	}

	return nil
}

func (x *NRFrequencyBandItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRFrequencyBandItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NRFrequencyBandItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NrFrequencyBand = new(NRFrequencyBand)
	err = x.NrFrequencyBand.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NrFrequencyBand error")
	}

	// optional field (optPresentFlag index: 0)
	if NRFrequencyBandItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerNRFrequencyBandItemExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
