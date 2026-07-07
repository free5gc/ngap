package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSServiceAreaInformationItem struct {
	MBSAreaSessionID          *MBSAreaSessionID
	MBSServiceAreaInformation *MBSServiceAreaInformation                                     // valueExt
	IEExtensions              *ProtocolExtensionContainerMBSServiceAreaInformationItemExtIEs // optional
}

func (x *MBSServiceAreaInformationItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSServiceAreaInformationItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSAreaSessionID == nil {
		return errors.Errorf("MBSAreaSessionID is missing")
	}
	// mandatory field
	if x.MBSServiceAreaInformation == nil {
		return errors.Errorf("MBSServiceAreaInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		MBSServiceAreaInformationItemOptPresentFlag = append(MBSServiceAreaInformationItemOptPresentFlag, true)
	} else {
		MBSServiceAreaInformationItemOptPresentFlag = append(MBSServiceAreaInformationItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSServiceAreaInformationItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.MBSAreaSessionID.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSAreaSessionID marshal failed")
	}

	// Write struct defined elsewhere (Pointer)
	err = x.MBSServiceAreaInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "MBSServiceAreaInformation marshal failed")
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

func (x *MBSServiceAreaInformationItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSServiceAreaInformationItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&MBSServiceAreaInformationItemOptPresentFlag, true)
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

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.MBSServiceAreaInformation = new(MBSServiceAreaInformation)
	err = x.MBSServiceAreaInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode MBSServiceAreaInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSServiceAreaInformationItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSServiceAreaInformationItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
