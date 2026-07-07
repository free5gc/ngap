package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type ngapIEType interface {
	Write(pd *aper.PerBitData) error
	Read(pd *aper.PerBitData) error
}

func MarshalBinary(unmarshalled ngapIEType) ([]byte, error) {
	pd := aper.NewPerBitData(nil)
	err := unmarshalled.Write(pd)
	return pd.Bytes(), errors.Wrap(err, "IE MarshalBinary failed")
}

func UnmarshalBinary(marshalled []byte, target_ie ngapIEType) error {
	pd := aper.NewPerBitData(marshalled)
	err := target_ie.Read(pd)
	return errors.Wrap(err, "IE UnmarshalBinary failed")
}

func UnmarshalUnknownIE(pd *aper.PerBitData) (*ProtocolIECriticality, error) {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	// no optional field -> no need to do ReadSequencePreambleBitMap
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err := protocolIECriticality.Read(pd)
	if err != nil {
		return nil, BuildTransferSyntaxErr(errors.Wrap(err, "UnmarshalAndSkipUnknownIE()"))
	}

	// sequence element: value (open type)
	_, err = pd.ReadOpenType()
	if err != nil {
		return nil, BuildTransferSyntaxErr(errors.Wrap(err, "UnmarshalAndSkipUnknownIE()"))
	}
	return &protocolIECriticality, nil
}
