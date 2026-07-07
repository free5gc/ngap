package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type TSCTrafficCharacteristics struct {
	TSCAssistanceInformationDL *TSCAssistanceInformation                                  // valueExt,optional
	TSCAssistanceInformationUL *TSCAssistanceInformation                                  // valueExt,optional
	IEExtensions               *ProtocolExtensionContainerTSCTrafficCharacteristicsExtIEs // optional
}

func (x *TSCTrafficCharacteristics) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TSCTrafficCharacteristicsOptPresentFlag := []bool{}
	// optional field
	if x.TSCAssistanceInformationDL != nil {
		TSCTrafficCharacteristicsOptPresentFlag = append(TSCTrafficCharacteristicsOptPresentFlag, true)
	} else {
		TSCTrafficCharacteristicsOptPresentFlag = append(TSCTrafficCharacteristicsOptPresentFlag, false)
	}
	// optional field
	if x.TSCAssistanceInformationUL != nil {
		TSCTrafficCharacteristicsOptPresentFlag = append(TSCTrafficCharacteristicsOptPresentFlag, true)
	} else {
		TSCTrafficCharacteristicsOptPresentFlag = append(TSCTrafficCharacteristicsOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		TSCTrafficCharacteristicsOptPresentFlag = append(TSCTrafficCharacteristicsOptPresentFlag, true)
	} else {
		TSCTrafficCharacteristicsOptPresentFlag = append(TSCTrafficCharacteristicsOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TSCTrafficCharacteristicsOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.TSCAssistanceInformationDL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TSCAssistanceInformationDL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TSCAssistanceInformationDL marshal failed")
		}
	}

	// optional field
	if x.TSCAssistanceInformationUL != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TSCAssistanceInformationUL.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TSCAssistanceInformationUL marshal failed")
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

func (x *TSCTrafficCharacteristics) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TSCTrafficCharacteristicsOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&TSCTrafficCharacteristicsOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if TSCTrafficCharacteristicsOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TSCAssistanceInformationDL = new(TSCAssistanceInformation)
		err = x.TSCAssistanceInformationDL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TSCAssistanceInformationDL error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if TSCTrafficCharacteristicsOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.TSCAssistanceInformationUL = new(TSCAssistanceInformation)
		err = x.TSCAssistanceInformationUL.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TSCAssistanceInformationUL error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if TSCTrafficCharacteristicsOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTSCTrafficCharacteristicsExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
