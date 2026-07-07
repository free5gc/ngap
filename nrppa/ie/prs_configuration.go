package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type PRSConfiguration struct {
	PRSResourceSetList *PRSResourceSetList
	IEExtensions       *ProtocolExtensionContainerPRSConfigurationExtIEs // optional
}

func (x *PRSConfiguration) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	PRSConfigurationOptPresentFlag := []bool{}
	// mandatory field
	if x.PRSResourceSetList == nil {
		return errors.Errorf("PRSResourceSetList is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		PRSConfigurationOptPresentFlag = append(PRSConfigurationOptPresentFlag, true)
	} else {
		PRSConfigurationOptPresentFlag = append(PRSConfigurationOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(PRSConfigurationOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.PRSResourceSetList.Write(pd)
	if err != nil {
		return errors.Wrap(err, "PRSResourceSetList marshal failed")
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

func (x *PRSConfiguration) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	PRSConfigurationOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&PRSConfigurationOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.PRSResourceSetList = new(PRSResourceSetList)
	err = x.PRSResourceSetList.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode PRSResourceSetList error")
	}

	// optional field (optPresentFlag index: 0)
	if PRSConfigurationOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerPRSConfigurationExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
