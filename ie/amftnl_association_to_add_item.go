package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type AMFTNLAssociationToAddItem struct {
	AMFTNLAssociationAddress *CPTransportLayerInformation // valueLB:0,valueUB:1
	TNLAssociationUsage      *TNLAssociationUsage         // valueExt,valueLB:0,valueUB:2,optional
	TNLAddressWeightFactor   *TNLAddressWeightFactor
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs // optional
}

func (x *AMFTNLAssociationToAddItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Check mandatory field or write optPresentFlag for optional field
	AMFTNLAssociationToAddItemOptPresentFlag := []bool{}
	// mandatory field
	if x.AMFTNLAssociationAddress == nil {
		return errors.Errorf("AMFTNLAssociationAddress is missing")
	}
	// optional field
	if x.TNLAssociationUsage != nil {
		AMFTNLAssociationToAddItemOptPresentFlag = append(AMFTNLAssociationToAddItemOptPresentFlag, true)
	} else {
		AMFTNLAssociationToAddItemOptPresentFlag = append(AMFTNLAssociationToAddItemOptPresentFlag, false)
	}
	// mandatory field
	if x.TNLAddressWeightFactor == nil {
		return errors.Errorf("TNLAddressWeightFactor is missing")
	}
	// optional field
	if x.IEExtensions != nil {
		AMFTNLAssociationToAddItemOptPresentFlag = append(AMFTNLAssociationToAddItemOptPresentFlag, true)
	} else {
		AMFTNLAssociationToAddItemOptPresentFlag = append(AMFTNLAssociationToAddItemOptPresentFlag, false)
	}

	err = pd.WriteSequencePreambleBitMap(AMFTNLAssociationToAddItemOptPresentFlag, true)
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
	if x.TNLAssociationUsage != nil {
		// Write struct defined elsewhere (Pointer)
		err = x.TNLAssociationUsage.Write(pd)
		if err != nil {
			return errors.Wrap(err, "TNLAssociationUsage marshal failed")
		}
	}

	// Write struct defined elsewhere (Pointer)
	err = x.TNLAddressWeightFactor.Write(pd)
	if err != nil {
		return errors.Wrap(err, "TNLAddressWeightFactor marshal failed")
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

func (x *AMFTNLAssociationToAddItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read optPresentFlag
	AMFTNLAssociationToAddItemOptPresentFlag := make([]bool, 2)
	err = pd.ReadSequencePreambleBitMap(&AMFTNLAssociationToAddItemOptPresentFlag, true)
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
	if AMFTNLAssociationToAddItemOptPresentFlag[0] {
		// Read struct defined elsewhere (Pointer)
		x.TNLAssociationUsage = new(TNLAssociationUsage)
		err = x.TNLAssociationUsage.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode TNLAssociationUsage error")
		}
	}

	// mandatory field
	// Read struct defined elsewhere (Pointer)
	x.TNLAddressWeightFactor = new(TNLAddressWeightFactor)
	err = x.TNLAddressWeightFactor.Read(pd)
	if err != nil {
		return errors.Wrap(err, "decode TNLAddressWeightFactor error")
	}

	// optional field (optPresentFlag index: 1)
	if AMFTNLAssociationToAddItemOptPresentFlag[1] {
		// Read struct defined elsewhere (Pointer)
		x.IEExtensions = new(ProtocolExtensionContainerAMFTNLAssociationToAddItemExtIEs)
		err = x.IEExtensions.Read(pd)
		if err != nil {
			return errors.Wrap(err, "decode IEExtensions error")
		}
	}

	return nil
}
