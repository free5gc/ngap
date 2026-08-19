package ie

import (
	"github.com/free5gc/ngap/aper"
	"github.com/pkg/errors"
)

type DLPRSMutingPattern struct {
	Choice DLPRSMutingPatternAlt
}

type DLPRSMutingPatternAlt interface {
	DLPRSMutingPatternAltIndex() int64
	Write(*aper.PerBitData) error
	Read(*aper.PerBitData) error
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type TwoForDLPRSMutingPattern struct {
	Value aper.BitString
}

func (alt0 *TwoForDLPRSMutingPattern) DLPRSMutingPatternAltIndex() int64 {
	return int64(0)
}

func (x *TwoForDLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 2, 2
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *TwoForDLPRSMutingPattern) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 2, 2
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type FourForDLPRSMutingPattern struct {
	Value aper.BitString
}

func (alt1 *FourForDLPRSMutingPattern) DLPRSMutingPatternAltIndex() int64 {
	return int64(1)
}

func (x *FourForDLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 4, 4
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *FourForDLPRSMutingPattern) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 4, 4
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type SixForDLPRSMutingPattern struct {
	Value aper.BitString
}

func (alt2 *SixForDLPRSMutingPattern) DLPRSMutingPatternAltIndex() int64 {
	return int64(2)
}

func (x *SixForDLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 6, 6
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *SixForDLPRSMutingPattern) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 6, 6
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type EightForDLPRSMutingPattern struct {
	Value aper.BitString
}

func (alt3 *EightForDLPRSMutingPattern) DLPRSMutingPatternAltIndex() int64 {
	return int64(3)
}

func (x *EightForDLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 8, 8
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *EightForDLPRSMutingPattern) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 8, 8
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type SixteenForDLPRSMutingPattern struct {
	Value aper.BitString
}

func (alt4 *SixteenForDLPRSMutingPattern) DLPRSMutingPatternAltIndex() int64 {
	return int64(4)
}

func (x *SixteenForDLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 16, 16
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *SixteenForDLPRSMutingPattern) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 16, 16
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice is an aper-defined/built-in type, complete interface implementation is required
type ThirtyTwoForDLPRSMutingPattern struct {
	Value aper.BitString
}

func (alt5 *ThirtyTwoForDLPRSMutingPattern) DLPRSMutingPatternAltIndex() int64 {
	return int64(5)
}

func (x *ThirtyTwoForDLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Write BitString
	*sLb, *sUb = 32, 32
	err = pd.WriteBitString(x.Value, false, sLb, sUb)
	if err != nil {
		return errors.Wrap(err, "bitString marshal failed")
	}
	return nil
}

func (x *ThirtyTwoForDLPRSMutingPattern) Read(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	// Read BitString
	*sLb, *sUb = 32, 32
	x.Value, err = pd.ReadBitString(false, sLb, sUb)
	if err != nil {
		return BuildTransferSyntaxErr(errors.Wrap(err, "asn.1 decode bitString error"))
	}
	return nil
}

// Choice type and its Read/Write is defined elsewhere
func (alt6 ProtocolIESingleContainerDLPRSMutingPatternExtIEs) DLPRSMutingPatternAltIndex() int64 {
	return int64(6)
}

// Choice Type Read/Write Functions

func (x *DLPRSMutingPattern) Write(pd *aper.PerBitData) error {
	var err error
	var sLb, sUb *uint64 = new(uint64), new(uint64)
	var vLb, vUb *int64 = new(int64), new(int64)

	// dummy function to avoid unused error
	foo(err, sLb, sUb, vLb, vUb)

	var choiceUb int64 = 6
	var option_idx int64 = x.Choice.DLPRSMutingPatternAltIndex()
	err = pd.WriteChoicePreambleBitMap(option_idx, false, &choiceUb)
	if err != nil {
		return errors.Wrap(err, "choice marshal failed")
	}

	// Write Choice
	err = x.Choice.Write(pd)
	return err
}

func (x *DLPRSMutingPattern) Read(pd *aper.PerBitData) error {
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
		x.Choice = new(ThirtyTwoForDLPRSMutingPattern)
	} else if option_idx == 1 {
		x.Choice = new(ThirtyTwoForDLPRSMutingPattern)
	} else if option_idx == 2 {
		x.Choice = new(ThirtyTwoForDLPRSMutingPattern)
	} else if option_idx == 3 {
		x.Choice = new(ThirtyTwoForDLPRSMutingPattern)
	} else if option_idx == 4 {
		x.Choice = new(ThirtyTwoForDLPRSMutingPattern)
	} else if option_idx == 5 {
		x.Choice = new(ThirtyTwoForDLPRSMutingPattern)
	} else if option_idx == 6 {
		x.Choice = new(ProtocolIESingleContainerDLPRSMutingPatternExtIEs)
	} else {
		return errors.Errorf("decoded option index is out of valid choice")
	}

	err = x.Choice.Read(pd)
	return err
}
