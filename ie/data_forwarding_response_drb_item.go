package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type DataForwardingResponseDRBItem struct {
	DRBID                        *DRBID
	DLForwardingUPTNLInformation *UPTransportLayerInformation                                   // valueLB:0,valueUB:1,optional
	ULForwardingUPTNLInformation *UPTransportLayerInformation                                   // valueLB:0,valueUB:1,optional
	IEExtensions                 *ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs // optional
}

func (x *DataForwardingResponseDRBItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	DataForwardingResponseDRBItemOptPresentFlag := []bool{}
	// mandatory field
	if x.DRBID == nil {
		return errors.Errorf("DRBID is missing")
	}
	// optional field
	if x.DLForwardingUPTNLInformation != nil {
		DataForwardingResponseDRBItemOptPresentFlag = append(DataForwardingResponseDRBItemOptPresentFlag, true)
	} else {
		DataForwardingResponseDRBItemOptPresentFlag = append(DataForwardingResponseDRBItemOptPresentFlag, false)
	}
	// optional field
	if x.ULForwardingUPTNLInformation != nil {
		DataForwardingResponseDRBItemOptPresentFlag = append(DataForwardingResponseDRBItemOptPresentFlag, true)
	} else {
		DataForwardingResponseDRBItemOptPresentFlag = append(DataForwardingResponseDRBItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		DataForwardingResponseDRBItemOptPresentFlag = append(DataForwardingResponseDRBItemOptPresentFlag, true)
	} else {
		DataForwardingResponseDRBItemOptPresentFlag = append(DataForwardingResponseDRBItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(DataForwardingResponseDRBItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.DRBID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "DRBID marshal failed")
	}

	// optional field
	if x.DLForwardingUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.DLForwardingUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "DLForwardingUPTNLInformation marshal failed")
		}
	}

	// optional field
	if x.ULForwardingUPTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.ULForwardingUPTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "ULForwardingUPTNLInformation marshal failed")
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

func (x *DataForwardingResponseDRBItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	DataForwardingResponseDRBItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&DataForwardingResponseDRBItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.DRBID = new(DRBID)
	err = x.DRBID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode DRBID error")
	}

	// optional field (optPresentFlag index: 0)
	if DataForwardingResponseDRBItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.DLForwardingUPTNLInformation = new(UPTransportLayerInformation)
		err = x.DLForwardingUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode DLForwardingUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if DataForwardingResponseDRBItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.ULForwardingUPTNLInformation = new(UPTransportLayerInformation)
		err = x.ULForwardingUPTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode ULForwardingUPTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if DataForwardingResponseDRBItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerDataForwardingResponseDRBItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
