package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type NRPRSBeamInformation struct {
	/* Sequence of = 35, FULL Name = struct NR_PRS_Beam_Information__nR_PRS_Beam_InformationList */
	/* Type Name = NRPRSBeamInformationItem */
	/* Sequence Of Embed */
	NRPRSBeamInformationList []NRPRSBeamInformationItem // valueExt,sizeLB:1,sizeUB:2
	/* Sequence of = 35, FULL Name = struct NR_PRS_Beam_Information__lCS_to_GCS_TranslationList */
	/* Type Name = LCSToGCSTranslationItem */
	/* Sequence Of Embed */
	LCSToGCSTranslationList []LCSToGCSTranslationItem                          // valueExt,sizeLB:1,sizeUB:3
	IEExtensions            *ProtocolExtensionContainerNRPRSBeamInformationIEs // optional
}

func (x *NRPRSBeamInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NRPRSBeamInformationOptPresentFlag := []bool{}
	// mandatory field
	if x.NRPRSBeamInformationList == nil {
		return errors.Errorf("NRPRSBeamInformationList is missing")
	}
	// mandatory field
	if x.LCSToGCSTranslationList == nil {
		return errors.Errorf("LCSToGCSTranslationList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NRPRSBeamInformationOptPresentFlag = append(NRPRSBeamInformationOptPresentFlag, true)
	} else {
		NRPRSBeamInformationOptPresentFlag = append(NRPRSBeamInformationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NRPRSBeamInformationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Sequence Of
	*sLb, *sUb = 1, 2
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.NRPRSBeamInformationList)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.NRPRSBeamInformationList {
		err = element.Write(pd)
		if err != nil {
			return errors.Wrap(err, "seqof marshal failed")
		}
	}

	// Write Sequence Of
	*sLb, *sUb = 1, 3
	err = pd.WriteSequenceOfPreambleBitMap(uint64(len(x.LCSToGCSTranslationList)), false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "seqof marshal failed")
	}
	for _, element := range x.LCSToGCSTranslationList {
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

func (x *NRPRSBeamInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NRPRSBeamInformationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NRPRSBeamInformationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 2
	var numElementsNRPRSBeamInformationList uint64
	numElementsNRPRSBeamInformationList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.NRPRSBeamInformationList = []NRPRSBeamInformationItem{}
	for i := 0; i < int(numElementsNRPRSBeamInformationList); i++ {
		var val NRPRSBeamInformationItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.NRPRSBeamInformationList = append(x.NRPRSBeamInformationList, val)
		}
	}

	// mandatory field
	// Read Sequence Of
	*sLb, *sUb = 1, 3
	var numElementsLCSToGCSTranslationList uint64
	numElementsLCSToGCSTranslationList, err = pd.ReadSequenceOfPreambleBitMap(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode seqof error"))
	}
	x.LCSToGCSTranslationList = []LCSToGCSTranslationItem{}
	for i := 0; i < int(numElementsLCSToGCSTranslationList); i++ {
		var val LCSToGCSTranslationItem
		if err = val.Read(pd); err != nil {
			return errors.Wrap(err, "seqof unmarshal failed")
		} else {
			x.LCSToGCSTranslationList = append(x.LCSToGCSTranslationList, val)
		}
	}

	// optional field (optPresentFlag index: 0)
	if NRPRSBeamInformationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNRPRSBeamInformationIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
