package ie

import (
	"github.com/pkg/errors"

	"github.com/free5gc/ngap/aper"
)

type MRBProgressInformation struct {
	Choice MRBProgressInformationAlt
}

type MRBProgressInformationAlt interface {
	MRBProgressInformationAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PDCPSNLength12ForMRBProgressInformation struct {
	Value int64
}

func (alt0 *PDCPSNLength12ForMRBProgressInformation) MRBProgressInformationAltIndex() int64 {
	return int64(0)
}

func (x *PDCPSNLength12ForMRBProgressInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 4095
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *PDCPSNLength12ForMRBProgressInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 4095
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type PDCPSNLength18ForMRBProgressInformation struct {
	Value int64
}

func (alt1 *PDCPSNLength18ForMRBProgressInformation) MRBProgressInformationAltIndex() int64 {
	return int64(1)
}

func (x *PDCPSNLength18ForMRBProgressInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write Integer
	*vLb, *vUb = 0, 262143
	err = pd.WriteInteger(x.Value, false, vLb, vUb)
	if err != nil {
		return errors.Wrap(err, "integer marshal failed")
	}
	return nil
}

func (x *PDCPSNLength18ForMRBProgressInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read Integer
	*vLb, *vUb = 0, 262143
	x.Value, err = pd.ReadInteger(false, vLb, vUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode integer error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt2 ProtocolIESingleContainerMRBProgressInformationExtIEs) MRBProgressInformationAltIndex() int64 {
	return int64(2)
}

// Choice Type Read/Write Functions

func (x *MRBProgressInformation) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64 = x.Choice.MRBProgressInformationAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *MRBProgressInformation) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 2
	var option_idx int64
	option_idx, err = pd.ReadChoicePreambleBitMap(false, &choiceUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode choice error"))
	}

	// Read Choice
	if option_idx == 0 {
		x.Choice = new(PDCPSNLength18ForMRBProgressInformation)
	} else if option_idx == 1 {
		x.Choice = new(PDCPSNLength18ForMRBProgressInformation)
	} else if option_idx == 2 {
		x.Choice = new(ProtocolIESingleContainerMRBProgressInformationExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
