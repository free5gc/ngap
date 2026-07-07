package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MBSSessionTNLInfo5GCItem struct {
	MBSAreaSessionID                 *MBSAreaSessionID
	SharedNGUMulticastTNLInformation *SharedNGUMulticastTNLInformation                         // valueExt
	IEExtensions                     *ProtocolExtensionContainerMBSSessionTNLInfo5GCItemExtIEs // optional
}

func (x *MBSSessionTNLInfo5GCItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	MBSSessionTNLInfo5GCItemOptPresentFlag := []bool{}
	// mandatory field
	if x.MBSAreaSessionID == nil {
		return errors.Errorf("MBSAreaSessionID is missing")
	}
	// mandatory field
	if x.SharedNGUMulticastTNLInformation == nil {
		return errors.Errorf("SharedNGUMulticastTNLInformation is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		MBSSessionTNLInfo5GCItemOptPresentFlag = append(MBSSessionTNLInfo5GCItemOptPresentFlag, true)
	} else {
		MBSSessionTNLInfo5GCItemOptPresentFlag = append(MBSSessionTNLInfo5GCItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(MBSSessionTNLInfo5GCItemOptPresentFlag, true)
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
	err = x.SharedNGUMulticastTNLInformation.Write(pd)
	if err != nil {
		return errors.Wrap(err, "SharedNGUMulticastTNLInformation marshal failed")
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

func (x *MBSSessionTNLInfo5GCItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	MBSSessionTNLInfo5GCItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&MBSSessionTNLInfo5GCItemOptPresentFlag, true)
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
	x.SharedNGUMulticastTNLInformation = new(SharedNGUMulticastTNLInformation)
	err = x.SharedNGUMulticastTNLInformation.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode SharedNGUMulticastTNLInformation error")
	}

	// optional field (optPresentFlag index: 0)
	if MBSSessionTNLInfo5GCItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerMBSSessionTNLInfo5GCItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
