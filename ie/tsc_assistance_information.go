package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TSCAssistanceInformation struct {
	Periodicity      *Periodicity
	BurstArrivalTime *BurstArrivalTime                                         // optional
	IEExtensions     *ProtocolExtensionContainerTSCAssistanceInformationExtIEs // optional
}

func (x *TSCAssistanceInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TSCAssistanceInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.Periodicity == nil {
		return errors.Errorf("Periodicity is missing")
	}
	// optional field
	if x.BurstArrivalTime != nil {
		TSCAssistanceInformationOptPresentFlag = append(TSCAssistanceInformationOptPresentFlag, true)
	} else {
		TSCAssistanceInformationOptPresentFlag = append(TSCAssistanceInformationOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TSCAssistanceInformationOptPresentFlag = append(TSCAssistanceInformationOptPresentFlag, true)
	} else {
		TSCAssistanceInformationOptPresentFlag = append(TSCAssistanceInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TSCAssistanceInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.Periodicity.Write(pd)
	if err != nil {
		return errors.Wrap(err, "Periodicity marshal failed")
	}

	// optional field
	if x.BurstArrivalTime != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.BurstArrivalTime.Write(pd)
		if err != nil {
			return errors.Wrap(err, "BurstArrivalTime marshal failed")
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

func (x *TSCAssistanceInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TSCAssistanceInformationOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TSCAssistanceInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.Periodicity = new(Periodicity)
	err = x.Periodicity.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode Periodicity error")
	}

	// optional field (optPresentFlag index: 0)
	if TSCAssistanceInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.BurstArrivalTime = new(BurstArrivalTime)
		err = x.BurstArrivalTime.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode BurstArrivalTime error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TSCAssistanceInformationOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTSCAssistanceInformationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
