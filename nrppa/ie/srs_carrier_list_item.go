package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type SRSCarrierListItem struct {
	PointA                    *int64 // valueLB:0,valueUB:3279165
	UplinkChannelBWPerSCSList *UplinkChannelBWPerSCSList
	ActiveULBWP               *ActiveULBWP                                        // valueExt
	PCINR                     *int64                                              // valueLB:0,valueUB:1007,optional
	IEExtensions              *ProtocolExtensionContainerSRSCarrierListItemExtIEs // optional
}

func (x *SRSCarrierListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	SRSCarrierListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.PointA == nil {
		return errors.Errorf("PointA is missing")
	}
	// mandatory field
	if x.UplinkChannelBWPerSCSList == nil {
		return errors.Errorf("UplinkChannelBWPerSCSList is missing")
	}
	// mandatory field
	if x.ActiveULBWP == nil {
		return errors.Errorf("ActiveULBWP is missing")
	}
	// optional field
	if x.PCINR != nil {
		SRSCarrierListItemOptPresentFlag = append(SRSCarrierListItemOptPresentFlag, true)
	} else {
		SRSCarrierListItemOptPresentFlag = append(SRSCarrierListItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		SRSCarrierListItemOptPresentFlag = append(SRSCarrierListItemOptPresentFlag, true)
	} else {
		SRSCarrierListItemOptPresentFlag = append(SRSCarrierListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(SRSCarrierListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(*(x.PointA), false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.UplinkChannelBWPerSCSList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "UplinkChannelBWPerSCSList marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.ActiveULBWP.Write(pd)
	if err != nil {
		return errors.Wrap(err, "ActiveULBWP marshal failed")
	}

	// optional field
	if x.PCINR != nil {
		// Write Integer (Pointer)
		*vLb, *vUb = 0, 1007
		err = pd.WriteInteger(*(x.PCINR), false, vLb, vUb)
		if err != nil {
			return errors.Wrap(err, "integer marshal failed")
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

func (x *SRSCarrierListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	SRSCarrierListItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&SRSCarrierListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read Integer (Pointer)
	*vLb, *vUb = 0, 3279165
	x.PointA = new(int64)
	*(x.PointA), err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.UplinkChannelBWPerSCSList = new(UplinkChannelBWPerSCSList)
	err = x.UplinkChannelBWPerSCSList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode UplinkChannelBWPerSCSList error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.ActiveULBWP = new(ActiveULBWP)
	err = x.ActiveULBWP.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode ActiveULBWP error")
	}

	// optional field (optPresentFlag index: 0)
	if SRSCarrierListItemOptPresentFlag[0] {
		// Read Integer (Pointer)
		*vLb, *vUb = 0, 1007
		x.PCINR = new(int64)
		*(x.PCINR), err = pd.ReadInteger(false, vLb, vUb)
		if err != nil {
			return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
		}
	}

	// optional field (optPresentFlag index: 1)
	if SRSCarrierListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerSRSCarrierListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
