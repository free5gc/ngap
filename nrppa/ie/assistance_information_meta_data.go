package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

const ( /* Enum Type */
	AssistanceInformationMetaDataEncryptedPresentTrue aper.Enumerated = 0
)

const ( /* Enum Type */
	AssistanceInformationMetaDataGNSSIDPresentGps     aper.Enumerated = 0
	AssistanceInformationMetaDataGNSSIDPresentSbas    aper.Enumerated = 1
	AssistanceInformationMetaDataGNSSIDPresentQzss    aper.Enumerated = 2
	AssistanceInformationMetaDataGNSSIDPresentGalileo aper.Enumerated = 3
	AssistanceInformationMetaDataGNSSIDPresentGlonass aper.Enumerated = 4
	AssistanceInformationMetaDataGNSSIDPresentBds     aper.Enumerated = 5
	AssistanceInformationMetaDataGNSSIDPresentNavic   aper.Enumerated = 6
)

const ( /* Enum Type */
	AssistanceInformationMetaDataSBASIDPresentWaas  aper.Enumerated = 0
	AssistanceInformationMetaDataSBASIDPresentEgnos aper.Enumerated = 1
	AssistanceInformationMetaDataSBASIDPresentMsas  aper.Enumerated = 2
	AssistanceInformationMetaDataSBASIDPresentGagan aper.Enumerated = 3
)

type AssistanceInformationMetaData struct {
	Encrypted    *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:0,optional
	GNSSID       *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:6,optional
	SBASID       *aper.Enumerated                                               // valueExt,valueLB:0,valueUB:3,optional
	IEExtensions *ProtocolExtensionContainerAssistanceInformationMetaDataExtIEs // optional
}

func (x *AssistanceInformationMetaData) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AssistanceInformationMetaDataOptPresentFlag := []bool{}
	// optional field
	if x.Encrypted != nil {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, true)
	} else {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, false)
	}
	// optional field
	if x.GNSSID != nil {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, true)
	} else {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, false)
	}
	// optional field
	if x.SBASID != nil {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, true)
	} else {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, true)
	} else {
		AssistanceInformationMetaDataOptPresentFlag = append(AssistanceInformationMetaDataOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AssistanceInformationMetaDataOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.Encrypted != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		err = pd.WriteEnumerated(*(x.Encrypted), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.GNSSID != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 6
		err = pd.WriteEnumerated(*(x.GNSSID), true, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "enumerated marshal failed")
		}
	}

	// optional field
	if x.SBASID != nil {
		// Write Enumerated (Pointer)
		*vLb, *vUb = 0, 3
		err = pd.WriteEnumerated(*(x.SBASID), true, vLb, vUb)
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

func (x *AssistanceInformationMetaData) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AssistanceInformationMetaDataOptPresentFlag := make([]bool, 4)
	err = pd.ReadSequencePreambleBitMap(&AssistanceInformationMetaDataOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if AssistanceInformationMetaDataOptPresentFlag[0] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 0
		x.Encrypted = new(aper.Enumerated)
		*(x.Encrypted), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if AssistanceInformationMetaDataOptPresentFlag[1] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 6
		x.GNSSID = new(aper.Enumerated)
		*(x.GNSSID), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 2)
	if AssistanceInformationMetaDataOptPresentFlag[2] {
		// Read Enumerated (Pointer)
		*vLb, *vUb = 0, 3
		x.SBASID = new(aper.Enumerated)
		*(x.SBASID), err = pd.ReadEnumerated(true, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
		}
	}

	// optional field (optPresentFlag index: 3)
	if AssistanceInformationMetaDataOptPresentFlag[3] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAssistanceInformationMetaDataExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
