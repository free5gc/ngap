package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSSessionTNLInfoNGRANItem struct {
	MBSAreaSessionID               *MBSAreaSessionID
	SharedNGUUnicastTNLInformation *UPTransportLayerInformation                                // valueLB:0,valueUB:1,optional
	IEExtensions                   *ProtocolExtensionContainerMBSSessionTNLInfoNGRANItemExtIEs // optional
}

func (x *MBSSessionTNLInfoNGRANItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSSessionTNLInfoNGRANItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSAreaSessionID == nil {
		return errors.Errorf("MBSAreaSessionID is missing")
	}
	// optional field
	if x.SharedNGUUnicastTNLInformation != nil {
		MBSSessionTNLInfoNGRANItemOptPresentFlag = append(MBSSessionTNLInfoNGRANItemOptPresentFlag, true)
	} else {
		MBSSessionTNLInfoNGRANItemOptPresentFlag = append(MBSSessionTNLInfoNGRANItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSSessionTNLInfoNGRANItemOptPresentFlag = append(MBSSessionTNLInfoNGRANItemOptPresentFlag, true)
	} else {
		MBSSessionTNLInfoNGRANItemOptPresentFlag = append(MBSSessionTNLInfoNGRANItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSSessionTNLInfoNGRANItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSAreaSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSAreaSessionID marshal failed")
	}

	// optional field
	if x.SharedNGUUnicastTNLInformation != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.SharedNGUUnicastTNLInformation.Write(pd)
		if err != nil {
			return errors.Wrap(err, "SharedNGUUnicastTNLInformation marshal failed")
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

func (x *MBSSessionTNLInfoNGRANItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSSessionTNLInfoNGRANItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MBSSessionTNLInfoNGRANItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSAreaSessionID = new(MBSAreaSessionID)
	err = x.MBSAreaSessionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSAreaSessionID error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSSessionTNLInfoNGRANItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.SharedNGUUnicastTNLInformation = new(UPTransportLayerInformation)
		err = x.SharedNGUUnicastTNLInformation.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode SharedNGUUnicastTNLInformation error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSSessionTNLInfoNGRANItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSSessionTNLInfoNGRANItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
