package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AMFTNLAssociationToRemoveItem struct {
	AMFTNLAssociationAddress *CPTransportLayerInformation                                   // valueLB:0,valueUB:1
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs // optional
}

func (x *AMFTNLAssociationToRemoveItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AMFTNLAssociationToRemoveItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AMFTNLAssociationAddress == nil {
		return errors.Errorf("AMFTNLAssociationAddress is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AMFTNLAssociationToRemoveItemOptPresentFlag = append(AMFTNLAssociationToRemoveItemOptPresentFlag, true)
	} else {
		AMFTNLAssociationToRemoveItemOptPresentFlag = append(AMFTNLAssociationToRemoveItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AMFTNLAssociationToRemoveItemOptPresentFlag, true)
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

func (x *AMFTNLAssociationToRemoveItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AMFTNLAssociationToRemoveItemOptPresentFlag := make([]bool, 1)
	err = pd.ReadSequencePreambleBitMap(&AMFTNLAssociationToRemoveItemOptPresentFlag, true)
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
	if AMFTNLAssociationToRemoveItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAMFTNLAssociationToRemoveItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
