package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type OverloadStartNSSAIItem struct {
	SliceOverloadList                   *SliceOverloadList
	SliceOverloadResponse               *OverloadResponse                                       // valueLB:0,valueUB:1,optional
	SliceTrafficLoadReductionIndication *TrafficLoadReductionIndication                         // optional
	IEExtensions                        *ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs // optional
}

func (x *OverloadStartNSSAIItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	OverloadStartNSSAIItemOptPresentFlag := []bool{}
	// mandatory field
	if x.SliceOverloadList == nil {
		return errors.Errorf("SliceOverloadList is missing")
	}
	// optional field
	if x.SliceOverloadResponse != nil {
		OverloadStartNSSAIItemOptPresentFlag = append(OverloadStartNSSAIItemOptPresentFlag, true)
	} else {
		OverloadStartNSSAIItemOptPresentFlag = append(OverloadStartNSSAIItemOptPresentFlag, false)
	}
	// optional field
	if x.SliceTrafficLoadReductionIndication != nil {
		OverloadStartNSSAIItemOptPresentFlag = append(OverloadStartNSSAIItemOptPresentFlag, true)
	} else {
		OverloadStartNSSAIItemOptPresentFlag = append(OverloadStartNSSAIItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		OverloadStartNSSAIItemOptPresentFlag = append(OverloadStartNSSAIItemOptPresentFlag, true)
	} else {
		OverloadStartNSSAIItemOptPresentFlag = append(OverloadStartNSSAIItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(OverloadStartNSSAIItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.SliceOverloadList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SliceOverloadList marshal failed")
	}

	// optional field
	if x.SliceOverloadResponse != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SliceOverloadResponse.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SliceOverloadResponse marshal failed")
		}
	}

	// optional field
	if x.SliceTrafficLoadReductionIndication != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SliceTrafficLoadReductionIndication.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SliceTrafficLoadReductionIndication marshal failed")
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

func (x *OverloadStartNSSAIItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	OverloadStartNSSAIItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&OverloadStartNSSAIItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.SliceOverloadList = new(SliceOverloadList)
	err = x.SliceOverloadList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SliceOverloadList error")
	}

	// optional field (optPresentFlag index: 0)
	if OverloadStartNSSAIItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SliceOverloadResponse = new(OverloadResponse)
		err = x.SliceOverloadResponse.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SliceOverloadResponse error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if OverloadStartNSSAIItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.SliceTrafficLoadReductionIndication = new(TrafficLoadReductionIndication)
		err = x.SliceTrafficLoadReductionIndication.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SliceTrafficLoadReductionIndication error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if OverloadStartNSSAIItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerOverloadStartNSSAIItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
