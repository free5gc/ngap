package ie

// Need to import "fmt" if it uses "fmt"
import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

func WriteOctetStringIE(pd *aper.PerBitData, id ProtocolIEID,
	criticality ProtocolIECriticality, octetString aper.OctetString) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	OptPresentFlag := []bool{} // no optional field
	err := pd.WriteSequencePreambleBitMap(OptPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "IE marshal failed")
	}
	// sequence element: id
	err = id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IE marshal failed")
	}
	// sequence element: criticality
	err = criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "IE marshal failed")
	}
	// sequence element: value (open type)
	pdOpenType := aper.NewPerBitData(nil)
	err = pdOpenType.WriteOctetString(octetString, false, nil, nil)
	if err != nil {
		return errors.Wrap(err, "IE marshal failed")
	}
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "IE marshal failed")
	}

	return nil
}

func ReadOctetStringIE(pd *aper.PerBitData, octetString *aper.OctetString) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	OptPresentFlag := []bool{} // no optional field
	err := pd.ReadSequencePreambleBitMap(&OptPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode ProtocolIE-Field (sequence) error"))
	}
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err = protocolIECriticality.Read(pd)
	if err != nil {
		return errors.Wrap(err, "IE unmarshal failed")
	}

	// sequence element: value (open type)
	var pdOpenTypeBytes []byte
	pdOpenTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode ProtocolIE-Field Value (opentype) error"))
	}
	pdOpenType := aper.NewPerBitData(pdOpenTypeBytes)
	(*octetString), err = pdOpenType.ReadOctetString(false, nil, nil)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err,
			"asn.1 decode ProtocolIE-Field Value opentype (octetstring) error"))
	}

	return nil
}
