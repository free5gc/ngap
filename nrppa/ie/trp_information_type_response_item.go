package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type TRPInformationTypeResponseItem struct {
	Choice TRPInformationTypeResponseItemAlt
}

type TRPInformationTypeResponseItemAlt interface {
	TRPInformationTypeResponseItemAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PCINRForTRPInformationTypeResponseItem struct {
	Value int64
}

func (alt0 *PCINRForTRPInformationTypeResponseItem) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(0)
}

func (x *PCINRForTRPInformationTypeResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 1007
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *PCINRForTRPInformationTypeResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 1007
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt1 CGINR) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(1)
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type ARFCNForTRPInformationTypeResponseItem struct {
	Value int64
}

func (alt2 *ARFCNForTRPInformationTypeResponseItem) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(2)
}

func (x *ARFCNForTRPInformationTypeResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 3279165
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *ARFCNForTRPInformationTypeResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 3279165
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt3 PRSConfiguration) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(3)
}

// Choice type and its Read/Write is defined elsewhere
func (alt4 SSBInfo) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(4)
}

// Choice type and its Read/Write is defined elsewhere
func (alt5 RelativeTime1900) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(5)
}

// Choice type and its Read/Write is defined elsewhere
func (alt6 SpatialDirectionInformation) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(6)
}

// Choice type and its Read/Write is defined elsewhere
func (alt7 GeographicalCoordinates) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(7)
}

// Choice type and its Read/Write is defined elsewhere
func (alt8 ProtocolIESingleContainerTRPInformationTypeResponseItemExtIEs) TRPInformationTypeResponseItemAltIndex() int64 {
	return int64(8)
}

// Choice Type Read/Write Functions

func (x *TRPInformationTypeResponseItem) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 8
	var option_idx int64 = x.Choice.TRPInformationTypeResponseItemAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *TRPInformationTypeResponseItem) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 8
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(ARFCNForTRPInformationTypeResponseItem)
	} else if option_idx == 1 {
		x.Choice = new(CGINR)
	} else if option_idx == 2 {
		x.Choice = new(ARFCNForTRPInformationTypeResponseItem)
	} else if option_idx == 3 {
		x.Choice = new(PRSConfiguration)
	} else if option_idx == 4 {
		x.Choice = new(SSBInfo)
	} else if option_idx == 5 {
		x.Choice = new(RelativeTime1900)
	} else if option_idx == 6 {
		x.Choice = new(SpatialDirectionInformation)
	} else if option_idx == 7 {
		x.Choice = new(GeographicalCoordinates)
	} else if option_idx == 8 {
		x.Choice = new(ProtocolIESingleContainerTRPInformationTypeResponseItemExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
