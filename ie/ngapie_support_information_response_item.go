package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

const ( /* Enum Type */
	NGAPIESupportInformationResponseItemNgapProtocolIESupportInfoPresentSupported    aper.Enumerated = 0
	NGAPIESupportInformationResponseItemNgapProtocolIESupportInfoPresentNotSupported aper.Enumerated = 1
)

const ( /* Enum Type */
	NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfoPresentPresent    aper.Enumerated = 0
	NGAPIESupportInformationResponseItemNgapProtocolIEPresenceInfoPresentNotPresent aper.Enumerated = 1
)

type NGAPIESupportInformationResponseItem struct {
	NgapProtocolIEId           *ProtocolIEID
	NgapProtocolIESupportInfo  *aper.Enumerated                                                      // valueExt,valueLB:0,valueUB:1
	NgapProtocolIEPresenceInfo *aper.Enumerated                                                      // valueExt,valueLB:0,valueUB:1
	IEExtensions               *ProtocolExtensionContainerNGAPIESupportInformationResponseItemExtIEs // optional
}

func (x *NGAPIESupportInformationResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGAPIESupportInformationResponseItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NgapProtocolIEId == nil {
		return errors.Errorf("NgapProtocolIEId is missing")
	}
	// mandatory field
	if x.NgapProtocolIESupportInfo == nil {
		return errors.Errorf("NgapProtocolIESupportInfo is missing")
	}
	// mandatory field
	if x.NgapProtocolIEPresenceInfo == nil {
		return errors.Errorf("NgapProtocolIEPresenceInfo is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGAPIESupportInformationResponseItemOptPresentFlag = append(NGAPIESupportInformationResponseItemOptPresentFlag, true)
	} else {
		NGAPIESupportInformationResponseItemOptPresentFlag = append(NGAPIESupportInformationResponseItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGAPIESupportInformationResponseItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NgapProtocolIEId.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NgapProtocolIEId marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.NgapProtocolIESupportInfo), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
	}

	// Write Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	err = pd.WriteEnumerated(*(x.NgapProtocolIEPresenceInfo), true, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "enumerated marshal failed")
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

func (x *NGAPIESupportInformationResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGAPIESupportInformationResponseItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGAPIESupportInformationResponseItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NgapProtocolIEId = new(ProtocolIEID)
	err = x.NgapProtocolIEId.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NgapProtocolIEId error")
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.NgapProtocolIESupportInfo = new(aper.Enumerated)
	*(x.NgapProtocolIESupportInfo), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// mandatory field
	// Read Enumerated (Pointer)
	*vLb, *vUb = 0, 1
	x.NgapProtocolIEPresenceInfo = new(aper.Enumerated)
	*(x.NgapProtocolIEPresenceInfo), err = pd.ReadEnumerated(true, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode enumerated error"))
	}

	// optional field (optPresentFlag index: 0)
	if NGAPIESupportInformationResponseItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGAPIESupportInformationResponseItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
