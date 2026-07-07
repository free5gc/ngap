package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGAPIESupportInformationRequestItem struct {
	NgapProtocolIEId *ProtocolIEID
	IEExtensions     *ProtocolExtensionContainerNGAPIESupportInformationRequestItemExtIEs // optional
}

func (x *NGAPIESupportInformationRequestItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGAPIESupportInformationRequestItemOptPresentFlag := []bool{}
	// mandatory field
	if x.NgapProtocolIEId == nil {
		return errors.Errorf("NgapProtocolIEId is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		NGAPIESupportInformationRequestItemOptPresentFlag = append(NGAPIESupportInformationRequestItemOptPresentFlag, true)
	} else {
		NGAPIESupportInformationRequestItemOptPresentFlag = append(NGAPIESupportInformationRequestItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGAPIESupportInformationRequestItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.NgapProtocolIEId.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NgapProtocolIEId marshal failed")
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

func (x *NGAPIESupportInformationRequestItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGAPIESupportInformationRequestItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&NGAPIESupportInformationRequestItemOptPresentFlag, true)
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

	// optional field (optPresentFlag index: 0)
	if NGAPIESupportInformationRequestItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGAPIESupportInformationRequestItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
