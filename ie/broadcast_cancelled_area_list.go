package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

// error occurs here if the type doesn't implement ProtocolIE interface correctly
var _ ProtocolIE = &BroadcastCancelledAreaList{}

type BroadcastCancelledAreaList struct {
	Choice BroadcastCancelledAreaListAlt
}

type BroadcastCancelledAreaListAlt interface {
	BroadcastCancelledAreaListAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice type and its Read/Write is defined elsewhere
func (alt0 CellIDCancelledEUTRA) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(0)
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 TAICancelledEUTRA) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(1)
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 EmergencyAreaIDCancelledEUTRA) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(2)
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 CellIDCancelledNR) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 TAICancelledNR) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(4)
}

// Choice type and its Read/Write is defined elsewhere
func (alt5 EmergencyAreaIDCancelledNR) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(5)
}

// Choice type and its Read/Write is defined elsewhere
func (alt6 ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs) BroadcastCancelledAreaListAltIndex() int64 {
	return int64(6)
}

// Choice Type Read/Write Functions

func (x *BroadcastCancelledAreaList) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 6
	var option_idx int64 = x.Choice.BroadcastCancelledAreaListAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *BroadcastCancelledAreaList) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 6
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(CellIDCancelledEUTRA)
	} else if option_idx == 1 {
		x.Choice = new(TAICancelledEUTRA)
	} else if option_idx == 2 {
		x.Choice = new(EmergencyAreaIDCancelledEUTRA)
	} else if option_idx == 3 {
		x.Choice = new(CellIDCancelledNR)
	} else if option_idx == 4 {
		x.Choice = new(TAICancelledNR)
	} else if option_idx == 5 {
		x.Choice = new(EmergencyAreaIDCancelledNR)
	} else if option_idx == 6 {
		x.Choice = new(ProtocolIESingleContainerBroadcastCancelledAreaListExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}

func (x *BroadcastCancelledAreaList) WriteIE(pd *aper.PerBitData, id ProtocolIEID, criticality ProtocolIECriticality) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.WriteSequencePreambleBitMap(optPresentFlag, false)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: id
	err = id.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: criticality
	err = criticality.Write(pd)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	// sequence element: value (open type)
	pdOpenType := aper.NewPerBitData(nil)
	err = x.Write(pdOpenType)
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}
	err = pd.WriteOpenType(pdOpenType.Bytes())
	if err != nil {
		return errors.Wrap(err, "write IE failed")
	}

	return nil
}

func (x *BroadcastCancelledAreaList) ReadIE(pd *aper.PerBitData) error {
	// Protocol IE (ProtocolIE-Field in 38.413) is a SEQUENCE type
	optPresentFlag := []bool{} // no optional field
	err := pd.ReadSequencePreambleBitMap(&optPresentFlag, false)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	// sequence element: id is read in message level

	// sequence element: criticality
	protocolIECriticality := ProtocolIECriticality{}
	err = protocolIECriticality.Read(pd)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	// sequence element: value (open type)
	var pdOpenTypeBytes []byte
	pdOpenTypeBytes, err = pd.ReadOpenType()
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}
	pdOpenType := aper.NewPerBitData(pdOpenTypeBytes)
	err = x.Read(pdOpenType)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode IE error"))
	}

	return nil
}
