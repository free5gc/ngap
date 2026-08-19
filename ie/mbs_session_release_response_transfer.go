package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSSessionReleaseResponseTransfer struct {
	MBSSessionTNLInfoNGRAN *MBSSessionTNLInfoNGRAN                                            // valueLB:0,valueUB:2,optional
	IEExtensions           *ProtocolExtensionContainerMBSSessionReleaseResponseTransferExtIEs // optional
}

func (x *MBSSessionReleaseResponseTransfer) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSSessionReleaseResponseTransferOptPresentFlag := []bool{}
	// optional field
	if x.MBSSessionTNLInfoNGRAN != nil {
		MBSSessionReleaseResponseTransferOptPresentFlag = append(MBSSessionReleaseResponseTransferOptPresentFlag, true)
	} else {
		MBSSessionReleaseResponseTransferOptPresentFlag = append(MBSSessionReleaseResponseTransferOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSSessionReleaseResponseTransferOptPresentFlag = append(MBSSessionReleaseResponseTransferOptPresentFlag, true)
	} else {
		MBSSessionReleaseResponseTransferOptPresentFlag = append(MBSSessionReleaseResponseTransferOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSSessionReleaseResponseTransferOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// optional field
	if x.MBSSessionTNLInfoNGRAN != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSSessionTNLInfoNGRAN.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSSessionTNLInfoNGRAN marshal failed")
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

func (x *MBSSessionReleaseResponseTransfer) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSSessionReleaseResponseTransferOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MBSSessionReleaseResponseTransferOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// optional field (optPresentFlag index: 0)
	if MBSSessionReleaseResponseTransferOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MBSSessionTNLInfoNGRAN = new(MBSSessionTNLInfoNGRAN)
		err = x.MBSSessionTNLInfoNGRAN.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSSessionTNLInfoNGRAN error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSSessionReleaseResponseTransferOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSSessionReleaseResponseTransferExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
