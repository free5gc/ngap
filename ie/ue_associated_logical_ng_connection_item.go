package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type UEAssociatedLogicalNGConnectionItem struct {
	AMFUENGAPID  *AMFUENGAPID                                                         // optional
	RANUENGAPID  *RANUENGAPID                                                         // optional
	IEExtensions *ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs // optional
}

func (x *UEAssociatedLogicalNGConnectionItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	UEAssociatedLogicalNGConnectionItemOptPresentFlag := []bool{}
	// optional field
	if x.AMFUENGAPID != nil {
		UEAssociatedLogicalNGConnectionItemOptPresentFlag = append(UEAssociatedLogicalNGConnectionItemOptPresentFlag, true)
	} else {
		UEAssociatedLogicalNGConnectionItemOptPresentFlag = append(UEAssociatedLogicalNGConnectionItemOptPresentFlag, false)
	}
	// optional field
	if x.RANUENGAPID != nil {
		UEAssociatedLogicalNGConnectionItemOptPresentFlag = append(UEAssociatedLogicalNGConnectionItemOptPresentFlag, true)
	} else {
		UEAssociatedLogicalNGConnectionItemOptPresentFlag = append(UEAssociatedLogicalNGConnectionItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		UEAssociatedLogicalNGConnectionItemOptPresentFlag = append(UEAssociatedLogicalNGConnectionItemOptPresentFlag, true)
	} else {
		UEAssociatedLogicalNGConnectionItemOptPresentFlag = append(UEAssociatedLogicalNGConnectionItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(UEAssociatedLogicalNGConnectionItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.AMFUENGAPID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.AMFUENGAPID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "AMFUENGAPID marshal failed")
		}
	}

	// optional field
	if x.RANUENGAPID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.RANUENGAPID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "RANUENGAPID marshal failed")
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

func (x *UEAssociatedLogicalNGConnectionItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	UEAssociatedLogicalNGConnectionItemOptPresentFlag := make([]bool, 3)
	err = pd.ReadSequencePreambleBitMap(&UEAssociatedLogicalNGConnectionItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if UEAssociatedLogicalNGConnectionItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.AMFUENGAPID = new(AMFUENGAPID)
		err = x.AMFUENGAPID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode AMFUENGAPID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if UEAssociatedLogicalNGConnectionItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.RANUENGAPID = new(RANUENGAPID)
		err = x.RANUENGAPID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode RANUENGAPID error")
		}
	}

	// optional field (optPresentFlag index: 2)
	if UEAssociatedLogicalNGConnectionItemOptPresentFlag[2] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
