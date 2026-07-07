package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AMFTNLAssociationSetupItem struct {
	AMFTNLAssociationAddress *CPTransportLayerInformation                                // valueLB:0,valueUB:1
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs // optional
}

func (x *AMFTNLAssociationSetupItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AMFTNLAssociationSetupItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AMFTNLAssociationAddress == nil {
		return errors.Errorf("AMFTNLAssociationAddress is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AMFTNLAssociationSetupItemOptPresentFlag = append(AMFTNLAssociationSetupItemOptPresentFlag, true)
	} else {
		AMFTNLAssociationSetupItemOptPresentFlag = append(AMFTNLAssociationSetupItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AMFTNLAssociationSetupItemOptPresentFlag, true)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.AMFTNLAssociationAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "AMFTNLAssociationAddress marshal failed")
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

func (x *AMFTNLAssociationSetupItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AMFTNLAssociationSetupItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AMFTNLAssociationSetupItemOptPresentFlag, true)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.AMFTNLAssociationAddress = new(CPTransportLayerInformation)
	err = x.AMFTNLAssociationAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode AMFTNLAssociationAddress error")
	}

	// optional field (optPresentFlag index: 0)
	if AMFTNLAssociationSetupItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAMFTNLAssociationSetupItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
