package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPPRSInformationListItem struct {
	TRPID            *TRPID
	NRPCI            *NRPCI
	CGINR            *CGINR                                                     // valueExt,optional
	PRSConfiguration *PRSConfiguration                                          // valueExt
	IEExtensions     *ProtocolExtensionContainerTRPPRSInformationListItemExtIEs // optional
}

func (x *TRPPRSInformationListItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	TRPPRSInformationListItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TRPID == nil {
		return errors.Errorf("TRPID is missing")
	}
	// mandatory field
	if x.NRPCI == nil {
		return errors.Errorf("NRPCI is missing")
	}
	// optional field
	if x.CGINR != nil {
		TRPPRSInformationListItemOptPresentFlag = append(TRPPRSInformationListItemOptPresentFlag, true)
	} else {
		TRPPRSInformationListItemOptPresentFlag = append(TRPPRSInformationListItemOptPresentFlag, false)
	}
	// mandatory field
	if x.PRSConfiguration == nil {
		return errors.Errorf("PRSConfiguration is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		TRPPRSInformationListItemOptPresentFlag = append(TRPPRSInformationListItemOptPresentFlag, true)
	} else {
		TRPPRSInformationListItemOptPresentFlag = append(TRPPRSInformationListItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(TRPPRSInformationListItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TRPID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TRPID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.NRPCI.Write(pd)
	if err != nil {
		return errors.Wrap(err, "NRPCI marshal failed")
	}

	// optional field
	if x.CGINR != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.CGINR.Write(pd)
		if err != nil {
			return errors.Wrap(err, "CGINR marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.PRSConfiguration.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSConfiguration marshal failed")
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

func (x *TRPPRSInformationListItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	TRPPRSInformationListItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&TRPPRSInformationListItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TRPID = new(TRPID)
	err = x.TRPID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TRPID error")
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.NRPCI = new(NRPCI)
	err = x.NRPCI.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode NRPCI error")
	}

	// optional field (optPresentFlag index: 0)
	if TRPPRSInformationListItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.CGINR = new(CGINR)
		err = x.CGINR.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode CGINR error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSConfiguration = new(PRSConfiguration)
	err = x.PRSConfiguration.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSConfiguration error")
	}

	// optional field (optPresentFlag index: 1)
	if TRPPRSInformationListItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerTRPPRSInformationListItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
