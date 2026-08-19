package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type NGRANTNLAssociationToRemoveItem struct {
	TNLAssociationTransportLayerAddress    *CPTransportLayerInformation                                     // valueLB:0,valueUB:1
	TNLAssociationTransportLayerAddressAMF *CPTransportLayerInformation                                     // valueLB:0,valueUB:1,optional
	IEExtensions                           *ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs // optional
}

func (x *NGRANTNLAssociationToRemoveItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	NGRANTNLAssociationToRemoveItemOptPresentFlag := []bool{}
	// mandatory field
	if x.TNLAssociationTransportLayerAddress == nil {
		return errors.Errorf("TNLAssociationTransportLayerAddress is missing")
	}
	// optional field
	if x.TNLAssociationTransportLayerAddressAMF != nil {
		NGRANTNLAssociationToRemoveItemOptPresentFlag = append(NGRANTNLAssociationToRemoveItemOptPresentFlag, true)
	} else {
		NGRANTNLAssociationToRemoveItemOptPresentFlag = append(NGRANTNLAssociationToRemoveItemOptPresentFlag, false)
	}
	// optional field
	if x.IEExtensions != nil {
		NGRANTNLAssociationToRemoveItemOptPresentFlag = append(NGRANTNLAssociationToRemoveItemOptPresentFlag, true)
	} else {
		NGRANTNLAssociationToRemoveItemOptPresentFlag = append(NGRANTNLAssociationToRemoveItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(NGRANTNLAssociationToRemoveItemOptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "sequence marshal failed")
	}

	// Write sequence elements

	// Write struct defined elsewhere (Pointer)
	err = x.TNLAssociationTransportLayerAddress.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TNLAssociationTransportLayerAddress marshal failed")
	}

	// optional field
	if x.TNLAssociationTransportLayerAddressAMF != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TNLAssociationTransportLayerAddressAMF.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TNLAssociationTransportLayerAddressAMF marshal failed")
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

func (x *NGRANTNLAssociationToRemoveItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	NGRANTNLAssociationToRemoveItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&NGRANTNLAssociationToRemoveItemOptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode sequence error"))
	}

	// Read sequence elements

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TNLAssociationTransportLayerAddress = new(CPTransportLayerInformation)
	err = x.TNLAssociationTransportLayerAddress.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TNLAssociationTransportLayerAddress error")
	}

	// optional field (optPresentFlag index: 0)
	if NGRANTNLAssociationToRemoveItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TNLAssociationTransportLayerAddressAMF = new(CPTransportLayerInformation)
		err = x.TNLAssociationTransportLayerAddressAMF.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TNLAssociationTransportLayerAddressAMF error")
		}
	}

	// optional field (optPresentFlag index: 1)
	if NGRANTNLAssociationToRemoveItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerNGRANTNLAssociationToRemoveItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
