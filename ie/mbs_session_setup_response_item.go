package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSSessionSetupResponseItem struct {
	MBSSessionID     *MBSSessionID                                                // valueExt
	MBSAreaSessionID *MBSAreaSessionID                                            // optional
	IEExtensions     *ProtocolExtensionContainerMBSSessionSetupResponseItemExtIEs // optional
}

func (x *MBSSessionSetupResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSSessionSetupResponseItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSSessionID == nil {
		return errors.Errorf("MBSSessionID is missing")
	}
	// optional field
	if x.MBSAreaSessionID != nil {
		MBSSessionSetupResponseItemOptPresentFlag = append(MBSSessionSetupResponseItemOptPresentFlag, true)
	} else {
		MBSSessionSetupResponseItemOptPresentFlag = append(MBSSessionSetupResponseItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		MBSSessionSetupResponseItemOptPresentFlag = append(MBSSessionSetupResponseItemOptPresentFlag, true)
	} else {
		MBSSessionSetupResponseItemOptPresentFlag = append(MBSSessionSetupResponseItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSSessionSetupResponseItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSSessionID marshal failed")
	}

	// optional field
	if x.MBSAreaSessionID != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.MBSAreaSessionID.Write(pd)
		if err != nil {
			return errors.Wrap(err, "MBSAreaSessionID marshal failed")
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

func (x *MBSSessionSetupResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSSessionSetupResponseItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&MBSSessionSetupResponseItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSSessionID = new(MBSSessionID)
	err = x.MBSSessionID.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSSessionID error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSSessionSetupResponseItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.MBSAreaSessionID = new(MBSAreaSessionID)
		err = x.MBSAreaSessionID.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode MBSAreaSessionID error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if MBSSessionSetupResponseItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSSessionSetupResponseItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
