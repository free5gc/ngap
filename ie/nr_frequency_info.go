package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NRFrequencyInfo struct {
	NrARFCN           *NRARFCN
	FrequencyBandList *NRFrequencyBandList
	IEExtension       *ProtocolExtensionContainerNRFrequencyInfoExtIEs // optional
}

func (x *NRFrequencyInfo) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRFrequencyInfoOptPresentFlag := []bool{}
	// mandatory field
	if x.NrARFCN == nil {
		return errors.Errorf("NrARFCN is missing")
	}
	// mandatory field
	if x.FrequencyBandList == nil {
		return errors.Errorf("FrequencyBandList is missing")
	}
	// optional field
	if x.IEExtension != nil {
		NRFrequencyInfoOptPresentFlag = append(NRFrequencyInfoOptPresentFlag, true)
	} else {
		NRFrequencyInfoOptPresentFlag = append(NRFrequencyInfoOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRFrequencyInfoOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NrARFCN.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NrARFCN marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.FrequencyBandList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "FrequencyBandList marshal failed")
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

func (x *NRFrequencyInfo) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRFrequencyInfoOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NRFrequencyInfoOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NrARFCN = new(NRARFCN)
	err = x.NrARFCN.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NrARFCN error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.FrequencyBandList = new(NRFrequencyBandList)
	err = x.FrequencyBandList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode FrequencyBandList error")
	}

	// optional field (optPresentFlag index: 0)
	if NRFrequencyInfoOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtension = new(ProtocolExtensionContainerNRFrequencyInfoExtIEs)
		err = x.IEExtension.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtension error")
		}
	}

	return nil
}
